package geoclue2

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

type Manager struct {
	manager // interface org.freedesktop.GeoClue2.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.GeoClue2", "/org/freedesktop/GeoClue2/Manager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Manager"
}

// method GetClient

func (v *manager) GoGetClient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetClient", flags, ch)
}

func (v *manager) GoGetClientWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetClient", flags, ch)
}

func (*manager) StoreGetClient(call *dbus.Call) (client dbus.ObjectPath, err error) {
	err = call.Store(&client)
	return
}

func (v *manager) GetClient(flags dbus.Flags) (client dbus.ObjectPath, err error) {
	return v.StoreGetClient(
		<-v.GoGetClient(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) GetClientWithTimeout(timeout time.Duration, flags dbus.Flags) (client dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetClientWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetClient(call)
}

// method AddAgent

func (v *manager) GoAddAgent(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAgent", flags, ch, id)
}

func (v *manager) GoAddAgentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddAgent", flags, ch, id)
}

func (v *manager) AddAgent(flags dbus.Flags, id string) error {
	return (<-v.GoAddAgent(flags, make(chan *dbus.Call, 1), id).Done).Err
}

func (v *manager) AddAgentWithTimeout(timeout time.Duration, flags dbus.Flags, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddAgentWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property InUse b

func (v *manager) InUse() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "InUse",
	}
}

// property AvailableAccuracyLevel u

func (v *manager) AvailableAccuracyLevel() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "AvailableAccuracyLevel",
	}
}

type Client struct {
	client // interface org.freedesktop.GeoClue2.Client
	proxy.Object
}

func NewClient(conn *dbus.Conn, path dbus.ObjectPath) (*Client, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Client)
	obj.Object.Init_(conn, "org.freedesktop.GeoClue2", path)
	return obj, nil
}

type client struct{}

func (v *client) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*client) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Client"
}

// method Start

func (v *client) GoStart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start", flags, ch)
}

func (v *client) GoStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Start", flags, ch)
}

func (v *client) Start(flags dbus.Flags) error {
	return (<-v.GoStart(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *client) StartWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Stop

func (v *client) GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Stop", flags, ch)
}

func (v *client) GoStopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Stop", flags, ch)
}

func (v *client) Stop(flags dbus.Flags) error {
	return (<-v.GoStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *client) StopWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal LocationUpdated

func (v *client) ConnectLocationUpdated(cb func(old dbus.ObjectPath, new dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LocationUpdated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LocationUpdated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old dbus.ObjectPath
		var new dbus.ObjectPath
		err := dbus.Store(sig.Body, &old, &new)
		if err == nil {
			cb(old, new)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Location o

func (v *client) Location() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Location",
	}
}

// property DistanceThreshold u

func (v *client) DistanceThreshold() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "DistanceThreshold",
	}
}

// property TimeThreshold u

func (v *client) TimeThreshold() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "TimeThreshold",
	}
}

// property DesktopId s

func (v *client) DesktopId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DesktopId",
	}
}

// property RequestedAccuracyLevel u

func (v *client) RequestedAccuracyLevel() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RequestedAccuracyLevel",
	}
}

// property Active b

func (v *client) Active() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Active",
	}
}

type Location struct {
	location // interface org.freedesktop.GeoClue2.Location
	proxy.Object
}

func NewLocation(conn *dbus.Conn, path dbus.ObjectPath) (*Location, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Location)
	obj.Object.Init_(conn, "org.freedesktop.GeoClue2", path)
	return obj, nil
}

type location struct{}

func (v *location) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*location) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Location"
}

// property Latitude d

func (v *location) Latitude() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Latitude",
	}
}

// property Longitude d

func (v *location) Longitude() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Longitude",
	}
}

// property Accuracy d

func (v *location) Accuracy() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Accuracy",
	}
}

// property Altitude d

func (v *location) Altitude() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Altitude",
	}
}

// property Speed d

func (v *location) Speed() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Speed",
	}
}

// property Heading d

func (v *location) Heading() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Heading",
	}
}

// property Description s

func (v *location) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property Timestamp (tt)

func (v *location) Timestamp() PropTimestamp {
	return PropTimestamp{
		Impl: v,
	}
}

type PropTimestamp struct {
	Impl proxy.Implementer
}

func (p PropTimestamp) Get(flags dbus.Flags) (value Timestamp, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Timestamp", &value)
	return
}

func (p PropTimestamp) ConnectChanged(cb func(hasValue bool, value Timestamp)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v Timestamp
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, Timestamp{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Timestamp", cb0)
}
