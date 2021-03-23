package shutdownfront

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type ShutdownFront struct {
	shutdownfront // interface com.deepin.dde.shutdownFront
	proxy.Object
}

func NewShutdownFront(conn *dbus.Conn) *ShutdownFront {
	obj := new(ShutdownFront)
	obj.Object.Init_(conn, "com.deepin.dde.shutdownFront", "/com/deepin/dde/shutdownFront")
	return obj
}

type shutdownfront struct{}

func (v *shutdownfront) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*shutdownfront) GetInterfaceName_() string {
	return "com.deepin.dde.shutdownFront"
}

// method Hibernate

func (v *shutdownfront) GoHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hibernate", flags, ch)
}

func (v *shutdownfront) Hibernate(flags dbus.Flags) error {
	return (<-v.GoHibernate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Lock

func (v *shutdownfront) GoLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Lock", flags, ch)
}

func (v *shutdownfront) Lock(flags dbus.Flags) error {
	return (<-v.GoLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Logout

func (v *shutdownfront) GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *shutdownfront) Logout(flags dbus.Flags) error {
	return (<-v.GoLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Restart

func (v *shutdownfront) GoRestart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Restart", flags, ch)
}

func (v *shutdownfront) Restart(flags dbus.Flags) error {
	return (<-v.GoRestart(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *shutdownfront) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *shutdownfront) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Shutdown

func (v *shutdownfront) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *shutdownfront) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Suspend

func (v *shutdownfront) GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Suspend", flags, ch)
}

func (v *shutdownfront) Suspend(flags dbus.Flags) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SwitchUser

func (v *shutdownfront) GoSwitchUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchUser", flags, ch)
}

func (v *shutdownfront) SwitchUser(flags dbus.Flags) error {
	return (<-v.GoSwitchUser(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal ChangKey

func (v *shutdownfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
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
