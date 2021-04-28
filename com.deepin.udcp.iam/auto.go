package iam

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

type UdcpCache struct {
	udcpCache // interface com.deepin.udcp.iam
	proxy.Object
}

func NewUdcpCache(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (*UdcpCache, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(UdcpCache)
	obj.Object.Init_(conn, serviceName, path)
	return obj, nil
}

type udcpCache struct{}

func (v *udcpCache) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*udcpCache) GetInterfaceName_() string {
	return "com.deepin.udcp.iam"
}

// method GetUserIdList

func (v *udcpCache) GoGetUserIdList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserIdList", flags, ch)
}

func (*udcpCache) StoreGetUserIdList(call *dbus.Call) (uIdList []uint32, err error) {
	err = call.Store(&uIdList)
	return
}

func (v *udcpCache) GetUserIdList(flags dbus.Flags) (uIdList []uint32, err error) {
	return v.StoreGetUserIdList(
		<-v.GoGetUserIdList(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetUserGroups

func (v *udcpCache) GoGetUserGroups(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserGroups", flags, ch, name)
}

func (*udcpCache) StoreGetUserGroups(call *dbus.Call) (groups []string, err error) {
	err = call.Store(&groups)
	return
}

func (v *udcpCache) GetUserGroups(flags dbus.Flags, name string) (groups []string, err error) {
	return v.StoreGetUserGroups(
		<-v.GoGetUserGroups(flags, make(chan *dbus.Call, 1), name).Done)
}
