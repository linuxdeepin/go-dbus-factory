// Code generated by "./generator ./com.deepin.devicemanager"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package devicemanager

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockDeviceManager struct {
	MockInterfaceDevicemanager // interface com.deepin.devicemanager
	proxy.MockObject
}

type MockInterfaceDevicemanager struct {
	mock.Mock
}

// method getInfo

func (v *MockInterfaceDevicemanager) GoGetInfo(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevicemanager) GetInfo(flags dbus.Flags, arg1 string) (string, error) {
	mockArgs := v.Called(flags, arg1)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method execCmd

func (v *MockInterfaceDevicemanager) GoExecCmd(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevicemanager) ExecCmd(flags dbus.Flags, arg1 string) (string, error) {
	mockArgs := v.Called(flags, arg1)

	return mockArgs.String(0), mockArgs.Error(1)
}
