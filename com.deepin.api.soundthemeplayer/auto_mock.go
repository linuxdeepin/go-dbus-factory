// Code generated by "./generator ./com.deepin.api.soundthemeplayer"; DO NOT EDIT.

package soundthemeplayer

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
)

type MockSoundThemePlayer struct {
	mockInterfaceSoundThemePlayer // interface com.deepin.api.SoundThemePlayer
}

type mockInterfaceSoundThemePlayer struct {
	mock.Mock
}

// method EnableSoundDesktopLogin

func (v *mockInterfaceSoundThemePlayer) GoEnableSoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSoundThemePlayer) EnableSoundDesktopLogin(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method Play

func (v *mockInterfaceSoundThemePlayer) GoPlay(flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, theme, event, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSoundThemePlayer) Play(flags dbus.Flags, theme string, event string, device string) error {
	mockArgs := v.Called(flags, theme, event, device)

	return mockArgs.Error(0)
}

// method PlaySoundDesktopLogin

func (v *mockInterfaceSoundThemePlayer) GoPlaySoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSoundThemePlayer) PlaySoundDesktopLogin(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SaveAudioState

func (v *mockInterfaceSoundThemePlayer) GoSaveAudioState(flags dbus.Flags, ch chan *dbus.Call, activePlayback map[string]dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, activePlayback)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSoundThemePlayer) SaveAudioState(flags dbus.Flags, activePlayback map[string]dbus.Variant) error {
	mockArgs := v.Called(flags, activePlayback)

	return mockArgs.Error(0)
}

// method SetSoundTheme

func (v *mockInterfaceSoundThemePlayer) GoSetSoundTheme(flags dbus.Flags, ch chan *dbus.Call, theme string) *dbus.Call {
	mockArgs := v.Called(flags, ch, theme)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceSoundThemePlayer) SetSoundTheme(flags dbus.Flags, theme string) error {
	mockArgs := v.Called(flags, theme)

	return mockArgs.Error(0)
}
