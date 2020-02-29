package cmd

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/store"
	"github.com/Frosin/shoplist-api-client-go/store/sqlc"
	"github.com/jmoiron/sqlx"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start and serve service",
	Long:  "it start service",
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("SHOPLIST_API_VERSION")
		port, err := cmd.Flags().GetString("port")
		log.Info("Serve command launched on port = ", port)
		if err != nil {
			log.Info("Error = ", err)
		}
		var myServer store.Server
		db, err := sqlx.Open("sqlite3", "store/db/"+viper.GetString("SHOPLIST_DB_FILE_NAME"))
		if err != nil {
			log.Fatal(err)
		}
		myServer.Queries = sqlc.New(db)
		myServer.Version = version
		myServer.DB = db
		e := echo.New()
		e.HTTPErrorHandler = errorHandler
		api.RegisterHandlers(e, &myServer)
		e.Logger.SetLevel(log.INFO)
		e.Logger.Debug(e.Start(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringP("port", "p", "80", "service port")

	dsn := viper.GetString("SHOPLIST_SENTRY_DSN")
	sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	})
}

func errorHandler(err error, ctx echo.Context) {
	log.Info("REQUEST='", ctx.Path(), "', ERROR=", err)
	version := viper.GetString("SHOPLIST_API_VERSION")
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

	if requestError, ok := err.(*openapi3filter.RequestError); ok {
		code = requestError.HTTPStatus()
		// Get original error
		err = requestError.Err
	}

	switch code {
	case http.StatusBadRequest:
		var response api.Error400
		response.Version = &version
		response.Message = err.Error()
		ctx.JSON(code, response)
	case http.StatusNotFound:
		var response api.Error404
		response.Version = &version
		response.Message = err.Error()
		ctx.JSON(code, response)
	case http.StatusMethodNotAllowed:
		var response api.Error405
		response.Version = &version
		errStr := err.Error()
		response.Message = &errStr
		ctx.JSON(code, response)
	case http.StatusInternalServerError:
		var response api.Error500
		response.Version = &version
		response.Message = err.Error()
		ctx.JSON(code, response)
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
