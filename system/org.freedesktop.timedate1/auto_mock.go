// Code generated by "./generator ./system/org.freedesktop.timedate1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package timedate1

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockTimedate struct {
	MockInterfaceTimedate // interface org.freedesktop.timedate1
	proxy.MockObject
}

type MockInterfaceTimedate struct {
	mock.Mock
}

// method SetTime

func (v *MockInterfaceTimedate) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, arg0 int64, arg1 bool, arg2 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetTime(flags dbus.Flags, arg0 int64, arg1 bool, arg2 bool) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2)

	return mockArgs.Error(0)
}

// method SetTimezone

func (v *MockInterfaceTimedate) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetTimezone(flags dbus.Flags, arg0 string, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetLocalRTC

func (v *MockInterfaceTimedate) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool, arg2 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetLocalRTC(flags dbus.Flags, arg0 bool, arg1 bool, arg2 bool) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2)

	return mockArgs.Error(0)
}

// method SetNTP

func (v *MockInterfaceTimedate) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, arg0 bool, arg1 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetNTP(flags dbus.Flags, arg0 bool, arg1 bool) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method ListTimezones

func (v *MockInterfaceTimedate) GoListTimezones(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) ListTimezones(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Timezone s

func (v *MockInterfaceTimedate) Timezone() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LocalRTC b

func (v *MockInterfaceTimedate) LocalRTC() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CanNTP b

func (v *MockInterfaceTimedate) CanNTP() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property NTP b

func (v *MockInterfaceTimedate) NTP() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property NTPSynchronized b

func (v *MockInterfaceTimedate) NTPSynchronized() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property TimeUSec t

func (v *MockInterfaceTimedate) TimeUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property RTCTimeUSec t

func (v *MockInterfaceTimedate) RTCTimeUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
