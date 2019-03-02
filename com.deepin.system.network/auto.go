package network

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

type Network struct {
	network // interface com.deepin.system.Network
	proxy.Object
}

func NewNetwork(conn *dbus.Conn) *Network {
	obj := new(Network)
	obj.Object.Init_(conn, "com.deepin.system.Network", "/com/deepin/system/Network")
	return obj
}

type network struct{}

func (v *network) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*network) GetInterfaceName_() string {
	return "com.deepin.system.Network"
}

// method EnableDevice

func (v *network) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableDevice", flags, ch, pathOrIface, enabled)
}

func (v *network) EnableDevice(flags dbus.Flags, pathOrIface string, enabled bool) error {
	return (<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done).Err
}

// method IsDeviceEnabled

func (v *network) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, pathOrIface)
}

func (*network) StoreIsDeviceEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsDeviceEnabled(flags dbus.Flags, pathOrIface string) (enabled bool, err error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), pathOrIface).Done)
}

// signal DeviceEnabled

func (v *network) ConnectDeviceEnabled(cb func(devPath dbus.ObjectPath, enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceEnabled", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceEnabled",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devPath dbus.ObjectPath
		var enabled bool
		err := dbus.Store(sig.Body, &devPath, &enabled)
		if err == nil {
			cb(devPath, enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property VpnEnabled b

func (v *network) VpnEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "VpnEnabled",
	}
}
