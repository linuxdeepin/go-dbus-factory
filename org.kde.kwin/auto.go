package kwin

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

func (v *compositing) GoSuspendWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".suspend", flags, ch)
}

func (v *compositing) Suspend(flags dbus.Flags) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *compositing) SuspendWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSuspendWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method resume

func (v *compositing) GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".resume", flags, ch)
}

func (v *compositing) GoResumeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".resume", flags, ch)
}

func (v *compositing) Resume(flags dbus.Flags) error {
	return (<-v.GoResume(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *compositing) ResumeWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoResumeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
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

type InputDeviceManager struct {
	inputDeviceManager // interface org.kde.KWin.InputDeviceManager
	proxy.Object
}

func NewInputDeviceManager(conn *dbus.Conn) *InputDeviceManager {
	obj := new(InputDeviceManager)
	obj.Object.Init_(conn, "org.kde.KWin", "/org/kde/KWin/InputDevice")
	return obj
}

type inputDeviceManager struct{}

func (v *inputDeviceManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*inputDeviceManager) GetInterfaceName_() string {
	return "org.kde.KWin.InputDeviceManager"
}

// signal deviceAdded

func (v *inputDeviceManager) ConnectDeviceAdded(cb func(sysName string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "deviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".deviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sysName string
		err := dbus.Store(sig.Body, &sysName)
		if err == nil {
			cb(sysName)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal deviceRemoved

func (v *inputDeviceManager) ConnectDeviceRemoved(cb func(sysName string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "deviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".deviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sysName string
		err := dbus.Store(sig.Body, &sysName)
		if err == nil {
			cb(sysName)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property devicesSysNames as

func (v *inputDeviceManager) DevicesSysNames() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "devicesSysNames",
	}
}

type InputDevice struct {
	inputDevice // interface org.kde.KWin.InputDevice
	proxy.Object
}

func NewInputDevice(conn *dbus.Conn, path dbus.ObjectPath) (*InputDevice, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(InputDevice)
	obj.Object.Init_(conn, "org.kde.KWin", path)
	return obj, nil
}

type inputDevice struct{}

func (v *inputDevice) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*inputDevice) GetInterfaceName_() string {
	return "org.kde.KWin.InputDevice"
}

// property keyboard b

func (v *inputDevice) Keyboard() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "keyboard",
	}
}

// property alphaNumericKeyboard b

func (v *inputDevice) AlphaNumericKeyboard() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "alphaNumericKeyboard",
	}
}

// property pointer b

func (v *inputDevice) Pointer() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "pointer",
	}
}

// property touchpad b

func (v *inputDevice) Touchpad() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "touchpad",
	}
}

// property touch b

func (v *inputDevice) Touch() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "touch",
	}
}

// property tabletTool b

func (v *inputDevice) TabletTool() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tabletTool",
	}
}

// property tabletPad b

func (v *inputDevice) TabletPad() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tabletPad",
	}
}

// property gestureSupport b

func (v *inputDevice) GestureSupport() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "gestureSupport",
	}
}

// property name s

func (v *inputDevice) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "name",
	}
}

// property sysName s

func (v *inputDevice) SysName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "sysName",
	}
}

// property outputName s

func (v *inputDevice) OutputName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "outputName",
	}
}

// property product u

func (v *inputDevice) Product() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "product",
	}
}

// property vendor u

func (v *inputDevice) Vendor() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "vendor",
	}
}

// property supportsDisableEvents b

func (v *inputDevice) SupportsDisableEvents() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsDisableEvents",
	}
}

// property enabled b

func (v *inputDevice) Enabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "enabled",
	}
}

// property supportedButtons i

func (v *inputDevice) SupportedButtons() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "supportedButtons",
	}
}

// property supportsCalibrationMatrix b

func (v *inputDevice) SupportsCalibrationMatrix() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsCalibrationMatrix",
	}
}

// property supportsLeftHanded b

func (v *inputDevice) SupportsLeftHanded() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsLeftHanded",
	}
}

// property leftHandedEnabledByDefault b

func (v *inputDevice) LeftHandedEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "leftHandedEnabledByDefault",
	}
}

// property leftHanded b

func (v *inputDevice) LeftHanded() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "leftHanded",
	}
}

// property supportsDisableEventsOnExternalMouse b

func (v *inputDevice) SupportsDisableEventsOnExternalMouse() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsDisableEventsOnExternalMouse",
	}
}

// property supportsDisableWhileTyping b

func (v *inputDevice) SupportsDisableWhileTyping() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsDisableWhileTyping",
	}
}

// property disableWhileTypingEnabledByDefault b

func (v *inputDevice) DisableWhileTypingEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "disableWhileTypingEnabledByDefault",
	}
}

// property disableWhileTyping b

func (v *inputDevice) DisableWhileTyping() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "disableWhileTyping",
	}
}

// property supportsPointerAcceleration b

func (v *inputDevice) SupportsPointerAcceleration() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsPointerAcceleration",
	}
}

// property defaultPointerAcceleration d

func (v *inputDevice) DefaultPointerAcceleration() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "defaultPointerAcceleration",
	}
}

// property pointerAcceleration d

func (v *inputDevice) PointerAcceleration() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "pointerAcceleration",
	}
}

// property supportsPointerAccelerationProfileFlat b

func (v *inputDevice) SupportsPointerAccelerationProfileFlat() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsPointerAccelerationProfileFlat",
	}
}

// property defaultPointerAccelerationProfileFlat b

func (v *inputDevice) DefaultPointerAccelerationProfileFlat() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "defaultPointerAccelerationProfileFlat",
	}
}

// property pointerAccelerationProfileFlat b

func (v *inputDevice) PointerAccelerationProfileFlat() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "pointerAccelerationProfileFlat",
	}
}

// property supportsPointerAccelerationProfileAdaptive b

func (v *inputDevice) SupportsPointerAccelerationProfileAdaptive() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsPointerAccelerationProfileAdaptive",
	}
}

// property defaultPointerAccelerationProfileAdaptive b

func (v *inputDevice) DefaultPointerAccelerationProfileAdaptive() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "defaultPointerAccelerationProfileAdaptive",
	}
}

// property pointerAccelerationProfileAdaptive b

func (v *inputDevice) PointerAccelerationProfileAdaptive() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "pointerAccelerationProfileAdaptive",
	}
}

// property tapFingerCount i

func (v *inputDevice) TapFingerCount() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "tapFingerCount",
	}
}

// property tapToClickEnabledByDefault b

func (v *inputDevice) TapToClickEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapToClickEnabledByDefault",
	}
}

// property tapToClick b

func (v *inputDevice) TapToClick() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapToClick",
	}
}

// property supportsLmrTapButtonMap b

func (v *inputDevice) SupportsLmrTapButtonMap() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsLmrTapButtonMap",
	}
}

// property lmrTapButtonMapEnabledByDefault b

func (v *inputDevice) LmrTapButtonMapEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "lmrTapButtonMapEnabledByDefault",
	}
}

// property lmrTapButtonMap b

func (v *inputDevice) LmrTapButtonMap() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "lmrTapButtonMap",
	}
}

// property tapAndDragEnabledByDefault b

func (v *inputDevice) TapAndDragEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapAndDragEnabledByDefault",
	}
}

// property tapAndDrag b

func (v *inputDevice) TapAndDrag() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapAndDrag",
	}
}

// property tapDragLockEnabledByDefault b

func (v *inputDevice) TapDragLockEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapDragLockEnabledByDefault",
	}
}

// property tapDragLock b

func (v *inputDevice) TapDragLock() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tapDragLock",
	}
}

// property supportsMiddleEmulation b

func (v *inputDevice) SupportsMiddleEmulation() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsMiddleEmulation",
	}
}

// property middleEmulationEnabledByDefault b

func (v *inputDevice) MiddleEmulationEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "middleEmulationEnabledByDefault",
	}
}

// property middleEmulation b

func (v *inputDevice) MiddleEmulation() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "middleEmulation",
	}
}

// property supportsNaturalScroll b

func (v *inputDevice) SupportsNaturalScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsNaturalScroll",
	}
}

// property naturalScrollEnabledByDefault b

func (v *inputDevice) NaturalScrollEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "naturalScrollEnabledByDefault",
	}
}

// property naturalScroll b

func (v *inputDevice) NaturalScroll() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "naturalScroll",
	}
}

// property supportsScrollTwoFinger b

func (v *inputDevice) SupportsScrollTwoFinger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsScrollTwoFinger",
	}
}

// property scrollTwoFingerEnabledByDefault b

func (v *inputDevice) ScrollTwoFingerEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollTwoFingerEnabledByDefault",
	}
}

// property scrollTwoFinger b

func (v *inputDevice) ScrollTwoFinger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollTwoFinger",
	}
}

// property supportsScrollEdge b

func (v *inputDevice) SupportsScrollEdge() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsScrollEdge",
	}
}

// property scrollEdgeEnabledByDefault b

func (v *inputDevice) ScrollEdgeEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollEdgeEnabledByDefault",
	}
}

// property scrollEdge b

func (v *inputDevice) ScrollEdge() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollEdge",
	}
}

// property supportsScrollOnButtonDown b

func (v *inputDevice) SupportsScrollOnButtonDown() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsScrollOnButtonDown",
	}
}

// property scrollOnButtonDownEnabledByDefault b

func (v *inputDevice) ScrollOnButtonDownEnabledByDefault() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollOnButtonDownEnabledByDefault",
	}
}

// property defaultScrollButton u

func (v *inputDevice) DefaultScrollButton() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "defaultScrollButton",
	}
}

// property scrollOnButtonDown b

func (v *inputDevice) ScrollOnButtonDown() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "scrollOnButtonDown",
	}
}

// property scrollButton u

func (v *inputDevice) ScrollButton() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "scrollButton",
	}
}

// property switchDevice b

func (v *inputDevice) SwitchDevice() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "switchDevice",
	}
}

// property lidSwitch b

func (v *inputDevice) LidSwitch() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "lidSwitch",
	}
}

// property tabletModeSwitch b

func (v *inputDevice) TabletModeSwitch() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "tabletModeSwitch",
	}
}

// property supportsClickMethodAreas b

func (v *inputDevice) SupportsClickMethodAreas() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsClickMethodAreas",
	}
}

// property defaultClickMethodAreas b

func (v *inputDevice) DefaultClickMethodAreas() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "defaultClickMethodAreas",
	}
}

// property clickMethodAreas b

func (v *inputDevice) ClickMethodAreas() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "clickMethodAreas",
	}
}

// property supportsClickMethodClickfinger b

func (v *inputDevice) SupportsClickMethodClickfinger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "supportsClickMethodClickfinger",
	}
}

// property defaultClickMethodClickfinger b

func (v *inputDevice) DefaultClickMethodClickfinger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "defaultClickMethodClickfinger",
	}
}

// property clickMethodClickfinger b

func (v *inputDevice) ClickMethodClickfinger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "clickMethodClickfinger",
	}
}
