// Code generated by "./generator ./com.deepin.api.xeventmonitor"; DO NOT EDIT.

package xeventmonitor

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type MockXEventMonitor struct {
	MockInterfaceXEventMonitor // interface com.deepin.api.XEventMonitor
	proxy.MockObject
}

type MockInterfaceXEventMonitor struct {
	mock.Mock
}

// method RegisterArea

func (v *MockInterfaceXEventMonitor) GoRegisterArea(flags dbus.Flags, ch chan *dbus.Call, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, x1, y1, x2, y2, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceXEventMonitor) RegisterArea(flags dbus.Flags, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) (string, error) {
	mockArgs := v.Called(flags, x1, y1, x2, y2, flag)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RegisterAreas

func (v *MockInterfaceXEventMonitor) GoRegisterAreas(flags dbus.Flags, ch chan *dbus.Call, areas []CoordinateRange, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, areas, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceXEventMonitor) RegisterAreas(flags dbus.Flags, areas []CoordinateRange, flag int32) (string, error) {
	mockArgs := v.Called(flags, areas, flag)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RegisterFullScreen

func (v *MockInterfaceXEventMonitor) GoRegisterFullScreen(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceXEventMonitor) RegisterFullScreen(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method UnregisterArea

func (v *MockInterfaceXEventMonitor) GoUnregisterArea(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceXEventMonitor) UnregisterArea(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// signal CancelAllArea

func (v *MockInterfaceXEventMonitor) ConnectCancelAllArea(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorInto

func (v *MockInterfaceXEventMonitor) ConnectCursorInto(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorOut

func (v *MockInterfaceXEventMonitor) ConnectCursorOut(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorMove

func (v *MockInterfaceXEventMonitor) ConnectCursorMove(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ButtonPress

func (v *MockInterfaceXEventMonitor) ConnectButtonPress(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ButtonRelease

func (v *MockInterfaceXEventMonitor) ConnectButtonRelease(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal KeyPress

func (v *MockInterfaceXEventMonitor) ConnectKeyPress(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal KeyRelease

func (v *MockInterfaceXEventMonitor) ConnectKeyRelease(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
