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
	daemon // interface com.deepin.userexperience.Daemon
	proxy.Object
}

func NewDaemon(conn *dbus.Conn) *Daemon {
	obj := new(Daemon)
	obj.Object.Init_(conn, "com.deepin.userexperience.Daemon", "/com/deepin/userexperience/Daemon")
	return obj
}

type daemon struct{}

func (v *daemon) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*daemon) GetInterfaceName_() string {
	return "com.deepin.userexperience.Daemon"
}

// method Enable

func (v *daemon) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *daemon) Enable(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method IsEnabled

func (v *daemon) GoIsEnabled(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsEnabled", flags, ch)
}

func (*daemon) StoreIsEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *daemon) IsEnabled(flags dbus.Flags) (enabled bool, err error) {
	return v.StoreIsEnabled(
		<-v.GoIsEnabled(flags, make(chan *dbus.Call, 1)).Done)
}

// method PostMessage

func (v *daemon) GoPostMessage(flags dbus.Flags, ch chan *dbus.Call, msg string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PostMessage", flags, ch, msg)
}

func (v *daemon) PostMessage(flags dbus.Flags, msg string) error {
	return (<-v.GoPostMessage(flags, make(chan *dbus.Call, 1), msg).Done).Err
}
