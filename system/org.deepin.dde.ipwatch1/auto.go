// Code generated by "./generator ./system/org.deepin.dde.ipwatch1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package ipwatch1

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type IPWatch interface {
	ipwatch // interface org.deepin.dde.IPWatch1
	proxy.Object
}

type objectIPWatch struct {
	interfaceIpwatch // interface org.deepin.dde.IPWatch1
	proxy.ImplObject
}

func NewIPWatch(conn *dbus.Conn) IPWatch {
	obj := new(objectIPWatch)
	obj.ImplObject.Init_(conn, "org.deepin.dde.IPWatch1", "/org/deepin/dde/IPWatch1")
	return obj
}

type ipwatch interface {
	GoRequestIPConflictCheck(flags dbus.Flags, ch chan *dbus.Call, ip string, ifc string) *dbus.Call
	RequestIPConflictCheck(flags dbus.Flags, ip string, ifc string) (string, error)
	ConnectIPConflict(cb func(ip string, smac string, dmac string)) (dbusutil.SignalHandlerId, error)
}

type interfaceIpwatch struct{}

func (v *interfaceIpwatch) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceIpwatch) GetInterfaceName_() string {
	return "org.deepin.dde.IPWatch1"
}

// method RequestIPConflictCheck

func (v *interfaceIpwatch) GoRequestIPConflictCheck(flags dbus.Flags, ch chan *dbus.Call, ip string, ifc string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestIPConflictCheck", flags, ch, ip, ifc)
}

func (*interfaceIpwatch) StoreRequestIPConflictCheck(call *dbus.Call) (mac string, err error) {
	err = call.Store(&mac)
	return
}

func (v *interfaceIpwatch) RequestIPConflictCheck(flags dbus.Flags, ip string, ifc string) (string, error) {
	return v.StoreRequestIPConflictCheck(
		<-v.GoRequestIPConflictCheck(flags, make(chan *dbus.Call, 1), ip, ifc).Done)
}

// signal IPConflict

func (v *interfaceIpwatch) ConnectIPConflict(cb func(ip string, smac string, dmac string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IPConflict", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IPConflict",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip string
		var smac string
		var dmac string
		err := dbus.Store(sig.Body, &ip, &smac, &dmac)
		if err == nil {
			cb(ip, smac, dmac)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
