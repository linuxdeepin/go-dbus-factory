// Code generated by "./generator ./session/org.deepin.dde.bluetooth1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package bluetooth1

import (
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Bluetooth interface {
	bluetooth // interface org.deepin.dde.Bluetooth1
	proxy.Object
}

type objectBluetooth struct {
	interfaceBluetooth // interface org.deepin.dde.Bluetooth1
	proxy.ImplObject
}

func NewBluetooth(conn *dbus.Conn) Bluetooth {
	obj := new(objectBluetooth)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Bluetooth1", "/org/deepin/dde/Bluetooth1")
	return obj
}

type bluetooth interface {
	GoDisconnectAllDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DisconnectAllDevices(flags dbus.Flags) error
}

type interfaceBluetooth struct{}

func (v *interfaceBluetooth) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceBluetooth) GetInterfaceName_() string {
	return "org.deepin.dde.Bluetooth1"
}

// method DisconnectAllDevices

func (v *interfaceBluetooth) GoDisconnectAllDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectAllDevices", flags, ch)
}

func (v *interfaceBluetooth) DisconnectAllDevices(flags dbus.Flags) error {
	return (<-v.GoDisconnectAllDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}
