package kwin

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

type Compositor struct {
	compositing // interface org.kde.kwin.Compositing
	proxy.Object
}

func NewCompositor(conn *dbus.Conn) *Compositor {
	obj := new(Compositor)
	obj.Object.Init_(conn, "org.kde.KWin", "/Compositor")
	return obj
}

type compositing struct{}

func (v *compositing) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*compositing) GetInterfaceName_() string {
	return "org.kde.kwin.Compositing"
}

// method suspend

func (v *compositing) GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".suspend", flags, ch)
}

func (v *compositing) Suspend(flags dbus.Flags) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method resume

func (v *compositing) GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".resume", flags, ch)
}

func (v *compositing) Resume(flags dbus.Flags) error {
	return (<-v.GoResume(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal compositingToggled

func (v *compositing) ConnectCompositingToggled(cb func(active bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "compositingToggled", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".compositingToggled",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var active bool
		err := dbus.Store(sig.Body, &active)
		if err == nil {
			cb(active)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property active b

func (v *compositing) Active() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "active",
	}
}

// property compositingPossible b

func (v *compositing) CompositingPossible() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "compositingPossible",
	}
}

// property compositingNotPossibleReason s

func (v *compositing) CompositingNotPossibleReason() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "compositingNotPossibleReason",
	}
}

// property openGLIsBroken b

func (v *compositing) OpenGLIsBroken() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "openGLIsBroken",
	}
}

// property compositingType s

func (v *compositing) CompositingType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "compositingType",
	}
}

// property supportedOpenGLPlatformInterfaces as

func (v *compositing) SupportedOpenGLPlatformInterfaces() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "supportedOpenGLPlatformInterfaces",
	}
}

// property platformRequiresCompositing b

func (v *compositing) PlatformRequiresCompositing() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "platformRequiresCompositing",
	}
}
