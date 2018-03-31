package inputdevices

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

type Keyboard struct {
	keyboard // interface com.deepin.daemon.InputDevice.Keyboard
	proxy.Object
}

func NewKeyboard(conn *dbus.Conn) *Keyboard {
	obj := new(Keyboard)
	obj.Object.Init_(conn, "com.deepin.daemon.InputDevices", "/com/deepin/daemon/InputDevice/Keyboard")
	return obj
}

type keyboard struct{}

func (v *keyboard) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*keyboard) GetInterfaceName_() string {
	return "com.deepin.daemon.InputDevice.Keyboard"
}

// method AddLayoutOption

func (v *keyboard) GoAddLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddLayoutOption", flags, ch, option)
}

func (v *keyboard) AddLayoutOption(flags dbus.Flags, option string) error {
	return (<-v.GoAddLayoutOption(flags, make(chan *dbus.Call, 1), option).Done).Err
}

// method AddUserLayout

func (v *keyboard) GoAddUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddUserLayout", flags, ch, layout)
}

func (v *keyboard) AddUserLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoAddUserLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method ClearLayoutOption

func (v *keyboard) GoClearLayoutOption(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearLayoutOption", flags, ch)
}

func (v *keyboard) ClearLayoutOption(flags dbus.Flags) error {
	return (<-v.GoClearLayoutOption(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DeleteLayoutOption

func (v *keyboard) GoDeleteLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteLayoutOption", flags, ch, option)
}

func (v *keyboard) DeleteLayoutOption(flags dbus.Flags, option string) error {
	return (<-v.GoDeleteLayoutOption(flags, make(chan *dbus.Call, 1), option).Done).Err
}

// method DeleteUserLayout

func (v *keyboard) GoDeleteUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteUserLayout", flags, ch, layout)
}

func (v *keyboard) DeleteUserLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoDeleteUserLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method GetLayoutDesc

func (v *keyboard) GoGetLayoutDesc(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLayoutDesc", flags, ch, layout)
}

func (*keyboard) StoreGetLayoutDesc(call *dbus.Call) (description string, err error) {
	err = call.Store(&description)
	return
}

func (v *keyboard) GetLayoutDesc(flags dbus.Flags, layout string) (description string, err error) {
	return v.StoreGetLayoutDesc(
		<-v.GoGetLayoutDesc(flags, make(chan *dbus.Call, 1), layout).Done)
}

// method LayoutList

func (v *keyboard) GoLayoutList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LayoutList", flags, ch)
}

func (*keyboard) StoreLayoutList(call *dbus.Call) (layout_list map[string]string, err error) {
	err = call.Store(&layout_list)
	return
}

func (v *keyboard) LayoutList(flags dbus.Flags) (layout_list map[string]string, err error) {
	return v.StoreLayoutList(
		<-v.GoLayoutList(flags, make(chan *dbus.Call, 1)).Done)
}

// method Reset

func (v *keyboard) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *keyboard) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property UserOptionList as

func (v *keyboard) UserOptionList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UserOptionList",
	}
}

// property RepeatEnabled b

func (v *keyboard) RepeatEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RepeatEnabled",
	}
}

// property CapslockToggle b

func (v *keyboard) CapslockToggle() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CapslockToggle",
	}
}

// property CursorBlink i

func (v *keyboard) CursorBlink() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "CursorBlink",
	}
}

// property RepeatInterval u

func (v *keyboard) RepeatInterval() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RepeatInterval",
	}
}

// property RepeatDelay u

func (v *keyboard) RepeatDelay() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "RepeatDelay",
	}
}

// property CurrentLayout s

func (v *keyboard) CurrentLayout() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CurrentLayout",
	}
}

// property UserLayoutList as

func (v *keyboard) UserLayoutList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UserLayoutList",
	}
}

type TouchPad struct {
	touchPad // interface com.deepin.daemon.InputDevice.TouchPad
	proxy.Object
}

func NewTouchPad(conn *dbus.Conn) *TouchPad {
	obj := new(TouchPad)
	obj.Object.Init_(conn, "com.deepin.daemon.InputDevices", "/com/deepin/daemon/InputDevice/TouchPad")
	return obj
}

type touchPad struct{}

func (v *touchPad) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*touchPad) GetInterfaceName_() string {
	return "com.deepin.daemon.InputDevice.TouchPad"
}

// method Reset

func (v *touchPad) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *touchPad) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property EdgeScroll b

func (v *touchPad) EdgeScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "EdgeScroll",
	}
}

// property PalmDetect b

func (v *touchPad) PalmDetect() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PalmDetect",
	}
}

// property MotionAcceleration d

func (v *touchPad) MotionAcceleration() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "MotionAcceleration",
	}
}

// property DeltaScroll i

func (v *touchPad) DeltaScroll() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DeltaScroll",
	}
}

// property DragThreshold i

func (v *touchPad) DragThreshold() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DragThreshold",
	}
}

// property LeftHanded b

func (v *touchPad) LeftHanded() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "LeftHanded",
	}
}

// property DisableIfTyping b

func (v *touchPad) DisableIfTyping() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "DisableIfTyping",
	}
}

// property NaturalScroll b

func (v *touchPad) NaturalScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "NaturalScroll",
	}
}

// property HorizScroll b

func (v *touchPad) HorizScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HorizScroll",
	}
}

// property VertScroll b

func (v *touchPad) VertScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "VertScroll",
	}
}

// property MotionThreshold d

func (v *touchPad) MotionThreshold() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "MotionThreshold",
	}
}

// property DoubleClick i

func (v *touchPad) DoubleClick() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "DoubleClick",
	}
}

// property DeviceList s

func (v *touchPad) DeviceList() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DeviceList",
	}
}

// property TPadEnable b

func (v *touchPad) TPadEnable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TPadEnable",
	}
}

// property PalmMinZ i

func (v *touchPad) PalmMinZ() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "PalmMinZ",
	}
}

// property Exist b

func (v *touchPad) Exist() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Exist",
	}
}

// property TapClick b

func (v *touchPad) TapClick() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "TapClick",
	}
}

// property MotionScaling d

func (v *touchPad) MotionScaling() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "MotionScaling",
	}
}

// property PalmMinWidth i

func (v *touchPad) PalmMinWidth() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "PalmMinWidth",
	}
}
