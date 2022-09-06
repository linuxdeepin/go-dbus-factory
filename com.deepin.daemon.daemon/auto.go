// Code generated by "./generator ./com.deepin.daemon.daemon"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package daemon

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Daemon interface {
	daemon // interface com.deepin.daemon.Daemon
	proxy.Object
}

type objectDaemon struct {
	interfaceDaemon // interface com.deepin.daemon.Daemon
	proxy.ImplObject
}

func NewDaemon(conn *dbus.Conn) Daemon {
	obj := new(objectDaemon)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Daemon", "/com/deepin/daemon/Daemon")
	return obj
}

type daemon interface {
	GoBluetoothGetDeviceTechnologies(flags dbus.Flags, ch chan *dbus.Call, adapter string, device string) *dbus.Call
	BluetoothGetDeviceTechnologies(flags dbus.Flags, adapter string, device string) ([]string, error)
	GoIsPidVirtualMachine(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call
	IsPidVirtualMachine(flags dbus.Flags, pid uint32) (bool, error)
	GoClearTtys(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearTtys(flags dbus.Flags) error
	GoClearTty(flags dbus.Flags, ch chan *dbus.Call, num uint32) *dbus.Call
	ClearTty(flags dbus.Flags, num uint32) error
	GoNetworkGetConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	NetworkGetConnections(flags dbus.Flags) ([]uint8, error)
	GoNetworkSetConnections(flags dbus.Flags, ch chan *dbus.Call, data []uint8) *dbus.Call
	NetworkSetConnections(flags dbus.Flags, data []uint8) error
	GoScalePlymouth(flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call
	ScalePlymouth(flags dbus.Flags, scale uint32) error
	GoSetLongPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call
	SetLongPressDuration(flags dbus.Flags, duration uint32) error
	GoSaveCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call
	SaveCustomWallPaper(flags dbus.Flags, username string, file string) (string, error)
	GoGetCustomWallPapers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	GetCustomWallPapers(flags dbus.Flags, username string) ([]string, error)
	GoDeleteCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call
	DeleteCustomWallPaper(flags dbus.Flags, username string, file string) error
	ConnectHandleForSleep(cb func(start bool)) (dbusutil.SignalHandlerId, error)
}

type interfaceDaemon struct{}

func (v *interfaceDaemon) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDaemon) GetInterfaceName_() string {
	return "com.deepin.daemon.Daemon"
}

// method BluetoothGetDeviceTechnologies

func (v *interfaceDaemon) GoBluetoothGetDeviceTechnologies(flags dbus.Flags, ch chan *dbus.Call, adapter string, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BluetoothGetDeviceTechnologies", flags, ch, adapter, device)
}

func (*interfaceDaemon) StoreBluetoothGetDeviceTechnologies(call *dbus.Call) (technologies []string, err error) {
	err = call.Store(&technologies)
	return
}

func (v *interfaceDaemon) BluetoothGetDeviceTechnologies(flags dbus.Flags, adapter string, device string) ([]string, error) {
	return v.StoreBluetoothGetDeviceTechnologies(
		<-v.GoBluetoothGetDeviceTechnologies(flags, make(chan *dbus.Call, 1), adapter, device).Done)
}

// method IsPidVirtualMachine

func (v *interfaceDaemon) GoIsPidVirtualMachine(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPidVirtualMachine", flags, ch, pid)
}

func (*interfaceDaemon) StoreIsPidVirtualMachine(call *dbus.Call) (ret bool, err error) {
	err = call.Store(&ret)
	return
}

func (v *interfaceDaemon) IsPidVirtualMachine(flags dbus.Flags, pid uint32) (bool, error) {
	return v.StoreIsPidVirtualMachine(
		<-v.GoIsPidVirtualMachine(flags, make(chan *dbus.Call, 1), pid).Done)
}

// method ClearTtys

func (v *interfaceDaemon) GoClearTtys(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearTtys", flags, ch)
}

func (v *interfaceDaemon) ClearTtys(flags dbus.Flags) error {
	return (<-v.GoClearTtys(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ClearTty

func (v *interfaceDaemon) GoClearTty(flags dbus.Flags, ch chan *dbus.Call, num uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearTty", flags, ch, num)
}

func (v *interfaceDaemon) ClearTty(flags dbus.Flags, num uint32) error {
	return (<-v.GoClearTty(flags, make(chan *dbus.Call, 1), num).Done).Err
}

// method NetworkGetConnections

func (v *interfaceDaemon) GoNetworkGetConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NetworkGetConnections", flags, ch)
}

func (*interfaceDaemon) StoreNetworkGetConnections(call *dbus.Call) (data []uint8, err error) {
	err = call.Store(&data)
	return
}

func (v *interfaceDaemon) NetworkGetConnections(flags dbus.Flags) ([]uint8, error) {
	return v.StoreNetworkGetConnections(
		<-v.GoNetworkGetConnections(flags, make(chan *dbus.Call, 1)).Done)
}

// method NetworkSetConnections

func (v *interfaceDaemon) GoNetworkSetConnections(flags dbus.Flags, ch chan *dbus.Call, data []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NetworkSetConnections", flags, ch, data)
}

func (v *interfaceDaemon) NetworkSetConnections(flags dbus.Flags, data []uint8) error {
	return (<-v.GoNetworkSetConnections(flags, make(chan *dbus.Call, 1), data).Done).Err
}

// method ScalePlymouth

func (v *interfaceDaemon) GoScalePlymouth(flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ScalePlymouth", flags, ch, scale)
}

func (v *interfaceDaemon) ScalePlymouth(flags dbus.Flags, scale uint32) error {
	return (<-v.GoScalePlymouth(flags, make(chan *dbus.Call, 1), scale).Done).Err
}

// method SetLongPressDuration

func (v *interfaceDaemon) GoSetLongPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongPressDuration", flags, ch, duration)
}

func (v *interfaceDaemon) SetLongPressDuration(flags dbus.Flags, duration uint32) error {
	return (<-v.GoSetLongPressDuration(flags, make(chan *dbus.Call, 1), duration).Done).Err
}

// method SaveCustomWallPaper

func (v *interfaceDaemon) GoSaveCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveCustomWallPaper", flags, ch, username, file)
}

func (*interfaceDaemon) StoreSaveCustomWallPaper(call *dbus.Call) (path string, err error) {
	err = call.Store(&path)
	return
}

func (v *interfaceDaemon) SaveCustomWallPaper(flags dbus.Flags, username string, file string) (string, error) {
	return v.StoreSaveCustomWallPaper(
		<-v.GoSaveCustomWallPaper(flags, make(chan *dbus.Call, 1), username, file).Done)
}

// method GetCustomWallPapers

func (v *interfaceDaemon) GoGetCustomWallPapers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCustomWallPapers", flags, ch, username)
}

func (*interfaceDaemon) StoreGetCustomWallPapers(call *dbus.Call) (files []string, err error) {
	err = call.Store(&files)
	return
}

func (v *interfaceDaemon) GetCustomWallPapers(flags dbus.Flags, username string) ([]string, error) {
	return v.StoreGetCustomWallPapers(
		<-v.GoGetCustomWallPapers(flags, make(chan *dbus.Call, 1), username).Done)
}

// method DeleteCustomWallPaper

func (v *interfaceDaemon) GoDeleteCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCustomWallPaper", flags, ch, username, file)
}

func (v *interfaceDaemon) DeleteCustomWallPaper(flags dbus.Flags, username string, file string) error {
	return (<-v.GoDeleteCustomWallPaper(flags, make(chan *dbus.Call, 1), username, file).Done).Err
}

// signal HandleForSleep

func (v *interfaceDaemon) ConnectHandleForSleep(cb func(start bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "HandleForSleep", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".HandleForSleep",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var start bool
		err := dbus.Store(sig.Body, &start)
		if err == nil {
			cb(start)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
