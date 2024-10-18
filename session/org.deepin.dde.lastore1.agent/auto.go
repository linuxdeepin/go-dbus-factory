// Code generated by "./generator ./session/org.deepin.dde.lastore1.agent"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package agent

import (
	"errors"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Agent interface {
	agent // interface org.deepin.dde.Lastore1.Agent
	proxy.Object
}

type objectAgent struct {
	interfaceAgent // interface org.deepin.dde.Lastore1.Agent
	proxy.ImplObject
}

func NewAgent(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (Agent, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectAgent)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type agent interface {
	GoGetManualProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetManualProxy(flags dbus.Flags) (map[string]string, error)
	GoSendNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call
	SendNotify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, error)
	GoReportLog(flags dbus.Flags, ch chan *dbus.Call, log string) *dbus.Call
	ReportLog(flags dbus.Flags, log string) error
	GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call
	CloseNotification(flags dbus.Flags, id uint32) error
}

type interfaceAgent struct{}

func (v *interfaceAgent) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceAgent) GetInterfaceName_() string {
	return "org.deepin.dde.Lastore1.Agent"
}

// method GetManualProxy

func (v *interfaceAgent) GoGetManualProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetManualProxy", flags, ch)
}

func (*interfaceAgent) StoreGetManualProxy(call *dbus.Call) (outArg0 map[string]string, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceAgent) GetManualProxy(flags dbus.Flags) (map[string]string, error) {
	return v.StoreGetManualProxy(
		<-v.GoGetManualProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method SendNotify

func (v *interfaceAgent) GoSendNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SendNotify", flags, ch, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)
}

func (*interfaceAgent) StoreSendNotify(call *dbus.Call) (id uint32, err error) {
	err = call.Store(&id)
	return
}

func (v *interfaceAgent) SendNotify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, error) {
	return v.StoreSendNotify(
		<-v.GoSendNotify(flags, make(chan *dbus.Call, 1), app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout).Done)
}

// method ReportLog

func (v *interfaceAgent) GoReportLog(flags dbus.Flags, ch chan *dbus.Call, log string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReportLog", flags, ch, log)
}

func (v *interfaceAgent) ReportLog(flags dbus.Flags, log string) error {
	return (<-v.GoReportLog(flags, make(chan *dbus.Call, 1), log).Done).Err
}

// method CloseNotification

func (v *interfaceAgent) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseNotification", flags, ch, id)
}

func (v *interfaceAgent) CloseNotification(flags dbus.Flags, id uint32) error {
	return (<-v.GoCloseNotification(flags, make(chan *dbus.Call, 1), id).Done).Err
}