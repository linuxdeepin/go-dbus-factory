package osd

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

type OSD struct {
	osd // interface com.deepin.dde.osd
	proxy.Object
}

func NewOSD(conn *dbus.Conn) *OSD {
	obj := new(OSD)
	obj.Object.Init_(conn, "com.deepin.dde.osd", "/")
	return obj
}

type osd struct{}

func (v *osd) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*osd) GetInterfaceName_() string {
	return "com.deepin.dde.osd"
}

// method ShowOSD

func (v *osd) GoShowOSD(flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowOSD", flags, ch, osd)
}

func (v *osd) GoShowOSDWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ShowOSD", flags, ch, osd)
}

func (v *osd) ShowOSD(flags dbus.Flags, osd string) error {
	return (<-v.GoShowOSD(flags, make(chan *dbus.Call, 1), osd).Done).Err
}

func (v *osd) ShowOSDWithTimeout(timeout time.Duration, flags dbus.Flags, osd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoShowOSDWithContext(ctx, flags, make(chan *dbus.Call, 1), osd).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
