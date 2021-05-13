package faceverify

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

type FaceVerify struct {
	faceVerify // interface com.deepin.daemon.FaceVerify
	proxy.Object
}

func NewFaceVerify(conn *dbus.Conn) *FaceVerify {
	obj := new(FaceVerify)
	obj.Object.Init_(conn, "com.deepin.daemon.FaceVerify", "/com/deepin/daemon/FaceVerify")
	return obj
}

type faceVerify struct{}

func (v *faceVerify) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*faceVerify) GetInterfaceName_() string {
	return "com.deepin.daemon.FaceVerify"
}

// method Enroll

func (v *faceVerify) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, name)
}

func (v *faceVerify) Enroll(flags dbus.Flags, name string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method Verify

func (v *faceVerify) GoVerify(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, name)
}

func (v *faceVerify) Verify(flags dbus.Flags, name string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method StopVerify

func (v *faceVerify) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *faceVerify) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopEnroll

func (v *faceVerify) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *faceVerify) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DeleteFace

func (v *faceVerify) GoDeleteFace(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFace", flags, ch, name)
}

func (v *faceVerify) DeleteFace(flags dbus.Flags, name string) error {
	return (<-v.GoDeleteFace(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method GetFaceStatus

func (v *faceVerify) GoGetFaceStatus(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetFaceStatus", flags, ch, name)
}

func (*faceVerify) StoreGetFaceStatus(call *dbus.Call) (isExist bool, err error) {
	err = call.Store(&isExist)
	return
}

func (v *faceVerify) GetFaceStatus(flags dbus.Flags, name string) (isExist bool, err error) {
	return v.StoreGetFaceStatus(
		<-v.GoGetFaceStatus(flags, make(chan *dbus.Call, 1), name).Done)
}

// signal EnrollStatus

func (v *faceVerify) ConnectEnrollStatus(cb func(name string, result int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var name string
		var result int32
		var msg string
		err := dbus.Store(sig.Body, &name, &result, &msg)
		if err == nil {
			cb(name, result, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *faceVerify) ConnectVerifyStatus(cb func(name string, result int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var name string
		var result int32
		var msg string
		err := dbus.Store(sig.Body, &name, &result, &msg)
		if err == nil {
			cb(name, result, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
