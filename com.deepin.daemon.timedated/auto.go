package timedated

import "errors"
import "fmt"
import "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf

type Timedated struct {
	timedated // interface com.deepin.daemon.Timedated
	proxy.Object
}

func NewTimedated(conn *dbus.Conn) *Timedated {
	obj := new(Timedated)
	obj.Object.Init_(conn, "com.deepin.daemon.Timedated", "/com/deepin/daemon/Timedated")
	return obj
}

type timedated struct{}

func (v *timedated) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*timedated) GetInterfaceName_() string {
	return "com.deepin.daemon.Timedated"
}

// method SetLocalRTC

func (v *timedated) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocalRTC", flags, ch, enabled, fixSystem, message)
}

func (v *timedated) SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), enabled, fixSystem, message).Done).Err
}

// method SetNTP

func (v *timedated) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, enabled, message)
}

func (v *timedated) SetNTP(flags dbus.Flags, enabled bool, message string) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), enabled, message).Done).Err
}

// method SetTime

func (v *timedated) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, usec, relative, message)
}

func (v *timedated) SetTime(flags dbus.Flags, usec int64, relative bool, message string) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), usec, relative, message).Done).Err
}

// method SetTimezone

func (v *timedated) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, timezone, message)
}

func (v *timedated) SetTimezone(flags dbus.Flags, timezone string, message string) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), timezone, message).Done).Err
}
