package fingerprint

import "context"
import "errors"
import "fmt"
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

func (v *fingerprint) GoSearchDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SearchDevice", flags, ch)
}

func (*fingerprint) StoreSearchDevice(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) SearchDevice(flags dbus.Flags) (result bool, err error) {
	return v.StoreSearchDevice(
		<-v.GoSearchDevice(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fingerprint) SearchDeviceWithTimeout(timeout time.Duration, flags dbus.Flags) (result bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSearchDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSearchDevice(call)
}

// method Identify

func (v *fingerprint) GoIdentify(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Identify", flags, ch, uuid)
}

func (v *fingerprint) GoIdentifyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Identify", flags, ch, uuid)
}

func (v *fingerprint) Identify(flags dbus.Flags, uuid string) error {
	return (<-v.GoIdentify(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

func (v *fingerprint) IdentifyWithTimeout(timeout time.Duration, flags dbus.Flags, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIdentifyWithContext(ctx, flags, make(chan *dbus.Call, 1), uuid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method IdentifyWithMultipleUser

func (v *fingerprint) GoIdentifyWithMultipleUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IdentifyWithMultipleUser", flags, ch)
}

func (v *fingerprint) GoIdentifyWithMultipleUserWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IdentifyWithMultipleUser", flags, ch)
}

func (v *fingerprint) IdentifyWithMultipleUser(flags dbus.Flags) error {
	return (<-v.GoIdentifyWithMultipleUser(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *fingerprint) IdentifyWithMultipleUserWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIdentifyWithMultipleUserWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetStatus

func (v *fingerprint) GoGetStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStatus", flags, ch)
}

func (v *fingerprint) GoGetStatusWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetStatus", flags, ch)
}

func (*fingerprint) StoreGetStatus(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) GetStatus(flags dbus.Flags) (result int32, err error) {
	return v.StoreGetStatus(
		<-v.GoGetStatus(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fingerprint) GetStatusWithTimeout(timeout time.Duration, flags dbus.Flags) (result int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetStatusWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetStatus(call)
}

// method Enroll

func (v *fingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, filePath string, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, filePath, uuid)
}

func (v *fingerprint) GoEnrollWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, filePath string, uuid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Enroll", flags, ch, filePath, uuid)
}

func (v *fingerprint) Enroll(flags dbus.Flags, filePath string, uuid string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), filePath, uuid).Done).Err
}

func (v *fingerprint) EnrollWithTimeout(timeout time.Duration, flags dbus.Flags, filePath string, uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnrollWithContext(ctx, flags, make(chan *dbus.Call, 1), filePath, uuid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Close

func (v *fingerprint) GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Close", flags, ch)
}

func (v *fingerprint) GoCloseWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Close", flags, ch)
}

func (*fingerprint) StoreClose(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) Close(flags dbus.Flags) (result int32, err error) {
	return v.StoreClose(
		<-v.GoClose(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fingerprint) CloseWithTimeout(timeout time.Duration, flags dbus.Flags) (result int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCloseWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreClose(call)
}

// method Reload

func (v *fingerprint) GoReload(flags dbus.Flags, ch chan *dbus.Call, deleteType int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, deleteType)
}

func (v *fingerprint) GoReloadWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, deleteType int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reload", flags, ch, deleteType)
}

func (*fingerprint) StoreReload(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) Reload(flags dbus.Flags, deleteType int32) (result int32, err error) {
	return v.StoreReload(
		<-v.GoReload(flags, make(chan *dbus.Call, 1), deleteType).Done)
}

func (v *fingerprint) ReloadWithTimeout(timeout time.Duration, flags dbus.Flags, deleteType int32) (result int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadWithContext(ctx, flags, make(chan *dbus.Call, 1), deleteType).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReload(call)
}

// method ClearPovImage

func (v *fingerprint) GoClearPovImage(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearPovImage", flags, ch)
}

func (v *fingerprint) GoClearPovImageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ClearPovImage", flags, ch)
}

func (*fingerprint) StoreClearPovImage(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *fingerprint) ClearPovImage(flags dbus.Flags) (result int32, err error) {
	return v.StoreClearPovImage(
		<-v.GoClearPovImage(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fingerprint) ClearPovImageWithTimeout(timeout time.Duration, flags dbus.Flags) (result int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClearPovImageWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreClearPovImage(call)
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

// signal IdentifyNoAccount

func (v *fingerprint) ConnectIdentifyNoAccount(cb func(result int32, userName string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdentifyNoAccount", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdentifyNoAccount",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result int32
		var userName string
		err := dbus.Store(sig.Body, &result, &userName)
		if err == nil {
			cb(result, userName)
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

// signal DeviceStatus

func (v *fingerprint) ConnectDeviceStatus(cb func(result bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result bool
		err := dbus.Store(sig.Body, &result)
		if err == nil {
			cb(result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
