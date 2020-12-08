package powermanager

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

type PowerManager struct {
	powerManager // interface com.deepin.daemon.PowerManager
	proxy.Object
}

func NewPowerManager(conn *dbus.Conn) *PowerManager {
	obj := new(PowerManager)
	obj.Object.Init_(conn, "com.deepin.daemon.PowerManager", "/com/deepin/daemon/PowerManager")
	return obj
}

type powerManager struct{}

func (v *powerManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*powerManager) GetInterfaceName_() string {
	return "com.deepin.daemon.PowerManager"
}

// method CanShutdown

func (v *powerManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (*powerManager) StoreCanShutdown(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *powerManager) CanShutdown(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanShutdown(
		<-v.GoCanShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanReboot

func (v *powerManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*powerManager) StoreCanReboot(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *powerManager) CanReboot(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSuspend

func (v *powerManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*powerManager) StoreCanSuspend(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *powerManager) CanSuspend(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanHibernate

func (v *powerManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*powerManager) StoreCanHibernate(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *powerManager) CanHibernate(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}
