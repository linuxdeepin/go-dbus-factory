// Code generated by "./generator ./com.deepin.daemon.timedated"; DO NOT EDIT.

package timedated

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Timedated interface {
	timedated // interface com.deepin.daemon.Timedated
	proxy.Object
}

type objectTimedated struct {
	interfaceTimedated // interface com.deepin.daemon.Timedated
	proxy.ImplObject
}

func NewTimedated(conn *dbus.Conn) Timedated {
	obj := new(objectTimedated)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Timedated", "/com/deepin/daemon/Timedated")
	return obj
}

type timedated interface {
	GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call
	SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error
	GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call
	SetNTP(flags dbus.Flags, enabled bool, message string) error
	GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call
	SetNTPServer(flags dbus.Flags, server string, message string) error
	GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call
	SetTime(flags dbus.Flags, usec int64, relative bool, message string) error
	GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call
	SetTimezone(flags dbus.Flags, timezone string, message string) error
	NTPServer() proxy.PropString
	Timezone() proxy.PropString
}

type interfaceTimedated struct{}

func (v *interfaceTimedated) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTimedated) GetInterfaceName_() string {
	return "com.deepin.daemon.Timedated"
}

// method SetLocalRTC

func (v *interfaceTimedated) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocalRTC", flags, ch, enabled, fixSystem, message)
}

func (v *interfaceTimedated) SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), enabled, fixSystem, message).Done).Err
}

// method SetNTP

func (v *interfaceTimedated) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, enabled, message)
}

func (v *interfaceTimedated) SetNTP(flags dbus.Flags, enabled bool, message string) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), enabled, message).Done).Err
}

// method SetNTPServer

func (v *interfaceTimedated) GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTPServer", flags, ch, server, message)
}

func (v *interfaceTimedated) SetNTPServer(flags dbus.Flags, server string, message string) error {
	return (<-v.GoSetNTPServer(flags, make(chan *dbus.Call, 1), server, message).Done).Err
}

// method SetTime

func (v *interfaceTimedated) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, usec, relative, message)
}

func (v *interfaceTimedated) SetTime(flags dbus.Flags, usec int64, relative bool, message string) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), usec, relative, message).Done).Err
}

// method SetTimezone

func (v *interfaceTimedated) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, timezone, message)
}

func (v *interfaceTimedated) SetTimezone(flags dbus.Flags, timezone string, message string) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), timezone, message).Done).Err
}

// property NTPServer s

func (v *interfaceTimedated) NTPServer() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "NTPServer",
	}
}

// property Timezone s

func (v *interfaceTimedated) Timezone() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Timezone",
	}
}
