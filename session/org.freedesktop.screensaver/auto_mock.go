// Code generated by "./generator ./session/org.freedesktop.screensaver"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package screensaver

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockScreenSaver struct {
	MockInterfaceScreenSaver // interface org.freedesktop.ScreenSaver
	proxy.MockObject
}

type MockInterfaceScreenSaver struct {
	mock.Mock
}

// method Inhibit

func (v *MockInterfaceScreenSaver) GoInhibit(flags dbus.Flags, ch chan *dbus.Call, name string, reason string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name, reason)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceScreenSaver) Inhibit(flags dbus.Flags, name string, reason string) (uint32, error) {
	mockArgs := v.Called(flags, name, reason)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetTimeout

func (v *MockInterfaceScreenSaver) GoSetTimeout(flags dbus.Flags, ch chan *dbus.Call, seconds uint32, interval uint32, blank bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, seconds, interval, blank)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceScreenSaver) SetTimeout(flags dbus.Flags, seconds uint32, interval uint32, blank bool) error {
	mockArgs := v.Called(flags, seconds, interval, blank)

	return mockArgs.Error(0)
}

// method SimulateUserActivity

func (v *MockInterfaceScreenSaver) GoSimulateUserActivity(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceScreenSaver) SimulateUserActivity(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method UnInhibit

func (v *MockInterfaceScreenSaver) GoUnInhibit(flags dbus.Flags, ch chan *dbus.Call, cookie uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, cookie)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceScreenSaver) UnInhibit(flags dbus.Flags, cookie uint32) error {
	mockArgs := v.Called(flags, cookie)

	return mockArgs.Error(0)
}

// signal IdleOn

func (v *MockInterfaceScreenSaver) ConnectIdleOn(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CycleActive

func (v *MockInterfaceScreenSaver) ConnectCycleActive(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal IdleOff

func (v *MockInterfaceScreenSaver) ConnectIdleOff(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
