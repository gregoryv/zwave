package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte{0xff, 0x00})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes", n)

}
