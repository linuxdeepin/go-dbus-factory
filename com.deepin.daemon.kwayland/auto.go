package kwayland

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

func (*outputManagement) StoreListOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) ListOutput(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreListOutput(
		<-v.GoListOutput(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetOutput

func (v *outputManagement) GoGetOutput(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetOutput", flags, ch, arg1)
}

func (*outputManagement) StoreGetOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) GetOutput(flags dbus.Flags, arg1 string) (arg0 string, err error) {
	return v.StoreGetOutput(
		<-v.GoGetOutput(flags, make(chan *dbus.Call, 1), arg1).Done)
}

// method Apply

func (v *outputManagement) GoApply(flags dbus.Flags, ch chan *dbus.Call, outputs string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Apply", flags, ch, outputs)
}

func (v *outputManagement) Apply(flags dbus.Flags, outputs string) error {
	return (<-v.GoApply(flags, make(chan *dbus.Call, 1), outputs).Done).Err
}

// method WlSimulateKey

func (v *outputManagement) GoWlSimulateKey(flags dbus.Flags, ch chan *dbus.Call, state int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WlSimulateKey", flags, ch, state)
}

func (v *outputManagement) WlSimulateKey(flags dbus.Flags, state int32) error {
	return (<-v.GoWlSimulateKey(flags, make(chan *dbus.Call, 1), state).Done).Err
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

func (*windowManager) StoreActiveWindow(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *windowManager) ActiveWindow(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreActiveWindow(
		<-v.GoActiveWindow(flags, make(chan *dbus.Call, 1)).Done)
}

// method Windows

func (v *windowManager) GoWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Windows", flags, ch)
}

func (*windowManager) StoreWindows(call *dbus.Call) (argout0 []dbus.Variant, err error) {
	err = call.Store(&argout0)
	return
}

func (v *windowManager) Windows(flags dbus.Flags) (argout0 []dbus.Variant, err error) {
	return v.StoreWindows(
		<-v.GoWindows(flags, make(chan *dbus.Call, 1)).Done)
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

func (*window) StoreAppId(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) AppId(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreAppId(
		<-v.GoAppId(flags, make(chan *dbus.Call, 1)).Done)
}

// method Geometry

func (v *window) GoGeometry(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Geometry", flags, ch)
}

func (*window) StoreGeometry(call *dbus.Call) (argout0 Rect, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Geometry(flags dbus.Flags) (argout0 Rect, err error) {
	return v.StoreGeometry(
		<-v.GoGeometry(flags, make(chan *dbus.Call, 1)).Done)
}

// method Icon

func (v *window) GoIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Icon", flags, ch)
}

func (*window) StoreIcon(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Icon(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreIcon(
		<-v.GoIcon(flags, make(chan *dbus.Call, 1)).Done)
}

// method InternalId

func (v *window) GoInternalId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InternalId", flags, ch)
}

func (*window) StoreInternalId(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) InternalId(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreInternalId(
		<-v.GoInternalId(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsActive

func (v *window) GoIsActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsActive", flags, ch)
}

func (*window) StoreIsActive(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsActive(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsActive(
		<-v.GoIsActive(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsCloseable

func (v *window) GoIsCloseable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsCloseable", flags, ch)
}

func (*window) StoreIsCloseable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsCloseable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsCloseable(
		<-v.GoIsCloseable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsDemandingAttention

func (v *window) GoIsDemandingAttention(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDemandingAttention", flags, ch)
}

func (*window) StoreIsDemandingAttention(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsDemandingAttention(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsDemandingAttention(
		<-v.GoIsDemandingAttention(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsFullscreen

func (v *window) GoIsFullscreen(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsFullscreen", flags, ch)
}

func (*window) StoreIsFullscreen(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsFullscreen(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsFullscreen(
		<-v.GoIsFullscreen(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsFullscreenable

func (v *window) GoIsFullscreenable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsFullscreenable", flags, ch)
}

func (*window) StoreIsFullscreenable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsFullscreenable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsFullscreenable(
		<-v.GoIsFullscreenable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsKeepAbove

func (v *window) GoIsKeepAbove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsKeepAbove", flags, ch)
}

func (*window) StoreIsKeepAbove(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsKeepAbove(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsKeepAbove(
		<-v.GoIsKeepAbove(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsMaximizeable

func (v *window) GoIsMaximizeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMaximizeable", flags, ch)
}

func (*window) StoreIsMaximizeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMaximizeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMaximizeable(
		<-v.GoIsMaximizeable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsMaximized

func (v *window) GoIsMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMaximized", flags, ch)
}

func (*window) StoreIsMaximized(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMaximized(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMaximized(
		<-v.GoIsMaximized(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsMinimizeable

func (v *window) GoIsMinimizeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMinimizeable", flags, ch)
}

func (*window) StoreIsMinimizeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMinimizeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMinimizeable(
		<-v.GoIsMinimizeable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsMinimized

func (v *window) GoIsMinimized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMinimized", flags, ch)
}

func (*window) StoreIsMinimized(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMinimized(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMinimized(
		<-v.GoIsMinimized(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsMovable

func (v *window) GoIsMovable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMovable", flags, ch)
}

func (*window) StoreIsMovable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsMovable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsMovable(
		<-v.GoIsMovable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsOnAllDesktops

func (v *window) GoIsOnAllDesktops(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsOnAllDesktops", flags, ch)
}

func (*window) StoreIsOnAllDesktops(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsOnAllDesktops(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsOnAllDesktops(
		<-v.GoIsOnAllDesktops(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsResizable

func (v *window) GoIsResizable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsResizable", flags, ch)
}

func (*window) StoreIsResizable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsResizable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsResizable(
		<-v.GoIsResizable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsShadeable

func (v *window) GoIsShadeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsShadeable", flags, ch)
}

func (*window) StoreIsShadeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsShadeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsShadeable(
		<-v.GoIsShadeable(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsShaded

func (v *window) GoIsShaded(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsShaded", flags, ch)
}

func (*window) StoreIsShaded(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsShaded(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsShaded(
		<-v.GoIsShaded(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsValid

func (v *window) GoIsValid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsValid", flags, ch)
}

func (*window) StoreIsValid(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsValid(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsValid(
		<-v.GoIsValid(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsVirtualDesktopChangeable

func (v *window) GoIsVirtualDesktopChangeable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsVirtualDesktopChangeable", flags, ch)
}

func (*window) StoreIsVirtualDesktopChangeable(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) IsVirtualDesktopChangeable(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreIsVirtualDesktopChangeable(
		<-v.GoIsVirtualDesktopChangeable(flags, make(chan *dbus.Call, 1)).Done)
}

// method Pid

func (v *window) GoPid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pid", flags, ch)
}

func (*window) StorePid(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Pid(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StorePid(
		<-v.GoPid(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestActivate

func (v *window) GoRequestActivate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestActivate", flags, ch)
}

func (v *window) RequestActivate(flags dbus.Flags) error {
	return (<-v.GoRequestActivate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestClose

func (v *window) GoRequestClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestClose", flags, ch)
}

func (v *window) RequestClose(flags dbus.Flags) error {
	return (<-v.GoRequestClose(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestEnterNewVirtualDesktop

func (v *window) GoRequestEnterNewVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestEnterNewVirtualDesktop", flags, ch)
}

func (v *window) RequestEnterNewVirtualDesktop(flags dbus.Flags) error {
	return (<-v.GoRequestEnterNewVirtualDesktop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestEnterVirtualDesktop

func (v *window) GoRequestEnterVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestEnterVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestEnterVirtualDesktop(flags dbus.Flags, argin0 string) error {
	return (<-v.GoRequestEnterVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

// method RequestLeaveVirtualDesktop

func (v *window) GoRequestLeaveVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLeaveVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestLeaveVirtualDesktop(flags dbus.Flags, argin0 string) error {
	return (<-v.GoRequestLeaveVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

// method RequestMove

func (v *window) GoRequestMove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestMove", flags, ch)
}

func (v *window) RequestMove(flags dbus.Flags) error {
	return (<-v.GoRequestMove(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestResize

func (v *window) GoRequestResize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestResize", flags, ch)
}

func (v *window) RequestResize(flags dbus.Flags) error {
	return (<-v.GoRequestResize(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestToggleKeepAbove

func (v *window) GoRequestToggleKeepAbove(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleKeepAbove", flags, ch)
}

func (v *window) RequestToggleKeepAbove(flags dbus.Flags) error {
	return (<-v.GoRequestToggleKeepAbove(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestToggleKeepBelow

func (v *window) GoRequestToggleKeepBelow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleKeepBelow", flags, ch)
}

func (v *window) RequestToggleKeepBelow(flags dbus.Flags) error {
	return (<-v.GoRequestToggleKeepBelow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestToggleMaximized

func (v *window) GoRequestToggleMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleMaximized", flags, ch)
}

func (v *window) RequestToggleMaximized(flags dbus.Flags) error {
	return (<-v.GoRequestToggleMaximized(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestToggleMinimized

func (v *window) GoRequestToggleMinimized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleMinimized", flags, ch)
}

func (v *window) RequestToggleMinimized(flags dbus.Flags) error {
	return (<-v.GoRequestToggleMinimized(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestToggleShaded

func (v *window) GoRequestToggleShaded(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestToggleShaded", flags, ch)
}

func (v *window) RequestToggleShaded(flags dbus.Flags) error {
	return (<-v.GoRequestToggleShaded(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestVirtualDesktop

func (v *window) GoRequestVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call, argin0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestVirtualDesktop", flags, ch, argin0)
}

func (v *window) RequestVirtualDesktop(flags dbus.Flags, argin0 uint32) error {
	return (<-v.GoRequestVirtualDesktop(flags, make(chan *dbus.Call, 1), argin0).Done).Err
}

// method SkipSwitcher

func (v *window) GoSkipSwitcher(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SkipSwitcher", flags, ch)
}

func (*window) StoreSkipSwitcher(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) SkipSwitcher(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreSkipSwitcher(
		<-v.GoSkipSwitcher(flags, make(chan *dbus.Call, 1)).Done)
}

// method SkipTaskbar

func (v *window) GoSkipTaskbar(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SkipTaskbar", flags, ch)
}

func (*window) StoreSkipTaskbar(call *dbus.Call) (argout0 bool, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) SkipTaskbar(flags dbus.Flags) (argout0 bool, err error) {
	return v.StoreSkipTaskbar(
		<-v.GoSkipTaskbar(flags, make(chan *dbus.Call, 1)).Done)
}

// method Title

func (v *window) GoTitle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Title", flags, ch)
}

func (*window) StoreTitle(call *dbus.Call) (argout0 string, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) Title(flags dbus.Flags) (argout0 string, err error) {
	return v.StoreTitle(
		<-v.GoTitle(flags, make(chan *dbus.Call, 1)).Done)
}

// method VirtualDesktop

func (v *window) GoVirtualDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VirtualDesktop", flags, ch)
}

func (*window) StoreVirtualDesktop(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) VirtualDesktop(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreVirtualDesktop(
		<-v.GoVirtualDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

// method WindowId

func (v *window) GoWindowId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WindowId", flags, ch)
}

func (*window) StoreWindowId(call *dbus.Call) (argout0 uint32, err error) {
	err = call.Store(&argout0)
	return
}

func (v *window) WindowId(flags dbus.Flags) (argout0 uint32, err error) {
	return v.StoreWindowId(
		<-v.GoWindowId(flags, make(chan *dbus.Call, 1)).Done)
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
