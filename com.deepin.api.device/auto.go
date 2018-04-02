package device

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

type Device struct {
	device // interface com.deepin.api.Device
	proxy.Object
}

func NewDevice(conn *dbus.Conn) *Device {
	obj := new(Device)
	obj.Object.Init_(conn, "com.deepin.api.Device", "/com/deepin/api/Device")
	return obj
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*device) GetInterfaceName_() string {
	return "com.deepin.api.Device"
}

// method BlockDevice

func (v *device) GoBlockDevice(flags dbus.Flags, ch chan *dbus.Call, deviceType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BlockDevice", flags, ch, deviceType)
}

func (v *device) BlockDevice(flags dbus.Flags, deviceType string) error {
	return (<-v.GoBlockDevice(flags, make(chan *dbus.Call, 1), deviceType).Done).Err
}

// method UnblockDevice

func (v *device) GoUnblockDevice(flags dbus.Flags, ch chan *dbus.Call, deviceType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnblockDevice", flags, ch, deviceType)
}

func (v *device) UnblockDevice(flags dbus.Flags, deviceType string) error {
	return (<-v.GoUnblockDevice(flags, make(chan *dbus.Call, 1), deviceType).Done).Err
}
