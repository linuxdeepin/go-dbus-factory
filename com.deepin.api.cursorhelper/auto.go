package cursorhelper

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

func (v *cursorHelper) GoSetWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Set", flags, ch, name)
}

func (v *cursorHelper) Set(flags dbus.Flags, name string) error {
	return (<-v.GoSet(flags, make(chan *dbus.Call, 1), name).Done).Err
}

func (v *cursorHelper) SetWithTimeout(timeout time.Duration, flags dbus.Flags, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetWithContext(ctx, flags, make(chan *dbus.Call, 1), name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
