package modemmanager1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-dbus-factory/object_manager"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Manager struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	manager                      // interface org.freedesktop.ModemManager1
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.ModemManager1", "/org/freedesktop/ModemManager1")
	return obj
}

func (obj *Manager) ModemManager1() *manager {
	return &obj.manager
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1"
}

// method ScanDevices

func (v *manager) GoScanDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ScanDevices", flags, ch)
}

func (v *manager) ScanDevices(flags dbus.Flags) error {
	return (<-v.GoScanDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetLogging

func (v *manager) GoSetLogging(flags dbus.Flags, ch chan *dbus.Call, level string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogging", flags, ch, level)
}

func (v *manager) SetLogging(flags dbus.Flags, level string) error {
	return (<-v.GoSetLogging(flags, make(chan *dbus.Call, 1), level).Done).Err
}

type Modem struct {
	modem          // interface org.freedesktop.ModemManager1.Modem
	modemFirmware  // interface org.freedesktop.ModemManager1.Modem.Firmware
	modemLocation  // interface org.freedesktop.ModemManager1.Modem.Location
	modemMessaging // interface org.freedesktop.ModemManager1.Modem.Messaging
	modem3gppUssd  // interface org.freedesktop.ModemManager1.Modem.Modem3gpp.Ussd
	modem3gpp      // interface org.freedesktop.ModemManager1.Modem.Modem3gpp
	modemCdma      // interface org.freedesktop.ModemManager1.Modem.ModemCdma
	modemOma       // interface org.freedesktop.ModemManager1.Modem.Oma
	modemSignal    // interface org.freedesktop.ModemManager1.Modem.Signal
	modemSimple    // interface org.freedesktop.ModemManager1.Modem.Simple
	modemTime      // interface org.freedesktop.ModemManager1.Modem.Time
	modemVoice     // interface org.freedesktop.ModemManager1.Modem.Voice
	proxy.Object
}

func NewModem(conn *dbus.Conn, path dbus.ObjectPath) (*Modem, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Modem)
	obj.Object.Init_(conn, "org.freedesktop.ModemManager1", path)
	return obj, nil
}

func (obj *Modem) Modem() *modem {
	return &obj.modem
}

type modem struct{}

func (v *modem) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modem) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem"
}

// method Enable

func (v *modem) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enable)
}

func (v *modem) Enable(flags dbus.Flags, enable bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method ListBearers

func (v *modem) GoListBearers(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListBearers", flags, ch)
}

func (*modem) StoreListBearers(call *dbus.Call) (bearers []dbus.ObjectPath, err error) {
	err = call.Store(&bearers)
	return
}

func (v *modem) ListBearers(flags dbus.Flags) (bearers []dbus.ObjectPath, err error) {
	return v.StoreListBearers(
		<-v.GoListBearers(flags, make(chan *dbus.Call, 1)).Done)
}

// method CreateBearer

func (v *modem) GoCreateBearer(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateBearer", flags, ch, properties)
}

func (*modem) StoreCreateBearer(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *modem) CreateBearer(flags dbus.Flags, properties map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreCreateBearer(
		<-v.GoCreateBearer(flags, make(chan *dbus.Call, 1), properties).Done)
}

// method DeleteBearer

func (v *modem) GoDeleteBearer(flags dbus.Flags, ch chan *dbus.Call, bearer dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteBearer", flags, ch, bearer)
}

func (v *modem) DeleteBearer(flags dbus.Flags, bearer dbus.ObjectPath) error {
	return (<-v.GoDeleteBearer(flags, make(chan *dbus.Call, 1), bearer).Done).Err
}

// method Reset

func (v *modem) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *modem) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method FactoryReset

func (v *modem) GoFactoryReset(flags dbus.Flags, ch chan *dbus.Call, code string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FactoryReset", flags, ch, code)
}

func (v *modem) FactoryReset(flags dbus.Flags, code string) error {
	return (<-v.GoFactoryReset(flags, make(chan *dbus.Call, 1), code).Done).Err
}

// method SetPowerState

func (v *modem) GoSetPowerState(flags dbus.Flags, ch chan *dbus.Call, state uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPowerState", flags, ch, state)
}

func (v *modem) SetPowerState(flags dbus.Flags, state uint32) error {
	return (<-v.GoSetPowerState(flags, make(chan *dbus.Call, 1), state).Done).Err
}

// method SetCurrentCapabilities

func (v *modem) GoSetCurrentCapabilities(flags dbus.Flags, ch chan *dbus.Call, capabilities uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentCapabilities", flags, ch, capabilities)
}

func (v *modem) SetCurrentCapabilities(flags dbus.Flags, capabilities uint32) error {
	return (<-v.GoSetCurrentCapabilities(flags, make(chan *dbus.Call, 1), capabilities).Done).Err
}

// method SetCurrentModes

func (v *modem) GoSetCurrentModes(flags dbus.Flags, ch chan *dbus.Call, modes ModemModes) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentModes", flags, ch, modes)
}

func (v *modem) SetCurrentModes(flags dbus.Flags, modes ModemModes) error {
	return (<-v.GoSetCurrentModes(flags, make(chan *dbus.Call, 1), modes).Done).Err
}

// method SetCurrentBands

func (v *modem) GoSetCurrentBands(flags dbus.Flags, ch chan *dbus.Call, bands []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentBands", flags, ch, bands)
}

func (v *modem) SetCurrentBands(flags dbus.Flags, bands []uint32) error {
	return (<-v.GoSetCurrentBands(flags, make(chan *dbus.Call, 1), bands).Done).Err
}

// method Command

func (v *modem) GoCommand(flags dbus.Flags, ch chan *dbus.Call, cmd string, timeout uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Command", flags, ch, cmd, timeout)
}

func (*modem) StoreCommand(call *dbus.Call) (response string, err error) {
	err = call.Store(&response)
	return
}

func (v *modem) Command(flags dbus.Flags, cmd string, timeout uint32) (response string, err error) {
	return v.StoreCommand(
		<-v.GoCommand(flags, make(chan *dbus.Call, 1), cmd, timeout).Done)
}

// signal StateChanged

func (v *modem) ConnectStateChanged(cb func(old int32, new int32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old int32
		var new int32
		var reason uint32
		err := dbus.Store(sig.Body, &old, &new, &reason)
		if err == nil {
			cb(old, new, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Sim o

func (v *modem) Sim() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Sim",
	}
}

// property Bearers ao

func (v *modem) Bearers() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Bearers",
	}
}

// property SupportedCapabilities au

func (v *modem) SupportedCapabilities() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "SupportedCapabilities",
	}
}

// property CurrentCapabilities u

func (v *modem) CurrentCapabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "CurrentCapabilities",
	}
}

// property MaxBearers u

func (v *modem) MaxBearers() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "MaxBearers",
	}
}

// property MaxActiveBearers u

func (v *modem) MaxActiveBearers() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "MaxActiveBearers",
	}
}

// property Manufacturer s

func (v *modem) Manufacturer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property Model s

func (v *modem) Model() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Model",
	}
}

// property Revision s

func (v *modem) Revision() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Revision",
	}
}

// property DeviceIdentifier s

func (v *modem) DeviceIdentifier() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DeviceIdentifier",
	}
}

// property Device s

func (v *modem) Device() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Device",
	}
}

// property Drivers as

func (v *modem) Drivers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Drivers",
	}
}

// property Plugin s

func (v *modem) Plugin() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Plugin",
	}
}

// property PrimaryPort s

func (v *modem) PrimaryPort() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PrimaryPort",
	}
}

// property Ports a(su)

func (v *modem) Ports() PropModemPorts {
	return PropModemPorts{
		Impl: v,
	}
}

type PropModemPorts struct {
	Impl proxy.Implementer
}

func (p PropModemPorts) Get(flags dbus.Flags) (value []ModemPort, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Ports", &value)
	return
}

func (p PropModemPorts) ConnectChanged(cb func(hasValue bool, value []ModemPort)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModemPort
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
		"Ports", cb0)
}

// property EquipmentIdentifier s

func (v *modem) EquipmentIdentifier() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "EquipmentIdentifier",
	}
}

// property UnlockRequired u

func (v *modem) UnlockRequired() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "UnlockRequired",
	}
}

// property UnlockRetries a{uu}

func (v *modem) UnlockRetries() PropModemUnlockRetries {
	return PropModemUnlockRetries{
		Impl: v,
	}
}

type PropModemUnlockRetries struct {
	Impl proxy.Implementer
}

func (p PropModemUnlockRetries) Get(flags dbus.Flags) (value map[uint32]uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"UnlockRetries", &value)
	return
}

func (p PropModemUnlockRetries) ConnectChanged(cb func(hasValue bool, value map[uint32]uint32)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[uint32]uint32
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
		"UnlockRetries", cb0)
}

// property State i

func (v *modem) State() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "State",
	}
}

// property StateFailedReason u

func (v *modem) StateFailedReason() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "StateFailedReason",
	}
}

// property AccessTechnologies u

func (v *modem) AccessTechnologies() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "AccessTechnologies",
	}
}

// property SignalQuality (ub)

func (v *modem) SignalQuality() PropModemSignalQuality {
	return PropModemSignalQuality{
		Impl: v,
	}
}

type PropModemSignalQuality struct {
	Impl proxy.Implementer
}

func (p PropModemSignalQuality) Get(flags dbus.Flags) (value ModemSignalQuality, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"SignalQuality", &value)
	return
}

func (p PropModemSignalQuality) ConnectChanged(cb func(hasValue bool, value ModemSignalQuality)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModemSignalQuality
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModemSignalQuality{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"SignalQuality", cb0)
}

// property OwnNumbers as

func (v *modem) OwnNumbers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "OwnNumbers",
	}
}

// property PowerState u

func (v *modem) PowerState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "PowerState",
	}
}

// property SupportedModes a(uu)

func (v *modem) SupportedModes() PropModemSupportedModes {
	return PropModemSupportedModes{
		Impl: v,
	}
}

type PropModemSupportedModes struct {
	Impl proxy.Implementer
}

func (p PropModemSupportedModes) Get(flags dbus.Flags) (value []ModemModes, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"SupportedModes", &value)
	return
}

func (p PropModemSupportedModes) ConnectChanged(cb func(hasValue bool, value []ModemModes)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModemModes
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
		"SupportedModes", cb0)
}

// property CurrentModes (uu)

func (v *modem) CurrentModes() PropModemCurrentModes {
	return PropModemCurrentModes{
		Impl: v,
	}
}

type PropModemCurrentModes struct {
	Impl proxy.Implementer
}

func (p PropModemCurrentModes) Get(flags dbus.Flags) (value ModemModes, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"CurrentModes", &value)
	return
}

func (p PropModemCurrentModes) ConnectChanged(cb func(hasValue bool, value ModemModes)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModemModes
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModemModes{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"CurrentModes", cb0)
}

// property SupportedBands au

func (v *modem) SupportedBands() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "SupportedBands",
	}
}

// property CurrentBands au

func (v *modem) CurrentBands() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "CurrentBands",
	}
}

// property SupportedIpFamilies u

func (v *modem) SupportedIpFamilies() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "SupportedIpFamilies",
	}
}

func (obj *Modem) Firmware() *modemFirmware {
	return &obj.modemFirmware
}

type modemFirmware struct{}

func (v *modemFirmware) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemFirmware) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Firmware"
}

// method List

func (v *modemFirmware) GoList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".List", flags, ch)
}

func (*modemFirmware) StoreList(call *dbus.Call) (selected string, installed []map[string]dbus.Variant, err error) {
	err = call.Store(&selected, &installed)
	return
}

func (v *modemFirmware) List(flags dbus.Flags) (selected string, installed []map[string]dbus.Variant, err error) {
	return v.StoreList(
		<-v.GoList(flags, make(chan *dbus.Call, 1)).Done)
}

// method Select

func (v *modemFirmware) GoSelect(flags dbus.Flags, ch chan *dbus.Call, uniqueid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Select", flags, ch, uniqueid)
}

func (v *modemFirmware) Select(flags dbus.Flags, uniqueid string) error {
	return (<-v.GoSelect(flags, make(chan *dbus.Call, 1), uniqueid).Done).Err
}

func (obj *Modem) Location() *modemLocation {
	return &obj.modemLocation
}

type modemLocation struct{}

func (v *modemLocation) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemLocation) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Location"
}

// method Setup

func (v *modemLocation) GoSetup(flags dbus.Flags, ch chan *dbus.Call, sources uint32, signal_location bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Setup", flags, ch, sources, signal_location)
}

func (v *modemLocation) Setup(flags dbus.Flags, sources uint32, signal_location bool) error {
	return (<-v.GoSetup(flags, make(chan *dbus.Call, 1), sources, signal_location).Done).Err
}

// method GetLocation

func (v *modemLocation) GoGetLocation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLocation", flags, ch)
}

func (*modemLocation) StoreGetLocation(call *dbus.Call) (Location map[uint32]dbus.Variant, err error) {
	err = call.Store(&Location)
	return
}

func (v *modemLocation) GetLocation(flags dbus.Flags) (Location map[uint32]dbus.Variant, err error) {
	return v.StoreGetLocation(
		<-v.GoGetLocation(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetSuplServer

func (v *modemLocation) GoSetSuplServer(flags dbus.Flags, ch chan *dbus.Call, supl string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSuplServer", flags, ch, supl)
}

func (v *modemLocation) SetSuplServer(flags dbus.Flags, supl string) error {
	return (<-v.GoSetSuplServer(flags, make(chan *dbus.Call, 1), supl).Done).Err
}

// method SetGpsRefreshRate

func (v *modemLocation) GoSetGpsRefreshRate(flags dbus.Flags, ch chan *dbus.Call, rate uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetGpsRefreshRate", flags, ch, rate)
}

func (v *modemLocation) SetGpsRefreshRate(flags dbus.Flags, rate uint32) error {
	return (<-v.GoSetGpsRefreshRate(flags, make(chan *dbus.Call, 1), rate).Done).Err
}

// property Capabilities u

func (v *modemLocation) Capabilities() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Capabilities",
	}
}

// property Enabled u

func (v *modemLocation) Enabled() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Enabled",
	}
}

// property SignalsLocation b

func (v *modemLocation) SignalsLocation() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SignalsLocation",
	}
}

// property Location a{uv}

func (v *modemLocation) Location() PropModemLocation {
	return PropModemLocation{
		Impl: v,
	}
}

type PropModemLocation struct {
	Impl proxy.Implementer
}

func (p PropModemLocation) Get(flags dbus.Flags) (value map[uint32]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Location", &value)
	return
}

func (p PropModemLocation) ConnectChanged(cb func(hasValue bool, value map[uint32]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[uint32]dbus.Variant
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
		"Location", cb0)
}

// property SuplServer s

func (v *modemLocation) SuplServer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SuplServer",
	}
}

// property GpsRefreshRate u

func (v *modemLocation) GpsRefreshRate() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "GpsRefreshRate",
	}
}

func (obj *Modem) Messaging() *modemMessaging {
	return &obj.modemMessaging
}

type modemMessaging struct{}

func (v *modemMessaging) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemMessaging) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Messaging"
}

// method List

func (v *modemMessaging) GoList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".List", flags, ch)
}

func (*modemMessaging) StoreList(call *dbus.Call) (result []dbus.ObjectPath, err error) {
	err = call.Store(&result)
	return
}

func (v *modemMessaging) List(flags dbus.Flags) (result []dbus.ObjectPath, err error) {
	return v.StoreList(
		<-v.GoList(flags, make(chan *dbus.Call, 1)).Done)
}

// method Delete

func (v *modemMessaging) GoDelete(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, path)
}

func (v *modemMessaging) Delete(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method Create

func (v *modemMessaging) GoCreate(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Create", flags, ch, properties)
}

func (*modemMessaging) StoreCreate(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *modemMessaging) Create(flags dbus.Flags, properties map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreCreate(
		<-v.GoCreate(flags, make(chan *dbus.Call, 1), properties).Done)
}

// signal Added

func (v *modemMessaging) ConnectAdded(cb func(path dbus.ObjectPath, received bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Added", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Added",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var received bool
		err := dbus.Store(sig.Body, &path, &received)
		if err == nil {
			cb(path, received)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Deleted

func (v *modemMessaging) ConnectDeleted(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Deleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Deleted",
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

// property Messages ao

func (v *modemMessaging) Messages() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Messages",
	}
}

// property SupportedStorages au

func (v *modemMessaging) SupportedStorages() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "SupportedStorages",
	}
}

// property DefaultStorage u

func (v *modemMessaging) DefaultStorage() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "DefaultStorage",
	}
}

func (obj *Modem) Modem3gppUssd() *modem3gppUssd {
	return &obj.modem3gppUssd
}

type modem3gppUssd struct{}

func (v *modem3gppUssd) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modem3gppUssd) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Modem3gpp.Ussd"
}

// method Initiate

func (v *modem3gppUssd) GoInitiate(flags dbus.Flags, ch chan *dbus.Call, command string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Initiate", flags, ch, command)
}

func (*modem3gppUssd) StoreInitiate(call *dbus.Call) (reply string, err error) {
	err = call.Store(&reply)
	return
}

func (v *modem3gppUssd) Initiate(flags dbus.Flags, command string) (reply string, err error) {
	return v.StoreInitiate(
		<-v.GoInitiate(flags, make(chan *dbus.Call, 1), command).Done)
}

// method Respond

func (v *modem3gppUssd) GoRespond(flags dbus.Flags, ch chan *dbus.Call, response string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Respond", flags, ch, response)
}

func (*modem3gppUssd) StoreRespond(call *dbus.Call) (reply string, err error) {
	err = call.Store(&reply)
	return
}

func (v *modem3gppUssd) Respond(flags dbus.Flags, response string) (reply string, err error) {
	return v.StoreRespond(
		<-v.GoRespond(flags, make(chan *dbus.Call, 1), response).Done)
}

// method Cancel

func (v *modem3gppUssd) GoCancel(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Cancel", flags, ch)
}

func (v *modem3gppUssd) Cancel(flags dbus.Flags) error {
	return (<-v.GoCancel(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property State u

func (v *modem3gppUssd) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property NetworkNotification s

func (v *modem3gppUssd) NetworkNotification() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "NetworkNotification",
	}
}

// property NetworkRequest s

func (v *modem3gppUssd) NetworkRequest() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "NetworkRequest",
	}
}

func (obj *Modem) Modem3gpp() *modem3gpp {
	return &obj.modem3gpp
}

type modem3gpp struct{}

func (v *modem3gpp) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modem3gpp) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Modem3gpp"
}

// method Register

func (v *modem3gpp) GoRegister(flags dbus.Flags, ch chan *dbus.Call, operator_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, operator_id)
}

func (v *modem3gpp) Register(flags dbus.Flags, operator_id string) error {
	return (<-v.GoRegister(flags, make(chan *dbus.Call, 1), operator_id).Done).Err
}

// method Scan

func (v *modem3gpp) GoScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Scan", flags, ch)
}

func (*modem3gpp) StoreScan(call *dbus.Call) (results []map[string]dbus.Variant, err error) {
	err = call.Store(&results)
	return
}

func (v *modem3gpp) Scan(flags dbus.Flags) (results []map[string]dbus.Variant, err error) {
	return v.StoreScan(
		<-v.GoScan(flags, make(chan *dbus.Call, 1)).Done)
}

// property Imei s

func (v *modem3gpp) Imei() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Imei",
	}
}

// property RegistrationState u

func (v *modem3gpp) RegistrationState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RegistrationState",
	}
}

// property OperatorCode s

func (v *modem3gpp) OperatorCode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OperatorCode",
	}
}

// property OperatorName s

func (v *modem3gpp) OperatorName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OperatorName",
	}
}

// property EnabledFacilityLocks u

func (v *modem3gpp) EnabledFacilityLocks() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "EnabledFacilityLocks",
	}
}

// property SubscriptionState u

func (v *modem3gpp) SubscriptionState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "SubscriptionState",
	}
}

func (obj *Modem) Cdma() *modemCdma {
	return &obj.modemCdma
}

type modemCdma struct{}

func (v *modemCdma) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemCdma) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.ModemCdma"
}

// method Activate

func (v *modemCdma) GoActivate(flags dbus.Flags, ch chan *dbus.Call, carrier_code string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Activate", flags, ch, carrier_code)
}

func (v *modemCdma) Activate(flags dbus.Flags, carrier_code string) error {
	return (<-v.GoActivate(flags, make(chan *dbus.Call, 1), carrier_code).Done).Err
}

// method ActivateManual

func (v *modemCdma) GoActivateManual(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateManual", flags, ch, properties)
}

func (v *modemCdma) ActivateManual(flags dbus.Flags, properties map[string]dbus.Variant) error {
	return (<-v.GoActivateManual(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

// signal ActivationStateChanged

func (v *modemCdma) ConnectActivationStateChanged(cb func(activation_state uint32, activation_error uint32, status_changes map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActivationStateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActivationStateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var activation_state uint32
		var activation_error uint32
		var status_changes map[string]dbus.Variant
		err := dbus.Store(sig.Body, &activation_state, &activation_error, &status_changes)
		if err == nil {
			cb(activation_state, activation_error, status_changes)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ActivationState u

func (v *modemCdma) ActivationState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ActivationState",
	}
}

// property Meid s

func (v *modemCdma) Meid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Meid",
	}
}

// property Esn s

func (v *modemCdma) Esn() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Esn",
	}
}

// property Sid u

func (v *modemCdma) Sid() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Sid",
	}
}

// property Nid u

func (v *modemCdma) Nid() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Nid",
	}
}

// property Cdma1xRegistrationState u

func (v *modemCdma) Cdma1xRegistrationState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Cdma1xRegistrationState",
	}
}

// property EvdoRegistrationState u

func (v *modemCdma) EvdoRegistrationState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "EvdoRegistrationState",
	}
}

func (obj *Modem) Oma() *modemOma {
	return &obj.modemOma
}

type modemOma struct{}

func (v *modemOma) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemOma) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Oma"
}

// method Setup

func (v *modemOma) GoSetup(flags dbus.Flags, ch chan *dbus.Call, features uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Setup", flags, ch, features)
}

func (v *modemOma) Setup(flags dbus.Flags, features uint32) error {
	return (<-v.GoSetup(flags, make(chan *dbus.Call, 1), features).Done).Err
}

// method StartClientInitiatedSession

func (v *modemOma) GoStartClientInitiatedSession(flags dbus.Flags, ch chan *dbus.Call, session_type uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartClientInitiatedSession", flags, ch, session_type)
}

func (v *modemOma) StartClientInitiatedSession(flags dbus.Flags, session_type uint32) error {
	return (<-v.GoStartClientInitiatedSession(flags, make(chan *dbus.Call, 1), session_type).Done).Err
}

// method AcceptNetworkInitiatedSession

func (v *modemOma) GoAcceptNetworkInitiatedSession(flags dbus.Flags, ch chan *dbus.Call, session_id uint32, accept bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AcceptNetworkInitiatedSession", flags, ch, session_id, accept)
}

func (v *modemOma) AcceptNetworkInitiatedSession(flags dbus.Flags, session_id uint32, accept bool) error {
	return (<-v.GoAcceptNetworkInitiatedSession(flags, make(chan *dbus.Call, 1), session_id, accept).Done).Err
}

// method CancelSession

func (v *modemOma) GoCancelSession(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelSession", flags, ch)
}

func (v *modemOma) CancelSession(flags dbus.Flags) error {
	return (<-v.GoCancelSession(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal SessionStateChanged

func (v *modemOma) ConnectSessionStateChanged(cb func(old_session_state int32, new_session_state int32, session_state_failed_reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SessionStateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SessionStateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old_session_state int32
		var new_session_state int32
		var session_state_failed_reason uint32
		err := dbus.Store(sig.Body, &old_session_state, &new_session_state, &session_state_failed_reason)
		if err == nil {
			cb(old_session_state, new_session_state, session_state_failed_reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Features u

func (v *modemOma) Features() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Features",
	}
}

// property PendingNetworkInitiatedSessions a(uu)

func (v *modemOma) PendingNetworkInitiatedSessions() PropModemOmaPendingNetworkInitiatedSessions {
	return PropModemOmaPendingNetworkInitiatedSessions{
		Impl: v,
	}
}

type PropModemOmaPendingNetworkInitiatedSessions struct {
	Impl proxy.Implementer
}

func (p PropModemOmaPendingNetworkInitiatedSessions) Get(flags dbus.Flags) (value []OmaSession, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"PendingNetworkInitiatedSessions", &value)
	return
}

func (p PropModemOmaPendingNetworkInitiatedSessions) ConnectChanged(cb func(hasValue bool, value []OmaSession)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []OmaSession
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
		"PendingNetworkInitiatedSessions", cb0)
}

// property SessionType u

func (v *modemOma) SessionType() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "SessionType",
	}
}

// property SessionState i

func (v *modemOma) SessionState() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SessionState",
	}
}

func (obj *Modem) Signal() *modemSignal {
	return &obj.modemSignal
}

type modemSignal struct{}

func (v *modemSignal) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemSignal) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Signal"
}

// method Setup

func (v *modemSignal) GoSetup(flags dbus.Flags, ch chan *dbus.Call, rate uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Setup", flags, ch, rate)
}

func (v *modemSignal) Setup(flags dbus.Flags, rate uint32) error {
	return (<-v.GoSetup(flags, make(chan *dbus.Call, 1), rate).Done).Err
}

// property Rate u

func (v *modemSignal) Rate() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Rate",
	}
}

// property Cdma a{sv}

func (v *modemSignal) Cdma() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Cdma",
	}
}

// property Evdo a{sv}

func (v *modemSignal) Evdo() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Evdo",
	}
}

// property Gsm a{sv}

func (v *modemSignal) Gsm() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Gsm",
	}
}

// property Umts a{sv}

func (v *modemSignal) Umts() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Umts",
	}
}

// property Lte a{sv}

func (v *modemSignal) Lte() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "Lte",
	}
}

func (obj *Modem) Simple() *modemSimple {
	return &obj.modemSimple
}

type modemSimple struct{}

func (v *modemSimple) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemSimple) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Simple"
}

// method Connect

func (v *modemSimple) GoConnect(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch, properties)
}

func (*modemSimple) StoreConnect(call *dbus.Call) (bearer dbus.ObjectPath, err error) {
	err = call.Store(&bearer)
	return
}

func (v *modemSimple) Connect(flags dbus.Flags, properties map[string]dbus.Variant) (bearer dbus.ObjectPath, err error) {
	return v.StoreConnect(
		<-v.GoConnect(flags, make(chan *dbus.Call, 1), properties).Done)
}

// method Disconnect

func (v *modemSimple) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call, bearer dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch, bearer)
}

func (v *modemSimple) Disconnect(flags dbus.Flags, bearer dbus.ObjectPath) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1), bearer).Done).Err
}

// method GetStatus

func (v *modemSimple) GoGetStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStatus", flags, ch)
}

func (*modemSimple) StoreGetStatus(call *dbus.Call) (properties map[string]dbus.Variant, err error) {
	err = call.Store(&properties)
	return
}

func (v *modemSimple) GetStatus(flags dbus.Flags) (properties map[string]dbus.Variant, err error) {
	return v.StoreGetStatus(
		<-v.GoGetStatus(flags, make(chan *dbus.Call, 1)).Done)
}

func (obj *Modem) Time() *modemTime {
	return &obj.modemTime
}

type modemTime struct{}

func (v *modemTime) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemTime) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Time"
}

// method GetNetworkTime

func (v *modemTime) GoGetNetworkTime(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNetworkTime", flags, ch)
}

func (*modemTime) StoreGetNetworkTime(call *dbus.Call) (time string, err error) {
	err = call.Store(&time)
	return
}

func (v *modemTime) GetNetworkTime(flags dbus.Flags) (time string, err error) {
	return v.StoreGetNetworkTime(
		<-v.GoGetNetworkTime(flags, make(chan *dbus.Call, 1)).Done)
}

// signal NetworkTimeChanged

func (v *modemTime) ConnectNetworkTimeChanged(cb func(time string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NetworkTimeChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NetworkTimeChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var time string
		err := dbus.Store(sig.Body, &time)
		if err == nil {
			cb(time)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property NetworkTimezone a{sv}

func (v *modemTime) NetworkTimezone() PropMapStringVariant {
	return PropMapStringVariant{
		Impl: v,
		Name: "NetworkTimezone",
	}
}

func (obj *Modem) Voice() *modemVoice {
	return &obj.modemVoice
}

type modemVoice struct{}

func (v *modemVoice) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*modemVoice) GetInterfaceName_() string {
	return "org.freedesktop.ModemManager1.Modem.Voice"
}

// method ListCalls

func (v *modemVoice) GoListCalls(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListCalls", flags, ch)
}

func (*modemVoice) StoreListCalls(call *dbus.Call) (result []dbus.ObjectPath, err error) {
	err = call.Store(&result)
	return
}

func (v *modemVoice) ListCalls(flags dbus.Flags) (result []dbus.ObjectPath, err error) {
	return v.StoreListCalls(
		<-v.GoListCalls(flags, make(chan *dbus.Call, 1)).Done)
}

// method DeleteCall

func (v *modemVoice) GoDeleteCall(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCall", flags, ch, path)
}

func (v *modemVoice) DeleteCall(flags dbus.Flags, path dbus.ObjectPath) error {
	return (<-v.GoDeleteCall(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method CreateCall

func (v *modemVoice) GoCreateCall(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateCall", flags, ch, properties)
}

func (*modemVoice) StoreCreateCall(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *modemVoice) CreateCall(flags dbus.Flags, properties map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreCreateCall(
		<-v.GoCreateCall(flags, make(chan *dbus.Call, 1), properties).Done)
}

// signal CallAdded

func (v *modemVoice) ConnectCallAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CallAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CallAdded",
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

// signal CallDeleted

func (v *modemVoice) ConnectCallDeleted(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CallDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CallDeleted",
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

// property Calls ao

func (v *modemVoice) Calls() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Calls",
	}
}

type PropMapStringVariant struct {
	Impl proxy.Implementer
	Name string
}

func (p PropMapStringVariant) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropMapStringVariant) Set(flags dbus.Flags, value map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropMapStringVariant) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
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
