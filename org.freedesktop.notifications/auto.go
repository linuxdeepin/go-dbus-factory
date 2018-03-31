package notifications

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

type Notifications struct {
	notifications // interface org.freedesktop.Notifications
	proxy.Object
}

func NewNotifications(conn *dbus.Conn) *Notifications {
	obj := new(Notifications)
	obj.Object.Init_(conn, "org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	return obj
}

type notifications struct{}

func (v *notifications) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*notifications) GetInterfaceName_() string {
	return "org.freedesktop.Notifications"
}

// method CloseNotification

func (v *notifications) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, in0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseNotification", flags, ch, in0)
}

func (v *notifications) CloseNotification(flags dbus.Flags, in0 uint32) error {
	return (<-v.GoCloseNotification(flags, make(chan *dbus.Call, 1), in0).Done).Err
}

// method GetCapbilities

func (v *notifications) GoGetCapbilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapbilities", flags, ch)
}

func (*notifications) StoreGetCapbilities(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notifications) GetCapbilities(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreGetCapbilities(
		<-v.GoGetCapbilities(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetServerInformation

func (v *notifications) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetServerInformation", flags, ch)
}

func (*notifications) StoreGetServerInformation(call *dbus.Call) (arg0 string, arg1 string, arg2 string, arg3 string, err error) {
	err = call.Store(&arg0, &arg1, &arg2, &arg3)
	return
}

func (v *notifications) GetServerInformation(flags dbus.Flags) (arg0 string, arg1 string, arg2 string, arg3 string, err error) {
	return v.StoreGetServerInformation(
		<-v.GoGetServerInformation(flags, make(chan *dbus.Call, 1)).Done)
}

// method Notify

func (v *notifications) GoNotify(flags dbus.Flags, ch chan *dbus.Call, arg1 string, arg2 uint32, arg3 string, arg4 string, arg5 string, arg6 []string, arg7 map[string]dbus.Variant, arg8 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Notify", flags, ch, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

func (*notifications) StoreNotify(call *dbus.Call) (arg0 uint32, err error) {
	err = call.Store(&arg0)
	return
}

func (v *notifications) Notify(flags dbus.Flags, arg1 string, arg2 uint32, arg3 string, arg4 string, arg5 string, arg6 []string, arg7 map[string]dbus.Variant, arg8 int32) (arg0 uint32, err error) {
	return v.StoreNotify(
		<-v.GoNotify(flags, make(chan *dbus.Call, 1), arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8).Done)
}

// signal NotificationClosed

func (v *notifications) ConnectNotificationClosed(cb func(in0 uint32, in1 uint32)) (dbusutil.SignalHandlerId, error) {
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
		var in0 uint32
		var in1 uint32
		err := dbus.Store(sig.Body, &in0, &in1)
		if err == nil {
			cb(in0, in1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
