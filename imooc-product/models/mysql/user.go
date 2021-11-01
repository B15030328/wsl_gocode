package mysql

import (
	"errors"
	"imooc-product/models/define"
	"log"

	"gorm.io/gorm"
)

type UserInstance struct {
	db *gorm.DB
}

func NewUserInstance() *UserInstance {
	return &UserInstance{
		db: defaultDb,
	}
}

// Select(userName string) (user *define.User, err error)
// Insert(user *define.User) (userId int64, err error)

func (u *UserInstance) Select(userName string) (user *define.User, err error) {
	if userName == "" {
		return &define.User{}, errors.New("条件不能为空！")
	}

	err = u.db.Where("userName = ?", userName).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Panic("user Select error,err:", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return &define.User{}, errors.New("用户不存在！")
	}
	return user, err
}

func (u *UserInstance) Insert(user *define.User) (userId int64, err error) {
	err = u.db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}
