package main

import (
	"golang/configs"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.Init(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
