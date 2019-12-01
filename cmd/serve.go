package cmd

import (
	"log"
	"shoplist/store"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start and serve service",
	Long:  "it start service",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		log.Println("Serve launched on port = ", port)
		if err != nil {
			log.Println("Error = ", err)
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
		route.Run(":" + port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "80", "service port")
}
