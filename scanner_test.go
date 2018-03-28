package zwave

import (
	"bytes"
	"testing"
)

func TestScanner(t *testing.T) {
	msg := []byte{SOF, byte(4), 0x00, 0x01, 0x00, byte(250)}

	s, messages := NewScanner()
	r := bytes.NewReader(msg[0:2])
	s.Scan(r)

	r = bytes.NewReader(msg[2:3])
	s.Scan(r)

	r = bytes.NewReader(msg[3:]) // the rest
	go s.Scan(r)
	<-messages
	if !s.Valid {
		t.Errorf("Should be valid")
	}

}

func Test_checksum(t *testing.T) {
	data := []struct {
		val []byte
		exp byte
	}{
		{[]byte{0}, byte(255)},
		{[]byte{0, 1}, byte(254)},
		{[]byte{3, 1}, byte(253)},
		{[]byte{byte(4), 0x00, 0x01, 0x00}, byte(250)},
	}

	for _, d := range data {
		sum := checksum(d.val)
		if d.exp != sum {
			t.Errorf("Checksum => %v expected %v", sum, d.exp)
		}
	}
}
