// Code generated by "./generator ./session/org.deepin.dde.startmanager1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package startmanager1

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockStartManager struct {
	MockInterfaceStartManager // interface org.deepin.dde.StartManager1
	proxy.MockObject
}

type MockInterfaceStartManager struct {
	mock.Mock
}

// method AddAutostart

func (v *MockInterfaceStartManager) GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) AddAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method AutostartList

func (v *MockInterfaceStartManager) GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) AutostartList(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DumpMemRecord

func (v *MockInterfaceStartManager) GoDumpMemRecord(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) DumpMemRecord(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetApps

func (v *MockInterfaceStartManager) GoGetApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) GetApps(flags dbus.Flags) (map[uint32]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(map[uint32]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsAutostart

func (v *MockInterfaceStartManager) GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) IsAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsMemSufficient

func (v *MockInterfaceStartManager) GoIsMemSufficient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) IsMemSufficient(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method Launch

func (v *MockInterfaceStartManager) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) Launch(flags dbus.Flags, arg0 string) (bool, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method LaunchApp

func (v *MockInterfaceStartManager) GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) LaunchApp(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2)

	return mockArgs.Error(0)
}

// method LaunchAppAction

func (v *MockInterfaceStartManager) GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) LaunchAppAction(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2)

	return mockArgs.Error(0)
}

// method LaunchAppWithOptions

func (v *MockInterfaceStartManager) GoLaunchAppWithOptions(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2, arg3)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) LaunchAppWithOptions(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2, arg3)

	return mockArgs.Error(0)
}

// method LaunchWithTimestamp

func (v *MockInterfaceStartManager) GoLaunchWithTimestamp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) LaunchWithTimestamp(flags dbus.Flags, arg0 string, arg1 uint32) (bool, error) {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RemoveAutostart

func (v *MockInterfaceStartManager) GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) RemoveAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RunCommand

func (v *MockInterfaceStartManager) GoRunCommand(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) RunCommand(flags dbus.Flags, arg0 string, arg1 []string) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method TryAgain

func (v *MockInterfaceStartManager) GoTryAgain(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceStartManager) TryAgain(flags dbus.Flags, arg0 bool) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// signal AutostartChanged

func (v *MockInterfaceStartManager) ConnectAutostartChanged(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property NeededMemory t

func (v *MockInterfaceStartManager) NeededMemory() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
