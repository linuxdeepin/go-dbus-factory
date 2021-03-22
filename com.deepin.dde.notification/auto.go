package notification

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Notification struct {
	notification // interface com.deepin.dde.Notification
	proxy.Object
}

func NewNotification(conn *dbus.Conn) *Notification {
	obj := new(Notification)
	obj.Object.Init_(conn, "com.deepin.dde.osd", "/com/deepin/dde/Notification")
	return obj
}

type notification struct{}

func (v *notification) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*notification) GetInterfaceName_() string {
	return "com.deepin.dde.Notification"
}

// method CloseNotification

func (v *notification) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseNotification", flags, ch, arg0)
}

func (v *notification) CloseNotification(flags dbus.Flags, arg0 uint32) error {
	return (<-v.GoCloseNotification(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetCapbilities

func (v *notification) GoGetCapbilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapbilities", flags, ch)
}

func (*notification) StoreGetCapbilities(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notification) GetCapbilities(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreGetCapbilities(
		<-v.GoGetCapbilities(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetServerInformation

func (v *notification) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetServerInformation", flags, ch)
}

func (*notification) StoreGetServerInformation(call *dbus.Call) (arg0 string, arg1 string, arg2 string, arg3 string, err error) {
	err = call.Store(&arg0, &arg1, &arg2, &arg3)
	return
}

func (v *notification) GetServerInformation(flags dbus.Flags) (arg0 string, arg1 string, arg2 string, arg3 string, err error) {
	return v.StoreGetServerInformation(
		<-v.GoGetServerInformation(flags, make(chan *dbus.Call, 1)).Done)
}

// method Notify

func (v *notification) GoNotify(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Notify", flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (*notification) StoreNotify(call *dbus.Call) (arg8 uint32, err error) {
	err = call.Store(&arg8)
	return
}

func (v *notification) Notify(flags dbus.Flags, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) (arg8 uint32, err error) {
	return v.StoreNotify(
		<-v.GoNotify(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7).Done)
}

// method GetAllRecords

func (v *notification) GoGetAllRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllRecords", flags, ch)
}

func (*notification) StoreGetAllRecords(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notification) GetAllRecords(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetAllRecords(
		<-v.GoGetAllRecords(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetRecordById

func (v *notification) GoGetRecordById(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetRecordById", flags, ch, arg0)
}

func (*notification) StoreGetRecordById(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *notification) GetRecordById(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetRecordById(
		<-v.GoGetRecordById(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetRecordsFromId

func (v *notification) GoGetRecordsFromId(flags dbus.Flags, ch chan *dbus.Call, arg0 int32, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetRecordsFromId", flags, ch, arg0, arg1)
}

func (*notification) StoreGetRecordsFromId(call *dbus.Call) (arg2 string, err error) {
	err = call.Store(&arg2)
	return
}

func (v *notification) GetRecordsFromId(flags dbus.Flags, arg0 int32, arg1 string) (arg2 string, err error) {
	return v.StoreGetRecordsFromId(
		<-v.GoGetRecordsFromId(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method RemoveRecord

func (v *notification) GoRemoveRecord(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveRecord", flags, ch, arg0)
}

func (v *notification) RemoveRecord(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRemoveRecord(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method ClearRecords

func (v *notification) GoClearRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearRecords", flags, ch)
}

func (v *notification) ClearRecords(flags dbus.Flags) error {
	return (<-v.GoClearRecords(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method getAppSetting

func (v *notification) GoGetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".getAppSetting", flags, ch, arg0)
}

func (*notification) StoreGetAppSetting(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *notification) GetAppSetting(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetAppSetting(
		<-v.GoGetAppSetting(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method Toggle

func (v *notification) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *notification) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *notification) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *notification) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Hide

func (v *notification) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *notification) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method recordCount

func (v *notification) GoRecordCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".recordCount", flags, ch)
}

func (*notification) StoreRecordCount(call *dbus.Call) (arg0 uint32, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notification) RecordCount(flags dbus.Flags) (arg0 uint32, err error) {
	return v.StoreRecordCount(
		<-v.GoRecordCount(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAppList

func (v *notification) GoGetAppList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAppList", flags, ch)
}

func (*notification) StoreGetAppList(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notification) GetAppList(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreGetAppList(
		<-v.GoGetAppList(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAppInfo

func (v *notification) GoGetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAppInfo", flags, ch, arg0, arg1)
}

func (*notification) StoreGetAppInfo(call *dbus.Call) (arg2 dbus.Variant, err error) {
	err = call.Store(&arg2)
	return
}

func (v *notification) GetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 dbus.Variant, err error) {
	return v.StoreGetAppInfo(
		<-v.GoGetAppInfo(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method GetSystemInfo

func (v *notification) GoGetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSystemInfo", flags, ch, arg0)
}

func (*notification) StoreGetSystemInfo(call *dbus.Call) (arg1 dbus.Variant, err error) {
	err = call.Store(&arg1)
	return
}

func (v *notification) GetSystemInfo(flags dbus.Flags, arg0 uint32) (arg1 dbus.Variant, err error) {
	return v.StoreGetSystemInfo(
		<-v.GoGetSystemInfo(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method SetAppInfo

func (v *notification) GoSetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAppInfo", flags, ch, arg0, arg1, arg2)
}

func (v *notification) SetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32, arg2 dbus.Variant) error {
	return (<-v.GoSetAppInfo(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetSystemInfo

func (v *notification) GoSetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32, arg1 dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSystemInfo", flags, ch, arg0, arg1)
}

func (v *notification) SetSystemInfo(flags dbus.Flags, arg0 uint32, arg1 dbus.Variant) error {
	return (<-v.GoSetSystemInfo(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method setAppSetting

func (v *notification) GoSetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".setAppSetting", flags, ch, arg0)
}

func (v *notification) SetAppSetting(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetAppSetting(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// signal NotificationClosed

func (v *notification) ConnectNotificationClosed(cb func(arg0 uint32, arg1 uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NotificationClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NotificationClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 uint32
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActionInvoked

func (v *notification) ConnectActionInvoked(cb func(arg0 uint32, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActionInvoked", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActionInvoked",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal RecordAdded

func (v *notification) ConnectRecordAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "RecordAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".RecordAdded",
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

// signal AppInfoChanged

func (v *notification) ConnectAppInfoChanged(cb func(arg0 string, arg1 uint32, arg2 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppInfoChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppInfoChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 uint32
		var arg2 dbus.Variant
		err := dbus.Store(sig.Body, &arg0, &arg1, &arg2)
		if err == nil {
			cb(arg0, arg1, arg2)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SystemInfoChanged

func (v *notification) ConnectSystemInfoChanged(cb func(arg0 uint32, arg1 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SystemInfoChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SystemInfoChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 dbus.Variant
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AppAddedSignal

func (v *notification) ConnectAppAddedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppAddedSignal", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppAddedSignal",
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

// signal AppRemovedSignal

func (v *notification) ConnectAppRemovedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppRemovedSignal", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppRemovedSignal",
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

// signal appRemoved

func (v *notification) ConnectAppRemoved(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appRemoved",
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

// signal appAdded

func (v *notification) ConnectAppAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appAdded",
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

// signal appSettingChanged

func (v *notification) ConnectAppSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appSettingChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appSettingChanged",
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

// signal systemSettingChanged

func (v *notification) ConnectSystemSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "systemSettingChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".systemSettingChanged",
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

// property allSetting s

func (v *notification) AllSetting() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "allSetting",
	}
}

// property systemSetting s

func (v *notification) SystemSetting() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "systemSetting",
	}
}
