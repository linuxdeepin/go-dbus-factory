package launcher

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

type Launcher struct {
	launcher // interface com.deepin.dde.Launcher
	proxy.Object
}

func NewLauncher(conn *dbus.Conn) *Launcher {
	obj := new(Launcher)
	obj.Object.Init_(conn, "com.deepin.dde.Launcher", "/com/deepin/dde/Launcher")
	return obj
}

type launcher struct{}

func (v *launcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*launcher) GetInterfaceName_() string {
	return "com.deepin.dde.Launcher"
}

// method Exit

func (v *launcher) GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *launcher) GoExitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *launcher) Exit(flags dbus.Flags) error {
	return (<-v.GoExit(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *launcher) ExitWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoExitWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Hide

func (v *launcher) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *launcher) GoHideWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *launcher) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *launcher) HideWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHideWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Show

func (v *launcher) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *launcher) GoShowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *launcher) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *launcher) ShowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowByMode

func (v *launcher) GoShowByMode(flags dbus.Flags, ch chan *dbus.Call, in0 int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowByMode", flags, ch, in0)
}

func (v *launcher) GoShowByModeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, in0 int64) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowByMode", flags, ch, in0)
}

func (v *launcher) ShowByMode(flags dbus.Flags, in0 int64) error {
	return (<-v.GoShowByMode(flags, make(chan *dbus.Call, 1), in0).Done).Err
}

func (v *launcher) ShowByModeWithTimeout(timeout time.Duration, flags dbus.Flags, in0 int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowByModeWithContext(ctx, flags, make(chan *dbus.Call, 1), in0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UninstallApp

func (v *launcher) GoUninstallApp(flags dbus.Flags, ch chan *dbus.Call, appKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UninstallApp", flags, ch, appKey)
}

func (v *launcher) GoUninstallAppWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, appKey string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UninstallApp", flags, ch, appKey)
}

func (v *launcher) UninstallApp(flags dbus.Flags, appKey string) error {
	return (<-v.GoUninstallApp(flags, make(chan *dbus.Call, 1), appKey).Done).Err
}

func (v *launcher) UninstallAppWithTimeout(timeout time.Duration, flags dbus.Flags, appKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUninstallAppWithContext(ctx, flags, make(chan *dbus.Call, 1), appKey).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Toggle

func (v *launcher) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *launcher) GoToggleWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *launcher) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *launcher) ToggleWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoToggleWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method IsVisible

func (v *launcher) GoIsVisible(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsVisible", flags, ch)
}

func (v *launcher) GoIsVisibleWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsVisible", flags, ch)
}

func (*launcher) StoreIsVisible(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *launcher) IsVisible(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreIsVisible(
		<-v.GoIsVisible(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *launcher) IsVisibleWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsVisibleWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsVisible(call)
}

// signal Closed

func (v *launcher) ConnectClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Closed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Closed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Shown

func (v *launcher) ConnectShown(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Shown", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Shown",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VisibleChanged

func (v *launcher) ConnectVisibleChanged(cb func(visible bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VisibleChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VisibleChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var visible bool
		err := dbus.Store(sig.Body, &visible)
		if err == nil {
			cb(visible)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Visible b

func (v *launcher) Visible() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Visible",
	}
}
