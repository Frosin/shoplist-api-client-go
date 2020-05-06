package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/ent"
	"github.com/Frosin/shoplist-api-client-go/ent/migrate"
	"github.com/Frosin/shoplist-api-client-go/store"
	entsql "github.com/facebookincubator/ent/dialect/sql"

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

		dbPath := viper.GetString("SHOPLIST_DB_PATH")
		// create db path if not exist
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			err := os.Mkdir(dbPath, 0777)
			if err != nil {
				log.Info(err)
			}
		}

		dbFullFileName := dbPath + "/" + viper.GetString("SHOPLIST_DB_FILE_NAME")
		db, err := sql.Open("sqlite3", dbFullFileName+"?_fk=1")
		if err != nil {
			log.Fatal(err)
		}

		entClient := ent.NewClient(ent.Driver(entsql.OpenDB("sqlite3", db)))

		//run migration
		if _, err := os.Stat(dbFullFileName); os.IsNotExist(err) {
			runMigration(entClient)
			// change file permissions
			err := os.Chmod(dbFullFileName, 0777)
			if err != nil {
				log.Info(err)
			}
		}

		server := store.NewServer(version, entClient, db)
		//fill fixtures
		//server.FillFixtures()

		e := echo.New()
		api.RegisterHandlers(e, server)
		e.HTTPErrorHandler = errorHandler
		e.Use(server.TokenHandler)

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

func runMigration(entClient *ent.Client) error {
	return entClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	)
}
