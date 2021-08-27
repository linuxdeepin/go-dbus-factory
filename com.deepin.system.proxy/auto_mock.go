// Code generated by "./generator ./com.deepin.system.proxy"; DO NOT EDIT.

package proxy

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockApp struct {
	MockInterfaceApp // interface com.deepin.system.proxy.App
	proxy.MockObject
}

type MockInterfaceApp struct {
	mock.Mock
}

// method AddProxy

func (v *MockInterfaceApp) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	mockArgs := v.Called(flags, ch, proto, name, proxy)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	mockArgs := v.Called(flags, proto, name, proxy)

	return mockArgs.Error(0)
}

// method AddProxyApps

func (v *MockInterfaceApp) GoAddProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, app)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) AddProxyApps(flags dbus.Flags, app []string) error {
	mockArgs := v.Called(flags, app)

	return mockArgs.Error(0)
}

// method ClearProxy

func (v *MockInterfaceApp) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) ClearProxy(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method DelProxyApps

func (v *MockInterfaceApp) GoDelProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, app)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) DelProxyApps(flags dbus.Flags, app []string) error {
	mockArgs := v.Called(flags, app)

	return mockArgs.Error(0)
}

// method GetProxy

func (v *MockInterfaceApp) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) GetProxy(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetProxies

func (v *MockInterfaceApp) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxies)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	mockArgs := v.Called(flags, proxies)

	return mockArgs.Error(0)
}

// method GetCGroups

func (v *MockInterfaceApp) GoGetCGroups(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) GetCGroups(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method AddProc

func (v *MockInterfaceApp) GoAddProc(flags dbus.Flags, ch chan *dbus.Call, pid int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, pid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) AddProc(flags dbus.Flags, pid int32) error {
	mockArgs := v.Called(flags, pid)

	return mockArgs.Error(0)
}

// method StartProxy

func (v *MockInterfaceApp) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, proto, name, udp)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	mockArgs := v.Called(flags, proto, name, udp)

	return mockArgs.Error(0)
}

// method StopProxy

func (v *MockInterfaceApp) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApp) StopProxy(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// signal Proxy

func (v *MockInterfaceApp) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockGlobal struct {
	MockInterfaceGlobal // interface com.deepin.system.proxy.Global
	proxy.MockObject
}

type MockInterfaceGlobal struct {
	mock.Mock
}

// method AddProxy

func (v *MockInterfaceGlobal) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	mockArgs := v.Called(flags, ch, proto, name, proxy)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	mockArgs := v.Called(flags, proto, name, proxy)

	return mockArgs.Error(0)
}

// method ClearProxy

func (v *MockInterfaceGlobal) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) ClearProxy(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetProxy

func (v *MockInterfaceGlobal) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) GetProxy(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IgnoreProxyApps

func (v *MockInterfaceGlobal) GoIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, app)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) IgnoreProxyApps(flags dbus.Flags, app []string) error {
	mockArgs := v.Called(flags, app)

	return mockArgs.Error(0)
}

// method SetProxies

func (v *MockInterfaceGlobal) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxies)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	mockArgs := v.Called(flags, proxies)

	return mockArgs.Error(0)
}

// method StartProxy

func (v *MockInterfaceGlobal) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, proto, name, udp)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	mockArgs := v.Called(flags, proto, name, udp)

	return mockArgs.Error(0)
}

// method StopProxy

func (v *MockInterfaceGlobal) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) StopProxy(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method UnIgnoreProxyApps

func (v *MockInterfaceGlobal) GoUnIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, app)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGlobal) UnIgnoreProxyApps(flags dbus.Flags, app []string) error {
	mockArgs := v.Called(flags, app)

	return mockArgs.Error(0)
}

// signal Proxy

func (v *MockInterfaceGlobal) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property IgnoreApp as

func (v *MockInterfaceGlobal) IgnoreApp() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
