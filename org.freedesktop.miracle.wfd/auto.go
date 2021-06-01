package wfd

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-dbus-factory/object_manager"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Wfd struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	wfd                          // interface org.freedesktop.miracle.wfd
	proxy.Object
}

func NewWfd(conn *dbus.Conn) *Wfd {
	obj := new(Wfd)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wfd", "/org/freedesktop/miracle/wfd")
	return obj
}

func (obj *Wfd) Wfd() *wfd {
	return &obj.wfd
}

type wfd struct{}

func (v *wfd) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wfd) GetInterfaceName_() string {
	return "org.freedesktop.miracle.wfd"
}

// method Shutdown

func (v *wfd) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *wfd) GoShutdownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *wfd) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wfd) ShutdownWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShutdownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

type Sink struct {
	sink // interface org.freedesktop.miracle.wfd.Sink
	proxy.Object
}

func NewSink(conn *dbus.Conn, path dbus.ObjectPath) (*Sink, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Sink)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wfd", path)
	return obj, nil
}

type sink struct{}

func (v *sink) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sink) GetInterfaceName_() string {
	return "org.freedesktop.miracle.wfd.Sink"
}

// method StartSession

func (v *sink) GoStartSession(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartSession", flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (v *sink) GoStartSessionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartSession", flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (*sink) StoreStartSession(call *dbus.Call) (arg7 dbus.ObjectPath, err error) {
	err = call.Store(&arg7)
	return
}

func (v *sink) StartSession(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) (arg7 dbus.ObjectPath, err error) {
	return v.StoreStartSession(
		<-v.GoStartSession(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3, arg4, arg5, arg6).Done)
}

func (v *sink) StartSessionWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) (arg7 dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartSessionWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3, arg4, arg5, arg6).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStartSession(call)
}

// property Session o

func (v *sink) Session() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Session",
	}
}

// property Peer o

func (v *sink) Peer() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Peer",
	}
}

type Session struct {
	session // interface org.freedesktop.miracle.wfd.Session
	proxy.Object
}

func NewSession(conn *dbus.Conn, path dbus.ObjectPath) (*Session, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Session)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wfd", path)
	return obj, nil
}

type session struct{}

func (v *session) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*session) GetInterfaceName_() string {
	return "org.freedesktop.miracle.wfd.Session"
}

// method Resume

func (v *session) GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Resume", flags, ch)
}

func (v *session) GoResumeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Resume", flags, ch)
}

func (v *session) Resume(flags dbus.Flags) error {
	return (<-v.GoResume(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *session) ResumeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoResumeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Pause

func (v *session) GoPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pause", flags, ch)
}

func (v *session) GoPauseWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Pause", flags, ch)
}

func (v *session) Pause(flags dbus.Flags) error {
	return (<-v.GoPause(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *session) PauseWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPauseWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Teardown

func (v *session) GoTeardown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Teardown", flags, ch)
}

func (v *session) GoTeardownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Teardown", flags, ch)
}

func (v *session) Teardown(flags dbus.Flags) error {
	return (<-v.GoTeardown(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *session) TeardownWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTeardownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Sink o

func (v *session) Sink() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Sink",
	}
}

// property Url s

func (v *session) Url() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Url",
	}
}

// property State i

func (v *session) State() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "State",
	}
}
