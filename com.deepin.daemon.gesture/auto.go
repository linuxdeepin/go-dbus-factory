package gesture

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

type Gesture struct {
	gesture // interface com.deepin.daemon.Gesture
	proxy.Object
}

func NewGesture(conn *dbus.Conn) *Gesture {
	obj := new(Gesture)
	obj.Object.Init_(conn, "com.deepin.daemon.Gesture", "/com/deepin/daemon/Gesture")
	return obj
}

type gesture struct{}

func (v *gesture) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*gesture) GetInterfaceName_() string {
	return "com.deepin.daemon.Gesture"
}

// signal Event

func (v *gesture) ConnectEvent(cb func(name string, direction string, fingers int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Event", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Event",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var name string
		var direction string
		var fingers int32
		err := dbus.Store(sig.Body, &name, &direction, &fingers)
		if err == nil {
			cb(name, direction, fingers)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
