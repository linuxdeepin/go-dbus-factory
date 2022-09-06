// Code generated by "./generator ./com.deepin.daemon.keyevent"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package keyevent

import "fmt"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockKeyEvent struct {
	MockInterfaceKeyEvent // interface com.deepin.daemon.KeyEvent
	proxy.MockObject
}

type MockInterfaceKeyEvent struct {
	mock.Mock
}

// signal KeyEvent

func (v *MockInterfaceKeyEvent) ConnectKeyEvent(cb func(keycode uint32, pressed bool, ctrlPressed bool, shiftPressed bool, altPressed bool, superPressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
