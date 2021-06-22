package iam

import "context"
import "errors"
import "fmt"
import dbus "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
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

func (v *udcpCache) GoGetUserIdListWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUserIdList", flags, ch)
}

func (*udcpCache) StoreGetUserIdList(call *dbus.Call) (uIdList []uint32, err error) {
	err = call.Store(&uIdList)
	return
}

func (v *udcpCache) GetUserIdList(flags dbus.Flags) (uIdList []uint32, err error) {
	return v.StoreGetUserIdList(
		<-v.GoGetUserIdList(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *udcpCache) GetUserIdListWithTimeout(timeout time.Duration, flags dbus.Flags) (uIdList []uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUserIdListWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUserIdList(call)
}

// method GetUserGroups

func (v *udcpCache) GoGetUserGroups(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserGroups", flags, ch, name)
}

func (v *udcpCache) GoGetUserGroupsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetUserGroups", flags, ch, name)
}

func (*udcpCache) StoreGetUserGroups(call *dbus.Call) (groups []string, err error) {
	err = call.Store(&groups)
	return
}

func (v *udcpCache) GetUserGroups(flags dbus.Flags, name string) (groups []string, err error) {
	return v.StoreGetUserGroups(
		<-v.GoGetUserGroups(flags, make(chan *dbus.Call, 1), name).Done)
}

func (v *udcpCache) GetUserGroupsWithTimeout(timeout time.Duration, flags dbus.Flags, name string) (groups []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetUserGroupsWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetUserGroups(call)
}

// method RemoveCacheFile

func (v *udcpCache) GoRemoveCacheFile(flags dbus.Flags, ch chan *dbus.Call, uId uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveCacheFile", flags, ch, uId)
}

func (v *udcpCache) GoRemoveCacheFileWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, uId uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveCacheFile", flags, ch, uId)
}

func (*udcpCache) StoreRemoveCacheFile(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *udcpCache) RemoveCacheFile(flags dbus.Flags, uId uint32) (result bool, err error) {
	return v.StoreRemoveCacheFile(
		<-v.GoRemoveCacheFile(flags, make(chan *dbus.Call, 1), uId).Done)
}

func (v *udcpCache) RemoveCacheFileWithTimeout(timeout time.Duration, flags dbus.Flags, uId uint32) (result bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveCacheFileWithContext(ctx, flags, make(chan *dbus.Call, 1), uId).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRemoveCacheFile(call)
}

// property Enable b

func (v *udcpCache) Enable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Enable",
	}
}
