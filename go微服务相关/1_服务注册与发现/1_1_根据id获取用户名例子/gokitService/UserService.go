package gokitService

import "errors"

/*
1、service层用于实现业务逻辑
*/

type IUserService interface {
	GetName(uid int) string
	DeleteName(uid int) error
}

type UserService struct {
}

func (userservice UserService) GetName(uid int) string {
	if uid == 100 {
		return "chory"
	}
	return "guest"
}

func (userservice UserService) DeleteName(uid int) error {
	if uid == 100 {
		return errors.New("no permission")
	}
	return nil
}
