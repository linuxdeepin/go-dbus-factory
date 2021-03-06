// Code generated by "./generator ./com.deepin.api.xeventmonitor"; DO NOT EDIT.

package xeventmonitor

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
)

type MockXEventMonitor struct {
	mockInterfaceXEventMonitor // interface com.deepin.api.XEventMonitor
}

type mockInterfaceXEventMonitor struct {
	mock.Mock
}

// method RegisterArea

func (v *mockInterfaceXEventMonitor) GoRegisterArea(flags dbus.Flags, ch chan *dbus.Call, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, x1, y1, x2, y2, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceXEventMonitor) RegisterArea(flags dbus.Flags, x1 int32, y1 int32, x2 int32, y2 int32, flag int32) (string, error) {
	mockArgs := v.Called(flags, x1, y1, x2, y2, flag)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RegisterAreas

func (v *mockInterfaceXEventMonitor) GoRegisterAreas(flags dbus.Flags, ch chan *dbus.Call, areas []CoordinateRange, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, areas, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceXEventMonitor) RegisterAreas(flags dbus.Flags, areas []CoordinateRange, flag int32) (string, error) {
	mockArgs := v.Called(flags, areas, flag)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RegisterFullScreen

func (v *mockInterfaceXEventMonitor) GoRegisterFullScreen(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceXEventMonitor) RegisterFullScreen(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method UnregisterArea

func (v *mockInterfaceXEventMonitor) GoUnregisterArea(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceXEventMonitor) UnregisterArea(flags dbus.Flags, id string) (bool, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CancelAllArea

func (v *mockInterfaceXEventMonitor) ConnectCancelAllArea(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorInto

func (v *mockInterfaceXEventMonitor) ConnectCursorInto(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorOut

func (v *mockInterfaceXEventMonitor) ConnectCursorOut(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CursorMove

func (v *mockInterfaceXEventMonitor) ConnectCursorMove(cb func(x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ButtonPress

func (v *mockInterfaceXEventMonitor) ConnectButtonPress(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ButtonRelease

func (v *mockInterfaceXEventMonitor) ConnectButtonRelease(cb func(button int32, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal KeyPress

func (v *mockInterfaceXEventMonitor) ConnectKeyPress(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal KeyRelease

func (v *mockInterfaceXEventMonitor) ConnectKeyRelease(cb func(key string, x int32, y int32, id string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
