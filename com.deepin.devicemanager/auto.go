// Code generated by "./generator com.deepin.devicemanager"; DO NOT EDIT.

package devicemanager

import "github.com/godbus/dbus"

import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

type DeviceManager interface {
	devicemanager // interface com.deepin.devicemanager
	proxy.Object
}

type objectDeviceManager struct {
	interfaceDevicemanager // interface com.deepin.devicemanager
	proxy.ImplObject
}

func NewDeviceManager(conn *dbus.Conn) DeviceManager {
	obj := new(objectDeviceManager)
	obj.ImplObject.Init_(conn, "com.deepin.devicemanager", "/com/deepin/devicemanager")
	return obj
}

type devicemanager interface {
	GoGetInfo(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call
	GetInfo(flags dbus.Flags, arg1 string) (string, error)
	GoExecCmd(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call
	ExecCmd(flags dbus.Flags, arg1 string) (string, error)
}

type interfaceDevicemanager struct{}

func (v *interfaceDevicemanager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDevicemanager) GetInterfaceName_() string {
	return "com.deepin.devicemanager"
}

// method getInfo

func (v *interfaceDevicemanager) GoGetInfo(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".getInfo", flags, ch, arg1)
}

func (*interfaceDevicemanager) StoreGetInfo(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDevicemanager) GetInfo(flags dbus.Flags, arg1 string) (string, error) {
	return v.StoreGetInfo(
		<-v.GoGetInfo(flags, make(chan *dbus.Call, 1), arg1).Done)
}

// method execCmd

func (v *interfaceDevicemanager) GoExecCmd(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".execCmd", flags, ch, arg1)
}

func (*interfaceDevicemanager) StoreExecCmd(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDevicemanager) ExecCmd(flags dbus.Flags, arg1 string) (string, error) {
	return v.StoreExecCmd(
		<-v.GoExecCmd(flags, make(chan *dbus.Call, 1), arg1).Done)
}
