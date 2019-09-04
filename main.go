package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	var router Router
	router.db.InitDBConfig().Open()
	defer router.db.gormDB.Close()
	router.db.gormDB = router.db.gormDB.Debug().Set("gorm:auto_preload", true)

	route := gin.Default()
	route.GET("/getGoods/:shoppingId", router.getGoods())
	route.GET("/getComingShoppings/:date", router.getComingShoppings())
	route.GET("/lastShopping", router.lastShopping())
	route.POST("/addItem", router.addItem())
	route.POST("/addShopping", router.addShopping())
	route.Run(port)
}
