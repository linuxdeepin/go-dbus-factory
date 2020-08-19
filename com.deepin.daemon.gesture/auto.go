package gesture

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

type Gesture struct {
	gesture // interface com.deepin.daemon.Gesture
	proxy.Object
}

func NewGesture(conn *dbus.Conn) *Gesture {
	obj := new(Gesture)
	obj.Object.Init_(conn, "com.deepin.daemon.Gesture", "/com/deepin/daemon/Gesture")
	return obj
}

type gesture struct{}

func (v *gesture) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*gesture) GetInterfaceName_() string {
	return "com.deepin.daemon.Gesture"
}

// method SetShortPressDuration

func (v *gesture) GoSetShortPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortPressDuration", flags, ch, duration)
}

func (v *gesture) SetShortPressDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetShortPressDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

// method SetEdgeMoveStopDuration

func (v *gesture) GoSetEdgeMoveStopDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetEdgeMoveStopDuration", flags, ch, duration)
}

func (v *gesture) SetEdgeMoveStopDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetEdgeMoveStopDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

// signal Event

func (v *gesture) ConnectEvent(cb func(name string, direction string, fingers int32)) (dbusutil.SignalHandlerId, error) {
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
		var direction string
		var fingers int32
		err := dbus.Store(sig.Body, &name, &direction, &fingers)
		if err == nil {
			cb(name, direction, fingers)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchEdgeEvent

func (v *gesture) ConnectTouchEdgeEvent(cb func(direction string, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchEdgeEvent", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchEdgeEvent",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var direction string
		var scalex float64
		var scaley float64
		err := dbus.Store(sig.Body, &direction, &scalex, &scaley)
		if err == nil {
			cb(direction, scalex, scaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchSinglePressTimeout

func (v *gesture) ConnectTouchSinglePressTimeout(cb func(time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchSinglePressTimeout", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchSinglePressTimeout",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var time int32
		var scalex float64
		var scaley float64
		err := dbus.Store(sig.Body, &time, &scalex, &scaley)
		if err == nil {
			cb(time, scalex, scaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchPressTimeout

func (v *gesture) ConnectTouchPressTimeout(cb func(fingers int32, time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchPressTimeout", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchPressTimeout",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var fingers int32
		var time int32
		var scalex float64
		var scaley float64
		err := dbus.Store(sig.Body, &fingers, &time, &scalex, &scaley)
		if err == nil {
			cb(fingers, time, scalex, scaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchUpOrCancel

func (v *gesture) ConnectTouchUpOrCancel(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchUpOrCancel", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchUpOrCancel",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var scalex float64
		var scaley float64
		err := dbus.Store(sig.Body, &scalex, &scaley)
		if err == nil {
			cb(scalex, scaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchEdgeMoveStop

func (v *gesture) ConnectTouchEdgeMoveStop(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchEdgeMoveStop", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchEdgeMoveStop",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var direction string
		var scalex float64
		var scaley float64
		var duration int32
		err := dbus.Store(sig.Body, &direction, &scalex, &scaley, &duration)
		if err == nil {
			cb(direction, scalex, scaley, duration)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchEdgeMoveStopLeave

func (v *gesture) ConnectTouchEdgeMoveStopLeave(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchEdgeMoveStopLeave", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchEdgeMoveStopLeave",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var direction string
		var scalex float64
		var scaley float64
		var duration int32
		err := dbus.Store(sig.Body, &direction, &scalex, &scaley, &duration)
		if err == nil {
			cb(direction, scalex, scaley, duration)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchMoving

func (v *gesture) ConnectTouchMoving(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchMoving", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchMoving",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var scalex float64
		var scaley float64
		err := dbus.Store(sig.Body, &scalex, &scaley)
		if err == nil {
			cb(scalex, scaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
