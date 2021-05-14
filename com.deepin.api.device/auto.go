package device

import "context"
import "errors"
import "fmt"
import dbus "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
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

func (v *device) GoHasBluetoothDeviceBlockedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".HasBluetoothDeviceBlocked", flags, ch)
}

func (*device) StoreHasBluetoothDeviceBlocked(call *dbus.Call) (has bool, err error) {
	err = call.Store(&has)
	return
}

func (v *device) HasBluetoothDeviceBlocked(flags dbus.Flags) (has bool, err error) {
	return v.StoreHasBluetoothDeviceBlocked(
		<-v.GoHasBluetoothDeviceBlocked(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *device) HasBluetoothDeviceBlockedWithTimeout(timeout time.Duration, flags dbus.Flags) (has bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHasBluetoothDeviceBlockedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreHasBluetoothDeviceBlocked(call)
}

// method UnblockBluetoothDevices

func (v *device) GoUnblockBluetoothDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnblockBluetoothDevices", flags, ch)
}

func (v *device) GoUnblockBluetoothDevicesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnblockBluetoothDevices", flags, ch)
}

func (v *device) UnblockBluetoothDevices(flags dbus.Flags) error {
	return (<-v.GoUnblockBluetoothDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) UnblockBluetoothDevicesWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnblockBluetoothDevicesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
