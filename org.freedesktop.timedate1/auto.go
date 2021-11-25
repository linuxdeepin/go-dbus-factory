// Code generated by "./generator ./org.freedesktop.timedate1"; DO NOT EDIT.

package timedate1

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Timedate interface {
	timedate // interface org.freedesktop.timedate1
	proxy.Object
}

type objectTimedate struct {
	interfaceTimedate // interface org.freedesktop.timedate1
	proxy.ImplObject
}

func NewTimedate(conn *dbus.Conn) Timedate {
	obj := new(objectTimedate)
	obj.ImplObject.Init_(conn, "org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	return obj
}

type timedate interface {
	GoSetTime(flags dbus.Flags, ch chan *dbus.Call, arg0 int64, arg1 bool, arg2 bool) *dbus.Call
	SetTime(flags dbus.Flags, arg0 int64, arg1 bool, arg2 bool) error
	GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetTimezone(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool, arg2 bool) *dbus.Call
	SetLocalRTC(flags dbus.Flags, arg0 bool, arg1 bool, arg2 bool) error
	GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool) *dbus.Call
	SetNTP(flags dbus.Flags, arg0 bool, arg1 bool) error
	Timezone() proxy.PropString
	LocalRTC() proxy.PropBool
	CanNTP() proxy.PropBool
	NTP() proxy.PropBool
	NTPSynchronized() proxy.PropBool
	TimeUSec() proxy.PropUint64
	RTCTimeUSec() proxy.PropUint64
}

type interfaceTimedate struct{}

func (v *interfaceTimedate) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTimedate) GetInterfaceName_() string {
	return "org.freedesktop.timedate1"
}

// method SetTime

func (v *interfaceTimedate) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, arg0 int64, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, arg0, arg1, arg2)
}

func (v *interfaceTimedate) SetTime(flags dbus.Flags, arg0 int64, arg1 bool, arg2 bool) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetTimezone

func (v *interfaceTimedate) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, arg0, arg1)
}

func (v *interfaceTimedate) SetTimezone(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetLocalRTC

func (v *interfaceTimedate) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocalRTC", flags, ch, arg0, arg1, arg2)
}

func (v *interfaceTimedate) SetLocalRTC(flags dbus.Flags, arg0 bool, arg1 bool, arg2 bool) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetNTP

func (v *interfaceTimedate) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, arg0, arg1)
}

func (v *interfaceTimedate) SetNTP(flags dbus.Flags, arg0 bool, arg1 bool) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// property Timezone s

func (v *interfaceTimedate) Timezone() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Timezone",
	}
}

// property LocalRTC b

func (v *interfaceTimedate) LocalRTC() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "LocalRTC",
	}
}

// property CanNTP b

func (v *interfaceTimedate) CanNTP() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanNTP",
	}
}

// property NTP b

func (v *interfaceTimedate) NTP() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "NTP",
	}
}

// property NTPSynchronized b

func (v *interfaceTimedate) NTPSynchronized() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "NTPSynchronized",
	}
}

// property TimeUSec t

func (v *interfaceTimedate) TimeUSec() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "TimeUSec",
	}
}

// property RTCTimeUSec t

func (v *interfaceTimedate) RTCTimeUSec() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "RTCTimeUSec",
	}
}
