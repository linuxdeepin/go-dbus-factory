package localehelper

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

type LocaleHelper struct {
	localeHelper // interface com.deepin.api.LocaleHelper
	proxy.Object
}

func NewLocaleHelper(conn *dbus.Conn) *LocaleHelper {
	obj := new(LocaleHelper)
	obj.Object.Init_(conn, "com.deepin.api.LocaleHelper", "/com/deepin/api/LocaleHelper")
	return obj
}

type localeHelper struct{}

func (v *localeHelper) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*localeHelper) GetInterfaceName_() string {
	return "com.deepin.api.LocaleHelper"
}

// method GenerateLocale

func (v *localeHelper) GoGenerateLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GenerateLocale", flags, ch, locale)
}

func (v *localeHelper) GenerateLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoGenerateLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// method SetLocale

func (v *localeHelper) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *localeHelper) SetLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoSetLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// signal Success

func (v *localeHelper) ConnectSuccess(cb func(ok bool, reason string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Success", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Success",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ok bool
		var reason string
		err := dbus.Store(sig.Body, &ok, &reason)
		if err == nil {
			cb(ok, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
