package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := newApp()
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}
