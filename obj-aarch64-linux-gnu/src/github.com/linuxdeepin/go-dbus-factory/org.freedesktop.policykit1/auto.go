package policykit1

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

type Authority struct {
	authority // interface org.freedesktop.PolicyKit1.Authority
	proxy.Object
}

func NewAuthority(conn *dbus.Conn) *Authority {
	obj := new(Authority)
	obj.Object.Init_(conn, "org.freedesktop.PolicyKit1", "/org/freedesktop/PolicyKit1/Authority")
	return obj
}

type authority struct{}

func (v *authority) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*authority) GetInterfaceName_() string {
	return "org.freedesktop.PolicyKit1.Authority"
}

// method EnumerateActions

func (v *authority) GoEnumerateActions(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnumerateActions", flags, ch, locale)
}

func (*authority) StoreEnumerateActions(call *dbus.Call) (action_descriptions []ActionDescription, err error) {
	err = call.Store(&action_descriptions)
	return
}

func (v *authority) EnumerateActions(flags dbus.Flags, locale string) (action_descriptions []ActionDescription, err error) {
	return v.StoreEnumerateActions(
		<-v.GoEnumerateActions(flags, make(chan *dbus.Call, 1), locale).Done)
}

// method CheckAuthorization

func (v *authority) GoCheckAuthorization(flags dbus.Flags, ch chan *dbus.Call, subject Subject, action_id string, details map[string]string, flags0 uint32, cancellation_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckAuthorization", flags, ch, subject, action_id, details, flags0, cancellation_id)
}

func (*authority) StoreCheckAuthorization(call *dbus.Call) (result AuthorizationResult, err error) {
	err = call.Store(&result)
	return
}

func (v *authority) CheckAuthorization(flags dbus.Flags, subject Subject, action_id string, details map[string]string, flags0 uint32, cancellation_id string) (result AuthorizationResult, err error) {
	return v.StoreCheckAuthorization(
		<-v.GoCheckAuthorization(flags, make(chan *dbus.Call, 1), subject, action_id, details, flags0, cancellation_id).Done)
}

// method CancelCheckAuthorization

func (v *authority) GoCancelCheckAuthorization(flags dbus.Flags, ch chan *dbus.Call, cancellation_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelCheckAuthorization", flags, ch, cancellation_id)
}

func (v *authority) CancelCheckAuthorization(flags dbus.Flags, cancellation_id string) error {
	return (<-v.GoCancelCheckAuthorization(flags, make(chan *dbus.Call, 1), cancellation_id).Done).Err
}

// method RegisterAuthenticationAgent

func (v *authority) GoRegisterAuthenticationAgent(flags dbus.Flags, ch chan *dbus.Call, subject Subject, locale string, object_path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAuthenticationAgent", flags, ch, subject, locale, object_path)
}

func (v *authority) RegisterAuthenticationAgent(flags dbus.Flags, subject Subject, locale string, object_path string) error {
	return (<-v.GoRegisterAuthenticationAgent(flags, make(chan *dbus.Call, 1), subject, locale, object_path).Done).Err
}

// method RegisterAuthenticationAgentWithOptions

func (v *authority) GoRegisterAuthenticationAgentWithOptions(flags dbus.Flags, ch chan *dbus.Call, subject Subject, locale string, object_path string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAuthenticationAgentWithOptions", flags, ch, subject, locale, object_path, options)
}

func (v *authority) RegisterAuthenticationAgentWithOptions(flags dbus.Flags, subject Subject, locale string, object_path string, options map[string]dbus.Variant) error {
	return (<-v.GoRegisterAuthenticationAgentWithOptions(flags, make(chan *dbus.Call, 1), subject, locale, object_path, options).Done).Err
}

// method UnregisterAuthenticationAgent

func (v *authority) GoUnregisterAuthenticationAgent(flags dbus.Flags, ch chan *dbus.Call, subject Subject, object_path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterAuthenticationAgent", flags, ch, subject, object_path)
}

func (v *authority) UnregisterAuthenticationAgent(flags dbus.Flags, subject Subject, object_path string) error {
	return (<-v.GoUnregisterAuthenticationAgent(flags, make(chan *dbus.Call, 1), subject, object_path).Done).Err
}

// method AuthenticationAgentResponse

func (v *authority) GoAuthenticationAgentResponse(flags dbus.Flags, ch chan *dbus.Call, cookie string, identity Identity) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AuthenticationAgentResponse", flags, ch, cookie, identity)
}

func (v *authority) AuthenticationAgentResponse(flags dbus.Flags, cookie string, identity Identity) error {
	return (<-v.GoAuthenticationAgentResponse(flags, make(chan *dbus.Call, 1), cookie, identity).Done).Err
}

// method AuthenticationAgentResponse2

func (v *authority) GoAuthenticationAgentResponse2(flags dbus.Flags, ch chan *dbus.Call, uid uint32, cookie string, identity Identity) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AuthenticationAgentResponse2", flags, ch, uid, cookie, identity)
}

func (v *authority) AuthenticationAgentResponse2(flags dbus.Flags, uid uint32, cookie string, identity Identity) error {
	return (<-v.GoAuthenticationAgentResponse2(flags, make(chan *dbus.Call, 1), uid, cookie, identity).Done).Err
}

// method EnumerateTemporaryAuthorizations

func (v *authority) GoEnumerateTemporaryAuthorizations(flags dbus.Flags, ch chan *dbus.Call, subject Subject) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnumerateTemporaryAuthorizations", flags, ch, subject)
}

func (*authority) StoreEnumerateTemporaryAuthorizations(call *dbus.Call) (temporary_authorizations TemporaryAuthorization, err error) {
	err = call.Store(&temporary_authorizations)
	return
}

func (v *authority) EnumerateTemporaryAuthorizations(flags dbus.Flags, subject Subject) (temporary_authorizations TemporaryAuthorization, err error) {
	return v.StoreEnumerateTemporaryAuthorizations(
		<-v.GoEnumerateTemporaryAuthorizations(flags, make(chan *dbus.Call, 1), subject).Done)
}

// method RevokeTemporaryAuthorizations

func (v *authority) GoRevokeTemporaryAuthorizations(flags dbus.Flags, ch chan *dbus.Call, subject Subject) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RevokeTemporaryAuthorizations", flags, ch, subject)
}

func (v *authority) RevokeTemporaryAuthorizations(flags dbus.Flags, subject Subject) error {
	return (<-v.GoRevokeTemporaryAuthorizations(flags, make(chan *dbus.Call, 1), subject).Done).Err
}

// method RevokeTemporaryAuthorizationById

func (v *authority) GoRevokeTemporaryAuthorizationById(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RevokeTemporaryAuthorizationById", flags, ch, id)
}

func (v *authority) RevokeTemporaryAuthorizationById(flags dbus.Flags, id string) error {
	return (<-v.GoRevokeTemporaryAuthorizationById(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// signal Changed

func (v *authority) ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Changed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Changed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property BackendName s

func (v *authority) BackendName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BackendName",
	}
}

// property BackendVersion s

func (v *authority) BackendVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BackendVersion",
	}
}

// property BackendFeatures u

func (v *authority) BackendFeatures() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "BackendFeatures",
	}
}
