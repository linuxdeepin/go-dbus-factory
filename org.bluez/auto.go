package bluez

import "context"
import "errors"
import "fmt"
import "github.com/linuxdeepin/go-dbus-factory/object_manager"
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

type ObjectManager struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.Object
}

func NewObjectManager(conn *dbus.Conn) *ObjectManager {
	obj := new(ObjectManager)
	obj.Object.Init_(conn, "org.bluez", "/")
	return obj
}

type Manager struct {
	agentManager   // interface org.bluez.AgentManager1
	healthManager  // interface org.bluez.HealthManager1
	profileManager // interface org.bluez.ProfileManager1
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.bluez", "/org/bluez")
	return obj
}

func (obj *Manager) AgentManager() *agentManager {
	return &obj.agentManager
}

type agentManager struct{}

func (v *agentManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*agentManager) GetInterfaceName_() string {
	return "org.bluez.AgentManager1"
}

// method RegisterAgent

func (v *agentManager) GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath, capability string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAgent", flags, ch, agent, capability)
}

func (v *agentManager) GoRegisterAgentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath, capability string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RegisterAgent", flags, ch, agent, capability)
}

func (v *agentManager) RegisterAgent(flags dbus.Flags, agent dbus.ObjectPath, capability string) error {
	return (<-v.GoRegisterAgent(flags, make(chan *dbus.Call, 1), agent, capability).Done).Err
}

func (v *agentManager) RegisterAgentWithTimeout(timeout time.Duration, flags dbus.Flags, agent dbus.ObjectPath, capability string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterAgentWithContext(ctx, flags, make(chan *dbus.Call, 1), agent, capability).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnregisterAgent

func (v *agentManager) GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterAgent", flags, ch, agent)
}

func (v *agentManager) GoUnregisterAgentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnregisterAgent", flags, ch, agent)
}

func (v *agentManager) UnregisterAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	return (<-v.GoUnregisterAgent(flags, make(chan *dbus.Call, 1), agent).Done).Err
}

func (v *agentManager) UnregisterAgentWithTimeout(timeout time.Duration, flags dbus.Flags, agent dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterAgentWithContext(ctx, flags, make(chan *dbus.Call, 1), agent).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestDefaultAgent

func (v *agentManager) GoRequestDefaultAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestDefaultAgent", flags, ch, agent)
}

func (v *agentManager) GoRequestDefaultAgentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestDefaultAgent", flags, ch, agent)
}

func (v *agentManager) RequestDefaultAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	return (<-v.GoRequestDefaultAgent(flags, make(chan *dbus.Call, 1), agent).Done).Err
}

func (v *agentManager) RequestDefaultAgentWithTimeout(timeout time.Duration, flags dbus.Flags, agent dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestDefaultAgentWithContext(ctx, flags, make(chan *dbus.Call, 1), agent).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

func (obj *Manager) HealthManager() *healthManager {
	return &obj.healthManager
}

type healthManager struct{}

func (v *healthManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*healthManager) GetInterfaceName_() string {
	return "org.bluez.HealthManager1"
}

// method CreateApplication

func (v *healthManager) GoCreateApplication(flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateApplication", flags, ch, config)
}

func (v *healthManager) GoCreateApplicationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, config map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CreateApplication", flags, ch, config)
}

func (*healthManager) StoreCreateApplication(call *dbus.Call) (application dbus.ObjectPath, err error) {
	err = call.Store(&application)
	return
}

func (v *healthManager) CreateApplication(flags dbus.Flags, config map[string]dbus.Variant) (application dbus.ObjectPath, err error) {
	return v.StoreCreateApplication(
		<-v.GoCreateApplication(flags, make(chan *dbus.Call, 1), config).Done)
}

func (v *healthManager) CreateApplicationWithTimeout(timeout time.Duration, flags dbus.Flags, config map[string]dbus.Variant) (application dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCreateApplicationWithContext(ctx, flags, make(chan *dbus.Call, 1), config).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCreateApplication(call)
}

// method DestroyApplication

func (v *healthManager) GoDestroyApplication(flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DestroyApplication", flags, ch, application)
}

func (v *healthManager) GoDestroyApplicationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DestroyApplication", flags, ch, application)
}

func (v *healthManager) DestroyApplication(flags dbus.Flags, application dbus.ObjectPath) error {
	return (<-v.GoDestroyApplication(flags, make(chan *dbus.Call, 1), application).Done).Err
}

func (v *healthManager) DestroyApplicationWithTimeout(timeout time.Duration, flags dbus.Flags, application dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDestroyApplicationWithContext(ctx, flags, make(chan *dbus.Call, 1), application).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

func (obj *Manager) ProfileManager() *profileManager {
	return &obj.profileManager
}

type profileManager struct{}

func (v *profileManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*profileManager) GetInterfaceName_() string {
	return "org.bluez.ProfileManager1"
}

// method RegisterProfile

func (v *profileManager) GoRegisterProfile(flags dbus.Flags, ch chan *dbus.Call, profile dbus.ObjectPath, UUID string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterProfile", flags, ch, profile, UUID, options)
}

func (v *profileManager) GoRegisterProfileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, profile dbus.ObjectPath, UUID string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RegisterProfile", flags, ch, profile, UUID, options)
}

func (v *profileManager) RegisterProfile(flags dbus.Flags, profile dbus.ObjectPath, UUID string, options map[string]dbus.Variant) error {
	return (<-v.GoRegisterProfile(flags, make(chan *dbus.Call, 1), profile, UUID, options).Done).Err
}

func (v *profileManager) RegisterProfileWithTimeout(timeout time.Duration, flags dbus.Flags, profile dbus.ObjectPath, UUID string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterProfileWithContext(ctx, flags, make(chan *dbus.Call, 1), profile, UUID, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnregisterProfile

func (v *profileManager) GoUnregisterProfile(flags dbus.Flags, ch chan *dbus.Call, profile dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterProfile", flags, ch, profile)
}

func (v *profileManager) GoUnregisterProfileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, profile dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnregisterProfile", flags, ch, profile)
}

func (v *profileManager) UnregisterProfile(flags dbus.Flags, profile dbus.ObjectPath) error {
	return (<-v.GoUnregisterProfile(flags, make(chan *dbus.Call, 1), profile).Done).Err
}

func (v *profileManager) UnregisterProfileWithTimeout(timeout time.Duration, flags dbus.Flags, profile dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterProfileWithContext(ctx, flags, make(chan *dbus.Call, 1), profile).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

type HCI struct {
	adapter       // interface org.bluez.Adapter1
	gattManager   // interface org.bluez.GattManager1
	media         // interface org.bluez.Media1
	networkServer // interface org.bluez.NetworkServer1
	proxy.Object
}

func NewHCI(conn *dbus.Conn, path dbus.ObjectPath) (*HCI, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(HCI)
	obj.Object.Init_(conn, "org.bluez", path)
	return obj, nil
}

func (obj *HCI) Adapter() *adapter {
	return &obj.adapter
}

type adapter struct{}

func (v *adapter) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*adapter) GetInterfaceName_() string {
	return "org.bluez.Adapter1"
}

// method StartDiscovery

func (v *adapter) GoStartDiscovery(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartDiscovery", flags, ch)
}

func (v *adapter) GoStartDiscoveryWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartDiscovery", flags, ch)
}

func (v *adapter) StartDiscovery(flags dbus.Flags) error {
	return (<-v.GoStartDiscovery(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *adapter) StartDiscoveryWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartDiscoveryWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetDiscoveryFilter

func (v *adapter) GoSetDiscoveryFilter(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDiscoveryFilter", flags, ch, properties)
}

func (v *adapter) GoSetDiscoveryFilterWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDiscoveryFilter", flags, ch, properties)
}

func (v *adapter) SetDiscoveryFilter(flags dbus.Flags, properties map[string]dbus.Variant) error {
	return (<-v.GoSetDiscoveryFilter(flags, make(chan *dbus.Call, 1), properties).Done).Err
}

func (v *adapter) SetDiscoveryFilterWithTimeout(timeout time.Duration, flags dbus.Flags, properties map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDiscoveryFilterWithContext(ctx, flags, make(chan *dbus.Call, 1), properties).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method StopDiscovery

func (v *adapter) GoStopDiscovery(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopDiscovery", flags, ch)
}

func (v *adapter) GoStopDiscoveryWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StopDiscovery", flags, ch)
}

func (v *adapter) StopDiscovery(flags dbus.Flags) error {
	return (<-v.GoStopDiscovery(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *adapter) StopDiscoveryWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopDiscoveryWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RemoveDevice

func (v *adapter) GoRemoveDevice(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveDevice", flags, ch, device)
}

func (v *adapter) GoRemoveDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveDevice", flags, ch, device)
}

func (v *adapter) RemoveDevice(flags dbus.Flags, device dbus.ObjectPath) error {
	return (<-v.GoRemoveDevice(flags, make(chan *dbus.Call, 1), device).Done).Err
}

func (v *adapter) RemoveDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, device dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), device).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Address s

func (v *adapter) Address() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Address",
	}
}

// property Name s

func (v *adapter) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Alias s

func (v *adapter) Alias() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Alias",
	}
}

// property Class u

func (v *adapter) Class() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Class",
	}
}

// property Powered b

func (v *adapter) Powered() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Powered",
	}
}

// property Discoverable b

func (v *adapter) Discoverable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Discoverable",
	}
}

// property DiscoverableTimeout u

func (v *adapter) DiscoverableTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "DiscoverableTimeout",
	}
}

// property Pairable b

func (v *adapter) Pairable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Pairable",
	}
}

// property PairableTimeout u

func (v *adapter) PairableTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "PairableTimeout",
	}
}

// property Discovering b

func (v *adapter) Discovering() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Discovering",
	}
}

// property UUIDs as

func (v *adapter) UUIDs() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UUIDs",
	}
}

// property Modalias s

func (v *adapter) Modalias() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Modalias",
	}
}

func (obj *HCI) GattManager() *gattManager {
	return &obj.gattManager
}

type gattManager struct{}

func (v *gattManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*gattManager) GetInterfaceName_() string {
	return "org.bluez.GattManager1"
}

// method RegisterApplication

func (v *gattManager) GoRegisterApplication(flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterApplication", flags, ch, application, options)
}

func (v *gattManager) GoRegisterApplicationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RegisterApplication", flags, ch, application, options)
}

func (v *gattManager) RegisterApplication(flags dbus.Flags, application dbus.ObjectPath, options map[string]dbus.Variant) error {
	return (<-v.GoRegisterApplication(flags, make(chan *dbus.Call, 1), application, options).Done).Err
}

func (v *gattManager) RegisterApplicationWithTimeout(timeout time.Duration, flags dbus.Flags, application dbus.ObjectPath, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterApplicationWithContext(ctx, flags, make(chan *dbus.Call, 1), application, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnregisterApplication

func (v *gattManager) GoUnregisterApplication(flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterApplication", flags, ch, application)
}

func (v *gattManager) GoUnregisterApplicationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, application dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnregisterApplication", flags, ch, application)
}

func (v *gattManager) UnregisterApplication(flags dbus.Flags, application dbus.ObjectPath) error {
	return (<-v.GoUnregisterApplication(flags, make(chan *dbus.Call, 1), application).Done).Err
}

func (v *gattManager) UnregisterApplicationWithTimeout(timeout time.Duration, flags dbus.Flags, application dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterApplicationWithContext(ctx, flags, make(chan *dbus.Call, 1), application).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

func (obj *HCI) Media() *media {
	return &obj.media
}

type media struct{}

func (v *media) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*media) GetInterfaceName_() string {
	return "org.bluez.Media1"
}

// method RegisterEndpoint

func (v *media) GoRegisterEndpoint(flags dbus.Flags, ch chan *dbus.Call, endpoint dbus.ObjectPath, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterEndpoint", flags, ch, endpoint, properties)
}

func (v *media) GoRegisterEndpointWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, endpoint dbus.ObjectPath, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RegisterEndpoint", flags, ch, endpoint, properties)
}

func (v *media) RegisterEndpoint(flags dbus.Flags, endpoint dbus.ObjectPath, properties map[string]dbus.Variant) error {
	return (<-v.GoRegisterEndpoint(flags, make(chan *dbus.Call, 1), endpoint, properties).Done).Err
}

func (v *media) RegisterEndpointWithTimeout(timeout time.Duration, flags dbus.Flags, endpoint dbus.ObjectPath, properties map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterEndpointWithContext(ctx, flags, make(chan *dbus.Call, 1), endpoint, properties).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnregisterEndpoint

func (v *media) GoUnregisterEndpoint(flags dbus.Flags, ch chan *dbus.Call, endpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterEndpoint", flags, ch, endpoint)
}

func (v *media) GoUnregisterEndpointWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, endpoint dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnregisterEndpoint", flags, ch, endpoint)
}

func (v *media) UnregisterEndpoint(flags dbus.Flags, endpoint dbus.ObjectPath) error {
	return (<-v.GoUnregisterEndpoint(flags, make(chan *dbus.Call, 1), endpoint).Done).Err
}

func (v *media) UnregisterEndpointWithTimeout(timeout time.Duration, flags dbus.Flags, endpoint dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterEndpointWithContext(ctx, flags, make(chan *dbus.Call, 1), endpoint).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RegisterPlayer

func (v *media) GoRegisterPlayer(flags dbus.Flags, ch chan *dbus.Call, player dbus.ObjectPath, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterPlayer", flags, ch, player, properties)
}

func (v *media) GoRegisterPlayerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, player dbus.ObjectPath, properties map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RegisterPlayer", flags, ch, player, properties)
}

func (v *media) RegisterPlayer(flags dbus.Flags, player dbus.ObjectPath, properties map[string]dbus.Variant) error {
	return (<-v.GoRegisterPlayer(flags, make(chan *dbus.Call, 1), player, properties).Done).Err
}

func (v *media) RegisterPlayerWithTimeout(timeout time.Duration, flags dbus.Flags, player dbus.ObjectPath, properties map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterPlayerWithContext(ctx, flags, make(chan *dbus.Call, 1), player, properties).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnregisterPlayer

func (v *media) GoUnregisterPlayer(flags dbus.Flags, ch chan *dbus.Call, player dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterPlayer", flags, ch, player)
}

func (v *media) GoUnregisterPlayerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, player dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnregisterPlayer", flags, ch, player)
}

func (v *media) UnregisterPlayer(flags dbus.Flags, player dbus.ObjectPath) error {
	return (<-v.GoUnregisterPlayer(flags, make(chan *dbus.Call, 1), player).Done).Err
}

func (v *media) UnregisterPlayerWithTimeout(timeout time.Duration, flags dbus.Flags, player dbus.ObjectPath) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterPlayerWithContext(ctx, flags, make(chan *dbus.Call, 1), player).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

func (obj *HCI) NetworkServer() *networkServer {
	return &obj.networkServer
}

type networkServer struct{}

func (v *networkServer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*networkServer) GetInterfaceName_() string {
	return "org.bluez.NetworkServer1"
}

// method Register

func (v *networkServer) GoRegister(flags dbus.Flags, ch chan *dbus.Call, uuid string, bridge string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, uuid, bridge)
}

func (v *networkServer) GoRegisterWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string, bridge string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Register", flags, ch, uuid, bridge)
}

func (v *networkServer) Register(flags dbus.Flags, uuid string, bridge string) error {
	return (<-v.GoRegister(flags, make(chan *dbus.Call, 1), uuid, bridge).Done).Err
}

func (v *networkServer) RegisterWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string, bridge string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid, bridge).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Unregister

func (v *networkServer) GoUnregister(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unregister", flags, ch, uuid)
}

func (v *networkServer) GoUnregisterWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Unregister", flags, ch, uuid)
}

func (v *networkServer) Unregister(flags dbus.Flags, uuid string) error {
	return (<-v.GoUnregister(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

func (v *networkServer) UnregisterWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnregisterWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

type Device struct {
	device // interface org.bluez.Device1
	proxy.Object
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Device)
	obj.Object.Init_(conn, "org.bluez", path)
	return obj, nil
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*device) GetInterfaceName_() string {
	return "org.bluez.Device1"
}

// method Disconnect

func (v *device) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *device) GoDisconnectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *device) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) DisconnectWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisconnectWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Connect

func (v *device) GoConnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch)
}

func (v *device) GoConnectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Connect", flags, ch)
}

func (v *device) Connect(flags dbus.Flags) error {
	return (<-v.GoConnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) ConnectWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoConnectWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ConnectProfile

func (v *device) GoConnectProfile(flags dbus.Flags, ch chan *dbus.Call, UUID string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ConnectProfile", flags, ch, UUID)
}

func (v *device) GoConnectProfileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, UUID string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ConnectProfile", flags, ch, UUID)
}

func (v *device) ConnectProfile(flags dbus.Flags, UUID string) error {
	return (<-v.GoConnectProfile(flags, make(chan *dbus.Call, 1), UUID).Done).Err
}

func (v *device) ConnectProfileWithTimeout(timeout time.Duration, flags dbus.Flags, UUID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoConnectProfileWithContext(ctx, flags, make(chan *dbus.Call, 1), UUID).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DisconnectProfile

func (v *device) GoDisconnectProfile(flags dbus.Flags, ch chan *dbus.Call, UUID string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectProfile", flags, ch, UUID)
}

func (v *device) GoDisconnectProfileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, UUID string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DisconnectProfile", flags, ch, UUID)
}

func (v *device) DisconnectProfile(flags dbus.Flags, UUID string) error {
	return (<-v.GoDisconnectProfile(flags, make(chan *dbus.Call, 1), UUID).Done).Err
}

func (v *device) DisconnectProfileWithTimeout(timeout time.Duration, flags dbus.Flags, UUID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisconnectProfileWithContext(ctx, flags, make(chan *dbus.Call, 1), UUID).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Pair

func (v *device) GoPair(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pair", flags, ch)
}

func (v *device) GoPairWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Pair", flags, ch)
}

func (v *device) Pair(flags dbus.Flags) error {
	return (<-v.GoPair(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) PairWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPairWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method CancelPairing

func (v *device) GoCancelPairing(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPairing", flags, ch)
}

func (v *device) GoCancelPairingWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CancelPairing", flags, ch)
}

func (v *device) CancelPairing(flags dbus.Flags) error {
	return (<-v.GoCancelPairing(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) CancelPairingWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCancelPairingWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Address s

func (v *device) Address() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Address",
	}
}

// property Name s

func (v *device) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Alias s

func (v *device) Alias() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Alias",
	}
}

// property Class u

func (v *device) Class() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Class",
	}
}

// property Appearance q

func (v *device) Appearance() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Appearance",
	}
}

// property Icon s

func (v *device) Icon() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Icon",
	}
}

// property Paired b

func (v *device) Paired() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Paired",
	}
}

// property Trusted b

func (v *device) Trusted() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Trusted",
	}
}

// property Blocked b

func (v *device) Blocked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Blocked",
	}
}

// property LegacyPairing b

func (v *device) LegacyPairing() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "LegacyPairing",
	}
}

// property RSSI n

func (v *device) RSSI() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "RSSI",
	}
}

// property Connected b

func (v *device) Connected() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property UUIDs as

func (v *device) UUIDs() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UUIDs",
	}
}

// property Modalias s

func (v *device) Modalias() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Modalias",
	}
}

// property Adapter o

func (v *device) Adapter() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Adapter",
	}
}

// property ManufacturerData a{qv}

func (v *device) ManufacturerData() PropDeviceManufacturerData {
	return PropDeviceManufacturerData{
		Impl: v,
	}
}

type PropDeviceManufacturerData struct {
	Impl proxy.Implementer
}

func (p PropDeviceManufacturerData) Get(flags dbus.Flags) (value map[uint16]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"ManufacturerData", &value)
	return
}

func (p PropDeviceManufacturerData) ConnectChanged(cb func(hasValue bool, value map[uint16]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[uint16]dbus.Variant
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
		"ManufacturerData", cb0)
}

// property ServiceData a{sv}

func (v *device) ServiceData() PropDeviceServiceData {
	return PropDeviceServiceData{
		Impl: v,
	}
}

type PropDeviceServiceData struct {
	Impl proxy.Implementer
}

func (p PropDeviceServiceData) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"ServiceData", &value)
	return
}

func (p PropDeviceServiceData) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
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
		"ServiceData", cb0)
}

// property TxPower n

func (v *device) TxPower() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "TxPower",
	}
}

// property ServicesResolved b

func (v *device) ServicesResolved() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ServicesResolved",
	}
}
