package main

type ShopList struct {
	Id          uint   `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    uint   `json:"quantity,omitempty"`
	CategoryId  uint   `json:"categoryId,omitempty"`
	Complete    int32  `json:"complete,omitempty"`
	ListId      uint   `json:"listId,omitempty"`
}

type Shop struct {
	Id   uint   `json:"id"`
	Name string `json:"name,omitempty"`
}

type Shopping struct {
	Id       uint   `json:"id"`
	Date     string `json:"date"`
	Sum      int32  `json:"sum"`
	ShopId   uint   `json:"shopId"`
	Complete int32  `json:"complete"`
	Time     string `json:"time"`
	OwnerId  uint   `json:"ownerId"`
	Shop     *Shop  `json:"shop"` //`gorm:"foreignkey:ShopId;association_foreignkey:Id"`
}

func (ShopList) TableName() string {
	return "shop_list"
}

func (Shop) TableName() string {
	return "shop"
}

func (Shopping) TableName() string {
	return "shopping"
}
