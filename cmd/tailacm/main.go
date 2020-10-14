package main

import (
	"fmt"
	"github.com/gregoryv/zwave"
	"log"
)

func main() {
	con := zwave.NewController("/dev/ttyACM0")
	messages := con.Start()
	for {
		m := <-messages
		handle(con, m)
	}
}

func handle(con *zwave.Controller, m zwave.Message) {
	switch m.Header() {
	case zwave.SOF:
		fmt.Print("\n")
		con.Send(zwave.Message{zwave.ACK})

		anycmd := m.Command()
		switch cmd := anycmd.(type) {
		case zwave.NotificationReport:
			log.Printf("%s", cmd)
		case zwave.SensorMultilevel:
			log.Printf("%s", cmd)
		default:
			log.Printf("UNHANDLED: %v %v", m, []byte(m))
		}

	default:
		log.Printf("Got other than SOF")
	}

}
