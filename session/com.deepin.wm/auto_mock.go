// Code generated by "./generator ./session/com.deepin.wm"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package wm

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockWm struct {
	MockInterfaceWm // interface com.deepin.wm
	proxy.MockObject
}

type MockInterfaceWm struct {
	mock.Mock
}

// method SwitchApplication

func (v *MockInterfaceWm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, backward)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SwitchApplication(flags dbus.Flags, backward bool) error {
	mockArgs := v.Called(flags, backward)

	return mockArgs.Error(0)
}

// method TileActiveWindow

func (v *MockInterfaceWm) GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, side)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) TileActiveWindow(flags dbus.Flags, side uint32) error {
	mockArgs := v.Called(flags, side)

	return mockArgs.Error(0)
}

// method BeginToMoveActiveWindow

func (v *MockInterfaceWm) GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) BeginToMoveActiveWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ToggleActiveWindowMaximize

func (v *MockInterfaceWm) GoToggleActiveWindowMaximize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ToggleActiveWindowMaximize(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method MinimizeActiveWindow

func (v *MockInterfaceWm) GoMinimizeActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) MinimizeActiveWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowWorkspace

func (v *MockInterfaceWm) GoShowWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ShowWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowWindow

func (v *MockInterfaceWm) GoShowWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ShowWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowAllWindow

func (v *MockInterfaceWm) GoShowAllWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ShowAllWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ClearMoveStatus

func (v *MockInterfaceWm) GoClearMoveStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ClearMoveStatus(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method TouchToMove

func (v *MockInterfaceWm) GoTouchToMove(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, x, y)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) TouchToMove(flags dbus.Flags, x int32, y int32) error {
	mockArgs := v.Called(flags, x, y)

	return mockArgs.Error(0)
}

// method PerformAction

func (v *MockInterfaceWm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, type0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) PerformAction(flags dbus.Flags, type0 int32) error {
	mockArgs := v.Called(flags, type0)

	return mockArgs.Error(0)
}

// method PreviewWindow

func (v *MockInterfaceWm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, xid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	mockArgs := v.Called(flags, xid)

	return mockArgs.Error(0)
}

// method CancelPreviewWindow

func (v *MockInterfaceWm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) CancelPreviewWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspaceBackground

func (v *MockInterfaceWm) GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetCurrentWorkspaceBackground(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method SetCurrentWorkspaceBackground

func (v *MockInterfaceWm) GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	mockArgs := v.Called(flags, uri)

	return mockArgs.Error(0)
}

// method GetWorkspaceBackground

func (v *MockInterfaceWm) GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, index)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetWorkspaceBackground(flags dbus.Flags, index int32) (string, error) {
	mockArgs := v.Called(flags, index)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method SetWorkspaceBackground

func (v *MockInterfaceWm) GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error {
	mockArgs := v.Called(flags, index, uri)

	return mockArgs.Error(0)
}

// method SetTransientBackground

func (v *MockInterfaceWm) GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetTransientBackground(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspaceBackgroundForMonitor

func (v *MockInterfaceWm) GoGetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, strMonitorName string) (string, error) {
	mockArgs := v.Called(flags, strMonitorName)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method SetCurrentWorkspaceBackgroundForMonitor

func (v *MockInterfaceWm) GoSetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	mockArgs := v.Called(flags, uri, strMonitorName)

	return mockArgs.Error(0)
}

// method GetWorkspaceBackgroundForMonitor

func (v *MockInterfaceWm) GoGetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string) (string, error) {
	mockArgs := v.Called(flags, index, strMonitorName)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method SetWorkspaceBackgroundForMonitor

func (v *MockInterfaceWm) GoSetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, strMonitorName, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string, uri string) error {
	mockArgs := v.Called(flags, index, strMonitorName, uri)

	return mockArgs.Error(0)
}

// method SetTransientBackgroundForMonitor

func (v *MockInterfaceWm) GoSetTransientBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetTransientBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	mockArgs := v.Called(flags, uri, strMonitorName)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspace

func (v *MockInterfaceWm) GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetCurrentWorkspace(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method WorkspaceCount

func (v *MockInterfaceWm) GoWorkspaceCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) WorkspaceCount(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetCurrentWorkspace

func (v *MockInterfaceWm) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, index)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetCurrentWorkspace(flags dbus.Flags, index int32) error {
	mockArgs := v.Called(flags, index)

	return mockArgs.Error(0)
}

// method PreviousWorkspace

func (v *MockInterfaceWm) GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) PreviousWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method NextWorkspace

func (v *MockInterfaceWm) GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) NextWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetAllAccels

func (v *MockInterfaceWm) GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetAllAccels(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetAccel

func (v *MockInterfaceWm) GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetAccel(flags dbus.Flags, id string) ([]string, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetDefaultAccel

func (v *MockInterfaceWm) GoGetDefaultAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetDefaultAccel(flags dbus.Flags, id string) ([]string, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetAccel

func (v *MockInterfaceWm) GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	mockArgs := v.Called(flags, ch, data)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetAccel(flags dbus.Flags, data string) (bool, error) {
	mockArgs := v.Called(flags, data)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RemoveAccel

func (v *MockInterfaceWm) GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) RemoveAccel(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method SetDecorationTheme

func (v *MockInterfaceWm) GoSetDecorationTheme(flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, themeType, themeName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetDecorationTheme(flags dbus.Flags, themeType string, themeName string) error {
	mockArgs := v.Called(flags, themeType, themeName)

	return mockArgs.Error(0)
}

// method SetDecorationDeepinTheme

func (v *MockInterfaceWm) GoSetDecorationDeepinTheme(flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, deepinThemeName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SetDecorationDeepinTheme(flags dbus.Flags, deepinThemeName string) error {
	mockArgs := v.Called(flags, deepinThemeName)

	return mockArgs.Error(0)
}

// method ChangeCurrentWorkspaceBackground

func (v *MockInterfaceWm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	mockArgs := v.Called(flags, uri)

	return mockArgs.Error(0)
}

// method SwitchToWorkspace

func (v *MockInterfaceWm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, backward)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	mockArgs := v.Called(flags, backward)

	return mockArgs.Error(0)
}

// method PresentWindows

func (v *MockInterfaceWm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, xids)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	mockArgs := v.Called(flags, xids)

	return mockArgs.Error(0)
}

// method EnableZoneDetected

func (v *MockInterfaceWm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) EnableZoneDetected(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method GetIsShowDesktop

func (v *MockInterfaceWm) GoGetIsShowDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetIsShowDesktop(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method GetMultiTaskingStatus

func (v *MockInterfaceWm) GoGetMultiTaskingStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWm) GetMultiTaskingStatus(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// signal WorkspaceBackgroundChanged

func (v *MockInterfaceWm) ConnectWorkspaceBackgroundChanged(cb func(index int32, newUri string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WorkspaceBackgroundChangedForMonitor

func (v *MockInterfaceWm) ConnectWorkspaceBackgroundChangedForMonitor(cb func(index int32, monitor string, newUri string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal compositingEnabledChanged

func (v *MockInterfaceWm) ConnectCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal wmCompositingEnabledChanged

func (v *MockInterfaceWm) ConnectWmCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal workspaceCountChanged

func (v *MockInterfaceWm) ConnectWorkspaceCountChanged(cb func(count int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WorkspaceSwitched

func (v *MockInterfaceWm) ConnectWorkspaceSwitched(cb func(from int32, to int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property compositingEnabled b

func (v *MockInterfaceWm) CompositingEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property compositingPossible b

func (v *MockInterfaceWm) CompositingPossible() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property compositingAllowSwitch b

func (v *MockInterfaceWm) CompositingAllowSwitch() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property zoneEnabled b

func (v *MockInterfaceWm) ZoneEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property cursorTheme s

func (v *MockInterfaceWm) CursorTheme() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property cursorSize i

func (v *MockInterfaceWm) CursorSize() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
