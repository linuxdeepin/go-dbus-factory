// Code generated by "./generator ./com.deepin.daemon.accounts"; DO NOT EDIT.

package accounts

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockAccounts struct {
	mockInterfaceAccounts // interface com.deepin.daemon.Accounts
}

type mockInterfaceAccounts struct {
	mock.Mock
}

// method AllowGuestAccount

func (v *mockInterfaceAccounts) GoAllowGuestAccount(flags dbus.Flags, ch chan *dbus.Call, allow bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, allow)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) AllowGuestAccount(flags dbus.Flags, allow bool) error {
	mockArgs := v.Called(flags, allow)

	return mockArgs.Error(0)
}

// method CreateGuestAccount

func (v *mockInterfaceAccounts) GoCreateGuestAccount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) CreateGuestAccount(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method CreateUser

func (v *mockInterfaceAccounts) GoCreateUser(flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, name, fullName, type0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) CreateUser(flags dbus.Flags, name string, fullName string, type0 int32) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, name, fullName, type0)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DeleteUser

func (v *mockInterfaceAccounts) GoDeleteUser(flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, name, rmFiles)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) DeleteUser(flags dbus.Flags, name string, rmFiles bool) error {
	mockArgs := v.Called(flags, name, rmFiles)

	return mockArgs.Error(0)
}

// method FindUserById

func (v *mockInterfaceAccounts) GoFindUserById(flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) FindUserById(flags dbus.Flags, uid string) (string, error) {
	mockArgs := v.Called(flags, uid)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method FindUserByName

func (v *mockInterfaceAccounts) GoFindUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) FindUserByName(flags dbus.Flags, name string) (string, error) {
	mockArgs := v.Called(flags, name)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsPasswordValid

func (v *mockInterfaceAccounts) GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	mockArgs := v.Called(flags, ch, password)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) IsPasswordValid(flags dbus.Flags, password string) (bool, string, int32, error) {
	mockArgs := v.Called(flags, password)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	ret2, ok := mockArgs.Get(2).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 2, mockArgs.Get(2)))
	}

	return ret0, ret1, ret2, mockArgs.Error(3)
}

// method IsUsernameValid

func (v *mockInterfaceAccounts) GoIsUsernameValid(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) IsUsernameValid(flags dbus.Flags, name string) (bool, string, int32, error) {
	mockArgs := v.Called(flags, name)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	ret2, ok := mockArgs.Get(2).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 2, mockArgs.Get(2)))
	}

	return ret0, ret1, ret2, mockArgs.Error(3)
}

// method RandUserIcon

func (v *mockInterfaceAccounts) GoRandUserIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceAccounts) RandUserIcon(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UserAdded

func (v *mockInterfaceAccounts) ConnectUserAdded(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UserDeleted

func (v *mockInterfaceAccounts) ConnectUserDeleted(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property UserList as

func (v *mockInterfaceAccounts) UserList() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property GuestIcon s

func (v *mockInterfaceAccounts) GuestIcon() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property AllowGuest b

func (v *mockInterfaceAccounts) AllowGuest() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockUser struct {
	mockInterfaceUser // interface com.deepin.daemon.Accounts.User
}

type mockInterfaceUser struct {
	mock.Mock
}

// method AddGroup

func (v *mockInterfaceUser) GoAddGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	mockArgs := v.Called(flags, ch, group)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) AddGroup(flags dbus.Flags, group string) error {
	mockArgs := v.Called(flags, group)

	return mockArgs.Error(0)
}

// method DeleteGroup

func (v *mockInterfaceUser) GoDeleteGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	mockArgs := v.Called(flags, ch, group)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) DeleteGroup(flags dbus.Flags, group string) error {
	mockArgs := v.Called(flags, group)

	return mockArgs.Error(0)
}

// method DeleteIconFile

func (v *mockInterfaceUser) GoDeleteIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, iconFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) DeleteIconFile(flags dbus.Flags, iconFile string) error {
	mockArgs := v.Called(flags, iconFile)

	return mockArgs.Error(0)
}

// method EnableNoPasswdLogin

func (v *mockInterfaceUser) GoEnableNoPasswdLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) EnableNoPasswdLogin(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method SetAutomaticLogin

func (v *mockInterfaceUser) GoSetAutomaticLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetAutomaticLogin(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method SetDesktopBackgrounds

func (v *mockInterfaceUser) GoSetDesktopBackgrounds(flags dbus.Flags, ch chan *dbus.Call, backgrounds []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, backgrounds)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetDesktopBackgrounds(flags dbus.Flags, backgrounds []string) error {
	mockArgs := v.Called(flags, backgrounds)

	return mockArgs.Error(0)
}

// method SetFullName

func (v *mockInterfaceUser) GoSetFullName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetFullName(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method SetGreeterBackground

func (v *mockInterfaceUser) GoSetGreeterBackground(flags dbus.Flags, ch chan *dbus.Call, background string) *dbus.Call {
	mockArgs := v.Called(flags, ch, background)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetGreeterBackground(flags dbus.Flags, background string) error {
	mockArgs := v.Called(flags, background)

	return mockArgs.Error(0)
}

// method SetCurrentWorkspace

func (v *mockInterfaceUser) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, currentWorkspace int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, currentWorkspace)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetCurrentWorkspace(flags dbus.Flags, currentWorkspace int32) error {
	mockArgs := v.Called(flags, currentWorkspace)

	return mockArgs.Error(0)
}

// method SetHistoryLayout

func (v *mockInterfaceUser) GoSetHistoryLayout(flags dbus.Flags, ch chan *dbus.Call, layouts []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, layouts)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetHistoryLayout(flags dbus.Flags, layouts []string) error {
	mockArgs := v.Called(flags, layouts)

	return mockArgs.Error(0)
}

// method SetHomeDir

func (v *mockInterfaceUser) GoSetHomeDir(flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call {
	mockArgs := v.Called(flags, ch, home)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetHomeDir(flags dbus.Flags, home string) error {
	mockArgs := v.Called(flags, home)

	return mockArgs.Error(0)
}

// method SetIconFile

func (v *mockInterfaceUser) GoSetIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, iconFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetIconFile(flags dbus.Flags, iconFile string) error {
	mockArgs := v.Called(flags, iconFile)

	return mockArgs.Error(0)
}

// method SetLayout

func (v *mockInterfaceUser) GoSetLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	mockArgs := v.Called(flags, ch, layout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetLayout(flags dbus.Flags, layout string) error {
	mockArgs := v.Called(flags, layout)

	return mockArgs.Error(0)
}

// method SetLocale

func (v *mockInterfaceUser) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	mockArgs := v.Called(flags, ch, locale)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetLocale(flags dbus.Flags, locale string) error {
	mockArgs := v.Called(flags, locale)

	return mockArgs.Error(0)
}

// method SetLocked

func (v *mockInterfaceUser) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, locked)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetLocked(flags dbus.Flags, locked bool) error {
	mockArgs := v.Called(flags, locked)

	return mockArgs.Error(0)
}

// method SetPassword

func (v *mockInterfaceUser) GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	mockArgs := v.Called(flags, ch, password)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetPassword(flags dbus.Flags, password string) error {
	mockArgs := v.Called(flags, password)

	return mockArgs.Error(0)
}

// method SetShell

func (v *mockInterfaceUser) GoSetShell(flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call {
	mockArgs := v.Called(flags, ch, shell)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetShell(flags dbus.Flags, shell string) error {
	mockArgs := v.Called(flags, shell)

	return mockArgs.Error(0)
}

// method SetUse24HourFormat

func (v *mockInterfaceUser) GoSetUse24HourFormat(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetUse24HourFormat(flags dbus.Flags, value bool) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetWeekdayFormat

func (v *mockInterfaceUser) GoSetWeekdayFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetWeekdayFormat(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetShortDateFormat

func (v *mockInterfaceUser) GoSetShortDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetShortDateFormat(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetLongDateFormat

func (v *mockInterfaceUser) GoSetLongDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetLongDateFormat(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetShortTimeFormat

func (v *mockInterfaceUser) GoSetShortTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetShortTimeFormat(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetLongTimeFormat

func (v *mockInterfaceUser) GoSetLongTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetLongTimeFormat(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetWeekBegins

func (v *mockInterfaceUser) GoSetWeekBegins(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUser) SetWeekBegins(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// property HistoryLayout as

func (v *mockInterfaceUser) HistoryLayout() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Gid s

func (v *mockInterfaceUser) Gid() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Groups as

func (v *mockInterfaceUser) Groups() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property XSession s

func (v *mockInterfaceUser) XSession() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PasswordStatus s

func (v *mockInterfaceUser) PasswordStatus() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LoginTime t

func (v *mockInterfaceUser) LoginTime() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property GreeterBackground s

func (v *mockInterfaceUser) GreeterBackground() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CreatedTime t

func (v *mockInterfaceUser) CreatedTime() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property UserName s

func (v *mockInterfaceUser) UserName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Shell s

func (v *mockInterfaceUser) Shell() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Layout s

func (v *mockInterfaceUser) Layout() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IconFile s

func (v *mockInterfaceUser) IconFile() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Use24HourFormat b

func (v *mockInterfaceUser) Use24HourFormat() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property AccountType i

func (v *mockInterfaceUser) AccountType() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HomeDir s

func (v *mockInterfaceUser) HomeDir() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Locale s

func (v *mockInterfaceUser) Locale() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DesktopBackgrounds as

func (v *mockInterfaceUser) DesktopBackgrounds() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Locked b

func (v *mockInterfaceUser) Locked() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property NoPasswdLogin b

func (v *mockInterfaceUser) NoPasswdLogin() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IconList as

func (v *mockInterfaceUser) IconList() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property UUID s

func (v *mockInterfaceUser) UUID() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property FullName s

func (v *mockInterfaceUser) FullName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Uid s

func (v *mockInterfaceUser) Uid() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property AutomaticLogin b

func (v *mockInterfaceUser) AutomaticLogin() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SystemAccount b

func (v *mockInterfaceUser) SystemAccount() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property WeekdayFormat i

func (v *mockInterfaceUser) WeekdayFormat() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ShortDateFormat i

func (v *mockInterfaceUser) ShortDateFormat() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LongDateFormat i

func (v *mockInterfaceUser) LongDateFormat() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ShortTimeFormat i

func (v *mockInterfaceUser) ShortTimeFormat() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LongTimeFormat i

func (v *mockInterfaceUser) LongTimeFormat() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property WeekBegins i

func (v *mockInterfaceUser) WeekBegins() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockImageBlur struct {
	mockInterfaceImageBlur // interface com.deepin.daemon.ImageBlur
}

type mockInterfaceImageBlur struct {
	mock.Mock
}

// method Delete

func (v *mockInterfaceImageBlur) GoDelete(flags dbus.Flags, ch chan *dbus.Call, file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceImageBlur) Delete(flags dbus.Flags, file string) error {
	mockArgs := v.Called(flags, file)

	return mockArgs.Error(0)
}

// method Get

func (v *mockInterfaceImageBlur) GoGet(flags dbus.Flags, ch chan *dbus.Call, source string) *dbus.Call {
	mockArgs := v.Called(flags, ch, source)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceImageBlur) Get(flags dbus.Flags, source string) (string, error) {
	mockArgs := v.Called(flags, source)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal BlurDone

func (v *mockInterfaceImageBlur) ConnectBlurDone(cb func(imgFile string, imgBlurFile string, ok bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
