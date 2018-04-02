package backlight

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

type Backlight struct {
	backlight // interface com.deepin.daemon.helper.Backlight
	proxy.Object
}

func NewBacklight(conn *dbus.Conn) *Backlight {
	obj := new(Backlight)
	obj.Object.Init_(conn, "com.deepin.daemon.helper.Backlight", "/com/deepin/daemon/helper/Backlight")
	return obj
}

type backlight struct{}

func (v *backlight) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*backlight) GetInterfaceName_() string {
	return "com.deepin.daemon.helper.Backlight"
}

// method SetBrightness

func (v *backlight) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, type0 uint8, name string, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, type0, name, value)
}

func (v *backlight) SetBrightness(flags dbus.Flags, type0 uint8, name string, value int32) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), type0, name, value).Done).Err
}
