package pinyin

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Pinyin struct {
	pinyin // interface com.deepin.api.Pinyin
	proxy.Object
}

func NewPinyin(conn *dbus.Conn) *Pinyin {
	obj := new(Pinyin)
	obj.Object.Init_(conn, "com.deepin.api.Pinyin", "/com/deepin/api/Pinyin")
	return obj
}

type pinyin struct{}

func (v *pinyin) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*pinyin) GetInterfaceName_() string {
	return "com.deepin.api.Pinyin"
}

// method Query

func (v *pinyin) GoQuery(flags dbus.Flags, ch chan *dbus.Call, hans string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Query", flags, ch, hans)
}

func (v *pinyin) GoQueryWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, hans string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Query", flags, ch, hans)
}

func (*pinyin) StoreQuery(call *dbus.Call) (pinyin []string, err error) {
	err = call.Store(&pinyin)
	return
}

func (v *pinyin) Query(flags dbus.Flags, hans string) (pinyin []string, err error) {
	return v.StoreQuery(
		<-v.GoQuery(flags, make(chan *dbus.Call, 1), hans).Done)
}

func (v *pinyin) QueryWithTimeout(timeout time.Duration, flags dbus.Flags, hans string) (pinyin []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoQueryWithContext(ctx, flags, make(chan *dbus.Call, 1), hans).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreQuery(call)
}

// method QueryList

func (v *pinyin) GoQueryList(flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".QueryList", flags, ch, hansList)
}

func (v *pinyin) GoQueryListWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".QueryList", flags, ch, hansList)
}

func (*pinyin) StoreQueryList(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *pinyin) QueryList(flags dbus.Flags, hansList []string) (json string, err error) {
	return v.StoreQueryList(
		<-v.GoQueryList(flags, make(chan *dbus.Call, 1), hansList).Done)
}

func (v *pinyin) QueryListWithTimeout(timeout time.Duration, flags dbus.Flags, hansList []string) (json string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoQueryListWithContext(ctx, flags, make(chan *dbus.Call, 1), hansList).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreQueryList(call)
}
