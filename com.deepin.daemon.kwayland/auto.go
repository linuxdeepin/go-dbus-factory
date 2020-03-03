package kwayland

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

type OutputManagement struct {
	outputManagement // interface com.deepin.daemon.KWayland.Output
	proxy.Object
}

func NewOutputManagement(conn *dbus.Conn) *OutputManagement {
	obj := new(OutputManagement)
	obj.Object.Init_(conn, "com.deepin.daemon.KWayland", "/com/deepin/daemon/KWayland/Output")
	return obj
}

type outputManagement struct{}

func (v *outputManagement) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*outputManagement) GetInterfaceName_() string {
	return "com.deepin.daemon.KWayland.Output"
}

// method ListOutput

func (v *outputManagement) GoListOutput(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutput", flags, ch)
}

func (*outputManagement) StoreListOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) ListOutput(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreListOutput(
		<-v.GoListOutput(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetOutput

func (v *outputManagement) GoGetOutput(flags dbus.Flags, ch chan *dbus.Call, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetOutput", flags, ch, arg1)
}

func (*outputManagement) StoreGetOutput(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *outputManagement) GetOutput(flags dbus.Flags, arg1 string) (arg0 string, err error) {
	return v.StoreGetOutput(
		<-v.GoGetOutput(flags, make(chan *dbus.Call, 1), arg1).Done)
}

// method Apply

func (v *outputManagement) GoApply(flags dbus.Flags, ch chan *dbus.Call, outputs string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Apply", flags, ch, outputs)
}

func (v *outputManagement) Apply(flags dbus.Flags, outputs string) error {
	return (<-v.GoApply(flags, make(chan *dbus.Call, 1), outputs).Done).Err
}

// signal OutputAdded

func (v *outputManagement) ConnectOutputAdded(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal OutputRemoved

func (v *outputManagement) ConnectOutputRemoved(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal OutputChanged

func (v *outputManagement) ConnectOutputChanged(cb func(output string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "OutputChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".OutputChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var output string
		err := dbus.Store(sig.Body, &output)
		if err == nil {
			cb(output)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
