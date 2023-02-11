package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
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
	userEvents := events.NewPub(nc)
	handlers := handlers.NewWeb(e, nc, userEvents)
	handlers.All()
	e.Logger.Fatal(e.Start(":1323"))
}
