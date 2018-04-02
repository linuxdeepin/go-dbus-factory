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

func (v *network) GoActivateAccessPoint(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 dbus.ObjectPath, arg2 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateAccessPoint", flags, ch, arg0, arg1, arg2)
}

func (*network) StoreActivateAccessPoint(call *dbus.Call) (arg3 dbus.ObjectPath, err error) {
	err = call.Store(&arg3)
	return
}

func (v *network) ActivateAccessPoint(flags dbus.Flags, arg0 string, arg1 dbus.ObjectPath, arg2 dbus.ObjectPath) (arg3 dbus.ObjectPath, err error) {
	return v.StoreActivateAccessPoint(
		<-v.GoActivateAccessPoint(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done)
}

// method ActivateConnection

func (v *network) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateConnection", flags, ch, arg0, arg1)
}

func (*network) StoreActivateConnection(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *network) ActivateConnection(flags dbus.Flags, arg0 string, arg1 dbus.ObjectPath) (arg2 dbus.ObjectPath, err error) {
	return v.StoreActivateConnection(
		<-v.GoActivateConnection(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method CancelSecret

func (v *network) GoCancelSecret(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelSecret", flags, ch, arg0, arg1)
}

func (v *network) CancelSecret(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoCancelSecret(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method CreateConnection

func (v *network) GoCreateConnection(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateConnection", flags, ch, arg0, arg1)
}

func (*network) StoreCreateConnection(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *network) CreateConnection(flags dbus.Flags, arg0 string, arg1 dbus.ObjectPath) (arg2 dbus.ObjectPath, err error) {
	return v.StoreCreateConnection(
		<-v.GoCreateConnection(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method CreateConnectionForAccessPoint

func (v *network) GoCreateConnectionForAccessPoint(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath, arg1 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateConnectionForAccessPoint", flags, ch, arg0, arg1)
}

func (*network) StoreCreateConnectionForAccessPoint(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *network) CreateConnectionForAccessPoint(flags dbus.Flags, arg0 dbus.ObjectPath, arg1 dbus.ObjectPath) (arg2 dbus.ObjectPath, err error) {
	return v.StoreCreateConnectionForAccessPoint(
		<-v.GoCreateConnectionForAccessPoint(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method DeactivateConnection

func (v *network) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeactivateConnection", flags, ch, arg0)
}

func (v *network) DeactivateConnection(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDeactivateConnection(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DeleteConnection

func (v *network) GoDeleteConnection(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteConnection", flags, ch, arg0)
}

func (v *network) DeleteConnection(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDeleteConnection(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DisableWirelessHotspotMode

func (v *network) GoDisableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisableWirelessHotspotMode", flags, ch, arg0)
}

func (v *network) DisableWirelessHotspotMode(flags dbus.Flags, arg0 dbus.ObjectPath) error {
	return (<-v.GoDisableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DisconnectDevice

func (v *network) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectDevice", flags, ch, arg0)
}

func (v *network) DisconnectDevice(flags dbus.Flags, arg0 dbus.ObjectPath) error {
	return (<-v.GoDisconnectDevice(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method EditConnection

func (v *network) GoEditConnection(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EditConnection", flags, ch, arg0, arg1)
}

func (*network) StoreEditConnection(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *network) EditConnection(flags dbus.Flags, arg0 string, arg1 dbus.ObjectPath) (arg2 dbus.ObjectPath, err error) {
	return v.StoreEditConnection(
		<-v.GoEditConnection(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method EnableDevice

func (v *network) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableDevice", flags, ch, arg0, arg1)
}

func (v *network) EnableDevice(flags dbus.Flags, arg0 dbus.ObjectPath, arg1 bool) error {
	return (<-v.GoEnableDevice(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method EnableWirelessHotspotMode

func (v *network) GoEnableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWirelessHotspotMode", flags, ch, arg0)
}

func (v *network) EnableWirelessHotspotMode(flags dbus.Flags, arg0 dbus.ObjectPath) error {
	return (<-v.GoEnableWirelessHotspotMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method FeedSecret

func (v *network) GoFeedSecret(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 string, arg3 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FeedSecret", flags, ch, arg0, arg1, arg2, arg3)
}

func (v *network) FeedSecret(flags dbus.Flags, arg0 string, arg1 string, arg2 string, arg3 bool) error {
	return (<-v.GoFeedSecret(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3).Done).Err
}

// method GetAccessPoints

func (v *network) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccessPoints", flags, ch, arg0)
}

func (*network) StoreGetAccessPoints(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *network) GetAccessPoints(flags dbus.Flags, arg0 dbus.ObjectPath) (arg1 string, err error) {
	return v.StoreGetAccessPoints(
		<-v.GoGetAccessPoints(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetActiveConnectionInfo

func (v *network) GoGetActiveConnectionInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetActiveConnectionInfo", flags, ch)
}

func (*network) StoreGetActiveConnectionInfo(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *network) GetActiveConnectionInfo(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetActiveConnectionInfo(
		<-v.GoGetActiveConnectionInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAutoProxy

func (v *network) GoGetAutoProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAutoProxy", flags, ch)
}

func (*network) StoreGetAutoProxy(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *network) GetAutoProxy(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetAutoProxy(
		<-v.GoGetAutoProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProxy

func (v *network) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch, arg0)
}

func (*network) StoreGetProxy(call *dbus.Call) (arg1 string, arg2 string, err error) {
	err = call.Store(&arg1, &arg2)
	return
}

func (v *network) GetProxy(flags dbus.Flags, arg0 string) (arg1 string, arg2 string, err error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetProxyIgnoreHosts

func (v *network) GoGetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyIgnoreHosts", flags, ch)
}

func (*network) StoreGetProxyIgnoreHosts(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *network) GetProxyIgnoreHosts(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetProxyIgnoreHosts(
		<-v.GoGetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProxyMethod

func (v *network) GoGetProxyMethod(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxyMethod", flags, ch)
}

func (*network) StoreGetProxyMethod(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *network) GetProxyMethod(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetProxyMethod(
		<-v.GoGetProxyMethod(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSupportedConnectionTypes

func (v *network) GoGetSupportedConnectionTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSupportedConnectionTypes", flags, ch)
}

func (*network) StoreGetSupportedConnectionTypes(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *network) GetSupportedConnectionTypes(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreGetSupportedConnectionTypes(
		<-v.GoGetSupportedConnectionTypes(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetWiredConnectionUuid

func (v *network) GoGetWiredConnectionUuid(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWiredConnectionUuid", flags, ch, arg0)
}

func (*network) StoreGetWiredConnectionUuid(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *network) GetWiredConnectionUuid(flags dbus.Flags, arg0 dbus.ObjectPath) (arg1 string, err error) {
	return v.StoreGetWiredConnectionUuid(
		<-v.GoGetWiredConnectionUuid(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method IsDeviceEnabled

func (v *network) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDeviceEnabled", flags, ch, arg0)
}

func (*network) StoreIsDeviceEnabled(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *network) IsDeviceEnabled(flags dbus.Flags, arg0 dbus.ObjectPath) (arg1 bool, err error) {
	return v.StoreIsDeviceEnabled(
		<-v.GoIsDeviceEnabled(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method IsPasswordValid

func (v *network) GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPasswordValid", flags, ch, arg0, arg1)
}

func (*network) StoreIsPasswordValid(call *dbus.Call) (arg2 bool, err error) {
	err = call.Store(&arg2)
	return
}

func (v *network) IsPasswordValid(flags dbus.Flags, arg0 string, arg1 string) (arg2 bool, err error) {
	return v.StoreIsPasswordValid(
		<-v.GoIsPasswordValid(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method IsWirelessHotspotModeEnabled

func (v *network) GoIsWirelessHotspotModeEnabled(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsWirelessHotspotModeEnabled", flags, ch, arg0)
}

func (*network) StoreIsWirelessHotspotModeEnabled(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *network) IsWirelessHotspotModeEnabled(flags dbus.Flags, arg0 dbus.ObjectPath) (arg1 bool, err error) {
	return v.StoreIsWirelessHotspotModeEnabled(
		<-v.GoIsWirelessHotspotModeEnabled(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListDeviceConnections

func (v *network) GoListDeviceConnections(flags dbus.Flags, ch chan *dbus.Call, arg0 dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListDeviceConnections", flags, ch, arg0)
}

func (*network) StoreListDeviceConnections(call *dbus.Call) (arg1 []dbus.ObjectPath, err error) {
	err = call.Store(&arg1)
	return
}

func (v *network) ListDeviceConnections(flags dbus.Flags, arg0 dbus.ObjectPath) (arg1 []dbus.ObjectPath, err error) {
	return v.StoreListDeviceConnections(
		<-v.GoListDeviceConnections(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RegisterSecretReceiver

func (v *network) GoRegisterSecretReceiver(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterSecretReceiver", flags, ch)
}

func (v *network) RegisterSecretReceiver(flags dbus.Flags) error {
	return (<-v.GoRegisterSecretReceiver(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestWirelessScan

func (v *network) GoRequestWirelessScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestWirelessScan", flags, ch)
}

func (v *network) RequestWirelessScan(flags dbus.Flags) error {
	return (<-v.GoRequestWirelessScan(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAutoProxy

func (v *network) GoSetAutoProxy(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoProxy", flags, ch, arg0)
}

func (v *network) SetAutoProxy(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetAutoProxy(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetDeviceManaged

func (v *network) GoSetDeviceManaged(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeviceManaged", flags, ch, arg0, arg1)
}

func (v *network) SetDeviceManaged(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetDeviceManaged(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetProxy

func (v *network) GoSetProxy(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxy", flags, ch, arg0, arg1, arg2)
}

func (v *network) SetProxy(flags dbus.Flags, arg0 string, arg1 string, arg2 string) error {
	return (<-v.GoSetProxy(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetProxyIgnoreHosts

func (v *network) GoSetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyIgnoreHosts", flags, ch, arg0)
}

func (v *network) SetProxyIgnoreHosts(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetProxyIgnoreHosts(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetProxyMethod

func (v *network) GoSetProxyMethod(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxyMethod", flags, ch, arg0)
}

func (v *network) SetProxyMethod(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetProxyMethod(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// signal NeedSecrets

func (v *network) ConnectNeedSecrets(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NeedSecrets", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NeedSecrets",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NeedSecretsFinished

func (v *network) ConnectNeedSecretsFinished(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NeedSecretsFinished", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NeedSecretsFinished",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointAdded

func (v *network) ConnectAccessPointAdded(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointRemoved

func (v *network) ConnectAccessPointRemoved(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AccessPointPropertiesChanged

func (v *network) ConnectAccessPointPropertiesChanged(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceEnabled

func (v *network) ConnectDeviceEnabled(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 bool
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property State u

func (v *network) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
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
