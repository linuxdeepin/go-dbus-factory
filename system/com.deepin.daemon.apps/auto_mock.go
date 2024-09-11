// Code generated by "./generator ./system/com.deepin.daemon.apps"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package apps

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockApps struct {
	MockInterfaceDesktopFileWatcher // interface com.deepin.daemon.Apps.DesktopFileWatcher
	MockInterfaceLaunchedRecorder   // interface com.deepin.daemon.Apps.LaunchedRecorder
	proxy.MockObject
}

type MockInterfaceDesktopFileWatcher struct {
	mock.Mock
}

// signal Event

func (v *MockInterfaceDesktopFileWatcher) ConnectEvent(cb func(name string, op uint32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockInterfaceLaunchedRecorder struct {
	mock.Mock
}

// method GetNew

func (v *MockInterfaceLaunchedRecorder) GoGetNew(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLaunchedRecorder) GetNew(flags dbus.Flags) (map[string][]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(map[string][]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method MarkLaunched

func (v *MockInterfaceLaunchedRecorder) GoMarkLaunched(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLaunchedRecorder) MarkLaunched(flags dbus.Flags, desktopFile string) error {
	mockArgs := v.Called(flags, desktopFile)

	return mockArgs.Error(0)
}

// method UninstallHints

func (v *MockInterfaceLaunchedRecorder) GoUninstallHints(flags dbus.Flags, ch chan *dbus.Call, desktopFiles []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFiles)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLaunchedRecorder) UninstallHints(flags dbus.Flags, desktopFiles []string) error {
	mockArgs := v.Called(flags, desktopFiles)

	return mockArgs.Error(0)
}

// method WatchDirs

func (v *MockInterfaceLaunchedRecorder) GoWatchDirs(flags dbus.Flags, ch chan *dbus.Call, dirs []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, dirs)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLaunchedRecorder) WatchDirs(flags dbus.Flags, dirs []string) error {
	mockArgs := v.Called(flags, dirs)

	return mockArgs.Error(0)
}

// signal Launched

func (v *MockInterfaceLaunchedRecorder) ConnectLaunched(cb func(file string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal StatusSaved

func (v *MockInterfaceLaunchedRecorder) ConnectStatusSaved(cb func(root string, file string, ok bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ServiceRestarted

func (v *MockInterfaceLaunchedRecorder) ConnectServiceRestarted(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
