package define

type Product struct {
	ID           int64  `json:"ID" gorm:"column:ID"`
	ProductName  string `json:"ProductName" gorm:"column:productName"`
	ProductNum   int64  `json:"ProductNum" gorm:"column:productNum"`
	ProductImage string `json:"ProductImage" gorm:"column:productImage"`
	ProductUrl   string `json:"ProductUrl" gorm:"column:productUrl"`
}

func (pro *Product) TableName() string {
	return "product"
}
