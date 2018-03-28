package zwave

import (
	"github.com/tarm/serial"
	"log"
)

type Controller struct {
	device string
	port   *serial.Port
}

func NewController(device string) *Controller {
	return &Controller{device: device}
}

func (con *Controller) Start() chan Message {
	cfg := &serial.Config{Name: con.device, Baud: 115200}
	port, err := serial.OpenPort(cfg)
	if err != nil {
		panic(err)
	}
	con.port = port
	scanner, messages := NewScanner()
	go func() {
		for {
			scanner.Scan(port)
		}
	}()
	con.Init()
	return messages
}

func (con *Controller) Init() {

	nak := Message{NAK}
	con.Send(nak)
	m := Message{SOF, 0, REQUEST, FUNCTION_GET_VERSION} // what about the checksum?
	con.Send(m)
}

func (con *Controller) Send(m Message) {
	n, err := con.port.Write([]byte(m))
	log.Printf("Send: %s", m)
	if err != nil {
		log.Fatal(err)
	} else {
		if n != len(m) {
			log.Print("Haven't written ACK yet")
		}
	}
}
