// Code generated by "./generator ./session/com.deepin.dde.lockfront"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package lockfront

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockLockFront struct {
	MockInterfaceLockfront // interface com.deepin.dde.lockFront
	proxy.MockObject
}

type MockInterfaceLockfront struct {
	mock.Mock
}

// method Show

func (v *MockInterfaceLockfront) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLockfront) Show(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowUserList

func (v *MockInterfaceLockfront) GoShowUserList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLockfront) ShowUserList(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowAuth

func (v *MockInterfaceLockfront) GoShowAuth(flags dbus.Flags, ch chan *dbus.Call, active bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, active)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLockfront) ShowAuth(flags dbus.Flags, active bool) error {
	mockArgs := v.Called(flags, active)

	return mockArgs.Error(0)
}

// method Suspend

func (v *MockInterfaceLockfront) GoSuspend(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enable)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLockfront) Suspend(flags dbus.Flags, enable bool) error {
	mockArgs := v.Called(flags, enable)

	return mockArgs.Error(0)
}

// method Hibernate

func (v *MockInterfaceLockfront) GoHibernate(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enable)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLockfront) Hibernate(flags dbus.Flags, enable bool) error {
	mockArgs := v.Called(flags, enable)

	return mockArgs.Error(0)
}

// signal ChangKey

func (v *MockInterfaceLockfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
