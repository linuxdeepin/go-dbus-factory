// Code generated by "./generator ./com.deepin.daemon.authenticate.fingerprint"; DO NOT EDIT.

package fingerprint

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockDevice struct {
	mockInterfaceDevice // interface com.deepin.daemon.Authenticate.Fingerprint.Device
}

type mockInterfaceDevice struct {
	mock.Mock
}

func (v *mockInterfaceDevice) SetInterfaceName_(string) {
}

// method Claim

func (v *mockInterfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	mockArgs := v.Called(flags, userId, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *mockInterfaceDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	mockArgs := v.Called(flags, userId)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *mockInterfaceDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	mockArgs := v.Called(flags, userId, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *mockInterfaceDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) Enroll(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *mockInterfaceDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	mockArgs := v.Called(flags, userId)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFinger

func (v *mockInterfaceDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	mockArgs := v.Called(flags, userId, finger, newName)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *mockInterfaceDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *mockInterfaceDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *mockInterfaceDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDevice) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *mockInterfaceDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *mockInterfaceDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *mockInterfaceDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *mockInterfaceDevice) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *mockInterfaceDevice) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type i

func (v *mockInterfaceDevice) Type() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *mockInterfaceDevice) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockCommonDevice struct {
	mockInterfaceCommonDevice // interface com.deepin.daemon.Authenticate.Fingerprint.CommonDevice
}

type mockInterfaceCommonDevice struct {
	mock.Mock
}

func (v *mockInterfaceCommonDevice) SetInterfaceName_(string) {
}

// method Claim

func (v *mockInterfaceCommonDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	mockArgs := v.Called(flags, userId, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *mockInterfaceCommonDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	mockArgs := v.Called(flags, userId)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *mockInterfaceCommonDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	mockArgs := v.Called(flags, userId, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *mockInterfaceCommonDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) Enroll(flags dbus.Flags, username string, finger string) error {
	mockArgs := v.Called(flags, username, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *mockInterfaceCommonDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	mockArgs := v.Called(flags, userId)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFinger

func (v *mockInterfaceCommonDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	mockArgs := v.Called(flags, userId, finger, newName)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *mockInterfaceCommonDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *mockInterfaceCommonDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *mockInterfaceCommonDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceCommonDevice) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *mockInterfaceCommonDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *mockInterfaceCommonDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *mockInterfaceCommonDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *mockInterfaceCommonDevice) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *mockInterfaceCommonDevice) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type i

func (v *mockInterfaceCommonDevice) Type() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *mockInterfaceCommonDevice) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
