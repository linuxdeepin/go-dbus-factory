package sessionmanager

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

func (*startManager) StoreAddAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) AddAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreAddAutostart(
		<-v.GoAddAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method AutostartList

func (v *startManager) GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AutostartList", flags, ch)
}

func (*startManager) StoreAutostartList(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) AutostartList(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreAutostartList(
		<-v.GoAutostartList(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetApps

func (v *startManager) GoGetApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetApps", flags, ch)
}

func (*startManager) StoreGetApps(call *dbus.Call) (arg0 map[uint32]string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *startManager) GetApps(flags dbus.Flags) (arg0 map[uint32]string, err error) {
	return v.StoreGetApps(
		<-v.GoGetApps(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsAutostart

func (v *startManager) GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsAutostart", flags, ch, arg0)
}

func (*startManager) StoreIsAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) IsAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreIsAutostart(
		<-v.GoIsAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method Launch

func (v *startManager) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Launch", flags, ch, arg0)
}

func (*startManager) StoreLaunch(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) Launch(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreLaunch(
		<-v.GoLaunch(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method LaunchApp

func (v *startManager) GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchApp", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) LaunchApp(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error {
	return (<-v.GoLaunchApp(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method LaunchAppAction

func (v *startManager) GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppAction", flags, ch, arg0, arg1, arg2)
}

func (v *startManager) LaunchAppAction(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error {
	return (<-v.GoLaunchAppAction(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method LaunchWithTimestamp

func (v *startManager) GoLaunchWithTimestamp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchWithTimestamp", flags, ch, arg0, arg1)
}

func (*startManager) StoreLaunchWithTimestamp(call *dbus.Call) (arg2 bool, err error) {
	err = call.Store(&arg2)
	return
}

func (v *startManager) LaunchWithTimestamp(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 bool, err error) {
	return v.StoreLaunchWithTimestamp(
		<-v.GoLaunchWithTimestamp(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method RemoveAutostart

func (v *startManager) GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAutostart", flags, ch, arg0)
}

func (*startManager) StoreRemoveAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *startManager) RemoveAutostart(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreRemoveAutostart(
		<-v.GoRemoveAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RunCommand

func (v *startManager) GoRunCommand(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RunCommand", flags, ch, arg0, arg1)
}

func (v *startManager) RunCommand(flags dbus.Flags, arg0 string, arg1 []string) error {
	return (<-v.GoRunCommand(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
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

// method CanHibernate

func (v *sessionManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*sessionManager) StoreCanHibernate(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanHibernate(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanLogout

func (v *sessionManager) GoCanLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanLogout", flags, ch)
}

func (*sessionManager) StoreCanLogout(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanLogout(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanLogout(
		<-v.GoCanLogout(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanReboot

func (v *sessionManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*sessionManager) StoreCanReboot(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanReboot(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanShutdown

func (v *sessionManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (*sessionManager) StoreCanShutdown(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanShutdown(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanShutdown(
		<-v.GoCanShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSuspend

func (v *sessionManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*sessionManager) StoreCanSuspend(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *sessionManager) CanSuspend(flags dbus.Flags) (arg0 bool, err error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

// method ForceLogout

func (v *sessionManager) GoForceLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceLogout", flags, ch)
}

func (v *sessionManager) ForceLogout(flags dbus.Flags) error {
	return (<-v.GoForceLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ForceReboot

func (v *sessionManager) GoForceReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceReboot", flags, ch)
}

func (v *sessionManager) ForceReboot(flags dbus.Flags) error {
	return (<-v.GoForceReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ForceShutdown

func (v *sessionManager) GoForceShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceShutdown", flags, ch)
}

func (v *sessionManager) ForceShutdown(flags dbus.Flags) error {
	return (<-v.GoForceShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Logout

func (v *sessionManager) GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *sessionManager) Logout(flags dbus.Flags) error {
	return (<-v.GoLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PowerOffChoose

func (v *sessionManager) GoPowerOffChoose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOffChoose", flags, ch)
}

func (v *sessionManager) PowerOffChoose(flags dbus.Flags) error {
	return (<-v.GoPowerOffChoose(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reboot

func (v *sessionManager) GoReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *sessionManager) Reboot(flags dbus.Flags) error {
	return (<-v.GoReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Register

func (v *sessionManager) GoRegister(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, arg0)
}

func (*sessionManager) StoreRegister(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *sessionManager) Register(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreRegister(
		<-v.GoRegister(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RequestHibernate

func (v *sessionManager) GoRequestHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestHibernate", flags, ch)
}

func (v *sessionManager) RequestHibernate(flags dbus.Flags) error {
	return (<-v.GoRequestHibernate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestLock

func (v *sessionManager) GoRequestLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLock", flags, ch)
}

func (v *sessionManager) RequestLock(flags dbus.Flags) error {
	return (<-v.GoRequestLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestLogout

func (v *sessionManager) GoRequestLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLogout", flags, ch)
}

func (v *sessionManager) RequestLogout(flags dbus.Flags) error {
	return (<-v.GoRequestLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestReboot

func (v *sessionManager) GoRequestReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestReboot", flags, ch)
}

func (v *sessionManager) RequestReboot(flags dbus.Flags) error {
	return (<-v.GoRequestReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestShutdown

func (v *sessionManager) GoRequestShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestShutdown", flags, ch)
}

func (v *sessionManager) RequestShutdown(flags dbus.Flags) error {
	return (<-v.GoRequestShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestSuspend

func (v *sessionManager) GoRequestSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSuspend", flags, ch)
}

func (v *sessionManager) RequestSuspend(flags dbus.Flags) error {
	return (<-v.GoRequestSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Shutdown

func (v *sessionManager) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *sessionManager) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ToggleDebug

func (v *sessionManager) GoToggleDebug(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleDebug", flags, ch)
}

func (v *sessionManager) ToggleDebug(flags dbus.Flags) error {
	return (<-v.GoToggleDebug(flags, make(chan *dbus.Call, 1)).Done).Err
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

func (*wmSwitcher) StoreCurrentWM(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *wmSwitcher) CurrentWM(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreCurrentWM(
		<-v.GoCurrentWM(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestSwitchWM

func (v *wmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *wmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	return (<-v.GoRequestSwitchWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RestartWM

func (v *wmSwitcher) GoRestartWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestartWM", flags, ch)
}

func (v *wmSwitcher) RestartWM(flags dbus.Flags) error {
	return (<-v.GoRestartWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Start2DWM

func (v *wmSwitcher) GoStart2DWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start2DWM", flags, ch)
}

func (v *wmSwitcher) Start2DWM(flags dbus.Flags) error {
	return (<-v.GoStart2DWM(flags, make(chan *dbus.Call, 1)).Done).Err
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

func (*xSettings) StoreGetColor(call *dbus.Call) (arg1 []int16, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetColor(flags dbus.Flags, arg0 string) (arg1 []int16, err error) {
	return v.StoreGetColor(
		<-v.GoGetColor(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetInteger

func (v *xSettings) GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetInteger", flags, ch, arg0)
}

func (*xSettings) StoreGetInteger(call *dbus.Call) (arg1 int32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetInteger(flags dbus.Flags, arg0 string) (arg1 int32, err error) {
	return v.StoreGetInteger(
		<-v.GoGetInteger(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetString

func (v *xSettings) GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetString", flags, ch, arg0)
}

func (*xSettings) StoreGetString(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *xSettings) GetString(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetString(
		<-v.GoGetString(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListProps

func (v *xSettings) GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListProps", flags, ch)
}

func (*xSettings) StoreListProps(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *xSettings) ListProps(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreListProps(
		<-v.GoListProps(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetColor

func (v *xSettings) GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetColor", flags, ch, arg0, arg1)
}

func (v *xSettings) SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error {
	return (<-v.GoSetColor(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetInteger

func (v *xSettings) GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetInteger", flags, ch, arg0, arg1)
}

func (v *xSettings) SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error {
	return (<-v.GoSetInteger(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetString

func (v *xSettings) GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetString", flags, ch, arg0, arg1)
}

func (v *xSettings) SetString(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoSetString(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

type Display struct {
	display // interface com.deepin.daemon.Display
	proxy.Object
}

func NewDisplay(conn *dbus.Conn) *Display {
	obj := new(Display)
	obj.Object.Init_(conn, "com.deepin.SessionManager", "/com/deepin/daemon/Display")
	return obj
}

type display struct{}

func (v *display) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*display) GetInterfaceName_() string {
	return "com.deepin.daemon.Display"
}

// method ApplyChanges

func (v *display) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplyChanges", flags, ch)
}

func (v *display) ApplyChanges(flags dbus.Flags) error {
	return (<-v.GoApplyChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AssociateTouch

func (v *display) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AssociateTouch", flags, ch, arg0, arg1)
}

func (v *display) AssociateTouch(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoAssociateTouch(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method ChangeBrightness

func (v *display) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeBrightness", flags, ch, arg0)
}

func (v *display) ChangeBrightness(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoChangeBrightness(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DeleteCustomMode

func (v *display) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCustomMode", flags, ch, arg0)
}

func (v *display) DeleteCustomMode(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDeleteCustomMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetBrightness

func (v *display) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch)
}

func (*display) StoreGetBrightness(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) GetBrightness(flags dbus.Flags) (arg0 map[string]float64, err error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputNames

func (v *display) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputNames", flags, ch)
}

func (*display) StoreListOutputNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) ListOutputNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListOutputNames(
		<-v.GoListOutputNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputsCommonModes

func (v *display) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputsCommonModes", flags, ch)
}

func (*display) StoreListOutputsCommonModes(call *dbus.Call) (arg0 []ModeInfo, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) ListOutputsCommonModes(flags dbus.Flags) (arg0 []ModeInfo, err error) {
	return v.StoreListOutputsCommonModes(
		<-v.GoListOutputsCommonModes(flags, make(chan *dbus.Call, 1)).Done)
}

// method ModifyConfigName

func (v *display) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ModifyConfigName", flags, ch, arg0, arg1)
}

func (v *display) ModifyConfigName(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoModifyConfigName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method RefreshBrightness

func (v *display) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBrightness", flags, ch)
}

func (v *display) RefreshBrightness(flags dbus.Flags) error {
	return (<-v.GoRefreshBrightness(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reset

func (v *display) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *display) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ResetChanges

func (v *display) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetChanges", flags, ch)
}

func (v *display) ResetChanges(flags dbus.Flags) error {
	return (<-v.GoResetChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *display) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *display) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAndSaveBrightness

func (v *display) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAndSaveBrightness", flags, ch, arg0, arg1)
}

func (v *display) SetAndSaveBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetAndSaveBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetBrightness

func (v *display) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, arg0, arg1)
}

func (v *display) SetBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPrimary

func (v *display) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrimary", flags, ch, arg0)
}

func (v *display) SetPrimary(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetPrimary(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SwitchMode

func (v *display) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint8, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchMode", flags, ch, arg0, arg1)
}

func (v *display) SwitchMode(flags dbus.Flags, arg0 uint8, arg1 string) error {
	return (<-v.GoSwitchMode(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// property HasChanged b

func (v *display) HasChanged() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasChanged",
	}
}

// property DisplayMode y

func (v *display) DisplayMode() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "DisplayMode",
	}
}

// property ScreenWidth q

func (v *display) ScreenWidth() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "ScreenWidth",
	}
}

// property ScreenHeight q

func (v *display) ScreenHeight() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "ScreenHeight",
	}
}

// property Primary s

func (v *display) Primary() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Primary",
	}
}

// property CurrentCustomId s

func (v *display) CurrentCustomId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CurrentCustomId",
	}
}

// property CustomIdList as

func (v *display) CustomIdList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "CustomIdList",
	}
}

// property PrimaryRect (nnqq)

func (v *display) PrimaryRect() PropDisplayPrimaryRect {
	return PropDisplayPrimaryRect{
		Impl: v,
	}
}

type PropDisplayPrimaryRect struct {
	Impl proxy.Implementer
}

func (p PropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"PrimaryRect", &value)
	return
}

func (p PropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v Rectangle
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, Rectangle{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"PrimaryRect", cb0)
}

// property Monitors ao

func (v *display) Monitors() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Monitors",
	}
}

// property Brightness a{sd}

func (v *display) Brightness() PropDisplayBrightness {
	return PropDisplayBrightness{
		Impl: v,
	}
}

type PropDisplayBrightness struct {
	Impl proxy.Implementer
}

func (p PropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Brightness", &value)
	return
}

func (p PropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]float64
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
		"Brightness", cb0)
}

// property TouchMap a{ss}

func (v *display) TouchMap() PropDisplayTouchMap {
	return PropDisplayTouchMap{
		Impl: v,
	}
}

type PropDisplayTouchMap struct {
	Impl proxy.Implementer
}

func (p PropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"TouchMap", &value)
	return
}

func (p PropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
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
		"TouchMap", cb0)
}

type Monitor struct {
	monitor // interface com.deepin.daemon.Display.Monitor
	proxy.Object
}

func NewMonitor(conn *dbus.Conn, path dbus.ObjectPath) (*Monitor, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Monitor)
	obj.Object.Init_(conn, "com.deepin.SessionManager", path)
	return obj, nil
}

type monitor struct{}

func (v *monitor) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*monitor) GetInterfaceName_() string {
	return "com.deepin.daemon.Display.Monitor"
}

// method Enable

func (v *monitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, arg0)
}

func (v *monitor) Enable(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetMode

func (v *monitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMode", flags, ch, arg0)
}

func (v *monitor) SetMode(flags dbus.Flags, arg0 uint32) error {
	return (<-v.GoSetMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetModeBySize

func (v *monitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16, arg1 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetModeBySize", flags, ch, arg0, arg1)
}

func (v *monitor) SetModeBySize(flags dbus.Flags, arg0 uint16, arg1 uint16) error {
	return (<-v.GoSetModeBySize(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPosition

func (v *monitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, arg0 int16, arg1 int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, arg0, arg1)
}

func (v *monitor) SetPosition(flags dbus.Flags, arg0 int16, arg1 int16) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetReflect

func (v *monitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetReflect", flags, ch, arg0)
}

func (v *monitor) SetReflect(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetReflect(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRefreshRate

func (v *monitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRefreshRate", flags, ch, arg0)
}

func (v *monitor) SetRefreshRate(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetRefreshRate(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRotation

func (v *monitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRotation", flags, ch, arg0)
}

func (v *monitor) SetRotation(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetRotation(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// property Name s

func (v *monitor) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Enabled b

func (v *monitor) Enabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property Connected b

func (v *monitor) Connected() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property X n

func (v *monitor) X() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "X",
	}
}

// property Y n

func (v *monitor) Y() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "Y",
	}
}

// property Width q

func (v *monitor) Width() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Width",
	}
}

// property Height q

func (v *monitor) Height() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Height",
	}
}

// property Rotation q

func (v *monitor) Rotation() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Rotation",
	}
}

// property Reflect q

func (v *monitor) Reflect() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Reflect",
	}
}

// property RefreshRate d

func (v *monitor) RefreshRate() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "RefreshRate",
	}
}

// property Rotations aq

func (v *monitor) Rotations() proxy.PropUint16Array {
	return proxy.PropUint16Array{
		Impl: v,
		Name: "Rotations",
	}
}

// property Reflects aq

func (v *monitor) Reflects() proxy.PropUint16Array {
	return proxy.PropUint16Array{
		Impl: v,
		Name: "Reflects",
	}
}

// property BestMode (uqqd)

func (v *monitor) BestMode() PropMonitorBestMode {
	return PropMonitorBestMode{
		Impl: v,
	}
}

type PropMonitorBestMode struct {
	Impl proxy.Implementer
}

func (p PropMonitorBestMode) Get(flags dbus.Flags) (value ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"BestMode", &value)
	return
}

func (p PropMonitorBestMode) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModeInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"BestMode", cb0)
}

// property CurrentMode (uqqd)

func (v *monitor) CurrentMode() PropMonitorCurrentMode {
	return PropMonitorCurrentMode{
		Impl: v,
	}
}

type PropMonitorCurrentMode struct {
	Impl proxy.Implementer
}

func (p PropMonitorCurrentMode) Get(flags dbus.Flags) (value ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"CurrentMode", &value)
	return
}

func (p PropMonitorCurrentMode) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModeInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"CurrentMode", cb0)
}

// property Modes a(uqqd)

func (v *monitor) Modes() PropMonitorModes {
	return PropMonitorModes{
		Impl: v,
	}
}

type PropMonitorModes struct {
	Impl proxy.Implementer
}

func (p PropMonitorModes) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Modes", &value)
	return
}

func (p PropMonitorModes) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModeInfo
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
		"Modes", cb0)
}

// property PreferredModes a(uqqd)

func (v *monitor) PreferredModes() PropMonitorPreferredModes {
	return PropMonitorPreferredModes{
		Impl: v,
	}
}

type PropMonitorPreferredModes struct {
	Impl proxy.Implementer
}

func (p PropMonitorPreferredModes) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"PreferredModes", &value)
	return
}

func (p PropMonitorPreferredModes) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModeInfo
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
		"PreferredModes", cb0)
}
