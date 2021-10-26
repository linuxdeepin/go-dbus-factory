// Code generated by "./generator ./com.deepin.system.displaycfg/"; DO NOT EDIT.

package displaycfg

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type DisplayCfg interface {
	displayCfg // interface com.deepin.system.DisplayCfg
	proxy.Object
}

type objectDisplayCfg struct {
	interfaceDisplayCfg // interface com.deepin.system.DisplayCfg
	proxy.ImplObject
}

func NewDisplayCfg(conn *dbus.Conn) DisplayCfg {
	obj := new(objectDisplayCfg)
	obj.ImplObject.Init_(conn, "com.deepin.system.DisplayCfg", "/com/deepin/system/DisplayCfg")
	return obj
}

type displayCfg interface {
	GoGet(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Get(flags dbus.Flags) (string, error)
	GoSet(flags dbus.Flags, ch chan *dbus.Call, cfgStr string) *dbus.Call
	Set(flags dbus.Flags, cfgStr string) error
	ConnectUpdated(cb func(updateAt string)) (dbusutil.SignalHandlerId, error)
}

type interfaceDisplayCfg struct{}

func (v *interfaceDisplayCfg) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDisplayCfg) GetInterfaceName_() string {
	return "com.deepin.system.DisplayCfg"
}

// method Get

func (v *interfaceDisplayCfg) GoGet(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch)
}

func (*interfaceDisplayCfg) StoreGet(call *dbus.Call) (cfgStr string, err error) {
	err = call.Store(&cfgStr)
	return
}

func (v *interfaceDisplayCfg) Get(flags dbus.Flags) (string, error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1)).Done)
}

// method Set

func (v *interfaceDisplayCfg) GoSet(flags dbus.Flags, ch chan *dbus.Call, cfgStr string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Set", flags, ch, cfgStr)
}

func (v *interfaceDisplayCfg) Set(flags dbus.Flags, cfgStr string) error {
	return (<-v.GoSet(flags, make(chan *dbus.Call, 1), cfgStr).Done).Err
}

// signal Updated

func (v *interfaceDisplayCfg) ConnectUpdated(cb func(updateAt string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Updated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Updated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var updateAt string
		err := dbus.Store(sig.Body, &updateAt)
		if err == nil {
			cb(updateAt)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
