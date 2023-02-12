package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/config"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	handlers "github.com/mksmstpck/to-rename/api-gateway/handlers/web"
	"github.com/nats-io/nats.go"
)

func main() {
	// config
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	//nats connection
	nc, err := nats.Connect(config.NatsUrl)
	if err != nil {
		panic(err)
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	//starts echo
	e := echo.New()
	pub := events.NewPub(nc)
	handlers := handlers.NewWeb(e, nc, pub)
	handlers.All()
	e.Logger.Fatal(e.Start(config.EchoUrl))
}
