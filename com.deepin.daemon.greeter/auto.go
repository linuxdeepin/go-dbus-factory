package greeter

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

type Greeter struct {
	greeter // interface com.deepin.daemon.Greeter
	proxy.Object
}

func NewGreeter(conn *dbus.Conn) *Greeter {
	obj := new(Greeter)
	obj.Object.Init_(conn, "com.deepin.daemon.Greeter", "/com/deepin/daemon/Greeter")
	return obj
}

type greeter struct{}

func (v *greeter) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*greeter) GetInterfaceName_() string {
	return "com.deepin.daemon.Greeter"
}

// method UpdateGreeterQtTheme

func (v *greeter) GoUpdateGreeterQtTheme(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateGreeterQtTheme", flags, ch, fd)
}

func (v *greeter) GoUpdateGreeterQtThemeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UpdateGreeterQtTheme", flags, ch, fd)
}

func (v *greeter) UpdateGreeterQtTheme(flags dbus.Flags, fd dbus.UnixFD) error {
	return (<-v.GoUpdateGreeterQtTheme(flags, make(chan *dbus.Call, 1), fd).Done).Err
}

func (v *greeter) UpdateGreeterQtThemeWithTimeout(timeout time.Duration, flags dbus.Flags, fd dbus.UnixFD) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUpdateGreeterQtThemeWithContext(ctx, flags, make(chan *dbus.Call, 1), fd).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
