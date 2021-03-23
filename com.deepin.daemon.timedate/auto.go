package timedate

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Timedate struct {
	timedate // interface com.deepin.daemon.Timedate
	proxy.Object
}

func NewTimedate(conn *dbus.Conn) *Timedate {
	obj := new(Timedate)
	obj.Object.Init_(conn, "com.deepin.daemon.Timedate", "/com/deepin/daemon/Timedate")
	return obj
}

type timedate struct{}

func (v *timedate) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*timedate) GetInterfaceName_() string {
	return "com.deepin.daemon.Timedate"
}

// signal TimeUpdate

func (v *timedate) ConnectTimeUpdate(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TimeUpdate", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TimeUpdate",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
