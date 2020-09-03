package shutdownfront

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

type ShutdownFront struct {
	shutdownfront // interface com.deepin.dde.shutdownFront
	proxy.Object
}

func NewShutdownFront(conn *dbus.Conn) *ShutdownFront {
	obj := new(ShutdownFront)
	obj.Object.Init_(conn, "com.deepin.dde.shutdownFront", "/com/deepin/dde/shutdownFront")
	return obj
}

type shutdownfront struct{}

func (v *shutdownfront) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*shutdownfront) GetInterfaceName_() string {
	return "com.deepin.dde.shutdownFront"
}

// signal ChangKey

func (v *shutdownfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChangKey", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChangKey",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keyEvent string
		err := dbus.Store(sig.Body, &keyEvent)
		if err == nil {
			cb(keyEvent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
