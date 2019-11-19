package fingerprint

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

type Fingerprint struct {
	fingerprint // interface com.huawei.Fingerprint
	proxy.Object
}

func NewFingerprint(conn *dbus.Conn) *Fingerprint {
	obj := new(Fingerprint)
	obj.Object.Init_(conn, "com.huawei.Fingerprint", "/com/huawei/Fingerprint")
	return obj
}

type fingerprint struct{}

func (v *fingerprint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*fingerprint) GetInterfaceName_() string {
	return "com.huawei.Fingerprint"
}

// method SearchDevice

func (v *fingerprint) GoSearchDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchDevice", flags, ch)
}

func (*fingerprint) StoreSearchDevice(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) SearchDevice(flags dbus.Flags) (result bool, err error) {
	return v.StoreSearchDevice(
		<-v.GoSearchDevice(flags, make(chan *dbus.Call, 1)).Done)
}

// method Identify

func (v *fingerprint) GoIdentify(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Identify", flags, ch, uuid)
}

func (v *fingerprint) Identify(flags dbus.Flags, uuid string) error {
	return (<-v.GoIdentify(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

// method GetStatus

func (v *fingerprint) GoGetStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStatus", flags, ch)
}

func (*fingerprint) StoreGetStatus(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) GetStatus(flags dbus.Flags) (result int32, err error) {
	return v.StoreGetStatus(
		<-v.GoGetStatus(flags, make(chan *dbus.Call, 1)).Done)
}

// method Enroll

func (v *fingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, filePath string, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, filePath, uuid)
}

func (v *fingerprint) Enroll(flags dbus.Flags, filePath string, uuid string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), filePath, uuid).Done).Err
}

// method Close

func (v *fingerprint) GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Close", flags, ch)
}

func (*fingerprint) StoreClose(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) Close(flags dbus.Flags) (result int32, err error) {
	return v.StoreClose(
		<-v.GoClose(flags, make(chan *dbus.Call, 1)).Done)
}

// method Reload

func (v *fingerprint) GoReload(flags dbus.Flags, ch chan *dbus.Call, deleteType int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, deleteType)
}

func (*fingerprint) StoreReload(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) Reload(flags dbus.Flags, deleteType int32) (result int32, err error) {
	return v.StoreReload(
		<-v.GoReload(flags, make(chan *dbus.Call, 1), deleteType).Done)
}

// signal EnrollStatus

func (v *fingerprint) ConnectEnrollStatus(cb func(progress int32, result int32)) (dbusutil.SignalHandlerId, error) {
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
		var progress int32
		var result int32
		err := dbus.Store(sig.Body, &progress, &result)
		if err == nil {
			cb(progress, result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal IdentifyStatus

func (v *fingerprint) ConnectIdentifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdentifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdentifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result int32
		err := dbus.Store(sig.Body, &result)
		if err == nil {
			cb(result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *fingerprint) ConnectVerifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error) {
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
		var result int32
		err := dbus.Store(sig.Body, &result)
		if err == nil {
			cb(result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
