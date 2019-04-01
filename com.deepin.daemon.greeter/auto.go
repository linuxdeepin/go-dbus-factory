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

// method GetScreenScaleFactors

func (v *greeter) GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScreenScaleFactors", flags, ch)
}

func (*greeter) StoreGetScreenScaleFactors(call *dbus.Call) (factors map[string]float64, err error) {
	err = call.Store(&factors)
	return
}

func (v *greeter) GetScreenScaleFactors(flags dbus.Flags) (factors map[string]float64, err error) {
	return v.StoreGetScreenScaleFactors(
		<-v.GoGetScreenScaleFactors(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetScaleFactor

func (v *greeter) GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, factor float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScaleFactor", flags, ch, factor)
}

func (v *greeter) SetScaleFactor(flags dbus.Flags, factor float64) error {
	return (<-v.GoSetScaleFactor(flags, make(chan *dbus.Call, 1), factor).Done).Err
}

// method SetScreenScaleFactors

func (v *greeter) GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, factors map[string]float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScreenScaleFactors", flags, ch, factors)
}

func (v *greeter) SetScreenScaleFactors(flags dbus.Flags, factors map[string]float64) error {
	return (<-v.GoSetScreenScaleFactors(flags, make(chan *dbus.Call, 1), factors).Done).Err
}
