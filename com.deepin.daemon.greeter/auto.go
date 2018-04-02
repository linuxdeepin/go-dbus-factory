package greeter

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

// method GetScaleFactor

func (v *greeter) GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScaleFactor", flags, ch)
}

func (*greeter) StoreGetScaleFactor(call *dbus.Call) (factor float64, err error) {
	err = call.Store(&factor)
	return
}

func (v *greeter) GetScaleFactor(flags dbus.Flags) (factor float64, err error) {
	return v.StoreGetScaleFactor(
		<-v.GoGetScaleFactor(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetScaleFactor

func (v *greeter) GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, factor float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScaleFactor", flags, ch, factor)
}

func (v *greeter) SetScaleFactor(flags dbus.Flags, factor float64) error {
	return (<-v.GoSetScaleFactor(flags, make(chan *dbus.Call, 1), factor).Done).Err
}
