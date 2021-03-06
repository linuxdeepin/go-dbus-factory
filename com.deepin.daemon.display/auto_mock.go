// Code generated by "./generator ./com.deepin.daemon.display"; DO NOT EDIT.

package display

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockDisplay struct {
	mockInterfaceDisplay // interface com.deepin.daemon.Display
}

type mockInterfaceDisplay struct {
	mock.Mock
}

// method ApplyChanges

func (v *mockInterfaceDisplay) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ApplyChanges(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method AssociateTouch

func (v *mockInterfaceDisplay) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) AssociateTouch(flags dbus.Flags, arg0 string, arg1 string) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method ChangeBrightness

func (v *mockInterfaceDisplay) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ChangeBrightness(flags dbus.Flags, arg0 bool) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method DeleteCustomMode

func (v *mockInterfaceDisplay) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) DeleteCustomMode(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method GetBrightness

func (v *mockInterfaceDisplay) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) GetBrightness(flags dbus.Flags) (map[string]float64, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(map[string]float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListOutputNames

func (v *mockInterfaceDisplay) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ListOutputNames(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListOutputsCommonModes

func (v *mockInterfaceDisplay) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ListOutputsCommonModes(flags dbus.Flags) ([]ModeInfo, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]ModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ModifyConfigName

func (v *mockInterfaceDisplay) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ModifyConfigName(flags dbus.Flags, arg0 string, arg1 string) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method RefreshBrightness

func (v *mockInterfaceDisplay) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) RefreshBrightness(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Reset

func (v *mockInterfaceDisplay) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) Reset(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ResetChanges

func (v *mockInterfaceDisplay) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) ResetChanges(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Save

func (v *mockInterfaceDisplay) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) Save(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SetAndSaveBrightness

func (v *mockInterfaceDisplay) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) SetAndSaveBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetBrightness

func (v *mockInterfaceDisplay) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) SetBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetPrimary

func (v *mockInterfaceDisplay) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) SetPrimary(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method SwitchMode

func (v *mockInterfaceDisplay) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint8, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceDisplay) SwitchMode(flags dbus.Flags, arg0 uint8, arg1 string) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// property HasChanged b

func (v *mockInterfaceDisplay) HasChanged() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DisplayMode y

func (v *mockInterfaceDisplay) DisplayMode() proxy.PropByte {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropByte)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ScreenWidth q

func (v *mockInterfaceDisplay) ScreenWidth() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ScreenHeight q

func (v *mockInterfaceDisplay) ScreenHeight() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Primary s

func (v *mockInterfaceDisplay) Primary() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentCustomId s

func (v *mockInterfaceDisplay) CurrentCustomId() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CustomIdList as

func (v *mockInterfaceDisplay) CustomIdList() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayPrimaryRect struct {
	mock.Mock
}

func (p MockPropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(Rectangle)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayPrimaryRect) Set(flags dbus.Flags, value Rectangle) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property PrimaryRect (nnqq)

func (v *mockInterfaceDisplay) PrimaryRect() PropDisplayPrimaryRect {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayPrimaryRect)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Monitors ao

func (v *mockInterfaceDisplay) Monitors() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayBrightness struct {
	mock.Mock
}

func (p MockPropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[string]float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayBrightness) Set(flags dbus.Flags, value map[string]float64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property Brightness a{sd}

func (v *mockInterfaceDisplay) Brightness() PropDisplayBrightness {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayBrightness)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayTouchMap struct {
	mock.Mock
}

func (p MockPropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[string]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayTouchMap) Set(flags dbus.Flags, value map[string]string) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property TouchMap a{ss}

func (v *mockInterfaceDisplay) TouchMap() PropDisplayTouchMap {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayTouchMap)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropTouchscreens struct {
	mock.Mock
}

func (p MockPropTouchscreens) Get(flags dbus.Flags) (value []Touchscreen, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]Touchscreen)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropTouchscreens) Set(flags dbus.Flags, value []Touchscreen) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropTouchscreens) ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property Touchscreens a(isss)

func (v *mockInterfaceDisplay) Touchscreens() PropTouchscreens {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropTouchscreens)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property MaxBacklightBrightness u

func (v *mockInterfaceDisplay) MaxBacklightBrightness() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ColorTemperatureMode i

func (v *mockInterfaceDisplay) ColorTemperatureMode() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ColorTemperatureManual i

func (v *mockInterfaceDisplay) ColorTemperatureManual() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockMonitor struct {
	mockInterfaceMonitor // interface com.deepin.daemon.Display.Monitor
}

type mockInterfaceMonitor struct {
	mock.Mock
}

// method Enable

func (v *mockInterfaceMonitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) Enable(flags dbus.Flags, arg0 bool) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method SetMode

func (v *mockInterfaceMonitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetMode(flags dbus.Flags, arg0 uint32) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method SetModeBySize

func (v *mockInterfaceMonitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16, arg1 uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetModeBySize(flags dbus.Flags, arg0 uint16, arg1 uint16) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetPosition

func (v *mockInterfaceMonitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, arg0 int16, arg1 int16) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetPosition(flags dbus.Flags, arg0 int16, arg1 int16) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method SetReflect

func (v *mockInterfaceMonitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetReflect(flags dbus.Flags, arg0 uint16) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method SetRefreshRate

func (v *mockInterfaceMonitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetRefreshRate(flags dbus.Flags, arg0 float64) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method SetRotation

func (v *mockInterfaceMonitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceMonitor) SetRotation(flags dbus.Flags, arg0 uint16) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// property Name s

func (v *mockInterfaceMonitor) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Enabled b

func (v *mockInterfaceMonitor) Enabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Connected b

func (v *mockInterfaceMonitor) Connected() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property X n

func (v *mockInterfaceMonitor) X() proxy.PropInt16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Y n

func (v *mockInterfaceMonitor) Y() proxy.PropInt16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Width q

func (v *mockInterfaceMonitor) Width() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Height q

func (v *mockInterfaceMonitor) Height() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Rotation q

func (v *mockInterfaceMonitor) Rotation() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Reflect q

func (v *mockInterfaceMonitor) Reflect() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property RefreshRate d

func (v *mockInterfaceMonitor) RefreshRate() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Rotations aq

func (v *mockInterfaceMonitor) Rotations() proxy.PropUint16Array {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16Array)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Reflects aq

func (v *mockInterfaceMonitor) Reflects() proxy.PropUint16Array {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16Array)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BestMode (uqqd)

func (v *mockInterfaceMonitor) BestMode() PropModeInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentMode (uqqd)

func (v *mockInterfaceMonitor) CurrentMode() PropModeInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Modes a(uqqd)

func (v *mockInterfaceMonitor) Modes() PropModeInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PreferredModes a(uqqd)

func (v *mockInterfaceMonitor) PreferredModes() PropModeInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropModeInfo struct {
	mock.Mock
}

func (p MockPropModeInfo) Get(flags dbus.Flags) (value ModeInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(ModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropModeInfo) Set(flags dbus.Flags, value ModeInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropModeInfo) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type MockPropModeInfoSlice struct {
	mock.Mock
}

func (p MockPropModeInfoSlice) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]ModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropModeInfoSlice) Set(flags dbus.Flags, value []ModeInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropModeInfoSlice) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}
