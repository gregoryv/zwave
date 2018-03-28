package zwave

/*
Message Format

Index Bytes Description
0     1     Header
1     1     Message length
2 0   1     Type, eg. Request or Response
3 1   1     Function
4 2   1     EMPTY=0
5 3   1     (Node ID)
6 4   1     Command length(N)   +
7 5   1     Command Class       +  (From 3.3 Command class frame format) 0x20 - 0xEE
8 6   1     Command             +
            Command data        +
            ...                 +
            Command data N      +
      1     Checksum
*/
import (
	"fmt"
)

// Message
type Message []byte

func NewMessage(m []byte) Message {
	return Message(m)
}

func (m Message) Header() byte   { return m[0] }
func (m Message) Length() byte   { return m[1] }
func (m Message) Type() byte     { return m[2] }
func (m Message) Function() byte { return m[3] } // Index 4 is empty
func (m Message) NodeId() byte   { return m[5] }

func (m Message) HeaderName() string   { return NamedHeader.Get(m.Header()) }
func (m Message) TypeName() string     { return NamedMessageType.Get(m.Type()) }
func (m Message) FunctionName() string { return NamedFunction.Get(m.Function()) }

// ToString converts a message excluding command information
func (m Message) String() string {
	str := ""
	l := len(m)
	if l > 0 {
		str = fmt.Sprintf("hdr:%s", m.HeaderName())
	}
	if l > 1 {
		str = fmt.Sprintf("%s len:%d", str, m[1])
	}
	if l > 2 {
		str = fmt.Sprintf("%s type:%q", str, m.TypeName())
	}
	if l > 3 {
		str = fmt.Sprintf("%s fn:%q", str, m.FunctionName())
	}
	if l > 5 {
		str = fmt.Sprintf("%s nid:%d", str, m.NodeId())
	}
	if l > 6 {
		str = fmt.Sprintf("%s sum:%d", str, m[len(m)-1])
	}

	return str
}

// Returns function, command class, command and data for further handling
func (m Message) FCD() (fnc, cls, cmd byte, data []byte) {
	l := len(m)
	if l > 3 {
		fnc = m[3]
	}
	if l > 7 {
		cls = m[7]
	}
	if l > 8 {
		cmd = m[8]
	}
	if l > 9 {
		data = m[9 : len(m)-1]
	}
	return
}

// Command returns the command specified by function class and command constants
func (m Message) Command() interface{} {
	fn, cls, cmd, dat := m.FCD()
	switch fn {
	case FUNCTION_APPLICATION_COMMAND_HANDLER:
		switch cls {
		case COMMAND_CLASS_NOTIFICATION:
			switch cmd {
			case NOTIFICATION_REPORT:
				return NotificationReport(dat)
			}
		case COMMAND_CLASS_SENSOR_MULTILEVEL:
			return SensorMultilevel(dat)
		}
	}
	return UnknownCommand(dat)
}
