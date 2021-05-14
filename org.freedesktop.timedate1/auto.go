package timedate1

import "context"
import "errors"
import "fmt"
import dbus "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Timedate struct {
	timedate // interface org.freedesktop.timedate1
	proxy.Object
}

func NewTimedate(conn *dbus.Conn) *Timedate {
	obj := new(Timedate)
	obj.Object.Init_(conn, "org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	return obj
}

type timedate struct{}

func (v *timedate) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*timedate) GetInterfaceName_() string {
	return "org.freedesktop.timedate1"
}

// method SetTime

func (v *timedate) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, arg0 int64, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, arg0, arg1, arg2)
}

func (v *timedate) GoSetTimeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 int64, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTime", flags, ch, arg0, arg1, arg2)
}

func (v *timedate) SetTime(flags dbus.Flags, arg0 int64, arg1 bool, arg2 bool) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

func (v *timedate) SetTimeWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 int64, arg1 bool, arg2 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTimeWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetTimezone

func (v *timedate) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, arg0, arg1)
}

func (v *timedate) GoSetTimezoneWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTimezone", flags, ch, arg0, arg1)
}

func (v *timedate) SetTimezone(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *timedate) SetTimezoneWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTimezoneWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLocalRTC

func (v *timedate) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocalRTC", flags, ch, arg0, arg1, arg2)
}

func (v *timedate) GoSetLocalRTCWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool, arg2 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLocalRTC", flags, ch, arg0, arg1, arg2)
}

func (v *timedate) SetLocalRTC(flags dbus.Flags, arg0 bool, arg1 bool, arg2 bool) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

func (v *timedate) SetLocalRTCWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 bool, arg1 bool, arg2 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLocalRTCWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetNTP

func (v *timedate) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, arg0, arg1)
}

func (v *timedate) GoSetNTPWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetNTP", flags, ch, arg0, arg1)
}

func (v *timedate) SetNTP(flags dbus.Flags, arg0 bool, arg1 bool) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *timedate) SetNTPWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 bool, arg1 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetNTPWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Timezone s

func (v *timedate) Timezone() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Timezone",
	}
}

// property LocalRTC b

func (v *timedate) LocalRTC() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "LocalRTC",
	}
}

// property CanNTP b

func (v *timedate) CanNTP() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanNTP",
	}
}

// property NTP b

func (v *timedate) NTP() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NTP",
	}
}

// property NTPSynchronized b

func (v *timedate) NTPSynchronized() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NTPSynchronized",
	}
}

// property TimeUSec t

func (v *timedate) TimeUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeUSec",
	}
}

// property RTCTimeUSec t

func (v *timedate) RTCTimeUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RTCTimeUSec",
	}
}
