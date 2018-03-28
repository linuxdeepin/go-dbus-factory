package xeventmonitor

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

type XEventMonitor struct {
	xEventMonitor // interface com.deepin.api.XEventMonitor
	proxy.Object
}

func NewXEventMonitor(conn *dbus.Conn) *XEventMonitor {
	obj := new(XEventMonitor)
	obj.Object.Init_(conn, "com.deepin.api.XEventMonitor", "/com/deepin/api/XEventMonitor")
	return obj
}

type xEventMonitor struct{}

func (v *xEventMonitor) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*xEventMonitor) GetInterfaceName_() string {
	return "com.deepin.api.XEventMonitor"
}

// method RegisterArea

func (v *xEventMonitor) GoRegisterArea(flags dbus.Flags, ch chan *dbus.Call, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterArea", flags, ch, x1, y1, x2, y2, flag)
}

func (*xEventMonitor) StoreRegisterArea(call *dbus.Call) (id string, err error) {
	err = call.Store(&id)
	return
}

func (v *xEventMonitor) RegisterArea(flags dbus.Flags, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) (id string, err error) {
	return v.StoreRegisterArea(
		<-v.GoRegisterArea(flags, make(chan *dbus.Call, 1), x1, y1, x2, y2, flag).Done)
}

// method RegisterAreas

func (v *xEventMonitor) GoRegisterAreas(flags dbus.Flags, ch chan *dbus.Call, areas []CoordinateRange, flag int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAreas", flags, ch, areas, flag)
}

func (*xEventMonitor) StoreRegisterAreas(call *dbus.Call) (id string, err error) {
	err = call.Store(&id)
	return
}

func (v *xEventMonitor) RegisterAreas(flags dbus.Flags, areas []CoordinateRange, flag int32) (id string, err error) {
	return v.StoreRegisterAreas(
		<-v.GoRegisterAreas(flags, make(chan *dbus.Call, 1), areas, flag).Done)
}

// method RegisterFullScreen

func (v *xEventMonitor) GoRegisterFullScreen(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterFullScreen", flags, ch)
}

func (*xEventMonitor) StoreRegisterFullScreen(call *dbus.Call) (id string, err error) {
	err = call.Store(&id)
	return
}

func (v *xEventMonitor) RegisterFullScreen(flags dbus.Flags) (id string, err error) {
	return v.StoreRegisterFullScreen(
		<-v.GoRegisterFullScreen(flags, make(chan *dbus.Call, 1)).Done)
}

// method UnregisterArea

func (v *xEventMonitor) GoUnregisterArea(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterArea", flags, ch, id)
}

func (*xEventMonitor) StoreUnregisterArea(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *xEventMonitor) UnregisterArea(flags dbus.Flags, id string) (ok bool, err error) {
	return v.StoreUnregisterArea(
		<-v.GoUnregisterArea(flags, make(chan *dbus.Call, 1), id).Done)
}

// signal CancelAllArea

func (v *xEventMonitor) ConnectCancelAllArea(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CancelAllArea", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CancelAllArea",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CursorInto

func (v *xEventMonitor) ConnectCursorInto(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CursorInto", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CursorInto",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &x, &y, &id)
		if err == nil {
			cb(x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CursorOut

func (v *xEventMonitor) ConnectCursorOut(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CursorOut", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CursorOut",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &x, &y, &id)
		if err == nil {
			cb(x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CursorMove

func (v *xEventMonitor) ConnectCursorMove(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CursorMove", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CursorMove",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &x, &y, &id)
		if err == nil {
			cb(x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ButtonPress

func (v *xEventMonitor) ConnectButtonPress(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ButtonPress", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ButtonPress",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var button int32
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &button, &x, &y, &id)
		if err == nil {
			cb(button, x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ButtonRelease

func (v *xEventMonitor) ConnectButtonRelease(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ButtonRelease", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ButtonRelease",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var button int32
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &button, &x, &y, &id)
		if err == nil {
			cb(button, x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal KeyPress

func (v *xEventMonitor) ConnectKeyPress(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeyPress", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeyPress",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var key string
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &key, &x, &y, &id)
		if err == nil {
			cb(key, x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal KeyRelease

func (v *xEventMonitor) ConnectKeyRelease(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeyRelease", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeyRelease",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var key string
		var x int32
		var y int32
		var id string
		err := dbus.Store(sig.Body, &key, &x, &y, &id)
		if err == nil {
			cb(key, x, y, id)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
