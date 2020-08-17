package object_manager

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type ObjectManager struct{}

func (v *ObjectManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ObjectManager) GetInterfaceName_() string {
	return "org.freedesktop.DBus.ObjectManager"
}

// method GetManagedObjects

func (v *ObjectManager) GoGetManagedObjects(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetManagedObjects", flags, ch)
}

func (*ObjectManager) StoreGetManagedObjects(call *dbus.Call) (object_paths_interfaces_and_properties map[dbus.ObjectPath]map[string]map[string]dbus.Variant, err error) {
	err = call.Store(&object_paths_interfaces_and_properties)
	return
}

func (v *ObjectManager) GetManagedObjects(flags dbus.Flags) (object_paths_interfaces_and_properties map[dbus.ObjectPath]map[string]map[string]dbus.Variant, err error) {
	return v.StoreGetManagedObjects(
		<-v.GoGetManagedObjects(flags, make(chan *dbus.Call, 1)).Done)
}

// signal InterfacesAdded

func (v *ObjectManager) ConnectInterfacesAdded(cb func(object_path dbus.ObjectPath, interfaces_and_properties map[string]map[string]dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InterfacesAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InterfacesAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		var interfaces_and_properties map[string]map[string]dbus.Variant
		err := dbus.Store(sig.Body, &object_path, &interfaces_and_properties)
		if err == nil {
			cb(object_path, interfaces_and_properties)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal InterfacesRemoved

func (v *ObjectManager) ConnectInterfacesRemoved(cb func(object_path dbus.ObjectPath, interfaces []string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "InterfacesRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".InterfacesRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		var interfaces []string
		err := dbus.Store(sig.Body, &object_path, &interfaces)
		if err == nil {
			cb(object_path, interfaces)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
