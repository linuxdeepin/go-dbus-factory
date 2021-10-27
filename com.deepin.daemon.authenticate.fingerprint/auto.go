package fingerprint

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

type Device struct {
	device // interface com.deepin.daemon.Authenticate.Fingerprint.Device
	proxy.Object
}

func NewDevice(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (*Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Device)
	obj.Object.Init_(conn, serviceName, path)
	return obj, nil
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (v *device) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *device) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Claim

func (v *device) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, userId, claimed)
}

func (v *device) Claim(flags dbus.Flags, userId string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), userId, claimed).Done).Err
}

// method DeleteAllFingers

func (v *device) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, userId)
}

func (v *device) DeleteAllFingers(flags dbus.Flags, userId string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), userId).Done).Err
}

// method DeleteFinger

func (v *device) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, userId, finger)
}

func (v *device) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), userId, finger).Done).Err
}

// method Enroll

func (v *device) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, finger)
}

func (v *device) Enroll(flags dbus.Flags, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// method ListFingers

func (v *device) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, userId)
}

func (*device) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *device) ListFingers(flags dbus.Flags, userId string) (fingers []string, err error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), userId).Done)
}

// method RenameFinger

func (v *device) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RenameFinger", flags, ch, userId, finger, newName)
}

func (v *device) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	return (<-v.GoRenameFinger(flags, make(chan *dbus.Call, 1), userId, finger, newName).Done).Err
}

// method StopEnroll

func (v *device) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *device) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *device) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *device) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Verify

func (v *device) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *device) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// signal EnrollStatus

func (v *device) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *device) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *device) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Touch", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Touch",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var pressed bool
		err := dbus.Store(sig.Body, &userId, &pressed)
		if err == nil {
			cb(userId, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Name s

func (v *device) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property State i

func (v *device) State() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "State",
	}
}

// property Type i

func (v *device) Type() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Type",
	}
}

// property Capability i

func (v *device) Capability() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Capability",
	}
}

type CommonDevice struct {
	commonDevice // interface com.deepin.daemon.Authenticate.Fingerprint.CommonDevice
	proxy.Object
}

func NewCommonDevice(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (*CommonDevice, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(CommonDevice)
	obj.Object.Init_(conn, serviceName, path)
	return obj, nil
}

type commonDevice struct{}

func (v *commonDevice) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (v *commonDevice) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *commonDevice) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Claim

func (v *commonDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, userId, claimed)
}

func (v *commonDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), userId, claimed).Done).Err
}

// method DeleteAllFingers

func (v *commonDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, userId)
}

func (v *commonDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), userId).Done).Err
}

// method DeleteFinger

func (v *commonDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, userId, finger)
}

func (v *commonDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), userId, finger).Done).Err
}

// method Enroll

func (v *commonDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, username, finger)
}

func (v *commonDevice) Enroll(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

// method ListFingers

func (v *commonDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, userId)
}

func (*commonDevice) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *commonDevice) ListFingers(flags dbus.Flags, userId string) (fingers []string, err error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), userId).Done)
}

// method RenameFinger

func (v *commonDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RenameFinger", flags, ch, userId, finger, newName)
}

func (v *commonDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	return (<-v.GoRenameFinger(flags, make(chan *dbus.Call, 1), userId, finger, newName).Done).Err
}

// method StopEnroll

func (v *commonDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *commonDevice) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *commonDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *commonDevice) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Verify

func (v *commonDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *commonDevice) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// signal EnrollStatus

func (v *commonDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *commonDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *commonDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Touch", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Touch",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var pressed bool
		err := dbus.Store(sig.Body, &userId, &pressed)
		if err == nil {
			cb(userId, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Name s

func (v *commonDevice) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property State i

func (v *commonDevice) State() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "State",
	}
}

// property Type i

func (v *commonDevice) Type() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Type",
	}
}

// property Capability i

func (v *commonDevice) Capability() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Capability",
	}
}
