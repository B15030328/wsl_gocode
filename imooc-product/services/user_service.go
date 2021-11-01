package services

import (
	"errors"
	"imooc-product/models/define"
	"imooc-product/models/mysql"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Ins mysql.UserInstance
}

func NewUserService() *UserService {
	return &UserService{
		Ins: *mysql.NewUserInstance(),
	}
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (user *define.User, isOk bool) {

	user, err := u.Ins.Select(userName)

	if err != nil {
		return
	}
	isOk, _ = ValidatePassword(pwd, user.HashPassword)

	if !isOk {
		return &define.User{}, false
	}

	return
}

func (u *UserService) AddUser(user *define.User) (userId int64, err error) {
	pwdByte, errPwd := GeneratePassword(user.HashPassword)
	if errPwd != nil {
		return userId, errPwd
	}
	user.HashPassword = string(pwdByte)
	return u.Ins.Insert(user)
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil

}
