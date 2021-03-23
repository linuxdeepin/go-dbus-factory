package wpa_supplicant1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type WPASupplicant struct {
	wpaSupplicant // interface fi.w1.wpa_supplicant1
	proxy.Object
}

func NewWPASupplicant(conn *dbus.Conn) *WPASupplicant {
	obj := new(WPASupplicant)
	obj.Object.Init_(conn, "fi.w1.wpa_supplicant1", "/fi/w1/wpa_supplicant1")
	return obj
}

type wpaSupplicant struct{}

func (v *wpaSupplicant) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wpaSupplicant) GetInterfaceName_() string {
	return "fi.w1.wpa_supplicant1"
}

// method CreateInterface

func (v *wpaSupplicant) GoCreateInterface(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateInterface", flags, ch, args)
}

func (*wpaSupplicant) StoreCreateInterface(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *wpaSupplicant) CreateInterface(flags dbus.Flags, args map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreCreateInterface(
		<-v.GoCreateInterface(flags, make(chan *dbus.Call, 1), args).Done)
}

// method RemoveInterface

func (v *wpaSupplicant) GoRemoveInterface(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveInterface", flags, ch, path)
}

func (v *wpaSupplicant) RemoveInterface(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoRemoveInterface(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method GetInterface

func (v *wpaSupplicant) GoGetInterface(flags dbus.Flags, ch chan *dbus.Call, ifname string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetInterface", flags, ch, ifname)
}

func (*wpaSupplicant) StoreGetInterface(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *wpaSupplicant) GetInterface(flags dbus.Flags, ifname string) (path dbus.ObjectPath, err error) {
	return v.StoreGetInterface(
		<-v.GoGetInterface(flags, make(chan *dbus.Call, 1), ifname).Done)
}

// method ExpectDisconnect

func (v *wpaSupplicant) GoExpectDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ExpectDisconnect", flags, ch)
}

func (v *wpaSupplicant) ExpectDisconnect(flags dbus.Flags) error {
	return (<-v.GoExpectDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal InterfaceAdded

func (v *wpaSupplicant) ConnectInterfaceAdded(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InterfaceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InterfaceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal InterfaceRemoved

func (v *wpaSupplicant) ConnectInterfaceRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InterfaceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InterfaceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *wpaSupplicant) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property DebugLevel s

func (v *wpaSupplicant) DebugLevel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DebugLevel",
	}
}

// property DebugTimestamp b

func (v *wpaSupplicant) DebugTimestamp() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DebugTimestamp",
	}
}

// property DebugShowKeys b

func (v *wpaSupplicant) DebugShowKeys() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DebugShowKeys",
	}
}

// property Interfaces ao

func (v *wpaSupplicant) Interfaces() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Interfaces",
	}
}

// property EapMethods as

func (v *wpaSupplicant) EapMethods() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "EapMethods",
	}
}

// property Capabilities as

func (v *wpaSupplicant) Capabilities() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Capabilities",
	}
}

// property WFDIEs ay

func (v *wpaSupplicant) WFDIEs() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "WFDIEs",
	}
}

type Interface struct {
	interface1         // interface fi.w1.wpa_supplicant1.Interface
	interfaceWPS       // interface fi.w1.wpa_supplicant1.Interface.WPS
	interfaceP2PDevice // interface fi.w1.wpa_supplicant1.Interface.P2PDevice
	proxy.Object
}

func NewInterface(conn *dbus.Conn, path dbus.ObjectPath) (*Interface, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Interface)
	obj.Object.Init_(conn, "fi.w1.wpa_supplicant1", path)
	return obj, nil
}

func (obj *Interface) Interface() *interface1 {
	return &obj.interface1
}

type interface1 struct{}

func (v *interface1) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*interface1) GetInterfaceName_() string {
	return "fi.w1.wpa_supplicant1.Interface"
}

// method Scan

func (v *interface1) GoScan(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Scan", flags, ch, args)
}

func (v *interface1) Scan(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoScan(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method SignalPoll

func (v *interface1) GoSignalPoll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SignalPoll", flags, ch)
}

func (*interface1) StoreSignalPoll(call *dbus.Call) (args map[string]dbus.Variant, err error) {
	err = call.Store(&args)
	return
}

func (v *interface1) SignalPoll(flags dbus.Flags) (args map[string]dbus.Variant, err error) {
	return v.StoreSignalPoll(
		<-v.GoSignalPoll(flags, make(chan *dbus.Call, 1)).Done)
}

// method Disconnect

func (v *interface1) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *interface1) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AddNetwork

func (v *interface1) GoAddNetwork(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddNetwork", flags, ch, args)
}

func (*interface1) StoreAddNetwork(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *interface1) AddNetwork(flags dbus.Flags, args map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddNetwork(
		<-v.GoAddNetwork(flags, make(chan *dbus.Call, 1), args).Done)
}

// method Reassociate

func (v *interface1) GoReassociate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reassociate", flags, ch)
}

func (v *interface1) Reassociate(flags dbus.Flags) error {
	return (<-v.GoReassociate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reattach

func (v *interface1) GoReattach(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reattach", flags, ch)
}

func (v *interface1) Reattach(flags dbus.Flags) error {
	return (<-v.GoReattach(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reconnect

func (v *interface1) GoReconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reconnect", flags, ch)
}

func (v *interface1) Reconnect(flags dbus.Flags) error {
	return (<-v.GoReconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RemoveNetwork

func (v *interface1) GoRemoveNetwork(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveNetwork", flags, ch, path)
}

func (v *interface1) RemoveNetwork(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoRemoveNetwork(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method RemoveAllNetworks

func (v *interface1) GoRemoveAllNetworks(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAllNetworks", flags, ch)
}

func (v *interface1) RemoveAllNetworks(flags dbus.Flags) error {
	return (<-v.GoRemoveAllNetworks(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SelectNetwork

func (v *interface1) GoSelectNetwork(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SelectNetwork", flags, ch, path)
}

func (v *interface1) SelectNetwork(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoSelectNetwork(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method NetworkReply

func (v *interface1) GoNetworkReply(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath, field string, value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NetworkReply", flags, ch, path, field, value)
}

func (v *interface1) NetworkReply(flags dbus.Flags, path dbus.ObjectPath, field string, value string) error {
	return (<-v.GoNetworkReply(flags, make(chan *dbus.Call, 1), path, field, value).Done).Err
}

// method AddBlob

func (v *interface1) GoAddBlob(flags dbus.Flags, ch chan *dbus.Call, name string, data []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddBlob", flags, ch, name, data)
}

func (v *interface1) AddBlob(flags dbus.Flags, name string, data []uint8) error {
	return (<-v.GoAddBlob(flags, make(chan *dbus.Call, 1), name, data).Done).Err
}

// method GetBlob

func (v *interface1) GoGetBlob(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBlob", flags, ch, name)
}

func (*interface1) StoreGetBlob(call *dbus.Call) (data []uint8, err error) {
	err = call.Store(&data)
	return
}

func (v *interface1) GetBlob(flags dbus.Flags, name string) (data []uint8, err error) {
	return v.StoreGetBlob(
		<-v.GoGetBlob(flags, make(chan *dbus.Call, 1), name).Done)
}

// method RemoveBlob

func (v *interface1) GoRemoveBlob(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveBlob", flags, ch, name)
}

func (v *interface1) RemoveBlob(flags dbus.Flags, name string) error {
	return (<-v.GoRemoveBlob(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetPKCS11EngineAndModulePath

func (v *interface1) GoSetPKCS11EngineAndModulePath(flags dbus.Flags, ch chan *dbus.Call, pkcs11_engine_path string, pkcs11_module_path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPKCS11EngineAndModulePath", flags, ch, pkcs11_engine_path, pkcs11_module_path)
}

func (v *interface1) SetPKCS11EngineAndModulePath(flags dbus.Flags, pkcs11_engine_path string, pkcs11_module_path string) error {
	return (<-v.GoSetPKCS11EngineAndModulePath(flags, make(chan *dbus.Call, 1), pkcs11_engine_path, pkcs11_module_path).Done).Err
}

// method FlushBSS

func (v *interface1) GoFlushBSS(flags dbus.Flags, ch chan *dbus.Call, age uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FlushBSS", flags, ch, age)
}

func (v *interface1) FlushBSS(flags dbus.Flags, age uint32) error {
	return (<-v.GoFlushBSS(flags, make(chan *dbus.Call, 1), age).Done).Err
}

// method SubscribeProbeReq

func (v *interface1) GoSubscribeProbeReq(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SubscribeProbeReq", flags, ch)
}

func (v *interface1) SubscribeProbeReq(flags dbus.Flags) error {
	return (<-v.GoSubscribeProbeReq(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method UnsubscribeProbeReq

func (v *interface1) GoUnsubscribeProbeReq(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnsubscribeProbeReq", flags, ch)
}

func (v *interface1) UnsubscribeProbeReq(flags dbus.Flags) error {
	return (<-v.GoUnsubscribeProbeReq(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method EAPLogoff

func (v *interface1) GoEAPLogoff(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EAPLogoff", flags, ch)
}

func (v *interface1) EAPLogoff(flags dbus.Flags) error {
	return (<-v.GoEAPLogoff(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method EAPLogon

func (v *interface1) GoEAPLogon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EAPLogon", flags, ch)
}

func (v *interface1) EAPLogon(flags dbus.Flags) error {
	return (<-v.GoEAPLogon(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AutoScan

func (v *interface1) GoAutoScan(flags dbus.Flags, ch chan *dbus.Call, arg string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AutoScan", flags, ch, arg)
}

func (v *interface1) AutoScan(flags dbus.Flags, arg string) error {
	return (<-v.GoAutoScan(flags, make(chan *dbus.Call, 1), arg).Done).Err
}

// method TDLSDiscover

func (v *interface1) GoTDLSDiscover(flags dbus.Flags, ch chan *dbus.Call, peer_address string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSDiscover", flags, ch, peer_address)
}

func (v *interface1) TDLSDiscover(flags dbus.Flags, peer_address string) error {
	return (<-v.GoTDLSDiscover(flags, make(chan *dbus.Call, 1), peer_address).Done).Err
}

// method TDLSSetup

func (v *interface1) GoTDLSSetup(flags dbus.Flags, ch chan *dbus.Call, peer_address string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSSetup", flags, ch, peer_address)
}

func (v *interface1) TDLSSetup(flags dbus.Flags, peer_address string) error {
	return (<-v.GoTDLSSetup(flags, make(chan *dbus.Call, 1), peer_address).Done).Err
}

// method TDLSStatus

func (v *interface1) GoTDLSStatus(flags dbus.Flags, ch chan *dbus.Call, peer_address string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSStatus", flags, ch, peer_address)
}

func (*interface1) StoreTDLSStatus(call *dbus.Call) (status string, err error) {
	err = call.Store(&status)
	return
}

func (v *interface1) TDLSStatus(flags dbus.Flags, peer_address string) (status string, err error) {
	return v.StoreTDLSStatus(
		<-v.GoTDLSStatus(flags, make(chan *dbus.Call, 1), peer_address).Done)
}

// method TDLSTeardown

func (v *interface1) GoTDLSTeardown(flags dbus.Flags, ch chan *dbus.Call, peer_address string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSTeardown", flags, ch, peer_address)
}

func (v *interface1) TDLSTeardown(flags dbus.Flags, peer_address string) error {
	return (<-v.GoTDLSTeardown(flags, make(chan *dbus.Call, 1), peer_address).Done).Err
}

// method TDLSChannelSwitch

func (v *interface1) GoTDLSChannelSwitch(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSChannelSwitch", flags, ch, args)
}

func (v *interface1) TDLSChannelSwitch(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoTDLSChannelSwitch(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method TDLSCancelChannelSwitch

func (v *interface1) GoTDLSCancelChannelSwitch(flags dbus.Flags, ch chan *dbus.Call, peer_address string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TDLSCancelChannelSwitch", flags, ch, peer_address)
}

func (v *interface1) TDLSCancelChannelSwitch(flags dbus.Flags, peer_address string) error {
	return (<-v.GoTDLSCancelChannelSwitch(flags, make(chan *dbus.Call, 1), peer_address).Done).Err
}

// method VendorElemAdd

func (v *interface1) GoVendorElemAdd(flags dbus.Flags, ch chan *dbus.Call, frame_id int32, ielems []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VendorElemAdd", flags, ch, frame_id, ielems)
}

func (v *interface1) VendorElemAdd(flags dbus.Flags, frame_id int32, ielems []uint8) error {
	return (<-v.GoVendorElemAdd(flags, make(chan *dbus.Call, 1), frame_id, ielems).Done).Err
}

// method VendorElemGet

func (v *interface1) GoVendorElemGet(flags dbus.Flags, ch chan *dbus.Call, frame_id int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VendorElemGet", flags, ch, frame_id)
}

func (*interface1) StoreVendorElemGet(call *dbus.Call) (ielems []uint8, err error) {
	err = call.Store(&ielems)
	return
}

func (v *interface1) VendorElemGet(flags dbus.Flags, frame_id int32) (ielems []uint8, err error) {
	return v.StoreVendorElemGet(
		<-v.GoVendorElemGet(flags, make(chan *dbus.Call, 1), frame_id).Done)
}

// method VendorElemRem

func (v *interface1) GoVendorElemRem(flags dbus.Flags, ch chan *dbus.Call, frame_id int32, ielems []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VendorElemRem", flags, ch, frame_id, ielems)
}

func (v *interface1) VendorElemRem(flags dbus.Flags, frame_id int32, ielems []uint8) error {
	return (<-v.GoVendorElemRem(flags, make(chan *dbus.Call, 1), frame_id, ielems).Done).Err
}

// method SaveConfig

func (v *interface1) GoSaveConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveConfig", flags, ch)
}

func (v *interface1) SaveConfig(flags dbus.Flags) error {
	return (<-v.GoSaveConfig(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AbortScan

func (v *interface1) GoAbortScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AbortScan", flags, ch)
}

func (v *interface1) AbortScan(flags dbus.Flags) error {
	return (<-v.GoAbortScan(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal ScanDone

func (v *interface1) ConnectScanDone(cb func(success bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ScanDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ScanDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var success bool
		err := dbus.Store(sig.Body, &success)
		if err == nil {
			cb(success)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BSSAdded

func (v *interface1) ConnectBSSAdded(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BSSAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BSSAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BSSRemoved

func (v *interface1) ConnectBSSRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BSSRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BSSRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BlobAdded

func (v *interface1) ConnectBlobAdded(cb func(name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BlobAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BlobAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		err := dbus.Store(sig.Body, &name)
		if err == nil {
			cb(name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BlobRemoved

func (v *interface1) ConnectBlobRemoved(cb func(name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BlobRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BlobRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		err := dbus.Store(sig.Body, &name)
		if err == nil {
			cb(name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NetworkAdded

func (v *interface1) ConnectNetworkAdded(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NetworkAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NetworkAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NetworkRemoved

func (v *interface1) ConnectNetworkRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NetworkRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NetworkRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NetworkSelected

func (v *interface1) ConnectNetworkSelected(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NetworkSelected", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NetworkSelected",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *interface1) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProbeRequest

func (v *interface1) ConnectProbeRequest(cb func(args map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProbeRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProbeRequest",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var args map[string]dbus.Variant
		err := dbus.Store(sig.Body, &args)
		if err == nil {
			cb(args)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Certification

func (v *interface1) ConnectCertification(cb func(certification map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Certification", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Certification",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var certification map[string]dbus.Variant
		err := dbus.Store(sig.Body, &certification)
		if err == nil {
			cb(certification)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal EAP

func (v *interface1) ConnectEAP(cb func(status string, parameter string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EAP", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EAP",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var parameter string
		err := dbus.Store(sig.Body, &status, &parameter)
		if err == nil {
			cb(status, parameter)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StaAuthorized

func (v *interface1) ConnectStaAuthorized(cb func(name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StaAuthorized", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StaAuthorized",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		err := dbus.Store(sig.Body, &name)
		if err == nil {
			cb(name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StaDeauthorized

func (v *interface1) ConnectStaDeauthorized(cb func(name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StaDeauthorized", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StaDeauthorized",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		err := dbus.Store(sig.Body, &name)
		if err == nil {
			cb(name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StationAdded

func (v *interface1) ConnectStationAdded(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StationAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StationAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StationRemoved

func (v *interface1) ConnectStationRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StationRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StationRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NetworkRequest

func (v *interface1) ConnectNetworkRequest(cb func(path dbus.ObjectPath, field string, text string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NetworkRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NetworkRequest",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var field string
		var text string
		err := dbus.Store(sig.Body, &path, &field, &text)
		if err == nil {
			cb(path, field, text)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Capabilities a{sv}

func (v *interface1) Capabilities() MapStrVariant {
	return MapStrVariant{
		Impl: v,
		Name: "Capabilities",
	}
}

// property State s

func (v *interface1) State() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "State",
	}
}

// property Scanning b

func (v *interface1) Scanning() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Scanning",
	}
}

// property ApScan u

func (v *interface1) ApScan() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ApScan",
	}
}

// property BSSExpireAge u

func (v *interface1) BSSExpireAge() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "BSSExpireAge",
	}
}

// property BSSExpireCount u

func (v *interface1) BSSExpireCount() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "BSSExpireCount",
	}
}

// property Country s

func (v *interface1) Country() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Country",
	}
}

// property Ifname s

func (v *interface1) Ifname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Ifname",
	}
}

// property Driver s

func (v *interface1) Driver() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Driver",
	}
}

// property BridgeIfname s

func (v *interface1) BridgeIfname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BridgeIfname",
	}
}

// property ConfigFile s

func (v *interface1) ConfigFile() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ConfigFile",
	}
}

// property CurrentBSS o

func (v *interface1) CurrentBSS() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "CurrentBSS",
	}
}

// property CurrentNetwork o

func (v *interface1) CurrentNetwork() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "CurrentNetwork",
	}
}

// property CurrentAuthMode s

func (v *interface1) CurrentAuthMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CurrentAuthMode",
	}
}

// property Blobs a{say}

func (v *interface1) Blobs() PropInterfaceBlobs {
	return PropInterfaceBlobs{
		Impl: v,
	}
}

type PropInterfaceBlobs struct {
	Impl proxy.Implementer
}

func (p PropInterfaceBlobs) Get(flags dbus.Flags) (value map[string][]byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Blobs", &value)
	return
}

func (p PropInterfaceBlobs) ConnectChanged(cb func(hasValue bool, value map[string][]byte)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string][]byte
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Blobs", cb0)
}

// property BSSs ao

func (v *interface1) BSSs() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "BSSs",
	}
}

// property Networks ao

func (v *interface1) Networks() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Networks",
	}
}

// property FastReauth b

func (v *interface1) FastReauth() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "FastReauth",
	}
}

// property ScanInterval i

func (v *interface1) ScanInterval() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "ScanInterval",
	}
}

// property PKCS11EnginePath s

func (v *interface1) PKCS11EnginePath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PKCS11EnginePath",
	}
}

// property PKCS11ModulePath s

func (v *interface1) PKCS11ModulePath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PKCS11ModulePath",
	}
}

// property DisconnectReason i

func (v *interface1) DisconnectReason() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DisconnectReason",
	}
}

// property AuthStatusCode i

func (v *interface1) AuthStatusCode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "AuthStatusCode",
	}
}

// property AssocStatusCode i

func (v *interface1) AssocStatusCode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "AssocStatusCode",
	}
}

// property Stations ao

func (v *interface1) Stations() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Stations",
	}
}

// property CtrlInterface s

func (v *interface1) CtrlInterface() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CtrlInterface",
	}
}

// property CtrlInterfaceGroup s

func (v *interface1) CtrlInterfaceGroup() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CtrlInterfaceGroup",
	}
}

// property EapolVersion s

func (v *interface1) EapolVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "EapolVersion",
	}
}

// property Bgscan s

func (v *interface1) Bgscan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Bgscan",
	}
}

// property DisableScanOffload s

func (v *interface1) DisableScanOffload() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DisableScanOffload",
	}
}

// property OpenscEnginePath s

func (v *interface1) OpenscEnginePath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OpenscEnginePath",
	}
}

// property OpensslCiphers s

func (v *interface1) OpensslCiphers() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OpensslCiphers",
	}
}

// property PcscReader s

func (v *interface1) PcscReader() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PcscReader",
	}
}

// property PcscPin s

func (v *interface1) PcscPin() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PcscPin",
	}
}

// property ExternalSim s

func (v *interface1) ExternalSim() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ExternalSim",
	}
}

// property DriverParam s

func (v *interface1) DriverParam() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DriverParam",
	}
}

// property Dot11RSNAConfigPMKLifetime s

func (v *interface1) Dot11RSNAConfigPMKLifetime() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Dot11RSNAConfigPMKLifetime",
	}
}

// property Dot11RSNAConfigPMKReauthThreshold s

func (v *interface1) Dot11RSNAConfigPMKReauthThreshold() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Dot11RSNAConfigPMKReauthThreshold",
	}
}

// property Dot11RSNAConfigSATimeout s

func (v *interface1) Dot11RSNAConfigSATimeout() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Dot11RSNAConfigSATimeout",
	}
}

// property UpdateConfig s

func (v *interface1) UpdateConfig() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UpdateConfig",
	}
}

// property Uuid s

func (v *interface1) Uuid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Uuid",
	}
}

// property AutoUuid s

func (v *interface1) AutoUuid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "AutoUuid",
	}
}

// property DeviceName s

func (v *interface1) DeviceName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DeviceName",
	}
}

// property Manufacturer s

func (v *interface1) Manufacturer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property ModelName s

func (v *interface1) ModelName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ModelName",
	}
}

// property ModelNumber s

func (v *interface1) ModelNumber() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ModelNumber",
	}
}

// property SerialNumber s

func (v *interface1) SerialNumber() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SerialNumber",
	}
}

// property DeviceType s

func (v *interface1) DeviceType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DeviceType",
	}
}

// property OsVersion s

func (v *interface1) OsVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OsVersion",
	}
}

// property ConfigMethods s

func (v *interface1) ConfigMethods() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ConfigMethods",
	}
}

// property WpsCredProcessing s

func (v *interface1) WpsCredProcessing() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsCredProcessing",
	}
}

// property WpsVendorExtM1 s

func (v *interface1) WpsVendorExtM1() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsVendorExtM1",
	}
}

// property SecDeviceType s

func (v *interface1) SecDeviceType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SecDeviceType",
	}
}

// property P2pListenRegClass s

func (v *interface1) P2pListenRegClass() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pListenRegClass",
	}
}

// property P2pListenChannel s

func (v *interface1) P2pListenChannel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pListenChannel",
	}
}

// property P2pOperRegClass s

func (v *interface1) P2pOperRegClass() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pOperRegClass",
	}
}

// property P2pOperChannel s

func (v *interface1) P2pOperChannel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pOperChannel",
	}
}

// property P2pGoIntent s

func (v *interface1) P2pGoIntent() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoIntent",
	}
}

// property P2pSsidPostfix s

func (v *interface1) P2pSsidPostfix() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pSsidPostfix",
	}
}

// property PersistentReconnect s

func (v *interface1) PersistentReconnect() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PersistentReconnect",
	}
}

// property P2pIntraBss s

func (v *interface1) P2pIntraBss() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pIntraBss",
	}
}

// property P2pGroupIdle s

func (v *interface1) P2pGroupIdle() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGroupIdle",
	}
}

// property P2pGoFreqChangePolicy s

func (v *interface1) P2pGoFreqChangePolicy() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoFreqChangePolicy",
	}
}

// property P2pPassphraseLen s

func (v *interface1) P2pPassphraseLen() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pPassphraseLen",
	}
}

// property P2pPrefChan s

func (v *interface1) P2pPrefChan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pPrefChan",
	}
}

// property P2pNoGoFreq s

func (v *interface1) P2pNoGoFreq() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pNoGoFreq",
	}
}

// property P2pAddCliChan s

func (v *interface1) P2pAddCliChan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pAddCliChan",
	}
}

// property P2pOptimizeListenChan s

func (v *interface1) P2pOptimizeListenChan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pOptimizeListenChan",
	}
}

// property P2pGoHt40 s

func (v *interface1) P2pGoHt40() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoHt40",
	}
}

// property P2pGoVht s

func (v *interface1) P2pGoVht() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoVht",
	}
}

// property P2pGoHe s

func (v *interface1) P2pGoHe() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoHe",
	}
}

// property P2pDisabled s

func (v *interface1) P2pDisabled() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pDisabled",
	}
}

// property P2pGoCtwindow s

func (v *interface1) P2pGoCtwindow() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoCtwindow",
	}
}

// property P2pNoGroupIface s

func (v *interface1) P2pNoGroupIface() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pNoGroupIface",
	}
}

// property P2pIgnoreSharedFreq s

func (v *interface1) P2pIgnoreSharedFreq() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pIgnoreSharedFreq",
	}
}

// property IpAddrGo s

func (v *interface1) IpAddrGo() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IpAddrGo",
	}
}

// property IpAddrMask s

func (v *interface1) IpAddrMask() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IpAddrMask",
	}
}

// property IpAddrStart s

func (v *interface1) IpAddrStart() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IpAddrStart",
	}
}

// property IpAddrEnd s

func (v *interface1) IpAddrEnd() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IpAddrEnd",
	}
}

// property P2pCliProbe s

func (v *interface1) P2pCliProbe() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pCliProbe",
	}
}

// property P2pDeviceRandomMacAddr s

func (v *interface1) P2pDeviceRandomMacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pDeviceRandomMacAddr",
	}
}

// property P2pDevicePersistentMacAddr s

func (v *interface1) P2pDevicePersistentMacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pDevicePersistentMacAddr",
	}
}

// property P2pInterfaceRandomMacAddr s

func (v *interface1) P2pInterfaceRandomMacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pInterfaceRandomMacAddr",
	}
}

// property BssMaxCount s

func (v *interface1) BssMaxCount() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BssMaxCount",
	}
}

// property FilterSsids s

func (v *interface1) FilterSsids() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FilterSsids",
	}
}

// property FilterRssi s

func (v *interface1) FilterRssi() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FilterRssi",
	}
}

// property MaxNumSta s

func (v *interface1) MaxNumSta() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MaxNumSta",
	}
}

// property ApIsolate s

func (v *interface1) ApIsolate() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ApIsolate",
	}
}

// property DisassocLowAck s

func (v *interface1) DisassocLowAck() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DisassocLowAck",
	}
}

// property Hs20 s

func (v *interface1) Hs20() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Hs20",
	}
}

// property Interworking s

func (v *interface1) Interworking() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Interworking",
	}
}

// property Hessid s

func (v *interface1) Hessid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Hessid",
	}
}

// property AccessNetworkType s

func (v *interface1) AccessNetworkType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "AccessNetworkType",
	}
}

// property GoInterworking s

func (v *interface1) GoInterworking() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GoInterworking",
	}
}

// property GoAccessNetworkType s

func (v *interface1) GoAccessNetworkType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GoAccessNetworkType",
	}
}

// property GoInternet s

func (v *interface1) GoInternet() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GoInternet",
	}
}

// property GoVenueGroup s

func (v *interface1) GoVenueGroup() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GoVenueGroup",
	}
}

// property GoVenueType s

func (v *interface1) GoVenueType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GoVenueType",
	}
}

// property PbcInM1 s

func (v *interface1) PbcInM1() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PbcInM1",
	}
}

// property Autoscan s

func (v *interface1) Autoscan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Autoscan",
	}
}

// property WpsNfcDevPwId s

func (v *interface1) WpsNfcDevPwId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsNfcDevPwId",
	}
}

// property WpsNfcDhPubkey s

func (v *interface1) WpsNfcDhPubkey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsNfcDhPubkey",
	}
}

// property WpsNfcDhPrivkey s

func (v *interface1) WpsNfcDhPrivkey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsNfcDhPrivkey",
	}
}

// property WpsNfcDevPw s

func (v *interface1) WpsNfcDevPw() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsNfcDevPw",
	}
}

// property ExtPasswordBackend s

func (v *interface1) ExtPasswordBackend() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ExtPasswordBackend",
	}
}

// property P2pGoMaxInactivity s

func (v *interface1) P2pGoMaxInactivity() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pGoMaxInactivity",
	}
}

// property AutoInterworking s

func (v *interface1) AutoInterworking() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "AutoInterworking",
	}
}

// property Okc s

func (v *interface1) Okc() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Okc",
	}
}

// property Pmf s

func (v *interface1) Pmf() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Pmf",
	}
}

// property SaeGroups s

func (v *interface1) SaeGroups() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SaeGroups",
	}
}

// property DtimPeriod s

func (v *interface1) DtimPeriod() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DtimPeriod",
	}
}

// property BeaconInt s

func (v *interface1) BeaconInt() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BeaconInt",
	}
}

// property ApVendorElements s

func (v *interface1) ApVendorElements() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ApVendorElements",
	}
}

// property IgnoreOldScanRes s

func (v *interface1) IgnoreOldScanRes() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IgnoreOldScanRes",
	}
}

// property FreqList s

func (v *interface1) FreqList() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FreqList",
	}
}

// property ScanCurFreq s

func (v *interface1) ScanCurFreq() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ScanCurFreq",
	}
}

// property SchedScanInterval s

func (v *interface1) SchedScanInterval() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SchedScanInterval",
	}
}

// property SchedScanStartDelay s

func (v *interface1) SchedScanStartDelay() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SchedScanStartDelay",
	}
}

// property TdlsExternalControl s

func (v *interface1) TdlsExternalControl() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "TdlsExternalControl",
	}
}

// property OsuDir s

func (v *interface1) OsuDir() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OsuDir",
	}
}

// property WowlanTriggers s

func (v *interface1) WowlanTriggers() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WowlanTriggers",
	}
}

// property P2pSearchDelay s

func (v *interface1) P2pSearchDelay() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2pSearchDelay",
	}
}

// property MacAddr s

func (v *interface1) MacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MacAddr",
	}
}

// property RandAddrLifetime s

func (v *interface1) RandAddrLifetime() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RandAddrLifetime",
	}
}

// property PreassocMacAddr s

func (v *interface1) PreassocMacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PreassocMacAddr",
	}
}

// property KeyMgmtOffload s

func (v *interface1) KeyMgmtOffload() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KeyMgmtOffload",
	}
}

// property PassiveScan s

func (v *interface1) PassiveScan() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PassiveScan",
	}
}

// property ReassocSameBssOptim s

func (v *interface1) ReassocSameBssOptim() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ReassocSameBssOptim",
	}
}

// property WpsPriority s

func (v *interface1) WpsPriority() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpsPriority",
	}
}

// property FstGroupId s

func (v *interface1) FstGroupId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FstGroupId",
	}
}

// property FstPriority s

func (v *interface1) FstPriority() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FstPriority",
	}
}

// property FstLlt s

func (v *interface1) FstLlt() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FstLlt",
	}
}

// property CertInCb s

func (v *interface1) CertInCb() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CertInCb",
	}
}

// property WpaRscRelaxation s

func (v *interface1) WpaRscRelaxation() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WpaRscRelaxation",
	}
}

// property SchedScanPlans s

func (v *interface1) SchedScanPlans() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SchedScanPlans",
	}
}

// property GasAddress3 s

func (v *interface1) GasAddress3() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GasAddress3",
	}
}

// property FtmResponder s

func (v *interface1) FtmResponder() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FtmResponder",
	}
}

// property FtmInitiator s

func (v *interface1) FtmInitiator() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FtmInitiator",
	}
}

// property GasRandAddrLifetime s

func (v *interface1) GasRandAddrLifetime() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GasRandAddrLifetime",
	}
}

// property GasRandMacAddr s

func (v *interface1) GasRandMacAddr() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GasRandMacAddr",
	}
}

// property DppConfigProcessing s

func (v *interface1) DppConfigProcessing() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DppConfigProcessing",
	}
}

// property ColocIntfReporting s

func (v *interface1) ColocIntfReporting() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ColocIntfReporting",
	}
}

func (obj *Interface) WPS() *interfaceWPS {
	return &obj.interfaceWPS
}

type interfaceWPS struct{}

func (v *interfaceWPS) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*interfaceWPS) GetInterfaceName_() string {
	return "fi.w1.wpa_supplicant1.Interface.WPS"
}

// method Start

func (v *interfaceWPS) GoStart(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start", flags, ch, args)
}

func (*interfaceWPS) StoreStart(call *dbus.Call) (output map[string]dbus.Variant, err error) {
	err = call.Store(&output)
	return
}

func (v *interfaceWPS) Start(flags dbus.Flags, args map[string]dbus.Variant) (output map[string]dbus.Variant, err error) {
	return v.StoreStart(
		<-v.GoStart(flags, make(chan *dbus.Call, 1), args).Done)
}

// method Cancel

func (v *interfaceWPS) GoCancel(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Cancel", flags, ch)
}

func (v *interfaceWPS) Cancel(flags dbus.Flags) error {
	return (<-v.GoCancel(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Event

func (v *interfaceWPS) ConnectEvent(cb func(name string, args map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Event", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Event",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		var args map[string]dbus.Variant
		err := dbus.Store(sig.Body, &name, &args)
		if err == nil {
			cb(name, args)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Credentials

func (v *interfaceWPS) ConnectCredentials(cb func(credentials map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Credentials", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Credentials",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var credentials map[string]dbus.Variant
		err := dbus.Store(sig.Body, &credentials)
		if err == nil {
			cb(credentials)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *interfaceWPS) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ProcessCredentials b

func (v *interfaceWPS) ProcessCredentials() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ProcessCredentials",
	}
}

// property ConfigMethods s

func (v *interfaceWPS) ConfigMethods() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ConfigMethods",
	}
}

// property DeviceName s

func (v *interfaceWPS) DeviceName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DeviceName",
	}
}

// property Manufacturer s

func (v *interfaceWPS) Manufacturer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property ModelName s

func (v *interfaceWPS) ModelName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ModelName",
	}
}

// property ModelNumber s

func (v *interfaceWPS) ModelNumber() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ModelNumber",
	}
}

// property SerialNumber s

func (v *interfaceWPS) SerialNumber() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SerialNumber",
	}
}

// property DeviceType ay

func (v *interfaceWPS) DeviceType() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "DeviceType",
	}
}

func (obj *Interface) P2PDevice() *interfaceP2PDevice {
	return &obj.interfaceP2PDevice
}

type interfaceP2PDevice struct{}

func (v *interfaceP2PDevice) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*interfaceP2PDevice) GetInterfaceName_() string {
	return "fi.w1.wpa_supplicant1.Interface.P2PDevice"
}

// method Find

func (v *interfaceP2PDevice) GoFind(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Find", flags, ch, args)
}

func (v *interfaceP2PDevice) Find(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoFind(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method StopFind

func (v *interfaceP2PDevice) GoStopFind(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopFind", flags, ch)
}

func (v *interfaceP2PDevice) StopFind(flags dbus.Flags) error {
	return (<-v.GoStopFind(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Listen

func (v *interfaceP2PDevice) GoListen(flags dbus.Flags, ch chan *dbus.Call, timeout int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Listen", flags, ch, timeout)
}

func (v *interfaceP2PDevice) Listen(flags dbus.Flags, timeout int32) error {
	return (<-v.GoListen(flags, make(chan *dbus.Call, 1), timeout).Done).Err
}

// method ExtendedListen

func (v *interfaceP2PDevice) GoExtendedListen(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ExtendedListen", flags, ch, args)
}

func (v *interfaceP2PDevice) ExtendedListen(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoExtendedListen(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method PresenceRequest

func (v *interfaceP2PDevice) GoPresenceRequest(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresenceRequest", flags, ch, args)
}

func (v *interfaceP2PDevice) PresenceRequest(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoPresenceRequest(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method ProvisionDiscoveryRequest

func (v *interfaceP2PDevice) GoProvisionDiscoveryRequest(flags dbus.Flags, ch chan *dbus.Call, peer dbus.ObjectPath, config_method string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ProvisionDiscoveryRequest", flags, ch, peer, config_method)
}

func (v *interfaceP2PDevice) ProvisionDiscoveryRequest(flags dbus.Flags, peer dbus.ObjectPath, config_method string) error {
	return (<-v.GoProvisionDiscoveryRequest(flags, make(chan *dbus.Call, 1), peer, config_method).Done).Err
}

// method Connect

func (v *interfaceP2PDevice) GoConnect(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch, args)
}

func (*interfaceP2PDevice) StoreConnect(call *dbus.Call) (generated_pin string, err error) {
	err = call.Store(&generated_pin)
	return
}

func (v *interfaceP2PDevice) Connect(flags dbus.Flags, args map[string]dbus.Variant) (generated_pin string, err error) {
	return v.StoreConnect(
		<-v.GoConnect(flags, make(chan *dbus.Call, 1), args).Done)
}

// method GroupAdd

func (v *interfaceP2PDevice) GoGroupAdd(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GroupAdd", flags, ch, args)
}

func (v *interfaceP2PDevice) GroupAdd(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoGroupAdd(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method Cancel

func (v *interfaceP2PDevice) GoCancel(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Cancel", flags, ch)
}

func (v *interfaceP2PDevice) Cancel(flags dbus.Flags) error {
	return (<-v.GoCancel(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Invite

func (v *interfaceP2PDevice) GoInvite(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Invite", flags, ch, args)
}

func (v *interfaceP2PDevice) Invite(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoInvite(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method Disconnect

func (v *interfaceP2PDevice) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *interfaceP2PDevice) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RejectPeer

func (v *interfaceP2PDevice) GoRejectPeer(flags dbus.Flags, ch chan *dbus.Call, peer dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RejectPeer", flags, ch, peer)
}

func (v *interfaceP2PDevice) RejectPeer(flags dbus.Flags, peer dbus.ObjectPath) error {
	return (<-v.GoRejectPeer(flags, make(chan *dbus.Call, 1), peer).Done).Err
}

// method RemoveClient

func (v *interfaceP2PDevice) GoRemoveClient(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveClient", flags, ch, args)
}

func (v *interfaceP2PDevice) RemoveClient(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoRemoveClient(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method Flush

func (v *interfaceP2PDevice) GoFlush(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Flush", flags, ch)
}

func (v *interfaceP2PDevice) Flush(flags dbus.Flags) error {
	return (<-v.GoFlush(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AddService

func (v *interfaceP2PDevice) GoAddService(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddService", flags, ch, args)
}

func (v *interfaceP2PDevice) AddService(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoAddService(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method DeleteService

func (v *interfaceP2PDevice) GoDeleteService(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteService", flags, ch, args)
}

func (v *interfaceP2PDevice) DeleteService(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoDeleteService(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method FlushService

func (v *interfaceP2PDevice) GoFlushService(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FlushService", flags, ch)
}

func (v *interfaceP2PDevice) FlushService(flags dbus.Flags) error {
	return (<-v.GoFlushService(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ServiceDiscoveryRequest

func (v *interfaceP2PDevice) GoServiceDiscoveryRequest(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ServiceDiscoveryRequest", flags, ch, args)
}

func (*interfaceP2PDevice) StoreServiceDiscoveryRequest(call *dbus.Call) (ref uint64, err error) {
	err = call.Store(&ref)
	return
}

func (v *interfaceP2PDevice) ServiceDiscoveryRequest(flags dbus.Flags, args map[string]dbus.Variant) (ref uint64, err error) {
	return v.StoreServiceDiscoveryRequest(
		<-v.GoServiceDiscoveryRequest(flags, make(chan *dbus.Call, 1), args).Done)
}

// method ServiceDiscoveryResponse

func (v *interfaceP2PDevice) GoServiceDiscoveryResponse(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ServiceDiscoveryResponse", flags, ch, args)
}

func (v *interfaceP2PDevice) ServiceDiscoveryResponse(flags dbus.Flags, args map[string]dbus.Variant) error {
	return (<-v.GoServiceDiscoveryResponse(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method ServiceDiscoveryCancelRequest

func (v *interfaceP2PDevice) GoServiceDiscoveryCancelRequest(flags dbus.Flags, ch chan *dbus.Call, args uint64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ServiceDiscoveryCancelRequest", flags, ch, args)
}

func (v *interfaceP2PDevice) ServiceDiscoveryCancelRequest(flags dbus.Flags, args uint64) error {
	return (<-v.GoServiceDiscoveryCancelRequest(flags, make(chan *dbus.Call, 1), args).Done).Err
}

// method ServiceUpdate

func (v *interfaceP2PDevice) GoServiceUpdate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ServiceUpdate", flags, ch)
}

func (v *interfaceP2PDevice) ServiceUpdate(flags dbus.Flags) error {
	return (<-v.GoServiceUpdate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ServiceDiscoveryExternal

func (v *interfaceP2PDevice) GoServiceDiscoveryExternal(flags dbus.Flags, ch chan *dbus.Call, arg int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ServiceDiscoveryExternal", flags, ch, arg)
}

func (v *interfaceP2PDevice) ServiceDiscoveryExternal(flags dbus.Flags, arg int32) error {
	return (<-v.GoServiceDiscoveryExternal(flags, make(chan *dbus.Call, 1), arg).Done).Err
}

// method AddPersistentGroup

func (v *interfaceP2PDevice) GoAddPersistentGroup(flags dbus.Flags, ch chan *dbus.Call, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddPersistentGroup", flags, ch, args)
}

func (*interfaceP2PDevice) StoreAddPersistentGroup(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *interfaceP2PDevice) AddPersistentGroup(flags dbus.Flags, args map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddPersistentGroup(
		<-v.GoAddPersistentGroup(flags, make(chan *dbus.Call, 1), args).Done)
}

// method RemovePersistentGroup

func (v *interfaceP2PDevice) GoRemovePersistentGroup(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemovePersistentGroup", flags, ch, path)
}

func (v *interfaceP2PDevice) RemovePersistentGroup(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoRemovePersistentGroup(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method RemoveAllPersistentGroups

func (v *interfaceP2PDevice) GoRemoveAllPersistentGroups(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAllPersistentGroups", flags, ch)
}

func (v *interfaceP2PDevice) RemoveAllPersistentGroups(flags dbus.Flags) error {
	return (<-v.GoRemoveAllPersistentGroups(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal DeviceFound

func (v *interfaceP2PDevice) ConnectDeviceFound(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceFound", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceFound",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceFoundProperties

func (v *interfaceP2PDevice) ConnectDeviceFoundProperties(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceFoundProperties", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceFoundProperties",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceLost

func (v *interfaceP2PDevice) ConnectDeviceLost(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceLost", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceLost",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal FindStopped

func (v *interfaceP2PDevice) ConnectFindStopped(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "FindStopped", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".FindStopped",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryRequestDisplayPin

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryRequestDisplayPin(cb func(peer_object dbus.ObjectPath, pin string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryRequestDisplayPin", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryRequestDisplayPin",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		var pin string
		err := dbus.Store(sig.Body, &peer_object, &pin)
		if err == nil {
			cb(peer_object, pin)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryResponseDisplayPin

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryResponseDisplayPin(cb func(peer_object dbus.ObjectPath, pin string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryResponseDisplayPin", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryResponseDisplayPin",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		var pin string
		err := dbus.Store(sig.Body, &peer_object, &pin)
		if err == nil {
			cb(peer_object, pin)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryRequestEnterPin

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryRequestEnterPin(cb func(peer_object dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryRequestEnterPin", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryRequestEnterPin",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		err := dbus.Store(sig.Body, &peer_object)
		if err == nil {
			cb(peer_object)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryResponseEnterPin

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryResponseEnterPin(cb func(peer_object dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryResponseEnterPin", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryResponseEnterPin",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		err := dbus.Store(sig.Body, &peer_object)
		if err == nil {
			cb(peer_object)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryPBCRequest

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryPBCRequest(cb func(peer_object dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryPBCRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryPBCRequest",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		err := dbus.Store(sig.Body, &peer_object)
		if err == nil {
			cb(peer_object)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryPBCResponse

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryPBCResponse(cb func(peer_object dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryPBCResponse", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryPBCResponse",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		err := dbus.Store(sig.Body, &peer_object)
		if err == nil {
			cb(peer_object)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProvisionDiscoveryFailure

func (v *interfaceP2PDevice) ConnectProvisionDiscoveryFailure(cb func(peer_object dbus.ObjectPath, status int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscoveryFailure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscoveryFailure",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var peer_object dbus.ObjectPath
		var status int32
		err := dbus.Store(sig.Body, &peer_object, &status)
		if err == nil {
			cb(peer_object, status)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GroupStarted

func (v *interfaceP2PDevice) ConnectGroupStarted(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GroupStarted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GroupStarted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GroupFormationFailure

func (v *interfaceP2PDevice) ConnectGroupFormationFailure(cb func(reason string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GroupFormationFailure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GroupFormationFailure",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var reason string
		err := dbus.Store(sig.Body, &reason)
		if err == nil {
			cb(reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GONegotiationSuccess

func (v *interfaceP2PDevice) ConnectGONegotiationSuccess(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GONegotiationSuccess", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GONegotiationSuccess",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GONegotiationFailure

func (v *interfaceP2PDevice) ConnectGONegotiationFailure(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GONegotiationFailure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GONegotiationFailure",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GONegotiationRequest

func (v *interfaceP2PDevice) ConnectGONegotiationRequest(cb func(path dbus.ObjectPath, dev_passwd_id uint16, device_go_intent uint8)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GONegotiationRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GONegotiationRequest",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var dev_passwd_id uint16
		var device_go_intent uint8
		err := dbus.Store(sig.Body, &path, &dev_passwd_id, &device_go_intent)
		if err == nil {
			cb(path, dev_passwd_id, device_go_intent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal InvitationResult

func (v *interfaceP2PDevice) ConnectInvitationResult(cb func(invite_result map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InvitationResult", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InvitationResult",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var invite_result map[string]dbus.Variant
		err := dbus.Store(sig.Body, &invite_result)
		if err == nil {
			cb(invite_result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GroupFinished

func (v *interfaceP2PDevice) ConnectGroupFinished(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GroupFinished", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GroupFinished",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ServiceDiscoveryRequest

func (v *interfaceP2PDevice) ConnectServiceDiscoveryRequest(cb func(sd_request map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ServiceDiscoveryRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ServiceDiscoveryRequest",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sd_request map[string]dbus.Variant
		err := dbus.Store(sig.Body, &sd_request)
		if err == nil {
			cb(sd_request)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ServiceDiscoveryResponse

func (v *interfaceP2PDevice) ConnectServiceDiscoveryResponse(cb func(sd_response map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ServiceDiscoveryResponse", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ServiceDiscoveryResponse",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sd_response map[string]dbus.Variant
		err := dbus.Store(sig.Body, &sd_response)
		if err == nil {
			cb(sd_response)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PersistentGroupAdded

func (v *interfaceP2PDevice) ConnectPersistentGroupAdded(cb func(path dbus.ObjectPath, properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PersistentGroupAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PersistentGroupAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &path, &properties)
		if err == nil {
			cb(path, properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PersistentGroupRemoved

func (v *interfaceP2PDevice) ConnectPersistentGroupRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PersistentGroupRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PersistentGroupRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WpsFailed

func (v *interfaceP2PDevice) ConnectWpsFailed(cb func(name string, args map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WpsFailed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WpsFailed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		var args map[string]dbus.Variant
		err := dbus.Store(sig.Body, &name, &args)
		if err == nil {
			cb(name, args)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal InvitationReceived

func (v *interfaceP2PDevice) ConnectInvitationReceived(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InvitationReceived", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InvitationReceived",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property P2PDeviceConfig a{sv}

func (v *interfaceP2PDevice) P2PDeviceConfig() MapStrVariant {
	return MapStrVariant{
		Impl: v,
		Name: "P2PDeviceConfig",
	}
}

// property Peers ao

func (v *interfaceP2PDevice) Peers() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Peers",
	}
}

// property Role s

func (v *interfaceP2PDevice) Role() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Role",
	}
}

// property Group o

func (v *interfaceP2PDevice) Group() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Group",
	}
}

// property PeerGO o

func (v *interfaceP2PDevice) PeerGO() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "PeerGO",
	}
}

// property PersistentGroups ao

func (v *interfaceP2PDevice) PersistentGroups() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "PersistentGroups",
	}
}

type BSS struct {
	bss // interface fi.w1.wpa_supplicant1.BSS
	proxy.Object
}

func NewBSS(conn *dbus.Conn, path dbus.ObjectPath) (*BSS, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(BSS)
	obj.Object.Init_(conn, "fi.w1.wpa_supplicant1", path)
	return obj, nil
}

type bss struct{}

func (v *bss) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*bss) GetInterfaceName_() string {
	return "fi.w1.wpa_supplicant1.BSS"
}

// signal PropertiesChanged

func (v *bss) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var properties map[string]dbus.Variant
		err := dbus.Store(sig.Body, &properties)
		if err == nil {
			cb(properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property SSID ay

func (v *bss) SSID() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "SSID",
	}
}

// property BSSID ay

func (v *bss) BSSID() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "BSSID",
	}
}

// property Privacy b

func (v *bss) Privacy() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Privacy",
	}
}

// property Mode s

func (v *bss) Mode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Mode",
	}
}

// property Signal n

func (v *bss) Signal() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "Signal",
	}
}

// property Frequency q

func (v *bss) Frequency() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Frequency",
	}
}

// property Rates au

func (v *bss) Rates() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "Rates",
	}
}

// property WPA a{sv}

func (v *bss) WPA() MapStrVariant {
	return MapStrVariant{
		Impl: v,
		Name: "WPA",
	}
}

// property RSN a{sv}

func (v *bss) RSN() MapStrVariant {
	return MapStrVariant{
		Impl: v,
		Name: "RSN",
	}
}

// property WPS a{sv}

func (v *bss) WPS() MapStrVariant {
	return MapStrVariant{
		Impl: v,
		Name: "WPS",
	}
}

// property IEs ay

func (v *bss) IEs() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "IEs",
	}
}

// property Age u

func (v *bss) Age() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Age",
	}
}

type MapStrVariant struct {
	Impl proxy.Implementer
	Name string
}

func (p MapStrVariant) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p MapStrVariant) Set(flags dbus.Flags, value map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p MapStrVariant) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]dbus.Variant
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
