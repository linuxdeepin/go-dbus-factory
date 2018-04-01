package bamf

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

type Control struct {
	control // interface org.ayatana.bamf.control
	proxy.Object
}

func NewControl(conn *dbus.Conn) *Control {
	obj := new(Control)
	obj.Object.Init_(conn, "org.ayatana.bamf", "/org/ayatana/bamf/control")
	return obj
}

type control struct{}

func (v *control) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*control) GetInterfaceName_() string {
	return "org.ayatana.bamf.control"
}

// method Quit

func (v *control) GoQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Quit", flags, ch)
}

func (v *control) Quit(flags dbus.Flags) error {
	return (<-v.GoQuit(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method InsertDesktopFile

func (v *control) GoInsertDesktopFile(flags dbus.Flags, ch chan *dbus.Call, desktop_path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InsertDesktopFile", flags, ch, desktop_path)
}

func (v *control) InsertDesktopFile(flags dbus.Flags, desktop_path string) error {
	return (<-v.GoInsertDesktopFile(flags, make(chan *dbus.Call, 1), desktop_path).Done).Err
}

// method RegisterApplicationForPid

func (v *control) GoRegisterApplicationForPid(flags dbus.Flags, ch chan *dbus.Call, application string, pid int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterApplicationForPid", flags, ch, application, pid)
}

func (v *control) RegisterApplicationForPid(flags dbus.Flags, application string, pid int32) error {
	return (<-v.GoRegisterApplicationForPid(flags, make(chan *dbus.Call, 1), application, pid).Done).Err
}

// method CreateLocalDesktopFile

func (v *control) GoCreateLocalDesktopFile(flags dbus.Flags, ch chan *dbus.Call, application string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateLocalDesktopFile", flags, ch, application)
}

func (v *control) CreateLocalDesktopFile(flags dbus.Flags, application string) error {
	return (<-v.GoCreateLocalDesktopFile(flags, make(chan *dbus.Call, 1), application).Done).Err
}

// method OmNomNomDesktopFile

func (v *control) GoOmNomNomDesktopFile(flags dbus.Flags, ch chan *dbus.Call, desktop_path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OmNomNomDesktopFile", flags, ch, desktop_path)
}

func (v *control) OmNomNomDesktopFile(flags dbus.Flags, desktop_path string) error {
	return (<-v.GoOmNomNomDesktopFile(flags, make(chan *dbus.Call, 1), desktop_path).Done).Err
}

type Matcher struct {
	matcher // interface org.ayatana.bamf.matcher
	proxy.Object
}

func NewMatcher(conn *dbus.Conn) *Matcher {
	obj := new(Matcher)
	obj.Object.Init_(conn, "org.ayatana.bamf", "/org/ayatana/bamf/matcher")
	return obj
}

type matcher struct{}

func (v *matcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*matcher) GetInterfaceName_() string {
	return "org.ayatana.bamf.matcher"
}

// method XidsForApplication

func (v *matcher) GoXidsForApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".XidsForApplication", flags, ch, desktop_file)
}

func (*matcher) StoreXidsForApplication(call *dbus.Call) (xids []uint32, err error) {
	err = call.Store(&xids)
	return
}

func (v *matcher) XidsForApplication(flags dbus.Flags, desktop_file string) (xids []uint32, err error) {
	return v.StoreXidsForApplication(
		<-v.GoXidsForApplication(flags, make(chan *dbus.Call, 1), desktop_file).Done)
}

// method TabPaths

func (v *matcher) GoTabPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TabPaths", flags, ch)
}

func (*matcher) StoreTabPaths(call *dbus.Call) (paths []string, err error) {
	err = call.Store(&paths)
	return
}

func (v *matcher) TabPaths(flags dbus.Flags) (paths []string, err error) {
	return v.StoreTabPaths(
		<-v.GoTabPaths(flags, make(chan *dbus.Call, 1)).Done)
}

// method RunningApplications

func (v *matcher) GoRunningApplications(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RunningApplications", flags, ch)
}

func (*matcher) StoreRunningApplications(call *dbus.Call) (paths []string, err error) {
	err = call.Store(&paths)
	return
}

func (v *matcher) RunningApplications(flags dbus.Flags) (paths []string, err error) {
	return v.StoreRunningApplications(
		<-v.GoRunningApplications(flags, make(chan *dbus.Call, 1)).Done)
}

// method RunningApplicationsDesktopFiles

func (v *matcher) GoRunningApplicationsDesktopFiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RunningApplicationsDesktopFiles", flags, ch)
}

func (*matcher) StoreRunningApplicationsDesktopFiles(call *dbus.Call) (paths []string, err error) {
	err = call.Store(&paths)
	return
}

func (v *matcher) RunningApplicationsDesktopFiles(flags dbus.Flags) (paths []string, err error) {
	return v.StoreRunningApplicationsDesktopFiles(
		<-v.GoRunningApplicationsDesktopFiles(flags, make(chan *dbus.Call, 1)).Done)
}

// method RegisterFavorites

func (v *matcher) GoRegisterFavorites(flags dbus.Flags, ch chan *dbus.Call, favorites []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterFavorites", flags, ch, favorites)
}

func (v *matcher) RegisterFavorites(flags dbus.Flags, favorites []string) error {
	return (<-v.GoRegisterFavorites(flags, make(chan *dbus.Call, 1), favorites).Done).Err
}

// method PathForApplication

func (v *matcher) GoPathForApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PathForApplication", flags, ch, desktop_file)
}

func (*matcher) StorePathForApplication(call *dbus.Call) (path string, err error) {
	err = call.Store(&path)
	return
}

func (v *matcher) PathForApplication(flags dbus.Flags, desktop_file string) (path string, err error) {
	return v.StorePathForApplication(
		<-v.GoPathForApplication(flags, make(chan *dbus.Call, 1), desktop_file).Done)
}

// method WindowPaths

func (v *matcher) GoWindowPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WindowPaths", flags, ch)
}

func (*matcher) StoreWindowPaths(call *dbus.Call) (paths []string, err error) {
	err = call.Store(&paths)
	return
}

func (v *matcher) WindowPaths(flags dbus.Flags) (paths []string, err error) {
	return v.StoreWindowPaths(
		<-v.GoWindowPaths(flags, make(chan *dbus.Call, 1)).Done)
}

// method ApplicationPaths

func (v *matcher) GoApplicationPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationPaths", flags, ch)
}

func (*matcher) StoreApplicationPaths(call *dbus.Call) (paths []string, err error) {
	err = call.Store(&paths)
	return
}

func (v *matcher) ApplicationPaths(flags dbus.Flags) (paths []string, err error) {
	return v.StoreApplicationPaths(
		<-v.GoApplicationPaths(flags, make(chan *dbus.Call, 1)).Done)
}

// method ApplicationIsRunning

func (v *matcher) GoApplicationIsRunning(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationIsRunning", flags, ch, desktop_file)
}

func (*matcher) StoreApplicationIsRunning(call *dbus.Call) (running bool, err error) {
	err = call.Store(&running)
	return
}

func (v *matcher) ApplicationIsRunning(flags dbus.Flags, desktop_file string) (running bool, err error) {
	return v.StoreApplicationIsRunning(
		<-v.GoApplicationIsRunning(flags, make(chan *dbus.Call, 1), desktop_file).Done)
}

// method ApplicationForXid

func (v *matcher) GoApplicationForXid(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationForXid", flags, ch, xid)
}

func (*matcher) StoreApplicationForXid(call *dbus.Call) (application string, err error) {
	err = call.Store(&application)
	return
}

func (v *matcher) ApplicationForXid(flags dbus.Flags, xid uint32) (application string, err error) {
	return v.StoreApplicationForXid(
		<-v.GoApplicationForXid(flags, make(chan *dbus.Call, 1), xid).Done)
}

// method ActiveWindow

func (v *matcher) GoActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActiveWindow", flags, ch)
}

func (*matcher) StoreActiveWindow(call *dbus.Call) (window string, err error) {
	err = call.Store(&window)
	return
}

func (v *matcher) ActiveWindow(flags dbus.Flags) (window string, err error) {
	return v.StoreActiveWindow(
		<-v.GoActiveWindow(flags, make(chan *dbus.Call, 1)).Done)
}

// method ActiveApplication

func (v *matcher) GoActiveApplication(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActiveApplication", flags, ch)
}

func (*matcher) StoreActiveApplication(call *dbus.Call) (application string, err error) {
	err = call.Store(&application)
	return
}

func (v *matcher) ActiveApplication(flags dbus.Flags) (application string, err error) {
	return v.StoreActiveApplication(
		<-v.GoActiveApplication(flags, make(chan *dbus.Call, 1)).Done)
}

// method WindowStackForMonitor

func (v *matcher) GoWindowStackForMonitor(flags dbus.Flags, ch chan *dbus.Call, monitor_id int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WindowStackForMonitor", flags, ch, monitor_id)
}

func (*matcher) StoreWindowStackForMonitor(call *dbus.Call) (window_list []string, err error) {
	err = call.Store(&window_list)
	return
}

func (v *matcher) WindowStackForMonitor(flags dbus.Flags, monitor_id int32) (window_list []string, err error) {
	return v.StoreWindowStackForMonitor(
		<-v.GoWindowStackForMonitor(flags, make(chan *dbus.Call, 1), monitor_id).Done)
}

// signal ActiveApplicationChanged

func (v *matcher) ConnectActiveApplicationChanged(cb func(old_app string, new_app string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActiveApplicationChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActiveApplicationChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old_app string
		var new_app string
		err := dbus.Store(sig.Body, &old_app, &new_app)
		if err == nil {
			cb(old_app, new_app)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActiveWindowChanged

func (v *matcher) ConnectActiveWindowChanged(cb func(old_win string, new_win string)) (dbusutil.SignalHandlerId, error) {
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
		var old_win string
		var new_win string
		err := dbus.Store(sig.Body, &old_win, &new_win)
		if err == nil {
			cb(old_win, new_win)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ViewClosed

func (v *matcher) ConnectViewClosed(cb func(path string, type0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ViewClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ViewClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		var type0 string
		err := dbus.Store(sig.Body, &path, &type0)
		if err == nil {
			cb(path, type0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ViewOpened

func (v *matcher) ConnectViewOpened(cb func(path string, type0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ViewOpened", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ViewOpened",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		var type0 string
		err := dbus.Store(sig.Body, &path, &type0)
		if err == nil {
			cb(path, type0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StackingOrderChanged

func (v *matcher) ConnectStackingOrderChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StackingOrderChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StackingOrderChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal RunningApplicationsChanged

func (v *matcher) ConnectRunningApplicationsChanged(cb func(opened_desktop_files []string, closed_desktop_files []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "RunningApplicationsChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".RunningApplicationsChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var opened_desktop_files []string
		var closed_desktop_files []string
		err := dbus.Store(sig.Body, &opened_desktop_files, &closed_desktop_files)
		if err == nil {
			cb(opened_desktop_files, closed_desktop_files)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Application struct {
	view        // interface org.ayatana.bamf.view
	application // interface org.ayatana.bamf.application
	proxy.Object
}

func NewApplication(conn *dbus.Conn, path dbus.ObjectPath) (*Application, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Application)
	obj.Object.Init_(conn, "org.ayatana.bamf", path)
	return obj, nil
}

type view struct{}

func (v *view) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*view) GetInterfaceName_() string {
	return "org.ayatana.bamf.view"
}

// method ViewType

func (v *view) GoViewType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ViewType", flags, ch)
}

func (*view) StoreViewType(call *dbus.Call) (view_type string, err error) {
	err = call.Store(&view_type)
	return
}

func (v *view) ViewType(flags dbus.Flags) (view_type string, err error) {
	return v.StoreViewType(
		<-v.GoViewType(flags, make(chan *dbus.Call, 1)).Done)
}

// method Icon

func (v *view) GoIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Icon", flags, ch)
}

func (*view) StoreIcon(call *dbus.Call) (name string, err error) {
	err = call.Store(&name)
	return
}

func (v *view) Icon(flags dbus.Flags) (name string, err error) {
	return v.StoreIcon(
		<-v.GoIcon(flags, make(chan *dbus.Call, 1)).Done)
}

// method Name

func (v *view) GoName(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Name", flags, ch)
}

func (*view) StoreName(call *dbus.Call) (name string, err error) {
	err = call.Store(&name)
	return
}

func (v *view) Name(flags dbus.Flags) (name string, err error) {
	return v.StoreName(
		<-v.GoName(flags, make(chan *dbus.Call, 1)).Done)
}

// method UserVisible

func (v *view) GoUserVisible(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UserVisible", flags, ch)
}

func (*view) StoreUserVisible(call *dbus.Call) (visible bool, err error) {
	err = call.Store(&visible)
	return
}

func (v *view) UserVisible(flags dbus.Flags) (visible bool, err error) {
	return v.StoreUserVisible(
		<-v.GoUserVisible(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsUrgent

func (v *view) GoIsUrgent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsUrgent", flags, ch)
}

func (*view) StoreIsUrgent(call *dbus.Call) (urgent bool, err error) {
	err = call.Store(&urgent)
	return
}

func (v *view) IsUrgent(flags dbus.Flags) (urgent bool, err error) {
	return v.StoreIsUrgent(
		<-v.GoIsUrgent(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsRunning

func (v *view) GoIsRunning(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsRunning", flags, ch)
}

func (*view) StoreIsRunning(call *dbus.Call) (running bool, err error) {
	err = call.Store(&running)
	return
}

func (v *view) IsRunning(flags dbus.Flags) (running bool, err error) {
	return v.StoreIsRunning(
		<-v.GoIsRunning(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsActive

func (v *view) GoIsActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsActive", flags, ch)
}

func (*view) StoreIsActive(call *dbus.Call) (active bool, err error) {
	err = call.Store(&active)
	return
}

func (v *view) IsActive(flags dbus.Flags) (active bool, err error) {
	return v.StoreIsActive(
		<-v.GoIsActive(flags, make(chan *dbus.Call, 1)).Done)
}

// method Parents

func (v *view) GoParents(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Parents", flags, ch)
}

func (*view) StoreParents(call *dbus.Call) (parents_paths []string, err error) {
	err = call.Store(&parents_paths)
	return
}

func (v *view) Parents(flags dbus.Flags) (parents_paths []string, err error) {
	return v.StoreParents(
		<-v.GoParents(flags, make(chan *dbus.Call, 1)).Done)
}

// method Children

func (v *view) GoChildren(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Children", flags, ch)
}

func (*view) StoreChildren(call *dbus.Call) (children_paths []string, err error) {
	err = call.Store(&children_paths)
	return
}

func (v *view) Children(flags dbus.Flags) (children_paths []string, err error) {
	return v.StoreChildren(
		<-v.GoChildren(flags, make(chan *dbus.Call, 1)).Done)
}

// signal NameChanged

func (v *view) ConnectNameChanged(cb func(old_name string, new_name string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old_name string
		var new_name string
		err := dbus.Store(sig.Body, &old_name, &new_name)
		if err == nil {
			cb(old_name, new_name)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UserVisibleChanged

func (v *view) ConnectUserVisibleChanged(cb func(user_visible bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserVisibleChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserVisibleChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var user_visible bool
		err := dbus.Store(sig.Body, &user_visible)
		if err == nil {
			cb(user_visible)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UrgentChanged

func (v *view) ConnectUrgentChanged(cb func(is_urgent bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UrgentChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UrgentChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var is_urgent bool
		err := dbus.Store(sig.Body, &is_urgent)
		if err == nil {
			cb(is_urgent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal RunningChanged

func (v *view) ConnectRunningChanged(cb func(is_running bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "RunningChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".RunningChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var is_running bool
		err := dbus.Store(sig.Body, &is_running)
		if err == nil {
			cb(is_running)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActiveChanged

func (v *view) ConnectActiveChanged(cb func(is_active bool)) (dbusutil.SignalHandlerId, error) {
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
		var is_active bool
		err := dbus.Store(sig.Body, &is_active)
		if err == nil {
			cb(is_active)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ChildRemoved

func (v *view) ConnectChildRemoved(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChildRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChildRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ChildAdded

func (v *view) ConnectChildAdded(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChildAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChildAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Closed

func (v *view) ConnectClosed(cb func()) (dbusutil.SignalHandlerId, error) {
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

// property Name s

func (v *view) PropName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Icon s

func (v *view) PropIcon() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Icon",
	}
}

// property UserVisible b

func (v *view) PropUserVisible() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "UserVisible",
	}
}

// property Running b

func (v *view) Running() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Running",
	}
}

// property Starting b

func (v *view) Starting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Starting",
	}
}

// property Urgent b

func (v *view) Urgent() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Urgent",
	}
}

// property Active b

func (v *view) Active() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Active",
	}
}

type application struct{}

func (v *application) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*application) GetInterfaceName_() string {
	return "org.ayatana.bamf.application"
}

// method ShowStubs

func (v *application) GoShowStubs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowStubs", flags, ch)
}

func (*application) StoreShowStubs(call *dbus.Call) (show_stubs bool, err error) {
	err = call.Store(&show_stubs)
	return
}

func (v *application) ShowStubs(flags dbus.Flags) (show_stubs bool, err error) {
	return v.StoreShowStubs(
		<-v.GoShowStubs(flags, make(chan *dbus.Call, 1)).Done)
}

// method Xids

func (v *application) GoXids(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Xids", flags, ch)
}

func (*application) StoreXids(call *dbus.Call) (xids []uint32, err error) {
	err = call.Store(&xids)
	return
}

func (v *application) Xids(flags dbus.Flags) (xids []uint32, err error) {
	return v.StoreXids(
		<-v.GoXids(flags, make(chan *dbus.Call, 1)).Done)
}

// method DesktopFile

func (v *application) GoDesktopFile(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DesktopFile", flags, ch)
}

func (*application) StoreDesktopFile(call *dbus.Call) (desktop_file string, err error) {
	err = call.Store(&desktop_file)
	return
}

func (v *application) DesktopFile(flags dbus.Flags) (desktop_file string, err error) {
	return v.StoreDesktopFile(
		<-v.GoDesktopFile(flags, make(chan *dbus.Call, 1)).Done)
}

// method SupportedMimeTypes

func (v *application) GoSupportedMimeTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SupportedMimeTypes", flags, ch)
}

func (*application) StoreSupportedMimeTypes(call *dbus.Call) (mime_types []string, err error) {
	err = call.Store(&mime_types)
	return
}

func (v *application) SupportedMimeTypes(flags dbus.Flags) (mime_types []string, err error) {
	return v.StoreSupportedMimeTypes(
		<-v.GoSupportedMimeTypes(flags, make(chan *dbus.Call, 1)).Done)
}

// method ApplicationType

func (v *application) GoApplicationType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationType", flags, ch)
}

func (*application) StoreApplicationType(call *dbus.Call) (type0 string, err error) {
	err = call.Store(&type0)
	return
}

func (v *application) ApplicationType(flags dbus.Flags) (type0 string, err error) {
	return v.StoreApplicationType(
		<-v.GoApplicationType(flags, make(chan *dbus.Call, 1)).Done)
}

// method ApplicationMenu

func (v *application) GoApplicationMenu(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationMenu", flags, ch)
}

func (*application) StoreApplicationMenu(call *dbus.Call) (busname string, objectpath string, err error) {
	err = call.Store(&busname, &objectpath)
	return
}

func (v *application) ApplicationMenu(flags dbus.Flags) (busname string, objectpath string, err error) {
	return v.StoreApplicationMenu(
		<-v.GoApplicationMenu(flags, make(chan *dbus.Call, 1)).Done)
}

// method FocusableChild

func (v *application) GoFocusableChild(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FocusableChild", flags, ch)
}

func (*application) StoreFocusableChild(call *dbus.Call) (path string, err error) {
	err = call.Store(&path)
	return
}

func (v *application) FocusableChild(flags dbus.Flags) (path string, err error) {
	return v.StoreFocusableChild(
		<-v.GoFocusableChild(flags, make(chan *dbus.Call, 1)).Done)
}

// signal WindowRemoved

func (v *application) ConnectWindowRemoved(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WindowRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WindowRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WindowAdded

func (v *application) ConnectWindowAdded(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WindowAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WindowAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path string
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SupportedMimeTypesChanged

func (v *application) ConnectSupportedMimeTypesChanged(cb func(dnd_mimes []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SupportedMimeTypesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SupportedMimeTypesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var dnd_mimes []string
		err := dbus.Store(sig.Body, &dnd_mimes)
		if err == nil {
			cb(dnd_mimes)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DesktopFileUpdated

func (v *application) ConnectDesktopFileUpdated(cb func(desktop_file string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DesktopFileUpdated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DesktopFileUpdated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var desktop_file string
		err := dbus.Store(sig.Body, &desktop_file)
		if err == nil {
			cb(desktop_file)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Window struct {
	view   // interface org.ayatana.bamf.view
	window // interface org.ayatana.bamf.window
	proxy.Object
}

func NewWindow(conn *dbus.Conn, path dbus.ObjectPath) (*Window, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Window)
	obj.Object.Init_(conn, "org.ayatana.bamf", path)
	return obj, nil
}

type window struct{}

func (v *window) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*window) GetInterfaceName_() string {
	return "org.ayatana.bamf.window"
}

// method GetXid

func (v *window) GoGetXid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetXid", flags, ch)
}

func (*window) StoreGetXid(call *dbus.Call) (xid uint32, err error) {
	err = call.Store(&xid)
	return
}

func (v *window) GetXid(flags dbus.Flags) (xid uint32, err error) {
	return v.StoreGetXid(
		<-v.GoGetXid(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetPid

func (v *window) GoGetPid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetPid", flags, ch)
}

func (*window) StoreGetPid(call *dbus.Call) (pid uint32, err error) {
	err = call.Store(&pid)
	return
}

func (v *window) GetPid(flags dbus.Flags) (pid uint32, err error) {
	return v.StoreGetPid(
		<-v.GoGetPid(flags, make(chan *dbus.Call, 1)).Done)
}

// method Transient

func (v *window) GoTransient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Transient", flags, ch)
}

func (*window) StoreTransient(call *dbus.Call) (path string, err error) {
	err = call.Store(&path)
	return
}

func (v *window) Transient(flags dbus.Flags) (path string, err error) {
	return v.StoreTransient(
		<-v.GoTransient(flags, make(chan *dbus.Call, 1)).Done)
}

// method WindowType

func (v *window) GoWindowType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WindowType", flags, ch)
}

func (*window) StoreWindowType(call *dbus.Call) (type0 uint32, err error) {
	err = call.Store(&type0)
	return
}

func (v *window) WindowType(flags dbus.Flags) (type0 uint32, err error) {
	return v.StoreWindowType(
		<-v.GoWindowType(flags, make(chan *dbus.Call, 1)).Done)
}

// method Xprop

func (v *window) GoXprop(flags dbus.Flags, ch chan *dbus.Call, xprop string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Xprop", flags, ch, xprop)
}

func (*window) StoreXprop(call *dbus.Call) (name string, err error) {
	err = call.Store(&name)
	return
}

func (v *window) Xprop(flags dbus.Flags, xprop string) (name string, err error) {
	return v.StoreXprop(
		<-v.GoXprop(flags, make(chan *dbus.Call, 1), xprop).Done)
}

// method Monitor

func (v *window) GoMonitor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Monitor", flags, ch)
}

func (*window) StoreMonitor(call *dbus.Call) (monitor_number int32, err error) {
	err = call.Store(&monitor_number)
	return
}

func (v *window) Monitor(flags dbus.Flags) (monitor_number int32, err error) {
	return v.StoreMonitor(
		<-v.GoMonitor(flags, make(chan *dbus.Call, 1)).Done)
}

// method Maximized

func (v *window) GoMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Maximized", flags, ch)
}

func (*window) StoreMaximized(call *dbus.Call) (maximized int32, err error) {
	err = call.Store(&maximized)
	return
}

func (v *window) Maximized(flags dbus.Flags) (maximized int32, err error) {
	return v.StoreMaximized(
		<-v.GoMaximized(flags, make(chan *dbus.Call, 1)).Done)
}

// signal MonitorChanged

func (v *window) ConnectMonitorChanged(cb func(old int32, new int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "MonitorChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".MonitorChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old int32
		var new int32
		err := dbus.Store(sig.Body, &old, &new)
		if err == nil {
			cb(old, new)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal MaximizedChanged

func (v *window) ConnectMaximizedChanged(cb func(old int32, new int32)) (dbusutil.SignalHandlerId, error) {
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
		var old int32
		var new int32
		err := dbus.Store(sig.Body, &old, &new)
		if err == nil {
			cb(old, new)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
