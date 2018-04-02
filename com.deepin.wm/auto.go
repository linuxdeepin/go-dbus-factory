package wm

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

type Wm struct {
	wm // interface com.deepin.wm
	proxy.Object
}

func NewWm(conn *dbus.Conn) *Wm {
	obj := new(Wm)
	obj.Object.Init_(conn, "com.deepin.wm", "/com/deepin/wm")
	return obj
}

type wm struct{}

func (v *wm) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wm) GetInterfaceName_() string {
	return "com.deepin.wm"
}

// method PerformAction

func (v *wm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PerformAction", flags, ch, type0)
}

func (v *wm) PerformAction(flags dbus.Flags, type0 int32) error {
	return (<-v.GoPerformAction(flags, make(chan *dbus.Call, 1), type0).Done).Err
}

// method ToggleDebug

func (v *wm) GoToggleDebug(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleDebug", flags, ch)
}

func (v *wm) ToggleDebug(flags dbus.Flags) error {
	return (<-v.GoToggleDebug(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method CancelPreviewWindow

func (v *wm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *wm) CancelPreviewWindow(flags dbus.Flags) error {
	return (<-v.GoCancelPreviewWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PreviewWindow

func (v *wm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviewWindow", flags, ch, xid)
}

func (v *wm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	return (<-v.GoPreviewWindow(flags, make(chan *dbus.Call, 1), xid).Done).Err
}

// method PresentWindows

func (v *wm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresentWindows", flags, ch, xids)
}

func (v *wm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	return (<-v.GoPresentWindows(flags, make(chan *dbus.Call, 1), xids).Done).Err
}

// method RequestHideWindows

func (v *wm) GoRequestHideWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestHideWindows", flags, ch)
}

func (v *wm) RequestHideWindows(flags dbus.Flags) error {
	return (<-v.GoRequestHideWindows(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method CancelHideWindows

func (v *wm) GoCancelHideWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelHideWindows", flags, ch)
}

func (v *wm) CancelHideWindows(flags dbus.Flags) error {
	return (<-v.GoCancelHideWindows(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ChangeCurrentWorkspaceBackground

func (v *wm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoChangeCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method SetTransientBackground

func (v *wm) GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTransientBackground", flags, ch, uri)
}

func (v *wm) SetTransientBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoSetTransientBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method GetCurrentWorkspaceBackground

func (v *wm) GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspaceBackground", flags, ch)
}

func (*wm) StoreGetCurrentWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetCurrentWorkspaceBackground(flags dbus.Flags) (result string, err error) {
	return v.StoreGetCurrentWorkspaceBackground(
		<-v.GoGetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1)).Done)
}

// method EnableZoneDetected

func (v *wm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, val bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableZoneDetected", flags, ch, val)
}

func (v *wm) EnableZoneDetected(flags dbus.Flags, val bool) error {
	return (<-v.GoEnableZoneDetected(flags, make(chan *dbus.Call, 1), val).Done).Err
}

// method SwitchApplication

func (v *wm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchApplication", flags, ch, backward)
}

func (v *wm) SwitchApplication(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchApplication(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

// method SwitchToWorkspace

func (v *wm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToWorkspace", flags, ch, backward)
}

func (v *wm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchToWorkspace(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

// method TileActiveWindow

func (v *wm) GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TileActiveWindow", flags, ch, side)
}

func (v *wm) TileActiveWindow(flags dbus.Flags, side uint32) error {
	return (<-v.GoTileActiveWindow(flags, make(chan *dbus.Call, 1), side).Done).Err
}

// method BeginToMoveActiveWindow

func (v *wm) GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BeginToMoveActiveWindow", flags, ch)
}

func (v *wm) BeginToMoveActiveWindow(flags dbus.Flags) error {
	return (<-v.GoBeginToMoveActiveWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal WorkspaceRemoved

func (v *wm) ConnectWorkspaceRemoved(cb func(index int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var index int32
		err := dbus.Store(sig.Body, &index)
		if err == nil {
			cb(index)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WorkspaceAdded

func (v *wm) ConnectWorkspaceAdded(cb func(index int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var index int32
		err := dbus.Store(sig.Body, &index)
		if err == nil {
			cb(index)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WorkspaceSwitched

func (v *wm) ConnectWorkspaceSwitched(cb func(from int32, to int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceSwitched", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceSwitched",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var from int32
		var to int32
		err := dbus.Store(sig.Body, &from, &to)
		if err == nil {
			cb(from, to)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
