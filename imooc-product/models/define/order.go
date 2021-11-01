package define

type Order struct {
	ID          int64 `json:"ID" gorm:"column:ID"`
	UserId      int64 `json:"UserId" gorm:"column:userID"`
	ProductId   int64 `json:"ProductId" gorm:"column:productID"`
	OrderStatus int   `json:"OrderStatus" gorm:"column:orderStatus"`
}

const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)

func (o *Order) TableName() string {
	return "order"
}
