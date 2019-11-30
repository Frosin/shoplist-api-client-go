package main

import (
	"fmt"
	"os"
	"shoplist/cmd"
	"shoplist/store"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://70d91cb8123d4b149c225c315849f53c@sentry.io/1840045",
	})

	fmt.Println("Start!")
	port := ":80"

	cmd.Execute()

	if len(os.Args) >= 2 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	var router store.Router
	router.DB.Open(viper.GetString("DB_FILE_NAME"), false)
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
