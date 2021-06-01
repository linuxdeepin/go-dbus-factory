package accounts

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
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

func (v *accounts) GoAllowGuestAccountWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, allow bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AllowGuestAccount", flags, ch, allow)
}

func (v *accounts) AllowGuestAccount(flags dbus.Flags, allow bool) error {
	return (<-v.GoAllowGuestAccount(flags, make(chan *dbus.Call, 1), allow).Done).Err
}

func (v *accounts) AllowGuestAccountWithTimeout(timeout time.Duration, flags dbus.Flags, allow bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAllowGuestAccountWithContext(ctx, flags, make(chan *dbus.Call, 1), allow).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method CreateGuestAccount

func (v *accounts) GoCreateGuestAccount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateGuestAccount", flags, ch)
}

func (v *accounts) GoCreateGuestAccountWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CreateGuestAccount", flags, ch)
}

func (*accounts) StoreCreateGuestAccount(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) CreateGuestAccount(flags dbus.Flags) (user string, err error) {
	return v.StoreCreateGuestAccount(
		<-v.GoCreateGuestAccount(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *accounts) CreateGuestAccountWithTimeout(timeout time.Duration, flags dbus.Flags) (user string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCreateGuestAccountWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCreateGuestAccount(call)
}

// method CreateUser

func (v *accounts) GoCreateUser(flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateUser", flags, ch, name, fullName, type0)
}

func (v *accounts) GoCreateUserWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CreateUser", flags, ch, name, fullName, type0)
}

func (*accounts) StoreCreateUser(call *dbus.Call) (user dbus.ObjectPath, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) CreateUser(flags dbus.Flags, name string, fullName string, type0 int32) (user dbus.ObjectPath, err error) {
	return v.StoreCreateUser(
		<-v.GoCreateUser(flags, make(chan *dbus.Call, 1), name, fullName, type0).Done)
}

func (v *accounts) CreateUserWithTimeout(timeout time.Duration, flags dbus.Flags, name string, fullName string, type0 int32) (user dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCreateUserWithContext(ctx, flags, make(chan *dbus.Call, 1), name, fullName, type0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCreateUser(call)
}

// method DeleteUser

func (v *accounts) GoDeleteUser(flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteUser", flags, ch, name, rmFiles)
}

func (v *accounts) GoDeleteUserWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteUser", flags, ch, name, rmFiles)
}

func (v *accounts) DeleteUser(flags dbus.Flags, name string, rmFiles bool) error {
	return (<-v.GoDeleteUser(flags, make(chan *dbus.Call, 1), name, rmFiles).Done).Err
}

func (v *accounts) DeleteUserWithTimeout(timeout time.Duration, flags dbus.Flags, name string, rmFiles bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteUserWithContext(ctx, flags, make(chan *dbus.Call, 1), name, rmFiles).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method FindUserById

func (v *accounts) GoFindUserById(flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserById", flags, ch, uid)
}

func (v *accounts) GoFindUserByIdWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".FindUserById", flags, ch, uid)
}

func (*accounts) StoreFindUserById(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) FindUserById(flags dbus.Flags, uid string) (user string, err error) {
	return v.StoreFindUserById(
		<-v.GoFindUserById(flags, make(chan *dbus.Call, 1), uid).Done)
}

func (v *accounts) FindUserByIdWithTimeout(timeout time.Duration, flags dbus.Flags, uid string) (user string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoFindUserByIdWithContext(ctx, flags, make(chan *dbus.Call, 1), uid).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreFindUserById(call)
}

// method FindUserByName

func (v *accounts) GoFindUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserByName", flags, ch, name)
}

func (v *accounts) GoFindUserByNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".FindUserByName", flags, ch, name)
}

func (*accounts) StoreFindUserByName(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *accounts) FindUserByName(flags dbus.Flags, name string) (user string, err error) {
	return v.StoreFindUserByName(
		<-v.GoFindUserByName(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *accounts) FindUserByNameWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (user string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoFindUserByNameWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreFindUserByName(call)
}

// method IsPasswordValid

func (v *accounts) GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPasswordValid", flags, ch, password)
}

func (v *accounts) GoIsPasswordValidWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsPasswordValid", flags, ch, password)
}

func (*accounts) StoreIsPasswordValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *accounts) IsPasswordValid(flags dbus.Flags, password string) (ok bool, errReason string, errCode int32, err error) {
	return v.StoreIsPasswordValid(
		<-v.GoIsPasswordValid(flags, make(chan *dbus.Call, 1), password).Done)
}

func (v *accounts) IsPasswordValidWithTimeout(timeout time.Duration, flags dbus.Flags, password string) (ok bool, errReason string, errCode int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsPasswordValidWithContext(ctx, flags, make(chan *dbus.Call, 1), password).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsPasswordValid(call)
}

// method IsUsernameValid

func (v *accounts) GoIsUsernameValid(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsUsernameValid", flags, ch, name)
}

func (v *accounts) GoIsUsernameValidWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".IsUsernameValid", flags, ch, name)
}

func (*accounts) StoreIsUsernameValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *accounts) IsUsernameValid(flags dbus.Flags, name string) (ok bool, errReason string, errCode int32, err error) {
	return v.StoreIsUsernameValid(
		<-v.GoIsUsernameValid(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *accounts) IsUsernameValidWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (ok bool, errReason string, errCode int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoIsUsernameValidWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreIsUsernameValid(call)
}

// method RandUserIcon

func (v *accounts) GoRandUserIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RandUserIcon", flags, ch)
}

func (v *accounts) GoRandUserIconWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RandUserIcon", flags, ch)
}

func (*accounts) StoreRandUserIcon(call *dbus.Call) (iconFile string, err error) {
	err = call.Store(&iconFile)
	return
}

func (v *accounts) RandUserIcon(flags dbus.Flags) (iconFile string, err error) {
	return v.StoreRandUserIcon(
		<-v.GoRandUserIcon(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *accounts) RandUserIconWithTimeout(timeout time.Duration, flags dbus.Flags) (iconFile string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRandUserIconWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRandUserIcon(call)
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

func (v *user) GoAddGroupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddGroup", flags, ch, group)
}

func (v *user) AddGroup(flags dbus.Flags, group string) error {
	return (<-v.GoAddGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

func (v *user) AddGroupWithTimeout(timeout time.Duration, flags dbus.Flags, group string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddGroupWithContext(ctx, flags, make(chan *dbus.Call, 1), group).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteGroup

func (v *user) GoDeleteGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteGroup", flags, ch, group)
}

func (v *user) GoDeleteGroupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteGroup", flags, ch, group)
}

func (v *user) DeleteGroup(flags dbus.Flags, group string) error {
	return (<-v.GoDeleteGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

func (v *user) DeleteGroupWithTimeout(timeout time.Duration, flags dbus.Flags, group string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteGroupWithContext(ctx, flags, make(chan *dbus.Call, 1), group).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteIconFile

func (v *user) GoDeleteIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteIconFile", flags, ch, iconFile)
}

func (v *user) GoDeleteIconFileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteIconFile", flags, ch, iconFile)
}

func (v *user) DeleteIconFile(flags dbus.Flags, iconFile string) error {
	return (<-v.GoDeleteIconFile(flags, make(chan *dbus.Call, 1), iconFile).Done).Err
}

func (v *user) DeleteIconFileWithTimeout(timeout time.Duration, flags dbus.Flags, iconFile string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteIconFileWithContext(ctx, flags, make(chan *dbus.Call, 1), iconFile).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnableNoPasswdLogin

func (v *user) GoEnableNoPasswdLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableNoPasswdLogin", flags, ch, enabled)
}

func (v *user) GoEnableNoPasswdLoginWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableNoPasswdLogin", flags, ch, enabled)
}

func (v *user) EnableNoPasswdLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableNoPasswdLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *user) EnableNoPasswdLoginWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableNoPasswdLoginWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetAutomaticLogin

func (v *user) GoSetAutomaticLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutomaticLogin", flags, ch, enabled)
}

func (v *user) GoSetAutomaticLoginWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAutomaticLogin", flags, ch, enabled)
}

func (v *user) SetAutomaticLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoSetAutomaticLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *user) SetAutomaticLoginWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAutomaticLoginWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetDesktopBackgrounds

func (v *user) GoSetDesktopBackgrounds(flags dbus.Flags, ch chan *dbus.Call, backgrounds []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDesktopBackgrounds", flags, ch, backgrounds)
}

func (v *user) GoSetDesktopBackgroundsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, backgrounds []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDesktopBackgrounds", flags, ch, backgrounds)
}

func (v *user) SetDesktopBackgrounds(flags dbus.Flags, backgrounds []string) error {
	return (<-v.GoSetDesktopBackgrounds(flags, make(chan *dbus.Call, 1), backgrounds).Done).Err
}

func (v *user) SetDesktopBackgroundsWithTimeout(timeout time.Duration, flags dbus.Flags, backgrounds []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDesktopBackgroundsWithContext(ctx, flags, make(chan *dbus.Call, 1), backgrounds).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetFullName

func (v *user) GoSetFullName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFullName", flags, ch, name)
}

func (v *user) GoSetFullNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetFullName", flags, ch, name)
}

func (v *user) SetFullName(flags dbus.Flags, name string) error {
	return (<-v.GoSetFullName(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *user) SetFullNameWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetFullNameWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetGreeterBackground

func (v *user) GoSetGreeterBackground(flags dbus.Flags, ch chan *dbus.Call, background string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetGreeterBackground", flags, ch, background)
}

func (v *user) GoSetGreeterBackgroundWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, background string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetGreeterBackground", flags, ch, background)
}

func (v *user) SetGreeterBackground(flags dbus.Flags, background string) error {
	return (<-v.GoSetGreeterBackground(flags, make(chan *dbus.Call, 1), background).Done).Err
}

func (v *user) SetGreeterBackgroundWithTimeout(timeout time.Duration, flags dbus.Flags, background string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetGreeterBackgroundWithContext(ctx, flags, make(chan *dbus.Call, 1), background).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetHistoryLayout

func (v *user) GoSetHistoryLayout(flags dbus.Flags, ch chan *dbus.Call, layouts []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHistoryLayout", flags, ch, layouts)
}

func (v *user) GoSetHistoryLayoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, layouts []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetHistoryLayout", flags, ch, layouts)
}

func (v *user) SetHistoryLayout(flags dbus.Flags, layouts []string) error {
	return (<-v.GoSetHistoryLayout(flags, make(chan *dbus.Call, 1), layouts).Done).Err
}

func (v *user) SetHistoryLayoutWithTimeout(timeout time.Duration, flags dbus.Flags, layouts []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetHistoryLayoutWithContext(ctx, flags, make(chan *dbus.Call, 1), layouts).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetHomeDir

func (v *user) GoSetHomeDir(flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHomeDir", flags, ch, home)
}

func (v *user) GoSetHomeDirWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetHomeDir", flags, ch, home)
}

func (v *user) SetHomeDir(flags dbus.Flags, home string) error {
	return (<-v.GoSetHomeDir(flags, make(chan *dbus.Call, 1), home).Done).Err
}

func (v *user) SetHomeDirWithTimeout(timeout time.Duration, flags dbus.Flags, home string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetHomeDirWithContext(ctx, flags, make(chan *dbus.Call, 1), home).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetIconFile

func (v *user) GoSetIconFile(flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIconFile", flags, ch, iconFile)
}

func (v *user) GoSetIconFileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, iconFile string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetIconFile", flags, ch, iconFile)
}

func (v *user) SetIconFile(flags dbus.Flags, iconFile string) error {
	return (<-v.GoSetIconFile(flags, make(chan *dbus.Call, 1), iconFile).Done).Err
}

func (v *user) SetIconFileWithTimeout(timeout time.Duration, flags dbus.Flags, iconFile string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetIconFileWithContext(ctx, flags, make(chan *dbus.Call, 1), iconFile).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLayout

func (v *user) GoSetLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLayout", flags, ch, layout)
}

func (v *user) GoSetLayoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLayout", flags, ch, layout)
}

func (v *user) SetLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoSetLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

func (v *user) SetLayoutWithTimeout(timeout time.Duration, flags dbus.Flags, layout string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLayoutWithContext(ctx, flags, make(chan *dbus.Call, 1), layout).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLocale

func (v *user) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *user) GoSetLocaleWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *user) SetLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoSetLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

func (v *user) SetLocaleWithTimeout(timeout time.Duration, flags dbus.Flags, locale string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLocaleWithContext(ctx, flags, make(chan *dbus.Call, 1), locale).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLocked

func (v *user) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocked", flags, ch, locked)
}

func (v *user) GoSetLockedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLocked", flags, ch, locked)
}

func (v *user) SetLocked(flags dbus.Flags, locked bool) error {
	return (<-v.GoSetLocked(flags, make(chan *dbus.Call, 1), locked).Done).Err
}

func (v *user) SetLockedWithTimeout(timeout time.Duration, flags dbus.Flags, locked bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLockedWithContext(ctx, flags, make(chan *dbus.Call, 1), locked).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetPassword

func (v *user) GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPassword", flags, ch, password)
}

func (v *user) GoSetPasswordWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetPassword", flags, ch, password)
}

func (v *user) SetPassword(flags dbus.Flags, password string) error {
	return (<-v.GoSetPassword(flags, make(chan *dbus.Call, 1), password).Done).Err
}

func (v *user) SetPasswordWithTimeout(timeout time.Duration, flags dbus.Flags, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetPasswordWithContext(ctx, flags, make(chan *dbus.Call, 1), password).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetShell

func (v *user) GoSetShell(flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShell", flags, ch, shell)
}

func (v *user) GoSetShellWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetShell", flags, ch, shell)
}

func (v *user) SetShell(flags dbus.Flags, shell string) error {
	return (<-v.GoSetShell(flags, make(chan *dbus.Call, 1), shell).Done).Err
}

func (v *user) SetShellWithTimeout(timeout time.Duration, flags dbus.Flags, shell string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetShellWithContext(ctx, flags, make(chan *dbus.Call, 1), shell).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetUse24HourFormat

func (v *user) GoSetUse24HourFormat(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUse24HourFormat", flags, ch, value)
}

func (v *user) GoSetUse24HourFormatWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetUse24HourFormat", flags, ch, value)
}

func (v *user) SetUse24HourFormat(flags dbus.Flags, value bool) error {
	return (<-v.GoSetUse24HourFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

func (v *user) SetUse24HourFormatWithTimeout(timeout time.Duration, flags dbus.Flags, value bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetUse24HourFormatWithContext(ctx, flags, make(chan *dbus.Call, 1), value).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
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

func (v *imageBlur) GoDeleteWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, file string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Delete", flags, ch, file)
}

func (v *imageBlur) Delete(flags dbus.Flags, file string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), file).Done).Err
}

func (v *imageBlur) DeleteWithTimeout(timeout time.Duration, flags dbus.Flags, file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteWithContext(ctx, flags, make(chan *dbus.Call, 1), file).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Get

func (v *imageBlur) GoGet(flags dbus.Flags, ch chan *dbus.Call, source string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch, source)
}

func (v *imageBlur) GoGetWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, source string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Get", flags, ch, source)
}

func (*imageBlur) StoreGet(call *dbus.Call) (blurred string, err error) {
	err = call.Store(&blurred)
	return
}

func (v *imageBlur) Get(flags dbus.Flags, source string) (blurred string, err error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1), source).Done)
}

func (v *imageBlur) GetWithTimeout(timeout time.Duration, flags dbus.Flags, source string) (blurred string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetWithContext(ctx, flags, make(chan *dbus.Call, 1), source).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGet(call)
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
