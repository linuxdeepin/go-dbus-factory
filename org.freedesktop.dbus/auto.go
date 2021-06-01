package dbus

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

type DBus struct {
	dbusIfc // interface org.freedesktop.DBus
	proxy.Object
}

func NewDBus(conn *dbus.Conn) *DBus {
	obj := new(DBus)
	obj.Object.Init_(conn, "org.freedesktop.DBus", "/org/freedesktop/DBus")
	return obj
}

type dbusIfc struct{}

func (v *dbusIfc) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dbusIfc) GetInterfaceName_() string {
	return "org.freedesktop.DBus"
}

// method Hello

func (v *dbusIfc) GoHello(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hello", flags, ch)
}

func (v *dbusIfc) GoHelloWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Hello", flags, ch)
}

func (*dbusIfc) StoreHello(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) Hello(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreHello(
		<-v.GoHello(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *dbusIfc) HelloWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHelloWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreHello(call)
}

// method RequestName

func (v *dbusIfc) GoRequestName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestName", flags, ch, arg0, arg1)
}

func (v *dbusIfc) GoRequestNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RequestName", flags, ch, arg0, arg1)
}

func (*dbusIfc) StoreRequestName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *dbusIfc) RequestName(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	return v.StoreRequestName(
		<-v.GoRequestName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

func (v *dbusIfc) RequestNameWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRequestNameWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRequestName(call)
}

// method ReleaseName

func (v *dbusIfc) GoReleaseName(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseName", flags, ch, arg0)
}

func (v *dbusIfc) GoReleaseNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReleaseName", flags, ch, arg0)
}

func (*dbusIfc) StoreReleaseName(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) ReleaseName(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreReleaseName(
		<-v.GoReleaseName(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) ReleaseNameWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReleaseNameWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReleaseName(call)
}

// method StartServiceByName

func (v *dbusIfc) GoStartServiceByName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartServiceByName", flags, ch, arg0, arg1)
}

func (v *dbusIfc) GoStartServiceByNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartServiceByName", flags, ch, arg0, arg1)
}

func (*dbusIfc) StoreStartServiceByName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *dbusIfc) StartServiceByName(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	return v.StoreStartServiceByName(
		<-v.GoStartServiceByName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

func (v *dbusIfc) StartServiceByNameWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartServiceByNameWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStartServiceByName(call)
}

// method UpdateActivationEnvironment

func (v *dbusIfc) GoUpdateActivationEnvironment(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateActivationEnvironment", flags, ch, arg0)
}

func (v *dbusIfc) GoUpdateActivationEnvironmentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UpdateActivationEnvironment", flags, ch, arg0)
}

func (v *dbusIfc) UpdateActivationEnvironment(flags dbus.Flags, arg0 map[string]string) error {
	return (<-v.GoUpdateActivationEnvironment(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *dbusIfc) UpdateActivationEnvironmentWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUpdateActivationEnvironmentWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method NameHasOwner

func (v *dbusIfc) GoNameHasOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NameHasOwner", flags, ch, arg0)
}

func (v *dbusIfc) GoNameHasOwnerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".NameHasOwner", flags, ch, arg0)
}

func (*dbusIfc) StoreNameHasOwner(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) NameHasOwner(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreNameHasOwner(
		<-v.GoNameHasOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) NameHasOwnerWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoNameHasOwnerWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreNameHasOwner(call)
}

// method ListNames

func (v *dbusIfc) GoListNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListNames", flags, ch)
}

func (v *dbusIfc) GoListNamesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListNames", flags, ch)
}

func (*dbusIfc) StoreListNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) ListNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListNames(
		<-v.GoListNames(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *dbusIfc) ListNamesWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListNamesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListNames(call)
}

// method ListActivatableNames

func (v *dbusIfc) GoListActivatableNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListActivatableNames", flags, ch)
}

func (v *dbusIfc) GoListActivatableNamesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListActivatableNames", flags, ch)
}

func (*dbusIfc) StoreListActivatableNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) ListActivatableNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListActivatableNames(
		<-v.GoListActivatableNames(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *dbusIfc) ListActivatableNamesWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListActivatableNamesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListActivatableNames(call)
}

// method AddMatch

func (v *dbusIfc) GoAddMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddMatch", flags, ch, arg0)
}

func (v *dbusIfc) GoAddMatchWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddMatch", flags, ch, arg0)
}

func (v *dbusIfc) AddMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoAddMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *dbusIfc) AddMatchWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddMatchWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RemoveMatch

func (v *dbusIfc) GoRemoveMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveMatch", flags, ch, arg0)
}

func (v *dbusIfc) GoRemoveMatchWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveMatch", flags, ch, arg0)
}

func (v *dbusIfc) RemoveMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRemoveMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

func (v *dbusIfc) RemoveMatchWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveMatchWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetNameOwner

func (v *dbusIfc) GoGetNameOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNameOwner", flags, ch, arg0)
}

func (v *dbusIfc) GoGetNameOwnerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetNameOwner", flags, ch, arg0)
}

func (*dbusIfc) StoreGetNameOwner(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetNameOwner(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetNameOwner(
		<-v.GoGetNameOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetNameOwnerWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetNameOwnerWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetNameOwner(call)
}

// method ListQueuedOwners

func (v *dbusIfc) GoListQueuedOwners(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListQueuedOwners", flags, ch, arg0)
}

func (v *dbusIfc) GoListQueuedOwnersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListQueuedOwners", flags, ch, arg0)
}

func (*dbusIfc) StoreListQueuedOwners(call *dbus.Call) (arg1 []string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) ListQueuedOwners(flags dbus.Flags, arg0 string) (arg1 []string, err error) {
	return v.StoreListQueuedOwners(
		<-v.GoListQueuedOwners(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) ListQueuedOwnersWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListQueuedOwnersWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListQueuedOwners(call)
}

// method GetConnectionUnixUser

func (v *dbusIfc) GoGetConnectionUnixUser(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixUser", flags, ch, arg0)
}

func (v *dbusIfc) GoGetConnectionUnixUserWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetConnectionUnixUser", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionUnixUser(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionUnixUser(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreGetConnectionUnixUser(
		<-v.GoGetConnectionUnixUser(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetConnectionUnixUserWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetConnectionUnixUserWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetConnectionUnixUser(call)
}

// method GetConnectionUnixProcessID

func (v *dbusIfc) GoGetConnectionUnixProcessID(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixProcessID", flags, ch, arg0)
}

func (v *dbusIfc) GoGetConnectionUnixProcessIDWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetConnectionUnixProcessID", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionUnixProcessID(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionUnixProcessID(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreGetConnectionUnixProcessID(
		<-v.GoGetConnectionUnixProcessID(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetConnectionUnixProcessIDWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetConnectionUnixProcessIDWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetConnectionUnixProcessID(call)
}

// method GetAdtAuditSessionData

func (v *dbusIfc) GoGetAdtAuditSessionData(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAdtAuditSessionData", flags, ch, arg0)
}

func (v *dbusIfc) GoGetAdtAuditSessionDataWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetAdtAuditSessionData", flags, ch, arg0)
}

func (*dbusIfc) StoreGetAdtAuditSessionData(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetAdtAuditSessionData(flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	return v.StoreGetAdtAuditSessionData(
		<-v.GoGetAdtAuditSessionData(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetAdtAuditSessionDataWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetAdtAuditSessionDataWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetAdtAuditSessionData(call)
}

// method GetConnectionSELinuxSecurityContext

func (v *dbusIfc) GoGetConnectionSELinuxSecurityContext(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionSELinuxSecurityContext", flags, ch, arg0)
}

func (v *dbusIfc) GoGetConnectionSELinuxSecurityContextWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetConnectionSELinuxSecurityContext", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionSELinuxSecurityContext(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionSELinuxSecurityContext(flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	return v.StoreGetConnectionSELinuxSecurityContext(
		<-v.GoGetConnectionSELinuxSecurityContext(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetConnectionSELinuxSecurityContextWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetConnectionSELinuxSecurityContextWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetConnectionSELinuxSecurityContext(call)
}

// method ReloadConfig

func (v *dbusIfc) GoReloadConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadConfig", flags, ch)
}

func (v *dbusIfc) GoReloadConfigWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadConfig", flags, ch)
}

func (v *dbusIfc) ReloadConfig(flags dbus.Flags) error {
	return (<-v.GoReloadConfig(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *dbusIfc) ReloadConfigWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadConfigWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetId

func (v *dbusIfc) GoGetId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetId", flags, ch)
}

func (v *dbusIfc) GoGetIdWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetId", flags, ch)
}

func (*dbusIfc) StoreGetId(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) GetId(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetId(
		<-v.GoGetId(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *dbusIfc) GetIdWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetIdWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetId(call)
}

// method GetConnectionCredentials

func (v *dbusIfc) GoGetConnectionCredentials(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionCredentials", flags, ch, arg0)
}

func (v *dbusIfc) GoGetConnectionCredentialsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetConnectionCredentials", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionCredentials(call *dbus.Call) (arg1 map[string]dbus.Variant, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionCredentials(flags dbus.Flags, arg0 string) (arg1 map[string]dbus.Variant, err error) {
	return v.StoreGetConnectionCredentials(
		<-v.GoGetConnectionCredentials(flags, make(chan *dbus.Call, 1), arg0).Done)
}

func (v *dbusIfc) GetConnectionCredentialsWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string) (arg1 map[string]dbus.Variant, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetConnectionCredentialsWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetConnectionCredentials(call)
}

// signal NameOwnerChanged

func (v *dbusIfc) ConnectNameOwnerChanged(cb func(arg0 string, arg1 string, arg2 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameOwnerChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameOwnerChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		var arg2 string
		err := dbus.Store(sig.Body, &arg0, &arg1, &arg2)
		if err == nil {
			cb(arg0, arg1, arg2)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameLost

func (v *dbusIfc) ConnectNameLost(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameLost", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameLost",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameAcquired

func (v *dbusIfc) ConnectNameAcquired(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameAcquired", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameAcquired",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
