package authenticate

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

type Authenticate struct {
	authenticate // interface com.deepin.daemon.Authenticate
	proxy.Object
}

func NewAuthenticate(conn *dbus.Conn) *Authenticate {
	obj := new(Authenticate)
	obj.Object.Init_(conn, "com.deepin.daemon.Authenticate", "/com/deepin/daemon/Authenticate")
	return obj
}

type authenticate struct{}

func (v *authenticate) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*authenticate) GetInterfaceName_() string {
	return "com.deepin.daemon.Authenticate"
}

// method Authenticate

func (v *authenticate) GoAuthenticate(flags dbus.Flags, ch chan *dbus.Call, username string, flag int32, timeout int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Authenticate", flags, ch, username, flag, timeout)
}

func (v *authenticate) GoAuthenticateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string, flag int32, timeout int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Authenticate", flags, ch, username, flag, timeout)
}

func (*authenticate) StoreAuthenticate(call *dbus.Call) (id string, err error) {
	err = call.Store(&id)
	return
}

func (v *authenticate) Authenticate(flags dbus.Flags, username string, flag int32, timeout int32) (id string, err error) {
	return v.StoreAuthenticate(
		<-v.GoAuthenticate(flags, make(chan *dbus.Call, 1), username, flag, timeout).Done)
}

func (v *authenticate) AuthenticateWithTimeout(timeout_1 time.Duration, flags dbus.Flags, username string, flag int32, timeout int32) (id string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout_1)
	defer cancel()
	call := <-v.GoAuthenticateWithContext(ctx, flags, make(chan *dbus.Call, 1), username, flag, timeout).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAuthenticate(call)
}

// method CancelAuthenticate

func (v *authenticate) GoCancelAuthenticate(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelAuthenticate", flags, ch, id)
}

func (v *authenticate) GoCancelAuthenticateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CancelAuthenticate", flags, ch, id)
}

func (v *authenticate) CancelAuthenticate(flags dbus.Flags, id string) error {
	return (<-v.GoCancelAuthenticate(flags, make(chan *dbus.Call, 1), id).Done).Err
}

func (v *authenticate) CancelAuthenticateWithTimeout(timeout time.Duration, flags dbus.Flags, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCancelAuthenticateWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetPassword

func (v *authenticate) GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, id string, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPassword", flags, ch, id, password)
}

func (v *authenticate) GoSetPasswordWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string, password string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetPassword", flags, ch, id, password)
}

func (v *authenticate) SetPassword(flags dbus.Flags, id string, password string) error {
	return (<-v.GoSetPassword(flags, make(chan *dbus.Call, 1), id, password).Done).Err
}

func (v *authenticate) SetPasswordWithTimeout(timeout time.Duration, flags dbus.Flags, id string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetPasswordWithContext(ctx, flags, make(chan *dbus.Call, 1), id, password).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal Status

func (v *authenticate) ConnectStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Status", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Status",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &id, &code, &msg)
		if err == nil {
			cb(id, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property SupportedFlags i

func (v *authenticate) SupportedFlags() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SupportedFlags",
	}
}

type Fingerprint struct {
	fingerprint // interface com.deepin.daemon.Authenticate.Fingerprint
	proxy.Object
}

func NewFingerprint(conn *dbus.Conn) *Fingerprint {
	obj := new(Fingerprint)
	obj.Object.Init_(conn, "com.deepin.daemon.Authenticate", "/com/deepin/daemon/Authenticate/Fingerprint")
	return obj
}

type fingerprint struct{}

func (v *fingerprint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*fingerprint) GetInterfaceName_() string {
	return "com.deepin.daemon.Authenticate.Fingerprint"
}

// method Claim

func (v *fingerprint) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, username, claimed)
}

func (v *fingerprint) GoClaimWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string, claimed bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Claim", flags, ch, username, claimed)
}

func (v *fingerprint) Claim(flags dbus.Flags, username string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), username, claimed).Done).Err
}

func (v *fingerprint) ClaimWithTimeout(timeout time.Duration, flags dbus.Flags, username string, claimed bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClaimWithContext(ctx, flags, make(chan *dbus.Call, 1), username, claimed).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteAllFingers

func (v *fingerprint) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, username)
}

func (v *fingerprint) GoDeleteAllFingersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, username)
}

func (v *fingerprint) DeleteAllFingers(flags dbus.Flags, username string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), username).Done).Err
}

func (v *fingerprint) DeleteAllFingersWithTimeout(timeout time.Duration, flags dbus.Flags, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteAllFingersWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteFinger

func (v *fingerprint) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, username, finger)
}

func (v *fingerprint) GoDeleteFingerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteFinger", flags, ch, username, finger)
}

func (v *fingerprint) DeleteFinger(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

func (v *fingerprint) DeleteFingerWithTimeout(timeout time.Duration, flags dbus.Flags, username string, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteFingerWithContext(ctx, flags, make(chan *dbus.Call, 1), username, finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Enroll

func (v *fingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, finger)
}

func (v *fingerprint) GoEnrollWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Enroll", flags, ch, finger)
}

func (v *fingerprint) Enroll(flags dbus.Flags, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

func (v *fingerprint) EnrollWithTimeout(timeout time.Duration, flags dbus.Flags, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnrollWithContext(ctx, flags, make(chan *dbus.Call, 1), finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ListFingers

func (v *fingerprint) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, username)
}

func (v *fingerprint) GoListFingersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListFingers", flags, ch, username)
}

func (*fingerprint) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *fingerprint) ListFingers(flags dbus.Flags, username string) (fingers []string, err error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

func (v *fingerprint) ListFingersWithTimeout(timeout time.Duration, flags dbus.Flags, username string) (fingers []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListFingersWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListFingers(call)
}

// method SetDefaultDevice

func (v *fingerprint) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultDevice", flags, ch, device)
}

func (v *fingerprint) GoSetDefaultDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDefaultDevice", flags, ch, device)
}

func (v *fingerprint) SetDefaultDevice(flags dbus.Flags, device string) error {
	return (<-v.GoSetDefaultDevice(flags, make(chan *dbus.Call, 1), device).Done).Err
}

func (v *fingerprint) SetDefaultDeviceWithTimeout(timeout time.Duration, flags dbus.Flags, device string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDefaultDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1), device).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method StopEnroll

func (v *fingerprint) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *fingerprint) GoStopEnrollWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *fingerprint) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *fingerprint) StopEnrollWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopEnrollWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method StopVerify

func (v *fingerprint) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *fingerprint) GoStopVerifyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *fingerprint) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *fingerprint) StopVerifyWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopVerifyWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Verify

func (v *fingerprint) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *fingerprint) GoVerifyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *fingerprint) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

func (v *fingerprint) VerifyWithTimeout(timeout time.Duration, flags dbus.Flags, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoVerifyWithContext(ctx, flags, make(chan *dbus.Call, 1), finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal EnrollStatus

func (v *fingerprint) ConnectEnrollStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var id string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &id, &code, &msg)
		if err == nil {
			cb(id, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *fingerprint) ConnectVerifyStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var id string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &id, &code, &msg)
		if err == nil {
			cb(id, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal FingerMouldStatus

func (v *fingerprint) ConnectFingerMouldStatus(cb func(code int32, username string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "FingerMouldStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".FingerMouldStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var code int32
		var username string
		err := dbus.Store(sig.Body, &code, &username)
		if err == nil {
			cb(code, username)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *fingerprint) ConnectTouch(cb func(id string, pressed bool)) (dbusutil.SignalHandlerId, error) {
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
		var id string
		var pressed bool
		err := dbus.Store(sig.Body, &id, &pressed)
		if err == nil {
			cb(id, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property DefaultDevice s

func (v *fingerprint) DefaultDevice() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DefaultDevice",
	}
}

// property Devices s

func (v *fingerprint) Devices() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Devices",
	}
}
