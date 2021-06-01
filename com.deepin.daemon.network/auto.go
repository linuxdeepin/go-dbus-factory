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
	network // interface com.deepin.daemon.Network
	proxy.Object
}

func NewNetwork(conn *dbus.Conn) *Network {
	obj := new(Network)
	obj.Object.Init_(conn, "com.deepin.daemon.Network", "/com/deepin/daemon/Network")
	return obj
}

type network struct{}

func (v *network) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*network) GetInterfaceName_() string {
	return "com.deepin.daemon.Network"
}

// method ActivateAccessPoint

func (v *network) GoActivateAccessPoint(flags dbus.Flags, ch chan *dbus.Call, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateAccessPoint", flags, ch, uuid, apPath, devPath)
}

func (v *network) GoActivateAccessPointWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ActivateAccessPoint", flags, ch, uuid, apPath, devPath)
}

func (*network) StoreActivateAccessPoint(call *dbus.Call) (cPath dbus.ObjectPath, err error) {
	err = call.Store(&cPath)
	return
}

func (v *network) ActivateAccessPoint(flags dbus.Flags, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	return v.StoreActivateAccessPoint(
		<-v.GoActivateAccessPoint(flags, make(chan *dbus.Call, 1), uuid, apPath, devPath).Done)
}

func (v *network) ActivateAccessPointWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoActivateAccessPointWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid, apPath, devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreActivateAccessPoint(call)
}

// method ActivateConnection

func (v *network) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateConnection", flags, ch, uuid, devPath)
}

func (v *network) GoActivateConnectionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ActivateConnection", flags, ch, uuid, devPath)
}

func (*network) StoreActivateConnection(call *dbus.Call) (cPath dbus.ObjectPath, err error) {
	err = call.Store(&cPath)
	return
}

func (v *network) ActivateConnection(flags dbus.Flags, uuid string, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	return v.StoreActivateConnection(
		<-v.GoActivateConnection(flags, make(chan *dbus.Call, 1), uuid, devPath).Done)
}

func (v *network) ActivateConnectionWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoActivateConnectionWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid, devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreActivateConnection(call)
}

// method DeactivateConnection

func (v *network) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeactivateConnection", flags, ch, uuid)
}

func (v *network) GoDeactivateConnectionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeactivateConnection", flags, ch, uuid)
}

func (v *network) DeactivateConnection(flags dbus.Flags, uuid string) error {
	return (<-v.GoDeactivateConnection(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

func (v *network) DeactivateConnectionWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeactivateConnectionWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteConnection

func (v *network) GoDeleteConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteConnection", flags, ch, uuid)
}

func (v *network) GoDeleteConnectionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteConnection", flags, ch, uuid)
}

func (v *network) DeleteConnection(flags dbus.Flags, uuid string) error {
	return (<-v.GoDeleteConnection(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

func (v *network) DeleteConnectionWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteConnectionWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DisableWirelessHotspotMode

func (v *network) GoDisableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) GoDisableWirelessHotspotModeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DisableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) DisableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoDisableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

func (v *network) DisableWirelessHotspotModeWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisableWirelessHotspotModeWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DisconnectDevice

func (v *network) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectDevice", flags, ch, devPath)
}

func (v *network) GoDisconnectDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DisconnectDevice", flags, ch, devPath)
}

func (v *network) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoDisconnectDevice(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

func (v *network) DisconnectDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisconnectDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnableDevice

func (v *network) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableDevice", flags, ch, devPath, enabled)
}

func (v *network) GoEnableDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableDevice", flags, ch, devPath, enabled)
}

func (*network) StoreEnableDevice(call *dbus.Call) (cpath dbus.ObjectPath, err error) {
	err = call.Store(&cpath)
	return
}

func (v *network) EnableDevice(flags dbus.Flags, devPath dbus.ObjectPath, enabled bool) (cpath dbus.ObjectPath, err error) {
	return v.StoreEnableDevice(
		<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), devPath, enabled).Done)
}

func (v *network) EnableDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath, enabled bool) (cpath dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath, enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreEnableDevice(call)
}

// method EnableWirelessHotspotMode

func (v *network) GoEnableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) GoEnableWirelessHotspotModeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) EnableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoEnableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

func (v *network) EnableWirelessHotspotModeWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableWirelessHotspotModeWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetAccessPoints

func (v *network) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccessPoints", flags, ch, path)
}

func (v *network) GoGetAccessPointsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetAccessPoints", flags, ch, path)
}

func (*network) StoreGetAccessPoints(call *dbus.Call) (apsJSON string, err error) {
	err = call.Store(&apsJSON)
	return
}

func (v *network) GetAccessPoints(flags dbus.Flags, path dbus.ObjectPath) (apsJSON string, err error) {
	return v.StoreGetAccessPoints(
		<-v.GoGetAccessPoints(flags, make(chan *dbus.Call, 1), path).Done)
}

func (v *network) GetAccessPointsWithTimeout(timeout time.Duration, flags dbus.Flags, path dbus.ObjectPath) (apsJSON string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAccessPointsWithContext(ctx, flags, make(chan *dbus.Call, 1), path).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetAccessPoints(call)
}

// method GetActiveConnectionInfo

func (v *network) GoGetActiveConnectionInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetActiveConnectionInfo", flags, ch)
}

func (v *network) GoGetActiveConnectionInfoWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetActiveConnectionInfo", flags, ch)
}

func (*network) StoreGetActiveConnectionInfo(call *dbus.Call) (acInfosJSON string, err error) {
	err = call.Store(&acInfosJSON)
	return
}

func (v *network) GetActiveConnectionInfo(flags dbus.Flags) (acInfosJSON string, err error) {
	return v.StoreGetActiveConnectionInfo(
		<-v.GoGetActiveConnectionInfo(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) GetActiveConnectionInfoWithTimeout(timeout time.Duration, flags dbus.Flags) (acInfosJSON string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetActiveConnectionInfoWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetActiveConnectionInfo(call)
}

// method GetAutoProxy

func (v *network) GoGetAutoProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAutoProxy", flags, ch)
}

func (v *network) GoGetAutoProxyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetAutoProxy", flags, ch)
}

func (*network) StoreGetAutoProxy(call *dbus.Call) (proxyAuto string, err error) {
	err = call.Store(&proxyAuto)
	return
}

func (v *network) GetAutoProxy(flags dbus.Flags) (proxyAuto string, err error) {
	return v.StoreGetAutoProxy(
		<-v.GoGetAutoProxy(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) GetAutoProxyWithTimeout(timeout time.Duration, flags dbus.Flags) (proxyAuto string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAutoProxyWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetAutoProxy(call)
}

// method GetProxy

func (v *network) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch, proxyType)
}

func (v *network) GoGetProxyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, proxyType string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetProxy", flags, ch, proxyType)
}

func (*network) StoreGetProxy(call *dbus.Call) (host string, port string, err error) {
	err = call.Store(&host, &port)
	return
}

func (v *network) GetProxy(flags dbus.Flags, proxyType string) (host string, port string, err error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1), proxyType).Done)
}

func (v *network) GetProxyWithTimeout(timeout time.Duration, flags dbus.Flags, proxyType string) (host string, port string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetProxyWithContext(ctx, flags, make(chan *dbus.Call, 1), proxyType).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetProxy(call)
}

// method GetProxyIgnoreHosts

func (v *network) GoGetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyIgnoreHosts", flags, ch)
}

func (v *network) GoGetProxyIgnoreHostsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetProxyIgnoreHosts", flags, ch)
}

func (*network) StoreGetProxyIgnoreHosts(call *dbus.Call) (ignoreHosts string, err error) {
	err = call.Store(&ignoreHosts)
	return
}

func (v *network) GetProxyIgnoreHosts(flags dbus.Flags) (ignoreHosts string, err error) {
	return v.StoreGetProxyIgnoreHosts(
		<-v.GoGetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) GetProxyIgnoreHostsWithTimeout(timeout time.Duration, flags dbus.Flags) (ignoreHosts string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetProxyIgnoreHostsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetProxyIgnoreHosts(call)
}

// method GetProxyMethod

func (v *network) GoGetProxyMethod(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyMethod", flags, ch)
}

func (v *network) GoGetProxyMethodWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetProxyMethod", flags, ch)
}

func (*network) StoreGetProxyMethod(call *dbus.Call) (proxyMode string, err error) {
	err = call.Store(&proxyMode)
	return
}

func (v *network) GetProxyMethod(flags dbus.Flags) (proxyMode string, err error) {
	return v.StoreGetProxyMethod(
		<-v.GoGetProxyMethod(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) GetProxyMethodWithTimeout(timeout time.Duration, flags dbus.Flags) (proxyMode string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetProxyMethodWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetProxyMethod(call)
}

// method GetSupportedConnectionTypes

func (v *network) GoGetSupportedConnectionTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSupportedConnectionTypes", flags, ch)
}

func (v *network) GoGetSupportedConnectionTypesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetSupportedConnectionTypes", flags, ch)
}

func (*network) StoreGetSupportedConnectionTypes(call *dbus.Call) (types []string, err error) {
	err = call.Store(&types)
	return
}

func (v *network) GetSupportedConnectionTypes(flags dbus.Flags) (types []string, err error) {
	return v.StoreGetSupportedConnectionTypes(
		<-v.GoGetSupportedConnectionTypes(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *network) GetSupportedConnectionTypesWithTimeout(timeout time.Duration, flags dbus.Flags) (types []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetSupportedConnectionTypesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetSupportedConnectionTypes(call)
}

// method IsDeviceEnabled

func (v *network) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, devPath)
}

func (v *network) GoIsDeviceEnabledWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, devPath)
}

func (*network) StoreIsDeviceEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsDeviceEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), devPath).Done)
}

func (v *network) IsDeviceEnabledWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsDeviceEnabledWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsDeviceEnabled(call)
}

// method IsWirelessHotspotModeEnabled

func (v *network) GoIsWirelessHotspotModeEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsWirelessHotspotModeEnabled", flags, ch, devPath)
}

func (v *network) GoIsWirelessHotspotModeEnabledWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsWirelessHotspotModeEnabled", flags, ch, devPath)
}

func (*network) StoreIsWirelessHotspotModeEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsWirelessHotspotModeEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	return v.StoreIsWirelessHotspotModeEnabled(
		<-v.GoIsWirelessHotspotModeEnabled(flags, make(chan *dbus.Call, 1), devPath).Done)
}

func (v *network) IsWirelessHotspotModeEnabledWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsWirelessHotspotModeEnabledWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsWirelessHotspotModeEnabled(call)
}

// method ListDeviceConnections

func (v *network) GoListDeviceConnections(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListDeviceConnections", flags, ch, devPath)
}

func (v *network) GoListDeviceConnectionsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListDeviceConnections", flags, ch, devPath)
}

func (*network) StoreListDeviceConnections(call *dbus.Call) (connections []dbus.ObjectPath, err error) {
	err = call.Store(&connections)
	return
}

func (v *network) ListDeviceConnections(flags dbus.Flags, devPath dbus.ObjectPath) (connections []dbus.ObjectPath, err error) {
	return v.StoreListDeviceConnections(
		<-v.GoListDeviceConnections(flags, make(chan *dbus.Call, 1), devPath).Done)
}

func (v *network) ListDeviceConnectionsWithTimeout(timeout time.Duration, flags dbus.Flags, devPath dbus.ObjectPath) (connections []dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListDeviceConnectionsWithContext(ctx, flags, make(chan *dbus.Call, 1), devPath).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListDeviceConnections(call)
}

// method RequestWirelessScan

func (v *network) GoRequestWirelessScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestWirelessScan", flags, ch)
}

func (v *network) GoRequestWirelessScanWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestWirelessScan", flags, ch)
}

func (v *network) RequestWirelessScan(flags dbus.Flags) error {
	return (<-v.GoRequestWirelessScan(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *network) RequestWirelessScanWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestWirelessScanWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetAutoProxy

func (v *network) GoSetAutoProxy(flags dbus.Flags, ch chan *dbus.Call, proxyAuto string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoProxy", flags, ch, proxyAuto)
}

func (v *network) GoSetAutoProxyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, proxyAuto string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAutoProxy", flags, ch, proxyAuto)
}

func (v *network) SetAutoProxy(flags dbus.Flags, proxyAuto string) error {
	return (<-v.GoSetAutoProxy(flags, make(chan *dbus.Call, 1), proxyAuto).Done).Err
}

func (v *network) SetAutoProxyWithTimeout(timeout time.Duration, flags dbus.Flags, proxyAuto string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAutoProxyWithContext(ctx, flags, make(chan *dbus.Call, 1), proxyAuto).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetDeviceManaged

func (v *network) GoSetDeviceManaged(flags dbus.Flags, ch chan *dbus.Call, devPathOrIfc string, managed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeviceManaged", flags, ch, devPathOrIfc, managed)
}

func (v *network) GoSetDeviceManagedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, devPathOrIfc string, managed bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDeviceManaged", flags, ch, devPathOrIfc, managed)
}

func (v *network) SetDeviceManaged(flags dbus.Flags, devPathOrIfc string, managed bool) error {
	return (<-v.GoSetDeviceManaged(flags, make(chan *dbus.Call, 1), devPathOrIfc, managed).Done).Err
}

func (v *network) SetDeviceManagedWithTimeout(timeout time.Duration, flags dbus.Flags, devPathOrIfc string, managed bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDeviceManagedWithContext(ctx, flags, make(chan *dbus.Call, 1), devPathOrIfc, managed).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetProxy

func (v *network) GoSetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string, host string, port string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxy", flags, ch, proxyType, host, port)
}

func (v *network) GoSetProxyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, proxyType string, host string, port string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetProxy", flags, ch, proxyType, host, port)
}

func (v *network) SetProxy(flags dbus.Flags, proxyType string, host string, port string) error {
	return (<-v.GoSetProxy(flags, make(chan *dbus.Call, 1), proxyType, host, port).Done).Err
}

func (v *network) SetProxyWithTimeout(timeout time.Duration, flags dbus.Flags, proxyType string, host string, port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetProxyWithContext(ctx, flags, make(chan *dbus.Call, 1), proxyType, host, port).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetProxyIgnoreHosts

func (v *network) GoSetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call, ignoreHosts string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyIgnoreHosts", flags, ch, ignoreHosts)
}

func (v *network) GoSetProxyIgnoreHostsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, ignoreHosts string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetProxyIgnoreHosts", flags, ch, ignoreHosts)
}

func (v *network) SetProxyIgnoreHosts(flags dbus.Flags, ignoreHosts string) error {
	return (<-v.GoSetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1), ignoreHosts).Done).Err
}

func (v *network) SetProxyIgnoreHostsWithTimeout(timeout time.Duration, flags dbus.Flags, ignoreHosts string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetProxyIgnoreHostsWithContext(ctx, flags, make(chan *dbus.Call, 1), ignoreHosts).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetProxyMethod

func (v *network) GoSetProxyMethod(flags dbus.Flags, ch chan *dbus.Call, proxyMode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyMethod", flags, ch, proxyMode)
}

func (v *network) GoSetProxyMethodWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, proxyMode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetProxyMethod", flags, ch, proxyMode)
}

func (v *network) SetProxyMethod(flags dbus.Flags, proxyMode string) error {
	return (<-v.GoSetProxyMethod(flags, make(chan *dbus.Call, 1), proxyMode).Done).Err
}

func (v *network) SetProxyMethodWithTimeout(timeout time.Duration, flags dbus.Flags, proxyMode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetProxyMethodWithContext(ctx, flags, make(chan *dbus.Call, 1), proxyMode).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal AccessPointAdded

func (v *network) ConnectAccessPointAdded(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AccessPointAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AccessPointAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devPath string
		var apJSON string
		err := dbus.Store(sig.Body, &devPath, &apJSON)
		if err == nil {
			cb(devPath, apJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointRemoved

func (v *network) ConnectAccessPointRemoved(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AccessPointRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AccessPointRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devPath string
		var apJSON string
		err := dbus.Store(sig.Body, &devPath, &apJSON)
		if err == nil {
			cb(devPath, apJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointPropertiesChanged

func (v *network) ConnectAccessPointPropertiesChanged(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AccessPointPropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AccessPointPropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devPath string
		var apJSON string
		err := dbus.Store(sig.Body, &devPath, &apJSON)
		if err == nil {
			cb(devPath, apJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceEnabled

func (v *network) ConnectDeviceEnabled(cb func(devPath string, enabled bool)) (dbusutil.SignalHandlerId, error) {
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
		var devPath string
		var enabled bool
		err := dbus.Store(sig.Body, &devPath, &enabled)
		if err == nil {
			cb(devPath, enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Connectivity u

func (v *network) Connectivity() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Connectivity",
	}
}

// property NetworkingEnabled b

func (v *network) NetworkingEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NetworkingEnabled",
	}
}

// property VpnEnabled b

func (v *network) VpnEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "VpnEnabled",
	}
}

// property Devices s

func (v *network) Devices() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Devices",
	}
}

// property Connections s

func (v *network) Connections() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Connections",
	}
}

// property ActiveConnections s

func (v *network) ActiveConnections() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ActiveConnections",
	}
}

// property State u

func (v *network) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}
