package store

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//DB basic database type
type DB struct {
	GormDB *gorm.DB
}

// Open mysqlite connection
func (db *DB) Open(dbFileName string, test bool) *DB {
	gdb, err := gorm.Open("sqlite3", "store/db/"+dbFileName)
	db.GormDB = gdb
	if err != nil {
		panic(err)
	}
	return db
}

/*GetGoodsByShoppingID using:
var goods []ShopList
goods = getGoodsByShoppingId(db, 53)

*/
func (db *DB) GetGoodsByShoppingID(id uint64) (goods []ShopList) {
	db.GormDB.Where("list_id = ?", id).Find(&goods)
	return
}

/*AddProductItem using:
err = addProductItem(db, &ShopList{
	ProductName: "Тестовые товары",
	Quantity:    5,
	ListId:      53,
})
*/
func (db *DB) AddProductItem(productItem *ShopList) (err error) {
	err = db.GormDB.Create(productItem).Error
	return
}

/*AddShopping using:
id, err := addShopping(db, &Shopping{
	Date: "2019-08-23",
	Shop: &Shop{
		Name: "Ашан4",
	},
})
*/
func (db *DB) AddShopping(shopping *Shopping) (err error) {
	if db.GormDB.Where("name = ?", shopping.Shop.Name).Find(&shopping.Shop).RecordNotFound() {
		err = db.GormDB.Create(shopping.Shop).Error
		if err != nil {
			return
		}
	}
	shopping.ShopID = shopping.Shop.ID
	shopping.Shop = nil
	err = db.GormDB.Create(shopping).Error
	return
}

/*GetLastShopping using:
var shopping Shopping
shopping = db.getLastShopping()
*/
func (db *DB) GetLastShopping() (shopping Shopping) {
	db.GormDB.Exec("SELECT * FROM 'shopping' ORDER BY rowid DESC LIMIT 1").Find(&shopping)
	//
	log.Println("**** TEST shopping ****", shopping)
	fmt.Println("**** TEST shopping ****", shopping)
	//
	return
}

/*GetComingShoppings using:
var shoppings []Shoppig
shoppings = db.getComingShoppings("2019-09-05")
*/
func (db *DB) GetComingShoppings(date string) (shoppings []Shopping) {
	db.GormDB.Where("date >= ?", date).Limit(5).Find(&shoppings)
	return
}
