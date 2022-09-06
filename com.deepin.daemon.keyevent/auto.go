// Code generated by "./generator ./com.deepin.daemon.keyevent"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package keyevent

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type KeyEvent interface {
	keyEvent // interface com.deepin.daemon.KeyEvent
	proxy.Object
}

type objectKeyEvent struct {
	interfaceKeyEvent // interface com.deepin.daemon.KeyEvent
	proxy.ImplObject
}

func NewKeyEvent(conn *dbus.Conn) KeyEvent {
	obj := new(objectKeyEvent)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.KeyEvent", "/com/deepin/daemon/KeyEvent")
	return obj
}

type keyEvent interface {
	ConnectKeyEvent(cb func(keycode uint32, pressed bool, ctrlPressed bool, shiftPressed bool, altPressed bool, superPressed bool)) (dbusutil.SignalHandlerId, error)
}

type interfaceKeyEvent struct{}

func (v *interfaceKeyEvent) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceKeyEvent) GetInterfaceName_() string {
	return "com.deepin.daemon.KeyEvent"
}

// signal KeyEvent

func (v *interfaceKeyEvent) ConnectKeyEvent(cb func(keycode uint32, pressed bool, ctrlPressed bool, shiftPressed bool, altPressed bool, superPressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeyEvent", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeyEvent",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keycode uint32
		var pressed bool
		var ctrlPressed bool
		var shiftPressed bool
		var altPressed bool
		var superPressed bool
		err := dbus.Store(sig.Body, &keycode, &pressed, &ctrlPressed, &shiftPressed, &altPressed, &superPressed)
		if err == nil {
			cb(keycode, pressed, ctrlPressed, shiftPressed, altPressed, superPressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
