package daemon

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

type Daemon struct {
	daemon // interface com.deepin.daemon.Daemon
	proxy.Object
}

func NewDaemon(conn *dbus.Conn) *Daemon {
	obj := new(Daemon)
	obj.Object.Init_(conn, "com.deepin.daemon.Daemon", "/com/deepin/daemon/Daemon")
	return obj
}

type daemon struct{}

func (v *daemon) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*daemon) GetInterfaceName_() string {
	return "com.deepin.daemon.Daemon"
}

// method BluetoothGetDeviceTechnologies

func (v *daemon) GoBluetoothGetDeviceTechnologies(flags dbus.Flags, ch chan *dbus.Call, adapter string, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BluetoothGetDeviceTechnologies", flags, ch, adapter, device)
}

func (v *daemon) GoBluetoothGetDeviceTechnologiesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, adapter string, device string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".BluetoothGetDeviceTechnologies", flags, ch, adapter, device)
}

func (*daemon) StoreBluetoothGetDeviceTechnologies(call *dbus.Call) (technologies []string, err error) {
	err = call.Store(&technologies)
	return
}

func (v *daemon) BluetoothGetDeviceTechnologies(flags dbus.Flags, adapter string, device string) (technologies []string, err error) {
	return v.StoreBluetoothGetDeviceTechnologies(
		<-v.GoBluetoothGetDeviceTechnologies(flags, make(chan *dbus.Call, 1), adapter, device).Done)
}

func (v *daemon) BluetoothGetDeviceTechnologiesWithTimeout(timeout time.Duration, flags dbus.Flags, adapter string, device string) (technologies []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoBluetoothGetDeviceTechnologiesWithContext(ctx, flags, make(chan *dbus.Call, 1), adapter, device).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreBluetoothGetDeviceTechnologies(call)
}

// method NetworkGetConnections

func (v *daemon) GoNetworkGetConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NetworkGetConnections", flags, ch)
}

func (v *daemon) GoNetworkGetConnectionsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".NetworkGetConnections", flags, ch)
}

func (*daemon) StoreNetworkGetConnections(call *dbus.Call) (data []uint8, err error) {
	err = call.Store(&data)
	return
}

func (v *daemon) NetworkGetConnections(flags dbus.Flags) (data []uint8, err error) {
	return v.StoreNetworkGetConnections(
		<-v.GoNetworkGetConnections(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *daemon) NetworkGetConnectionsWithTimeout(timeout time.Duration, flags dbus.Flags) (data []uint8, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoNetworkGetConnectionsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreNetworkGetConnections(call)
}

// method NetworkSetConnections

func (v *daemon) GoNetworkSetConnections(flags dbus.Flags, ch chan *dbus.Call, data []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NetworkSetConnections", flags, ch, data)
}

func (v *daemon) GoNetworkSetConnectionsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, data []uint8) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".NetworkSetConnections", flags, ch, data)
}

func (v *daemon) NetworkSetConnections(flags dbus.Flags, data []uint8) error {
	return (<-v.GoNetworkSetConnections(flags, make(chan *dbus.Call, 1), data).Done).Err
}

func (v *daemon) NetworkSetConnectionsWithTimeout(timeout time.Duration, flags dbus.Flags, data []uint8) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoNetworkSetConnectionsWithContext(ctx, flags, make(chan *dbus.Call, 1), data).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ScalePlymouth

func (v *daemon) GoScalePlymouth(flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ScalePlymouth", flags, ch, scale)
}

func (v *daemon) GoScalePlymouthWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ScalePlymouth", flags, ch, scale)
}

func (v *daemon) ScalePlymouth(flags dbus.Flags, scale uint32) error {
	return (<-v.GoScalePlymouth(flags, make(chan *dbus.Call, 1), scale).Done).Err
}

func (v *daemon) ScalePlymouthWithTimeout(timeout time.Duration, flags dbus.Flags, scale uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoScalePlymouthWithContext(ctx, flags, make(chan *dbus.Call, 1), scale).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLongPressDuration

func (v *daemon) GoSetLongPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongPressDuration", flags, ch, duration)
}

func (v *daemon) GoSetLongPressDurationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLongPressDuration", flags, ch, duration)
}

func (v *daemon) SetLongPressDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetLongPressDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

func (v *daemon) SetLongPressDurationWithTimeout(timeout time.Duration, flags dbus.Flags, duration uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLongPressDurationWithContext(ctx, flags, make(chan *dbus.Call, 1), duration).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
