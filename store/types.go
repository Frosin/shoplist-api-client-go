package store

// ShopList contains products and thei attributes
type ShopList struct {
	ID          uint   `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    uint   `json:"quantity,omitempty"`
	CategoryID  uint   `json:"categoryId,omitempty"`
	Complete    int32  `json:"complete,omitempty"`
	ListID      uint   `json:"listId,omitempty"`
}

// Shop contains shops descriptions
type Shop struct {
	ID   uint   `json:"id"`
	Name string `json:"name,omitempty"`
}

// Shopping contains information of shoppings
type Shopping struct {
	ID       uint   `json:"id"`
	Date     string `json:"date"`
	Sum      int32  `json:"sum"`
	ShopID   uint   `json:"shopId"`
	Complete int32  `json:"complete"`
	Time     string `json:"time"`
	OwnerID  uint   `json:"ownerId"`
	Shop     *Shop  `json:"shop"` //`gorm:"foreignkey:ShopId;association_foreignkey:Id"`
}

// TableName returns tableName need for GORM
func (ShopList) TableName() string {
	return "shop_list"
}

// TableName returns tableName need for GORM
func (Shop) TableName() string {
	return "shop"
}

// TableName returns tableName need for GORM
func (Shopping) TableName() string {
	return "shopping"
}
