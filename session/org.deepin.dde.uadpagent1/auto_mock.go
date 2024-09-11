// Code generated by "./generator ./session/org.deepin.dde.uadpagent1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package uadpagent1

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockUadpAgent struct {
	MockInterfaceUadpagent // interface org.deepin.dde.UadpAgent1
	proxy.MockObject
}

type MockInterfaceUadpagent struct {
	mock.Mock
}

// method SetDataKey

func (v *MockInterfaceUadpagent) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string, dataKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, keyName, dataKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadpagent) SetDataKey(flags dbus.Flags, keyName string, dataKey string) error {
	mockArgs := v.Called(flags, keyName, dataKey)

	return mockArgs.Error(0)
}

// method GetDataKey

func (v *MockInterfaceUadpagent) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, keyName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadpagent) GetDataKey(flags dbus.Flags, keyName string) (string, error) {
	mockArgs := v.Called(flags, keyName)

	return mockArgs.String(0), mockArgs.Error(1)
}
