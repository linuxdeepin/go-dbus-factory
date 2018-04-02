package networkmanager

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

type AccessPoint struct {
	accessPoint // interface org.freedesktop.NetworkManager.AccessPoint
	proxy.Object
}

func NewAccessPoint(conn *dbus.Conn, path dbus.ObjectPath) (*AccessPoint, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(AccessPoint)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type accessPoint struct{}

func (v *accessPoint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*accessPoint) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.AccessPoint"
}

// signal PropertiesChanged

func (v *accessPoint) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Flags u

func (v *accessPoint) Flags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Flags",
	}
}

// property WpaFlags u

func (v *accessPoint) WpaFlags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "WpaFlags",
	}
}

// property RsnFlags u

func (v *accessPoint) RsnFlags() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RsnFlags",
	}
}

// property Ssid ay

func (v *accessPoint) Ssid() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "Ssid",
	}
}

// property Frequency u

func (v *accessPoint) Frequency() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Frequency",
	}
}

// property HwAddress s

func (v *accessPoint) HwAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HwAddress",
	}
}

// property Mode u

func (v *accessPoint) Mode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Mode",
	}
}

// property MaxBitrate u

func (v *accessPoint) MaxBitrate() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "MaxBitrate",
	}
}

// property Strength y

func (v *accessPoint) Strength() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Strength",
	}
}

// property LastSeen i

func (v *accessPoint) LastSeen() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "LastSeen",
	}
}

type AgentManager struct {
	agentManager // interface org.freedesktop.NetworkManager.AgentManager
	proxy.Object
}

func NewAgentManager(conn *dbus.Conn) *AgentManager {
	obj := new(AgentManager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/AgentManager")
	return obj
}

type agentManager struct{}

func (v *agentManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*agentManager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.AgentManager"
}

// method Register

func (v *agentManager) GoRegister(flags dbus.Flags, ch chan *dbus.Call, identifier string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, identifier)
}

func (v *agentManager) Register(flags dbus.Flags, identifier string) error {
	return (<-v.GoRegister(flags, make(chan *dbus.Call, 1), identifier).Done).Err
}

// method RegisterWithCapabilities

func (v *agentManager) GoRegisterWithCapabilities(flags dbus.Flags, ch chan *dbus.Call, identifier string, capabilities uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterWithCapabilities", flags, ch, identifier, capabilities)
}

func (v *agentManager) RegisterWithCapabilities(flags dbus.Flags, identifier string, capabilities uint32) error {
	return (<-v.GoRegisterWithCapabilities(flags, make(chan *dbus.Call, 1), identifier, capabilities).Done).Err
}

// method Unregister

func (v *agentManager) GoUnregister(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unregister", flags, ch)
}

func (v *agentManager) Unregister(flags dbus.Flags) error {
	return (<-v.GoUnregister(flags, make(chan *dbus.Call, 1)).Done).Err
}

type Checkpoint struct {
	checkpoint // interface org.freedesktop.NetworkManager.Checkpoint
	proxy.Object
}

func NewCheckpoint(conn *dbus.Conn, path dbus.ObjectPath) (*Checkpoint, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Checkpoint)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type checkpoint struct{}

func (v *checkpoint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*checkpoint) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Checkpoint"
}

// signal PropertiesChanged

func (v *checkpoint) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Devices ao

func (v *checkpoint) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property Created x

func (v *checkpoint) Created() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Created",
	}
}

// property RollbackTimeout u

func (v *checkpoint) RollbackTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RollbackTimeout",
	}
}

type ActiveConnection struct {
	activeConnection // interface org.freedesktop.NetworkManager.Connection.Active
	proxy.Object
}

func NewActiveConnection(conn *dbus.Conn, path dbus.ObjectPath) (*ActiveConnection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(ActiveConnection)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type activeConnection struct{}

func (v *activeConnection) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*activeConnection) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Connection.Active"
}

// signal StateChanged

func (v *activeConnection) ConnectStateChanged(cb func(state uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
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
		var state uint32
		var reason uint32
		err := dbus.Store(sig.Body, &state, &reason)
		if err == nil {
			cb(state, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *activeConnection) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Connection o

func (v *activeConnection) Connection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Connection",
	}
}

// property SpecificObject o

func (v *activeConnection) SpecificObject() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "SpecificObject",
	}
}

// property Id s

func (v *activeConnection) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Uuid s

func (v *activeConnection) Uuid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Uuid",
	}
}

// property Type s

func (v *activeConnection) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Devices ao

func (v *activeConnection) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property State u

func (v *activeConnection) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property Default b

func (v *activeConnection) Default() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Default",
	}
}

// property Ip4Config o

func (v *activeConnection) Ip4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip4Config",
	}
}

// property Dhcp4Config o

func (v *activeConnection) Dhcp4Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp4Config",
	}
}

// property Default6 b

func (v *activeConnection) Default6() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Default6",
	}
}

// property Ip6Config o

func (v *activeConnection) Ip6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Ip6Config",
	}
}

// property Dhcp6Config o

func (v *activeConnection) Dhcp6Config() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Dhcp6Config",
	}
}

// property Vpn b

func (v *activeConnection) Vpn() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Vpn",
	}
}

// property Master o

func (v *activeConnection) Master() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Master",
	}
}

type Dhcp4Config struct {
	proxy.Object
}

func NewDhcp4Config(conn *dbus.Conn, path dbus.ObjectPath) (*Dhcp4Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Dhcp4Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type Dhcp6Config struct {
	proxy.Object
}

func NewDhcp6Config(conn *dbus.Conn, path dbus.ObjectPath) (*Dhcp6Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Dhcp6Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type DnsManager struct {
	dnsManager // interface org.freedesktop.NetworkManager.DnsManager
	proxy.Object
}

func NewDnsManager(conn *dbus.Conn) *DnsManager {
	obj := new(DnsManager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/DnsManager")
	return obj
}

type dnsManager struct{}

func (v *dnsManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dnsManager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.DnsManager"
}

// property Mode s

func (v *dnsManager) Mode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Mode",
	}
}

// property RcManager s

func (v *dnsManager) RcManager() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RcManager",
	}
}

// property Configuration aa{sv}

func (v *dnsManager) Configuration() PropDnsManagerConfiguration {
	return PropDnsManagerConfiguration{
		Impl: v,
	}
}

type PropDnsManagerConfiguration struct {
	Impl proxy.Implementer
}

func (p PropDnsManagerConfiguration) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Configuration", &value)
	return
}

func (p PropDnsManagerConfiguration) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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
		"Configuration", cb0)
}

type IP4Config struct {
	ip4Config // interface org.freedesktop.NetworkManager.IP4Config
	proxy.Object
}

func NewIP4Config(conn *dbus.Conn, path dbus.ObjectPath) (*IP4Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(IP4Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ip4Config struct{}

func (v *ip4Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ip4Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.IP4Config"
}

// signal PropertiesChanged

func (v *ip4Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Addresses aau

func (v *ip4Config) Addresses() PropIP4ConfigAddresses {
	return PropIP4ConfigAddresses{
		Impl: v,
	}
}

type PropIP4ConfigAddresses struct {
	Impl proxy.Implementer
}

func (p PropIP4ConfigAddresses) Get(flags dbus.Flags) (value [][]uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Addresses", &value)
	return
}

func (p PropIP4ConfigAddresses) ConnectChanged(cb func(hasValue bool, value [][]uint32)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]uint32
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
		"Addresses", cb0)
}

// property AddressData aa{sv}

func (v *ip4Config) AddressData() PropIP4ConfigAddressData {
	return PropIP4ConfigAddressData{
		Impl: v,
	}
}

type PropIP4ConfigAddressData struct {
	Impl proxy.Implementer
}

func (p PropIP4ConfigAddressData) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"AddressData", &value)
	return
}

func (p PropIP4ConfigAddressData) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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
		"AddressData", cb0)
}

// property Gateway s

func (v *ip4Config) Gateway() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Gateway",
	}
}

// property Routes aau

func (v *ip4Config) Routes() PropIP4ConfigRoutes {
	return PropIP4ConfigRoutes{
		Impl: v,
	}
}

type PropIP4ConfigRoutes struct {
	Impl proxy.Implementer
}

func (p PropIP4ConfigRoutes) Get(flags dbus.Flags) (value [][]uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Routes", &value)
	return
}

func (p PropIP4ConfigRoutes) ConnectChanged(cb func(hasValue bool, value [][]uint32)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]uint32
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
		"Routes", cb0)
}

// property RouteData aa{sv}

func (v *ip4Config) RouteData() PropIP4ConfigRouteData {
	return PropIP4ConfigRouteData{
		Impl: v,
	}
}

type PropIP4ConfigRouteData struct {
	Impl proxy.Implementer
}

func (p PropIP4ConfigRouteData) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"RouteData", &value)
	return
}

func (p PropIP4ConfigRouteData) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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
		"RouteData", cb0)
}

// property Nameservers au

func (v *ip4Config) Nameservers() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "Nameservers",
	}
}

// property Domains as

func (v *ip4Config) Domains() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Domains",
	}
}

// property Searches as

func (v *ip4Config) Searches() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Searches",
	}
}

// property DnsOptions as

func (v *ip4Config) DnsOptions() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DnsOptions",
	}
}

// property DnsPriority i

func (v *ip4Config) DnsPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DnsPriority",
	}
}

// property WinsServers au

func (v *ip4Config) WinsServers() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "WinsServers",
	}
}

type IP6Config struct {
	ip6Config // interface org.freedesktop.NetworkManager.IP6Config
	proxy.Object
}

func NewIP6Config(conn *dbus.Conn, path dbus.ObjectPath) (*IP6Config, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(IP6Config)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ip6Config struct{}

func (v *ip6Config) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ip6Config) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.IP6Config"
}

// signal PropertiesChanged

func (v *ip6Config) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Addresses a(ayuay)

func (v *ip6Config) Addresses() PropIP6ConfigAddresses {
	return PropIP6ConfigAddresses{
		Impl: v,
	}
}

type PropIP6ConfigAddresses struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigAddresses) Get(flags dbus.Flags) (value []IP6Address, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Addresses", &value)
	return
}

func (p PropIP6ConfigAddresses) ConnectChanged(cb func(hasValue bool, value []IP6Address)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IP6Address
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
		"Addresses", cb0)
}

// property AddressData aa{sv}

func (v *ip6Config) AddressData() PropIP6ConfigAddressData {
	return PropIP6ConfigAddressData{
		Impl: v,
	}
}

type PropIP6ConfigAddressData struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigAddressData) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"AddressData", &value)
	return
}

func (p PropIP6ConfigAddressData) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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
		"AddressData", cb0)
}

// property Gateway s

func (v *ip6Config) Gateway() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Gateway",
	}
}

// property Routes a(ayuayu)

func (v *ip6Config) Routes() PropIP6ConfigRoutes {
	return PropIP6ConfigRoutes{
		Impl: v,
	}
}

type PropIP6ConfigRoutes struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigRoutes) Get(flags dbus.Flags) (value []IP6Route, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Routes", &value)
	return
}

func (p PropIP6ConfigRoutes) ConnectChanged(cb func(hasValue bool, value []IP6Route)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IP6Route
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
		"Routes", cb0)
}

// property RouteData aa{sv}

func (v *ip6Config) RouteData() PropIP6ConfigRouteData {
	return PropIP6ConfigRouteData{
		Impl: v,
	}
}

type PropIP6ConfigRouteData struct {
	Impl proxy.Implementer
}

func (p PropIP6ConfigRouteData) Get(flags dbus.Flags) (value []map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"RouteData", &value)
	return
}

func (p PropIP6ConfigRouteData) ConnectChanged(cb func(hasValue bool, value []map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []map[string]dbus.Variant
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
		"RouteData", cb0)
}

// property Nameservers aay

func (v *ip6Config) Nameservers() PropIP6NameServers {
	return PropIP6NameServers{
		Impl: v,
	}
}

type PropIP6NameServers struct {
	Impl proxy.Implementer
}

func (p PropIP6NameServers) Get(flags dbus.Flags) (value [][]byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Nameservers", &value)
	return
}

func (p PropIP6NameServers) ConnectChanged(cb func(hasValue bool, value [][]byte)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]byte
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
		"Nameservers", cb0)
}

// property Domains as

func (v *ip6Config) Domains() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Domains",
	}
}

// property Searches as

func (v *ip6Config) Searches() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Searches",
	}
}

// property DnsOptions as

func (v *ip6Config) DnsOptions() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DnsOptions",
	}
}

// property DnsPriority i

func (v *ip6Config) DnsPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DnsPriority",
	}
}

type Manager struct {
	manager // interface org.freedesktop.NetworkManager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager"
}

// method Reload

func (v *manager) GoReload(flags dbus.Flags, ch chan *dbus.Call, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, flags0)
}

func (v *manager) Reload(flags dbus.Flags, flags0 uint32) error {
	return (<-v.GoReload(flags, make(chan *dbus.Call, 1), flags0).Done).Err
}

// method GetDevices

func (v *manager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*manager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAllDevices

func (v *manager) GoGetAllDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllDevices", flags, ch)
}

func (*manager) StoreGetAllDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetAllDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetAllDevices(
		<-v.GoGetAllDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDeviceByIpIface

func (v *manager) GoGetDeviceByIpIface(flags dbus.Flags, ch chan *dbus.Call, iface string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDeviceByIpIface", flags, ch, iface)
}

func (*manager) StoreGetDeviceByIpIface(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *manager) GetDeviceByIpIface(flags dbus.Flags, iface string) (device dbus.ObjectPath, err error) {
	return v.StoreGetDeviceByIpIface(
		<-v.GoGetDeviceByIpIface(flags, make(chan *dbus.Call, 1), iface).Done)
}

// method ActivateConnection

func (v *manager) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, connection dbus.ObjectPath, device dbus.ObjectPath, specific_object dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateConnection", flags, ch, connection, device, specific_object)
}

func (*manager) StoreActivateConnection(call *dbus.Call) (active_connection dbus.ObjectPath, err error) {
	err = call.Store(&active_connection)
	return
}

func (v *manager) ActivateConnection(flags dbus.Flags, connection dbus.ObjectPath, device dbus.ObjectPath, specific_object dbus.ObjectPath) (active_connection dbus.ObjectPath, err error) {
	return v.StoreActivateConnection(
		<-v.GoActivateConnection(flags, make(chan *dbus.Call, 1), connection, device, specific_object).Done)
}

// method AddAndActivateConnection

func (v *manager) GoAddAndActivateConnection(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, device dbus.ObjectPath, specific_object dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAndActivateConnection", flags, ch, connection, device, specific_object)
}

func (*manager) StoreAddAndActivateConnection(call *dbus.Call) (path dbus.ObjectPath, active_connection dbus.ObjectPath, err error) {
	err = call.Store(&path, &active_connection)
	return
}

func (v *manager) AddAndActivateConnection(flags dbus.Flags, connection map[string]map[string]dbus.Variant, device dbus.ObjectPath, specific_object dbus.ObjectPath) (path dbus.ObjectPath, active_connection dbus.ObjectPath, err error) {
	return v.StoreAddAndActivateConnection(
		<-v.GoAddAndActivateConnection(flags, make(chan *dbus.Call, 1), connection, device, specific_object).Done)
}

// method DeactivateConnection

func (v *manager) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, active_connection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeactivateConnection", flags, ch, active_connection)
}

func (v *manager) DeactivateConnection(flags dbus.Flags, active_connection dbus.ObjectPath) error {
	return (<-v.GoDeactivateConnection(flags, make(chan *dbus.Call, 1), active_connection).Done).Err
}

// method Sleep

func (v *manager) GoSleep(flags dbus.Flags, ch chan *dbus.Call, sleep bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Sleep", flags, ch, sleep)
}

func (v *manager) Sleep(flags dbus.Flags, sleep bool) error {
	return (<-v.GoSleep(flags, make(chan *dbus.Call, 1), sleep).Done).Err
}

// method Enable

func (v *manager) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enable)
}

func (v *manager) Enable(flags dbus.Flags, enable bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method GetPermissions

func (v *manager) GoGetPermissions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetPermissions", flags, ch)
}

func (*manager) StoreGetPermissions(call *dbus.Call) (permissions map[string]string, err error) {
	err = call.Store(&permissions)
	return
}

func (v *manager) GetPermissions(flags dbus.Flags) (permissions map[string]string, err error) {
	return v.StoreGetPermissions(
		<-v.GoGetPermissions(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetLogging

func (v *manager) GoSetLogging(flags dbus.Flags, ch chan *dbus.Call, level string, domains string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogging", flags, ch, level, domains)
}

func (v *manager) SetLogging(flags dbus.Flags, level string, domains string) error {
	return (<-v.GoSetLogging(flags, make(chan *dbus.Call, 1), level, domains).Done).Err
}

// method GetLogging

func (v *manager) GoGetLogging(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLogging", flags, ch)
}

func (*manager) StoreGetLogging(call *dbus.Call) (level string, domains string, err error) {
	err = call.Store(&level, &domains)
	return
}

func (v *manager) GetLogging(flags dbus.Flags) (level string, domains string, err error) {
	return v.StoreGetLogging(
		<-v.GoGetLogging(flags, make(chan *dbus.Call, 1)).Done)
}

// method CheckConnectivity

func (v *manager) GoCheckConnectivity(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckConnectivity", flags, ch)
}

func (*manager) StoreCheckConnectivity(call *dbus.Call) (connectivity uint32, err error) {
	err = call.Store(&connectivity)
	return
}

func (v *manager) CheckConnectivity(flags dbus.Flags) (connectivity uint32, err error) {
	return v.StoreCheckConnectivity(
		<-v.GoCheckConnectivity(flags, make(chan *dbus.Call, 1)).Done)
}

// method state

func (v *manager) Gostate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".state", flags, ch)
}

func (*manager) Storestate(call *dbus.Call) (state uint32, err error) {
	err = call.Store(&state)
	return
}

func (v *manager) state(flags dbus.Flags) (state uint32, err error) {
	return v.Storestate(
		<-v.Gostate(flags, make(chan *dbus.Call, 1)).Done)
}

// method CheckpointCreate

func (v *manager) GoCheckpointCreate(flags dbus.Flags, ch chan *dbus.Call, devices []dbus.ObjectPath, rollback_timeout uint32, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointCreate", flags, ch, devices, rollback_timeout, flags0)
}

func (*manager) StoreCheckpointCreate(call *dbus.Call) (checkpoint dbus.ObjectPath, err error) {
	err = call.Store(&checkpoint)
	return
}

func (v *manager) CheckpointCreate(flags dbus.Flags, devices []dbus.ObjectPath, rollback_timeout uint32, flags0 uint32) (checkpoint dbus.ObjectPath, err error) {
	return v.StoreCheckpointCreate(
		<-v.GoCheckpointCreate(flags, make(chan *dbus.Call, 1), devices, rollback_timeout, flags0).Done)
}

// method CheckpointDestroy

func (v *manager) GoCheckpointDestroy(flags dbus.Flags, ch chan *dbus.Call, checkpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointDestroy", flags, ch, checkpoint)
}

func (v *manager) CheckpointDestroy(flags dbus.Flags, checkpoint dbus.ObjectPath) error {
	return (<-v.GoCheckpointDestroy(flags, make(chan *dbus.Call, 1), checkpoint).Done).Err
}

// method CheckpointRollback

func (v *manager) GoCheckpointRollback(flags dbus.Flags, ch chan *dbus.Call, checkpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckpointRollback", flags, ch, checkpoint)
}

func (*manager) StoreCheckpointRollback(call *dbus.Call) (result map[string]uint32, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CheckpointRollback(flags dbus.Flags, checkpoint dbus.ObjectPath) (result map[string]uint32, err error) {
	return v.StoreCheckpointRollback(
		<-v.GoCheckpointRollback(flags, make(chan *dbus.Call, 1), checkpoint).Done)
}

// signal CheckPermissions

func (v *manager) ConnectCheckPermissions(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CheckPermissions", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CheckPermissions",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StateChanged

func (v *manager) ConnectStateChanged(cb func(state uint32)) (dbusutil.SignalHandlerId, error) {
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
		var state uint32
		err := dbus.Store(sig.Body, &state)
		if err == nil {
			cb(state)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *manager) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// signal DeviceAdded

func (v *manager) ConnectDeviceAdded(cb func(device_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var device_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &device_path)
		if err == nil {
			cb(device_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceRemoved

func (v *manager) ConnectDeviceRemoved(cb func(device_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var device_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &device_path)
		if err == nil {
			cb(device_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Devices ao

func (v *manager) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

// property AllDevices ao

func (v *manager) AllDevices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "AllDevices",
	}
}

// property NetworkingEnabled b

func (v *manager) NetworkingEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NetworkingEnabled",
	}
}

// property WirelessEnabled b

func (v *manager) WirelessEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WirelessEnabled",
	}
}

// property WirelessHardwareEnabled b

func (v *manager) WirelessHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WirelessHardwareEnabled",
	}
}

// property WwanEnabled b

func (v *manager) WwanEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WwanEnabled",
	}
}

// property WwanHardwareEnabled b

func (v *manager) WwanHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WwanHardwareEnabled",
	}
}

// property WimaxEnabled b

func (v *manager) WimaxEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WimaxEnabled",
	}
}

// property WimaxHardwareEnabled b

func (v *manager) WimaxHardwareEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WimaxHardwareEnabled",
	}
}

// property ActiveConnections ao

func (v *manager) ActiveConnections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "ActiveConnections",
	}
}

// property PrimaryConnection o

func (v *manager) PrimaryConnection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "PrimaryConnection",
	}
}

// property PrimaryConnectionType s

func (v *manager) PrimaryConnectionType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PrimaryConnectionType",
	}
}

// property Metered u

func (v *manager) Metered() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Metered",
	}
}

// property ActivatingConnection o

func (v *manager) ActivatingConnection() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "ActivatingConnection",
	}
}

// property Startup b

func (v *manager) Startup() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Startup",
	}
}

// property Version s

func (v *manager) Version() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Version",
	}
}

// property Capabilities au

func (v *manager) Capabilities() proxy.PropUint32Array {
	return proxy.PropUint32Array{
		Impl: v,
		Name: "Capabilities",
	}
}

// property State u

func (v *manager) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}

// property Connectivity u

func (v *manager) Connectivity() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Connectivity",
	}
}

// property GlobalDnsConfiguration a{sv}

func (v *manager) GlobalDnsConfiguration() PropManagerGlobalDnsConfiguration {
	return PropManagerGlobalDnsConfiguration{
		Impl: v,
	}
}

type PropManagerGlobalDnsConfiguration struct {
	Impl proxy.Implementer
}

func (p PropManagerGlobalDnsConfiguration) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"GlobalDnsConfiguration", &value)
	return
}

func (p PropManagerGlobalDnsConfiguration) Set(flags dbus.Flags, value map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), "GlobalDnsConfiguration", value)
}

func (p PropManagerGlobalDnsConfiguration) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
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
		"GlobalDnsConfiguration", cb0)
}

type PPP struct {
	ppp // interface org.freedesktop.NetworkManager.PPP
	proxy.Object
}

func NewPPP(conn *dbus.Conn, path dbus.ObjectPath) (*PPP, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(PPP)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type ppp struct{}

func (v *ppp) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ppp) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.PPP"
}

// method NeedSecrets

func (v *ppp) GoNeedSecrets(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NeedSecrets", flags, ch)
}

func (*ppp) StoreNeedSecrets(call *dbus.Call) (username string, password string, err error) {
	err = call.Store(&username, &password)
	return
}

func (v *ppp) NeedSecrets(flags dbus.Flags) (username string, password string, err error) {
	return v.StoreNeedSecrets(
		<-v.GoNeedSecrets(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetIp4Config

func (v *ppp) GoSetIp4Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp4Config", flags, ch, config)
}

func (v *ppp) SetIp4Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp4Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp6Config

func (v *ppp) GoSetIp6Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp6Config", flags, ch, config)
}

func (v *ppp) SetIp6Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp6Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetState

func (v *ppp) GoSetState(flags dbus.Flags, ch chan *dbus.Call, state uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetState", flags, ch, state)
}

func (v *ppp) SetState(flags dbus.Flags, state uint32) error {
	return (<-v.GoSetState(flags, make(chan *dbus.Call, 1), state).Done).Err
}

type SecretAgent struct {
	secretAgent // interface org.freedesktop.NetworkManager.SecretAgent
	proxy.Object
}

func NewSecretAgent(conn *dbus.Conn, path dbus.ObjectPath) (*SecretAgent, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(SecretAgent)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type secretAgent struct{}

func (v *secretAgent) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*secretAgent) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.SecretAgent"
}

// method GetSecrets

func (v *secretAgent) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath, setting_name string, hints []string, flags0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, connection, connection_path, setting_name, hints, flags0)
}

func (*secretAgent) StoreGetSecrets(call *dbus.Call) (secrets map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&secrets)
	return
}

func (v *secretAgent) GetSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath, setting_name string, hints []string, flags0 uint32) (secrets map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path, setting_name, hints, flags0).Done)
}

// method CancelGetSecrets

func (v *secretAgent) GoCancelGetSecrets(flags dbus.Flags, ch chan *dbus.Call, connection_path dbus.ObjectPath, setting_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelGetSecrets", flags, ch, connection_path, setting_name)
}

func (v *secretAgent) CancelGetSecrets(flags dbus.Flags, connection_path dbus.ObjectPath, setting_name string) error {
	return (<-v.GoCancelGetSecrets(flags, make(chan *dbus.Call, 1), connection_path, setting_name).Done).Err
}

// method SaveSecrets

func (v *secretAgent) GoSaveSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveSecrets", flags, ch, connection, connection_path)
}

func (v *secretAgent) SaveSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) error {
	return (<-v.GoSaveSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path).Done).Err
}

// method DeleteSecrets

func (v *secretAgent) GoDeleteSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteSecrets", flags, ch, connection, connection_path)
}

func (v *secretAgent) DeleteSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) error {
	return (<-v.GoDeleteSecrets(flags, make(chan *dbus.Call, 1), connection, connection_path).Done).Err
}

type ConnectionSettings struct {
	connectionSettings // interface org.freedesktop.NetworkManager.Settings.Connection
	proxy.Object
}

func NewConnectionSettings(conn *dbus.Conn, path dbus.ObjectPath) (*ConnectionSettings, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(ConnectionSettings)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type connectionSettings struct{}

func (v *connectionSettings) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*connectionSettings) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Settings.Connection"
}

// method Update

func (v *connectionSettings) GoUpdate(flags dbus.Flags, ch chan *dbus.Call, properties map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Update", flags, ch, properties)
}

func (v *connectionSettings) Update(flags dbus.Flags, properties map[string]map[string]dbus.Variant) error {
	return (<-v.GoUpdate(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

// method UpdateUnsaved

func (v *connectionSettings) GoUpdateUnsaved(flags dbus.Flags, ch chan *dbus.Call, properties map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateUnsaved", flags, ch, properties)
}

func (v *connectionSettings) UpdateUnsaved(flags dbus.Flags, properties map[string]map[string]dbus.Variant) error {
	return (<-v.GoUpdateUnsaved(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

// method Delete

func (v *connectionSettings) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (v *connectionSettings) Delete(flags dbus.Flags) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetSettings

func (v *connectionSettings) GoGetSettings(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSettings", flags, ch)
}

func (*connectionSettings) StoreGetSettings(call *dbus.Call) (settings map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&settings)
	return
}

func (v *connectionSettings) GetSettings(flags dbus.Flags) (settings map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSettings(
		<-v.GoGetSettings(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSecrets

func (v *connectionSettings) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, setting_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, setting_name)
}

func (*connectionSettings) StoreGetSecrets(call *dbus.Call) (secrets map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&secrets)
	return
}

func (v *connectionSettings) GetSecrets(flags dbus.Flags, setting_name string) (secrets map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), setting_name).Done)
}

// method ClearSecrets

func (v *connectionSettings) GoClearSecrets(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearSecrets", flags, ch)
}

func (v *connectionSettings) ClearSecrets(flags dbus.Flags) error {
	return (<-v.GoClearSecrets(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *connectionSettings) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *connectionSettings) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Updated

func (v *connectionSettings) ConnectUpdated(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Updated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Updated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Removed

func (v *connectionSettings) ConnectRemoved(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Removed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Removed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PropertiesChanged

func (v *connectionSettings) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// property Unsaved b

func (v *connectionSettings) Unsaved() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Unsaved",
	}
}

type Settings struct {
	settings // interface org.freedesktop.NetworkManager.Settings
	proxy.Object
}

func NewSettings(conn *dbus.Conn) *Settings {
	obj := new(Settings)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")
	return obj
}

type settings struct{}

func (v *settings) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*settings) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.Settings"
}

// method ListConnections

func (v *settings) GoListConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListConnections", flags, ch)
}

func (*settings) StoreListConnections(call *dbus.Call) (connections []dbus.ObjectPath, err error) {
	err = call.Store(&connections)
	return
}

func (v *settings) ListConnections(flags dbus.Flags) (connections []dbus.ObjectPath, err error) {
	return v.StoreListConnections(
		<-v.GoListConnections(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetConnectionByUuid

func (v *settings) GoGetConnectionByUuid(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionByUuid", flags, ch, uuid)
}

func (*settings) StoreGetConnectionByUuid(call *dbus.Call) (connection dbus.ObjectPath, err error) {
	err = call.Store(&connection)
	return
}

func (v *settings) GetConnectionByUuid(flags dbus.Flags, uuid string) (connection dbus.ObjectPath, err error) {
	return v.StoreGetConnectionByUuid(
		<-v.GoGetConnectionByUuid(flags, make(chan *dbus.Call, 1), uuid).Done)
}

// method AddConnection

func (v *settings) GoAddConnection(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddConnection", flags, ch, connection)
}

func (*settings) StoreAddConnection(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *settings) AddConnection(flags dbus.Flags, connection map[string]map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddConnection(
		<-v.GoAddConnection(flags, make(chan *dbus.Call, 1), connection).Done)
}

// method AddConnectionUnsaved

func (v *settings) GoAddConnectionUnsaved(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddConnectionUnsaved", flags, ch, connection)
}

func (*settings) StoreAddConnectionUnsaved(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *settings) AddConnectionUnsaved(flags dbus.Flags, connection map[string]map[string]dbus.Variant) (path dbus.ObjectPath, err error) {
	return v.StoreAddConnectionUnsaved(
		<-v.GoAddConnectionUnsaved(flags, make(chan *dbus.Call, 1), connection).Done)
}

// method LoadConnections

func (v *settings) GoLoadConnections(flags dbus.Flags, ch chan *dbus.Call, filenames []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LoadConnections", flags, ch, filenames)
}

func (*settings) StoreLoadConnections(call *dbus.Call) (status bool, failures []string, err error) {
	err = call.Store(&status, &failures)
	return
}

func (v *settings) LoadConnections(flags dbus.Flags, filenames []string) (status bool, failures []string, err error) {
	return v.StoreLoadConnections(
		<-v.GoLoadConnections(flags, make(chan *dbus.Call, 1), filenames).Done)
}

// method ReloadConnections

func (v *settings) GoReloadConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadConnections", flags, ch)
}

func (*settings) StoreReloadConnections(call *dbus.Call) (status bool, err error) {
	err = call.Store(&status)
	return
}

func (v *settings) ReloadConnections(flags dbus.Flags) (status bool, err error) {
	return v.StoreReloadConnections(
		<-v.GoReloadConnections(flags, make(chan *dbus.Call, 1)).Done)
}

// method SaveHostname

func (v *settings) GoSaveHostname(flags dbus.Flags, ch chan *dbus.Call, hostname string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveHostname", flags, ch, hostname)
}

func (v *settings) SaveHostname(flags dbus.Flags, hostname string) error {
	return (<-v.GoSaveHostname(flags, make(chan *dbus.Call, 1), hostname).Done).Err
}

// signal PropertiesChanged

func (v *settings) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// signal NewConnection

func (v *settings) ConnectNewConnection(cb func(connection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NewConnection", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NewConnection",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var connection dbus.ObjectPath
		err := dbus.Store(sig.Body, &connection)
		if err == nil {
			cb(connection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ConnectionRemoved

func (v *settings) ConnectConnectionRemoved(cb func(connection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ConnectionRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ConnectionRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var connection dbus.ObjectPath
		err := dbus.Store(sig.Body, &connection)
		if err == nil {
			cb(connection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Connections ao

func (v *settings) Connections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Connections",
	}
}

// property Hostname s

func (v *settings) Hostname() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Hostname",
	}
}

// property CanModify b

func (v *settings) CanModify() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanModify",
	}
}

type VpnConnection struct {
	vpnConnection // interface org.freedesktop.NetworkManager.VPN.Connection
	proxy.Object
}

func NewVpnConnection(conn *dbus.Conn, path dbus.ObjectPath) (*VpnConnection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(VpnConnection)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type vpnConnection struct{}

func (v *vpnConnection) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*vpnConnection) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.VPN.Connection"
}

// signal PropertiesChanged

func (v *vpnConnection) ConnectPropertiesChanged(cb func(properties map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
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

// signal VpnStateChanged

func (v *vpnConnection) ConnectVpnStateChanged(cb func(state uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VpnStateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VpnStateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var state uint32
		var reason uint32
		err := dbus.Store(sig.Body, &state, &reason)
		if err == nil {
			cb(state, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property VpnState u

func (v *vpnConnection) VpnState() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "VpnState",
	}
}

// property Banner s

func (v *vpnConnection) Banner() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Banner",
	}
}

type VpnPlugin struct {
	vpnPlugin // interface org.freedesktop.NetworkManager.VPN.Plugin
	proxy.Object
}

func NewVpnPlugin(conn *dbus.Conn, path dbus.ObjectPath) (*VpnPlugin, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(VpnPlugin)
	obj.Object.Init_(conn, "org.freedesktop.NetworkManager", path)
	return obj, nil
}

type vpnPlugin struct{}

func (v *vpnPlugin) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*vpnPlugin) GetInterfaceName_() string {
	return "org.freedesktop.NetworkManager.VPN.Plugin"
}

// method Connect

func (v *vpnPlugin) GoConnect(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch, connection)
}

func (v *vpnPlugin) Connect(flags dbus.Flags, connection map[string]map[string]dbus.Variant) error {
	return (<-v.GoConnect(flags, make(chan *dbus.Call, 1), connection).Done).Err
}

// method ConnectInteractive

func (v *vpnPlugin) GoConnectInteractive(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant, details map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ConnectInteractive", flags, ch, connection, details)
}

func (v *vpnPlugin) ConnectInteractive(flags dbus.Flags, connection map[string]map[string]dbus.Variant, details map[string]dbus.Variant) error {
	return (<-v.GoConnectInteractive(flags, make(chan *dbus.Call, 1), connection, details).Done).Err
}

// method NeedSecrets

func (v *vpnPlugin) GoNeedSecrets(flags dbus.Flags, ch chan *dbus.Call, settings map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NeedSecrets", flags, ch, settings)
}

func (*vpnPlugin) StoreNeedSecrets(call *dbus.Call) (setting_name string, err error) {
	err = call.Store(&setting_name)
	return
}

func (v *vpnPlugin) NeedSecrets(flags dbus.Flags, settings map[string]map[string]dbus.Variant) (setting_name string, err error) {
	return v.StoreNeedSecrets(
		<-v.GoNeedSecrets(flags, make(chan *dbus.Call, 1), settings).Done)
}

// method Disconnect

func (v *vpnPlugin) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *vpnPlugin) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetConfig

func (v *vpnPlugin) GoSetConfig(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetConfig", flags, ch, config)
}

func (v *vpnPlugin) SetConfig(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetConfig(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp4Config

func (v *vpnPlugin) GoSetIp4Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp4Config", flags, ch, config)
}

func (v *vpnPlugin) SetIp4Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp4Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetIp6Config

func (v *vpnPlugin) GoSetIp6Config(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIp6Config", flags, ch, config)
}

func (v *vpnPlugin) SetIp6Config(flags dbus.Flags, config map[string]dbus.Variant) error {
	return (<-v.GoSetIp6Config(flags, make(chan *dbus.Call, 1), config).Done).Err
}

// method SetFailure

func (v *vpnPlugin) GoSetFailure(flags dbus.Flags, ch chan *dbus.Call, reason string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFailure", flags, ch, reason)
}

func (v *vpnPlugin) SetFailure(flags dbus.Flags, reason string) error {
	return (<-v.GoSetFailure(flags, make(chan *dbus.Call, 1), reason).Done).Err
}

// method NewSecrets

func (v *vpnPlugin) GoNewSecrets(flags dbus.Flags, ch chan *dbus.Call, connection map[string]map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NewSecrets", flags, ch, connection)
}

func (v *vpnPlugin) NewSecrets(flags dbus.Flags, connection map[string]map[string]dbus.Variant) error {
	return (<-v.GoNewSecrets(flags, make(chan *dbus.Call, 1), connection).Done).Err
}

// signal StateChanged

func (v *vpnPlugin) ConnectStateChanged(cb func(state uint32)) (dbusutil.SignalHandlerId, error) {
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
		var state uint32
		err := dbus.Store(sig.Body, &state)
		if err == nil {
			cb(state)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SecretsRequired

func (v *vpnPlugin) ConnectSecretsRequired(cb func(message string, secrets []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SecretsRequired", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SecretsRequired",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var message string
		var secrets []string
		err := dbus.Store(sig.Body, &message, &secrets)
		if err == nil {
			cb(message, secrets)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Config

func (v *vpnPlugin) ConnectConfig(cb func(config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &config)
		if err == nil {
			cb(config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Ip4Config

func (v *vpnPlugin) ConnectIp4Config(cb func(ip4config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Ip4Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Ip4Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip4config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &ip4config)
		if err == nil {
			cb(ip4config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Ip6Config

func (v *vpnPlugin) ConnectIp6Config(cb func(ip6config map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Ip6Config", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Ip6Config",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip6config map[string]dbus.Variant
		err := dbus.Store(sig.Body, &ip6config)
		if err == nil {
			cb(ip6config)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LoginBanner

func (v *vpnPlugin) ConnectLoginBanner(cb func(banner string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LoginBanner", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LoginBanner",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var banner string
		err := dbus.Store(sig.Body, &banner)
		if err == nil {
			cb(banner)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Failure

func (v *vpnPlugin) ConnectFailure(cb func(reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Failure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Failure",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var reason uint32
		err := dbus.Store(sig.Body, &reason)
		if err == nil {
			cb(reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property State u

func (v *vpnPlugin) State() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "State",
	}
}
