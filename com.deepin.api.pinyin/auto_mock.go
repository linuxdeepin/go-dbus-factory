// Code generated by "./generator ./com.deepin.api.pinyin"; DO NOT EDIT.

package pinyin

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type MockPinyin struct {
	MockInterfacePinyin // interface com.deepin.api.Pinyin
	proxy.MockObject
}

type MockInterfacePinyin struct {
	mock.Mock
}

// method Query

func (v *MockInterfacePinyin) GoQuery(flags dbus.Flags, ch chan *dbus.Call, hans string) *dbus.Call {
	mockArgs := v.Called(flags, ch, hans)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePinyin) Query(flags dbus.Flags, hans string) ([]string, error) {
	mockArgs := v.Called(flags, hans)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method QueryList

func (v *MockInterfacePinyin) GoQueryList(flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, hansList)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePinyin) QueryList(flags dbus.Flags, hansList []string) (string, error) {
	mockArgs := v.Called(flags, hansList)

	return mockArgs.String(0), mockArgs.Error(1)
}
