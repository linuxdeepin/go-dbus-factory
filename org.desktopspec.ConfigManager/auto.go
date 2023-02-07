// Code generated by "./generator ./org.desktopspec.ConfigManager"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package ConfigManager

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type ConfigManager interface {
	configManager // interface org.desktopspec.ConfigManager
	proxy.Object
}

type objectConfigManager struct {
	interfaceConfigManager // interface org.desktopspec.ConfigManager
	proxy.ImplObject
}

func NewConfigManager(conn *dbus.Conn) ConfigManager {
	obj := new(objectConfigManager)
	obj.ImplObject.Init_(conn, "org.desktopspec.ConfigManager", "/")
	return obj
}

type configManager interface {
	GoAcquireManager(flags dbus.Flags, ch chan *dbus.Call, appid string, name string, subpath string) *dbus.Call
	AcquireManager(flags dbus.Flags, appid string, name string, subpath string) (dbus.ObjectPath, error)
	GoUpdate(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call
	Update(flags dbus.Flags, path string) error
	GoSync(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call
	Sync(flags dbus.Flags, path string) error
	GoSetDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call, time int32) *dbus.Call
	SetDelayReleaseTime(flags dbus.Flags, time int32) error
	GoDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DelayReleaseTime(flags dbus.Flags) (int32, error)
}

type interfaceConfigManager struct{}

func (v *interfaceConfigManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceConfigManager) GetInterfaceName_() string {
	return "org.desktopspec.ConfigManager"
}

// method acquireManager

func (v *interfaceConfigManager) GoAcquireManager(flags dbus.Flags, ch chan *dbus.Call, appid string, name string, subpath string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".acquireManager", flags, ch, appid, name, subpath)
}

func (*interfaceConfigManager) StoreAcquireManager(call *dbus.Call) (path dbus.ObjectPath, err error) {
	err = call.Store(&path)
	return
}

func (v *interfaceConfigManager) AcquireManager(flags dbus.Flags, appid string, name string, subpath string) (dbus.ObjectPath, error) {
	return v.StoreAcquireManager(
		<-v.GoAcquireManager(flags, make(chan *dbus.Call, 1), appid, name, subpath).Done)
}

// method update

func (v *interfaceConfigManager) GoUpdate(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".update", flags, ch, path)
}

func (v *interfaceConfigManager) Update(flags dbus.Flags, path string) error {
	return (<-v.GoUpdate(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method sync

func (v *interfaceConfigManager) GoSync(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".sync", flags, ch, path)
}

func (v *interfaceConfigManager) Sync(flags dbus.Flags, path string) error {
	return (<-v.GoSync(flags, make(chan *dbus.Call, 1), path).Done).Err
}

// method setDelayReleaseTime

func (v *interfaceConfigManager) GoSetDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call, time int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".setDelayReleaseTime", flags, ch, time)
}

func (v *interfaceConfigManager) SetDelayReleaseTime(flags dbus.Flags, time int32) error {
	return (<-v.GoSetDelayReleaseTime(flags, make(chan *dbus.Call, 1), time).Done).Err
}

// method delayReleaseTime

func (v *interfaceConfigManager) GoDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".delayReleaseTime", flags, ch)
}

func (*interfaceConfigManager) StoreDelayReleaseTime(call *dbus.Call) (time int32, err error) {
	err = call.Store(&time)
	return
}

func (v *interfaceConfigManager) DelayReleaseTime(flags dbus.Flags) (int32, error) {
	return v.StoreDelayReleaseTime(
		<-v.GoDelayReleaseTime(flags, make(chan *dbus.Call, 1)).Done)
}

type Manager interface {
	manager // interface org.desktopspec.ConfigManager.Manager
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface org.desktopspec.ConfigManager.Manager
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn, path dbus.ObjectPath) (Manager, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "org.desktopspec.ConfigManager", path)
	return obj, nil
}

type manager interface {
	GoValue(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Value(flags dbus.Flags, key string) (dbus.Variant, error)
	GoSetValue(flags dbus.Flags, ch chan *dbus.Call, key string, value dbus.Variant) *dbus.Call
	SetValue(flags dbus.Flags, key string, value dbus.Variant) error
	GoReset(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Reset(flags dbus.Flags, key string) error
	GoName(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call
	Name(flags dbus.Flags, key string, language string) (string, error)
	GoDescription(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call
	Description(flags dbus.Flags, key string, language string) (string, error)
	GoVisibility(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Visibility(flags dbus.Flags, key string) (string, error)
	GoPermissions(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Permissions(flags dbus.Flags, key string) (string, error)
	GoFlags(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call
	Flags(flags dbus.Flags, key string) (int32, error)
	GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Release(flags dbus.Flags) error
	ConnectValueChanged(cb func(key string)) (dbusutil.SignalHandlerId, error)
	Version() proxy.PropString
	KeyList() proxy.PropStringArray
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "org.desktopspec.ConfigManager.Manager"
}

// method value

func (v *interfaceManager) GoValue(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".value", flags, ch, key)
}

func (*interfaceManager) StoreValue(call *dbus.Call) (value dbus.Variant, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceManager) Value(flags dbus.Flags, key string) (dbus.Variant, error) {
	return v.StoreValue(
		<-v.GoValue(flags, make(chan *dbus.Call, 1), key).Done)
}

// method setValue

func (v *interfaceManager) GoSetValue(flags dbus.Flags, ch chan *dbus.Call, key string, value dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".setValue", flags, ch, key, value)
}

func (v *interfaceManager) SetValue(flags dbus.Flags, key string, value dbus.Variant) error {
	return (<-v.GoSetValue(flags, make(chan *dbus.Call, 1), key, value).Done).Err
}

// method reset

func (v *interfaceManager) GoReset(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".reset", flags, ch, key)
}

func (v *interfaceManager) Reset(flags dbus.Flags, key string) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1), key).Done).Err
}

// method name

func (v *interfaceManager) GoName(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".name", flags, ch, key, language)
}

func (*interfaceManager) StoreName(call *dbus.Call) (name string, err error) {
	err = call.Store(&name)
	return
}

func (v *interfaceManager) Name(flags dbus.Flags, key string, language string) (string, error) {
	return v.StoreName(
		<-v.GoName(flags, make(chan *dbus.Call, 1), key, language).Done)
}

// method description

func (v *interfaceManager) GoDescription(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".description", flags, ch, key, language)
}

func (*interfaceManager) StoreDescription(call *dbus.Call) (description string, err error) {
	err = call.Store(&description)
	return
}

func (v *interfaceManager) Description(flags dbus.Flags, key string, language string) (string, error) {
	return v.StoreDescription(
		<-v.GoDescription(flags, make(chan *dbus.Call, 1), key, language).Done)
}

// method visibility

func (v *interfaceManager) GoVisibility(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".visibility", flags, ch, key)
}

func (*interfaceManager) StoreVisibility(call *dbus.Call) (visibility string, err error) {
	err = call.Store(&visibility)
	return
}

func (v *interfaceManager) Visibility(flags dbus.Flags, key string) (string, error) {
	return v.StoreVisibility(
		<-v.GoVisibility(flags, make(chan *dbus.Call, 1), key).Done)
}

// method permissions

func (v *interfaceManager) GoPermissions(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".permissions", flags, ch, key)
}

func (*interfaceManager) StorePermissions(call *dbus.Call) (permissions string, err error) {
	err = call.Store(&permissions)
	return
}

func (v *interfaceManager) Permissions(flags dbus.Flags, key string) (string, error) {
	return v.StorePermissions(
		<-v.GoPermissions(flags, make(chan *dbus.Call, 1), key).Done)
}

// method flags

func (v *interfaceManager) GoFlags(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".flags", flags, ch, key)
}

func (*interfaceManager) StoreFlags(call *dbus.Call) (flags0 int32, err error) {
	err = call.Store(&flags0)
	return
}

func (v *interfaceManager) Flags(flags dbus.Flags, key string) (int32, error) {
	return v.StoreFlags(
		<-v.GoFlags(flags, make(chan *dbus.Call, 1), key).Done)
}

// method release

func (v *interfaceManager) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".release", flags, ch)
}

func (v *interfaceManager) Release(flags dbus.Flags) error {
	return (<-v.GoRelease(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal valueChanged

func (v *interfaceManager) ConnectValueChanged(cb func(key string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "valueChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".valueChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var key string
		err := dbus.Store(sig.Body, &key)
		if err == nil {
			cb(key)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property version s

func (v *interfaceManager) Version() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "version",
	}
}

// property keyList as

func (v *interfaceManager) KeyList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "keyList",
	}
}
