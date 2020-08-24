package cursorhelper

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

type CursorHelper struct {
	cursorHelper // interface com.deepin.api.CursorHelper
	proxy.Object
}

func NewCursorHelper(conn *dbus.Conn) *CursorHelper {
	obj := new(CursorHelper)
	obj.Object.Init_(conn, "com.deepin.api.CursorHelper", "/com/deepin/api/CursorHelper")
	return obj
}

type cursorHelper struct{}

func (v *cursorHelper) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*cursorHelper) GetInterfaceName_() string {
	return "com.deepin.api.CursorHelper"
}

// method Set

func (v *cursorHelper) GoSet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Set", flags, ch, name)
}

func (v *cursorHelper) Set(flags dbus.Flags, name string) error {
	return (<-v.GoSet(flags, make(chan *dbus.Call, 1), name).Done).Err
}
