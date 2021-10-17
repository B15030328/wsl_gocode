package def

import (
	"github.com/streadway/amqp"
)

const Mqurl = "amqp://guest:guest@localhost:5672/"

type Rabbitmq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	Mqurl string
}

// 建立rabbitmq新连接
func NewRabbitMQ(queuename, exchange, key string) *Rabbitmq {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial(Mqurl)
	failOnError(err, "Failed to connect to RabbitMQ")
	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return &Rabbitmq{
		Conn:      conn,
		Channel:   ch,
		QueueName: queuename,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     Mqurl,
	}
}

// 断开连接
func (r *Rabbitmq) Destory() {
	r.Conn.Close()
	r.Channel.Close()
}
