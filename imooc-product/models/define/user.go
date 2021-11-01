package define

type User struct {
	ID           int64  `json:"ID" gorm:"column:ID" form:"ID"`
	NickName     string `json:"nickName" gorm:"column:nickName" form:"nickName"`
	UserName     string `json:"userName" gorm:"column:userName" form:"userName"`
	HashPassword string `json:"-" gorm:"column:passWord" form:"passWord"`
}

func (u *User) TableName() string {
	return "user"
}
