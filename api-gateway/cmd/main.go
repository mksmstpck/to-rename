package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maxotka/to-rename/api-gateway/handlers"
)

func main() {
	e := echo.New()
	handlers := handlers.NewHandler(e)
	handlers.All(e)
	e.Logger.Fatal(e.Start(":1323"))
}
