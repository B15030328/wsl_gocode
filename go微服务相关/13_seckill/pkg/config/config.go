package conf

import (
	"fmt"
	"net/http"
	"os"
	"seckill/pkg/bootstrap"
	"seckill/pkg/discover"
	"strconv"

	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

const (
	kConfigType = "CONFIG_TYPE"
)

var (
	Logger log.Logger
)

// 初始化viper默认配置
func initDefault() {
	viper.SetDefault(kConfigType, "yaml")
}

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()

	// 从远程获取配置
	if err := LoadRemoteConfig(); err != nil {
		Logger.Log("Fail to load remote config", err)
	}

	// 链路追踪暂未实现
}

// 从远程获取配置
func LoadRemoteConfig() (err error) {
	serviceInstance, err := discover.DiscoveryService(bootstrap.ConfigServerConfig.Id)
	if err != nil {
		return err
	}
	configServer := "http://" + serviceInstance.Host + ":" + strconv.Itoa(serviceInstance.Port)
	confAddr := fmt.Sprintf("%v/%v/%v-%v.%v",
		configServer, bootstrap.ConfigServerConfig.Label,
		bootstrap.DiscoverConfig.ServiceName, bootstrap.ConfigServerConfig.Profile,
		viper.Get(kConfigType))
	resp, err := http.Get(confAddr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	viper.SetConfigType(viper.GetString(kConfigType))
	if err = viper.ReadConfig(resp.Body); err != nil {
		return
	}
	Logger.Log("Load config from: ", confAddr)
	return
}
