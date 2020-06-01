package lockfront

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

type LockFront struct {
	lockfront // interface com.deepin.dde.lockFront
	proxy.Object
}

func NewLockFront(conn *dbus.Conn) *LockFront {
	obj := new(LockFront)
	obj.Object.Init_(conn, "com.deepin.dde.lockFront", "/com/deepin/dde/lockFront")
	return obj
}

type lockfront struct{}

func (v *lockfront) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*lockfront) GetInterfaceName_() string {
	return "com.deepin.dde.lockFront"
}

// signal ChangKey

func (v *lockfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
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
