package apps

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

type Apps struct {
	desktopFileWatcher // interface com.deepin.daemon.Apps.DesktopFileWatcher
	launchedRecorder   // interface com.deepin.daemon.Apps.LaunchedRecorder
	proxy.Object
}

func NewApps(conn *dbus.Conn) *Apps {
	obj := new(Apps)
	obj.Object.Init_(conn, "com.deepin.daemon.Apps", "/com/deepin/daemon/Apps")
	return obj
}

func (obj *Apps) DesktopFileWatcher() *desktopFileWatcher {
	return &obj.desktopFileWatcher
}

type desktopFileWatcher struct{}

func (v *desktopFileWatcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*desktopFileWatcher) GetInterfaceName_() string {
	return "com.deepin.daemon.Apps.DesktopFileWatcher"
}

// signal Event

func (v *desktopFileWatcher) ConnectEvent(cb func(name string, op uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Event", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Event",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		var op uint32
		err := dbus.Store(sig.Body, &name, &op)
		if err == nil {
			cb(name, op)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

func (obj *Apps) LaunchedRecorder() *launchedRecorder {
	return &obj.launchedRecorder
}

type launchedRecorder struct{}

func (v *launchedRecorder) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*launchedRecorder) GetInterfaceName_() string {
	return "com.deepin.daemon.Apps.LaunchedRecorder"
}

// method GetNew

func (v *launchedRecorder) GoGetNew(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNew", flags, ch)
}

func (v *launchedRecorder) GoGetNewWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetNew", flags, ch)
}

func (*launchedRecorder) StoreGetNew(call *dbus.Call) (newApps map[string][]string, err error) {
	err = call.Store(&newApps)
	return
}

func (v *launchedRecorder) GetNew(flags dbus.Flags) (newApps map[string][]string, err error) {
	return v.StoreGetNew(
		<-v.GoGetNew(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *launchedRecorder) GetNewWithTimeout(timeout time.Duration, flags dbus.Flags) (newApps map[string][]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetNewWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetNew(call)
}

// method MarkLaunched

func (v *launchedRecorder) GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MarkLaunched", flags, ch, desktopFile)
}

func (v *launchedRecorder) GoMarkLaunchedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MarkLaunched", flags, ch, desktopFile)
}

func (v *launchedRecorder) MarkLaunched(flags dbus.Flags, desktopFile string) error {
	return (<-v.GoMarkLaunched(flags, make(chan *dbus.Call, 1), desktopFile).Done).Err
}

func (v *launchedRecorder) MarkLaunchedWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFile string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMarkLaunchedWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFile).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UninstallHints

func (v *launchedRecorder) GoUninstallHints(flags dbus.Flags, ch chan *dbus.Call, desktopFiles []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UninstallHints", flags, ch, desktopFiles)
}

func (v *launchedRecorder) GoUninstallHintsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, desktopFiles []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UninstallHints", flags, ch, desktopFiles)
}

func (v *launchedRecorder) UninstallHints(flags dbus.Flags, desktopFiles []string) error {
	return (<-v.GoUninstallHints(flags, make(chan *dbus.Call, 1), desktopFiles).Done).Err
}

func (v *launchedRecorder) UninstallHintsWithTimeout(timeout time.Duration, flags dbus.Flags, desktopFiles []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUninstallHintsWithContext(ctx, flags, make(chan *dbus.Call, 1), desktopFiles).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method WatchDirs

func (v *launchedRecorder) GoWatchDirs(flags dbus.Flags, ch chan *dbus.Call, dirs []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WatchDirs", flags, ch, dirs)
}

func (v *launchedRecorder) GoWatchDirsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, dirs []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".WatchDirs", flags, ch, dirs)
}

func (v *launchedRecorder) WatchDirs(flags dbus.Flags, dirs []string) error {
	return (<-v.GoWatchDirs(flags, make(chan *dbus.Call, 1), dirs).Done).Err
}

func (v *launchedRecorder) WatchDirsWithTimeout(timeout time.Duration, flags dbus.Flags, dirs []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoWatchDirsWithContext(ctx, flags, make(chan *dbus.Call, 1), dirs).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal Launched

func (v *launchedRecorder) ConnectLaunched(cb func(file string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Launched", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Launched",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var file string
		err := dbus.Store(sig.Body, &file)
		if err == nil {
			cb(file)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StatusSaved

func (v *launchedRecorder) ConnectStatusSaved(cb func(root string, file string, ok bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StatusSaved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StatusSaved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var root string
		var file string
		var ok bool
		err := dbus.Store(sig.Body, &root, &file, &ok)
		if err == nil {
			cb(root, file, ok)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ServiceRestarted

func (v *launchedRecorder) ConnectServiceRestarted(cb func()) (dbusutil.SignalHandlerId, error) {
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
