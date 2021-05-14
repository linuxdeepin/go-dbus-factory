package sessionwatcher

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

type SessionWatcher struct {
	sessionWatcher // interface com.deepin.daemon.SessionWatcher
	proxy.Object
}

func NewSessionWatcher(conn *dbus.Conn) *SessionWatcher {
	obj := new(SessionWatcher)
	obj.Object.Init_(conn, "com.deepin.daemon.SessionWatcher", "/com/deepin/daemon/SessionWatcher")
	return obj
}

type sessionWatcher struct{}

func (v *sessionWatcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sessionWatcher) GetInterfaceName_() string {
	return "com.deepin.daemon.SessionWatcher"
}

// method GetSessions

func (v *sessionWatcher) GoGetSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSessions", flags, ch)
}

func (v *sessionWatcher) GoGetSessionsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetSessions", flags, ch)
}

func (*sessionWatcher) StoreGetSessions(call *dbus.Call) (sessions []dbus.ObjectPath, err error) {
	err = call.Store(&sessions)
	return
}

func (v *sessionWatcher) GetSessions(flags dbus.Flags) (sessions []dbus.ObjectPath, err error) {
	return v.StoreGetSessions(
		<-v.GoGetSessions(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionWatcher) GetSessionsWithTimeout(timeout time.Duration, flags dbus.Flags) (sessions []dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetSessionsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetSessions(call)
}

// method IsX11SessionActive

func (v *sessionWatcher) GoIsX11SessionActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsX11SessionActive", flags, ch)
}

func (v *sessionWatcher) GoIsX11SessionActiveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsX11SessionActive", flags, ch)
}

func (*sessionWatcher) StoreIsX11SessionActive(call *dbus.Call) (is_active bool, err error) {
	err = call.Store(&is_active)
	return
}

func (v *sessionWatcher) IsX11SessionActive(flags dbus.Flags) (is_active bool, err error) {
	return v.StoreIsX11SessionActive(
		<-v.GoIsX11SessionActive(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionWatcher) IsX11SessionActiveWithTimeout(timeout time.Duration, flags dbus.Flags) (is_active bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsX11SessionActiveWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsX11SessionActive(call)
}

// property IsActive b

func (v *sessionWatcher) IsActive() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsActive",
	}
}
