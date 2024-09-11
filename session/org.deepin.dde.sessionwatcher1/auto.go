// Code generated by "./generator ./session/org.deepin.dde.sessionwatcher1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package sessionwatcher1

import (
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type SessionWatcher interface {
	sessionWatcher // interface org.deepin.dde.SessionWatcher1
	proxy.Object
}

type objectSessionWatcher struct {
	interfaceSessionWatcher // interface org.deepin.dde.SessionWatcher1
	proxy.ImplObject
}

func NewSessionWatcher(conn *dbus.Conn) SessionWatcher {
	obj := new(objectSessionWatcher)
	obj.ImplObject.Init_(conn, "org.deepin.dde.SessionWatcher1", "/org/deepin/dde/SessionWatcher1")
	return obj
}

type sessionWatcher interface {
	GoGetSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetSessions(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoIsX11SessionActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IsX11SessionActive(flags dbus.Flags) (bool, error)
	IsActive() proxy.PropBool
}

type interfaceSessionWatcher struct{}

func (v *interfaceSessionWatcher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSessionWatcher) GetInterfaceName_() string {
	return "org.deepin.dde.SessionWatcher1"
}

// method GetSessions

func (v *interfaceSessionWatcher) GoGetSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSessions", flags, ch)
}

func (*interfaceSessionWatcher) StoreGetSessions(call *dbus.Call) (sessions []dbus.ObjectPath, err error) {
	err = call.Store(&sessions)
	return
}

func (v *interfaceSessionWatcher) GetSessions(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetSessions(
		<-v.GoGetSessions(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsX11SessionActive

func (v *interfaceSessionWatcher) GoIsX11SessionActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsX11SessionActive", flags, ch)
}

func (*interfaceSessionWatcher) StoreIsX11SessionActive(call *dbus.Call) (is_active bool, err error) {
	err = call.Store(&is_active)
	return
}

func (v *interfaceSessionWatcher) IsX11SessionActive(flags dbus.Flags) (bool, error) {
	return v.StoreIsX11SessionActive(
		<-v.GoIsX11SessionActive(flags, make(chan *dbus.Call, 1)).Done)
}

// property IsActive b

func (v *interfaceSessionWatcher) IsActive() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsActive",
	}
}
