package events

import "github.com/nats-io/nats.go"

type Pub struct {
	conn *nats.Conn
}

func NewPub(conn *nats.Conn) *Pub {
	return &Pub{
		conn: conn,
	}
}
