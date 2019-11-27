package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"shoplist/store"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Start!")
	port := ":80"
	if len(os.Args) >= 2 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	var router store.Router
	router.DB.Open(false)
	defer router.DB.GormDB.Close()
	router.DB.GormDB = router.DB.GormDB.Debug().Set("gorm:auto_preload", true)

	route := gin.Default()
	route.GET("/getGoods/:shoppingID", router.GetGoods())
	route.GET("/getComingShoppings/:date", router.GetComingShoppings())
	route.GET("/lastShopping", router.LastShopping())
	route.POST("/addItem", router.AddItem())
	route.POST("/addShopping", router.AddShopping())
	route.Run(port)
}
