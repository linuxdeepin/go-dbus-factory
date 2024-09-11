// Code generated by "./generator ./system/org.deepin.dde.gesture1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package gesture1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Gesture interface {
	gesture // interface org.deepin.dde.Gesture1
	proxy.Object
}

type objectGesture struct {
	interfaceGesture // interface org.deepin.dde.Gesture1
	proxy.ImplObject
}

func NewGesture(conn *dbus.Conn) Gesture {
	obj := new(objectGesture)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Gesture1", "/org/deepin/dde/Gesture1")
	return obj
}

type gesture interface {
	GoSetShortPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call
	SetShortPressDuration(flags dbus.Flags, duration uint32) error
	GoSetEdgeMoveStopDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call
	SetEdgeMoveStopDuration(flags dbus.Flags, duration uint32) error
	GoSetInputIgnore(flags dbus.Flags, ch chan *dbus.Call, node string, ignore bool) *dbus.Call
	SetInputIgnore(flags dbus.Flags, node string, ignore bool) error
	ConnectEvent(cb func(name string, direction string, fingers int32)) (dbusutil.SignalHandlerId, error)
	ConnectDbclickDown(cb func(fingers int32)) (dbusutil.SignalHandlerId, error)
	ConnectSwipeMoving(cb func(fingers int32, accelX float64, accely float64)) (dbusutil.SignalHandlerId, error)
	ConnectSwipeStop(cb func(fingers int32)) (dbusutil.SignalHandlerId, error)
	ConnectTouchEdgeEvent(cb func(direction string, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error)
	ConnectTouchSinglePressTimeout(cb func(time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error)
	ConnectTouchPressTimeout(cb func(fingers int32, time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error)
	ConnectTouchUpOrCancel(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error)
	ConnectTouchEdgeMoveStop(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error)
	ConnectTouchEdgeMoveStopLeave(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error)
	ConnectTouchMoving(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error)
	ConnectTouchMovementEvent(cb func(duration string, fingers int32, startScalex float64, startScaley float64, endScalex float64, endScaley float64)) (dbusutil.SignalHandlerId, error)
}

type interfaceGesture struct{}

func (v *interfaceGesture) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceGesture) GetInterfaceName_() string {
	return "org.deepin.dde.Gesture1"
}

// method SetShortPressDuration

func (v *interfaceGesture) GoSetShortPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortPressDuration", flags, ch, duration)
}

func (v *interfaceGesture) SetShortPressDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetShortPressDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

// method SetEdgeMoveStopDuration

func (v *interfaceGesture) GoSetEdgeMoveStopDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetEdgeMoveStopDuration", flags, ch, duration)
}

func (v *interfaceGesture) SetEdgeMoveStopDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetEdgeMoveStopDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

// method SetInputIgnore

func (v *interfaceGesture) GoSetInputIgnore(flags dbus.Flags, ch chan *dbus.Call, node string, ignore bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetInputIgnore", flags, ch, node, ignore)
}

func (v *interfaceGesture) SetInputIgnore(flags dbus.Flags, node string, ignore bool) error {
	return (<-v.GoSetInputIgnore(flags, make(chan *dbus.Call, 1), node, ignore).Done).Err
}

// signal Event

func (v *interfaceGesture) ConnectEvent(cb func(name string, direction string, fingers int32)) (dbusutil.SignalHandlerId, error) {
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

// signal DbclickDown

func (v *interfaceGesture) ConnectDbclickDown(cb func(fingers int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DbclickDown", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DbclickDown",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var fingers int32
		err := dbus.Store(sig.Body, &fingers)
		if err == nil {
			cb(fingers)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SwipeMoving

func (v *interfaceGesture) ConnectSwipeMoving(cb func(fingers int32, accelX float64, accely float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SwipeMoving", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SwipeMoving",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var fingers int32
		var accelX float64
		var accely float64
		err := dbus.Store(sig.Body, &fingers, &accelX, &accely)
		if err == nil {
			cb(fingers, accelX, accely)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SwipeStop

func (v *interfaceGesture) ConnectSwipeStop(cb func(fingers int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SwipeStop", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SwipeStop",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var fingers int32
		err := dbus.Store(sig.Body, &fingers)
		if err == nil {
			cb(fingers)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchEdgeEvent

func (v *interfaceGesture) ConnectTouchEdgeEvent(cb func(direction string, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchSinglePressTimeout(cb func(time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchPressTimeout(cb func(fingers int32, time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchUpOrCancel(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchEdgeMoveStop(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchEdgeMoveStopLeave(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceGesture) ConnectTouchMoving(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
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

// signal TouchMovementEvent

func (v *interfaceGesture) ConnectTouchMovementEvent(cb func(duration string, fingers int32, startScalex float64, startScaley float64, endScalex float64, endScaley float64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchMovementEvent", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchMovementEvent",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var duration string
		var fingers int32
		var startScalex float64
		var startScaley float64
		var endScalex float64
		var endScaley float64
		err := dbus.Store(sig.Body, &duration, &fingers, &startScalex, &startScaley, &endScalex, &endScaley)
		if err == nil {
			cb(duration, fingers, startScalex, startScaley, endScalex, endScaley)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
