package main

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/mksmstpck/to-rename/api-gateway/handlers/web"
	"github.com/nats-io/nats.go"
)

func main() {
	//nats connection
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	//starts echo
	e := echo.New()
	handlers := handlers.NewWeb(e, nc)
	handlers.All(e)
	e.Logger.Fatal(e.Start(":1323"))
}
