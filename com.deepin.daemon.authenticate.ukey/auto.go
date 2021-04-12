package ukey

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

type UKey struct {
	ukey // interface com.deepin.daemon.Authenticate.UKey.Device
	proxy.Object
}

func NewUKey(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (*UKey, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(UKey)
	obj.Object.Init_(conn, serviceName, path)
	return obj, nil
}

type ukey struct{}

func (v *ukey) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (v *ukey) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *ukey) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method SetPin

func (v *ukey) GoSetPin(flags dbus.Flags, ch chan *dbus.Call, username string, id string, pin string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPin", flags, ch, username, id, pin)
}

func (v *ukey) SetPin(flags dbus.Flags, username string, id string, pin string) error {
	return (<-v.GoSetPin(flags, make(chan *dbus.Call, 1), username, id, pin).Done).Err
}

// method SetSessionPath

func (v *ukey) GoSetSessionPath(flags dbus.Flags, ch chan *dbus.Call, username string, id string, path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSessionPath", flags, ch, username, id, path)
}

func (v *ukey) SetSessionPath(flags dbus.Flags, username string, id string, path string) error {
	return (<-v.GoSetSessionPath(flags, make(chan *dbus.Call, 1), username, id, path).Done).Err
}

// method StopVerify

func (v *ukey) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, username string, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch, username, id)
}

func (v *ukey) StopVerify(flags dbus.Flags, username string, id string) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1), username, id).Done).Err
}

// method Verify

func (v *ukey) GoVerify(flags dbus.Flags, ch chan *dbus.Call, username string, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, username, id)
}

func (v *ukey) Verify(flags dbus.Flags, username string, id string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), username, id).Done).Err
}

// signal VerifyResult

func (v *ukey) ConnectVerifyResult(cb func(id string, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyResult", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyResult",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var msg string
		err := dbus.Store(sig.Body, &id, &msg)
		if err == nil {
			cb(id, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property State i

func (v *ukey) State() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "State",
	}
}

// property Type i

func (v *ukey) Type() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Type",
	}
}

// property Capability i

func (v *ukey) Capability() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Capability",
	}
}

// property Name s

func (v *ukey) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}
