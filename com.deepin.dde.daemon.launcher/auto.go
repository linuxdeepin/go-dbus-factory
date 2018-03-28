package launcher

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

type Launcher struct {
	launcher // interface com.deepin.dde.daemon.Launcher
	proxy.Object
}

func NewLauncher(conn *dbus.Conn) *Launcher {
	obj := new(Launcher)
	obj.Object.Init_(conn, "com.deepin.dde.daemon.Launcher", "/com/deepin/dde/daemon/Launcher")
	return obj
}

type launcher struct{}

func (v *launcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*launcher) GetInterfaceName_() string {
	return "com.deepin.dde.daemon.Launcher"
}

// method GetAllItemInfos

func (v *launcher) GoGetAllItemInfos(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllItemInfos", flags, ch)
}

func (*launcher) StoreGetAllItemInfos(call *dbus.Call) (itemInfoList []ItemInfo, err error) {
	err = call.Store(&itemInfoList)
	return
}

func (v *launcher) GetAllItemInfos(flags dbus.Flags) (itemInfoList []ItemInfo, err error) {
	return v.StoreGetAllItemInfos(
		<-v.GoGetAllItemInfos(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAllNewInstalledApps

func (v *launcher) GoGetAllNewInstalledApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllNewInstalledApps", flags, ch)
}

func (*launcher) StoreGetAllNewInstalledApps(call *dbus.Call) (apps []string, err error) {
	err = call.Store(&apps)
	return
}

func (v *launcher) GetAllNewInstalledApps(flags dbus.Flags) (apps []string, err error) {
	return v.StoreGetAllNewInstalledApps(
		<-v.GoGetAllNewInstalledApps(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDisableScaling

func (v *launcher) GoGetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDisableScaling", flags, ch, id)
}

func (*launcher) StoreGetDisableScaling(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *launcher) GetDisableScaling(flags dbus.Flags, id string) (value bool, err error) {
	return v.StoreGetDisableScaling(
		<-v.GoGetDisableScaling(flags, make(chan *dbus.Call, 1), id).Done)
}

// method GetItemInfo

func (v *launcher) GoGetItemInfo(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetItemInfo", flags, ch, id)
}

func (*launcher) StoreGetItemInfo(call *dbus.Call) (itemInfo ItemInfo, err error) {
	err = call.Store(&itemInfo)
	return
}

func (v *launcher) GetItemInfo(flags dbus.Flags, id string) (itemInfo ItemInfo, err error) {
	return v.StoreGetItemInfo(
		<-v.GoGetItemInfo(flags, make(chan *dbus.Call, 1), id).Done)
}

// method GetUseProxy

func (v *launcher) GoGetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUseProxy", flags, ch, id)
}

func (*launcher) StoreGetUseProxy(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *launcher) GetUseProxy(flags dbus.Flags, id string) (value bool, err error) {
	return v.StoreGetUseProxy(
		<-v.GoGetUseProxy(flags, make(chan *dbus.Call, 1), id).Done)
}

// method IsItemOnDesktop

func (v *launcher) GoIsItemOnDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsItemOnDesktop", flags, ch, id)
}

func (*launcher) StoreIsItemOnDesktop(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *launcher) IsItemOnDesktop(flags dbus.Flags, id string) (result bool, err error) {
	return v.StoreIsItemOnDesktop(
		<-v.GoIsItemOnDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method MarkLaunched

func (v *launcher) GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MarkLaunched", flags, ch, id)
}

func (v *launcher) MarkLaunched(flags dbus.Flags, id string) error {
	return (<-v.GoMarkLaunched(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method RequestRemoveFromDesktop

func (v *launcher) GoRequestRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestRemoveFromDesktop", flags, ch, id)
}

func (*launcher) StoreRequestRemoveFromDesktop(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *launcher) RequestRemoveFromDesktop(flags dbus.Flags, id string) (ok bool, err error) {
	return v.StoreRequestRemoveFromDesktop(
		<-v.GoRequestRemoveFromDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method RequestSendToDesktop

func (v *launcher) GoRequestSendToDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSendToDesktop", flags, ch, id)
}

func (*launcher) StoreRequestSendToDesktop(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *launcher) RequestSendToDesktop(flags dbus.Flags, id string) (ok bool, err error) {
	return v.StoreRequestSendToDesktop(
		<-v.GoRequestSendToDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method RequestUninstall

func (v *launcher) GoRequestUninstall(flags dbus.Flags, ch chan *dbus.Call, id string, purge bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestUninstall", flags, ch, id, purge)
}

func (v *launcher) RequestUninstall(flags dbus.Flags, id string, purge bool) error {
	return (<-v.GoRequestUninstall(flags, make(chan *dbus.Call, 1), id, purge).Done).Err
}

// method Search

func (v *launcher) GoSearch(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Search", flags, ch, key)
}

func (v *launcher) Search(flags dbus.Flags, key string) error {
	return (<-v.GoSearch(flags, make(chan *dbus.Call, 1), key).Done).Err
}

// method SetDisableScaling

func (v *launcher) GoSetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDisableScaling", flags, ch, id, value)
}

func (v *launcher) SetDisableScaling(flags dbus.Flags, id string, value bool) error {
	return (<-v.GoSetDisableScaling(flags, make(chan *dbus.Call, 1), id, value).Done).Err
}

// method SetUseProxy

func (v *launcher) GoSetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUseProxy", flags, ch, id, value)
}

func (v *launcher) SetUseProxy(flags dbus.Flags, id string, value bool) error {
	return (<-v.GoSetUseProxy(flags, make(chan *dbus.Call, 1), id, value).Done).Err
}

// signal SearchDone

func (v *launcher) ConnectSearchDone(cb func(apps []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SearchDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SearchDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var apps []string
		err := dbus.Store(sig.Body, &apps)
		if err == nil {
			cb(apps)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemChanged

func (v *launcher) ConnectItemChanged(cb func(status string, itemInfo ItemInfo, categoryID int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var itemInfo ItemInfo
		var categoryID int64
		err := dbus.Store(sig.Body, &status, &itemInfo, &categoryID)
		if err == nil {
			cb(status, itemInfo, categoryID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NewAppLaunched

func (v *launcher) ConnectNewAppLaunched(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NewAppLaunched", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NewAppLaunched",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appID string
		err := dbus.Store(sig.Body, &appID)
		if err == nil {
			cb(appID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UninstallSuccess

func (v *launcher) ConnectUninstallSuccess(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UninstallSuccess", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UninstallSuccess",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appID string
		err := dbus.Store(sig.Body, &appID)
		if err == nil {
			cb(appID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UninstallFailed

func (v *launcher) ConnectUninstallFailed(cb func(appId string, errMsg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UninstallFailed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UninstallFailed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appId string
		var errMsg string
		err := dbus.Store(sig.Body, &appId, &errMsg)
		if err == nil {
			cb(appId, errMsg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Fullscreen b

func (v *launcher) Fullscreen() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Fullscreen",
	}
}

// property DisplayMode i

func (v *launcher) DisplayMode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DisplayMode",
	}
}
