// Code generated by "./generator ./net.reactivated.fprint"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package fprint

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Manager interface {
	manager // interface net.reactivated.Fprint.Manager
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface net.reactivated.Fprint.Manager
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn) Manager {
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "net.reactivated.Fprint", "/net/reactivated/Fprint/Manager")
	return obj
}

type manager interface {
	GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error)
	GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error)
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "net.reactivated.Fprint.Manager"
}

// method GetDefaultDevice

func (v *interfaceManager) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (*interfaceManager) StoreGetDefaultDevice(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *interfaceManager) GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreGetDefaultDevice(
		<-v.GoGetDefaultDevice(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevices

func (v *interfaceManager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*interfaceManager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

type Device interface {
	device // interface net.reactivated.Fprint.Device
	proxy.Object
}

type objectDevice struct {
	interfaceDevice // interface net.reactivated.Fprint.Device
	proxy.ImplObject
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectDevice)
	obj.ImplObject.Init_(conn, "net.reactivated.Fprint", path)
	return obj, nil
}

type device interface {
	GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	EnrollStop(flags dbus.Flags) error
	GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call
	EnrollStart(flags dbus.Flags, finger_name string) error
	GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	VerifyStop(flags dbus.Flags) error
	GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call
	VerifyStart(flags dbus.Flags, finger_name string) error
	GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Release(flags dbus.Flags) error
	GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	Claim(flags dbus.Flags, username string) error
	GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	DeleteEnrolledFingers(flags dbus.Flags, username string) error
	GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger_name string) *dbus.Call
	DeleteEnrolledFinger(flags dbus.Flags, username string, finger_name string) error
	GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error)
	ConnectEnrollStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyFingerSelected(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ScanType() proxy.PropString
	NumEnrollStages() proxy.PropInt32
	Name() proxy.PropString
}

type interfaceDevice struct{}

func (v *interfaceDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDevice) GetInterfaceName_() string {
	return "net.reactivated.Fprint.Device"
}

// method EnrollStop

func (v *interfaceDevice) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStop", flags, ch)
}

func (v *interfaceDevice) EnrollStop(flags dbus.Flags) error {
	return (<-v.GoEnrollStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method EnrollStart

func (v *interfaceDevice) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStart", flags, ch, finger_name)
}

func (v *interfaceDevice) EnrollStart(flags dbus.Flags, finger_name string) error {
	return (<-v.GoEnrollStart(flags, make(chan *dbus.Call, 1), finger_name).Done).Err
}

// method VerifyStop

func (v *interfaceDevice) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStop", flags, ch)
}

func (v *interfaceDevice) VerifyStop(flags dbus.Flags) error {
	return (<-v.GoVerifyStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method VerifyStart

func (v *interfaceDevice) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStart", flags, ch, finger_name)
}

func (v *interfaceDevice) VerifyStart(flags dbus.Flags, finger_name string) error {
	return (<-v.GoVerifyStart(flags, make(chan *dbus.Call, 1), finger_name).Done).Err
}

// method Release

func (v *interfaceDevice) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Release", flags, ch)
}

func (v *interfaceDevice) Release(flags dbus.Flags) error {
	return (<-v.GoRelease(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Claim

func (v *interfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, username)
}

func (v *interfaceDevice) Claim(flags dbus.Flags, username string) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method DeleteEnrolledFingers

func (v *interfaceDevice) GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFingers", flags, ch, username)
}

func (v *interfaceDevice) DeleteEnrolledFingers(flags dbus.Flags, username string) error {
	return (<-v.GoDeleteEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method DeleteEnrolledFinger

func (v *interfaceDevice) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger_name)
}

func (v *interfaceDevice) DeleteEnrolledFinger(flags dbus.Flags, username string, finger_name string) error {
	return (<-v.GoDeleteEnrolledFinger(flags, make(chan *dbus.Call, 1), username, finger_name).Done).Err
}

// method ListEnrolledFingers

func (v *interfaceDevice) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (*interfaceDevice) StoreListEnrolledFingers(call *dbus.Call) (enrolled_fingers []string, err error) {
	err = call.Store(&enrolled_fingers)
	return
}

func (v *interfaceDevice) ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error) {
	return v.StoreListEnrolledFingers(
		<-v.GoListEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

// signal EnrollStatus

func (v *interfaceDevice) ConnectEnrollStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 bool
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceDevice) ConnectVerifyStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 bool
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyFingerSelected

func (v *interfaceDevice) ConnectVerifyFingerSelected(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyFingerSelected", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyFingerSelected",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property scan-type s

func (v *interfaceDevice) ScanType() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "scan-type",
	}
}

// property num-enroll-stages i

func (v *interfaceDevice) NumEnrollStages() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "num-enroll-stages",
	}
}

// property name s

func (v *interfaceDevice) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "name",
	}
}
