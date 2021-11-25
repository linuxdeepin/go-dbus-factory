// Code generated by "./generator ./com.deepin.system.ipwatchd"; DO NOT EDIT.

package ipwatchd

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type MockIPWatchD struct {
	MockInterfaceIpwatchd // interface com.deepin.system.IPWatchD
	proxy.MockObject
}

type MockInterfaceIpwatchd struct {
	mock.Mock
}

// method RequestIPConflictCheck

func (v *MockInterfaceIpwatchd) GoRequestIPConflictCheck(flags dbus.Flags, ch chan *dbus.Call, ip string, ifc string) *dbus.Call {
	mockArgs := v.Called(flags, ch, ip, ifc)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceIpwatchd) RequestIPConflictCheck(flags dbus.Flags, ip string, ifc string) (string, error) {
	mockArgs := v.Called(flags, ip, ifc)

	return mockArgs.String(0), mockArgs.Error(1)
}

// signal IPConflict

func (v *MockInterfaceIpwatchd) ConnectIPConflict(cb func(ip string, smac string, dmac string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
