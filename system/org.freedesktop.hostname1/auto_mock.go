// Code generated by "./generator ./system/org.freedesktop.hostname1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package hostname1

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockHostname struct {
	MockInterfaceHostname // interface org.freedesktop.hostname1
	proxy.MockObject
}

type MockInterfaceHostname struct {
	mock.Mock
}

// method SetHostname

func (v *MockInterfaceHostname) GoSetHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetStaticHostname

func (v *MockInterfaceHostname) GoSetStaticHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetStaticHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetPrettyHostname

func (v *MockInterfaceHostname) GoSetPrettyHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetPrettyHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetIconName

func (v *MockInterfaceHostname) GoSetIconName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetIconName(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetChassis

func (v *MockInterfaceHostname) GoSetChassis(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetChassis(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetDeployment

func (v *MockInterfaceHostname) GoSetDeployment(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetDeployment(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetLocation

func (v *MockInterfaceHostname) GoSetLocation(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) SetLocation(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method GetProductUUID

func (v *MockInterfaceHostname) GoGetProductUUID(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceHostname) GetProductUUID(flags dbus.Flags, arg0 bool) ([]uint8, error) {
	mockArgs := v.Called(flags, arg0)

	ret0, ok := mockArgs.Get(0).([]uint8)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Hostname s

func (v *MockInterfaceHostname) Hostname() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property StaticHostname s

func (v *MockInterfaceHostname) StaticHostname() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PrettyHostname s

func (v *MockInterfaceHostname) PrettyHostname() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IconName s

func (v *MockInterfaceHostname) IconName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Chassis s

func (v *MockInterfaceHostname) Chassis() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Deployment s

func (v *MockInterfaceHostname) Deployment() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Location s

func (v *MockInterfaceHostname) Location() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property KernelName s

func (v *MockInterfaceHostname) KernelName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property KernelRelease s

func (v *MockInterfaceHostname) KernelRelease() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property KernelVersion s

func (v *MockInterfaceHostname) KernelVersion() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property OperatingSystemPrettyName s

func (v *MockInterfaceHostname) OperatingSystemPrettyName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property OperatingSystemCPEName s

func (v *MockInterfaceHostname) OperatingSystemCPEName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HomeURL s

func (v *MockInterfaceHostname) HomeURL() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
