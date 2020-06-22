package main

import (
	"github.com/maitungmn/bookstore_items-api/app"
	"os"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}
