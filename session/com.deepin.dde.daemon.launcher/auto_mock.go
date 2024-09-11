// Code generated by "./generator ./session/com.deepin.dde.daemon.launcher"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package launcher

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockLauncher struct {
	MockInterfaceLauncher // interface com.deepin.dde.daemon.Launcher
	proxy.MockObject
}

type MockInterfaceLauncher struct {
	mock.Mock
}

// method GetAllItemInfos

func (v *MockInterfaceLauncher) GoGetAllItemInfos(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) GetAllItemInfos(flags dbus.Flags) ([]ItemInfo, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]ItemInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAllNewInstalledApps

func (v *MockInterfaceLauncher) GoGetAllNewInstalledApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) GetAllNewInstalledApps(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetDisableScaling

func (v *MockInterfaceLauncher) GoGetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) GetDisableScaling(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method GetItemInfo

func (v *MockInterfaceLauncher) GoGetItemInfo(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) GetItemInfo(flags dbus.Flags, id string) (ItemInfo, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).(ItemInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetUseProxy

func (v *MockInterfaceLauncher) GoGetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) GetUseProxy(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsItemOnDesktop

func (v *MockInterfaceLauncher) GoIsItemOnDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) IsItemOnDesktop(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method MarkLaunched

func (v *MockInterfaceLauncher) GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) MarkLaunched(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method RequestRemoveFromDesktop

func (v *MockInterfaceLauncher) GoRequestRemoveFromDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) RequestRemoveFromDesktop(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RequestSendToDesktop

func (v *MockInterfaceLauncher) GoRequestSendToDesktop(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) RequestSendToDesktop(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RequestUninstall

func (v *MockInterfaceLauncher) GoRequestUninstall(flags dbus.Flags, ch chan *dbus.Call, id string, purge bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, id, purge)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) RequestUninstall(flags dbus.Flags, id string, purge bool) error {
	mockArgs := v.Called(flags, id, purge)

	return mockArgs.Error(0)
}

// method Search

func (v *MockInterfaceLauncher) GoSearch(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) Search(flags dbus.Flags, key string) error {
	mockArgs := v.Called(flags, key)

	return mockArgs.Error(0)
}

// method SetDisableScaling

func (v *MockInterfaceLauncher) GoSetDisableScaling(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, id, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) SetDisableScaling(flags dbus.Flags, id string, value bool) error {
	mockArgs := v.Called(flags, id, value)

	return mockArgs.Error(0)
}

// method SetUseProxy

func (v *MockInterfaceLauncher) GoSetUseProxy(flags dbus.Flags, ch chan *dbus.Call, id string, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, id, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLauncher) SetUseProxy(flags dbus.Flags, id string, value bool) error {
	mockArgs := v.Called(flags, id, value)

	return mockArgs.Error(0)
}

// signal SearchDone

func (v *MockInterfaceLauncher) ConnectSearchDone(cb func(apps []string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ItemChanged

func (v *MockInterfaceLauncher) ConnectItemChanged(cb func(status string, itemInfo ItemInfo, categoryID int64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal NewAppLaunched

func (v *MockInterfaceLauncher) ConnectNewAppLaunched(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UninstallSuccess

func (v *MockInterfaceLauncher) ConnectUninstallSuccess(cb func(appID string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UninstallFailed

func (v *MockInterfaceLauncher) ConnectUninstallFailed(cb func(appId string, errMsg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Fullscreen b

func (v *MockInterfaceLauncher) Fullscreen() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DisplayMode i

func (v *MockInterfaceLauncher) DisplayMode() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
