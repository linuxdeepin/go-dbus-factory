// Code generated by "./generator ./system/net.reactivated.fprint"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package fprint

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockManager struct {
	MockInterfaceManager // interface net.reactivated.Fprint.Manager
	proxy.MockObject
}

type MockInterfaceManager struct {
	mock.Mock
}

// method GetDefaultDevice

func (v *MockInterfaceManager) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetDevices

func (v *MockInterfaceManager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockDevice struct {
	MockInterfaceDevice // interface net.reactivated.Fprint.Device
	proxy.MockObject
}

type MockInterfaceDevice struct {
	mock.Mock
}

// method EnrollStop

func (v *MockInterfaceDevice) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) EnrollStop(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method EnrollStart

func (v *MockInterfaceDevice) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger_name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) EnrollStart(flags dbus.Flags, finger_name string) error {
	mockArgs := v.Called(flags, finger_name)

	return mockArgs.Error(0)
}

// method VerifyStop

func (v *MockInterfaceDevice) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) VerifyStop(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method VerifyStart

func (v *MockInterfaceDevice) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger_name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) VerifyStart(flags dbus.Flags, finger_name string) error {
	mockArgs := v.Called(flags, finger_name)

	return mockArgs.Error(0)
}

// method Release

func (v *MockInterfaceDevice) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Release(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Claim

func (v *MockInterfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Claim(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method DeleteEnrolledFingers

func (v *MockInterfaceDevice) GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteEnrolledFingers(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method DeleteEnrolledFinger

func (v *MockInterfaceDevice) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger_name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteEnrolledFinger(flags dbus.Flags, finger_name string) error {
	mockArgs := v.Called(flags, finger_name)

	return mockArgs.Error(0)
}

// method ListEnrolledFingers

func (v *MockInterfaceDevice) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal EnrollStatus

func (v *MockInterfaceDevice) ConnectEnrollStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceDevice) ConnectVerifyStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyFingerSelected

func (v *MockInterfaceDevice) ConnectVerifyFingerSelected(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property scan-type s

func (v *MockInterfaceDevice) ScanType() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property num-enroll-stages i

func (v *MockInterfaceDevice) NumEnrollStages() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property name s

func (v *MockInterfaceDevice) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
