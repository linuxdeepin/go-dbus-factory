// Code generated by "./generator ./session/org.kde.kwin"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package kwin

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Compositor interface {
	compositing // interface org.kde.kwin.Compositing
	proxy.Object
}

type objectCompositor struct {
	interfaceCompositing // interface org.kde.kwin.Compositing
	proxy.ImplObject
}

func NewCompositor(conn *dbus.Conn) Compositor {
	obj := new(objectCompositor)
	obj.ImplObject.Init_(conn, "org.kde.KWin", "/Compositor")
	return obj
}

type compositing interface {
	GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Suspend(flags dbus.Flags) error
	GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Resume(flags dbus.Flags) error
	ConnectCompositingToggled(cb func(active bool)) (dbusutil.SignalHandlerId, error)
	Active() proxy.PropBool
	CompositingPossible() proxy.PropBool
	CompositingNotPossibleReason() proxy.PropString
	OpenGLIsBroken() proxy.PropBool
	CompositingType() proxy.PropString
	SupportedOpenGLPlatformInterfaces() proxy.PropStringArray
	PlatformRequiresCompositing() proxy.PropBool
}

type interfaceCompositing struct{}

func (v *interfaceCompositing) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceCompositing) GetInterfaceName_() string {
	return "org.kde.kwin.Compositing"
}

// method suspend

func (v *interfaceCompositing) GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".suspend", flags, ch)
}

func (v *interfaceCompositing) Suspend(flags dbus.Flags) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method resume

func (v *interfaceCompositing) GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".resume", flags, ch)
}

func (v *interfaceCompositing) Resume(flags dbus.Flags) error {
	return (<-v.GoResume(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal compositingToggled

func (v *interfaceCompositing) ConnectCompositingToggled(cb func(active bool)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceCompositing) Active() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "active",
	}
}

// property compositingPossible b

func (v *interfaceCompositing) CompositingPossible() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "compositingPossible",
	}
}

// property compositingNotPossibleReason s

func (v *interfaceCompositing) CompositingNotPossibleReason() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "compositingNotPossibleReason",
	}
}

// property openGLIsBroken b

func (v *interfaceCompositing) OpenGLIsBroken() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "openGLIsBroken",
	}
}

// property compositingType s

func (v *interfaceCompositing) CompositingType() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "compositingType",
	}
}

// property supportedOpenGLPlatformInterfaces as

func (v *interfaceCompositing) SupportedOpenGLPlatformInterfaces() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "supportedOpenGLPlatformInterfaces",
	}
}

// property platformRequiresCompositing b

func (v *interfaceCompositing) PlatformRequiresCompositing() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "platformRequiresCompositing",
	}
}

type InputDeviceManager interface {
	inputDeviceManager // interface org.kde.KWin.InputDeviceManager
	proxy.Object
}

type objectInputDeviceManager struct {
	interfaceInputDeviceManager // interface org.kde.KWin.InputDeviceManager
	proxy.ImplObject
}

func NewInputDeviceManager(conn *dbus.Conn) InputDeviceManager {
	obj := new(objectInputDeviceManager)
	obj.ImplObject.Init_(conn, "org.kde.KWin", "/org/kde/KWin/InputDevice")
	return obj
}

type inputDeviceManager interface {
	ConnectDeviceAdded(cb func(sysName string)) (dbusutil.SignalHandlerId, error)
	ConnectDeviceRemoved(cb func(sysName string)) (dbusutil.SignalHandlerId, error)
	DevicesSysNames() proxy.PropStringArray
}

type interfaceInputDeviceManager struct{}

func (v *interfaceInputDeviceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceInputDeviceManager) GetInterfaceName_() string {
	return "org.kde.KWin.InputDeviceManager"
}

// signal deviceAdded

func (v *interfaceInputDeviceManager) ConnectDeviceAdded(cb func(sysName string)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceInputDeviceManager) ConnectDeviceRemoved(cb func(sysName string)) (dbusutil.SignalHandlerId, error) {
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

func (v *interfaceInputDeviceManager) DevicesSysNames() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "devicesSysNames",
	}
}

type InputDevice interface {
	inputDevice // interface org.kde.KWin.InputDevice
	proxy.Object
}

type objectInputDevice struct {
	interfaceInputDevice // interface org.kde.KWin.InputDevice
	proxy.ImplObject
}

func NewInputDevice(conn *dbus.Conn, path dbus.ObjectPath) (InputDevice, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectInputDevice)
	obj.ImplObject.Init_(conn, "org.kde.KWin", path)
	return obj, nil
}

type inputDevice interface {
	Keyboard() proxy.PropBool
	AlphaNumericKeyboard() proxy.PropBool
	Pointer() proxy.PropBool
	Touchpad() proxy.PropBool
	Touch() proxy.PropBool
	TabletTool() proxy.PropBool
	TabletPad() proxy.PropBool
	GestureSupport() proxy.PropBool
	Name() proxy.PropString
	SysName() proxy.PropString
	OutputName() proxy.PropString
	Product() proxy.PropUint32
	Vendor() proxy.PropUint32
	SupportsDisableEvents() proxy.PropBool
	Enabled() proxy.PropBool
	SupportedButtons() proxy.PropInt32
	SupportsCalibrationMatrix() proxy.PropBool
	SupportsLeftHanded() proxy.PropBool
	LeftHandedEnabledByDefault() proxy.PropBool
	LeftHanded() proxy.PropBool
	SupportsDisableEventsOnExternalMouse() proxy.PropBool
	SupportsDisableWhileTyping() proxy.PropBool
	DisableWhileTypingEnabledByDefault() proxy.PropBool
	DisableWhileTyping() proxy.PropBool
	SupportsPointerAcceleration() proxy.PropBool
	DefaultPointerAcceleration() proxy.PropDouble
	PointerAcceleration() proxy.PropDouble
	SupportsPointerAccelerationProfileFlat() proxy.PropBool
	DefaultPointerAccelerationProfileFlat() proxy.PropBool
	PointerAccelerationProfileFlat() proxy.PropBool
	SupportsPointerAccelerationProfileAdaptive() proxy.PropBool
	DefaultPointerAccelerationProfileAdaptive() proxy.PropBool
	PointerAccelerationProfileAdaptive() proxy.PropBool
	TapFingerCount() proxy.PropInt32
	TapToClickEnabledByDefault() proxy.PropBool
	TapToClick() proxy.PropBool
	SupportsLmrTapButtonMap() proxy.PropBool
	LmrTapButtonMapEnabledByDefault() proxy.PropBool
	LmrTapButtonMap() proxy.PropBool
	TapAndDragEnabledByDefault() proxy.PropBool
	TapAndDrag() proxy.PropBool
	TapDragLockEnabledByDefault() proxy.PropBool
	TapDragLock() proxy.PropBool
	SupportsMiddleEmulation() proxy.PropBool
	MiddleEmulationEnabledByDefault() proxy.PropBool
	MiddleEmulation() proxy.PropBool
	SupportsNaturalScroll() proxy.PropBool
	NaturalScrollEnabledByDefault() proxy.PropBool
	NaturalScroll() proxy.PropBool
	SupportsScrollTwoFinger() proxy.PropBool
	ScrollTwoFingerEnabledByDefault() proxy.PropBool
	ScrollTwoFinger() proxy.PropBool
	SupportsScrollEdge() proxy.PropBool
	ScrollEdgeEnabledByDefault() proxy.PropBool
	ScrollEdge() proxy.PropBool
	SupportsScrollOnButtonDown() proxy.PropBool
	ScrollOnButtonDownEnabledByDefault() proxy.PropBool
	DefaultScrollButton() proxy.PropUint32
	ScrollOnButtonDown() proxy.PropBool
	ScrollButton() proxy.PropUint32
	SwitchDevice() proxy.PropBool
	LidSwitch() proxy.PropBool
	TabletModeSwitch() proxy.PropBool
	SupportsClickMethodAreas() proxy.PropBool
	DefaultClickMethodAreas() proxy.PropBool
	ClickMethodAreas() proxy.PropBool
	SupportsClickMethodClickfinger() proxy.PropBool
	DefaultClickMethodClickfinger() proxy.PropBool
	ClickMethodClickfinger() proxy.PropBool
}

type interfaceInputDevice struct{}

func (v *interfaceInputDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceInputDevice) GetInterfaceName_() string {
	return "org.kde.KWin.InputDevice"
}

// property keyboard b

func (v *interfaceInputDevice) Keyboard() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "keyboard",
	}
}

// property alphaNumericKeyboard b

func (v *interfaceInputDevice) AlphaNumericKeyboard() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "alphaNumericKeyboard",
	}
}

// property pointer b

func (v *interfaceInputDevice) Pointer() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "pointer",
	}
}

// property touchpad b

func (v *interfaceInputDevice) Touchpad() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "touchpad",
	}
}

// property touch b

func (v *interfaceInputDevice) Touch() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "touch",
	}
}

// property tabletTool b

func (v *interfaceInputDevice) TabletTool() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tabletTool",
	}
}

// property tabletPad b

func (v *interfaceInputDevice) TabletPad() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tabletPad",
	}
}

// property gestureSupport b

func (v *interfaceInputDevice) GestureSupport() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "gestureSupport",
	}
}

// property name s

func (v *interfaceInputDevice) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "name",
	}
}

// property sysName s

func (v *interfaceInputDevice) SysName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "sysName",
	}
}

// property outputName s

func (v *interfaceInputDevice) OutputName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "outputName",
	}
}

// property product u

func (v *interfaceInputDevice) Product() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "product",
	}
}

// property vendor u

func (v *interfaceInputDevice) Vendor() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "vendor",
	}
}

// property supportsDisableEvents b

func (v *interfaceInputDevice) SupportsDisableEvents() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsDisableEvents",
	}
}

// property enabled b

func (v *interfaceInputDevice) Enabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "enabled",
	}
}

// property supportedButtons i

func (v *interfaceInputDevice) SupportedButtons() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "supportedButtons",
	}
}

// property supportsCalibrationMatrix b

func (v *interfaceInputDevice) SupportsCalibrationMatrix() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsCalibrationMatrix",
	}
}

// property supportsLeftHanded b

func (v *interfaceInputDevice) SupportsLeftHanded() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsLeftHanded",
	}
}

// property leftHandedEnabledByDefault b

func (v *interfaceInputDevice) LeftHandedEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "leftHandedEnabledByDefault",
	}
}

// property leftHanded b

func (v *interfaceInputDevice) LeftHanded() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "leftHanded",
	}
}

// property supportsDisableEventsOnExternalMouse b

func (v *interfaceInputDevice) SupportsDisableEventsOnExternalMouse() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsDisableEventsOnExternalMouse",
	}
}

// property supportsDisableWhileTyping b

func (v *interfaceInputDevice) SupportsDisableWhileTyping() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsDisableWhileTyping",
	}
}

// property disableWhileTypingEnabledByDefault b

func (v *interfaceInputDevice) DisableWhileTypingEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "disableWhileTypingEnabledByDefault",
	}
}

// property disableWhileTyping b

func (v *interfaceInputDevice) DisableWhileTyping() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "disableWhileTyping",
	}
}

// property supportsPointerAcceleration b

func (v *interfaceInputDevice) SupportsPointerAcceleration() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsPointerAcceleration",
	}
}

// property defaultPointerAcceleration d

func (v *interfaceInputDevice) DefaultPointerAcceleration() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "defaultPointerAcceleration",
	}
}

// property pointerAcceleration d

func (v *interfaceInputDevice) PointerAcceleration() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "pointerAcceleration",
	}
}

// property supportsPointerAccelerationProfileFlat b

func (v *interfaceInputDevice) SupportsPointerAccelerationProfileFlat() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsPointerAccelerationProfileFlat",
	}
}

// property defaultPointerAccelerationProfileFlat b

func (v *interfaceInputDevice) DefaultPointerAccelerationProfileFlat() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "defaultPointerAccelerationProfileFlat",
	}
}

// property pointerAccelerationProfileFlat b

func (v *interfaceInputDevice) PointerAccelerationProfileFlat() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "pointerAccelerationProfileFlat",
	}
}

// property supportsPointerAccelerationProfileAdaptive b

func (v *interfaceInputDevice) SupportsPointerAccelerationProfileAdaptive() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsPointerAccelerationProfileAdaptive",
	}
}

// property defaultPointerAccelerationProfileAdaptive b

func (v *interfaceInputDevice) DefaultPointerAccelerationProfileAdaptive() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "defaultPointerAccelerationProfileAdaptive",
	}
}

// property pointerAccelerationProfileAdaptive b

func (v *interfaceInputDevice) PointerAccelerationProfileAdaptive() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "pointerAccelerationProfileAdaptive",
	}
}

// property tapFingerCount i

func (v *interfaceInputDevice) TapFingerCount() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "tapFingerCount",
	}
}

// property tapToClickEnabledByDefault b

func (v *interfaceInputDevice) TapToClickEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapToClickEnabledByDefault",
	}
}

// property tapToClick b

func (v *interfaceInputDevice) TapToClick() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapToClick",
	}
}

// property supportsLmrTapButtonMap b

func (v *interfaceInputDevice) SupportsLmrTapButtonMap() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsLmrTapButtonMap",
	}
}

// property lmrTapButtonMapEnabledByDefault b

func (v *interfaceInputDevice) LmrTapButtonMapEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "lmrTapButtonMapEnabledByDefault",
	}
}

// property lmrTapButtonMap b

func (v *interfaceInputDevice) LmrTapButtonMap() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "lmrTapButtonMap",
	}
}

// property tapAndDragEnabledByDefault b

func (v *interfaceInputDevice) TapAndDragEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapAndDragEnabledByDefault",
	}
}

// property tapAndDrag b

func (v *interfaceInputDevice) TapAndDrag() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapAndDrag",
	}
}

// property tapDragLockEnabledByDefault b

func (v *interfaceInputDevice) TapDragLockEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapDragLockEnabledByDefault",
	}
}

// property tapDragLock b

func (v *interfaceInputDevice) TapDragLock() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tapDragLock",
	}
}

// property supportsMiddleEmulation b

func (v *interfaceInputDevice) SupportsMiddleEmulation() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsMiddleEmulation",
	}
}

// property middleEmulationEnabledByDefault b

func (v *interfaceInputDevice) MiddleEmulationEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "middleEmulationEnabledByDefault",
	}
}

// property middleEmulation b

func (v *interfaceInputDevice) MiddleEmulation() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "middleEmulation",
	}
}

// property supportsNaturalScroll b

func (v *interfaceInputDevice) SupportsNaturalScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsNaturalScroll",
	}
}

// property naturalScrollEnabledByDefault b

func (v *interfaceInputDevice) NaturalScrollEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "naturalScrollEnabledByDefault",
	}
}

// property naturalScroll b

func (v *interfaceInputDevice) NaturalScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "naturalScroll",
	}
}

// property supportsScrollTwoFinger b

func (v *interfaceInputDevice) SupportsScrollTwoFinger() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsScrollTwoFinger",
	}
}

// property scrollTwoFingerEnabledByDefault b

func (v *interfaceInputDevice) ScrollTwoFingerEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollTwoFingerEnabledByDefault",
	}
}

// property scrollTwoFinger b

func (v *interfaceInputDevice) ScrollTwoFinger() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollTwoFinger",
	}
}

// property supportsScrollEdge b

func (v *interfaceInputDevice) SupportsScrollEdge() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsScrollEdge",
	}
}

// property scrollEdgeEnabledByDefault b

func (v *interfaceInputDevice) ScrollEdgeEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollEdgeEnabledByDefault",
	}
}

// property scrollEdge b

func (v *interfaceInputDevice) ScrollEdge() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollEdge",
	}
}

// property supportsScrollOnButtonDown b

func (v *interfaceInputDevice) SupportsScrollOnButtonDown() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsScrollOnButtonDown",
	}
}

// property scrollOnButtonDownEnabledByDefault b

func (v *interfaceInputDevice) ScrollOnButtonDownEnabledByDefault() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollOnButtonDownEnabledByDefault",
	}
}

// property defaultScrollButton u

func (v *interfaceInputDevice) DefaultScrollButton() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "defaultScrollButton",
	}
}

// property scrollOnButtonDown b

func (v *interfaceInputDevice) ScrollOnButtonDown() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "scrollOnButtonDown",
	}
}

// property scrollButton u

func (v *interfaceInputDevice) ScrollButton() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "scrollButton",
	}
}

// property switchDevice b

func (v *interfaceInputDevice) SwitchDevice() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "switchDevice",
	}
}

// property lidSwitch b

func (v *interfaceInputDevice) LidSwitch() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "lidSwitch",
	}
}

// property tabletModeSwitch b

func (v *interfaceInputDevice) TabletModeSwitch() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "tabletModeSwitch",
	}
}

// property supportsClickMethodAreas b

func (v *interfaceInputDevice) SupportsClickMethodAreas() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsClickMethodAreas",
	}
}

// property defaultClickMethodAreas b

func (v *interfaceInputDevice) DefaultClickMethodAreas() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "defaultClickMethodAreas",
	}
}

// property clickMethodAreas b

func (v *interfaceInputDevice) ClickMethodAreas() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "clickMethodAreas",
	}
}

// property supportsClickMethodClickfinger b

func (v *interfaceInputDevice) SupportsClickMethodClickfinger() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "supportsClickMethodClickfinger",
	}
}

// property defaultClickMethodClickfinger b

func (v *interfaceInputDevice) DefaultClickMethodClickfinger() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "defaultClickMethodClickfinger",
	}
}

// property clickMethodClickfinger b

func (v *interfaceInputDevice) ClickMethodClickfinger() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "clickMethodClickfinger",
	}
}
