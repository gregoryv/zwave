package zwave

import (
	"fmt"
	"io"
	"log"
)

type Scanner struct {
	buf      []byte // Temporary read buffer
	n        int    // last number of bytes read
	i        int    // Index of byte to process next
	err      error  // last error
	Start    bool
	Full     []byte
	Length   int // Set to the payload length
	Valid    bool
	Checksum byte
	messages chan Message
}

func NewScanner() (s *Scanner, messages chan Message) {
	s = &Scanner{
		buf:      make([]byte, 128),
		Start:    false,
		i:        0,
		Valid:    false,
		Full:     make([]byte, 254),
		messages: make(chan Message),
	}
	return s, s.messages
}

// Scan reads from the reader
// does not handle if starting to read in the middle of a message
func (s *Scanner) Scan(r io.Reader) {
	s.n, s.err = r.Read(s.buf)
	for _, b := range s.buf[:s.n] {
		if !s.Start && b == ACK {
			s.messages <- NewMessage([]byte{ACK})
			// No need to reset here
		}
		if !s.Start && b == NAK {
			log.Print("Got a NAK????")
			// No need to reset here
		}
		if !s.Start && b == CAN {
			log.Print("CAN!!!")
			// No need to reset here
		}

		if !s.Start && b != SOF { // Skip until we have a SOF
			continue
		}
		if !s.Start && b == SOF { // We found a SOF
			s.Full[0] = SOF
			s.Start = true
			continue
		}
		if s.Start && s.Length == 0 { // Length of message, create it
			s.Full[1] = b
			s.Length = int(b)
			s.i++
			continue
		}
		if s.i < s.Length { // Message still incomplete
			s.Full[s.i+1] = b
			s.i++
			continue
		}

		if s.i == s.Length { // Done with entire message, do checksum validation
			s.Checksum = checksum(s.Full[1 : s.Length+1])
			s.Full[s.i] = b
			if b == s.Checksum { // Exclude the checksum
				s.Valid = true
				m := make([]byte, s.Length+1)
				copy(m, s.Full[:s.Length+1])
				s.messages <- NewMessage(m)
			} else {
				fmt.Printf("%#v\n", s.buf)
			}
			s.Reset()
		}
	}
}

func (s *Scanner) Reset() {
	s.Start = false
	s.i = 0
	s.Length = 0
}

// XOR each value in sequence and negate it
func checksum(data []byte) byte {
	sum := byte(0xff)
	for _, b := range data {
		sum ^= b
	}
	return sum
}
