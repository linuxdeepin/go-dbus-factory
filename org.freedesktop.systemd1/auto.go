package systemd1

import "context"
import "errors"
import "fmt"
import dbus "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Manager struct {
	manager // interface org.freedesktop.systemd1.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.systemd1.Manager"
}

// method GetUnit

func (v *manager) GoGetUnit(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnit", flags, ch, name)
}

func (v *manager) GoGetUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnit", flags, ch, name)
}

func (*manager) StoreGetUnit(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) GetUnit(flags dbus.Flags, name string) (unit dbus.ObjectPath, err error) {
	return v.StoreGetUnit(
		<-v.GoGetUnit(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *manager) GetUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnit(call)
}

// method GetUnitByPID

func (v *manager) GoGetUnitByPID(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitByPID", flags, ch, pid)
}

func (v *manager) GoGetUnitByPIDWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitByPID", flags, ch, pid)
}

func (*manager) StoreGetUnitByPID(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) GetUnitByPID(flags dbus.Flags, pid uint32) (unit dbus.ObjectPath, err error) {
	return v.StoreGetUnitByPID(
		<-v.GoGetUnitByPID(flags, make(chan *dbus.Call, 1), pid).Done)
}

func (v *manager) GetUnitByPIDWithTimeout(timeout time.Duration, flags dbus.Flags, pid uint32) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitByPIDWithContext(ctx, flags, make(chan *dbus.Call, 1), pid).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitByPID(call)
}

// method GetUnitByInvocationID

func (v *manager) GoGetUnitByInvocationID(flags dbus.Flags, ch chan *dbus.Call, invocationID []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitByInvocationID", flags, ch, invocationID)
}

func (v *manager) GoGetUnitByInvocationIDWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, invocationID []uint8) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitByInvocationID", flags, ch, invocationID)
}

func (*manager) StoreGetUnitByInvocationID(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) GetUnitByInvocationID(flags dbus.Flags, invocationID []uint8) (unit dbus.ObjectPath, err error) {
	return v.StoreGetUnitByInvocationID(
		<-v.GoGetUnitByInvocationID(flags, make(chan *dbus.Call, 1), invocationID).Done)
}

func (v *manager) GetUnitByInvocationIDWithTimeout(timeout time.Duration, flags dbus.Flags, invocationID []uint8) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitByInvocationIDWithContext(ctx, flags, make(chan *dbus.Call, 1), invocationID).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitByInvocationID(call)
}

// method GetUnitByControlGroup

func (v *manager) GoGetUnitByControlGroup(flags dbus.Flags, ch chan *dbus.Call, ctrlGroup string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitByControlGroup", flags, ch, ctrlGroup)
}

func (v *manager) GoGetUnitByControlGroupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, ctrlGroup string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitByControlGroup", flags, ch, ctrlGroup)
}

func (*manager) StoreGetUnitByControlGroup(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) GetUnitByControlGroup(flags dbus.Flags, ctrlGroup string) (unit dbus.ObjectPath, err error) {
	return v.StoreGetUnitByControlGroup(
		<-v.GoGetUnitByControlGroup(flags, make(chan *dbus.Call, 1), ctrlGroup).Done)
}

func (v *manager) GetUnitByControlGroupWithTimeout(timeout time.Duration, flags dbus.Flags, ctrlGroup string) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitByControlGroupWithContext(ctx, flags, make(chan *dbus.Call, 1), ctrlGroup).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitByControlGroup(call)
}

// method LoadUnit

func (v *manager) GoLoadUnit(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LoadUnit", flags, ch, name)
}

func (v *manager) GoLoadUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LoadUnit", flags, ch, name)
}

func (*manager) StoreLoadUnit(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) LoadUnit(flags dbus.Flags, name string) (unit dbus.ObjectPath, err error) {
	return v.StoreLoadUnit(
		<-v.GoLoadUnit(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *manager) LoadUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLoadUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLoadUnit(call)
}

// method StartUnit

func (v *manager) GoStartUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartUnit", flags, ch, name, mode)
}

func (v *manager) GoStartUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartUnit", flags, ch, name, mode)
}

func (*manager) StoreStartUnit(call *dbus.Call) (unit dbus.ObjectPath, err error) {
	err = call.Store(&unit)
	return
}

func (v *manager) StartUnit(flags dbus.Flags, name string, mode string) (unit dbus.ObjectPath, err error) {
	return v.StoreStartUnit(
		<-v.GoStartUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) StartUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (unit dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStartUnit(call)
}

// method StartUnitReplace

func (v *manager) GoStartUnitReplace(flags dbus.Flags, ch chan *dbus.Call, oldUnit string, newUnit string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartUnitReplace", flags, ch, oldUnit, newUnit, mode)
}

func (v *manager) GoStartUnitReplaceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, oldUnit string, newUnit string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartUnitReplace", flags, ch, oldUnit, newUnit, mode)
}

func (*manager) StoreStartUnitReplace(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) StartUnitReplace(flags dbus.Flags, oldUnit string, newUnit string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreStartUnitReplace(
		<-v.GoStartUnitReplace(flags, make(chan *dbus.Call, 1), oldUnit, newUnit, mode).Done)
}

func (v *manager) StartUnitReplaceWithTimeout(timeout time.Duration, flags dbus.Flags, oldUnit string, newUnit string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartUnitReplaceWithContext(ctx, flags, make(chan *dbus.Call, 1), oldUnit, newUnit, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStartUnitReplace(call)
}

// method StopUnit

func (v *manager) GoStopUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopUnit", flags, ch, name, mode)
}

func (v *manager) GoStopUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StopUnit", flags, ch, name, mode)
}

func (*manager) StoreStopUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) StopUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreStopUnit(
		<-v.GoStopUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) StopUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStopUnit(call)
}

// method ReloadUnit

func (v *manager) GoReloadUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadUnit", flags, ch, name, mode)
}

func (v *manager) GoReloadUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadUnit", flags, ch, name, mode)
}

func (*manager) StoreReloadUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) ReloadUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReloadUnit(
		<-v.GoReloadUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) ReloadUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReloadUnit(call)
}

// method RestartUnit

func (v *manager) GoRestartUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestartUnit", flags, ch, name, mode)
}

func (v *manager) GoRestartUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RestartUnit", flags, ch, name, mode)
}

func (*manager) StoreRestartUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) RestartUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreRestartUnit(
		<-v.GoRestartUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) RestartUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRestartUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRestartUnit(call)
}

// method TryRestartUnit

func (v *manager) GoTryRestartUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TryRestartUnit", flags, ch, name, mode)
}

func (v *manager) GoTryRestartUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".TryRestartUnit", flags, ch, name, mode)
}

func (*manager) StoreTryRestartUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) TryRestartUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreTryRestartUnit(
		<-v.GoTryRestartUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) TryRestartUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTryRestartUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreTryRestartUnit(call)
}

// method ReloadOrRestartUnit

func (v *manager) GoReloadOrRestartUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadOrRestartUnit", flags, ch, name, mode)
}

func (v *manager) GoReloadOrRestartUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadOrRestartUnit", flags, ch, name, mode)
}

func (*manager) StoreReloadOrRestartUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) ReloadOrRestartUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReloadOrRestartUnit(
		<-v.GoReloadOrRestartUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) ReloadOrRestartUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadOrRestartUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReloadOrRestartUnit(call)
}

// method ReloadOrTryRestartUnit

func (v *manager) GoReloadOrTryRestartUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadOrTryRestartUnit", flags, ch, name, mode)
}

func (v *manager) GoReloadOrTryRestartUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadOrTryRestartUnit", flags, ch, name, mode)
}

func (*manager) StoreReloadOrTryRestartUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) ReloadOrTryRestartUnit(flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReloadOrTryRestartUnit(
		<-v.GoReloadOrTryRestartUnit(flags, make(chan *dbus.Call, 1), name, mode).Done)
}

func (v *manager) ReloadOrTryRestartUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadOrTryRestartUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReloadOrTryRestartUnit(call)
}

// method KillUnit

func (v *manager) GoKillUnit(flags dbus.Flags, ch chan *dbus.Call, name string, who string, signal int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".KillUnit", flags, ch, name, who, signal)
}

func (v *manager) GoKillUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, who string, signal int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".KillUnit", flags, ch, name, who, signal)
}

func (v *manager) KillUnit(flags dbus.Flags, name string, who string, signal int32) error {
	return (<-v.GoKillUnit(flags, make(chan *dbus.Call, 1), name, who, signal).Done).Err
}

func (v *manager) KillUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, who string, signal int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoKillUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, who, signal).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ResetFailedUnit

func (v *manager) GoResetFailedUnit(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetFailedUnit", flags, ch, name)
}

func (v *manager) GoResetFailedUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ResetFailedUnit", flags, ch, name)
}

func (v *manager) ResetFailedUnit(flags dbus.Flags, name string) error {
	return (<-v.GoResetFailedUnit(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *manager) ResetFailedUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoResetFailedUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetUnitProperties

func (v *manager) GoSetUnitProperties(flags dbus.Flags, ch chan *dbus.Call, name string, runtime bool, properties []Property) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUnitProperties", flags, ch, name, runtime, properties)
}

func (v *manager) GoSetUnitPropertiesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, runtime bool, properties []Property) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetUnitProperties", flags, ch, name, runtime, properties)
}

func (v *manager) SetUnitProperties(flags dbus.Flags, name string, runtime bool, properties []Property) error {
	return (<-v.GoSetUnitProperties(flags, make(chan *dbus.Call, 1), name, runtime, properties).Done).Err
}

func (v *manager) SetUnitPropertiesWithTimeout(timeout time.Duration, flags dbus.Flags, name string, runtime bool, properties []Property) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetUnitPropertiesWithContext(ctx, flags, make(chan *dbus.Call, 1), name, runtime, properties).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RefUnit

func (v *manager) GoRefUnit(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefUnit", flags, ch, name)
}

func (v *manager) GoRefUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RefUnit", flags, ch, name)
}

func (v *manager) RefUnit(flags dbus.Flags, name string) error {
	return (<-v.GoRefUnit(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *manager) RefUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRefUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnrefUnit

func (v *manager) GoUnrefUnit(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnrefUnit", flags, ch, name)
}

func (v *manager) GoUnrefUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnrefUnit", flags, ch, name)
}

func (v *manager) UnrefUnit(flags dbus.Flags, name string) error {
	return (<-v.GoUnrefUnit(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *manager) UnrefUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnrefUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method StartTransientUnit

func (v *manager) GoStartTransientUnit(flags dbus.Flags, ch chan *dbus.Call, name string, mode string, properties []Property, aux []PropertyCollection) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartTransientUnit", flags, ch, name, mode, properties, aux)
}

func (v *manager) GoStartTransientUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, mode string, properties []Property, aux []PropertyCollection) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartTransientUnit", flags, ch, name, mode, properties, aux)
}

func (*manager) StoreStartTransientUnit(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) StartTransientUnit(flags dbus.Flags, name string, mode string, properties []Property, aux []PropertyCollection) (job dbus.ObjectPath, err error) {
	return v.StoreStartTransientUnit(
		<-v.GoStartTransientUnit(flags, make(chan *dbus.Call, 1), name, mode, properties, aux).Done)
}

func (v *manager) StartTransientUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, mode string, properties []Property, aux []PropertyCollection) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartTransientUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, mode, properties, aux).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStartTransientUnit(call)
}

// method GetUnitProcesses

func (v *manager) GoGetUnitProcesses(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitProcesses", flags, ch, name)
}

func (v *manager) GoGetUnitProcessesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitProcesses", flags, ch, name)
}

func (*manager) StoreGetUnitProcesses(call *dbus.Call) (processes []UnitProcess, err error) {
	err = call.Store(&processes)
	return
}

func (v *manager) GetUnitProcesses(flags dbus.Flags, name string) (processes []UnitProcess, err error) {
	return v.StoreGetUnitProcesses(
		<-v.GoGetUnitProcesses(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *manager) GetUnitProcessesWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (processes []UnitProcess, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitProcessesWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitProcesses(call)
}

// method AttachProcessesToUnit

func (v *manager) GoAttachProcessesToUnit(flags dbus.Flags, ch chan *dbus.Call, name string, path string, pids []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AttachProcessesToUnit", flags, ch, name, path, pids)
}

func (v *manager) GoAttachProcessesToUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, path string, pids []uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AttachProcessesToUnit", flags, ch, name, path, pids)
}

func (v *manager) AttachProcessesToUnit(flags dbus.Flags, name string, path string, pids []uint32) error {
	return (<-v.GoAttachProcessesToUnit(flags, make(chan *dbus.Call, 1), name, path, pids).Done).Err
}

func (v *manager) AttachProcessesToUnitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, path string, pids []uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAttachProcessesToUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, path, pids).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method AbandonScope

func (v *manager) GoAbandonScope(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AbandonScope", flags, ch, name)
}

func (v *manager) GoAbandonScopeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AbandonScope", flags, ch, name)
}

func (v *manager) AbandonScope(flags dbus.Flags, name string) error {
	return (<-v.GoAbandonScope(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *manager) AbandonScopeWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAbandonScopeWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetJob

func (v *manager) GoGetJob(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetJob", flags, ch, id)
}

func (v *manager) GoGetJobWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetJob", flags, ch, id)
}

func (*manager) StoreGetJob(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) GetJob(flags dbus.Flags, id uint32) (job dbus.ObjectPath, err error) {
	return v.StoreGetJob(
		<-v.GoGetJob(flags, make(chan *dbus.Call, 1), id).Done)
}

func (v *manager) GetJobWithTimeout(timeout time.Duration, flags dbus.Flags, id uint32) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetJobWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetJob(call)
}

// method GetJobAfter

func (v *manager) GoGetJobAfter(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetJobAfter", flags, ch, id)
}

func (v *manager) GoGetJobAfterWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetJobAfter", flags, ch, id)
}

func (*manager) StoreGetJobAfter(call *dbus.Call) (jobs []JobInfo, err error) {
	err = call.Store(&jobs)
	return
}

func (v *manager) GetJobAfter(flags dbus.Flags, id uint32) (jobs []JobInfo, err error) {
	return v.StoreGetJobAfter(
		<-v.GoGetJobAfter(flags, make(chan *dbus.Call, 1), id).Done)
}

func (v *manager) GetJobAfterWithTimeout(timeout time.Duration, flags dbus.Flags, id uint32) (jobs []JobInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetJobAfterWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetJobAfter(call)
}

// method GetJobBefore

func (v *manager) GoGetJobBefore(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetJobBefore", flags, ch, id)
}

func (v *manager) GoGetJobBeforeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetJobBefore", flags, ch, id)
}

func (*manager) StoreGetJobBefore(call *dbus.Call) (jobs []JobInfo, err error) {
	err = call.Store(&jobs)
	return
}

func (v *manager) GetJobBefore(flags dbus.Flags, id uint32) (jobs []JobInfo, err error) {
	return v.StoreGetJobBefore(
		<-v.GoGetJobBefore(flags, make(chan *dbus.Call, 1), id).Done)
}

func (v *manager) GetJobBeforeWithTimeout(timeout time.Duration, flags dbus.Flags, id uint32) (jobs []JobInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetJobBeforeWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetJobBefore(call)
}

// method CancelJob

func (v *manager) GoCancelJob(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelJob", flags, ch, id)
}

func (v *manager) GoCancelJobWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CancelJob", flags, ch, id)
}

func (v *manager) CancelJob(flags dbus.Flags, id uint32) error {
	return (<-v.GoCancelJob(flags, make(chan *dbus.Call, 1), id).Done).Err
}

func (v *manager) CancelJobWithTimeout(timeout time.Duration, flags dbus.Flags, id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCancelJobWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ClearJobs

func (v *manager) GoClearJobs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearJobs", flags, ch)
}

func (v *manager) GoClearJobsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ClearJobs", flags, ch)
}

func (v *manager) ClearJobs(flags dbus.Flags) error {
	return (<-v.GoClearJobs(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) ClearJobsWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClearJobsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ResetFailed

func (v *manager) GoResetFailed(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetFailed", flags, ch)
}

func (v *manager) GoResetFailedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ResetFailed", flags, ch)
}

func (v *manager) ResetFailed(flags dbus.Flags) error {
	return (<-v.GoResetFailed(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) ResetFailedWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoResetFailedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ListUnits

func (v *manager) GoListUnits(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnits", flags, ch)
}

func (v *manager) GoListUnitsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnits", flags, ch)
}

func (*manager) StoreListUnits(call *dbus.Call) (units []UnitInfo, err error) {
	err = call.Store(&units)
	return
}

func (v *manager) ListUnits(flags dbus.Flags) (units []UnitInfo, err error) {
	return v.StoreListUnits(
		<-v.GoListUnits(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) ListUnitsWithTimeout(timeout time.Duration, flags dbus.Flags) (units []UnitInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnits(call)
}

// method ListUnitsFiltered

func (v *manager) GoListUnitsFiltered(flags dbus.Flags, ch chan *dbus.Call, states []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnitsFiltered", flags, ch, states)
}

func (v *manager) GoListUnitsFilteredWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, states []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnitsFiltered", flags, ch, states)
}

func (*manager) StoreListUnitsFiltered(call *dbus.Call) (units []UnitInfo, err error) {
	err = call.Store(&units)
	return
}

func (v *manager) ListUnitsFiltered(flags dbus.Flags, states []string) (units []UnitInfo, err error) {
	return v.StoreListUnitsFiltered(
		<-v.GoListUnitsFiltered(flags, make(chan *dbus.Call, 1), states).Done)
}

func (v *manager) ListUnitsFilteredWithTimeout(timeout time.Duration, flags dbus.Flags, states []string) (units []UnitInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitsFilteredWithContext(ctx, flags, make(chan *dbus.Call, 1), states).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnitsFiltered(call)
}

// method ListUnitsByPatterns

func (v *manager) GoListUnitsByPatterns(flags dbus.Flags, ch chan *dbus.Call, states []string, patterns []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnitsByPatterns", flags, ch, states, patterns)
}

func (v *manager) GoListUnitsByPatternsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, states []string, patterns []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnitsByPatterns", flags, ch, states, patterns)
}

func (*manager) StoreListUnitsByPatterns(call *dbus.Call) (units []UnitInfo, err error) {
	err = call.Store(&units)
	return
}

func (v *manager) ListUnitsByPatterns(flags dbus.Flags, states []string, patterns []string) (units []UnitInfo, err error) {
	return v.StoreListUnitsByPatterns(
		<-v.GoListUnitsByPatterns(flags, make(chan *dbus.Call, 1), states, patterns).Done)
}

func (v *manager) ListUnitsByPatternsWithTimeout(timeout time.Duration, flags dbus.Flags, states []string, patterns []string) (units []UnitInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitsByPatternsWithContext(ctx, flags, make(chan *dbus.Call, 1), states, patterns).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnitsByPatterns(call)
}

// method ListUnitsByNames

func (v *manager) GoListUnitsByNames(flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnitsByNames", flags, ch, names)
}

func (v *manager) GoListUnitsByNamesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnitsByNames", flags, ch, names)
}

func (*manager) StoreListUnitsByNames(call *dbus.Call) (units []UnitInfo, err error) {
	err = call.Store(&units)
	return
}

func (v *manager) ListUnitsByNames(flags dbus.Flags, names []string) (units []UnitInfo, err error) {
	return v.StoreListUnitsByNames(
		<-v.GoListUnitsByNames(flags, make(chan *dbus.Call, 1), names).Done)
}

func (v *manager) ListUnitsByNamesWithTimeout(timeout time.Duration, flags dbus.Flags, names []string) (units []UnitInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitsByNamesWithContext(ctx, flags, make(chan *dbus.Call, 1), names).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnitsByNames(call)
}

// method ListJobs

func (v *manager) GoListJobs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListJobs", flags, ch)
}

func (v *manager) GoListJobsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListJobs", flags, ch)
}

func (*manager) StoreListJobs(call *dbus.Call) (jobs []JobInfo, err error) {
	err = call.Store(&jobs)
	return
}

func (v *manager) ListJobs(flags dbus.Flags) (jobs []JobInfo, err error) {
	return v.StoreListJobs(
		<-v.GoListJobs(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) ListJobsWithTimeout(timeout time.Duration, flags dbus.Flags) (jobs []JobInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListJobsWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListJobs(call)
}

// method Subscribe

func (v *manager) GoSubscribe(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Subscribe", flags, ch)
}

func (v *manager) GoSubscribeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Subscribe", flags, ch)
}

func (v *manager) Subscribe(flags dbus.Flags) error {
	return (<-v.GoSubscribe(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) SubscribeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSubscribeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Unsubscribe

func (v *manager) GoUnsubscribe(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unsubscribe", flags, ch)
}

func (v *manager) GoUnsubscribeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Unsubscribe", flags, ch)
}

func (v *manager) Unsubscribe(flags dbus.Flags) error {
	return (<-v.GoUnsubscribe(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) UnsubscribeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnsubscribeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Dump

func (v *manager) GoDump(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Dump", flags, ch)
}

func (v *manager) GoDumpWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Dump", flags, ch)
}

func (*manager) StoreDump(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *manager) Dump(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreDump(
		<-v.GoDump(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) DumpWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDumpWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreDump(call)
}

// method DumpByFileDescriptor

func (v *manager) GoDumpByFileDescriptor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpByFileDescriptor", flags, ch)
}

func (v *manager) GoDumpByFileDescriptorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DumpByFileDescriptor", flags, ch)
}

func (*manager) StoreDumpByFileDescriptor(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *manager) DumpByFileDescriptor(flags dbus.Flags) (fd dbus.UnixFD, err error) {
	return v.StoreDumpByFileDescriptor(
		<-v.GoDumpByFileDescriptor(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) DumpByFileDescriptorWithTimeout(timeout time.Duration, flags dbus.Flags) (fd dbus.UnixFD, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDumpByFileDescriptorWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreDumpByFileDescriptor(call)
}

// method Reload

func (v *manager) GoReload(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch)
}

func (v *manager) GoReloadWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reload", flags, ch)
}

func (v *manager) Reload(flags dbus.Flags) error {
	return (<-v.GoReload(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) ReloadWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Reexecute

func (v *manager) GoReexecute(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reexecute", flags, ch)
}

func (v *manager) GoReexecuteWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reexecute", flags, ch)
}

func (v *manager) Reexecute(flags dbus.Flags) error {
	return (<-v.GoReexecute(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) ReexecuteWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReexecuteWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Exit

func (v *manager) GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *manager) GoExitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *manager) Exit(flags dbus.Flags) error {
	return (<-v.GoExit(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) ExitWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoExitWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Reboot

func (v *manager) GoReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *manager) GoRebootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *manager) Reboot(flags dbus.Flags) error {
	return (<-v.GoReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) RebootWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRebootWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PowerOff

func (v *manager) GoPowerOff(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOff", flags, ch)
}

func (v *manager) GoPowerOffWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PowerOff", flags, ch)
}

func (v *manager) PowerOff(flags dbus.Flags) error {
	return (<-v.GoPowerOff(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) PowerOffWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPowerOffWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Halt

func (v *manager) GoHalt(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Halt", flags, ch)
}

func (v *manager) GoHaltWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Halt", flags, ch)
}

func (v *manager) Halt(flags dbus.Flags) error {
	return (<-v.GoHalt(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) HaltWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHaltWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method KExec

func (v *manager) GoKExec(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".KExec", flags, ch)
}

func (v *manager) GoKExecWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".KExec", flags, ch)
}

func (v *manager) KExec(flags dbus.Flags) error {
	return (<-v.GoKExec(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *manager) KExecWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoKExecWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SwitchRoot

func (v *manager) GoSwitchRoot(flags dbus.Flags, ch chan *dbus.Call, newRoot string, init string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchRoot", flags, ch, newRoot, init)
}

func (v *manager) GoSwitchRootWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, newRoot string, init string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SwitchRoot", flags, ch, newRoot, init)
}

func (v *manager) SwitchRoot(flags dbus.Flags, newRoot string, init string) error {
	return (<-v.GoSwitchRoot(flags, make(chan *dbus.Call, 1), newRoot, init).Done).Err
}

func (v *manager) SwitchRootWithTimeout(timeout time.Duration, flags dbus.Flags, newRoot string, init string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSwitchRootWithContext(ctx, flags, make(chan *dbus.Call, 1), newRoot, init).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetEnvironment

func (v *manager) GoSetEnvironment(flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetEnvironment", flags, ch, names)
}

func (v *manager) GoSetEnvironmentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetEnvironment", flags, ch, names)
}

func (v *manager) SetEnvironment(flags dbus.Flags, names []string) error {
	return (<-v.GoSetEnvironment(flags, make(chan *dbus.Call, 1), names).Done).Err
}

func (v *manager) SetEnvironmentWithTimeout(timeout time.Duration, flags dbus.Flags, names []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetEnvironmentWithContext(ctx, flags, make(chan *dbus.Call, 1), names).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnsetEnvironment

func (v *manager) GoUnsetEnvironment(flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnsetEnvironment", flags, ch, names)
}

func (v *manager) GoUnsetEnvironmentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, names []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnsetEnvironment", flags, ch, names)
}

func (v *manager) UnsetEnvironment(flags dbus.Flags, names []string) error {
	return (<-v.GoUnsetEnvironment(flags, make(chan *dbus.Call, 1), names).Done).Err
}

func (v *manager) UnsetEnvironmentWithTimeout(timeout time.Duration, flags dbus.Flags, names []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnsetEnvironmentWithContext(ctx, flags, make(chan *dbus.Call, 1), names).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnsetAndSetEnvironment

func (v *manager) GoUnsetAndSetEnvironment(flags dbus.Flags, ch chan *dbus.Call, unset []string, set []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnsetAndSetEnvironment", flags, ch, unset, set)
}

func (v *manager) GoUnsetAndSetEnvironmentWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, unset []string, set []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnsetAndSetEnvironment", flags, ch, unset, set)
}

func (v *manager) UnsetAndSetEnvironment(flags dbus.Flags, unset []string, set []string) error {
	return (<-v.GoUnsetAndSetEnvironment(flags, make(chan *dbus.Call, 1), unset, set).Done).Err
}

func (v *manager) UnsetAndSetEnvironmentWithTimeout(timeout time.Duration, flags dbus.Flags, unset []string, set []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnsetAndSetEnvironmentWithContext(ctx, flags, make(chan *dbus.Call, 1), unset, set).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ListUnitFiles

func (v *manager) GoListUnitFiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnitFiles", flags, ch)
}

func (v *manager) GoListUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnitFiles", flags, ch)
}

func (*manager) StoreListUnitFiles(call *dbus.Call) (files []UnitFile, err error) {
	err = call.Store(&files)
	return
}

func (v *manager) ListUnitFiles(flags dbus.Flags) (files []UnitFile, err error) {
	return v.StoreListUnitFiles(
		<-v.GoListUnitFiles(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) ListUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags) (files []UnitFile, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnitFiles(call)
}

// method ListUnitFilesByPatterns

func (v *manager) GoListUnitFilesByPatterns(flags dbus.Flags, ch chan *dbus.Call, states []string, patterns []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUnitFilesByPatterns", flags, ch, states, patterns)
}

func (v *manager) GoListUnitFilesByPatternsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, states []string, patterns []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListUnitFilesByPatterns", flags, ch, states, patterns)
}

func (*manager) StoreListUnitFilesByPatterns(call *dbus.Call) (files []UnitFile, err error) {
	err = call.Store(&files)
	return
}

func (v *manager) ListUnitFilesByPatterns(flags dbus.Flags, states []string, patterns []string) (files []UnitFile, err error) {
	return v.StoreListUnitFilesByPatterns(
		<-v.GoListUnitFilesByPatterns(flags, make(chan *dbus.Call, 1), states, patterns).Done)
}

func (v *manager) ListUnitFilesByPatternsWithTimeout(timeout time.Duration, flags dbus.Flags, states []string, patterns []string) (files []UnitFile, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListUnitFilesByPatternsWithContext(ctx, flags, make(chan *dbus.Call, 1), states, patterns).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListUnitFilesByPatterns(call)
}

// method GetUnitFileState

func (v *manager) GoGetUnitFileState(flags dbus.Flags, ch chan *dbus.Call, unit string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitFileState", flags, ch, unit)
}

func (v *manager) GoGetUnitFileStateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, unit string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitFileState", flags, ch, unit)
}

func (*manager) StoreGetUnitFileState(call *dbus.Call) (state string, err error) {
	err = call.Store(&state)
	return
}

func (v *manager) GetUnitFileState(flags dbus.Flags, unit string) (state string, err error) {
	return v.StoreGetUnitFileState(
		<-v.GoGetUnitFileState(flags, make(chan *dbus.Call, 1), unit).Done)
}

func (v *manager) GetUnitFileStateWithTimeout(timeout time.Duration, flags dbus.Flags, unit string) (state string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitFileStateWithContext(ctx, flags, make(chan *dbus.Call, 1), unit).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitFileState(call)
}

// method EnableUnitFiles

func (v *manager) GoEnableUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableUnitFiles", flags, ch, files, runtime, force)
}

func (v *manager) GoEnableUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableUnitFiles", flags, ch, files, runtime, force)
}

func (*manager) StoreEnableUnitFiles(call *dbus.Call) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	err = call.Store(&carriesInstallInfo, &changes)
	return
}

func (v *manager) EnableUnitFiles(flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	return v.StoreEnableUnitFiles(
		<-v.GoEnableUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime, force).Done)
}

func (v *manager) EnableUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreEnableUnitFiles(call)
}

// method DisableUnitFiles

func (v *manager) GoDisableUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DisableUnitFiles", flags, ch, files, runtime)
}

func (v *manager) GoDisableUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DisableUnitFiles", flags, ch, files, runtime)
}

func (*manager) StoreDisableUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) DisableUnitFiles(flags dbus.Flags, files []string, runtime bool) (changes []UnitFileChange, err error) {
	return v.StoreDisableUnitFiles(
		<-v.GoDisableUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime).Done)
}

func (v *manager) DisableUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDisableUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreDisableUnitFiles(call)
}

// method ReenableUnitFiles

func (v *manager) GoReenableUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReenableUnitFiles", flags, ch, files, runtime, force)
}

func (v *manager) GoReenableUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReenableUnitFiles", flags, ch, files, runtime, force)
}

func (*manager) StoreReenableUnitFiles(call *dbus.Call) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	err = call.Store(&carriesInstallInfo, &changes)
	return
}

func (v *manager) ReenableUnitFiles(flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	return v.StoreReenableUnitFiles(
		<-v.GoReenableUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime, force).Done)
}

func (v *manager) ReenableUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReenableUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReenableUnitFiles(call)
}

// method LinkUnitFiles

func (v *manager) GoLinkUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LinkUnitFiles", flags, ch, files, runtime, force)
}

func (v *manager) GoLinkUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LinkUnitFiles", flags, ch, files, runtime, force)
}

func (*manager) StoreLinkUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) LinkUnitFiles(flags dbus.Flags, files []string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	return v.StoreLinkUnitFiles(
		<-v.GoLinkUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime, force).Done)
}

func (v *manager) LinkUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLinkUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLinkUnitFiles(call)
}

// method PresetUnitFiles

func (v *manager) GoPresetUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresetUnitFiles", flags, ch, files, runtime, force)
}

func (v *manager) GoPresetUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PresetUnitFiles", flags, ch, files, runtime, force)
}

func (*manager) StorePresetUnitFiles(call *dbus.Call) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	err = call.Store(&carriesInstallInfo, &changes)
	return
}

func (v *manager) PresetUnitFiles(flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	return v.StorePresetUnitFiles(
		<-v.GoPresetUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime, force).Done)
}

func (v *manager) PresetUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPresetUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePresetUnitFiles(call)
}

// method PresetUnitFilesWithMode

func (v *manager) GoPresetUnitFilesWithMode(flags dbus.Flags, ch chan *dbus.Call, files []string, mode string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresetUnitFilesWithMode", flags, ch, files, mode, runtime, force)
}

func (v *manager) GoPresetUnitFilesWithModeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, mode string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PresetUnitFilesWithMode", flags, ch, files, mode, runtime, force)
}

func (*manager) StorePresetUnitFilesWithMode(call *dbus.Call) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	err = call.Store(&carriesInstallInfo, &changes)
	return
}

func (v *manager) PresetUnitFilesWithMode(flags dbus.Flags, files []string, mode string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	return v.StorePresetUnitFilesWithMode(
		<-v.GoPresetUnitFilesWithMode(flags, make(chan *dbus.Call, 1), files, mode, runtime, force).Done)
}

func (v *manager) PresetUnitFilesWithModeWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, mode string, runtime bool, force bool) (carriesInstallInfo bool, changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPresetUnitFilesWithModeWithContext(ctx, flags, make(chan *dbus.Call, 1), files, mode, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePresetUnitFilesWithMode(call)
}

// method MaskUnitFiles

func (v *manager) GoMaskUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MaskUnitFiles", flags, ch, files, runtime, force)
}

func (v *manager) GoMaskUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MaskUnitFiles", flags, ch, files, runtime, force)
}

func (*manager) StoreMaskUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) MaskUnitFiles(flags dbus.Flags, files []string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	return v.StoreMaskUnitFiles(
		<-v.GoMaskUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime, force).Done)
}

func (v *manager) MaskUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMaskUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreMaskUnitFiles(call)
}

// method UnmaskUnitFiles

func (v *manager) GoUnmaskUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnmaskUnitFiles", flags, ch, files, runtime)
}

func (v *manager) GoUnmaskUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, runtime bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnmaskUnitFiles", flags, ch, files, runtime)
}

func (*manager) StoreUnmaskUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) UnmaskUnitFiles(flags dbus.Flags, files []string, runtime bool) (changes []UnitFileChange, err error) {
	return v.StoreUnmaskUnitFiles(
		<-v.GoUnmaskUnitFiles(flags, make(chan *dbus.Call, 1), files, runtime).Done)
}

func (v *manager) UnmaskUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, runtime bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnmaskUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, runtime).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreUnmaskUnitFiles(call)
}

// method RevertUnitFiles

func (v *manager) GoRevertUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RevertUnitFiles", flags, ch, files)
}

func (v *manager) GoRevertUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RevertUnitFiles", flags, ch, files)
}

func (*manager) StoreRevertUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) RevertUnitFiles(flags dbus.Flags, files []string) (changes []UnitFileChange, err error) {
	return v.StoreRevertUnitFiles(
		<-v.GoRevertUnitFiles(flags, make(chan *dbus.Call, 1), files).Done)
}

func (v *manager) RevertUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRevertUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRevertUnitFiles(call)
}

// method SetDefaultTarget

func (v *manager) GoSetDefaultTarget(flags dbus.Flags, ch chan *dbus.Call, name string, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultTarget", flags, ch, name, force)
}

func (v *manager) GoSetDefaultTargetWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetDefaultTarget", flags, ch, name, force)
}

func (*manager) StoreSetDefaultTarget(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) SetDefaultTarget(flags dbus.Flags, name string, force bool) (changes []UnitFileChange, err error) {
	return v.StoreSetDefaultTarget(
		<-v.GoSetDefaultTarget(flags, make(chan *dbus.Call, 1), name, force).Done)
}

func (v *manager) SetDefaultTargetWithTimeout(timeout time.Duration, flags dbus.Flags, name string, force bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetDefaultTargetWithContext(ctx, flags, make(chan *dbus.Call, 1), name, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSetDefaultTarget(call)
}

// method GetDefaultTarget

func (v *manager) GoGetDefaultTarget(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultTarget", flags, ch)
}

func (v *manager) GoGetDefaultTargetWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDefaultTarget", flags, ch)
}

func (*manager) StoreGetDefaultTarget(call *dbus.Call) (name string, err error) {
	err = call.Store(&name)
	return
}

func (v *manager) GetDefaultTarget(flags dbus.Flags) (name string, err error) {
	return v.StoreGetDefaultTarget(
		<-v.GoGetDefaultTarget(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) GetDefaultTargetWithTimeout(timeout time.Duration, flags dbus.Flags) (name string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetDefaultTargetWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetDefaultTarget(call)
}

// method PresetAllUnitFiles

func (v *manager) GoPresetAllUnitFiles(flags dbus.Flags, ch chan *dbus.Call, mode string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresetAllUnitFiles", flags, ch, mode, runtime, force)
}

func (v *manager) GoPresetAllUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PresetAllUnitFiles", flags, ch, mode, runtime, force)
}

func (*manager) StorePresetAllUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) PresetAllUnitFiles(flags dbus.Flags, mode string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	return v.StorePresetAllUnitFiles(
		<-v.GoPresetAllUnitFiles(flags, make(chan *dbus.Call, 1), mode, runtime, force).Done)
}

func (v *manager) PresetAllUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, mode string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPresetAllUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), mode, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePresetAllUnitFiles(call)
}

// method AddDependencyUnitFiles

func (v *manager) GoAddDependencyUnitFiles(flags dbus.Flags, ch chan *dbus.Call, files []string, target string, type0 string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddDependencyUnitFiles", flags, ch, files, target, type0, runtime, force)
}

func (v *manager) GoAddDependencyUnitFilesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, files []string, target string, type0 string, runtime bool, force bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddDependencyUnitFiles", flags, ch, files, target, type0, runtime, force)
}

func (*manager) StoreAddDependencyUnitFiles(call *dbus.Call) (changes []UnitFileChange, err error) {
	err = call.Store(&changes)
	return
}

func (v *manager) AddDependencyUnitFiles(flags dbus.Flags, files []string, target string, type0 string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	return v.StoreAddDependencyUnitFiles(
		<-v.GoAddDependencyUnitFiles(flags, make(chan *dbus.Call, 1), files, target, type0, runtime, force).Done)
}

func (v *manager) AddDependencyUnitFilesWithTimeout(timeout time.Duration, flags dbus.Flags, files []string, target string, type0 string, runtime bool, force bool) (changes []UnitFileChange, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddDependencyUnitFilesWithContext(ctx, flags, make(chan *dbus.Call, 1), files, target, type0, runtime, force).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreAddDependencyUnitFiles(call)
}

// method GetUnitFileLinks

func (v *manager) GoGetUnitFileLinks(flags dbus.Flags, ch chan *dbus.Call, name string, runtime bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUnitFileLinks", flags, ch, name, runtime)
}

func (v *manager) GoGetUnitFileLinksWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, runtime bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUnitFileLinks", flags, ch, name, runtime)
}

func (*manager) StoreGetUnitFileLinks(call *dbus.Call) (links []string, err error) {
	err = call.Store(&links)
	return
}

func (v *manager) GetUnitFileLinks(flags dbus.Flags, name string, runtime bool) (links []string, err error) {
	return v.StoreGetUnitFileLinks(
		<-v.GoGetUnitFileLinks(flags, make(chan *dbus.Call, 1), name, runtime).Done)
}

func (v *manager) GetUnitFileLinksWithTimeout(timeout time.Duration, flags dbus.Flags, name string, runtime bool) (links []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUnitFileLinksWithContext(ctx, flags, make(chan *dbus.Call, 1), name, runtime).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUnitFileLinks(call)
}

// method SetExitCode

func (v *manager) GoSetExitCode(flags dbus.Flags, ch chan *dbus.Call, exitCode uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetExitCode", flags, ch, exitCode)
}

func (v *manager) GoSetExitCodeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, exitCode uint8) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetExitCode", flags, ch, exitCode)
}

func (v *manager) SetExitCode(flags dbus.Flags, exitCode uint8) error {
	return (<-v.GoSetExitCode(flags, make(chan *dbus.Call, 1), exitCode).Done).Err
}

func (v *manager) SetExitCodeWithTimeout(timeout time.Duration, flags dbus.Flags, exitCode uint8) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetExitCodeWithContext(ctx, flags, make(chan *dbus.Call, 1), exitCode).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method LookupDynamicUserByName

func (v *manager) GoLookupDynamicUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LookupDynamicUserByName", flags, ch, name)
}

func (v *manager) GoLookupDynamicUserByNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LookupDynamicUserByName", flags, ch, name)
}

func (*manager) StoreLookupDynamicUserByName(call *dbus.Call) (user uint32, err error) {
	err = call.Store(&user)
	return
}

func (v *manager) LookupDynamicUserByName(flags dbus.Flags, name string) (user uint32, err error) {
	return v.StoreLookupDynamicUserByName(
		<-v.GoLookupDynamicUserByName(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *manager) LookupDynamicUserByNameWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (user uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLookupDynamicUserByNameWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLookupDynamicUserByName(call)
}

// method LookupDynamicUserByUID

func (v *manager) GoLookupDynamicUserByUID(flags dbus.Flags, ch chan *dbus.Call, uid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LookupDynamicUserByUID", flags, ch, uid)
}

func (v *manager) GoLookupDynamicUserByUIDWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uid uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LookupDynamicUserByUID", flags, ch, uid)
}

func (*manager) StoreLookupDynamicUserByUID(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *manager) LookupDynamicUserByUID(flags dbus.Flags, uid uint32) (user string, err error) {
	return v.StoreLookupDynamicUserByUID(
		<-v.GoLookupDynamicUserByUID(flags, make(chan *dbus.Call, 1), uid).Done)
}

func (v *manager) LookupDynamicUserByUIDWithTimeout(timeout time.Duration, flags dbus.Flags, uid uint32) (user string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLookupDynamicUserByUIDWithContext(ctx, flags, make(chan *dbus.Call, 1), uid).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLookupDynamicUserByUID(call)
}

// method GetDynamicUsers

func (v *manager) GoGetDynamicUsers(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDynamicUsers", flags, ch)
}

func (v *manager) GoGetDynamicUsersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDynamicUsers", flags, ch)
}

func (*manager) StoreGetDynamicUsers(call *dbus.Call) (users []DynamicUser, err error) {
	err = call.Store(&users)
	return
}

func (v *manager) GetDynamicUsers(flags dbus.Flags) (users []DynamicUser, err error) {
	return v.StoreGetDynamicUsers(
		<-v.GoGetDynamicUsers(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) GetDynamicUsersWithTimeout(timeout time.Duration, flags dbus.Flags) (users []DynamicUser, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetDynamicUsersWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetDynamicUsers(call)
}

// signal UnitNew

func (v *manager) ConnectUnitNew(cb func(id string, unit dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UnitNew", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UnitNew",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var unit dbus.ObjectPath
		err := dbus.Store(sig.Body, &id, &unit)
		if err == nil {
			cb(id, unit)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UnitRemoved

func (v *manager) ConnectUnitRemoved(cb func(id string, unit dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UnitRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UnitRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var unit dbus.ObjectPath
		err := dbus.Store(sig.Body, &id, &unit)
		if err == nil {
			cb(id, unit)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal JobNew

func (v *manager) ConnectJobNew(cb func(id uint32, job dbus.ObjectPath, unit string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "JobNew", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".JobNew",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id uint32
		var job dbus.ObjectPath
		var unit string
		err := dbus.Store(sig.Body, &id, &job, &unit)
		if err == nil {
			cb(id, job, unit)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal JobRemoved

func (v *manager) ConnectJobRemoved(cb func(id uint32, job dbus.ObjectPath, unit string, result string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "JobRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".JobRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id uint32
		var job dbus.ObjectPath
		var unit string
		var result string
		err := dbus.Store(sig.Body, &id, &job, &unit, &result)
		if err == nil {
			cb(id, job, unit, result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal StartupFinished

func (v *manager) ConnectStartupFinished(cb func(firmware uint64, loader uint64, kernel uint64, initrd uint64, userspace uint64, total uint64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StartupFinished", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StartupFinished",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var firmware uint64
		var loader uint64
		var kernel uint64
		var initrd uint64
		var userspace uint64
		var total uint64
		err := dbus.Store(sig.Body, &firmware, &loader, &kernel, &initrd, &userspace, &total)
		if err == nil {
			cb(firmware, loader, kernel, initrd, userspace, total)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UnitFilesChanged

func (v *manager) ConnectUnitFilesChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UnitFilesChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UnitFilesChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Reloading

func (v *manager) ConnectReloading(cb func(active bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Reloading", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Reloading",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var active bool
		err := dbus.Store(sig.Body, &active)
		if err == nil {
			cb(active)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Version s

func (v *manager) Version() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Version",
	}
}

// property Features s

func (v *manager) Features() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Features",
	}
}

// property Virtualization s

func (v *manager) Virtualization() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Virtualization",
	}
}

// property Architecture s

func (v *manager) Architecture() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Architecture",
	}
}

// property Tainted s

func (v *manager) Tainted() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Tainted",
	}
}

// property FirmwareTimestamp t

func (v *manager) FirmwareTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "FirmwareTimestamp",
	}
}

// property FirmwareTimestampMonotonic t

func (v *manager) FirmwareTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "FirmwareTimestampMonotonic",
	}
}

// property LoaderTimestamp t

func (v *manager) LoaderTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LoaderTimestamp",
	}
}

// property LoaderTimestampMonotonic t

func (v *manager) LoaderTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LoaderTimestampMonotonic",
	}
}

// property KernelTimestamp t

func (v *manager) KernelTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "KernelTimestamp",
	}
}

// property KernelTimestampMonotonic t

func (v *manager) KernelTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "KernelTimestampMonotonic",
	}
}

// property InitRDTimestamp t

func (v *manager) InitRDTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDTimestamp",
	}
}

// property InitRDTimestampMonotonic t

func (v *manager) InitRDTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDTimestampMonotonic",
	}
}

// property UserspaceTimestamp t

func (v *manager) UserspaceTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UserspaceTimestamp",
	}
}

// property UserspaceTimestampMonotonic t

func (v *manager) UserspaceTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UserspaceTimestampMonotonic",
	}
}

// property FinishTimestamp t

func (v *manager) FinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "FinishTimestamp",
	}
}

// property FinishTimestampMonotonic t

func (v *manager) FinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "FinishTimestampMonotonic",
	}
}

// property SecurityStartTimestamp t

func (v *manager) SecurityStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SecurityStartTimestamp",
	}
}

// property SecurityStartTimestampMonotonic t

func (v *manager) SecurityStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SecurityStartTimestampMonotonic",
	}
}

// property SecurityFinishTimestamp t

func (v *manager) SecurityFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SecurityFinishTimestamp",
	}
}

// property SecurityFinishTimestampMonotonic t

func (v *manager) SecurityFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SecurityFinishTimestampMonotonic",
	}
}

// property GeneratorsStartTimestamp t

func (v *manager) GeneratorsStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "GeneratorsStartTimestamp",
	}
}

// property GeneratorsStartTimestampMonotonic t

func (v *manager) GeneratorsStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "GeneratorsStartTimestampMonotonic",
	}
}

// property GeneratorsFinishTimestamp t

func (v *manager) GeneratorsFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "GeneratorsFinishTimestamp",
	}
}

// property GeneratorsFinishTimestampMonotonic t

func (v *manager) GeneratorsFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "GeneratorsFinishTimestampMonotonic",
	}
}

// property UnitsLoadStartTimestamp t

func (v *manager) UnitsLoadStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UnitsLoadStartTimestamp",
	}
}

// property UnitsLoadStartTimestampMonotonic t

func (v *manager) UnitsLoadStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UnitsLoadStartTimestampMonotonic",
	}
}

// property UnitsLoadFinishTimestamp t

func (v *manager) UnitsLoadFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UnitsLoadFinishTimestamp",
	}
}

// property UnitsLoadFinishTimestampMonotonic t

func (v *manager) UnitsLoadFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UnitsLoadFinishTimestampMonotonic",
	}
}

// property InitRDSecurityStartTimestamp t

func (v *manager) InitRDSecurityStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDSecurityStartTimestamp",
	}
}

// property InitRDSecurityStartTimestampMonotonic t

func (v *manager) InitRDSecurityStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDSecurityStartTimestampMonotonic",
	}
}

// property InitRDSecurityFinishTimestamp t

func (v *manager) InitRDSecurityFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDSecurityFinishTimestamp",
	}
}

// property InitRDSecurityFinishTimestampMonotonic t

func (v *manager) InitRDSecurityFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDSecurityFinishTimestampMonotonic",
	}
}

// property InitRDGeneratorsStartTimestamp t

func (v *manager) InitRDGeneratorsStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDGeneratorsStartTimestamp",
	}
}

// property InitRDGeneratorsStartTimestampMonotonic t

func (v *manager) InitRDGeneratorsStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDGeneratorsStartTimestampMonotonic",
	}
}

// property InitRDGeneratorsFinishTimestamp t

func (v *manager) InitRDGeneratorsFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDGeneratorsFinishTimestamp",
	}
}

// property InitRDGeneratorsFinishTimestampMonotonic t

func (v *manager) InitRDGeneratorsFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDGeneratorsFinishTimestampMonotonic",
	}
}

// property InitRDUnitsLoadStartTimestamp t

func (v *manager) InitRDUnitsLoadStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDUnitsLoadStartTimestamp",
	}
}

// property InitRDUnitsLoadStartTimestampMonotonic t

func (v *manager) InitRDUnitsLoadStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDUnitsLoadStartTimestampMonotonic",
	}
}

// property InitRDUnitsLoadFinishTimestamp t

func (v *manager) InitRDUnitsLoadFinishTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDUnitsLoadFinishTimestamp",
	}
}

// property InitRDUnitsLoadFinishTimestampMonotonic t

func (v *manager) InitRDUnitsLoadFinishTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InitRDUnitsLoadFinishTimestampMonotonic",
	}
}

// property LogLevel s

func (v *manager) LogLevel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "LogLevel",
	}
}

// property LogTarget s

func (v *manager) LogTarget() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "LogTarget",
	}
}

// property NNames u

func (v *manager) NNames() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NNames",
	}
}

// property NFailedUnits u

func (v *manager) NFailedUnits() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NFailedUnits",
	}
}

// property NJobs u

func (v *manager) NJobs() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NJobs",
	}
}

// property NInstalledJobs u

func (v *manager) NInstalledJobs() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NInstalledJobs",
	}
}

// property NFailedJobs u

func (v *manager) NFailedJobs() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NFailedJobs",
	}
}

// property Progress d

func (v *manager) Progress() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Progress",
	}
}

// property Environment as

func (v *manager) Environment() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Environment",
	}
}

// property ConfirmSpawn b

func (v *manager) ConfirmSpawn() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ConfirmSpawn",
	}
}

// property ShowStatus b

func (v *manager) ShowStatus() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ShowStatus",
	}
}

// property UnitPath as

func (v *manager) UnitPath() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UnitPath",
	}
}

// property DefaultStandardOutput s

func (v *manager) DefaultStandardOutput() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DefaultStandardOutput",
	}
}

// property DefaultStandardError s

func (v *manager) DefaultStandardError() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DefaultStandardError",
	}
}

// property RuntimeWatchdogUSec t

func (v *manager) RuntimeWatchdogUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RuntimeWatchdogUSec",
	}
}

// property ShutdownWatchdogUSec t

func (v *manager) ShutdownWatchdogUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ShutdownWatchdogUSec",
	}
}

// property ServiceWatchdogs b

func (v *manager) ServiceWatchdogs() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ServiceWatchdogs",
	}
}

// property ControlGroup s

func (v *manager) ControlGroup() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ControlGroup",
	}
}

// property SystemState s

func (v *manager) SystemState() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SystemState",
	}
}

// property ExitCode y

func (v *manager) ExitCode() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "ExitCode",
	}
}

// property DefaultTimerAccuracyUSec t

func (v *manager) DefaultTimerAccuracyUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultTimerAccuracyUSec",
	}
}

// property DefaultTimeoutStartUSec t

func (v *manager) DefaultTimeoutStartUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultTimeoutStartUSec",
	}
}

// property DefaultTimeoutStopUSec t

func (v *manager) DefaultTimeoutStopUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultTimeoutStopUSec",
	}
}

// property DefaultRestartUSec t

func (v *manager) DefaultRestartUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultRestartUSec",
	}
}

// property DefaultStartLimitIntervalUSec t

func (v *manager) DefaultStartLimitIntervalUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultStartLimitIntervalUSec",
	}
}

// property DefaultStartLimitBurst u

func (v *manager) DefaultStartLimitBurst() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "DefaultStartLimitBurst",
	}
}

// property DefaultCPUAccounting b

func (v *manager) DefaultCPUAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DefaultCPUAccounting",
	}
}

// property DefaultBlockIOAccounting b

func (v *manager) DefaultBlockIOAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DefaultBlockIOAccounting",
	}
}

// property DefaultMemoryAccounting b

func (v *manager) DefaultMemoryAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DefaultMemoryAccounting",
	}
}

// property DefaultTasksAccounting b

func (v *manager) DefaultTasksAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DefaultTasksAccounting",
	}
}

// property DefaultLimitCPU t

func (v *manager) DefaultLimitCPU() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitCPU",
	}
}

// property DefaultLimitCPUSoft t

func (v *manager) DefaultLimitCPUSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitCPUSoft",
	}
}

// property DefaultLimitFSIZE t

func (v *manager) DefaultLimitFSIZE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitFSIZE",
	}
}

// property DefaultLimitFSIZESoft t

func (v *manager) DefaultLimitFSIZESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitFSIZESoft",
	}
}

// property DefaultLimitDATA t

func (v *manager) DefaultLimitDATA() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitDATA",
	}
}

// property DefaultLimitDATASoft t

func (v *manager) DefaultLimitDATASoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitDATASoft",
	}
}

// property DefaultLimitSTACK t

func (v *manager) DefaultLimitSTACK() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitSTACK",
	}
}

// property DefaultLimitSTACKSoft t

func (v *manager) DefaultLimitSTACKSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitSTACKSoft",
	}
}

// property DefaultLimitCORE t

func (v *manager) DefaultLimitCORE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitCORE",
	}
}

// property DefaultLimitCORESoft t

func (v *manager) DefaultLimitCORESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitCORESoft",
	}
}

// property DefaultLimitRSS t

func (v *manager) DefaultLimitRSS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRSS",
	}
}

// property DefaultLimitRSSSoft t

func (v *manager) DefaultLimitRSSSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRSSSoft",
	}
}

// property DefaultLimitNOFILE t

func (v *manager) DefaultLimitNOFILE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNOFILE",
	}
}

// property DefaultLimitNOFILESoft t

func (v *manager) DefaultLimitNOFILESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNOFILESoft",
	}
}

// property DefaultLimitAS t

func (v *manager) DefaultLimitAS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitAS",
	}
}

// property DefaultLimitASSoft t

func (v *manager) DefaultLimitASSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitASSoft",
	}
}

// property DefaultLimitNPROC t

func (v *manager) DefaultLimitNPROC() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNPROC",
	}
}

// property DefaultLimitNPROCSoft t

func (v *manager) DefaultLimitNPROCSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNPROCSoft",
	}
}

// property DefaultLimitMEMLOCK t

func (v *manager) DefaultLimitMEMLOCK() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitMEMLOCK",
	}
}

// property DefaultLimitMEMLOCKSoft t

func (v *manager) DefaultLimitMEMLOCKSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitMEMLOCKSoft",
	}
}

// property DefaultLimitLOCKS t

func (v *manager) DefaultLimitLOCKS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitLOCKS",
	}
}

// property DefaultLimitLOCKSSoft t

func (v *manager) DefaultLimitLOCKSSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitLOCKSSoft",
	}
}

// property DefaultLimitSIGPENDING t

func (v *manager) DefaultLimitSIGPENDING() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitSIGPENDING",
	}
}

// property DefaultLimitSIGPENDINGSoft t

func (v *manager) DefaultLimitSIGPENDINGSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitSIGPENDINGSoft",
	}
}

// property DefaultLimitMSGQUEUE t

func (v *manager) DefaultLimitMSGQUEUE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitMSGQUEUE",
	}
}

// property DefaultLimitMSGQUEUESoft t

func (v *manager) DefaultLimitMSGQUEUESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitMSGQUEUESoft",
	}
}

// property DefaultLimitNICE t

func (v *manager) DefaultLimitNICE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNICE",
	}
}

// property DefaultLimitNICESoft t

func (v *manager) DefaultLimitNICESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitNICESoft",
	}
}

// property DefaultLimitRTPRIO t

func (v *manager) DefaultLimitRTPRIO() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRTPRIO",
	}
}

// property DefaultLimitRTPRIOSoft t

func (v *manager) DefaultLimitRTPRIOSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRTPRIOSoft",
	}
}

// property DefaultLimitRTTIME t

func (v *manager) DefaultLimitRTTIME() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRTTIME",
	}
}

// property DefaultLimitRTTIMESoft t

func (v *manager) DefaultLimitRTTIMESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultLimitRTTIMESoft",
	}
}

// property DefaultTasksMax t

func (v *manager) DefaultTasksMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DefaultTasksMax",
	}
}

// property TimerSlackNSec t

func (v *manager) TimerSlackNSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimerSlackNSec",
	}
}

type Unit struct {
	unit    // interface org.freedesktop.systemd1.Unit
	service // interface org.freedesktop.systemd1.Service
	proxy.Object
}

func NewUnit(conn *dbus.Conn, path dbus.ObjectPath) (*Unit, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Unit)
	obj.Object.Init_(conn, "org.freedesktop.systemd1", path)
	return obj, nil
}

func (obj *Unit) Unit() *unit {
	return &obj.unit
}

type unit struct{}

func (v *unit) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*unit) GetInterfaceName_() string {
	return "org.freedesktop.systemd1.Unit"
}

// method Start

func (v *unit) GoStart(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start", flags, ch, mode)
}

func (v *unit) GoStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Start", flags, ch, mode)
}

func (*unit) StoreStart(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) Start(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreStart(
		<-v.GoStart(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) StartWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStart(call)
}

// method Stop

func (v *unit) GoStop(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Stop", flags, ch, mode)
}

func (v *unit) GoStopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Stop", flags, ch, mode)
}

func (*unit) StoreStop(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) Stop(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreStop(
		<-v.GoStop(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) StopWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStopWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreStop(call)
}

// method Reload

func (v *unit) GoReload(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, mode)
}

func (v *unit) GoReloadWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Reload", flags, ch, mode)
}

func (*unit) StoreReload(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) Reload(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReload(
		<-v.GoReload(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) ReloadWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReload(call)
}

// method Restart

func (v *unit) GoRestart(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Restart", flags, ch, mode)
}

func (v *unit) GoRestartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Restart", flags, ch, mode)
}

func (*unit) StoreRestart(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) Restart(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreRestart(
		<-v.GoRestart(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) RestartWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRestartWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRestart(call)
}

// method TryRestart

func (v *unit) GoTryRestart(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TryRestart", flags, ch, mode)
}

func (v *unit) GoTryRestartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".TryRestart", flags, ch, mode)
}

func (*unit) StoreTryRestart(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) TryRestart(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreTryRestart(
		<-v.GoTryRestart(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) TryRestartWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTryRestartWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreTryRestart(call)
}

// method ReloadOrRestart

func (v *unit) GoReloadOrRestart(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadOrRestart", flags, ch, mode)
}

func (v *unit) GoReloadOrRestartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadOrRestart", flags, ch, mode)
}

func (*unit) StoreReloadOrRestart(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) ReloadOrRestart(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReloadOrRestart(
		<-v.GoReloadOrRestart(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) ReloadOrRestartWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadOrRestartWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReloadOrRestart(call)
}

// method ReloadOrTryRestart

func (v *unit) GoReloadOrTryRestart(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadOrTryRestart", flags, ch, mode)
}

func (v *unit) GoReloadOrTryRestartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReloadOrTryRestart", flags, ch, mode)
}

func (*unit) StoreReloadOrTryRestart(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *unit) ReloadOrTryRestart(flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	return v.StoreReloadOrTryRestart(
		<-v.GoReloadOrTryRestart(flags, make(chan *dbus.Call, 1), mode).Done)
}

func (v *unit) ReloadOrTryRestartWithTimeout(timeout time.Duration, flags dbus.Flags, mode string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReloadOrTryRestartWithContext(ctx, flags, make(chan *dbus.Call, 1), mode).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreReloadOrTryRestart(call)
}

// method Kill

func (v *unit) GoKill(flags dbus.Flags, ch chan *dbus.Call, who string, signal int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Kill", flags, ch, who, signal)
}

func (v *unit) GoKillWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, who string, signal int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Kill", flags, ch, who, signal)
}

func (v *unit) Kill(flags dbus.Flags, who string, signal int32) error {
	return (<-v.GoKill(flags, make(chan *dbus.Call, 1), who, signal).Done).Err
}

func (v *unit) KillWithTimeout(timeout time.Duration, flags dbus.Flags, who string, signal int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoKillWithContext(ctx, flags, make(chan *dbus.Call, 1), who, signal).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ResetFailed

func (v *unit) GoResetFailed(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetFailed", flags, ch)
}

func (v *unit) GoResetFailedWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ResetFailed", flags, ch)
}

func (v *unit) ResetFailed(flags dbus.Flags) error {
	return (<-v.GoResetFailed(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *unit) ResetFailedWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoResetFailedWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetProperties

func (v *unit) GoSetProperties(flags dbus.Flags, ch chan *dbus.Call, runtime bool, properties []Property) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProperties", flags, ch, runtime, properties)
}

func (v *unit) GoSetPropertiesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, runtime bool, properties []Property) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetProperties", flags, ch, runtime, properties)
}

func (v *unit) SetProperties(flags dbus.Flags, runtime bool, properties []Property) error {
	return (<-v.GoSetProperties(flags, make(chan *dbus.Call, 1), runtime, properties).Done).Err
}

func (v *unit) SetPropertiesWithTimeout(timeout time.Duration, flags dbus.Flags, runtime bool, properties []Property) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetPropertiesWithContext(ctx, flags, make(chan *dbus.Call, 1), runtime, properties).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Ref

func (v *unit) GoRef(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Ref", flags, ch)
}

func (v *unit) GoRefWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Ref", flags, ch)
}

func (v *unit) Ref(flags dbus.Flags) error {
	return (<-v.GoRef(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *unit) RefWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRefWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Unref

func (v *unit) GoUnref(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unref", flags, ch)
}

func (v *unit) GoUnrefWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Unref", flags, ch)
}

func (v *unit) Unref(flags dbus.Flags) error {
	return (<-v.GoUnref(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *unit) UnrefWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnrefWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Id s

func (v *unit) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Names as

func (v *unit) Names() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Names",
	}
}

// property Following s

func (v *unit) Following() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Following",
	}
}

// property Requires as

func (v *unit) Requires() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Requires",
	}
}

// property Requisite as

func (v *unit) Requisite() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Requisite",
	}
}

// property Wants as

func (v *unit) Wants() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Wants",
	}
}

// property BindsTo as

func (v *unit) BindsTo() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "BindsTo",
	}
}

// property PartOf as

func (v *unit) PartOf() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "PartOf",
	}
}

// property RequiredBy as

func (v *unit) RequiredBy() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "RequiredBy",
	}
}

// property RequisiteOf as

func (v *unit) RequisiteOf() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "RequisiteOf",
	}
}

// property WantedBy as

func (v *unit) WantedBy() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "WantedBy",
	}
}

// property BoundBy as

func (v *unit) BoundBy() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "BoundBy",
	}
}

// property ConsistsOf as

func (v *unit) ConsistsOf() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ConsistsOf",
	}
}

// property Conflicts as

func (v *unit) Conflicts() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Conflicts",
	}
}

// property ConflictedBy as

func (v *unit) ConflictedBy() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ConflictedBy",
	}
}

// property Before as

func (v *unit) Before() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Before",
	}
}

// property After as

func (v *unit) After() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "After",
	}
}

// property OnFailure as

func (v *unit) OnFailure() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "OnFailure",
	}
}

// property Triggers as

func (v *unit) Triggers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Triggers",
	}
}

// property TriggeredBy as

func (v *unit) TriggeredBy() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "TriggeredBy",
	}
}

// property PropagatesReloadTo as

func (v *unit) PropagatesReloadTo() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "PropagatesReloadTo",
	}
}

// property ReloadPropagatedFrom as

func (v *unit) ReloadPropagatedFrom() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ReloadPropagatedFrom",
	}
}

// property JoinsNamespaceOf as

func (v *unit) JoinsNamespaceOf() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "JoinsNamespaceOf",
	}
}

// property RequiresMountsFor as

func (v *unit) RequiresMountsFor() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "RequiresMountsFor",
	}
}

// property Documentation as

func (v *unit) Documentation() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Documentation",
	}
}

// property Description s

func (v *unit) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property LoadState s

func (v *unit) LoadState() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "LoadState",
	}
}

// property ActiveState s

func (v *unit) ActiveState() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ActiveState",
	}
}

// property SubState s

func (v *unit) SubState() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SubState",
	}
}

// property FragmentPath s

func (v *unit) FragmentPath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FragmentPath",
	}
}

// property SourcePath s

func (v *unit) SourcePath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SourcePath",
	}
}

// property DropInPaths as

func (v *unit) DropInPaths() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DropInPaths",
	}
}

// property UnitFileState s

func (v *unit) UnitFileState() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UnitFileState",
	}
}

// property UnitFilePreset s

func (v *unit) UnitFilePreset() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UnitFilePreset",
	}
}

// property StateChangeTimestamp t

func (v *unit) StateChangeTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StateChangeTimestamp",
	}
}

// property StateChangeTimestampMonotonic t

func (v *unit) StateChangeTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StateChangeTimestampMonotonic",
	}
}

// property InactiveExitTimestamp t

func (v *unit) InactiveExitTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InactiveExitTimestamp",
	}
}

// property InactiveExitTimestampMonotonic t

func (v *unit) InactiveExitTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InactiveExitTimestampMonotonic",
	}
}

// property ActiveEnterTimestamp t

func (v *unit) ActiveEnterTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ActiveEnterTimestamp",
	}
}

// property ActiveEnterTimestampMonotonic t

func (v *unit) ActiveEnterTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ActiveEnterTimestampMonotonic",
	}
}

// property ActiveExitTimestamp t

func (v *unit) ActiveExitTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ActiveExitTimestamp",
	}
}

// property ActiveExitTimestampMonotonic t

func (v *unit) ActiveExitTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ActiveExitTimestampMonotonic",
	}
}

// property InactiveEnterTimestamp t

func (v *unit) InactiveEnterTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InactiveEnterTimestamp",
	}
}

// property InactiveEnterTimestampMonotonic t

func (v *unit) InactiveEnterTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InactiveEnterTimestampMonotonic",
	}
}

// property CanStart b

func (v *unit) CanStart() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanStart",
	}
}

// property CanStop b

func (v *unit) CanStop() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanStop",
	}
}

// property CanReload b

func (v *unit) CanReload() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanReload",
	}
}

// property CanIsolate b

func (v *unit) CanIsolate() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanIsolate",
	}
}

// property Job (uo)

func (v *unit) Job() PropUnitJob {
	return PropUnitJob{
		Impl: v,
	}
}

type PropUnitJob struct {
	Impl proxy.Implementer
}

func (p PropUnitJob) Get(flags dbus.Flags) (value JobIdPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Job", &value)
	return
}

func (p PropUnitJob) ConnectChanged(cb func(hasValue bool, value JobIdPath)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v JobIdPath
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, JobIdPath{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Job", cb0)
}

// property StopWhenUnneeded b

func (v *unit) StopWhenUnneeded() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "StopWhenUnneeded",
	}
}

// property RefuseManualStart b

func (v *unit) RefuseManualStart() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RefuseManualStart",
	}
}

// property RefuseManualStop b

func (v *unit) RefuseManualStop() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RefuseManualStop",
	}
}

// property AllowIsolate b

func (v *unit) AllowIsolate() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AllowIsolate",
	}
}

// property DefaultDependencies b

func (v *unit) DefaultDependencies() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DefaultDependencies",
	}
}

// property OnFailureJobMode s

func (v *unit) OnFailureJobMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OnFailureJobMode",
	}
}

// property IgnoreOnIsolate b

func (v *unit) IgnoreOnIsolate() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IgnoreOnIsolate",
	}
}

// property NeedDaemonReload b

func (v *unit) NeedDaemonReload() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NeedDaemonReload",
	}
}

// property JobTimeoutUSec t

func (v *unit) JobTimeoutUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "JobTimeoutUSec",
	}
}

// property JobRunningTimeoutUSec t

func (v *unit) JobRunningTimeoutUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "JobRunningTimeoutUSec",
	}
}

// property JobTimeoutAction s

func (v *unit) JobTimeoutAction() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "JobTimeoutAction",
	}
}

// property JobTimeoutRebootArgument s

func (v *unit) JobTimeoutRebootArgument() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "JobTimeoutRebootArgument",
	}
}

// property ConditionResult b

func (v *unit) ConditionResult() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ConditionResult",
	}
}

// property AssertResult b

func (v *unit) AssertResult() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AssertResult",
	}
}

// property ConditionTimestamp t

func (v *unit) ConditionTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ConditionTimestamp",
	}
}

// property ConditionTimestampMonotonic t

func (v *unit) ConditionTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ConditionTimestampMonotonic",
	}
}

// property AssertTimestamp t

func (v *unit) AssertTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "AssertTimestamp",
	}
}

// property AssertTimestampMonotonic t

func (v *unit) AssertTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "AssertTimestampMonotonic",
	}
}

// property Conditions a(sbbsi)

func (v *unit) Conditions() PropUnitConditions {
	return PropUnitConditions{
		Impl: v,
	}
}

type PropUnitConditions struct {
	Impl proxy.Implementer
}

func (p PropUnitConditions) Get(flags dbus.Flags) (value []Condition, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Conditions", &value)
	return
}

func (p PropUnitConditions) ConnectChanged(cb func(hasValue bool, value []Condition)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []Condition
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Conditions", cb0)
}

// property Asserts a(sbbsi)

func (v *unit) Asserts() PropUnitAsserts {
	return PropUnitAsserts{
		Impl: v,
	}
}

type PropUnitAsserts struct {
	Impl proxy.Implementer
}

func (p PropUnitAsserts) Get(flags dbus.Flags) (value []Assert, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Asserts", &value)
	return
}

func (p PropUnitAsserts) ConnectChanged(cb func(hasValue bool, value []Assert)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []Assert
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Asserts", cb0)
}

// property LoadError (ss)

func (v *unit) LoadError() PropUnitLoadError {
	return PropUnitLoadError{
		Impl: v,
	}
}

type PropUnitLoadError struct {
	Impl proxy.Implementer
}

func (p PropUnitLoadError) Get(flags dbus.Flags) (value LoadError, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"LoadError", &value)
	return
}

func (p PropUnitLoadError) ConnectChanged(cb func(hasValue bool, value LoadError)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v LoadError
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, LoadError{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"LoadError", cb0)
}

// property Transient b

func (v *unit) Transient() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Transient",
	}
}

// property Perpetual b

func (v *unit) Perpetual() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Perpetual",
	}
}

// property StartLimitIntervalUSec t

func (v *unit) StartLimitIntervalUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StartLimitIntervalUSec",
	}
}

// property StartLimitBurst u

func (v *unit) StartLimitBurst() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "StartLimitBurst",
	}
}

// property StartLimitAction s

func (v *unit) StartLimitAction() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StartLimitAction",
	}
}

// property FailureAction s

func (v *unit) FailureAction() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "FailureAction",
	}
}

// property FailureActionExitStatus i

func (v *unit) FailureActionExitStatus() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "FailureActionExitStatus",
	}
}

// property SuccessAction s

func (v *unit) SuccessAction() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SuccessAction",
	}
}

// property SuccessActionExitStatus i

func (v *unit) SuccessActionExitStatus() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SuccessActionExitStatus",
	}
}

// property RebootArgument s

func (v *unit) RebootArgument() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RebootArgument",
	}
}

// property InvocationID ay

func (v *unit) InvocationID() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "InvocationID",
	}
}

// property CollectMode s

func (v *unit) CollectMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CollectMode",
	}
}

// property Refs as

func (v *unit) Refs() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Refs",
	}
}

func (obj *Unit) Service() *service {
	return &obj.service
}

type service struct{}

func (v *service) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*service) GetInterfaceName_() string {
	return "org.freedesktop.systemd1.Service"
}

// method GetProcesses

func (v *service) GoGetProcesses(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProcesses", flags, ch)
}

func (v *service) GoGetProcessesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetProcesses", flags, ch)
}

func (*service) StoreGetProcesses(call *dbus.Call) (processes []UnitProcess, err error) {
	err = call.Store(&processes)
	return
}

func (v *service) GetProcesses(flags dbus.Flags) (processes []UnitProcess, err error) {
	return v.StoreGetProcesses(
		<-v.GoGetProcesses(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *service) GetProcessesWithTimeout(timeout time.Duration, flags dbus.Flags) (processes []UnitProcess, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetProcessesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetProcesses(call)
}

// method AttachProcesses

func (v *service) GoAttachProcesses(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AttachProcesses", flags, ch, arg0, arg1)
}

func (v *service) GoAttachProcessesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AttachProcesses", flags, ch, arg0, arg1)
}

func (v *service) AttachProcesses(flags dbus.Flags, arg0 string, arg1 []uint32) error {
	return (<-v.GoAttachProcesses(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

func (v *service) AttachProcessesWithTimeout(timeout time.Duration, flags dbus.Flags, arg0 string, arg1 []uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAttachProcessesWithContext(ctx, flags, make(chan *dbus.Call, 1), arg0, arg1).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Type s

func (v *service) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Restart s

func (v *service) Restart() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Restart",
	}
}

// property PIDFile s

func (v *service) PIDFile() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PIDFile",
	}
}

// property NotifyAccess s

func (v *service) NotifyAccess() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "NotifyAccess",
	}
}

// property RestartUSec t

func (v *service) RestartUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RestartUSec",
	}
}

// property TimeoutStartUSec t

func (v *service) TimeoutStartUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeoutStartUSec",
	}
}

// property TimeoutStopUSec t

func (v *service) TimeoutStopUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeoutStopUSec",
	}
}

// property RuntimeMaxUSec t

func (v *service) RuntimeMaxUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RuntimeMaxUSec",
	}
}

// property WatchdogUSec t

func (v *service) WatchdogUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "WatchdogUSec",
	}
}

// property WatchdogTimestamp t

func (v *service) WatchdogTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "WatchdogTimestamp",
	}
}

// property WatchdogTimestampMonotonic t

func (v *service) WatchdogTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "WatchdogTimestampMonotonic",
	}
}

// property RootDirectoryStartOnly b

func (v *service) RootDirectoryStartOnly() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RootDirectoryStartOnly",
	}
}

// property RemainAfterExit b

func (v *service) RemainAfterExit() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RemainAfterExit",
	}
}

// property GuessMainPID b

func (v *service) GuessMainPID() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "GuessMainPID",
	}
}

// property RestartPreventExitStatus (aiai)

func (v *service) RestartPreventExitStatus() PropExitStatus {
	return PropExitStatus{
		Impl: v,
		Name: "RestartPreventExitStatus",
	}
}

// property RestartForceExitStatus (aiai)

func (v *service) RestartForceExitStatus() PropExitStatus {
	return PropExitStatus{
		Impl: v,
		Name: "RestartForceExitStatus",
	}
}

// property SuccessExitStatus (aiai)

func (v *service) SuccessExitStatus() PropExitStatus {
	return PropExitStatus{
		Impl: v,
		Name: "SuccessExitStatus",
	}
}

// property MainPID u

func (v *service) MainPID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "MainPID",
	}
}

// property ControlPID u

func (v *service) ControlPID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ControlPID",
	}
}

// property BusName s

func (v *service) BusName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BusName",
	}
}

// property FileDescriptorStoreMax u

func (v *service) FileDescriptorStoreMax() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "FileDescriptorStoreMax",
	}
}

// property NFileDescriptorStore u

func (v *service) NFileDescriptorStore() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NFileDescriptorStore",
	}
}

// property StatusText s

func (v *service) StatusText() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StatusText",
	}
}

// property StatusErrno i

func (v *service) StatusErrno() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "StatusErrno",
	}
}

// property Result s

func (v *service) Result() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Result",
	}
}

// property USBFunctionDescriptors s

func (v *service) USBFunctionDescriptors() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "USBFunctionDescriptors",
	}
}

// property USBFunctionStrings s

func (v *service) USBFunctionStrings() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "USBFunctionStrings",
	}
}

// property UID u

func (v *service) UID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "UID",
	}
}

// property GID u

func (v *service) GID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "GID",
	}
}

// property NRestarts u

func (v *service) NRestarts() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NRestarts",
	}
}

// property ExecMainStartTimestamp t

func (v *service) ExecMainStartTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ExecMainStartTimestamp",
	}
}

// property ExecMainStartTimestampMonotonic t

func (v *service) ExecMainStartTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ExecMainStartTimestampMonotonic",
	}
}

// property ExecMainExitTimestamp t

func (v *service) ExecMainExitTimestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ExecMainExitTimestamp",
	}
}

// property ExecMainExitTimestampMonotonic t

func (v *service) ExecMainExitTimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "ExecMainExitTimestampMonotonic",
	}
}

// property ExecMainPID u

func (v *service) ExecMainPID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ExecMainPID",
	}
}

// property ExecMainCode i

func (v *service) ExecMainCode() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "ExecMainCode",
	}
}

// property ExecMainStatus i

func (v *service) ExecMainStatus() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "ExecMainStatus",
	}
}

// property ExecStartPre a(sasbttttuii)

func (v *service) ExecStartPre() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecStartPre",
	}
}

// property ExecStart a(sasbttttuii)

func (v *service) ExecStart() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecStart",
	}
}

// property ExecStartPost a(sasbttttuii)

func (v *service) ExecStartPost() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecStartPost",
	}
}

// property ExecReload a(sasbttttuii)

func (v *service) ExecReload() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecReload",
	}
}

// property ExecStop a(sasbttttuii)

func (v *service) ExecStop() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecStop",
	}
}

// property ExecStopPost a(sasbttttuii)

func (v *service) ExecStopPost() PropExecInfos {
	return PropExecInfos{
		Impl: v,
		Name: "ExecStopPost",
	}
}

// property Slice s

func (v *service) Slice() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Slice",
	}
}

// property ControlGroup s

func (v *service) ControlGroup() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ControlGroup",
	}
}

// property MemoryCurrent t

func (v *service) MemoryCurrent() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryCurrent",
	}
}

// property CPUUsageNSec t

func (v *service) CPUUsageNSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CPUUsageNSec",
	}
}

// property TasksCurrent t

func (v *service) TasksCurrent() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TasksCurrent",
	}
}

// property IPIngressBytes t

func (v *service) IPIngressBytes() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IPIngressBytes",
	}
}

// property IPIngressPackets t

func (v *service) IPIngressPackets() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IPIngressPackets",
	}
}

// property IPEgressBytes t

func (v *service) IPEgressBytes() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IPEgressBytes",
	}
}

// property IPEgressPackets t

func (v *service) IPEgressPackets() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IPEgressPackets",
	}
}

// property Delegate b

func (v *service) Delegate() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Delegate",
	}
}

// property DelegateControllers as

func (v *service) DelegateControllers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DelegateControllers",
	}
}

// property CPUAccounting b

func (v *service) CPUAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CPUAccounting",
	}
}

// property CPUWeight t

func (v *service) CPUWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CPUWeight",
	}
}

// property StartupCPUWeight t

func (v *service) StartupCPUWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StartupCPUWeight",
	}
}

// property CPUShares t

func (v *service) CPUShares() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CPUShares",
	}
}

// property StartupCPUShares t

func (v *service) StartupCPUShares() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StartupCPUShares",
	}
}

// property CPUQuotaPerSecUSec t

func (v *service) CPUQuotaPerSecUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CPUQuotaPerSecUSec",
	}
}

// property IOAccounting b

func (v *service) IOAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IOAccounting",
	}
}

// property IOWeight t

func (v *service) IOWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IOWeight",
	}
}

// property StartupIOWeight t

func (v *service) StartupIOWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StartupIOWeight",
	}
}

// property IODeviceWeight a(st)

func (v *service) IODeviceWeight() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IODeviceWeight",
	}
}

// property IOReadBandwidthMax a(st)

func (v *service) IOReadBandwidthMax() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IOReadBandwidthMax",
	}
}

// property IOWriteBandwidthMax a(st)

func (v *service) IOWriteBandwidthMax() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IOWriteBandwidthMax",
	}
}

// property IOReadIOPSMax a(st)

func (v *service) IOReadIOPSMax() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IOReadIOPSMax",
	}
}

// property IOWriteIOPSMax a(st)

func (v *service) IOWriteIOPSMax() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IOWriteIOPSMax",
	}
}

// property IODeviceLatencyTargetUSec a(st)

func (v *service) IODeviceLatencyTargetUSec() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "IODeviceLatencyTargetUSec",
	}
}

// property BlockIOAccounting b

func (v *service) BlockIOAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "BlockIOAccounting",
	}
}

// property BlockIOWeight t

func (v *service) BlockIOWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "BlockIOWeight",
	}
}

// property StartupBlockIOWeight t

func (v *service) StartupBlockIOWeight() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "StartupBlockIOWeight",
	}
}

// property BlockIODeviceWeight a(st)

func (v *service) BlockIODeviceWeight() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "BlockIODeviceWeight",
	}
}

// property BlockIOReadBandwidth a(st)

func (v *service) BlockIOReadBandwidth() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "BlockIOReadBandwidth",
	}
}

// property BlockIOWriteBandwidth a(st)

func (v *service) BlockIOWriteBandwidth() PropIOParams {
	return PropIOParams{
		Impl: v,
		Name: "BlockIOWriteBandwidth",
	}
}

// property MemoryAccounting b

func (v *service) MemoryAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MemoryAccounting",
	}
}

// property MemoryMin t

func (v *service) MemoryMin() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryMin",
	}
}

// property MemoryLow t

func (v *service) MemoryLow() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryLow",
	}
}

// property MemoryHigh t

func (v *service) MemoryHigh() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryHigh",
	}
}

// property MemoryMax t

func (v *service) MemoryMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryMax",
	}
}

// property MemorySwapMax t

func (v *service) MemorySwapMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemorySwapMax",
	}
}

// property MemoryLimit t

func (v *service) MemoryLimit() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemoryLimit",
	}
}

// property DevicePolicy s

func (v *service) DevicePolicy() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DevicePolicy",
	}
}

// property DeviceAllow a(ss)

func (v *service) DeviceAllow() PropServiceDeviceAllow {
	return PropServiceDeviceAllow{
		Impl: v,
	}
}

type PropServiceDeviceAllow struct {
	Impl proxy.Implementer
}

func (p PropServiceDeviceAllow) Get(flags dbus.Flags) (value []DeviceAllowItem, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"DeviceAllow", &value)
	return
}

func (p PropServiceDeviceAllow) ConnectChanged(cb func(hasValue bool, value []DeviceAllowItem)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []DeviceAllowItem
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"DeviceAllow", cb0)
}

// property TasksAccounting b

func (v *service) TasksAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TasksAccounting",
	}
}

// property TasksMax t

func (v *service) TasksMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TasksMax",
	}
}

// property IPAccounting b

func (v *service) IPAccounting() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IPAccounting",
	}
}

// property IPAddressAllow a(iayu)

func (v *service) IPAddressAllow() PropServiceIPAddressAllow {
	return PropServiceIPAddressAllow{
		Impl: v,
	}
}

type PropServiceIPAddressAllow struct {
	Impl proxy.Implementer
}

func (p PropServiceIPAddressAllow) Get(flags dbus.Flags) (value []IPAddressAllowItem, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"IPAddressAllow", &value)
	return
}

func (p PropServiceIPAddressAllow) ConnectChanged(cb func(hasValue bool, value []IPAddressAllowItem)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IPAddressAllowItem
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"IPAddressAllow", cb0)
}

// property IPAddressDeny a(iayu)

func (v *service) IPAddressDeny() PropServiceIPAddressDeny {
	return PropServiceIPAddressDeny{
		Impl: v,
	}
}

type PropServiceIPAddressDeny struct {
	Impl proxy.Implementer
}

func (p PropServiceIPAddressDeny) Get(flags dbus.Flags) (value []IPAddressDenyItem, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"IPAddressDeny", &value)
	return
}

func (p PropServiceIPAddressDeny) ConnectChanged(cb func(hasValue bool, value []IPAddressDenyItem)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IPAddressDenyItem
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"IPAddressDeny", cb0)
}

// property DisableControllers as

func (v *service) DisableControllers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "DisableControllers",
	}
}

// property Environment as

func (v *service) Environment() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Environment",
	}
}

// property EnvironmentFiles a(sb)

func (v *service) EnvironmentFiles() PropEnvironmentFiles {
	return PropEnvironmentFiles{
		Impl: v,
		Name: "EnvironmentFiles",
	}
}

// property PassEnvironment as

func (v *service) PassEnvironment() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "PassEnvironment",
	}
}

// property UnsetEnvironment as

func (v *service) UnsetEnvironment() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UnsetEnvironment",
	}
}

// property UMask u

func (v *service) UMask() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "UMask",
	}
}

// property LimitCPU t

func (v *service) LimitCPU() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitCPU",
	}
}

// property LimitCPUSoft t

func (v *service) LimitCPUSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitCPUSoft",
	}
}

// property LimitFSIZE t

func (v *service) LimitFSIZE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitFSIZE",
	}
}

// property LimitFSIZESoft t

func (v *service) LimitFSIZESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitFSIZESoft",
	}
}

// property LimitDATA t

func (v *service) LimitDATA() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitDATA",
	}
}

// property LimitDATASoft t

func (v *service) LimitDATASoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitDATASoft",
	}
}

// property LimitSTACK t

func (v *service) LimitSTACK() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitSTACK",
	}
}

// property LimitSTACKSoft t

func (v *service) LimitSTACKSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitSTACKSoft",
	}
}

// property LimitCORE t

func (v *service) LimitCORE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitCORE",
	}
}

// property LimitCORESoft t

func (v *service) LimitCORESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitCORESoft",
	}
}

// property LimitRSS t

func (v *service) LimitRSS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRSS",
	}
}

// property LimitRSSSoft t

func (v *service) LimitRSSSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRSSSoft",
	}
}

// property LimitNOFILE t

func (v *service) LimitNOFILE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNOFILE",
	}
}

// property LimitNOFILESoft t

func (v *service) LimitNOFILESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNOFILESoft",
	}
}

// property LimitAS t

func (v *service) LimitAS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitAS",
	}
}

// property LimitASSoft t

func (v *service) LimitASSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitASSoft",
	}
}

// property LimitNPROC t

func (v *service) LimitNPROC() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNPROC",
	}
}

// property LimitNPROCSoft t

func (v *service) LimitNPROCSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNPROCSoft",
	}
}

// property LimitMEMLOCK t

func (v *service) LimitMEMLOCK() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitMEMLOCK",
	}
}

// property LimitMEMLOCKSoft t

func (v *service) LimitMEMLOCKSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitMEMLOCKSoft",
	}
}

// property LimitLOCKS t

func (v *service) LimitLOCKS() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitLOCKS",
	}
}

// property LimitLOCKSSoft t

func (v *service) LimitLOCKSSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitLOCKSSoft",
	}
}

// property LimitSIGPENDING t

func (v *service) LimitSIGPENDING() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitSIGPENDING",
	}
}

// property LimitSIGPENDINGSoft t

func (v *service) LimitSIGPENDINGSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitSIGPENDINGSoft",
	}
}

// property LimitMSGQUEUE t

func (v *service) LimitMSGQUEUE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitMSGQUEUE",
	}
}

// property LimitMSGQUEUESoft t

func (v *service) LimitMSGQUEUESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitMSGQUEUESoft",
	}
}

// property LimitNICE t

func (v *service) LimitNICE() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNICE",
	}
}

// property LimitNICESoft t

func (v *service) LimitNICESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitNICESoft",
	}
}

// property LimitRTPRIO t

func (v *service) LimitRTPRIO() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRTPRIO",
	}
}

// property LimitRTPRIOSoft t

func (v *service) LimitRTPRIOSoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRTPRIOSoft",
	}
}

// property LimitRTTIME t

func (v *service) LimitRTTIME() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRTTIME",
	}
}

// property LimitRTTIMESoft t

func (v *service) LimitRTTIMESoft() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LimitRTTIMESoft",
	}
}

// property WorkingDirectory s

func (v *service) WorkingDirectory() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WorkingDirectory",
	}
}

// property RootDirectory s

func (v *service) RootDirectory() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RootDirectory",
	}
}

// property RootImage s

func (v *service) RootImage() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RootImage",
	}
}

// property OOMScoreAdjust i

func (v *service) OOMScoreAdjust() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "OOMScoreAdjust",
	}
}

// property Nice i

func (v *service) Nice() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "Nice",
	}
}

// property IOSchedulingClass i

func (v *service) IOSchedulingClass() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "IOSchedulingClass",
	}
}

// property IOSchedulingPriority i

func (v *service) IOSchedulingPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "IOSchedulingPriority",
	}
}

// property CPUSchedulingPolicy i

func (v *service) CPUSchedulingPolicy() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "CPUSchedulingPolicy",
	}
}

// property CPUSchedulingPriority i

func (v *service) CPUSchedulingPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "CPUSchedulingPriority",
	}
}

// property CPUAffinity ay

func (v *service) CPUAffinity() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "CPUAffinity",
	}
}

// property TimerSlackNSec t

func (v *service) TimerSlackNSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimerSlackNSec",
	}
}

// property CPUSchedulingResetOnFork b

func (v *service) CPUSchedulingResetOnFork() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CPUSchedulingResetOnFork",
	}
}

// property NonBlocking b

func (v *service) NonBlocking() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NonBlocking",
	}
}

// property StandardInput s

func (v *service) StandardInput() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardInput",
	}
}

// property StandardInputFileDescriptorName s

func (v *service) StandardInputFileDescriptorName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardInputFileDescriptorName",
	}
}

// property StandardInputData ay

func (v *service) StandardInputData() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "StandardInputData",
	}
}

// property StandardOutput s

func (v *service) StandardOutput() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardOutput",
	}
}

// property StandardOutputFileDescriptorName s

func (v *service) StandardOutputFileDescriptorName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardOutputFileDescriptorName",
	}
}

// property StandardError s

func (v *service) StandardError() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardError",
	}
}

// property StandardErrorFileDescriptorName s

func (v *service) StandardErrorFileDescriptorName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "StandardErrorFileDescriptorName",
	}
}

// property TTYPath s

func (v *service) TTYPath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "TTYPath",
	}
}

// property TTYReset b

func (v *service) TTYReset() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TTYReset",
	}
}

// property TTYVHangup b

func (v *service) TTYVHangup() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TTYVHangup",
	}
}

// property TTYVTDisallocate b

func (v *service) TTYVTDisallocate() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TTYVTDisallocate",
	}
}

// property SyslogPriority i

func (v *service) SyslogPriority() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SyslogPriority",
	}
}

// property SyslogIdentifier s

func (v *service) SyslogIdentifier() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SyslogIdentifier",
	}
}

// property SyslogLevelPrefix b

func (v *service) SyslogLevelPrefix() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SyslogLevelPrefix",
	}
}

// property SyslogLevel i

func (v *service) SyslogLevel() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SyslogLevel",
	}
}

// property SyslogFacility i

func (v *service) SyslogFacility() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SyslogFacility",
	}
}

// property LogLevelMax i

func (v *service) LogLevelMax() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "LogLevelMax",
	}
}

// property LogRateLimitIntervalUSec t

func (v *service) LogRateLimitIntervalUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "LogRateLimitIntervalUSec",
	}
}

// property LogRateLimitBurst u

func (v *service) LogRateLimitBurst() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "LogRateLimitBurst",
	}
}

// property LogExtraFields aay

func (v *service) LogExtraFields() PropServiceLogExtraFields {
	return PropServiceLogExtraFields{
		Impl: v,
	}
}

type PropServiceLogExtraFields struct {
	Impl proxy.Implementer
}

func (p PropServiceLogExtraFields) Get(flags dbus.Flags) (value [][]byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"LogExtraFields", &value)
	return
}

func (p PropServiceLogExtraFields) ConnectChanged(cb func(hasValue bool, value [][]byte)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]byte
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"LogExtraFields", cb0)
}

// property SecureBits i

func (v *service) SecureBits() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SecureBits",
	}
}

// property CapabilityBoundingSet t

func (v *service) CapabilityBoundingSet() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CapabilityBoundingSet",
	}
}

// property AmbientCapabilities t

func (v *service) AmbientCapabilities() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "AmbientCapabilities",
	}
}

// property User s

func (v *service) User() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "User",
	}
}

// property Group s

func (v *service) Group() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Group",
	}
}

// property DynamicUser b

func (v *service) DynamicUser() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DynamicUser",
	}
}

// property RemoveIPC b

func (v *service) RemoveIPC() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RemoveIPC",
	}
}

// property SupplementaryGroups as

func (v *service) SupplementaryGroups() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "SupplementaryGroups",
	}
}

// property PAMName s

func (v *service) PAMName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PAMName",
	}
}

// property ReadWritePaths as

func (v *service) ReadWritePaths() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ReadWritePaths",
	}
}

// property ReadOnlyPaths as

func (v *service) ReadOnlyPaths() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ReadOnlyPaths",
	}
}

// property InaccessiblePaths as

func (v *service) InaccessiblePaths() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "InaccessiblePaths",
	}
}

// property MountFlags t

func (v *service) MountFlags() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MountFlags",
	}
}

// property PrivateTmp b

func (v *service) PrivateTmp() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PrivateTmp",
	}
}

// property PrivateDevices b

func (v *service) PrivateDevices() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PrivateDevices",
	}
}

// property ProtectKernelTunables b

func (v *service) ProtectKernelTunables() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ProtectKernelTunables",
	}
}

// property ProtectKernelModules b

func (v *service) ProtectKernelModules() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ProtectKernelModules",
	}
}

// property ProtectControlGroups b

func (v *service) ProtectControlGroups() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ProtectControlGroups",
	}
}

// property PrivateNetwork b

func (v *service) PrivateNetwork() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PrivateNetwork",
	}
}

// property PrivateUsers b

func (v *service) PrivateUsers() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PrivateUsers",
	}
}

// property PrivateMounts b

func (v *service) PrivateMounts() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PrivateMounts",
	}
}

// property ProtectHome s

func (v *service) ProtectHome() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ProtectHome",
	}
}

// property ProtectSystem s

func (v *service) ProtectSystem() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ProtectSystem",
	}
}

// property SameProcessGroup b

func (v *service) SameProcessGroup() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SameProcessGroup",
	}
}

// property UtmpIdentifier s

func (v *service) UtmpIdentifier() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UtmpIdentifier",
	}
}

// property UtmpMode s

func (v *service) UtmpMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UtmpMode",
	}
}

// property SELinuxContext (bs)

func (v *service) SELinuxContext() PropBS {
	return PropBS{
		Impl: v,
		Name: "SELinuxContext",
	}
}

// property AppArmorProfile (bs)

func (v *service) AppArmorProfile() PropBS {
	return PropBS{
		Impl: v,
		Name: "AppArmorProfile",
	}
}

// property SmackProcessLabel (bs)

func (v *service) SmackProcessLabel() PropBS {
	return PropBS{
		Impl: v,
		Name: "SmackProcessLabel",
	}
}

// property IgnoreSIGPIPE b

func (v *service) IgnoreSIGPIPE() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IgnoreSIGPIPE",
	}
}

// property NoNewPrivileges b

func (v *service) NoNewPrivileges() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NoNewPrivileges",
	}
}

// property SystemCallFilter (bas)

func (v *service) SystemCallFilter() PropServiceSystemCallFilter {
	return PropServiceSystemCallFilter{
		Impl: v,
	}
}

type PropServiceSystemCallFilter struct {
	Impl proxy.Implementer
}

func (p PropServiceSystemCallFilter) Get(flags dbus.Flags) (value SystemCallFilter, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"SystemCallFilter", &value)
	return
}

func (p PropServiceSystemCallFilter) ConnectChanged(cb func(hasValue bool, value SystemCallFilter)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v SystemCallFilter
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, SystemCallFilter{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"SystemCallFilter", cb0)
}

// property SystemCallArchitectures as

func (v *service) SystemCallArchitectures() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "SystemCallArchitectures",
	}
}

// property SystemCallErrorNumber i

func (v *service) SystemCallErrorNumber() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SystemCallErrorNumber",
	}
}

// property Personality s

func (v *service) Personality() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Personality",
	}
}

// property LockPersonality b

func (v *service) LockPersonality() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "LockPersonality",
	}
}

// property RestrictAddressFamilies (bas)

func (v *service) RestrictAddressFamilies() PropServiceRestrictAddressFamilies {
	return PropServiceRestrictAddressFamilies{
		Impl: v,
	}
}

type PropServiceRestrictAddressFamilies struct {
	Impl proxy.Implementer
}

func (p PropServiceRestrictAddressFamilies) Get(flags dbus.Flags) (value RestrictAddressFamilies, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"RestrictAddressFamilies", &value)
	return
}

func (p PropServiceRestrictAddressFamilies) ConnectChanged(cb func(hasValue bool, value RestrictAddressFamilies)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v RestrictAddressFamilies
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, RestrictAddressFamilies{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"RestrictAddressFamilies", cb0)
}

// property RuntimeDirectoryPreserve s

func (v *service) RuntimeDirectoryPreserve() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RuntimeDirectoryPreserve",
	}
}

// property RuntimeDirectoryMode u

func (v *service) RuntimeDirectoryMode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RuntimeDirectoryMode",
	}
}

// property RuntimeDirectory as

func (v *service) RuntimeDirectory() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "RuntimeDirectory",
	}
}

// property StateDirectoryMode u

func (v *service) StateDirectoryMode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "StateDirectoryMode",
	}
}

// property StateDirectory as

func (v *service) StateDirectory() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "StateDirectory",
	}
}

// property CacheDirectoryMode u

func (v *service) CacheDirectoryMode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "CacheDirectoryMode",
	}
}

// property CacheDirectory as

func (v *service) CacheDirectory() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "CacheDirectory",
	}
}

// property LogsDirectoryMode u

func (v *service) LogsDirectoryMode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "LogsDirectoryMode",
	}
}

// property LogsDirectory as

func (v *service) LogsDirectory() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "LogsDirectory",
	}
}

// property ConfigurationDirectoryMode u

func (v *service) ConfigurationDirectoryMode() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "ConfigurationDirectoryMode",
	}
}

// property ConfigurationDirectory as

func (v *service) ConfigurationDirectory() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "ConfigurationDirectory",
	}
}

// property MemoryDenyWriteExecute b

func (v *service) MemoryDenyWriteExecute() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MemoryDenyWriteExecute",
	}
}

// property RestrictRealtime b

func (v *service) RestrictRealtime() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RestrictRealtime",
	}
}

// property RestrictNamespaces t

func (v *service) RestrictNamespaces() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RestrictNamespaces",
	}
}

// property BindPaths a(ssbt)

func (v *service) BindPaths() PropBindPaths {
	return PropBindPaths{
		Impl: v,
		Name: "BindPaths",
	}
}

// property BindReadOnlyPaths a(ssbt)

func (v *service) BindReadOnlyPaths() PropBindPaths {
	return PropBindPaths{
		Impl: v,
		Name: "BindReadOnlyPaths",
	}
}

// property TemporaryFileSystem a(ss)

func (v *service) TemporaryFileSystem() PropServiceTemporaryFileSystem {
	return PropServiceTemporaryFileSystem{
		Impl: v,
	}
}

type PropServiceTemporaryFileSystem struct {
	Impl proxy.Implementer
}

func (p PropServiceTemporaryFileSystem) Get(flags dbus.Flags) (value []TemporaryFileSystemItem, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"TemporaryFileSystem", &value)
	return
}

func (p PropServiceTemporaryFileSystem) ConnectChanged(cb func(hasValue bool, value []TemporaryFileSystemItem)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []TemporaryFileSystemItem
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"TemporaryFileSystem", cb0)
}

// property MountAPIVFS b

func (v *service) MountAPIVFS() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MountAPIVFS",
	}
}

// property KeyringMode s

func (v *service) KeyringMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KeyringMode",
	}
}

// property KillMode s

func (v *service) KillMode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "KillMode",
	}
}

// property KillSignal i

func (v *service) KillSignal() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "KillSignal",
	}
}

// property FinalKillSignal i

func (v *service) FinalKillSignal() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "FinalKillSignal",
	}
}

// property SendSIGKILL b

func (v *service) SendSIGKILL() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SendSIGKILL",
	}
}

// property SendSIGHUP b

func (v *service) SendSIGHUP() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SendSIGHUP",
	}
}

// property WatchdogSignal i

func (v *service) WatchdogSignal() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "WatchdogSignal",
	}
}

type PropExitStatus struct {
	Impl proxy.Implementer
	Name string
}

func (p PropExitStatus) Get(flags dbus.Flags) (value ExitStatus, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropExitStatus) Set(flags dbus.Flags, value ExitStatus) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropExitStatus) ConnectChanged(cb func(hasValue bool, value ExitStatus)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ExitStatus
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ExitStatus{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropExecInfos struct {
	Impl proxy.Implementer
	Name string
}

func (p PropExecInfos) Get(flags dbus.Flags) (value []ExecInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropExecInfos) Set(flags dbus.Flags, value []ExecInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropExecInfos) ConnectChanged(cb func(hasValue bool, value []ExecInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ExecInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropIOParams struct {
	Impl proxy.Implementer
	Name string
}

func (p PropIOParams) Get(flags dbus.Flags) (value []IOParam, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropIOParams) Set(flags dbus.Flags, value []IOParam) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropIOParams) ConnectChanged(cb func(hasValue bool, value []IOParam)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []IOParam
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropEnvironmentFiles struct {
	Impl proxy.Implementer
	Name string
}

func (p PropEnvironmentFiles) Get(flags dbus.Flags) (value []EnvironmentFile, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropEnvironmentFiles) Set(flags dbus.Flags, value []EnvironmentFile) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropEnvironmentFiles) ConnectChanged(cb func(hasValue bool, value []EnvironmentFile)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []EnvironmentFile
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropBS struct {
	Impl proxy.Implementer
	Name string
}

func (p PropBS) Get(flags dbus.Flags) (value BS, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropBS) Set(flags dbus.Flags, value BS) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropBS) ConnectChanged(cb func(hasValue bool, value BS)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v BS
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, BS{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropBindPaths struct {
	Impl proxy.Implementer
	Name string
}

func (p PropBindPaths) Get(flags dbus.Flags) (value []BindPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropBindPaths) Set(flags dbus.Flags, value []BindPath) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropBindPaths) ConnectChanged(cb func(hasValue bool, value []BindPath)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []BindPath
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
