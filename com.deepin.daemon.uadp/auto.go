package uadp

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Uadp struct {
	uadp // interface com.deepin.daemon.Uadp
	proxy.Object
}

func NewUadp(conn *dbus.Conn) *Uadp {
	obj := new(Uadp)
	obj.Object.Init_(conn, "com.deepin.daemon.Uadp", "/com/deepin/daemon/Uadp")
	return obj
}

type uadp struct{}

func (v *uadp) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*uadp) GetInterfaceName_() string {
	return "com.deepin.daemon.Uadp"
}

// method SetDataKey

func (v *uadp) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, dataKey string, keyringKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDataKey", flags, ch, exePath, keyName, dataKey, keyringKey)
}

func (v *uadp) SetDataKey(flags dbus.Flags, exePath string, keyName string, dataKey string, keyringKey string) error {
	return (<-v.GoSetDataKey(flags, make(chan *dbus.Call, 1), exePath, keyName, dataKey, keyringKey).Done).Err
}

// method GetDataKey

func (v *uadp) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, keyringKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDataKey", flags, ch, exePath, keyName, keyringKey)
}

func (*uadp) StoreGetDataKey(call *dbus.Call) (dataKey string, err error) {
	err = call.Store(&dataKey)
	return
}

func (v *uadp) GetDataKey(flags dbus.Flags, exePath string, keyName string, keyringKey string) (dataKey string, err error) {
	return v.StoreGetDataKey(
		<-v.GoGetDataKey(flags, make(chan *dbus.Call, 1), exePath, keyName, keyringKey).Done)
}
