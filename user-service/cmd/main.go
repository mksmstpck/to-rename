package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/mkskstpck/to-rename/user-service/config"
	"github.com/mkskstpck/to-rename/user-service/database"
	"github.com/mkskstpck/to-rename/user-service/handlers"
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

	// postgres connection
	port, err := strconv.Atoi(config.PSQLport)
	if err != nil {
		panic(err)
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PSQLhost,
		port,
		config.PSQLuser,
		config.PSQLpass,
		config.PSQLdb,
	)
	// open database
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Print(err)
	}
	log.Println("Connected!")
	// handle requests
	user := database.NewUserDB(db)
	handlers.NewHandler(c, user).HandleAll()

	<-make(chan int)
}
