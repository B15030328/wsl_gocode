package define

type Message struct {
	ProductID int64
	UserID    int64
}

// 创建结构体
func NewMessage(userid, productid int64) *Message {
	return &Message{ProductID: productid, UserID: userid}
}
