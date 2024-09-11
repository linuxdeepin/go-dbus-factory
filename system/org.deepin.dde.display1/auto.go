// Code generated by "./generator ./system/org.deepin.dde.display1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package display1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Display interface {
	display // interface org.deepin.dde.Display1
	proxy.Object
}

type objectDisplay struct {
	interfaceDisplay // interface org.deepin.dde.Display1
	proxy.ImplObject
}

func NewDisplay(conn *dbus.Conn) Display {
	obj := new(objectDisplay)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Display1", "/org/deepin/dde/Display1")
	return obj
}

type display interface {
	GoGetConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetConfig(flags dbus.Flags) (string, error)
	GoSetConfig(flags dbus.Flags, ch chan *dbus.Call, cfgStr string) *dbus.Call
	SetConfig(flags dbus.Flags, cfgStr string) error
	ConnectConfigUpdated(cb func(updateAt string)) (dbusutil.SignalHandlerId, error)
}

type interfaceDisplay struct{}

func (v *interfaceDisplay) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDisplay) GetInterfaceName_() string {
	return "org.deepin.dde.Display1"
}

// method GetConfig

func (v *interfaceDisplay) GoGetConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConfig", flags, ch)
}

func (*interfaceDisplay) StoreGetConfig(call *dbus.Call) (cfgStr string, err error) {
	err = call.Store(&cfgStr)
	return
}

func (v *interfaceDisplay) GetConfig(flags dbus.Flags) (string, error) {
	return v.StoreGetConfig(
		<-v.GoGetConfig(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetConfig

func (v *interfaceDisplay) GoSetConfig(flags dbus.Flags, ch chan *dbus.Call, cfgStr string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetConfig", flags, ch, cfgStr)
}

func (v *interfaceDisplay) SetConfig(flags dbus.Flags, cfgStr string) error {
	return (<-v.GoSetConfig(flags, make(chan *dbus.Call, 1), cfgStr).Done).Err
}

// signal ConfigUpdated

func (v *interfaceDisplay) ConnectConfigUpdated(cb func(updateAt string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ConfigUpdated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ConfigUpdated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var updateAt string
		err := dbus.Store(sig.Body, &updateAt)
		if err == nil {
			cb(updateAt)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
