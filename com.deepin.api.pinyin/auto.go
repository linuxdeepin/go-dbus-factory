package pinyin

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

func (*pinyin) StoreQuery(call *dbus.Call) (pinyin []string, err error) {
	err = call.Store(&pinyin)
	return
}

func (v *pinyin) Query(flags dbus.Flags, hans string) (pinyin []string, err error) {
	return v.StoreQuery(
		<-v.GoQuery(flags, make(chan *dbus.Call, 1), hans).Done)
}

// method QueryList

func (v *pinyin) GoQueryList(flags dbus.Flags, ch chan *dbus.Call, hansList []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".QueryList", flags, ch, hansList)
}

func (*pinyin) StoreQueryList(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *pinyin) QueryList(flags dbus.Flags, hansList []string) (json string, err error) {
	return v.StoreQueryList(
		<-v.GoQueryList(flags, make(chan *dbus.Call, 1), hansList).Done)
}
