package main

import (
	"fmt"
	"shoplist/cmd"

	"github.com/getsentry/sentry-go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	cmd.Execute()
}
