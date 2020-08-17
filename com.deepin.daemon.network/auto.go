package network

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

func (*network) StoreActivateAccessPoint(call *dbus.Call) (cPath dbus.ObjectPath, err error) {
	err = call.Store(&cPath)
	return
}

func (v *network) ActivateAccessPoint(flags dbus.Flags, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	return v.StoreActivateAccessPoint(
		<-v.GoActivateAccessPoint(flags, make(chan *dbus.Call, 1), uuid, apPath, devPath).Done)
}

// method ActivateConnection

func (v *network) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateConnection", flags, ch, uuid, devPath)
}

func (*network) StoreActivateConnection(call *dbus.Call) (cPath dbus.ObjectPath, err error) {
	err = call.Store(&cPath)
	return
}

func (v *network) ActivateConnection(flags dbus.Flags, uuid string, devPath dbus.ObjectPath) (cPath dbus.ObjectPath, err error) {
	return v.StoreActivateConnection(
		<-v.GoActivateConnection(flags, make(chan *dbus.Call, 1), uuid, devPath).Done)
}

// method DeactivateConnection

func (v *network) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeactivateConnection", flags, ch, uuid)
}

func (v *network) DeactivateConnection(flags dbus.Flags, uuid string) error {
	return (<-v.GoDeactivateConnection(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

// method DeleteConnection

func (v *network) GoDeleteConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteConnection", flags, ch, uuid)
}

func (v *network) DeleteConnection(flags dbus.Flags, uuid string) error {
	return (<-v.GoDeleteConnection(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

// method DisableWirelessHotspotMode

func (v *network) GoDisableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) DisableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoDisableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

// method DisconnectDevice

func (v *network) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectDevice", flags, ch, devPath)
}

func (v *network) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoDisconnectDevice(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

// method EnableDevice

func (v *network) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableDevice", flags, ch, devPath, enabled)
}

func (*network) StoreEnableDevice(call *dbus.Call) (cpath dbus.ObjectPath, err error) {
	err = call.Store(&cpath)
	return
}

func (v *network) EnableDevice(flags dbus.Flags, devPath dbus.ObjectPath, enabled bool) (cpath dbus.ObjectPath, err error) {
	return v.StoreEnableDevice(
		<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), devPath, enabled).Done)
}

// method EnableWirelessHotspotMode

func (v *network) GoEnableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWirelessHotspotMode", flags, ch, devPath)
}

func (v *network) EnableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoEnableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

// method GetAccessPoints

func (v *network) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccessPoints", flags, ch, path)
}

func (*network) StoreGetAccessPoints(call *dbus.Call) (apsJSON string, err error) {
	err = call.Store(&apsJSON)
	return
}

func (v *network) GetAccessPoints(flags dbus.Flags, path dbus.ObjectPath) (apsJSON string, err error) {
	return v.StoreGetAccessPoints(
		<-v.GoGetAccessPoints(flags, make(chan *dbus.Call, 1), path).Done)
}

// method GetActiveConnectionInfo

func (v *network) GoGetActiveConnectionInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetActiveConnectionInfo", flags, ch)
}

func (*network) StoreGetActiveConnectionInfo(call *dbus.Call) (acInfosJSON string, err error) {
	err = call.Store(&acInfosJSON)
	return
}

func (v *network) GetActiveConnectionInfo(flags dbus.Flags) (acInfosJSON string, err error) {
	return v.StoreGetActiveConnectionInfo(
		<-v.GoGetActiveConnectionInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAutoProxy

func (v *network) GoGetAutoProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAutoProxy", flags, ch)
}

func (*network) StoreGetAutoProxy(call *dbus.Call) (proxyAuto string, err error) {
	err = call.Store(&proxyAuto)
	return
}

func (v *network) GetAutoProxy(flags dbus.Flags) (proxyAuto string, err error) {
	return v.StoreGetAutoProxy(
		<-v.GoGetAutoProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProxy

func (v *network) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch, proxyType)
}

func (*network) StoreGetProxy(call *dbus.Call) (host string, port string, err error) {
	err = call.Store(&host, &port)
	return
}

func (v *network) GetProxy(flags dbus.Flags, proxyType string) (host string, port string, err error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1), proxyType).Done)
}

// method GetProxyIgnoreHosts

func (v *network) GoGetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyIgnoreHosts", flags, ch)
}

func (*network) StoreGetProxyIgnoreHosts(call *dbus.Call) (ignoreHosts string, err error) {
	err = call.Store(&ignoreHosts)
	return
}

func (v *network) GetProxyIgnoreHosts(flags dbus.Flags) (ignoreHosts string, err error) {
	return v.StoreGetProxyIgnoreHosts(
		<-v.GoGetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProxyMethod

func (v *network) GoGetProxyMethod(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyMethod", flags, ch)
}

func (*network) StoreGetProxyMethod(call *dbus.Call) (proxyMode string, err error) {
	err = call.Store(&proxyMode)
	return
}

func (v *network) GetProxyMethod(flags dbus.Flags) (proxyMode string, err error) {
	return v.StoreGetProxyMethod(
		<-v.GoGetProxyMethod(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSupportedConnectionTypes

func (v *network) GoGetSupportedConnectionTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSupportedConnectionTypes", flags, ch)
}

func (*network) StoreGetSupportedConnectionTypes(call *dbus.Call) (types []string, err error) {
	err = call.Store(&types)
	return
}

func (v *network) GetSupportedConnectionTypes(flags dbus.Flags) (types []string, err error) {
	return v.StoreGetSupportedConnectionTypes(
		<-v.GoGetSupportedConnectionTypes(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsDeviceEnabled

func (v *network) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, devPath)
}

func (*network) StoreIsDeviceEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsDeviceEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), devPath).Done)
}

// method IsWirelessHotspotModeEnabled

func (v *network) GoIsWirelessHotspotModeEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsWirelessHotspotModeEnabled", flags, ch, devPath)
}

func (*network) StoreIsWirelessHotspotModeEnabled(call *dbus.Call) (enabled bool, err error) {
	err = call.Store(&enabled)
	return
}

func (v *network) IsWirelessHotspotModeEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (enabled bool, err error) {
	return v.StoreIsWirelessHotspotModeEnabled(
		<-v.GoIsWirelessHotspotModeEnabled(flags, make(chan *dbus.Call, 1), devPath).Done)
}

// method ListDeviceConnections

func (v *network) GoListDeviceConnections(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListDeviceConnections", flags, ch, devPath)
}

func (*network) StoreListDeviceConnections(call *dbus.Call) (connections []dbus.ObjectPath, err error) {
	err = call.Store(&connections)
	return
}

func (v *network) ListDeviceConnections(flags dbus.Flags, devPath dbus.ObjectPath) (connections []dbus.ObjectPath, err error) {
	return v.StoreListDeviceConnections(
		<-v.GoListDeviceConnections(flags, make(chan *dbus.Call, 1), devPath).Done)
}

// method RequestWirelessScan

func (v *network) GoRequestWirelessScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestWirelessScan", flags, ch)
}

func (v *network) RequestWirelessScan(flags dbus.Flags) error {
	return (<-v.GoRequestWirelessScan(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestWirelessScanHasReturn

func (v *network) GoRequestWirelessScanHasReturn(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestWirelessScanHasReturn", flags, ch)
}

func (*network) StoreRequestWirelessScanHasReturn(call *dbus.Call) (apsJSONs string, err error) {
	err = call.Store(&apsJSONs)
	return
}

func (v *network) RequestWirelessScanHasReturn(flags dbus.Flags) (apsJSONs string, err error) {
	return v.StoreRequestWirelessScanHasReturn(
		<-v.GoRequestWirelessScanHasReturn(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetAutoProxy

func (v *network) GoSetAutoProxy(flags dbus.Flags, ch chan *dbus.Call, proxyAuto string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoProxy", flags, ch, proxyAuto)
}

func (v *network) SetAutoProxy(flags dbus.Flags, proxyAuto string) error {
	return (<-v.GoSetAutoProxy(flags, make(chan *dbus.Call, 1), proxyAuto).Done).Err
}

// method SetDeviceManaged

func (v *network) GoSetDeviceManaged(flags dbus.Flags, ch chan *dbus.Call, devPathOrIfc string, managed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeviceManaged", flags, ch, devPathOrIfc, managed)
}

func (v *network) SetDeviceManaged(flags dbus.Flags, devPathOrIfc string, managed bool) error {
	return (<-v.GoSetDeviceManaged(flags, make(chan *dbus.Call, 1), devPathOrIfc, managed).Done).Err
}

// method SetProxy

func (v *network) GoSetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string, host string, port string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxy", flags, ch, proxyType, host, port)
}

func (v *network) SetProxy(flags dbus.Flags, proxyType string, host string, port string) error {
	return (<-v.GoSetProxy(flags, make(chan *dbus.Call, 1), proxyType, host, port).Done).Err
}

// method SetProxyIgnoreHosts

func (v *network) GoSetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call, ignoreHosts string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyIgnoreHosts", flags, ch, ignoreHosts)
}

func (v *network) SetProxyIgnoreHosts(flags dbus.Flags, ignoreHosts string) error {
	return (<-v.GoSetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1), ignoreHosts).Done).Err
}

// method SetProxyMethod

func (v *network) GoSetProxyMethod(flags dbus.Flags, ch chan *dbus.Call, proxyMode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyMethod", flags, ch, proxyMode)
}

func (v *network) SetProxyMethod(flags dbus.Flags, proxyMode string) error {
	return (<-v.GoSetProxyMethod(flags, make(chan *dbus.Call, 1), proxyMode).Done).Err
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
