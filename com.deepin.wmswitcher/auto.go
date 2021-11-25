// Code generated by "./generator ./com.deepin.wmswitcher"; DO NOT EDIT.

package wmswitcher

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type WMSwitcher interface {
	wmSwitcher // interface com.deepin.WMSwitcher
	proxy.Object
}

type objectWMSwitcher struct {
	interfaceWmSwitcher // interface com.deepin.WMSwitcher
	proxy.ImplObject
}

func NewWMSwitcher(conn *dbus.Conn) WMSwitcher {
	obj := new(objectWMSwitcher)
	obj.ImplObject.Init_(conn, "com.deepin.WMSwitcher", "/com/deepin/WMSwitcher")
	return obj
}

type wmSwitcher interface {
	GoAllowSwitch(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	AllowSwitch(flags dbus.Flags) (bool, error)
	GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CurrentWM(flags dbus.Flags) (string, error)
	GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestSwitchWM(flags dbus.Flags) error
	ConnectWMChanged(cb func(wmName string)) (dbusutil.SignalHandlerId, error)
}

type interfaceWmSwitcher struct{}

func (v *interfaceWmSwitcher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceWmSwitcher) GetInterfaceName_() string {
	return "com.deepin.WMSwitcher"
}

// method AllowSwitch

func (v *interfaceWmSwitcher) GoAllowSwitch(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowSwitch", flags, ch)
}

func (*interfaceWmSwitcher) StoreAllowSwitch(call *dbus.Call) (allow bool, err error) {
	err = call.Store(&allow)
	return
}

func (v *interfaceWmSwitcher) AllowSwitch(flags dbus.Flags) (bool, error) {
	return v.StoreAllowSwitch(
		<-v.GoAllowSwitch(flags, make(chan *dbus.Call, 1)).Done)
}

// method CurrentWM

func (v *interfaceWmSwitcher) GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CurrentWM", flags, ch)
}

func (*interfaceWmSwitcher) StoreCurrentWM(call *dbus.Call) (wmName string, err error) {
	err = call.Store(&wmName)
	return
}

func (v *interfaceWmSwitcher) CurrentWM(flags dbus.Flags) (string, error) {
	return v.StoreCurrentWM(
		<-v.GoCurrentWM(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestSwitchWM

func (v *interfaceWmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *interfaceWmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	return (<-v.GoRequestSwitchWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal WMChanged

func (v *interfaceWmSwitcher) ConnectWMChanged(cb func(wmName string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WMChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WMChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var wmName string
		err := dbus.Store(sig.Body, &wmName)
		if err == nil {
			cb(wmName)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
