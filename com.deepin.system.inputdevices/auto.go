package inputdevices

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type InputDevices struct {
	inputDevices // interface com.deepin.system.InputDevices
	proxy.Object
}

func NewInputDevices(conn *dbus.Conn) *InputDevices {
	obj := new(InputDevices)
	obj.Object.Init_(conn, "com.deepin.system.InputDevices", "/com/deepin/system/InputDevices")
	return obj
}

type inputDevices struct{}

func (v *inputDevices) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*inputDevices) GetInterfaceName_() string {
	return "com.deepin.system.InputDevices"
}

// signal TouchscreenAdded

func (v *inputDevices) ConnectTouchscreenAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchscreenAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchscreenAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchscreenRemoved

func (v *inputDevices) ConnectTouchscreenRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchscreenRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchscreenRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Touchscreens ao

func (v *inputDevices) Touchscreens() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Touchscreens",
	}
}

type Touchscreen struct {
	touchscreen // interface com.deepin.system.InputDevices.Touchscreen
	proxy.Object
}

func NewTouchscreen(conn *dbus.Conn, path dbus.ObjectPath) (*Touchscreen, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Touchscreen)
	obj.Object.Init_(conn, "com.deepin.system.InputDevices", path)
	return obj, nil
}

type touchscreen struct{}

func (v *touchscreen) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*touchscreen) GetInterfaceName_() string {
	return "com.deepin.system.InputDevices.Touchscreen"
}

// property DevNode s

func (v *touchscreen) DevNode() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DevNode",
	}
}

// property BusType s

func (v *touchscreen) BusType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BusType",
	}
}

// property UUID s

func (v *touchscreen) UUID() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UUID",
	}
}

// property Phys s

func (v *touchscreen) Phys() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Phys",
	}
}

// property OutputName s

func (v *touchscreen) OutputName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "OutputName",
	}
}

// property Width d

func (v *touchscreen) Width() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Width",
	}
}

// property Height d

func (v *touchscreen) Height() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Height",
	}
}

// property Name s

func (v *touchscreen) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Serial s

func (v *touchscreen) Serial() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Serial",
	}
}
