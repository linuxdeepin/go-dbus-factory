// Code generated by "./generator ./com.deepin.daemon.audio"; DO NOT EDIT.

package audio

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockAudio struct {
	mockInterfaceAudio // interface com.deepin.daemon.Audio
}

type mockInterfaceAudio struct {
	mock.Mock
}

// method SetDefaultSink

func (v *mockInterfaceAudio) GoSetDefaultSink(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAudio) SetDefaultSink(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method SetDefaultSource

func (v *mockInterfaceAudio) GoSetDefaultSource(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAudio) SetDefaultSource(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method SetPort

func (v *mockInterfaceAudio) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, cardId uint32, portName string, direction int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, cardId, portName, direction)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAudio) SetPort(flags dbus.Flags, cardId uint32, portName string, direction int32) error {
	mockArgs := v.Called(flags, cardId, portName, direction)

	return mockArgs.Error(0)
}

// method SetPortEnabled

func (v *mockInterfaceAudio) GoSetPortEnabled(flags dbus.Flags, ch chan *dbus.Call, cardId uint32, portName string, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, cardId, portName, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAudio) SetPortEnabled(flags dbus.Flags, cardId uint32, portName string, enabled bool) error {
	mockArgs := v.Called(flags, cardId, portName, enabled)

	return mockArgs.Error(0)
}

// method IsPortEnabled

func (v *mockInterfaceAudio) GoIsPortEnabled(flags dbus.Flags, ch chan *dbus.Call, cardId uint32, portName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, cardId, portName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAudio) IsPortEnabled(flags dbus.Flags, cardId uint32, portName string) (bool, error) {
	mockArgs := v.Called(flags, cardId, portName)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property MaxUIVolume d

func (v *mockInterfaceAudio) MaxUIVolume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SinkInputs ao

func (v *mockInterfaceAudio) SinkInputs() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultSink o

func (v *mockInterfaceAudio) DefaultSink() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultSource o

func (v *mockInterfaceAudio) DefaultSource() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Cards s

func (v *mockInterfaceAudio) Cards() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CardsWithoutUnavailable s

func (v *mockInterfaceAudio) CardsWithoutUnavailable() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockSink struct {
	mockInterfaceSink // interface com.deepin.daemon.Audio.Sink
}

type mockInterfaceSink struct {
	mock.Mock
}

// method GetMeter

func (v *mockInterfaceSink) GoGetMeter(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) GetMeter(flags dbus.Flags) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetBalance

func (v *mockInterfaceSink) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// method SetFade

func (v *mockInterfaceSink) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) SetFade(flags dbus.Flags, value float64) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetMute

func (v *mockInterfaceSink) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) SetMute(flags dbus.Flags, value bool) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetPort

func (v *mockInterfaceSink) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) SetPort(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method SetVolume

func (v *mockInterfaceSink) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSink) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// property SupportBalance b

func (v *mockInterfaceSink) SupportBalance() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Ports a(ssy)

func (v *mockInterfaceSink) Ports() PropPortInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropPortInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *mockInterfaceSink) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Mute b

func (v *mockInterfaceSink) Mute() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Volume d

func (v *mockInterfaceSink) Volume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Balance d

func (v *mockInterfaceSink) Balance() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ActivePort (ssy)

func (v *mockInterfaceSink) ActivePort() PropPortInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropPortInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Card u

func (v *mockInterfaceSink) Card() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Description s

func (v *mockInterfaceSink) Description() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BaseVolume d

func (v *mockInterfaceSink) BaseVolume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Fade d

func (v *mockInterfaceSink) Fade() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportFade b

func (v *mockInterfaceSink) SupportFade() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockSource struct {
	mockInterfaceSource // interface com.deepin.daemon.Audio.Source
}

type mockInterfaceSource struct {
	mock.Mock
}

// method GetMeter

func (v *mockInterfaceSource) GoGetMeter(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) GetMeter(flags dbus.Flags) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetBalance

func (v *mockInterfaceSource) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// method SetFade

func (v *mockInterfaceSource) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) SetFade(flags dbus.Flags, value float64) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetMute

func (v *mockInterfaceSource) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) SetMute(flags dbus.Flags, value bool) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetPort

func (v *mockInterfaceSource) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) SetPort(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method SetVolume

func (v *mockInterfaceSource) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSource) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// property Mute b

func (v *mockInterfaceSource) Mute() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Balance d

func (v *mockInterfaceSource) Balance() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportBalance b

func (v *mockInterfaceSource) SupportBalance() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Fade d

func (v *mockInterfaceSource) Fade() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Ports a(ssy)

func (v *mockInterfaceSource) Ports() PropPortInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropPortInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Card u

func (v *mockInterfaceSource) Card() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BaseVolume d

func (v *mockInterfaceSource) BaseVolume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Description s

func (v *mockInterfaceSource) Description() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Volume d

func (v *mockInterfaceSource) Volume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportFade b

func (v *mockInterfaceSource) SupportFade() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ActivePort (ssy)

func (v *mockInterfaceSource) ActivePort() PropPortInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropPortInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *mockInterfaceSource) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockSinkInput struct {
	mockInterfaceSinkInput // interface com.deepin.daemon.Audio.SinkInput
}

type mockInterfaceSinkInput struct {
	mock.Mock
}

// method SetBalance

func (v *mockInterfaceSinkInput) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSinkInput) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// method SetFade

func (v *mockInterfaceSinkInput) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSinkInput) SetFade(flags dbus.Flags, value float64) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetMute

func (v *mockInterfaceSinkInput) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSinkInput) SetMute(flags dbus.Flags, value bool) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetVolume

func (v *mockInterfaceSinkInput) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value, isPlay)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSinkInput) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	mockArgs := v.Called(flags, value, isPlay)

	return mockArgs.Error(0)
}

// property Volume d

func (v *mockInterfaceSinkInput) Volume() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Balance d

func (v *mockInterfaceSinkInput) Balance() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportBalance b

func (v *mockInterfaceSinkInput) SupportBalance() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Fade d

func (v *mockInterfaceSinkInput) Fade() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportFade b

func (v *mockInterfaceSinkInput) SupportFade() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *mockInterfaceSinkInput) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Icon s

func (v *mockInterfaceSinkInput) Icon() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Mute b

func (v *mockInterfaceSinkInput) Mute() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropPortInfoSlice struct {
	mock.Mock
}

func (p MockPropPortInfoSlice) Get(flags dbus.Flags) (value []PortInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]PortInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropPortInfoSlice) Set(flags dbus.Flags, value []PortInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropPortInfoSlice) ConnectChanged(cb func(hasValue bool, value []PortInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type MockPropPortInfo struct {
	mock.Mock
}

func (p MockPropPortInfo) Get(flags dbus.Flags) (value PortInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(PortInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropPortInfo) Set(flags dbus.Flags, value PortInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropPortInfo) ConnectChanged(cb func(hasValue bool, value PortInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}
