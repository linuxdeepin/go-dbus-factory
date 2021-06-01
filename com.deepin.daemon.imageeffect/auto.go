package imageeffect

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

type ImageEffect struct {
	imageEffect // interface com.deepin.daemon.ImageEffect
	proxy.Object
}

func NewImageEffect(conn *dbus.Conn) *ImageEffect {
	obj := new(ImageEffect)
	obj.Object.Init_(conn, "com.deepin.daemon.ImageEffect", "/com/deepin/daemon/ImageEffect")
	return obj
}

type imageEffect struct{}

func (v *imageEffect) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*imageEffect) GetInterfaceName_() string {
	return "com.deepin.daemon.ImageEffect"
}

// method Delete

func (v *imageEffect) GoDelete(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, effect, filename)
}

func (v *imageEffect) GoDeleteWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Delete", flags, ch, effect, filename)
}

func (v *imageEffect) Delete(flags dbus.Flags, effect string, filename string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), effect, filename).Done).Err
}

func (v *imageEffect) DeleteWithTimeout(timeout time.Duration, flags dbus.Flags, effect string, filename string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteWithContext(ctx, flags, make(chan *dbus.Call, 1), effect, filename).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Get

func (v *imageEffect) GoGet(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch, effect, filename)
}

func (v *imageEffect) GoGetWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Get", flags, ch, effect, filename)
}

func (*imageEffect) StoreGet(call *dbus.Call) (outputFile string, err error) {
	err = call.Store(&outputFile)
	return
}

func (v *imageEffect) Get(flags dbus.Flags, effect string, filename string) (outputFile string, err error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1), effect, filename).Done)
}

func (v *imageEffect) GetWithTimeout(timeout time.Duration, flags dbus.Flags, effect string, filename string) (outputFile string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetWithContext(ctx, flags, make(chan *dbus.Call, 1), effect, filename).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGet(call)
}
