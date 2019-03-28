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

// method ChangeCurrentWorkspaceBackground

func (v *wm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoChangeCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method SwitchToWorkspace

func (v *wm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToWorkspace", flags, ch, backward)
}

func (v *wm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchToWorkspace(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

// method SwitchApplication

func (v *wm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchApplication", flags, ch, backward)
}

func (v *wm) SwitchApplication(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchApplication(flags, make(chan *dbus.Call, 1), backward).Done).Err
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

// method PerformAction

func (v *wm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PerformAction", flags, ch, type0)
}

func (v *wm) PerformAction(flags dbus.Flags, type0 int32) error {
	return (<-v.GoPerformAction(flags, make(chan *dbus.Call, 1), type0).Done).Err
}

// method PreviewWindow

func (v *wm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviewWindow", flags, ch, xid)
}

func (v *wm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	return (<-v.GoPreviewWindow(flags, make(chan *dbus.Call, 1), xid).Done).Err
}

// method CancelPreviewWindow

func (v *wm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *wm) CancelPreviewWindow(flags dbus.Flags) error {
	return (<-v.GoCancelPreviewWindow(flags, make(chan *dbus.Call, 1)).Done).Err
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

// method SetCurrentWorkspaceBackground

func (v *wm) GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoSetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method GetWorkspaceBackground

func (v *wm) GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWorkspaceBackground", flags, ch, index)
}

func (*wm) StoreGetWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetWorkspaceBackground(flags dbus.Flags, index int32) (result string, err error) {
	return v.StoreGetWorkspaceBackground(
		<-v.GoGetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index).Done)
}

// method SetWorkspaceBackground

func (v *wm) GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWorkspaceBackground", flags, ch, index, uri)
}

func (v *wm) SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error {
	return (<-v.GoSetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index, uri).Done).Err
}

// method GetCurrentWorkspace

func (v *wm) GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspace", flags, ch)
}

func (*wm) StoreGetCurrentWorkspace(call *dbus.Call) (index int32, err error) {
	err = call.Store(&index)
	return
}

func (v *wm) GetCurrentWorkspace(flags dbus.Flags) (index int32, err error) {
	return v.StoreGetCurrentWorkspace(
		<-v.GoGetCurrentWorkspace(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetCurrentWorkspace

func (v *wm) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspace", flags, ch, index)
}

func (v *wm) SetCurrentWorkspace(flags dbus.Flags, index int32) error {
	return (<-v.GoSetCurrentWorkspace(flags, make(chan *dbus.Call, 1), index).Done).Err
}

// method PreviousWorkspace

func (v *wm) GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviousWorkspace", flags, ch)
}

func (v *wm) PreviousWorkspace(flags dbus.Flags) error {
	return (<-v.GoPreviousWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method NextWorkspace

func (v *wm) GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NextWorkspace", flags, ch)
}

func (v *wm) NextWorkspace(flags dbus.Flags) error {
	return (<-v.GoNextWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetAllAccels

func (v *wm) GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllAccels", flags, ch)
}

func (*wm) StoreGetAllAccels(call *dbus.Call) (data string, err error) {
	err = call.Store(&data)
	return
}

func (v *wm) GetAllAccels(flags dbus.Flags) (data string, err error) {
	return v.StoreGetAllAccels(
		<-v.GoGetAllAccels(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAccel

func (v *wm) GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccel", flags, ch, id)
}

func (*wm) StoreGetAccel(call *dbus.Call) (data []string, err error) {
	err = call.Store(&data)
	return
}

func (v *wm) GetAccel(flags dbus.Flags, id string) (data []string, err error) {
	return v.StoreGetAccel(
		<-v.GoGetAccel(flags, make(chan *dbus.Call, 1), id).Done)
}

// method SetAccel

func (v *wm) GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAccel", flags, ch, data)
}

func (*wm) StoreSetAccel(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) SetAccel(flags dbus.Flags, data string) (result bool, err error) {
	return v.StoreSetAccel(
		<-v.GoSetAccel(flags, make(chan *dbus.Call, 1), data).Done)
}

// method RemoveAccel

func (v *wm) GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAccel", flags, ch, id)
}

func (v *wm) RemoveAccel(flags dbus.Flags, id string) error {
	return (<-v.GoRemoveAccel(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method PresentWindows

func (v *wm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresentWindows", flags, ch, xids)
}

func (v *wm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	return (<-v.GoPresentWindows(flags, make(chan *dbus.Call, 1), xids).Done).Err
}

// method EnableZoneDetected

func (v *wm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, val bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableZoneDetected", flags, ch, val)
}

func (v *wm) EnableZoneDetected(flags dbus.Flags, val bool) error {
	return (<-v.GoEnableZoneDetected(flags, make(chan *dbus.Call, 1), val).Done).Err
}

// signal WorkspaceBackgroundChanged

func (v *wm) ConnectWorkspaceBackgroundChanged(cb func(index int32, newUri string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceBackgroundChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceBackgroundChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var index int32
		var newUri string
		err := dbus.Store(sig.Body, &index, &newUri)
		if err == nil {
			cb(index, newUri)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StartupReady

func (v *wm) ConnectStartupReady(cb func(wm_name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StartupReady", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StartupReady",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var wm_name string
		err := dbus.Store(sig.Body, &wm_name)
		if err == nil {
			cb(wm_name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
