package daemon

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
var _ unsafe.Pointer

type Daemon struct {
	daemon // interface com.deepin.daemon.Daemon
	proxy.Object
}

func NewDaemon(conn *dbus.Conn) *Daemon {
	obj := new(Daemon)
	obj.Object.Init_(conn, "com.deepin.daemon.Daemon", "/com/deepin/daemon/Daemon")
	return obj
}

type daemon struct{}

func (v *daemon) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*daemon) GetInterfaceName_() string {
	return "com.deepin.daemon.Daemon"
}

// method ScalePlymouth

func (v *daemon) GoScalePlymouth(flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ScalePlymouth", flags, ch, scale)
}

func (v *daemon) ScalePlymouth(flags dbus.Flags, scale uint32) error {
	return (<-v.GoScalePlymouth(flags, make(chan *dbus.Call, 1), scale).Done).Err
}

// method SetLongPressDuration

func (v *daemon) GoSetLongPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongPressDuration", flags, ch, duration)
}

func (v *daemon) SetLongPressDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetLongPressDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}
