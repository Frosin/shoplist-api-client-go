package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type DB struct {
	u      string
	p      string
	h      string
	dbName string
	gormDB *gorm.DB
}

func (db *DB) InitDBConfig() *DB {
	file, err := os.Open("sett.txt")
	defer file.Close()

	if err != nil {
		panic("Db file error!")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 4 {
		panic("Invalid config db file!")
	}
	db.u = lines[0]
	db.p = lines[1]
	db.h = lines[2]
	db.dbName = lines[3]
	return db
}

func (db *DB) Open() *DB {
	var err error
	config := fmt.Sprintf("%s:%s@%s/%s", db.u, db.p, db.h, db.dbName)
	db.gormDB, err = gorm.Open("mysql", config)
	panicIfError(err)
	return db
}

/*
	var goods []ShopList
	goods = getGoodsByShoppingId(db, 53)

*/
func (db *DB) getGoodsByShoppingId(id uint64) (goods []ShopList) {
	db.gormDB.Where("list_id=?", id).Find(&goods)
	return
}

/*
	err = addProductItem(db, &ShopList{
		ProductName: "Тестовые товары",
		Quantity:    5,
		ListId:      53,
	})
*/
func (db *DB) addProductItem(productItem *ShopList) (err error) {
	err = db.gormDB.Create(productItem).Error
	return
}

/*
	id, err := addShopping(db, &Shopping{
		Date: "2019-08-23",
		Shop: &Shop{
			Name: "Ашан4",
		},
	})
*/
func (db *DB) addShopping(shopping *Shopping) (err error) {
	if db.gormDB.Where("name = ?", shopping.Shop.Name).Find(&shopping.Shop).RecordNotFound() {
		err = db.gormDB.Create(shopping.Shop).Error
		if err != nil {
			return
		}
	}
	shopping.ShopId = shopping.Shop.Id
	shopping.Shop = nil
	err = db.gormDB.Create(shopping).Error
	return
}

/*
	var shopping Shopping
	shopping = db.getLastShopping()
*/
func (db *DB) getLastShopping() (shopping Shopping) {
	db.gormDB.Last(&shopping)
	return
}

/*
	var shoppings []Shoppig
	shoppings = db.getComingShoppings("2019-09-05")

*/
func (db *DB) getComingShoppings(date string) (shoppings []Shopping) {
	db.gormDB.Where("date >= ?", date).Limit(5).Find(&shoppings)
	return
}
