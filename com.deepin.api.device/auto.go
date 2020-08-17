package device

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

// method HasBluetoothDeviceBlocked

func (v *device) GoHasBluetoothDeviceBlocked(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".HasBluetoothDeviceBlocked", flags, ch)
}

func (*device) StoreHasBluetoothDeviceBlocked(call *dbus.Call) (has bool, err error) {
	err = call.Store(&has)
	return
}

func (v *device) HasBluetoothDeviceBlocked(flags dbus.Flags) (has bool, err error) {
	return v.StoreHasBluetoothDeviceBlocked(
		<-v.GoHasBluetoothDeviceBlocked(flags, make(chan *dbus.Call, 1)).Done)
}

// method UnblockBluetoothDevices

func (v *device) GoUnblockBluetoothDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnblockBluetoothDevices", flags, ch)
}

func (v *device) UnblockBluetoothDevices(flags dbus.Flags) error {
	return (<-v.GoUnblockBluetoothDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}
