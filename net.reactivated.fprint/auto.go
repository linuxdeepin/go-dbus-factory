package fprint

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

type Manager struct {
	manager // interface net.reactivated.Fprint.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "net.reactivated.Fprint", "/net/reactivated/Fprint/Manager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "net.reactivated.Fprint.Manager"
}

// method GetDefaultDevice

func (v *manager) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (*manager) StoreGetDefaultDevice(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *manager) GetDefaultDevice(flags dbus.Flags) (device dbus.ObjectPath, err error) {
	return v.StoreGetDefaultDevice(
		<-v.GoGetDefaultDevice(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevices

func (v *manager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*manager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}
