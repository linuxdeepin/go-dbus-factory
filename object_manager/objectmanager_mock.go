package object_manager

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/stretchr/testify/mock"
)

type MockInterfaceObjectManager struct {
	mock.Mock
}

// method GetManagedObjects

func (v *MockInterfaceObjectManager) GetManagedObjects(flags dbus.Flags) (map[dbus.ObjectPath]map[string]map[string]dbus.Variant, error) {
	args := v.Called(flags)

	ret0, ok := args.Get(0).(map[dbus.ObjectPath]map[string]map[string]dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return ret0, args.Error(1)
}

// signal InterfacesAdded

func (v *MockInterfaceObjectManager) ConnectInterfacesAdded(cb func(object_path dbus.ObjectPath, interfaces_and_properties map[string]map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	args := v.Called(cb)

	ret0, ok := args.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return ret0, args.Error(1)
}

// signal InterfacesRemoved

func (v *MockInterfaceObjectManager) ConnectInterfacesRemoved(cb func(object_path dbus.ObjectPath, interfaces []string)) (dbusutil.SignalHandlerId, error) {
	args := v.Called(cb)

	ret0, ok := args.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return ret0, args.Error(1)
}
