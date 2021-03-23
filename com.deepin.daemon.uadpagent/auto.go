package uadpagent

import (
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type UadpAgent struct {
	uadpagent // interface com.deepin.daemon.UadpAgent
	proxy.Object
}

func NewUadpAgent(conn *dbus.Conn) *UadpAgent {
	obj := new(UadpAgent)
	obj.Object.Init_(conn, "com.deepin.daemon.UadpAgent", "/com/deepin/daemon/UadpAgent")
	return obj
}

type uadpagent struct{}

func (v *uadpagent) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*uadpagent) GetInterfaceName_() string {
	return "com.deepin.daemon.UadpAgent"
}

// method SetDataKey

func (v *uadpagent) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string, dataKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDataKey", flags, ch, keyName, dataKey)
}

func (v *uadpagent) SetDataKey(flags dbus.Flags, keyName string, dataKey string) error {
	return (<-v.GoSetDataKey(flags, make(chan *dbus.Call, 1), keyName, dataKey).Done).Err
}

// method GetDataKey

func (v *uadpagent) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDataKey", flags, ch, keyName)
}

func (*uadpagent) StoreGetDataKey(call *dbus.Call) (dataKey string, err error) {
	err = call.Store(&dataKey)
	return
}

func (v *uadpagent) GetDataKey(flags dbus.Flags, keyName string) (dataKey string, err error) {
	return v.StoreGetDataKey(
		<-v.GoGetDataKey(flags, make(chan *dbus.Call, 1), keyName).Done)
}
