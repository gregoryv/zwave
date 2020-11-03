package zwave

import "fmt"

type UnknownCommand []byte

type SensorMultilevel []byte

func (dat SensorMultilevel) String() string {
	return fmt.Sprintf("Implement sensor multilevel: %v", []byte(dat))
}

type NotificationReport []byte

func (dat NotificationReport) String() string {
	return fmt.Sprintf(
		"type:%q level:%d source:%d status:%q ztype:%q event:%q count:%d",
		dat.TypeName(),
		dat.Level(),
		dat.Source(),
		dat.StatusName(),
		dat.ZTypeName(),
		dat.ZEventName(),
		dat.EventParamCount(),
	)
}

func (dat NotificationReport) Type() byte   { return dat[0] } // 4.3.4 Alarm Report Command v1 specific
func (dat NotificationReport) Level() byte  { return dat[1] } //                            v1 specific
func (dat NotificationReport) Source() byte { return dat[2] }
func (dat NotificationReport) Status() byte { return dat[3] }
func (dat NotificationReport) ZType() byte  { return dat[4] } // Z-Wave Alarm Type
func (dat NotificationReport) ZEvent() byte { return dat[5] }
func (dat NotificationReport) EventParamCount() byte {
	if len(dat) >= 7 { // fibaro motion sensor didn't send this one, spec says nothing of it being optional though
		return dat[6]
	} else {
		return 0
	}
}

func (dat NotificationReport) TypeName() string   { return NamedCommandClass.Get(dat.Type()) }
func (dat NotificationReport) StatusName() string { return NamedZAlarmStatus.Get(dat.Status()) }
func (dat NotificationReport) ZTypeName() string  { return NamedZAlarmType.Get(dat.ZType()) }
func (dat NotificationReport) ZEventName() string { return NamedBurglarEvent.Get(dat.ZEvent()) }
