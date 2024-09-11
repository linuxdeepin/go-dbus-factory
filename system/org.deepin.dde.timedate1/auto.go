// Code generated by "./generator ./system/org.deepin.dde.timedate1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package timedate1

import (
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Timedate interface {
	timedate // interface org.deepin.dde.Timedate1
	proxy.Object
}

type objectTimedate struct {
	interfaceTimedate // interface org.deepin.dde.Timedate1
	proxy.ImplObject
}

func NewTimedate(conn *dbus.Conn) Timedate {
	obj := new(objectTimedate)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Timedate1", "/org/deepin/dde/Timedate1")
	return obj
}

type timedate interface {
	GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call
	SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error
	GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call
	SetNTP(flags dbus.Flags, enabled bool, message string) error
	GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call
	SetNTPServer(flags dbus.Flags, server string, message string) error
	GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call
	SetTime(flags dbus.Flags, usec int64, relative bool, message string) error
	GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call
	SetTimezone(flags dbus.Flags, timezone string, message string) error
	NTPServer() proxy.PropString
	Timezone() proxy.PropString
}

type interfaceTimedate struct{}

func (v *interfaceTimedate) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTimedate) GetInterfaceName_() string {
	return "org.deepin.dde.Timedate1"
}

// method SetLocalRTC

func (v *interfaceTimedate) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocalRTC", flags, ch, enabled, fixSystem, message)
}

func (v *interfaceTimedate) SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	return (<-v.GoSetLocalRTC(flags, make(chan *dbus.Call, 1), enabled, fixSystem, message).Done).Err
}

// method SetNTP

func (v *interfaceTimedate) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTP", flags, ch, enabled, message)
}

func (v *interfaceTimedate) SetNTP(flags dbus.Flags, enabled bool, message string) error {
	return (<-v.GoSetNTP(flags, make(chan *dbus.Call, 1), enabled, message).Done).Err
}

// method SetNTPServer

func (v *interfaceTimedate) GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetNTPServer", flags, ch, server, message)
}

func (v *interfaceTimedate) SetNTPServer(flags dbus.Flags, server string, message string) error {
	return (<-v.GoSetNTPServer(flags, make(chan *dbus.Call, 1), server, message).Done).Err
}

// method SetTime

func (v *interfaceTimedate) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTime", flags, ch, usec, relative, message)
}

func (v *interfaceTimedate) SetTime(flags dbus.Flags, usec int64, relative bool, message string) error {
	return (<-v.GoSetTime(flags, make(chan *dbus.Call, 1), usec, relative, message).Done).Err
}

// method SetTimezone

func (v *interfaceTimedate) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimezone", flags, ch, timezone, message)
}

func (v *interfaceTimedate) SetTimezone(flags dbus.Flags, timezone string, message string) error {
	return (<-v.GoSetTimezone(flags, make(chan *dbus.Call, 1), timezone, message).Done).Err
}

// property NTPServer s

func (v *interfaceTimedate) NTPServer() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "NTPServer",
	}
}

// property Timezone s

func (v *interfaceTimedate) Timezone() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Timezone",
	}
}
