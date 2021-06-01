package dock

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

type Dock struct {
	dock // interface com.deepin.dde.daemon.Dock
	proxy.Object
}

func NewDock(conn *dbus.Conn) *Dock {
	obj := new(Dock)
	obj.Object.Init_(conn, "com.deepin.dde.daemon.Dock", "/com/deepin/dde/daemon/Dock")
	return obj
}

type dock struct{}

func (v *dock) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dock) GetInterfaceName_() string {
	return "com.deepin.dde.daemon.Dock"
}

// method ActivateWindow

func (v *dock) GoActivateWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateWindow", flags, ch, win)
}

func (v *dock) GoActivateWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ActivateWindow", flags, ch, win)
}

func (v *dock) ActivateWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoActivateWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) ActivateWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoActivateWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method CancelPreviewWindow

func (v *dock) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *dock) GoCancelPreviewWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *dock) CancelPreviewWindow(flags dbus.Flags) error {
	return (<-v.GoCancelPreviewWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *dock) CancelPreviewWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCancelPreviewWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method CloseWindow

func (v *dock) GoCloseWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseWindow", flags, ch, win)
}

func (v *dock) GoCloseWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CloseWindow", flags, ch, win)
}

func (v *dock) CloseWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoCloseWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) CloseWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCloseWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetEntryIDs

func (v *dock) GoGetEntryIDs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetEntryIDs", flags, ch)
}

func (v *dock) GoGetEntryIDsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetEntryIDs", flags, ch)
}

func (*dock) StoreGetEntryIDs(call *dbus.Call) (list []string, err error) {
	err = call.Store(&list)
	return
}

func (v *dock) GetEntryIDs(flags dbus.Flags) (list []string, err error) {
	return v.StoreGetEntryIDs(
		<-v.GoGetEntryIDs(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *dock) GetEntryIDsWithTimeout(timeout time.Duration, flags dbus.Flags) (list []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetEntryIDsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetEntryIDs(call)
}

// method IsDocked

func (v *dock) GoIsDocked(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsDocked", flags, ch, desktopFile)
}

func (v *dock) GoIsDockedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsDocked", flags, ch, desktopFile)
}

func (*dock) StoreIsDocked(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *dock) IsDocked(flags dbus.Flags, desktopFile string) (value bool, err error) {
	return v.StoreIsDocked(
		<-v.GoIsDocked(flags, make(chan *dbus.Call, 1), desktopFile).Done)
}

func (v *dock) IsDockedWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFile string) (value bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsDockedWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFile).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsDocked(call)
}

// method IsOnDock

func (v *dock) GoIsOnDock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsOnDock", flags, ch, desktopFile)
}

func (v *dock) GoIsOnDockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsOnDock", flags, ch, desktopFile)
}

func (*dock) StoreIsOnDock(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *dock) IsOnDock(flags dbus.Flags, desktopFile string) (value bool, err error) {
	return v.StoreIsOnDock(
		<-v.GoIsOnDock(flags, make(chan *dbus.Call, 1), desktopFile).Done)
}

func (v *dock) IsOnDockWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFile string) (value bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsOnDockWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFile).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsOnDock(call)
}

// method MakeWindowAbove

func (v *dock) GoMakeWindowAbove(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MakeWindowAbove", flags, ch, win)
}

func (v *dock) GoMakeWindowAboveWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MakeWindowAbove", flags, ch, win)
}

func (v *dock) MakeWindowAbove(flags dbus.Flags, win uint32) error {
	return (<-v.GoMakeWindowAbove(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) MakeWindowAboveWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMakeWindowAboveWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method MaximizeWindow

func (v *dock) GoMaximizeWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MaximizeWindow", flags, ch, win)
}

func (v *dock) GoMaximizeWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MaximizeWindow", flags, ch, win)
}

func (v *dock) MaximizeWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoMaximizeWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) MaximizeWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMaximizeWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method MinimizeWindow

func (v *dock) GoMinimizeWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MinimizeWindow", flags, ch, win)
}

func (v *dock) GoMinimizeWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MinimizeWindow", flags, ch, win)
}

func (v *dock) MinimizeWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoMinimizeWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) MinimizeWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMinimizeWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method MoveEntry

func (v *dock) GoMoveEntry(flags dbus.Flags, ch chan *dbus.Call, index int32, newIndex int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MoveEntry", flags, ch, index, newIndex)
}

func (v *dock) GoMoveEntryWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32, newIndex int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MoveEntry", flags, ch, index, newIndex)
}

func (v *dock) MoveEntry(flags dbus.Flags, index int32, newIndex int32) error {
	return (<-v.GoMoveEntry(flags, make(chan *dbus.Call, 1), index, newIndex).Done).Err
}

func (v *dock) MoveEntryWithTimeout(timeout time.Duration, flags dbus.Flags, index int32, newIndex int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMoveEntryWithContext(ctx, flags, make(chan *dbus.Call, 1), index, newIndex).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method MoveWindow

func (v *dock) GoMoveWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MoveWindow", flags, ch, win)
}

func (v *dock) GoMoveWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MoveWindow", flags, ch, win)
}

func (v *dock) MoveWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoMoveWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) MoveWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMoveWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PreviewWindow

func (v *dock) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviewWindow", flags, ch, win)
}

func (v *dock) GoPreviewWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PreviewWindow", flags, ch, win)
}

func (v *dock) PreviewWindow(flags dbus.Flags, win uint32) error {
	return (<-v.GoPreviewWindow(flags, make(chan *dbus.Call, 1), win).Done).Err
}

func (v *dock) PreviewWindowWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPreviewWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method QueryWindowIdentifyMethod

func (v *dock) GoQueryWindowIdentifyMethod(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".QueryWindowIdentifyMethod", flags, ch, win)
}

func (v *dock) GoQueryWindowIdentifyMethodWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".QueryWindowIdentifyMethod", flags, ch, win)
}

func (*dock) StoreQueryWindowIdentifyMethod(call *dbus.Call) (identifyMethod string, err error) {
	err = call.Store(&identifyMethod)
	return
}

func (v *dock) QueryWindowIdentifyMethod(flags dbus.Flags, win uint32) (identifyMethod string, err error) {
	return v.StoreQueryWindowIdentifyMethod(
		<-v.GoQueryWindowIdentifyMethod(flags, make(chan *dbus.Call, 1), win).Done)
}

func (v *dock) QueryWindowIdentifyMethodWithTimeout(timeout time.Duration, flags dbus.Flags, win uint32) (identifyMethod string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoQueryWindowIdentifyMethodWithContext(ctx, flags, make(chan *dbus.Call, 1), win).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreQueryWindowIdentifyMethod(call)
}

// method RequestDock

func (v *dock) GoRequestDock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestDock", flags, ch, desktopFile, index)
}

func (v *dock) GoRequestDockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFile string, index int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestDock", flags, ch, desktopFile, index)
}

func (*dock) StoreRequestDock(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *dock) RequestDock(flags dbus.Flags, desktopFile string, index int32) (ok bool, err error) {
	return v.StoreRequestDock(
		<-v.GoRequestDock(flags, make(chan *dbus.Call, 1), desktopFile, index).Done)
}

func (v *dock) RequestDockWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFile string, index int32) (ok bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestDockWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFile, index).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRequestDock(call)
}

// method RequestUndock

func (v *dock) GoRequestUndock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestUndock", flags, ch, desktopFile)
}

func (v *dock) GoRequestUndockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestUndock", flags, ch, desktopFile)
}

func (*dock) StoreRequestUndock(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *dock) RequestUndock(flags dbus.Flags, desktopFile string) (ok bool, err error) {
	return v.StoreRequestUndock(
		<-v.GoRequestUndock(flags, make(chan *dbus.Call, 1), desktopFile).Done)
}

func (v *dock) RequestUndockWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFile string) (ok bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestUndockWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFile).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRequestUndock(call)
}

// method SetFrontendWindowRect

func (v *dock) GoSetFrontendWindowRect(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32, width uint32, height uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFrontendWindowRect", flags, ch, x, y, width, height)
}

func (v *dock) GoSetFrontendWindowRectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, x int32, y int32, width uint32, height uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetFrontendWindowRect", flags, ch, x, y, width, height)
}

func (v *dock) SetFrontendWindowRect(flags dbus.Flags, x int32, y int32, width uint32, height uint32) error {
	return (<-v.GoSetFrontendWindowRect(flags, make(chan *dbus.Call, 1), x, y, width, height).Done).Err
}

func (v *dock) SetFrontendWindowRectWithTimeout(timeout time.Duration, flags dbus.Flags, x int32, y int32, width uint32, height uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetFrontendWindowRectWithContext(ctx, flags, make(chan *dbus.Call, 1), x, y, width, height).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal ServiceRestarted

func (v *dock) ConnectServiceRestarted(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ServiceRestarted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ServiceRestarted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal EntryAdded

func (v *dock) ConnectEntryAdded(cb func(path dbus.ObjectPath, index int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EntryAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EntryAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		var index int32
		err := dbus.Store(sig.Body, &path, &index)
		if err == nil {
			cb(path, index)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal EntryRemoved

func (v *dock) ConnectEntryRemoved(cb func(entryId string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EntryRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EntryRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var entryId string
		err := dbus.Store(sig.Body, &entryId)
		if err == nil {
			cb(entryId)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ShowTimeout u

func (v *dock) ShowTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ShowTimeout",
	}
}

// property HideTimeout u

func (v *dock) HideTimeout() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "HideTimeout",
	}
}

// property FrontendWindowRect (iiuu)

func (v *dock) FrontendWindowRect() PropDockFrontendWindowRect {
	return PropDockFrontendWindowRect{
		Impl: v,
	}
}

type PropDockFrontendWindowRect struct {
	Impl proxy.Implementer
}

func (p PropDockFrontendWindowRect) Get(flags dbus.Flags) (value FrontendWindowRect, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"FrontendWindowRect", &value)
	return
}

func (p PropDockFrontendWindowRect) ConnectChanged(cb func(hasValue bool, value FrontendWindowRect)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v FrontendWindowRect
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, FrontendWindowRect{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"FrontendWindowRect", cb0)
}

// property Entries ao

func (v *dock) Entries() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Entries",
	}
}

// property HideMode i

func (v *dock) HideMode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "HideMode",
	}
}

// property DisplayMode i

func (v *dock) DisplayMode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DisplayMode",
	}
}

// property HideState i

func (v *dock) HideState() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "HideState",
	}
}

// property Position i

func (v *dock) Position() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Position",
	}
}

// property IconSize u

func (v *dock) IconSize() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "IconSize",
	}
}

// property DockedApps as

func (v *dock) DockedApps() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DockedApps",
	}
}

type Entry struct {
	entry // interface com.deepin.dde.daemon.Dock.Entry
	proxy.Object
}

func NewEntry(conn *dbus.Conn, path dbus.ObjectPath) (*Entry, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Entry)
	obj.Object.Init_(conn, "com.deepin.dde.daemon.Dock", path)
	return obj, nil
}

type entry struct{}

func (v *entry) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*entry) GetInterfaceName_() string {
	return "com.deepin.dde.daemon.Dock.Entry"
}

// method Activate

func (v *entry) GoActivate(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Activate", flags, ch, timestamp)
}

func (v *entry) GoActivateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Activate", flags, ch, timestamp)
}

func (v *entry) Activate(flags dbus.Flags, timestamp uint32) error {
	return (<-v.GoActivate(flags, make(chan *dbus.Call, 1), timestamp).Done).Err
}

func (v *entry) ActivateWithTimeout(timeout time.Duration, flags dbus.Flags, timestamp uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoActivateWithContext(ctx, flags, make(chan *dbus.Call, 1), timestamp).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Check

func (v *entry) GoCheck(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Check", flags, ch)
}

func (v *entry) GoCheckWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Check", flags, ch)
}

func (v *entry) Check(flags dbus.Flags) error {
	return (<-v.GoCheck(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *entry) CheckWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCheckWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ForceQuit

func (v *entry) GoForceQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceQuit", flags, ch)
}

func (v *entry) GoForceQuitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ForceQuit", flags, ch)
}

func (v *entry) ForceQuit(flags dbus.Flags) error {
	return (<-v.GoForceQuit(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *entry) ForceQuitWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoForceQuitWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method HandleDragDrop

func (v *entry) GoHandleDragDrop(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, files []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".HandleDragDrop", flags, ch, timestamp, files)
}

func (v *entry) GoHandleDragDropWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, files []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".HandleDragDrop", flags, ch, timestamp, files)
}

func (v *entry) HandleDragDrop(flags dbus.Flags, timestamp uint32, files []string) error {
	return (<-v.GoHandleDragDrop(flags, make(chan *dbus.Call, 1), timestamp, files).Done).Err
}

func (v *entry) HandleDragDropWithTimeout(timeout time.Duration, flags dbus.Flags, timestamp uint32, files []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHandleDragDropWithContext(ctx, flags, make(chan *dbus.Call, 1), timestamp, files).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method HandleMenuItem

func (v *entry) GoHandleMenuItem(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".HandleMenuItem", flags, ch, timestamp, id)
}

func (v *entry) GoHandleMenuItemWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".HandleMenuItem", flags, ch, timestamp, id)
}

func (v *entry) HandleMenuItem(flags dbus.Flags, timestamp uint32, id string) error {
	return (<-v.GoHandleMenuItem(flags, make(chan *dbus.Call, 1), timestamp, id).Done).Err
}

func (v *entry) HandleMenuItemWithTimeout(timeout time.Duration, flags dbus.Flags, timestamp uint32, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHandleMenuItemWithContext(ctx, flags, make(chan *dbus.Call, 1), timestamp, id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method NewInstance

func (v *entry) GoNewInstance(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NewInstance", flags, ch, timestamp)
}

func (v *entry) GoNewInstanceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".NewInstance", flags, ch, timestamp)
}

func (v *entry) NewInstance(flags dbus.Flags, timestamp uint32) error {
	return (<-v.GoNewInstance(flags, make(chan *dbus.Call, 1), timestamp).Done).Err
}

func (v *entry) NewInstanceWithTimeout(timeout time.Duration, flags dbus.Flags, timestamp uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoNewInstanceWithContext(ctx, flags, make(chan *dbus.Call, 1), timestamp).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PresentWindows

func (v *entry) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresentWindows", flags, ch)
}

func (v *entry) GoPresentWindowsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PresentWindows", flags, ch)
}

func (v *entry) PresentWindows(flags dbus.Flags) error {
	return (<-v.GoPresentWindows(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *entry) PresentWindowsWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPresentWindowsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestDock

func (v *entry) GoRequestDock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestDock", flags, ch)
}

func (v *entry) GoRequestDockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestDock", flags, ch)
}

func (v *entry) RequestDock(flags dbus.Flags) error {
	return (<-v.GoRequestDock(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *entry) RequestDockWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestDockWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestUndock

func (v *entry) GoRequestUndock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestUndock", flags, ch)
}

func (v *entry) GoRequestUndockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestUndock", flags, ch)
}

func (v *entry) RequestUndock(flags dbus.Flags) error {
	return (<-v.GoRequestUndock(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *entry) RequestUndockWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestUndockWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Name s

func (v *entry) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Icon s

func (v *entry) Icon() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Icon",
	}
}

// property Id s

func (v *entry) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property IsActive b

func (v *entry) IsActive() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsActive",
	}
}

// property CurrentWindow u

func (v *entry) CurrentWindow() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "CurrentWindow",
	}
}

// property IsDocked b

func (v *entry) IsDocked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsDocked",
	}
}

// property WindowInfos a{u(sb)}

func (v *entry) WindowInfos() PropEntryWindowInfos {
	return PropEntryWindowInfos{
		Impl: v,
	}
}

type PropEntryWindowInfos struct {
	Impl proxy.Implementer
}

func (p PropEntryWindowInfos) Get(flags dbus.Flags) (value map[uint32]WindowInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"WindowInfos", &value)
	return
}

func (p PropEntryWindowInfos) ConnectChanged(cb func(hasValue bool, value map[uint32]WindowInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[uint32]WindowInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"WindowInfos", cb0)
}

// property Menu s

func (v *entry) Menu() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Menu",
	}
}

// property DesktopFile s

func (v *entry) DesktopFile() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DesktopFile",
	}
}
