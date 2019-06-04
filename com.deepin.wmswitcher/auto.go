package wmswitcher

import "errors"
import "fmt"
import "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type WMSwitcher struct {
	wmSwitcher // interface com.deepin.WMSwitcher
	proxy.Object
}

func NewWMSwitcher(conn *dbus.Conn) *WMSwitcher {
	obj := new(WMSwitcher)
	obj.Object.Init_(conn, "com.deepin.WMSwitcher", "/com/deepin/WMSwitcher")
	return obj
}

type wmSwitcher struct{}

func (v *wmSwitcher) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*wmSwitcher) GetInterfaceName_() string {
	return "com.deepin.WMSwitcher"
}

// method AllowSwitch

func (v *wmSwitcher) GoAllowSwitch(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowSwitch", flags, ch)
}

func (*wmSwitcher) StoreAllowSwitch(call *dbus.Call) (allow bool, err error) {
	err = call.Store(&allow)
	return
}

func (v *wmSwitcher) AllowSwitch(flags dbus.Flags) (allow bool, err error) {
	return v.StoreAllowSwitch(
		<-v.GoAllowSwitch(flags, make(chan *dbus.Call, 1)).Done)
}

// method CurrentWM

func (v *wmSwitcher) GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CurrentWM", flags, ch)
}

func (*wmSwitcher) StoreCurrentWM(call *dbus.Call) (wmName string, err error) {
	err = call.Store(&wmName)
	return
}

func (v *wmSwitcher) CurrentWM(flags dbus.Flags) (wmName string, err error) {
	return v.StoreCurrentWM(
		<-v.GoCurrentWM(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestSwitchWM

func (v *wmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *wmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	return (<-v.GoRequestSwitchWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal WMChanged

func (v *wmSwitcher) ConnectWMChanged(cb func(wmName string)) (dbusutil.SignalHandlerId, error) {
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
