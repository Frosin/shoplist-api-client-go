package main

import (
	"fmt"
	"shoplist/cmd"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Start!")
	cmd.Execute()
}
