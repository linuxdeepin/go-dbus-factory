package wm

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

// method SwitchApplication

func (v *wm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchApplication", flags, ch, backward)
}

func (v *wm) GoSwitchApplicationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SwitchApplication", flags, ch, backward)
}

func (v *wm) SwitchApplication(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchApplication(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

func (v *wm) SwitchApplicationWithTimeout(timeout time.Duration, flags dbus.Flags, backward bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSwitchApplicationWithContext(ctx, flags, make(chan *dbus.Call, 1), backward).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method TileActiveWindow

func (v *wm) GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TileActiveWindow", flags, ch, side)
}

func (v *wm) GoTileActiveWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".TileActiveWindow", flags, ch, side)
}

func (v *wm) TileActiveWindow(flags dbus.Flags, side uint32) error {
	return (<-v.GoTileActiveWindow(flags, make(chan *dbus.Call, 1), side).Done).Err
}

func (v *wm) TileActiveWindowWithTimeout(timeout time.Duration, flags dbus.Flags, side uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTileActiveWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), side).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method BeginToMoveActiveWindow

func (v *wm) GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BeginToMoveActiveWindow", flags, ch)
}

func (v *wm) GoBeginToMoveActiveWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".BeginToMoveActiveWindow", flags, ch)
}

func (v *wm) BeginToMoveActiveWindow(flags dbus.Flags) error {
	return (<-v.GoBeginToMoveActiveWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) BeginToMoveActiveWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoBeginToMoveActiveWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ToggleActiveWindowMaximize

func (v *wm) GoToggleActiveWindowMaximize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleActiveWindowMaximize", flags, ch)
}

func (v *wm) GoToggleActiveWindowMaximizeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ToggleActiveWindowMaximize", flags, ch)
}

func (v *wm) ToggleActiveWindowMaximize(flags dbus.Flags) error {
	return (<-v.GoToggleActiveWindowMaximize(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) ToggleActiveWindowMaximizeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoToggleActiveWindowMaximizeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method MinimizeActiveWindow

func (v *wm) GoMinimizeActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MinimizeActiveWindow", flags, ch)
}

func (v *wm) GoMinimizeActiveWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MinimizeActiveWindow", flags, ch)
}

func (v *wm) MinimizeActiveWindow(flags dbus.Flags) error {
	return (<-v.GoMinimizeActiveWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) MinimizeActiveWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMinimizeActiveWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowWorkspace

func (v *wm) GoShowWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowWorkspace", flags, ch)
}

func (v *wm) GoShowWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowWorkspace", flags, ch)
}

func (v *wm) ShowWorkspace(flags dbus.Flags) error {
	return (<-v.GoShowWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) ShowWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowWindow

func (v *wm) GoShowWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowWindow", flags, ch)
}

func (v *wm) GoShowWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowWindow", flags, ch)
}

func (v *wm) ShowWindow(flags dbus.Flags) error {
	return (<-v.GoShowWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) ShowWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowAllWindow

func (v *wm) GoShowAllWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowAllWindow", flags, ch)
}

func (v *wm) GoShowAllWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowAllWindow", flags, ch)
}

func (v *wm) ShowAllWindow(flags dbus.Flags) error {
	return (<-v.GoShowAllWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) ShowAllWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowAllWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PerformAction

func (v *wm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PerformAction", flags, ch, type0)
}

func (v *wm) GoPerformActionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PerformAction", flags, ch, type0)
}

func (v *wm) PerformAction(flags dbus.Flags, type0 int32) error {
	return (<-v.GoPerformAction(flags, make(chan *dbus.Call, 1), type0).Done).Err
}

func (v *wm) PerformActionWithTimeout(timeout time.Duration, flags dbus.Flags, type0 int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPerformActionWithContext(ctx, flags, make(chan *dbus.Call, 1), type0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PreviewWindow

func (v *wm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviewWindow", flags, ch, xid)
}

func (v *wm) GoPreviewWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PreviewWindow", flags, ch, xid)
}

func (v *wm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	return (<-v.GoPreviewWindow(flags, make(chan *dbus.Call, 1), xid).Done).Err
}

func (v *wm) PreviewWindowWithTimeout(timeout time.Duration, flags dbus.Flags, xid uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPreviewWindowWithContext(ctx, flags, make(chan *dbus.Call, 1), xid).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method CancelPreviewWindow

func (v *wm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *wm) GoCancelPreviewWindowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *wm) CancelPreviewWindow(flags dbus.Flags) error {
	return (<-v.GoCancelPreviewWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) CancelPreviewWindowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCancelPreviewWindowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetCurrentWorkspaceBackground

func (v *wm) GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspaceBackground", flags, ch)
}

func (v *wm) GoGetCurrentWorkspaceBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetCurrentWorkspaceBackground", flags, ch)
}

func (*wm) StoreGetCurrentWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetCurrentWorkspaceBackground(flags dbus.Flags) (result string, err error) {
	return v.StoreGetCurrentWorkspaceBackground(
		<-v.GoGetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *wm) GetCurrentWorkspaceBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetCurrentWorkspaceBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetCurrentWorkspaceBackground(call)
}

// method SetCurrentWorkspaceBackground

func (v *wm) GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) GoSetCurrentWorkspaceBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoSetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

func (v *wm) SetCurrentWorkspaceBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetCurrentWorkspaceBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), uri).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetWorkspaceBackground

func (v *wm) GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWorkspaceBackground", flags, ch, index)
}

func (v *wm) GoGetWorkspaceBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetWorkspaceBackground", flags, ch, index)
}

func (*wm) StoreGetWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetWorkspaceBackground(flags dbus.Flags, index int32) (result string, err error) {
	return v.StoreGetWorkspaceBackground(
		<-v.GoGetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index).Done)
}

func (v *wm) GetWorkspaceBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, index int32) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetWorkspaceBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), index).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetWorkspaceBackground(call)
}

// method SetWorkspaceBackground

func (v *wm) GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWorkspaceBackground", flags, ch, index, uri)
}

func (v *wm) GoSetWorkspaceBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetWorkspaceBackground", flags, ch, index, uri)
}

func (v *wm) SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error {
	return (<-v.GoSetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index, uri).Done).Err
}

func (v *wm) SetWorkspaceBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, index int32, uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetWorkspaceBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), index, uri).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetTransientBackground

func (v *wm) GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTransientBackground", flags, ch, arg0)
}

func (v *wm) GoSetTransientBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTransientBackground", flags, ch, arg0)
}

func (v *wm) SetTransientBackground(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetTransientBackground(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *wm) SetTransientBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTransientBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetCurrentWorkspaceBackgroundForMonitor

func (v *wm) GoGetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspaceBackgroundForMonitor", flags, ch, strMonitorName)
}

func (v *wm) GoGetCurrentWorkspaceBackgroundForMonitorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetCurrentWorkspaceBackgroundForMonitor", flags, ch, strMonitorName)
}

func (*wm) StoreGetCurrentWorkspaceBackgroundForMonitor(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, strMonitorName string) (result string, err error) {
	return v.StoreGetCurrentWorkspaceBackgroundForMonitor(
		<-v.GoGetCurrentWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), strMonitorName).Done)
}

func (v *wm) GetCurrentWorkspaceBackgroundForMonitorWithTimeout(timeout time.Duration, flags dbus.Flags, strMonitorName string) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetCurrentWorkspaceBackgroundForMonitorWithContext(ctx, flags, make(chan *dbus.Call, 1), strMonitorName).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetCurrentWorkspaceBackgroundForMonitor(call)
}

// method SetCurrentWorkspaceBackgroundForMonitor

func (v *wm) GoSetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspaceBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *wm) GoSetCurrentWorkspaceBackgroundForMonitorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetCurrentWorkspaceBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *wm) SetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	return (<-v.GoSetCurrentWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done).Err
}

func (v *wm) SetCurrentWorkspaceBackgroundForMonitorWithTimeout(timeout time.Duration, flags dbus.Flags, uri string, strMonitorName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetCurrentWorkspaceBackgroundForMonitorWithContext(ctx, flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetWorkspaceBackgroundForMonitor

func (v *wm) GoGetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName)
}

func (v *wm) GoGetWorkspaceBackgroundForMonitorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName)
}

func (*wm) StoreGetWorkspaceBackgroundForMonitor(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) GetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string) (result string, err error) {
	return v.StoreGetWorkspaceBackgroundForMonitor(
		<-v.GoGetWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), index, strMonitorName).Done)
}

func (v *wm) GetWorkspaceBackgroundForMonitorWithTimeout(timeout time.Duration, flags dbus.Flags, index int32, strMonitorName string) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetWorkspaceBackgroundForMonitorWithContext(ctx, flags, make(chan *dbus.Call, 1), index, strMonitorName).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetWorkspaceBackgroundForMonitor(call)
}

// method SetWorkspaceBackgroundForMonitor

func (v *wm) GoSetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName, uri)
}

func (v *wm) GoSetWorkspaceBackgroundForMonitorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName, uri)
}

func (v *wm) SetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string, uri string) error {
	return (<-v.GoSetWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), index, strMonitorName, uri).Done).Err
}

func (v *wm) SetWorkspaceBackgroundForMonitorWithTimeout(timeout time.Duration, flags dbus.Flags, index int32, strMonitorName string, uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetWorkspaceBackgroundForMonitorWithContext(ctx, flags, make(chan *dbus.Call, 1), index, strMonitorName, uri).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetTransientBackgroundForMonitor

func (v *wm) GoSetTransientBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTransientBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *wm) GoSetTransientBackgroundForMonitorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTransientBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *wm) SetTransientBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	return (<-v.GoSetTransientBackgroundForMonitor(flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done).Err
}

func (v *wm) SetTransientBackgroundForMonitorWithTimeout(timeout time.Duration, flags dbus.Flags, uri string, strMonitorName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTransientBackgroundForMonitorWithContext(ctx, flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetCurrentWorkspace

func (v *wm) GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspace", flags, ch)
}

func (v *wm) GoGetCurrentWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetCurrentWorkspace", flags, ch)
}

func (*wm) StoreGetCurrentWorkspace(call *dbus.Call) (index int32, err error) {
	err = call.Store(&index)
	return
}

func (v *wm) GetCurrentWorkspace(flags dbus.Flags) (index int32, err error) {
	return v.StoreGetCurrentWorkspace(
		<-v.GoGetCurrentWorkspace(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *wm) GetCurrentWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags) (index int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetCurrentWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetCurrentWorkspace(call)
}

// method WorkspaceCount

func (v *wm) GoWorkspaceCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WorkspaceCount", flags, ch)
}

func (v *wm) GoWorkspaceCountWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".WorkspaceCount", flags, ch)
}

func (*wm) StoreWorkspaceCount(call *dbus.Call) (count int32, err error) {
	err = call.Store(&count)
	return
}

func (v *wm) WorkspaceCount(flags dbus.Flags) (count int32, err error) {
	return v.StoreWorkspaceCount(
		<-v.GoWorkspaceCount(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *wm) WorkspaceCountWithTimeout(timeout time.Duration, flags dbus.Flags) (count int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoWorkspaceCountWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreWorkspaceCount(call)
}

// method SetCurrentWorkspace

func (v *wm) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspace", flags, ch, index)
}

func (v *wm) GoSetCurrentWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetCurrentWorkspace", flags, ch, index)
}

func (v *wm) SetCurrentWorkspace(flags dbus.Flags, index int32) error {
	return (<-v.GoSetCurrentWorkspace(flags, make(chan *dbus.Call, 1), index).Done).Err
}

func (v *wm) SetCurrentWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags, index int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetCurrentWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1), index).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PreviousWorkspace

func (v *wm) GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviousWorkspace", flags, ch)
}

func (v *wm) GoPreviousWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PreviousWorkspace", flags, ch)
}

func (v *wm) PreviousWorkspace(flags dbus.Flags) error {
	return (<-v.GoPreviousWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) PreviousWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPreviousWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method NextWorkspace

func (v *wm) GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NextWorkspace", flags, ch)
}

func (v *wm) GoNextWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".NextWorkspace", flags, ch)
}

func (v *wm) NextWorkspace(flags dbus.Flags) error {
	return (<-v.GoNextWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wm) NextWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoNextWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetAllAccels

func (v *wm) GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllAccels", flags, ch)
}

func (v *wm) GoGetAllAccelsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetAllAccels", flags, ch)
}

func (*wm) StoreGetAllAccels(call *dbus.Call) (data string, err error) {
	err = call.Store(&data)
	return
}

func (v *wm) GetAllAccels(flags dbus.Flags) (data string, err error) {
	return v.StoreGetAllAccels(
		<-v.GoGetAllAccels(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *wm) GetAllAccelsWithTimeout(timeout time.Duration, flags dbus.Flags) (data string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAllAccelsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetAllAccels(call)
}

// method GetAccel

func (v *wm) GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccel", flags, ch, id)
}

func (v *wm) GoGetAccelWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetAccel", flags, ch, id)
}

func (*wm) StoreGetAccel(call *dbus.Call) (data []string, err error) {
	err = call.Store(&data)
	return
}

func (v *wm) GetAccel(flags dbus.Flags, id string) (data []string, err error) {
	return v.StoreGetAccel(
		<-v.GoGetAccel(flags, make(chan *dbus.Call, 1), id).Done)
}

func (v *wm) GetAccelWithTimeout(timeout time.Duration, flags dbus.Flags, id string) (data []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAccelWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetAccel(call)
}

// method GetDefaultAccel

func (v *wm) GoGetDefaultAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultAccel", flags, ch, id)
}

func (v *wm) GoGetDefaultAccelWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDefaultAccel", flags, ch, id)
}

func (*wm) StoreGetDefaultAccel(call *dbus.Call) (data []string, err error) {
	err = call.Store(&data)
	return
}

func (v *wm) GetDefaultAccel(flags dbus.Flags, id string) (data []string, err error) {
	return v.StoreGetDefaultAccel(
		<-v.GoGetDefaultAccel(flags, make(chan *dbus.Call, 1), id).Done)
}

func (v *wm) GetDefaultAccelWithTimeout(timeout time.Duration, flags dbus.Flags, id string) (data []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetDefaultAccelWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetDefaultAccel(call)
}

// method SetAccel

func (v *wm) GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAccel", flags, ch, data)
}

func (v *wm) GoSetAccelWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAccel", flags, ch, data)
}

func (*wm) StoreSetAccel(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *wm) SetAccel(flags dbus.Flags, data string) (result bool, err error) {
	return v.StoreSetAccel(
		<-v.GoSetAccel(flags, make(chan *dbus.Call, 1), data).Done)
}

func (v *wm) SetAccelWithTimeout(timeout time.Duration, flags dbus.Flags, data string) (result bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAccelWithContext(ctx, flags, make(chan *dbus.Call, 1), data).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSetAccel(call)
}

// method RemoveAccel

func (v *wm) GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAccel", flags, ch, id)
}

func (v *wm) GoRemoveAccelWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveAccel", flags, ch, id)
}

func (v *wm) RemoveAccel(flags dbus.Flags, id string) error {
	return (<-v.GoRemoveAccel(flags, make(chan *dbus.Call, 1), id).Done).Err
}

func (v *wm) RemoveAccelWithTimeout(timeout time.Duration, flags dbus.Flags, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveAccelWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetDecorationTheme

func (v *wm) GoSetDecorationTheme(flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDecorationTheme", flags, ch, themeType, themeName)
}

func (v *wm) GoSetDecorationThemeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDecorationTheme", flags, ch, themeType, themeName)
}

func (v *wm) SetDecorationTheme(flags dbus.Flags, themeType string, themeName string) error {
	return (<-v.GoSetDecorationTheme(flags, make(chan *dbus.Call, 1), themeType, themeName).Done).Err
}

func (v *wm) SetDecorationThemeWithTimeout(timeout time.Duration, flags dbus.Flags, themeType string, themeName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDecorationThemeWithContext(ctx, flags, make(chan *dbus.Call, 1), themeType, themeName).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetDecorationDeepinTheme

func (v *wm) GoSetDecorationDeepinTheme(flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDecorationDeepinTheme", flags, ch, deepinThemeName)
}

func (v *wm) GoSetDecorationDeepinThemeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDecorationDeepinTheme", flags, ch, deepinThemeName)
}

func (v *wm) SetDecorationDeepinTheme(flags dbus.Flags, deepinThemeName string) error {
	return (<-v.GoSetDecorationDeepinTheme(flags, make(chan *dbus.Call, 1), deepinThemeName).Done).Err
}

func (v *wm) SetDecorationDeepinThemeWithTimeout(timeout time.Duration, flags dbus.Flags, deepinThemeName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDecorationDeepinThemeWithContext(ctx, flags, make(chan *dbus.Call, 1), deepinThemeName).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ChangeCurrentWorkspaceBackground

func (v *wm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) GoChangeCurrentWorkspaceBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ChangeCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *wm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoChangeCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

func (v *wm) ChangeCurrentWorkspaceBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoChangeCurrentWorkspaceBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), uri).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SwitchToWorkspace

func (v *wm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToWorkspace", flags, ch, backward)
}

func (v *wm) GoSwitchToWorkspaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SwitchToWorkspace", flags, ch, backward)
}

func (v *wm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchToWorkspace(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

func (v *wm) SwitchToWorkspaceWithTimeout(timeout time.Duration, flags dbus.Flags, backward bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSwitchToWorkspaceWithContext(ctx, flags, make(chan *dbus.Call, 1), backward).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PresentWindows

func (v *wm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresentWindows", flags, ch, xids)
}

func (v *wm) GoPresentWindowsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PresentWindows", flags, ch, xids)
}

func (v *wm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	return (<-v.GoPresentWindows(flags, make(chan *dbus.Call, 1), xids).Done).Err
}

func (v *wm) PresentWindowsWithTimeout(timeout time.Duration, flags dbus.Flags, xids []uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPresentWindowsWithContext(ctx, flags, make(chan *dbus.Call, 1), xids).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnableZoneDetected

func (v *wm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableZoneDetected", flags, ch, enabled)
}

func (v *wm) GoEnableZoneDetectedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableZoneDetected", flags, ch, enabled)
}

func (v *wm) EnableZoneDetected(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableZoneDetected(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *wm) EnableZoneDetectedWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableZoneDetectedWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
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

// signal compositingEnabledChanged

func (v *wm) ConnectCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "compositingEnabledChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".compositingEnabledChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var enabled bool
		err := dbus.Store(sig.Body, &enabled)
		if err == nil {
			cb(enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal wmCompositingEnabledChanged

func (v *wm) ConnectWmCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "wmCompositingEnabledChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".wmCompositingEnabledChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var enabled bool
		err := dbus.Store(sig.Body, &enabled)
		if err == nil {
			cb(enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal workspaceCountChanged

func (v *wm) ConnectWorkspaceCountChanged(cb func(count int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "workspaceCountChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".workspaceCountChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var count int32
		err := dbus.Store(sig.Body, &count)
		if err == nil {
			cb(count)
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

// property compositingEnabled b

func (v *wm) CompositingEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "compositingEnabled",
	}
}

// property compositingPossible b

func (v *wm) CompositingPossible() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "compositingPossible",
	}
}

// property compositingAllowSwitch b

func (v *wm) CompositingAllowSwitch() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "compositingAllowSwitch",
	}
}

// property zoneEnabled b

func (v *wm) ZoneEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "zoneEnabled",
	}
}

// property cursorTheme s

func (v *wm) CursorTheme() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "cursorTheme",
	}
}

// property cursorSize i

func (v *wm) CursorSize() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "cursorSize",
	}
}
