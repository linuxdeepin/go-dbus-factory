package lockfront

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

type LockFront struct {
	lockfront // interface com.deepin.dde.lockFront
	proxy.Object
}

func NewLockFront(conn *dbus.Conn) *LockFront {
	obj := new(LockFront)
	obj.Object.Init_(conn, "com.deepin.dde.lockFront", "/com/deepin/dde/lockFront")
	return obj
}

type lockfront struct{}

func (v *lockfront) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*lockfront) GetInterfaceName_() string {
	return "com.deepin.dde.lockFront"
}

// method Show

func (v *lockfront) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *lockfront) GoShowWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *lockfront) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *lockfront) ShowWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowUserList

func (v *lockfront) GoShowUserList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowUserList", flags, ch)
}

func (v *lockfront) GoShowUserListWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowUserList", flags, ch)
}

func (v *lockfront) ShowUserList(flags dbus.Flags) error {
	return (<-v.GoShowUserList(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *lockfront) ShowUserListWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowUserListWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ShowAuth

func (v *lockfront) GoShowAuth(flags dbus.Flags, ch chan *dbus.Call, active bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowAuth", flags, ch, active)
}

func (v *lockfront) GoShowAuthWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, active bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowAuth", flags, ch, active)
}

func (v *lockfront) ShowAuth(flags dbus.Flags, active bool) error {
	return (<-v.GoShowAuth(flags, make(chan *dbus.Call, 1), active).Done).Err
}

func (v *lockfront) ShowAuthWithTimeout(timeout time.Duration, flags dbus.Flags, active bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowAuthWithContext(ctx, flags, make(chan *dbus.Call, 1), active).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Suspend

func (v *lockfront) GoSuspend(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Suspend", flags, ch, enable)
}

func (v *lockfront) GoSuspendWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Suspend", flags, ch, enable)
}

func (v *lockfront) Suspend(flags dbus.Flags, enable bool) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

func (v *lockfront) SuspendWithTimeout(timeout time.Duration, flags dbus.Flags, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSuspendWithContext(ctx, flags, make(chan *dbus.Call, 1), enable).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Hibernate

func (v *lockfront) GoHibernate(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hibernate", flags, ch, enable)
}

func (v *lockfront) GoHibernateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Hibernate", flags, ch, enable)
}

func (v *lockfront) Hibernate(flags dbus.Flags, enable bool) error {
	return (<-v.GoHibernate(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

func (v *lockfront) HibernateWithTimeout(timeout time.Duration, flags dbus.Flags, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoHibernateWithContext(ctx, flags, make(chan *dbus.Call, 1), enable).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal ChangKey

func (v *lockfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChangKey", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChangKey",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keyEvent string
		err := dbus.Store(sig.Body, &keyEvent)
		if err == nil {
			cb(keyEvent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
