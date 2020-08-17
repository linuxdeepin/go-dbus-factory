package accounts

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Accounts struct {
	accounts // interface com.deepin.daemon.Accounts
	proxy.Object
}

func NewAccounts(conn *dbus.Conn) *Accounts {
	obj := new(Accounts)
	obj.Object.Init_(conn, "com.deepin.daemon.Accounts", "/com/deepin/daemon/Accounts")
	return obj
}

type accounts struct{}

func (v *accounts) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*accounts) GetInterfaceName_() string {
	return "com.deepin.daemon.Accounts"
}

// method AllowGuestAccount

func (v *accounts) GoAllowGuestAccount(flags dbus.Flags, ch chan *dbus.Call, allow bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowGuestAccount", flags, ch, allow)
}

func (v *accounts) AllowGuestAccount(flags dbus.Flags, allow bool) error {
	return (<-v.GoAllowGuestAccount(flags, make(chan *dbus.Call, 1), allow).Done).Err
}

// method CreateGuestAccount

func (v *accounts) GoCreateGuestAccount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateGuestAccount", flags, ch)
}

func (*accounts) StoreCreateGuestAccount(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) CreateGuestAccount(flags dbus.Flags) (user string, err error) {
	return v.StoreCreateGuestAccount(
		<-v.GoCreateGuestAccount(flags, make(chan *dbus.Call, 1)).Done)
}

// method CreateUser

func (v *accounts) GoCreateUser(flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateUser", flags, ch, name, fullName, type0)
}

func (*accounts) StoreCreateUser(call *dbus.Call) (user dbus.ObjectPath, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) CreateUser(flags dbus.Flags, name string, fullName string, type0 int32) (user dbus.ObjectPath, err error) {
	return v.StoreCreateUser(
		<-v.GoCreateUser(flags, make(chan *dbus.Call, 1), name, fullName, type0).Done)
}

// method DeleteUser

func (v *accounts) GoDeleteUser(flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteUser", flags, ch, name, rmFiles)
}

func (v *accounts) DeleteUser(flags dbus.Flags, name string, rmFiles bool) error {
	return (<-v.GoDeleteUser(flags, make(chan *dbus.Call, 1), name, rmFiles).Done).Err
}

// method FindUserById

func (v *accounts) GoFindUserById(flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserById", flags, ch, uid)
}

func (*accounts) StoreFindUserById(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) FindUserById(flags dbus.Flags, uid string) (user string, err error) {
	return v.StoreFindUserById(
		<-v.GoFindUserById(flags, make(chan *dbus.Call, 1), uid).Done)
}

// method FindUserByName

func (v *accounts) GoFindUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserByName", flags, ch, name)
}

func (*accounts) StoreFindUserByName(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) FindUserByName(flags dbus.Flags, name string) (user string, err error) {
	return v.StoreFindUserByName(
		<-v.GoFindUserByName(flags, make(chan *dbus.Call, 1), name).Done)
}

// method IsPasswordValid

func (v *accounts) GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPasswordValid", flags, ch, password)
}

func (*accounts) StoreIsPasswordValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *accounts) IsPasswordValid(flags dbus.Flags, password string) (ok bool, errReason string, errCode int32, err error) {
	return v.StoreIsPasswordValid(
		<-v.GoIsPasswordValid(flags, make(chan *dbus.Call, 1), password).Done)
}

// method IsUsernameValid

func (v *accounts) GoIsUsernameValid(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsUsernameValid", flags, ch, name)
}

func (*accounts) StoreIsUsernameValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *accounts) IsUsernameValid(flags dbus.Flags, name string) (ok bool, errReason string, errCode int32, err error) {
	return v.StoreIsUsernameValid(
		<-v.GoIsUsernameValid(flags, make(chan *dbus.Call, 1), name).Done)
}

// method RandUserIcon

func (v *accounts) GoRandUserIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RandUserIcon", flags, ch)
}

func (*accounts) StoreRandUserIcon(call *dbus.Call) (iconFile string, err error) {
	err = call.Store(&iconFile)
	return
}

func (v *accounts) RandUserIcon(flags dbus.Flags) (iconFile string, err error) {
	return v.StoreRandUserIcon(
		<-v.GoRandUserIcon(flags, make(chan *dbus.Call, 1)).Done)
}

// signal UserAdded

func (v *accounts) ConnectUserAdded(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var objPath string
		err := dbus.Store(sig.Body, &objPath)
		if err == nil {
			cb(objPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UserDeleted

func (v *accounts) ConnectUserDeleted(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var objPath string
		err := dbus.Store(sig.Body, &objPath)
		if err == nil {
			cb(objPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property UserList as

func (v *accounts) UserList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UserList",
	}
}

// property GuestIcon s

func (v *accounts) GuestIcon() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GuestIcon",
	}
}

// property AllowGuest b

func (v *accounts) AllowGuest() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AllowGuest",
	}
}

type User struct {
	user // interface com.deepin.daemon.Accounts.User
	proxy.Object
}

func NewUser(conn *dbus.Conn, path dbus.ObjectPath) (*User, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(User)
	obj.Object.Init_(conn, "com.deepin.daemon.Accounts", path)
	return obj, nil
}

type user struct{}

func (v *user) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*user) GetInterfaceName_() string {
	return "com.deepin.daemon.Accounts.User"
}

// method AddGroup

func (v *user) GoAddGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddGroup", flags, ch, group)
}

func (v *user) AddGroup(flags dbus.Flags, group string) error {
	return (<-v.GoAddGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

// method DeleteGroup

func (v *user) GoDeleteGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteGroup", flags, ch, group)
}

func (v *user) DeleteGroup(flags dbus.Flags, group string) error {
	return (<-v.GoDeleteGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

// method DeleteIconFile

func (v *user) GoDeleteIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteIconFile", flags, ch, iconFile)
}

func (v *user) DeleteIconFile(flags dbus.Flags, iconFile string) error {
	return (<-v.GoDeleteIconFile(flags, make(chan *dbus.Call, 1), iconFile).Done).Err
}

// method EnableNoPasswdLogin

func (v *user) GoEnableNoPasswdLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableNoPasswdLogin", flags, ch, enabled)
}

func (v *user) EnableNoPasswdLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableNoPasswdLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method SetAutomaticLogin

func (v *user) GoSetAutomaticLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutomaticLogin", flags, ch, enabled)
}

func (v *user) SetAutomaticLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoSetAutomaticLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method SetDesktopBackgrounds

func (v *user) GoSetDesktopBackgrounds(flags dbus.Flags, ch chan *dbus.Call, backgrounds []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDesktopBackgrounds", flags, ch, backgrounds)
}

func (v *user) SetDesktopBackgrounds(flags dbus.Flags, backgrounds []string) error {
	return (<-v.GoSetDesktopBackgrounds(flags, make(chan *dbus.Call, 1), backgrounds).Done).Err
}

// method SetFullName

func (v *user) GoSetFullName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFullName", flags, ch, name)
}

func (v *user) SetFullName(flags dbus.Flags, name string) error {
	return (<-v.GoSetFullName(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetGreeterBackground

func (v *user) GoSetGreeterBackground(flags dbus.Flags, ch chan *dbus.Call, background string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetGreeterBackground", flags, ch, background)
}

func (v *user) SetGreeterBackground(flags dbus.Flags, background string) error {
	return (<-v.GoSetGreeterBackground(flags, make(chan *dbus.Call, 1), background).Done).Err
}

// method SetHistoryLayout

func (v *user) GoSetHistoryLayout(flags dbus.Flags, ch chan *dbus.Call, layouts []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHistoryLayout", flags, ch, layouts)
}

func (v *user) SetHistoryLayout(flags dbus.Flags, layouts []string) error {
	return (<-v.GoSetHistoryLayout(flags, make(chan *dbus.Call, 1), layouts).Done).Err
}

// method SetHomeDir

func (v *user) GoSetHomeDir(flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHomeDir", flags, ch, home)
}

func (v *user) SetHomeDir(flags dbus.Flags, home string) error {
	return (<-v.GoSetHomeDir(flags, make(chan *dbus.Call, 1), home).Done).Err
}

// method SetIconFile

func (v *user) GoSetIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIconFile", flags, ch, iconFile)
}

func (v *user) SetIconFile(flags dbus.Flags, iconFile string) error {
	return (<-v.GoSetIconFile(flags, make(chan *dbus.Call, 1), iconFile).Done).Err
}

// method SetLayout

func (v *user) GoSetLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLayout", flags, ch, layout)
}

func (v *user) SetLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoSetLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method SetLocale

func (v *user) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *user) SetLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoSetLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// method SetLocked

func (v *user) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocked", flags, ch, locked)
}

func (v *user) SetLocked(flags dbus.Flags, locked bool) error {
	return (<-v.GoSetLocked(flags, make(chan *dbus.Call, 1), locked).Done).Err
}

// method SetPassword

func (v *user) GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPassword", flags, ch, password)
}

func (v *user) SetPassword(flags dbus.Flags, password string) error {
	return (<-v.GoSetPassword(flags, make(chan *dbus.Call, 1), password).Done).Err
}

// method SetShell

func (v *user) GoSetShell(flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShell", flags, ch, shell)
}

func (v *user) SetShell(flags dbus.Flags, shell string) error {
	return (<-v.GoSetShell(flags, make(chan *dbus.Call, 1), shell).Done).Err
}

// method SetUse24HourFormat

func (v *user) GoSetUse24HourFormat(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUse24HourFormat", flags, ch, value)
}

func (v *user) SetUse24HourFormat(flags dbus.Flags, value bool) error {
	return (<-v.GoSetUse24HourFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetWeekdayFormat

func (v *user) GoSetWeekdayFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWeekdayFormat", flags, ch, value)
}

func (v *user) SetWeekdayFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetWeekdayFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetShortDateFormat

func (v *user) GoSetShortDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortDateFormat", flags, ch, value)
}

func (v *user) SetShortDateFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetShortDateFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetLongDateFormat

func (v *user) GoSetLongDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongDateFormat", flags, ch, value)
}

func (v *user) SetLongDateFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetLongDateFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetShortTimeFormat

func (v *user) GoSetShortTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortTimeFormat", flags, ch, value)
}

func (v *user) SetShortTimeFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetShortTimeFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetLongTimeFormat

func (v *user) GoSetLongTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongTimeFormat", flags, ch, value)
}

func (v *user) SetLongTimeFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetLongTimeFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetWeekBegins

func (v *user) GoSetWeekBegins(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWeekBegins", flags, ch, value)
}

func (v *user) SetWeekBegins(flags dbus.Flags, value int32) error {
	return (<-v.GoSetWeekBegins(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// property HistoryLayout as

func (v *user) HistoryLayout() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "HistoryLayout",
	}
}

// property Gid s

func (v *user) Gid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Gid",
	}
}

// property Groups as

func (v *user) Groups() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Groups",
	}
}

// property XSession s

func (v *user) XSession() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "XSession",
	}
}

// property PasswordStatus s

func (v *user) PasswordStatus() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PasswordStatus",
	}
}

// property LoginTime t

func (v *user) LoginTime() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LoginTime",
	}
}

// property GreeterBackground s

func (v *user) GreeterBackground() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "GreeterBackground",
	}
}

// property CreatedTime t

func (v *user) CreatedTime() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CreatedTime",
	}
}

// property UserName s

func (v *user) UserName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UserName",
	}
}

// property Shell s

func (v *user) Shell() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Shell",
	}
}

// property Layout s

func (v *user) Layout() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Layout",
	}
}

// property IconFile s

func (v *user) IconFile() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IconFile",
	}
}

// property Use24HourFormat b

func (v *user) Use24HourFormat() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Use24HourFormat",
	}
}

// property AccountType i

func (v *user) AccountType() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "AccountType",
	}
}

// property HomeDir s

func (v *user) HomeDir() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HomeDir",
	}
}

// property Locale s

func (v *user) Locale() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Locale",
	}
}

// property DesktopBackgrounds as

func (v *user) DesktopBackgrounds() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DesktopBackgrounds",
	}
}

// property Locked b

func (v *user) Locked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property NoPasswdLogin b

func (v *user) NoPasswdLogin() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NoPasswdLogin",
	}
}

// property IconList as

func (v *user) IconList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "IconList",
	}
}

// property UUID s

func (v *user) UUID() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UUID",
	}
}

// property FullName s

func (v *user) FullName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FullName",
	}
}

// property Uid s

func (v *user) Uid() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Uid",
	}
}

// property AutomaticLogin b

func (v *user) AutomaticLogin() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutomaticLogin",
	}
}

// property SystemAccount b

func (v *user) SystemAccount() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SystemAccount",
	}
}

// property WeekdayFormat i

func (v *user) WeekdayFormat() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "WeekdayFormat",
	}
}

// property ShortDateFormat i

func (v *user) ShortDateFormat() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "ShortDateFormat",
	}
}

// property LongDateFormat i

func (v *user) LongDateFormat() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "LongDateFormat",
	}
}

// property ShortTimeFormat i

func (v *user) ShortTimeFormat() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "ShortTimeFormat",
	}
}

// property LongTimeFormat i

func (v *user) LongTimeFormat() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "LongTimeFormat",
	}
}

// property WeekBegins i

func (v *user) WeekBegins() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "WeekBegins",
	}
}

type ImageBlur struct {
	imageBlur // interface com.deepin.daemon.ImageBlur
	proxy.Object
}

func NewImageBlur(conn *dbus.Conn) *ImageBlur {
	obj := new(ImageBlur)
	obj.Object.Init_(conn, "com.deepin.daemon.Accounts", "/com/deepin/daemon/ImageBlur")
	return obj
}

type imageBlur struct{}

func (v *imageBlur) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*imageBlur) GetInterfaceName_() string {
	return "com.deepin.daemon.ImageBlur"
}

// method Delete

func (v *imageBlur) GoDelete(flags dbus.Flags, ch chan *dbus.Call, file string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, file)
}

func (v *imageBlur) Delete(flags dbus.Flags, file string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), file).Done).Err
}

// method Get

func (v *imageBlur) GoGet(flags dbus.Flags, ch chan *dbus.Call, source string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch, source)
}

func (*imageBlur) StoreGet(call *dbus.Call) (blurred string, err error) {
	err = call.Store(&blurred)
	return
}

func (v *imageBlur) Get(flags dbus.Flags, source string) (blurred string, err error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1), source).Done)
}

// signal BlurDone

func (v *imageBlur) ConnectBlurDone(cb func(imgFile string, imgBlurFile string, ok bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BlurDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BlurDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var imgFile string
		var imgBlurFile string
		var ok bool
		err := dbus.Store(sig.Body, &imgFile, &imgBlurFile, &ok)
		if err == nil {
			cb(imgFile, imgBlurFile, ok)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
