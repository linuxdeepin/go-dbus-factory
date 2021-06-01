package network

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
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

func (v *network) GoEnableDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableDevice", flags, ch, pathOrIface, enabled)
}

func (*network) StoreEnableDevice(call *dbus.Call) (cpath dbus.ObjectPath, err error) {
	err = call.Store(&cpath)
	return
}

func (v *network) EnableDevice(flags dbus.Flags, pathOrIface string, enabled bool) (cpath dbus.ObjectPath, err error) {
	return v.StoreEnableDevice(
		<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done)
}

func (v *network) EnableDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, pathOrIface string, enabled bool) (cpath dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreEnableDevice(call)
}

// method IsDeviceEnabled

func (v *network) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, pathOrIface)
}

func (v *network) GoIsDeviceEnabledWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pathOrIface string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, pathOrIface)
}

func (*network) StoreIsDeviceEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsDeviceEnabled(flags dbus.Flags, pathOrIface string) (enabled bool, err error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), pathOrIface).Done)
}

func (v *network) IsDeviceEnabledWithTimeout(timeout time.Duration, flags dbus.Flags, pathOrIface string) (enabled bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsDeviceEnabledWithContext(ctx, flags, make(chan *dbus.Call, 1), pathOrIface).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsDeviceEnabled(call)
}

// method Ping

func (v *network) GoPing(flags dbus.Flags, ch chan *dbus.Call, host string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Ping", flags, ch, host)
}

func (v *network) GoPingWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, host string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Ping", flags, ch, host)
}

func (v *network) Ping(flags dbus.Flags, host string) error {
	return (<-v.GoPing(flags, make(chan *dbus.Call, 1), host).Done).Err
}

func (v *network) PingWithTimeout(timeout time.Duration, flags dbus.Flags, host string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPingWithContext(ctx, flags, make(chan *dbus.Call, 1), host).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ToggleWirelessEnabled

func (v *network) GoToggleWirelessEnabled(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleWirelessEnabled", flags, ch)
}

func (v *network) GoToggleWirelessEnabledWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ToggleWirelessEnabled", flags, ch)
}

func (*network) StoreToggleWirelessEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) ToggleWirelessEnabled(flags dbus.Flags) (enabled bool, err error) {
	return v.StoreToggleWirelessEnabled(
		<-v.GoToggleWirelessEnabled(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) ToggleWirelessEnabledWithTimeout(timeout time.Duration, flags dbus.Flags) (enabled bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoToggleWirelessEnabledWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreToggleWirelessEnabled(call)
}

// method OpenWifiDevice

func (v *network) GoOpenWifiDevice(flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenWifiDevice", flags, ch, pathOrIface, enabled)
}

func (v *network) GoOpenWifiDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pathOrIface string, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".OpenWifiDevice", flags, ch, pathOrIface, enabled)
}

func (v *network) OpenWifiDevice(flags dbus.Flags, pathOrIface string, enabled bool) error {
	return (<-v.GoOpenWifiDevice(flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done).Err
}

func (v *network) OpenWifiDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, pathOrIface string, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoOpenWifiDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), pathOrIface, enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
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
