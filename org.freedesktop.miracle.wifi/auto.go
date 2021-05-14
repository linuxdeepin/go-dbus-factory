package wifi

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

type Wifi struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.Object
}

func NewWifi(conn *dbus.Conn) *Wifi {
	obj := new(Wifi)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wifi", "/org/freedesktop/miracle/wifi")
	return obj
}

type Link struct {
	link // interface org.freedesktop.miracle.wifi.Link
	proxy.Object
}

func NewLink(conn *dbus.Conn, path dbus.ObjectPath) (*Link, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Link)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wifi", path)
	return obj, nil
}

type link struct{}

func (v *link) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*link) GetInterfaceName_() string {
	return "org.freedesktop.miracle.wifi.Link"
}

// method Manage

func (v *link) GoManage(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Manage", flags, ch)
}

func (v *link) GoManageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Manage", flags, ch)
}

func (v *link) Manage(flags dbus.Flags) error {
	return (<-v.GoManage(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *link) ManageWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoManageWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Unmanage

func (v *link) GoUnmanage(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unmanage", flags, ch)
}

func (v *link) GoUnmanageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Unmanage", flags, ch)
}

func (v *link) Unmanage(flags dbus.Flags) error {
	return (<-v.GoUnmanage(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *link) UnmanageWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnmanageWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property InterfaceIndex u

func (v *link) InterfaceIndex() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "InterfaceIndex",
	}
}

// property MACAddress s

func (v *link) MACAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MACAddress",
	}
}

// property InterfaceName s

func (v *link) InterfaceName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "InterfaceName",
	}
}

// property FriendlyName s

func (v *link) FriendlyName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FriendlyName",
	}
}

// property Managed b

func (v *link) Managed() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Managed",
	}
}

// property P2PState i

func (v *link) P2PState() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "P2PState",
	}
}

// property P2PScanning b

func (v *link) P2PScanning() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "P2PScanning",
	}
}

// property WfdSubelements s

func (v *link) WfdSubelements() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WfdSubelements",
	}
}

type Peer struct {
	peer // interface org.freedesktop.miracle.wifi.Peer
	proxy.Object
}

func NewPeer(conn *dbus.Conn, path dbus.ObjectPath) (*Peer, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Peer)
	obj.Object.Init_(conn, "org.freedesktop.miracle.wifi", path)
	return obj, nil
}

type peer struct{}

func (v *peer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*peer) GetInterfaceName_() string {
	return "org.freedesktop.miracle.wifi.Peer"
}

// method Connect

func (v *peer) GoConnect(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Connect", flags, ch, arg0, arg1)
}

func (v *peer) GoConnectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Connect", flags, ch, arg0, arg1)
}

func (v *peer) Connect(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoConnect(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *peer) ConnectWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoConnectWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Disconnect

func (v *peer) GoDisconnect(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *peer) GoDisconnectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Disconnect", flags, ch)
}

func (v *peer) Disconnect(flags dbus.Flags) error {
	return (<-v.GoDisconnect(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *peer) DisconnectWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisconnectWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal ProvisionDiscovery

func (v *peer) ConnectProvisionDiscovery(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProvisionDiscovery", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProvisionDiscovery",
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

// signal GoNegRequest

func (v *peer) ConnectGoNegRequest(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GoNegRequest", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GoNegRequest",
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

// signal FormationFailure

func (v *peer) ConnectFormationFailure(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "FormationFailure", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".FormationFailure",
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

// property Link o

func (v *peer) Link() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Link",
	}
}

// property P2PMac s

func (v *peer) P2PMac() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "P2PMac",
	}
}

// property FriendlyName s

func (v *peer) FriendlyName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FriendlyName",
	}
}

// property Connected b

func (v *peer) Connected() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property Interface s

func (v *peer) Interface() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Interface",
	}
}

// property LocalAddress s

func (v *peer) LocalAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "LocalAddress",
	}
}

// property RemoteAddress s

func (v *peer) RemoteAddress() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RemoteAddress",
	}
}

// property WfdSubelements s

func (v *peer) WfdSubelements() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WfdSubelements",
	}
}
