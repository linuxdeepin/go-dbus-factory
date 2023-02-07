// Code generated by "./generator ./session/org.deepin.dde.bluetooth1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package bluetooth1

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockBluetooth struct {
	MockInterfaceBluetooth // interface org.deepin.dde.Bluetooth1
	proxy.MockObject
}

type MockInterfaceBluetooth struct {
	mock.Mock
}

// method DisconnectAllDevices

func (v *MockInterfaceBluetooth) GoDisconnectAllDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceBluetooth) DisconnectAllDevices(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}
