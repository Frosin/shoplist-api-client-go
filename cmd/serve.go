package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/store"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	version = "0.1"
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

		var myServer store.Server

		myServer.DB.Open(viper.GetString("DB_FILE_NAME"), false)
		defer myServer.DB.GormDB.Close()
		myServer.DB.GormDB = myServer.DB.GormDB.Debug().Set("gorm:auto_preload", true)

		e := echo.New()
		e.HTTPErrorHandler = errorHandler
		//e.Use(middleware.Logger())

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
	//serveCmd.Flags().StringP("port", "p", "80", "service port")
	serveCmd.PersistentFlags().StringP("port", "p", "80", "service port")

	sentry.Init(sentry.ClientOptions{
		Dsn: "https://70d91cb8123d4b149c225c315849f53c@sentry.io/1840045",
	})
}

func errorHandler(err error, ctx echo.Context) {
	stacktrace := sentry.NewStacktrace()

	event := sentry.Event{
		User:    sentry.User{},
		Request: sentryRequestFromHTTP(ctx.Request()),
		Exception: []sentry.Exception{{
			Type:       fmt.Sprintf("%T", err),
			Value:      err.Error(),
			Stacktrace: stacktrace,
		}},
	}
	_ = sentry.CaptureEvent(&event)

	code := http.StatusInternalServerError

	//
	spew.Dump(ctx)
	//

	if requestError, ok := err.(*openapi3filter.RequestError); ok {
		code = requestError.HTTPStatus()

		// Get original error
		err = requestError.Err
	}

	switch code {
	case http.StatusBadRequest:
		ctx.JSON(code, api.Error400{
			Message: err.Error(),
		})
	case http.StatusNotFound:
		ctx.JSON(code, api.Error404{
			Message: err.Error(),
		})
	case http.StatusMethodNotAllowed:
		errStr := err.Error()
		ctx.JSON(code, api.Error405{
			Message: &errStr,
		})
	case http.StatusInternalServerError:
		ctx.JSON(code, api.Error500{
			Message: err.Error(),
		})
	}
	return

}

func sentryRequestFromHTTP(r *http.Request) sentry.Request {
	proto := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		proto = "https"
	}

	sentryRequest := sentry.Request{
		URL:         proto + "://" + r.Host + r.URL.Path,
		Method:      r.Method,
		QueryString: r.RequestURI,
		Cookies:     r.Header.Get("Cookie"),
		Headers:     map[string]string{},
	}

	for k, v := range r.Header {
		sentryRequest.Headers[k] = strings.Join(v, ",")
	}

	if addr, port, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		sentryRequest.Env = map[string]string{
			"REMOTE_ADDR": addr,
			"REMOTE_PORT": port,
		}
	}

	return sentryRequest
}
