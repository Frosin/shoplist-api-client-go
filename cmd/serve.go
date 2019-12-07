package cmd

import (
	"log"
	"shoplist/api"
	"shoplist/store"

	"github.com/labstack/echo/v4"
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

		e := echo.New()

		//e.Use(middleware.Logger())
		var myServer store.Server

		api.RegisterHandlers(e, &myServer)

		//
		// echo.GET("/getGoods/:shoppingID", router.GetGoods())
		// echo.GET("/getComingShoppings/:date", router.GetComingShoppings())
		// echo.GET("/lastShopping", router.LastShopping())
		// echo.POST("/addItem", router.AddItem())
		// echo.POST("/addShopping", router.AddShopping())
		e.Logger.Debug(e.Start(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "80", "service port")
}
