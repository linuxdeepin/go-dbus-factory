// Code generated by "./generator ./system/org.deepin.dde.bluetooth1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package bluetooth1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
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
	GoClearUnpairedDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearUnpairedDevice(flags dbus.Flags) error
	GoConnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) *dbus.Call
	ConnectDevice(flags dbus.Flags, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) error
	GoDebugInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DebugInfo(flags dbus.Flags) (string, error)
	GoDisconnectAudioDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DisconnectAudioDevices(flags dbus.Flags) error
	GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call
	DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error
	GoGetAdapters(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAdapters(flags dbus.Flags) (string, error)
	GoGetDevices(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call
	GetDevices(flags dbus.Flags, adapterPath dbus.ObjectPath) (string, error)
	GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call
	RegisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error
	GoRemoveDevice(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call
	RemoveDevice(flags dbus.Flags, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) error
	GoRequestDiscovery(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call
	RequestDiscovery(flags dbus.Flags, adapterPath dbus.ObjectPath) error
	GoSetAdapterAlias(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, alias string) *dbus.Call
	SetAdapterAlias(flags dbus.Flags, adapterPath dbus.ObjectPath, alias string) error
	GoSetAdapterDiscoverable(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverable bool) *dbus.Call
	SetAdapterDiscoverable(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverable bool) error
	GoSetAdapterDiscoverableTimeout(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverableTimeout uint32) *dbus.Call
	SetAdapterDiscoverableTimeout(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverableTimeout uint32) error
	GoSetAdapterDiscovering(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discovering bool) *dbus.Call
	SetAdapterDiscovering(flags dbus.Flags, adapterPath dbus.ObjectPath, discovering bool) error
	GoSetAdapterPowered(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, powered bool) *dbus.Call
	SetAdapterPowered(flags dbus.Flags, adapterPath dbus.ObjectPath, powered bool) error
	GoSetDeviceAlias(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, alias string) *dbus.Call
	SetDeviceAlias(flags dbus.Flags, device dbus.ObjectPath, alias string) error
	GoSetDeviceTrusted(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, trusted bool) *dbus.Call
	SetDeviceTrusted(flags dbus.Flags, device dbus.ObjectPath, trusted bool) error
	GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call
	UnregisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error
	ConnectAdapterAdded(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error)
	ConnectAdapterRemoved(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error)
	ConnectAdapterPropertiesChanged(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error)
	ConnectDeviceAdded(cb func(devJSON string)) (dbusutil.SignalHandlerId, error)
	ConnectDeviceRemoved(cb func(devJSON string)) (dbusutil.SignalHandlerId, error)
	ConnectDevicePropertiesChanged(cb func(devJSON string)) (dbusutil.SignalHandlerId, error)
	CanSendFile() proxy.PropBool
	State() proxy.PropUint32
}

type interfaceBluetooth struct{}

func (v *interfaceBluetooth) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceBluetooth) GetInterfaceName_() string {
	return "org.deepin.dde.Bluetooth1"
}

// method ClearUnpairedDevice

func (v *interfaceBluetooth) GoClearUnpairedDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearUnpairedDevice", flags, ch)
}

func (v *interfaceBluetooth) ClearUnpairedDevice(flags dbus.Flags) error {
	return (<-v.GoClearUnpairedDevice(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ConnectDevice

func (v *interfaceBluetooth) GoConnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ConnectDevice", flags, ch, devPath, adapterPath)
}

func (v *interfaceBluetooth) ConnectDevice(flags dbus.Flags, devPath dbus.ObjectPath, adapterPath dbus.ObjectPath) error {
	return (<-v.GoConnectDevice(flags, make(chan *dbus.Call, 1), devPath, adapterPath).Done).Err
}

// method DebugInfo

func (v *interfaceBluetooth) GoDebugInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DebugInfo", flags, ch)
}

func (*interfaceBluetooth) StoreDebugInfo(call *dbus.Call) (info string, err error) {
	err = call.Store(&info)
	return
}

func (v *interfaceBluetooth) DebugInfo(flags dbus.Flags) (string, error) {
	return v.StoreDebugInfo(
		<-v.GoDebugInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method DisconnectAudioDevices

func (v *interfaceBluetooth) GoDisconnectAudioDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectAudioDevices", flags, ch)
}

func (v *interfaceBluetooth) DisconnectAudioDevices(flags dbus.Flags) error {
	return (<-v.GoDisconnectAudioDevices(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DisconnectDevice

func (v *interfaceBluetooth) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisconnectDevice", flags, ch, devPath)
}

func (v *interfaceBluetooth) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	return (<-v.GoDisconnectDevice(flags, make(chan *dbus.Call, 1), devPath).Done).Err
}

// method GetAdapters

func (v *interfaceBluetooth) GoGetAdapters(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAdapters", flags, ch)
}

func (*interfaceBluetooth) StoreGetAdapters(call *dbus.Call) (adaptersJSON string, err error) {
	err = call.Store(&adaptersJSON)
	return
}

func (v *interfaceBluetooth) GetAdapters(flags dbus.Flags) (string, error) {
	return v.StoreGetAdapters(
		<-v.GoGetAdapters(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevices

func (v *interfaceBluetooth) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch, adapterPath)
}

func (*interfaceBluetooth) StoreGetDevices(call *dbus.Call) (devicesJSON string, err error) {
	err = call.Store(&devicesJSON)
	return
}

func (v *interfaceBluetooth) GetDevices(flags dbus.Flags, adapterPath dbus.ObjectPath) (string, error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1), adapterPath).Done)
}

// method RegisterAgent

func (v *interfaceBluetooth) GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAgent", flags, ch, agentPath)
}

func (v *interfaceBluetooth) RegisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error {
	return (<-v.GoRegisterAgent(flags, make(chan *dbus.Call, 1), agentPath).Done).Err
}

// method RemoveDevice

func (v *interfaceBluetooth) GoRemoveDevice(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveDevice", flags, ch, adapterPath, devPath)
}

func (v *interfaceBluetooth) RemoveDevice(flags dbus.Flags, adapterPath dbus.ObjectPath, devPath dbus.ObjectPath) error {
	return (<-v.GoRemoveDevice(flags, make(chan *dbus.Call, 1), adapterPath, devPath).Done).Err
}

// method RequestDiscovery

func (v *interfaceBluetooth) GoRequestDiscovery(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestDiscovery", flags, ch, adapterPath)
}

func (v *interfaceBluetooth) RequestDiscovery(flags dbus.Flags, adapterPath dbus.ObjectPath) error {
	return (<-v.GoRequestDiscovery(flags, make(chan *dbus.Call, 1), adapterPath).Done).Err
}

// method SetAdapterAlias

func (v *interfaceBluetooth) GoSetAdapterAlias(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, alias string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAdapterAlias", flags, ch, adapterPath, alias)
}

func (v *interfaceBluetooth) SetAdapterAlias(flags dbus.Flags, adapterPath dbus.ObjectPath, alias string) error {
	return (<-v.GoSetAdapterAlias(flags, make(chan *dbus.Call, 1), adapterPath, alias).Done).Err
}

// method SetAdapterDiscoverable

func (v *interfaceBluetooth) GoSetAdapterDiscoverable(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAdapterDiscoverable", flags, ch, adapterPath, discoverable)
}

func (v *interfaceBluetooth) SetAdapterDiscoverable(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverable bool) error {
	return (<-v.GoSetAdapterDiscoverable(flags, make(chan *dbus.Call, 1), adapterPath, discoverable).Done).Err
}

// method SetAdapterDiscoverableTimeout

func (v *interfaceBluetooth) GoSetAdapterDiscoverableTimeout(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discoverableTimeout uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAdapterDiscoverableTimeout", flags, ch, adapterPath, discoverableTimeout)
}

func (v *interfaceBluetooth) SetAdapterDiscoverableTimeout(flags dbus.Flags, adapterPath dbus.ObjectPath, discoverableTimeout uint32) error {
	return (<-v.GoSetAdapterDiscoverableTimeout(flags, make(chan *dbus.Call, 1), adapterPath, discoverableTimeout).Done).Err
}

// method SetAdapterDiscovering

func (v *interfaceBluetooth) GoSetAdapterDiscovering(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, discovering bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAdapterDiscovering", flags, ch, adapterPath, discovering)
}

func (v *interfaceBluetooth) SetAdapterDiscovering(flags dbus.Flags, adapterPath dbus.ObjectPath, discovering bool) error {
	return (<-v.GoSetAdapterDiscovering(flags, make(chan *dbus.Call, 1), adapterPath, discovering).Done).Err
}

// method SetAdapterPowered

func (v *interfaceBluetooth) GoSetAdapterPowered(flags dbus.Flags, ch chan *dbus.Call, adapterPath dbus.ObjectPath, powered bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAdapterPowered", flags, ch, adapterPath, powered)
}

func (v *interfaceBluetooth) SetAdapterPowered(flags dbus.Flags, adapterPath dbus.ObjectPath, powered bool) error {
	return (<-v.GoSetAdapterPowered(flags, make(chan *dbus.Call, 1), adapterPath, powered).Done).Err
}

// method SetDeviceAlias

func (v *interfaceBluetooth) GoSetDeviceAlias(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, alias string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeviceAlias", flags, ch, device, alias)
}

func (v *interfaceBluetooth) SetDeviceAlias(flags dbus.Flags, device dbus.ObjectPath, alias string) error {
	return (<-v.GoSetDeviceAlias(flags, make(chan *dbus.Call, 1), device, alias).Done).Err
}

// method SetDeviceTrusted

func (v *interfaceBluetooth) GoSetDeviceTrusted(flags dbus.Flags, ch chan *dbus.Call, device dbus.ObjectPath, trusted bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeviceTrusted", flags, ch, device, trusted)
}

func (v *interfaceBluetooth) SetDeviceTrusted(flags dbus.Flags, device dbus.ObjectPath, trusted bool) error {
	return (<-v.GoSetDeviceTrusted(flags, make(chan *dbus.Call, 1), device, trusted).Done).Err
}

// method UnregisterAgent

func (v *interfaceBluetooth) GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agentPath dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterAgent", flags, ch, agentPath)
}

func (v *interfaceBluetooth) UnregisterAgent(flags dbus.Flags, agentPath dbus.ObjectPath) error {
	return (<-v.GoUnregisterAgent(flags, make(chan *dbus.Call, 1), agentPath).Done).Err
}

// signal AdapterAdded

func (v *interfaceBluetooth) ConnectAdapterAdded(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AdapterAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AdapterAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var adapterJSON string
		err := dbus.Store(sig.Body, &adapterJSON)
		if err == nil {
			cb(adapterJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AdapterRemoved

func (v *interfaceBluetooth) ConnectAdapterRemoved(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AdapterRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AdapterRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var adapterJSON string
		err := dbus.Store(sig.Body, &adapterJSON)
		if err == nil {
			cb(adapterJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AdapterPropertiesChanged

func (v *interfaceBluetooth) ConnectAdapterPropertiesChanged(cb func(adapterJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AdapterPropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AdapterPropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var adapterJSON string
		err := dbus.Store(sig.Body, &adapterJSON)
		if err == nil {
			cb(adapterJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceAdded

func (v *interfaceBluetooth) ConnectDeviceAdded(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devJSON string
		err := dbus.Store(sig.Body, &devJSON)
		if err == nil {
			cb(devJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceRemoved

func (v *interfaceBluetooth) ConnectDeviceRemoved(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devJSON string
		err := dbus.Store(sig.Body, &devJSON)
		if err == nil {
			cb(devJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DevicePropertiesChanged

func (v *interfaceBluetooth) ConnectDevicePropertiesChanged(cb func(devJSON string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DevicePropertiesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DevicePropertiesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var devJSON string
		err := dbus.Store(sig.Body, &devJSON)
		if err == nil {
			cb(devJSON)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property CanSendFile b

func (v *interfaceBluetooth) CanSendFile() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanSendFile",
	}
}

// property State u

func (v *interfaceBluetooth) State() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "State",
	}
}
