package sessionmanager

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

type StartManager struct {
	startManager // interface com.deepin.StartManager
	proxy.Object
}

func NewStartManager(conn *dbus.Conn) *StartManager {
	obj := new(StartManager)
	obj.Object.Init_(conn, "com.deepin.SessionManager", "/com/deepin/StartManager")
	return obj
}

type startManager struct{}

func (v *startManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*startManager) GetInterfaceName_() string {
	return "com.deepin.StartManager"
}

// method AddAutostart

func (v *startManager) GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAutostart", flags, ch, arg0)
}

func (v *startManager) GoAddAutostartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddAutostart", flags, ch, arg0)
}

func (*startManager) StoreAddAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) AddAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreAddAutostart(
		<-v.GoAddAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *startManager) AddAutostartWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddAutostartWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAddAutostart(call)
}

// method AutostartList

func (v *startManager) GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AutostartList", flags, ch)
}

func (v *startManager) GoAutostartListWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AutostartList", flags, ch)
}

func (*startManager) StoreAutostartList(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) AutostartList(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreAutostartList(
		<-v.GoAutostartList(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *startManager) AutostartListWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAutostartListWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAutostartList(call)
}

// method DumpMemRecord

func (v *startManager) GoDumpMemRecord(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpMemRecord", flags, ch)
}

func (v *startManager) GoDumpMemRecordWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DumpMemRecord", flags, ch)
}

func (*startManager) StoreDumpMemRecord(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) DumpMemRecord(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreDumpMemRecord(
		<-v.GoDumpMemRecord(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *startManager) DumpMemRecordWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDumpMemRecordWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreDumpMemRecord(call)
}

// method GetApps

func (v *startManager) GoGetApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetApps", flags, ch)
}

func (v *startManager) GoGetAppsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetApps", flags, ch)
}

func (*startManager) StoreGetApps(call *dbus.Call) (arg0 map[uint32]string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) GetApps(flags dbus.Flags) (arg0 map[uint32]string, err error) {
	return v.StoreGetApps(
		<-v.GoGetApps(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *startManager) GetAppsWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 map[uint32]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAppsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetApps(call)
}

// method IsAutostart

func (v *startManager) GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsAutostart", flags, ch, arg0)
}

func (v *startManager) GoIsAutostartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsAutostart", flags, ch, arg0)
}

func (*startManager) StoreIsAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) IsAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreIsAutostart(
		<-v.GoIsAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *startManager) IsAutostartWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsAutostartWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsAutostart(call)
}

// method IsMemSufficient

func (v *startManager) GoIsMemSufficient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMemSufficient", flags, ch)
}

func (v *startManager) GoIsMemSufficientWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsMemSufficient", flags, ch)
}

func (*startManager) StoreIsMemSufficient(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) IsMemSufficient(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreIsMemSufficient(
		<-v.GoIsMemSufficient(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *startManager) IsMemSufficientWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsMemSufficientWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsMemSufficient(call)
}

// method Launch

func (v *startManager) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Launch", flags, ch, arg0)
}

func (v *startManager) GoLaunchWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Launch", flags, ch, arg0)
}

func (*startManager) StoreLaunch(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) Launch(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreLaunch(
		<-v.GoLaunch(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *startManager) LaunchWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLaunchWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLaunch(call)
}

// method LaunchApp

func (v *startManager) GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchApp", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) GoLaunchAppWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LaunchApp", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) LaunchApp(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error {
	return (<-v.GoLaunchApp(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

func (v *startManager) LaunchAppWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLaunchAppWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method LaunchAppAction

func (v *startManager) GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppAction", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) GoLaunchAppActionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LaunchAppAction", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) LaunchAppAction(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error {
	return (<-v.GoLaunchAppAction(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

func (v *startManager) LaunchAppActionWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLaunchAppActionWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method LaunchAppWithOptions

func (v *startManager) GoLaunchAppWithOptions(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppWithOptions", flags, ch, arg0, arg1, arg2, arg3)
}

func (v *startManager) GoLaunchAppWithOptionsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LaunchAppWithOptions", flags, ch, arg0, arg1, arg2, arg3)
}

func (v *startManager) LaunchAppWithOptions(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) error {
	return (<-v.GoLaunchAppWithOptions(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3).Done).Err
}

func (v *startManager) LaunchAppWithOptionsWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLaunchAppWithOptionsWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method LaunchWithTimestamp

func (v *startManager) GoLaunchWithTimestamp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchWithTimestamp", flags, ch, arg0, arg1)
}

func (v *startManager) GoLaunchWithTimestampWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LaunchWithTimestamp", flags, ch, arg0, arg1)
}

func (*startManager) StoreLaunchWithTimestamp(call *dbus.Call) (arg2 bool, err error) {
	err = call.Store(&arg2)
	return
}

func (v *startManager) LaunchWithTimestamp(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 bool, err error) {
	return v.StoreLaunchWithTimestamp(
		<-v.GoLaunchWithTimestamp(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

func (v *startManager) LaunchWithTimestampWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 uint32) (arg2 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLaunchWithTimestampWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLaunchWithTimestamp(call)
}

// method RemoveAutostart

func (v *startManager) GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAutostart", flags, ch, arg0)
}

func (v *startManager) GoRemoveAutostartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveAutostart", flags, ch, arg0)
}

func (*startManager) StoreRemoveAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) RemoveAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreRemoveAutostart(
		<-v.GoRemoveAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *startManager) RemoveAutostartWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveAutostartWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRemoveAutostart(call)
}

// method RunCommand

func (v *startManager) GoRunCommand(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RunCommand", flags, ch, arg0, arg1)
}

func (v *startManager) GoRunCommandWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RunCommand", flags, ch, arg0, arg1)
}

func (v *startManager) RunCommand(flags dbus.Flags, arg0 string, arg1 []string) error {
	return (<-v.GoRunCommand(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *startManager) RunCommandWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRunCommandWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method TryAgain

func (v *startManager) GoTryAgain(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TryAgain", flags, ch, arg0)
}

func (v *startManager) GoTryAgainWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".TryAgain", flags, ch, arg0)
}

func (v *startManager) TryAgain(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoTryAgain(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *startManager) TryAgainWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTryAgainWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal AutostartChanged

func (v *startManager) ConnectAutostartChanged(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AutostartChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AutostartChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property NeededMemory t

func (v *startManager) NeededMemory() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "NeededMemory",
	}
}

type SessionManager struct {
	sessionManager // interface com.deepin.SessionManager
	proxy.Object
}

func NewSessionManager(conn *dbus.Conn) *SessionManager {
	obj := new(SessionManager)
	obj.Object.Init_(conn, "com.deepin.SessionManager", "/com/deepin/SessionManager")
	return obj
}

type sessionManager struct{}

func (v *sessionManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sessionManager) GetInterfaceName_() string {
	return "com.deepin.SessionManager"
}

// method AllowSessionDaemonRun

func (v *sessionManager) GoAllowSessionDaemonRun(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowSessionDaemonRun", flags, ch)
}

func (v *sessionManager) GoAllowSessionDaemonRunWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AllowSessionDaemonRun", flags, ch)
}

func (*sessionManager) StoreAllowSessionDaemonRun(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) AllowSessionDaemonRun(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreAllowSessionDaemonRun(
		<-v.GoAllowSessionDaemonRun(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) AllowSessionDaemonRunWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAllowSessionDaemonRunWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAllowSessionDaemonRun(call)
}

// method CanHibernate

func (v *sessionManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (v *sessionManager) GoCanHibernateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*sessionManager) StoreCanHibernate(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanHibernate(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) CanHibernateWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCanHibernateWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCanHibernate(call)
}

// method CanLogout

func (v *sessionManager) GoCanLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanLogout", flags, ch)
}

func (v *sessionManager) GoCanLogoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CanLogout", flags, ch)
}

func (*sessionManager) StoreCanLogout(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanLogout(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanLogout(
		<-v.GoCanLogout(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) CanLogoutWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCanLogoutWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCanLogout(call)
}

// method CanReboot

func (v *sessionManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (v *sessionManager) GoCanRebootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*sessionManager) StoreCanReboot(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanReboot(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) CanRebootWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCanRebootWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCanReboot(call)
}

// method CanShutdown

func (v *sessionManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (v *sessionManager) GoCanShutdownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (*sessionManager) StoreCanShutdown(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanShutdown(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanShutdown(
		<-v.GoCanShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) CanShutdownWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCanShutdownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCanShutdown(call)
}

// method CanSuspend

func (v *sessionManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (v *sessionManager) GoCanSuspendWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*sessionManager) StoreCanSuspend(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanSuspend(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *sessionManager) CanSuspendWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCanSuspendWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCanSuspend(call)
}

// method ForceLogout

func (v *sessionManager) GoForceLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceLogout", flags, ch)
}

func (v *sessionManager) GoForceLogoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ForceLogout", flags, ch)
}

func (v *sessionManager) ForceLogout(flags dbus.Flags) error {
	return (<-v.GoForceLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) ForceLogoutWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoForceLogoutWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ForceReboot

func (v *sessionManager) GoForceReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceReboot", flags, ch)
}

func (v *sessionManager) GoForceRebootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ForceReboot", flags, ch)
}

func (v *sessionManager) ForceReboot(flags dbus.Flags) error {
	return (<-v.GoForceReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) ForceRebootWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoForceRebootWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ForceShutdown

func (v *sessionManager) GoForceShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceShutdown", flags, ch)
}

func (v *sessionManager) GoForceShutdownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ForceShutdown", flags, ch)
}

func (v *sessionManager) ForceShutdown(flags dbus.Flags) error {
	return (<-v.GoForceShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) ForceShutdownWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoForceShutdownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Logout

func (v *sessionManager) GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *sessionManager) GoLogoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *sessionManager) Logout(flags dbus.Flags) error {
	return (<-v.GoLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) LogoutWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLogoutWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PowerOffChoose

func (v *sessionManager) GoPowerOffChoose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOffChoose", flags, ch)
}

func (v *sessionManager) GoPowerOffChooseWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PowerOffChoose", flags, ch)
}

func (v *sessionManager) PowerOffChoose(flags dbus.Flags) error {
	return (<-v.GoPowerOffChoose(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) PowerOffChooseWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPowerOffChooseWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Reboot

func (v *sessionManager) GoReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *sessionManager) GoRebootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *sessionManager) Reboot(flags dbus.Flags) error {
	return (<-v.GoReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RebootWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRebootWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Register

func (v *sessionManager) GoRegister(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, arg0)
}

func (v *sessionManager) GoRegisterWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Register", flags, ch, arg0)
}

func (*sessionManager) StoreRegister(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *sessionManager) Register(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreRegister(
		<-v.GoRegister(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *sessionManager) RegisterWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRegisterWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRegister(call)
}

// method RequestHibernate

func (v *sessionManager) GoRequestHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestHibernate", flags, ch)
}

func (v *sessionManager) GoRequestHibernateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestHibernate", flags, ch)
}

func (v *sessionManager) RequestHibernate(flags dbus.Flags) error {
	return (<-v.GoRequestHibernate(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestHibernateWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestHibernateWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestLock

func (v *sessionManager) GoRequestLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLock", flags, ch)
}

func (v *sessionManager) GoRequestLockWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestLock", flags, ch)
}

func (v *sessionManager) RequestLock(flags dbus.Flags) error {
	return (<-v.GoRequestLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestLockWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestLockWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestLogout

func (v *sessionManager) GoRequestLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLogout", flags, ch)
}

func (v *sessionManager) GoRequestLogoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestLogout", flags, ch)
}

func (v *sessionManager) RequestLogout(flags dbus.Flags) error {
	return (<-v.GoRequestLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestLogoutWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestLogoutWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestReboot

func (v *sessionManager) GoRequestReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestReboot", flags, ch)
}

func (v *sessionManager) GoRequestRebootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestReboot", flags, ch)
}

func (v *sessionManager) RequestReboot(flags dbus.Flags) error {
	return (<-v.GoRequestReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestRebootWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestRebootWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestShutdown

func (v *sessionManager) GoRequestShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestShutdown", flags, ch)
}

func (v *sessionManager) GoRequestShutdownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestShutdown", flags, ch)
}

func (v *sessionManager) RequestShutdown(flags dbus.Flags) error {
	return (<-v.GoRequestShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestShutdownWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestShutdownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RequestSuspend

func (v *sessionManager) GoRequestSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSuspend", flags, ch)
}

func (v *sessionManager) GoRequestSuspendWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestSuspend", flags, ch)
}

func (v *sessionManager) RequestSuspend(flags dbus.Flags) error {
	return (<-v.GoRequestSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) RequestSuspendWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestSuspendWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLocked

func (v *sessionManager) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocked", flags, ch, arg0)
}

func (v *sessionManager) GoSetLockedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLocked", flags, ch, arg0)
}

func (v *sessionManager) SetLocked(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoSetLocked(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *sessionManager) SetLockedWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLockedWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Shutdown

func (v *sessionManager) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *sessionManager) GoShutdownWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *sessionManager) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) ShutdownWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShutdownWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ToggleDebug

func (v *sessionManager) GoToggleDebug(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleDebug", flags, ch)
}

func (v *sessionManager) GoToggleDebugWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ToggleDebug", flags, ch)
}

func (v *sessionManager) ToggleDebug(flags dbus.Flags) error {
	return (<-v.GoToggleDebug(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sessionManager) ToggleDebugWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoToggleDebugWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal Unlock

func (v *sessionManager) ConnectUnlock(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Unlock", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Unlock",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Locked b

func (v *sessionManager) Locked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property CurrentUid s

func (v *sessionManager) CurrentUid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CurrentUid",
	}
}

// property Stage i

func (v *sessionManager) Stage() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Stage",
	}
}

type WMSwitcher struct {
	wmSwitcher // interface com.deepin.WMSwitcher
	proxy.Object
}

func NewWMSwitcher(conn *dbus.Conn) *WMSwitcher {
	obj := new(WMSwitcher)
	obj.Object.Init_(conn, "com.deepin.SessionManager", "/com/deepin/WMSwitcher")
	return obj
}

type wmSwitcher struct{}

func (v *wmSwitcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wmSwitcher) GetInterfaceName_() string {
	return "com.deepin.WMSwitcher"
}

// method CurrentWM

func (v *wmSwitcher) GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CurrentWM", flags, ch)
}

func (v *wmSwitcher) GoCurrentWMWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CurrentWM", flags, ch)
}

func (*wmSwitcher) StoreCurrentWM(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *wmSwitcher) CurrentWM(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreCurrentWM(
		<-v.GoCurrentWM(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *wmSwitcher) CurrentWMWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCurrentWMWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCurrentWM(call)
}

// method RequestSwitchWM

func (v *wmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *wmSwitcher) GoRequestSwitchWMWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *wmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	return (<-v.GoRequestSwitchWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wmSwitcher) RequestSwitchWMWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestSwitchWMWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RestartWM

func (v *wmSwitcher) GoRestartWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestartWM", flags, ch)
}

func (v *wmSwitcher) GoRestartWMWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RestartWM", flags, ch)
}

func (v *wmSwitcher) RestartWM(flags dbus.Flags) error {
	return (<-v.GoRestartWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wmSwitcher) RestartWMWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRestartWMWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Start2DWM

func (v *wmSwitcher) GoStart2DWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start2DWM", flags, ch)
}

func (v *wmSwitcher) GoStart2DWMWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Start2DWM", flags, ch)
}

func (v *wmSwitcher) Start2DWM(flags dbus.Flags) error {
	return (<-v.GoStart2DWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *wmSwitcher) Start2DWMWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStart2DWMWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal WMChanged

func (v *wmSwitcher) ConnectWMChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WMChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WMChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type XSettings struct {
	xSettings // interface com.deepin.XSettings
	proxy.Object
}

func NewXSettings(conn *dbus.Conn) *XSettings {
	obj := new(XSettings)
	obj.Object.Init_(conn, "com.deepin.SessionManager", "/com/deepin/XSettings")
	return obj
}

type xSettings struct{}

func (v *xSettings) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*xSettings) GetInterfaceName_() string {
	return "com.deepin.XSettings"
}

// method GetColor

func (v *xSettings) GoGetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetColor", flags, ch, arg0)
}

func (v *xSettings) GoGetColorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetColor", flags, ch, arg0)
}

func (*xSettings) StoreGetColor(call *dbus.Call) (arg1 []int16, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetColor(flags dbus.Flags, arg0 string) (arg1 []int16, err error) {
	return v.StoreGetColor(
		<-v.GoGetColor(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *xSettings) GetColorWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 []int16, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetColorWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetColor(call)
}

// method GetInteger

func (v *xSettings) GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetInteger", flags, ch, arg0)
}

func (v *xSettings) GoGetIntegerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetInteger", flags, ch, arg0)
}

func (*xSettings) StoreGetInteger(call *dbus.Call) (arg1 int32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetInteger(flags dbus.Flags, arg0 string) (arg1 int32, err error) {
	return v.StoreGetInteger(
		<-v.GoGetInteger(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *xSettings) GetIntegerWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetIntegerWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetInteger(call)
}

// method GetScaleFactor

func (v *xSettings) GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScaleFactor", flags, ch)
}

func (v *xSettings) GoGetScaleFactorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetScaleFactor", flags, ch)
}

func (*xSettings) StoreGetScaleFactor(call *dbus.Call) (arg0 float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *xSettings) GetScaleFactor(flags dbus.Flags) (arg0 float64, err error) {
	return v.StoreGetScaleFactor(
		<-v.GoGetScaleFactor(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *xSettings) GetScaleFactorWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 float64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetScaleFactorWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetScaleFactor(call)
}

// method GetScreenScaleFactors

func (v *xSettings) GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScreenScaleFactors", flags, ch)
}

func (v *xSettings) GoGetScreenScaleFactorsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetScreenScaleFactors", flags, ch)
}

func (*xSettings) StoreGetScreenScaleFactors(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *xSettings) GetScreenScaleFactors(flags dbus.Flags) (arg0 map[string]float64, err error) {
	return v.StoreGetScreenScaleFactors(
		<-v.GoGetScreenScaleFactors(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *xSettings) GetScreenScaleFactorsWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 map[string]float64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetScreenScaleFactorsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetScreenScaleFactors(call)
}

// method GetString

func (v *xSettings) GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetString", flags, ch, arg0)
}

func (v *xSettings) GoGetStringWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetString", flags, ch, arg0)
}

func (*xSettings) StoreGetString(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetString(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetString(
		<-v.GoGetString(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *xSettings) GetStringWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetStringWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetString(call)
}

// method ListProps

func (v *xSettings) GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListProps", flags, ch)
}

func (v *xSettings) GoListPropsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListProps", flags, ch)
}

func (*xSettings) StoreListProps(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *xSettings) ListProps(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreListProps(
		<-v.GoListProps(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *xSettings) ListPropsWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListPropsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListProps(call)
}

// method SetColor

func (v *xSettings) GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetColor", flags, ch, arg0, arg1)
}

func (v *xSettings) GoSetColorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetColor", flags, ch, arg0, arg1)
}

func (v *xSettings) SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error {
	return (<-v.GoSetColor(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *xSettings) SetColorWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 []int16) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetColorWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetInteger

func (v *xSettings) GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetInteger", flags, ch, arg0, arg1)
}

func (v *xSettings) GoSetIntegerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetInteger", flags, ch, arg0, arg1)
}

func (v *xSettings) SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error {
	return (<-v.GoSetInteger(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *xSettings) SetIntegerWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetIntegerWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetScaleFactor

func (v *xSettings) GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScaleFactor", flags, ch, arg0)
}

func (v *xSettings) GoSetScaleFactorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetScaleFactor", flags, ch, arg0)
}

func (v *xSettings) SetScaleFactor(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetScaleFactor(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *xSettings) SetScaleFactorWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetScaleFactorWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetScreenScaleFactors

func (v *xSettings) GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScreenScaleFactors", flags, ch, arg0)
}

func (v *xSettings) GoSetScreenScaleFactorsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetScreenScaleFactors", flags, ch, arg0)
}

func (v *xSettings) SetScreenScaleFactors(flags dbus.Flags, arg0 map[string]float64) error {
	return (<-v.GoSetScreenScaleFactors(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *xSettings) SetScreenScaleFactorsWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 map[string]float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetScreenScaleFactorsWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetString

func (v *xSettings) GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetString", flags, ch, arg0, arg1)
}

func (v *xSettings) GoSetStringWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetString", flags, ch, arg0, arg1)
}

func (v *xSettings) SetString(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoSetString(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *xSettings) SetStringWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetStringWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal SetScaleFactorStarted

func (v *xSettings) ConnectSetScaleFactorStarted(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorStarted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorStarted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SetScaleFactorDone

func (v *xSettings) ConnectSetScaleFactorDone(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
