package timedated

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

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

func (v *timedated) GoSetLocalRTCWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLocalRTC", flags, ch, enabled, fixSystem, message)
}

func (v *timedated) SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), enabled, fixSystem, message).Done).Err
}

func (v *timedated) SetLocalRTCWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLocalRTCWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled, fixSystem, message).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetNTP

func (v *timedated) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, enabled, message)
}

func (v *timedated) GoSetNTPWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetNTP", flags, ch, enabled, message)
}

func (v *timedated) SetNTP(flags dbus.Flags, enabled bool, message string) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), enabled, message).Done).Err
}

func (v *timedated) SetNTPWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetNTPWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled, message).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetNTPServer

func (v *timedated) GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTPServer", flags, ch, server, message)
}

func (v *timedated) GoSetNTPServerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetNTPServer", flags, ch, server, message)
}

func (v *timedated) SetNTPServer(flags dbus.Flags, server string, message string) error {
	return (<-v.GoSetNTPServer(flags, make(chan *dbus.Call, 1), server, message).Done).Err
}

func (v *timedated) SetNTPServerWithTimeout(timeout time.Duration, flags dbus.Flags, server string, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetNTPServerWithContext(ctx, flags, make(chan *dbus.Call, 1), server, message).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetTime

func (v *timedated) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, usec, relative, message)
}

func (v *timedated) GoSetTimeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTime", flags, ch, usec, relative, message)
}

func (v *timedated) SetTime(flags dbus.Flags, usec int64, relative bool, message string) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), usec, relative, message).Done).Err
}

func (v *timedated) SetTimeWithTimeout(timeout time.Duration, flags dbus.Flags, usec int64, relative bool, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTimeWithContext(ctx, flags, make(chan *dbus.Call, 1), usec, relative, message).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetTimezone

func (v *timedated) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, timezone, message)
}

func (v *timedated) GoSetTimezoneWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTimezone", flags, ch, timezone, message)
}

func (v *timedated) SetTimezone(flags dbus.Flags, timezone string, message string) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), timezone, message).Done).Err
}

func (v *timedated) SetTimezoneWithTimeout(timeout time.Duration, flags dbus.Flags, timezone string, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTimezoneWithContext(ctx, flags, make(chan *dbus.Call, 1), timezone, message).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property NTPServer s

func (v *timedated) NTPServer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "NTPServer",
	}
}
