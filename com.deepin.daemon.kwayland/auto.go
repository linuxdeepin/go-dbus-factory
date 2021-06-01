package kwayland

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

type OutputManagement struct {
	outputManagement // interface com.deepin.daemon.KWayland.Output
	proxy.Object
}

func NewOutputManagement(conn *dbus.Conn) *OutputManagement {
	obj := new(OutputManagement)
	obj.Object.Init_(conn, "com.deepin.daemon.KWayland", "/com/deepin/daemon/KWayland/Output")
	return obj
}

type outputManagement struct{}

func (v *outputManagement) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*outputManagement) GetInterfaceName_() string {
	return "com.deepin.daemon.KWayland.Output"
}

// method ListOutput

func (v *outputManagement) GoListOutput(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutput", flags, ch)
}

func (v *outputManagement) GoListOutputWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListOutput", flags, ch)
}

func (*outputManagement) StoreListOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) ListOutput(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreListOutput(
		<-v.GoListOutput(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *outputManagement) ListOutputWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListOutputWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListOutput(call)
}

// method GetOutput

func (v *outputManagement) GoGetOutput(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetOutput", flags, ch, arg1)
}

func (v *outputManagement) GoGetOutputWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetOutput", flags, ch, arg1)
}

func (*outputManagement) StoreGetOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) GetOutput(flags dbus.Flags, arg1 string) (arg0 string, err error) {
	return v.StoreGetOutput(
		<-v.GoGetOutput(flags, make(chan *dbus.Call, 1), arg1).Done)
}

func (v *outputManagement) GetOutputWithTimeout(timeout time.Duration, flags dbus.Flags, arg1 string) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetOutputWithContext(ctx, flags, make(chan *dbus.Call, 1), arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetOutput(call)
}

// method Apply

func (v *outputManagement) GoApply(flags dbus.Flags, ch chan *dbus.Call, outputs string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Apply", flags, ch, outputs)
}

func (v *outputManagement) GoApplyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, outputs string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Apply", flags, ch, outputs)
}

func (v *outputManagement) Apply(flags dbus.Flags, outputs string) error {
	return (<-v.GoApply(flags, make(chan *dbus.Call, 1), outputs).Done).Err
}

func (v *outputManagement) ApplyWithTimeout(timeout time.Duration, flags dbus.Flags, outputs string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoApplyWithContext(ctx, flags, make(chan *dbus.Call, 1), outputs).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method WlSimulateKey

func (v *outputManagement) GoWlSimulateKey(flags dbus.Flags, ch chan *dbus.Call, state int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WlSimulateKey", flags, ch, state)
}

func (v *outputManagement) GoWlSimulateKeyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, state int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".WlSimulateKey", flags, ch, state)
}

func (v *outputManagement) WlSimulateKey(flags dbus.Flags, state int32) error {
	return (<-v.GoWlSimulateKey(flags, make(chan *dbus.Call, 1), state).Done).Err
}

func (v *outputManagement) WlSimulateKeyWithTimeout(timeout time.Duration, flags dbus.Flags, state int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoWlSimulateKeyWithContext(ctx, flags, make(chan *dbus.Call, 1), state).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal OutputAdded

func (v *outputManagement) ConnectOutputAdded(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal OutputRemoved

func (v *outputManagement) ConnectOutputRemoved(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal OutputChanged

func (v *outputManagement) ConnectOutputChanged(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type WindowManager struct {
	windowManager // interface com.deepin.daemon.KWayland.WindowManager
	proxy.Object
}

func NewWindowManager(conn *dbus.Conn) *WindowManager {
	obj := new(WindowManager)
	obj.Object.Init_(conn, "com.deepin.daemon.KWayland", "/com/deepin/daemon/KWayland/WindowManager")
	return obj
}

type windowManager struct{}

func (v *windowManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*windowManager) GetInterfaceName_() string {
	return "com.deepin.daemon.KWayland.WindowManager"
}

// method ActiveWindow

func (v *windowManager) GoActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActiveWindow", flags, ch)
}

func (v *windowManager) GoActiveWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ActiveWindow", flags, ch)
}

func (*windowManager) StoreActiveWindow(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *windowManager) ActiveWindow(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreActiveWindow(
		<-v.GoActiveWindow(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *windowManager) ActiveWindowWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoActiveWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreActiveWindow(call)
}

// method IsShowingDesktop

func (v *windowManager) GoIsShowingDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsShowingDesktop", flags, ch)
}

func (v *windowManager) GoIsShowingDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsShowingDesktop", flags, ch)
}

func (*windowManager) StoreIsShowingDesktop(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *windowManager) IsShowingDesktop(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsShowingDesktop(
		<-v.GoIsShowingDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *windowManager) IsShowingDesktopWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsShowingDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsShowingDesktop(call)
}

// method Windows

func (v *windowManager) GoWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Windows", flags, ch)
}

func (v *windowManager) GoWindowsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Windows", flags, ch)
}

func (*windowManager) StoreWindows(call *dbus.Call) (argout0 []dbus.Variant, err error) {
	err = call.Store(&argout0)
	return
}

func (v *windowManager) Windows(flags dbus.Flags) (argout0 []dbus.Variant, err error) {
	return v.StoreWindows(
		<-v.GoWindows(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *windowManager) WindowsWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 []dbus.Variant, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoWindowsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreWindows(call)
}

// signal WindowCreated

func (v *windowManager) ConnectWindowCreated(cb func(ObjPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WindowCreated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WindowCreated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ObjPath string
		err := dbus.Store(sig.Body, &ObjPath)
		if err == nil {
			cb(ObjPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WindowRemove

func (v *windowManager) ConnectWindowRemove(cb func(ObjPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WindowRemove", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WindowRemove",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ObjPath string
		err := dbus.Store(sig.Body, &ObjPath)
		if err == nil {
			cb(ObjPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActiveWindowChanged

func (v *windowManager) ConnectActiveWindowChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActiveWindowChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActiveWindowChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Window struct {
	window // interface com.deepin.daemon.KWayland.PlasmaWindow
	proxy.Object
}

func NewWindow(conn *dbus.Conn, path dbus.ObjectPath) (*Window, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Window)
	obj.Object.Init_(conn, "com.deepin.daemon.KWayland", path)
	return obj, nil
}

type window struct{}

func (v *window) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*window) GetInterfaceName_() string {
	return "com.deepin.daemon.KWayland.PlasmaWindow"
}

// method AppId

func (v *window) GoAppId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AppId", flags, ch)
}

func (v *window) GoAppIdWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AppId", flags, ch)
}

func (*window) StoreAppId(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) AppId(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreAppId(
		<-v.GoAppId(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) AppIdWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAppIdWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAppId(call)
}

// method Geometry

func (v *window) GoGeometry(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Geometry", flags, ch)
}

func (v *window) GoGeometryWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Geometry", flags, ch)
}

func (*window) StoreGeometry(call *dbus.Call) (argout0 Rect, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Geometry(flags dbus.Flags) (argout0 Rect, err error) {
	return v.StoreGeometry(
		<-v.GoGeometry(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) GeometryWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 Rect, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGeometryWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGeometry(call)
}

// method Icon

func (v *window) GoIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Icon", flags, ch)
}

func (v *window) GoIconWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Icon", flags, ch)
}

func (*window) StoreIcon(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Icon(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreIcon(
		<-v.GoIcon(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IconWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIconWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIcon(call)
}

// method InternalId

func (v *window) GoInternalId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InternalId", flags, ch)
}

func (v *window) GoInternalIdWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".InternalId", flags, ch)
}

func (*window) StoreInternalId(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) InternalId(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreInternalId(
		<-v.GoInternalId(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) InternalIdWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoInternalIdWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreInternalId(call)
}

// method IsActive

func (v *window) GoIsActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsActive", flags, ch)
}

func (v *window) GoIsActiveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsActive", flags, ch)
}

func (*window) StoreIsActive(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsActive(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsActive(
		<-v.GoIsActive(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsActiveWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsActiveWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsActive(call)
}

// method IsCloseable

func (v *window) GoIsCloseable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsCloseable", flags, ch)
}

func (v *window) GoIsCloseableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsCloseable", flags, ch)
}

func (*window) StoreIsCloseable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsCloseable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsCloseable(
		<-v.GoIsCloseable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsCloseableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsCloseableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsCloseable(call)
}

// method IsDemandingAttention

func (v *window) GoIsDemandingAttention(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDemandingAttention", flags, ch)
}

func (v *window) GoIsDemandingAttentionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsDemandingAttention", flags, ch)
}

func (*window) StoreIsDemandingAttention(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsDemandingAttention(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsDemandingAttention(
		<-v.GoIsDemandingAttention(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsDemandingAttentionWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsDemandingAttentionWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsDemandingAttention(call)
}

// method IsFullscreen

func (v *window) GoIsFullscreen(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsFullscreen", flags, ch)
}

func (v *window) GoIsFullscreenWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsFullscreen", flags, ch)
}

func (*window) StoreIsFullscreen(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsFullscreen(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsFullscreen(
		<-v.GoIsFullscreen(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsFullscreenWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsFullscreenWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsFullscreen(call)
}

// method IsFullscreenable

func (v *window) GoIsFullscreenable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsFullscreenable", flags, ch)
}

func (v *window) GoIsFullscreenableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsFullscreenable", flags, ch)
}

func (*window) StoreIsFullscreenable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsFullscreenable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsFullscreenable(
		<-v.GoIsFullscreenable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsFullscreenableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsFullscreenableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsFullscreenable(call)
}

// method IsKeepAbove

func (v *window) GoIsKeepAbove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsKeepAbove", flags, ch)
}

func (v *window) GoIsKeepAboveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsKeepAbove", flags, ch)
}

func (*window) StoreIsKeepAbove(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsKeepAbove(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsKeepAbove(
		<-v.GoIsKeepAbove(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsKeepAboveWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsKeepAboveWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsKeepAbove(call)
}

// method IsMaximizeable

func (v *window) GoIsMaximizeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMaximizeable", flags, ch)
}

func (v *window) GoIsMaximizeableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMaximizeable", flags, ch)
}

func (*window) StoreIsMaximizeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMaximizeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMaximizeable(
		<-v.GoIsMaximizeable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsMaximizeableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMaximizeableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMaximizeable(call)
}

// method IsMaximized

func (v *window) GoIsMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMaximized", flags, ch)
}

func (v *window) GoIsMaximizedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMaximized", flags, ch)
}

func (*window) StoreIsMaximized(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMaximized(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMaximized(
		<-v.GoIsMaximized(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsMaximizedWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMaximizedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMaximized(call)
}

// method IsMinimizeable

func (v *window) GoIsMinimizeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMinimizeable", flags, ch)
}

func (v *window) GoIsMinimizeableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMinimizeable", flags, ch)
}

func (*window) StoreIsMinimizeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMinimizeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMinimizeable(
		<-v.GoIsMinimizeable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsMinimizeableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMinimizeableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMinimizeable(call)
}

// method IsMinimized

func (v *window) GoIsMinimized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMinimized", flags, ch)
}

func (v *window) GoIsMinimizedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMinimized", flags, ch)
}

func (*window) StoreIsMinimized(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMinimized(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMinimized(
		<-v.GoIsMinimized(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsMinimizedWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMinimizedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMinimized(call)
}

// method IsMovable

func (v *window) GoIsMovable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMovable", flags, ch)
}

func (v *window) GoIsMovableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMovable", flags, ch)
}

func (*window) StoreIsMovable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMovable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMovable(
		<-v.GoIsMovable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsMovableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMovableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMovable(call)
}

// method IsOnAllDesktops

func (v *window) GoIsOnAllDesktops(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsOnAllDesktops", flags, ch)
}

func (v *window) GoIsOnAllDesktopsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsOnAllDesktops", flags, ch)
}

func (*window) StoreIsOnAllDesktops(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsOnAllDesktops(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsOnAllDesktops(
		<-v.GoIsOnAllDesktops(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsOnAllDesktopsWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsOnAllDesktopsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsOnAllDesktops(call)
}

// method IsResizable

func (v *window) GoIsResizable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsResizable", flags, ch)
}

func (v *window) GoIsResizableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsResizable", flags, ch)
}

func (*window) StoreIsResizable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsResizable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsResizable(
		<-v.GoIsResizable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsResizableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsResizableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsResizable(call)
}

// method IsShadeable

func (v *window) GoIsShadeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsShadeable", flags, ch)
}

func (v *window) GoIsShadeableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsShadeable", flags, ch)
}

func (*window) StoreIsShadeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsShadeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsShadeable(
		<-v.GoIsShadeable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsShadeableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsShadeableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsShadeable(call)
}

// method IsShaded

func (v *window) GoIsShaded(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsShaded", flags, ch)
}

func (v *window) GoIsShadedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsShaded", flags, ch)
}

func (*window) StoreIsShaded(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsShaded(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsShaded(
		<-v.GoIsShaded(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsShadedWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsShadedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsShaded(call)
}

// method IsValid

func (v *window) GoIsValid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsValid", flags, ch)
}

func (v *window) GoIsValidWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsValid", flags, ch)
}

func (*window) StoreIsValid(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsValid(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsValid(
		<-v.GoIsValid(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsValidWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsValidWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsValid(call)
}

// method IsVirtualDesktopChangeable

func (v *window) GoIsVirtualDesktopChangeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsVirtualDesktopChangeable", flags, ch)
}

func (v *window) GoIsVirtualDesktopChangeableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsVirtualDesktopChangeable", flags, ch)
}

func (*window) StoreIsVirtualDesktopChangeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsVirtualDesktopChangeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsVirtualDesktopChangeable(
		<-v.GoIsVirtualDesktopChangeable(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) IsVirtualDesktopChangeableWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsVirtualDesktopChangeableWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsVirtualDesktopChangeable(call)
}

// method Pid

func (v *window) GoPid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pid", flags, ch)
}

func (v *window) GoPidWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Pid", flags, ch)
}

func (*window) StorePid(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Pid(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StorePid(
		<-v.GoPid(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) PidWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPidWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePid(call)
}

// method RequestActivate

func (v *window) GoRequestActivate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestActivate", flags, ch)
}

func (v *window) GoRequestActivateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestActivate", flags, ch)
}

func (v *window) RequestActivate(flags dbus.Flags) error {
	return (<-v.GoRequestActivate(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestActivateWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestActivateWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestClose

func (v *window) GoRequestClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestClose", flags, ch)
}

func (v *window) GoRequestCloseWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestClose", flags, ch)
}

func (v *window) RequestClose(flags dbus.Flags) error {
	return (<-v.GoRequestClose(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestCloseWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestCloseWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestEnterNewVirtualDesktop

func (v *window) GoRequestEnterNewVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestEnterNewVirtualDesktop", flags, ch)
}

func (v *window) GoRequestEnterNewVirtualDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestEnterNewVirtualDesktop", flags, ch)
}

func (v *window) RequestEnterNewVirtualDesktop(flags dbus.Flags) error {
	return (<-v.GoRequestEnterNewVirtualDesktop(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestEnterNewVirtualDesktopWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestEnterNewVirtualDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestEnterVirtualDesktop

func (v *window) GoRequestEnterVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestEnterVirtualDesktop", flags, ch, argin0)
}

func (v *window) GoRequestEnterVirtualDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestEnterVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestEnterVirtualDesktop(flags dbus.Flags, argin0 string) error {
	return (<-v.GoRequestEnterVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

func (v *window) RequestEnterVirtualDesktopWithTimeout(timeout time.Duration, flags dbus.Flags, argin0 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestEnterVirtualDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1), argin0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestLeaveVirtualDesktop

func (v *window) GoRequestLeaveVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLeaveVirtualDesktop", flags, ch, argin0)
}

func (v *window) GoRequestLeaveVirtualDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestLeaveVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestLeaveVirtualDesktop(flags dbus.Flags, argin0 string) error {
	return (<-v.GoRequestLeaveVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

func (v *window) RequestLeaveVirtualDesktopWithTimeout(timeout time.Duration, flags dbus.Flags, argin0 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestLeaveVirtualDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1), argin0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestMove

func (v *window) GoRequestMove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestMove", flags, ch)
}

func (v *window) GoRequestMoveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestMove", flags, ch)
}

func (v *window) RequestMove(flags dbus.Flags) error {
	return (<-v.GoRequestMove(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestMoveWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestMoveWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestResize

func (v *window) GoRequestResize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestResize", flags, ch)
}

func (v *window) GoRequestResizeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestResize", flags, ch)
}

func (v *window) RequestResize(flags dbus.Flags) error {
	return (<-v.GoRequestResize(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestResizeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestResizeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestToggleKeepAbove

func (v *window) GoRequestToggleKeepAbove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleKeepAbove", flags, ch)
}

func (v *window) GoRequestToggleKeepAboveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestToggleKeepAbove", flags, ch)
}

func (v *window) RequestToggleKeepAbove(flags dbus.Flags) error {
	return (<-v.GoRequestToggleKeepAbove(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestToggleKeepAboveWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestToggleKeepAboveWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestToggleKeepBelow

func (v *window) GoRequestToggleKeepBelow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleKeepBelow", flags, ch)
}

func (v *window) GoRequestToggleKeepBelowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestToggleKeepBelow", flags, ch)
}

func (v *window) RequestToggleKeepBelow(flags dbus.Flags) error {
	return (<-v.GoRequestToggleKeepBelow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestToggleKeepBelowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestToggleKeepBelowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestToggleMaximized

func (v *window) GoRequestToggleMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleMaximized", flags, ch)
}

func (v *window) GoRequestToggleMaximizedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestToggleMaximized", flags, ch)
}

func (v *window) RequestToggleMaximized(flags dbus.Flags) error {
	return (<-v.GoRequestToggleMaximized(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestToggleMaximizedWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestToggleMaximizedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestToggleMinimized

func (v *window) GoRequestToggleMinimized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleMinimized", flags, ch)
}

func (v *window) GoRequestToggleMinimizedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestToggleMinimized", flags, ch)
}

func (v *window) RequestToggleMinimized(flags dbus.Flags) error {
	return (<-v.GoRequestToggleMinimized(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestToggleMinimizedWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestToggleMinimizedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestToggleShaded

func (v *window) GoRequestToggleShaded(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleShaded", flags, ch)
}

func (v *window) GoRequestToggleShadedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestToggleShaded", flags, ch)
}

func (v *window) RequestToggleShaded(flags dbus.Flags) error {
	return (<-v.GoRequestToggleShaded(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *window) RequestToggleShadedWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestToggleShadedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestVirtualDesktop

func (v *window) GoRequestVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestVirtualDesktop", flags, ch, argin0)
}

func (v *window) GoRequestVirtualDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, argin0 uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestVirtualDesktop(flags dbus.Flags, argin0 uint32) error {
	return (<-v.GoRequestVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

func (v *window) RequestVirtualDesktopWithTimeout(timeout time.Duration, flags dbus.Flags, argin0 uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestVirtualDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1), argin0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SkipSwitcher

func (v *window) GoSkipSwitcher(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SkipSwitcher", flags, ch)
}

func (v *window) GoSkipSwitcherWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SkipSwitcher", flags, ch)
}

func (*window) StoreSkipSwitcher(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) SkipSwitcher(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreSkipSwitcher(
		<-v.GoSkipSwitcher(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) SkipSwitcherWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSkipSwitcherWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSkipSwitcher(call)
}

// method SkipTaskbar

func (v *window) GoSkipTaskbar(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SkipTaskbar", flags, ch)
}

func (v *window) GoSkipTaskbarWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SkipTaskbar", flags, ch)
}

func (*window) StoreSkipTaskbar(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) SkipTaskbar(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreSkipTaskbar(
		<-v.GoSkipTaskbar(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) SkipTaskbarWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSkipTaskbarWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSkipTaskbar(call)
}

// method Title

func (v *window) GoTitle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Title", flags, ch)
}

func (v *window) GoTitleWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Title", flags, ch)
}

func (*window) StoreTitle(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Title(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreTitle(
		<-v.GoTitle(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) TitleWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTitleWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreTitle(call)
}

// method VirtualDesktop

func (v *window) GoVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VirtualDesktop", flags, ch)
}

func (v *window) GoVirtualDesktopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".VirtualDesktop", flags, ch)
}

func (*window) StoreVirtualDesktop(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) VirtualDesktop(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreVirtualDesktop(
		<-v.GoVirtualDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) VirtualDesktopWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoVirtualDesktopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreVirtualDesktop(call)
}

// method WindowId

func (v *window) GoWindowId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WindowId", flags, ch)
}

func (v *window) GoWindowIdWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".WindowId", flags, ch)
}

func (*window) StoreWindowId(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) WindowId(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreWindowId(
		<-v.GoWindowId(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *window) WindowIdWithTimeout(timeout time.Duration, flags dbus.Flags) (argout0 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoWindowIdWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreWindowId(call)
}

// signal ActiveChanged

func (v *window) ConnectActiveChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActiveChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActiveChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AppIdChanged

func (v *window) ConnectAppIdChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppIdChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppIdChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CloseableChanged

func (v *window) ConnectCloseableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CloseableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CloseableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DemandsAttentionChanged

func (v *window) ConnectDemandsAttentionChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DemandsAttentionChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DemandsAttentionChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal FullscreenableChanged

func (v *window) ConnectFullscreenableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "FullscreenableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".FullscreenableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal FullscreenChanged

func (v *window) ConnectFullscreenChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "FullscreenChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".FullscreenChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal GeometryChanged

func (v *window) ConnectGeometryChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "GeometryChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".GeometryChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal IconChanged

func (v *window) ConnectIconChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IconChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IconChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal KeepAboveChanged

func (v *window) ConnectKeepAboveChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeepAboveChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeepAboveChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal KeepBelowChanged

func (v *window) ConnectKeepBelowChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeepBelowChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeepBelowChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MaximizeableChanged

func (v *window) ConnectMaximizeableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MaximizeableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MaximizeableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MaximizedChanged

func (v *window) ConnectMaximizedChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MaximizedChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MaximizedChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MinimizeableChanged

func (v *window) ConnectMinimizeableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MinimizeableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MinimizeableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MinimizedChanged

func (v *window) ConnectMinimizedChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MinimizedChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MinimizedChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MovableChanged

func (v *window) ConnectMovableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MovableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MovableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal OnAllDesktopsChanged

func (v *window) ConnectOnAllDesktopsChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OnAllDesktopsChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OnAllDesktopsChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ParentWindowChanged

func (v *window) ConnectParentWindowChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ParentWindowChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ParentWindowChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ResizableChanged

func (v *window) ConnectResizableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ResizableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ResizableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ShadeableChanged

func (v *window) ConnectShadeableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ShadeableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ShadeableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ShadedChanged

func (v *window) ConnectShadedChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ShadedChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ShadedChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SkipSwitcherChanged

func (v *window) ConnectSkipSwitcherChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SkipSwitcherChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SkipSwitcherChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SkipTaskbarChanged

func (v *window) ConnectSkipTaskbarChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SkipTaskbarChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SkipTaskbarChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TitleChanged

func (v *window) ConnectTitleChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TitleChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TitleChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Unmapped

func (v *window) ConnectUnmapped(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Unmapped", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Unmapped",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VirtualDesktopChangeableChanged

func (v *window) ConnectVirtualDesktopChangeableChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VirtualDesktopChangeableChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VirtualDesktopChangeableChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VirtualDesktopChanged

func (v *window) ConnectVirtualDesktopChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VirtualDesktopChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VirtualDesktopChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WindowIdChanged

func (v *window) ConnectWindowIdChanged(cb func(wid uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WindowIdChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WindowIdChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var wid uint32
		err := dbus.Store(sig.Body, &wid)
		if err == nil {
			cb(wid)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
