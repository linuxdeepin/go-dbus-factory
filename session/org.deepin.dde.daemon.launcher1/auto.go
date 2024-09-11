// Code generated by "./generator ./session/org.deepin.dde.daemon.launcher1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package launcher1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Launcher interface {
	launcher // interface org.deepin.dde.daemon.Launcher1
	proxy.Object
}

type objectLauncher struct {
	interfaceLauncher // interface org.deepin.dde.daemon.Launcher1
	proxy.ImplObject
}

func NewLauncher(conn *dbus.Conn) Launcher {
	obj := new(objectLauncher)
	obj.ImplObject.Init_(conn, "org.deepin.dde.daemon.Launcher1", "/org/deepin/dde/daemon/Launcher1")
	return obj
}

type launcher interface {
	GoGetAllItemInfos(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAllItemInfos(flags dbus.Flags) ([]ItemInfo, error)
	GoGetAllNewInstalledApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAllNewInstalledApps(flags dbus.Flags) ([]string, error)
	GoGetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	GetDisableScaling(flags dbus.Flags, id string) (bool, error)
	GoGetItemInfo(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	GetItemInfo(flags dbus.Flags, id string) (ItemInfo, error)
	GoGetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	GetUseProxy(flags dbus.Flags, id string) (bool, error)
	GoIsItemOnDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	IsItemOnDesktop(flags dbus.Flags, id string) (bool, error)
	GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	MarkLaunched(flags dbus.Flags, id string) error
	GoRequestRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	RequestRemoveFromDesktop(flags dbus.Flags, id string) (bool, error)
	GoRequestSendToDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	RequestSendToDesktop(flags dbus.Flags, id string) (bool, error)
	GoRequestUninstall(flags dbus.Flags, ch chan *dbus.Call, id string, purge bool) *dbus.Call
	RequestUninstall(flags dbus.Flags, id string, purge bool) error
	GoSearch(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Search(flags dbus.Flags, key string) error
	GoSetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call
	SetDisableScaling(flags dbus.Flags, id string, value bool) error
	GoSetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call
	SetUseProxy(flags dbus.Flags, id string, value bool) error
	ConnectSearchDone(cb func(apps []string)) (dbusutil.SignalHandlerId, error)
	ConnectItemChanged(cb func(status string, itemInfo ItemInfo, categoryID int64)) (dbusutil.SignalHandlerId, error)
	ConnectNewAppLaunched(cb func(appID string)) (dbusutil.SignalHandlerId, error)
	ConnectUninstallSuccess(cb func(appID string)) (dbusutil.SignalHandlerId, error)
	ConnectUninstallFailed(cb func(appId string, errMsg string)) (dbusutil.SignalHandlerId, error)
	Fullscreen() proxy.PropBool
	DisplayMode() proxy.PropInt32
}

type interfaceLauncher struct{}

func (v *interfaceLauncher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLauncher) GetInterfaceName_() string {
	return "org.deepin.dde.daemon.Launcher1"
}

// method GetAllItemInfos

func (v *interfaceLauncher) GoGetAllItemInfos(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllItemInfos", flags, ch)
}

func (*interfaceLauncher) StoreGetAllItemInfos(call *dbus.Call) (itemInfoList []ItemInfo, err error) {
	err = call.Store(&itemInfoList)
	return
}

func (v *interfaceLauncher) GetAllItemInfos(flags dbus.Flags) ([]ItemInfo, error) {
	return v.StoreGetAllItemInfos(
		<-v.GoGetAllItemInfos(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAllNewInstalledApps

func (v *interfaceLauncher) GoGetAllNewInstalledApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllNewInstalledApps", flags, ch)
}

func (*interfaceLauncher) StoreGetAllNewInstalledApps(call *dbus.Call) (apps []string, err error) {
	err = call.Store(&apps)
	return
}

func (v *interfaceLauncher) GetAllNewInstalledApps(flags dbus.Flags) ([]string, error) {
	return v.StoreGetAllNewInstalledApps(
		<-v.GoGetAllNewInstalledApps(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDisableScaling

func (v *interfaceLauncher) GoGetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDisableScaling", flags, ch, id)
}

func (*interfaceLauncher) StoreGetDisableScaling(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceLauncher) GetDisableScaling(flags dbus.Flags, id string) (bool, error) {
	return v.StoreGetDisableScaling(
		<-v.GoGetDisableScaling(flags, make(chan *dbus.Call, 1), id).Done)
}

// method GetItemInfo

func (v *interfaceLauncher) GoGetItemInfo(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetItemInfo", flags, ch, id)
}

func (*interfaceLauncher) StoreGetItemInfo(call *dbus.Call) (itemInfo ItemInfo, err error) {
	err = call.Store(&itemInfo)
	return
}

func (v *interfaceLauncher) GetItemInfo(flags dbus.Flags, id string) (ItemInfo, error) {
	return v.StoreGetItemInfo(
		<-v.GoGetItemInfo(flags, make(chan *dbus.Call, 1), id).Done)
}

// method GetUseProxy

func (v *interfaceLauncher) GoGetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUseProxy", flags, ch, id)
}

func (*interfaceLauncher) StoreGetUseProxy(call *dbus.Call) (value bool, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceLauncher) GetUseProxy(flags dbus.Flags, id string) (bool, error) {
	return v.StoreGetUseProxy(
		<-v.GoGetUseProxy(flags, make(chan *dbus.Call, 1), id).Done)
}

// method IsItemOnDesktop

func (v *interfaceLauncher) GoIsItemOnDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsItemOnDesktop", flags, ch, id)
}

func (*interfaceLauncher) StoreIsItemOnDesktop(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceLauncher) IsItemOnDesktop(flags dbus.Flags, id string) (bool, error) {
	return v.StoreIsItemOnDesktop(
		<-v.GoIsItemOnDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method MarkLaunched

func (v *interfaceLauncher) GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MarkLaunched", flags, ch, id)
}

func (v *interfaceLauncher) MarkLaunched(flags dbus.Flags, id string) error {
	return (<-v.GoMarkLaunched(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method RequestRemoveFromDesktop

func (v *interfaceLauncher) GoRequestRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestRemoveFromDesktop", flags, ch, id)
}

func (*interfaceLauncher) StoreRequestRemoveFromDesktop(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *interfaceLauncher) RequestRemoveFromDesktop(flags dbus.Flags, id string) (bool, error) {
	return v.StoreRequestRemoveFromDesktop(
		<-v.GoRequestRemoveFromDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method RequestSendToDesktop

func (v *interfaceLauncher) GoRequestSendToDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSendToDesktop", flags, ch, id)
}

func (*interfaceLauncher) StoreRequestSendToDesktop(call *dbus.Call) (ok bool, err error) {
	err = call.Store(&ok)
	return
}

func (v *interfaceLauncher) RequestSendToDesktop(flags dbus.Flags, id string) (bool, error) {
	return v.StoreRequestSendToDesktop(
		<-v.GoRequestSendToDesktop(flags, make(chan *dbus.Call, 1), id).Done)
}

// method RequestUninstall

func (v *interfaceLauncher) GoRequestUninstall(flags dbus.Flags, ch chan *dbus.Call, id string, purge bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestUninstall", flags, ch, id, purge)
}

func (v *interfaceLauncher) RequestUninstall(flags dbus.Flags, id string, purge bool) error {
	return (<-v.GoRequestUninstall(flags, make(chan *dbus.Call, 1), id, purge).Done).Err
}

// method Search

func (v *interfaceLauncher) GoSearch(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Search", flags, ch, key)
}

func (v *interfaceLauncher) Search(flags dbus.Flags, key string) error {
	return (<-v.GoSearch(flags, make(chan *dbus.Call, 1), key).Done).Err
}

// method SetDisableScaling

func (v *interfaceLauncher) GoSetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDisableScaling", flags, ch, id, value)
}

func (v *interfaceLauncher) SetDisableScaling(flags dbus.Flags, id string, value bool) error {
	return (<-v.GoSetDisableScaling(flags, make(chan *dbus.Call, 1), id, value).Done).Err
}

// method SetUseProxy

func (v *interfaceLauncher) GoSetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUseProxy", flags, ch, id, value)
}

func (v *interfaceLauncher) SetUseProxy(flags dbus.Flags, id string, value bool) error {
	return (<-v.GoSetUseProxy(flags, make(chan *dbus.Call, 1), id, value).Done).Err
}

// signal SearchDone

func (v *interfaceLauncher) ConnectSearchDone(cb func(apps []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SearchDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SearchDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var apps []string
		err := dbus.Store(sig.Body, &apps)
		if err == nil {
			cb(apps)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemChanged

func (v *interfaceLauncher) ConnectItemChanged(cb func(status string, itemInfo ItemInfo, categoryID int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var itemInfo ItemInfo
		var categoryID int64
		err := dbus.Store(sig.Body, &status, &itemInfo, &categoryID)
		if err == nil {
			cb(status, itemInfo, categoryID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NewAppLaunched

func (v *interfaceLauncher) ConnectNewAppLaunched(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NewAppLaunched", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NewAppLaunched",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appID string
		err := dbus.Store(sig.Body, &appID)
		if err == nil {
			cb(appID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UninstallSuccess

func (v *interfaceLauncher) ConnectUninstallSuccess(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UninstallSuccess", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UninstallSuccess",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appID string
		err := dbus.Store(sig.Body, &appID)
		if err == nil {
			cb(appID)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UninstallFailed

func (v *interfaceLauncher) ConnectUninstallFailed(cb func(appId string, errMsg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UninstallFailed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UninstallFailed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var appId string
		var errMsg string
		err := dbus.Store(sig.Body, &appId, &errMsg)
		if err == nil {
			cb(appId, errMsg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Fullscreen b

func (v *interfaceLauncher) Fullscreen() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Fullscreen",
	}
}

// property DisplayMode i

func (v *interfaceLauncher) DisplayMode() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "DisplayMode",
	}
}
