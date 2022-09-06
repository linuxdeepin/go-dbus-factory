// Code generated by "./generator ./com.deepin.daemon.display"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package display

import "errors"

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Display interface {
	display // interface com.deepin.daemon.Display
	proxy.Object
}

type objectDisplay struct {
	interfaceDisplay // interface com.deepin.daemon.Display
	proxy.ImplObject
}

func NewDisplay(conn *dbus.Conn) Display {
	obj := new(objectDisplay)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Display", "/com/deepin/daemon/Display")
	return obj
}

type display interface {
	GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ApplyChanges(flags dbus.Flags) error
	GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, outputName string, touchSerial string) *dbus.Call
	AssociateTouch(flags dbus.Flags, outputName string, touchSerial string) error
	GoCanRotate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanRotate(flags dbus.Flags) (bool, error)
	GoCanSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call
	CanSetBrightness(flags dbus.Flags, outputName string) (bool, error)
	GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, raised bool) *dbus.Call
	ChangeBrightness(flags dbus.Flags, raised bool) error
	GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	DeleteCustomMode(flags dbus.Flags, name string) error
	GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetBrightness(flags dbus.Flags) (map[string]float64, error)
	GoGetBuiltinMonitor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetBuiltinMonitor(flags dbus.Flags) (string, dbus.ObjectPath, error)
	GoGetRealDisplayMode(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetRealDisplayMode(flags dbus.Flags) (uint8, error)
	GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListOutputNames(flags dbus.Flags) ([]string, error)
	GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListOutputsCommonModes(flags dbus.Flags) ([][]interface{}, error)
	GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, name string, newName string) *dbus.Call
	ModifyConfigName(flags dbus.Flags, name string, newName string) error
	GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RefreshBrightness(flags dbus.Flags) error
	GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Reset(flags dbus.Flags) error
	GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ResetChanges(flags dbus.Flags) error
	GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Save(flags dbus.Flags) error
	GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call
	SetAndSaveBrightness(flags dbus.Flags, outputName string, value float64) error
	GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call
	SetBrightness(flags dbus.Flags, outputName string, value float64) error
	GoSetColorTemperature(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetColorTemperature(flags dbus.Flags, value int32) error
	GoSetMethodAdjustCCT(flags dbus.Flags, ch chan *dbus.Call, adjustMethod int32) *dbus.Call
	SetMethodAdjustCCT(flags dbus.Flags, adjustMethod int32) error
	GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call
	SetPrimary(flags dbus.Flags, outputName string) error
	GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, mode uint8, name string) *dbus.Call
	SwitchMode(flags dbus.Flags, mode uint8, name string) error
	ColorTemperatureMode() proxy.PropInt32
	HasChanged() proxy.PropBool
	TouchMap() PropDisplayTouchMap
	PrimaryRect() PropDisplayPrimaryRect
	ScreenWidth() proxy.PropUint16
	MaxBacklightBrightness() proxy.PropUint32
	ColorTemperatureManual() proxy.PropInt32
	DisplayMode() proxy.PropByte
	CustomIdList() proxy.PropStringArray
	CurrentCustomId() proxy.PropString
	ScreenHeight() proxy.PropUint16
	Monitors() proxy.PropObjectPathArray
	Touchscreens() PropTouchscreens
	Primary() proxy.PropString
	Brightness() PropDisplayBrightness
}

type interfaceDisplay struct{}

func (v *interfaceDisplay) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDisplay) GetInterfaceName_() string {
	return "com.deepin.daemon.Display"
}

// method ApplyChanges

func (v *interfaceDisplay) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplyChanges", flags, ch)
}

func (v *interfaceDisplay) ApplyChanges(flags dbus.Flags) error {
	return (<-v.GoApplyChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AssociateTouch

func (v *interfaceDisplay) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, outputName string, touchSerial string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AssociateTouch", flags, ch, outputName, touchSerial)
}

func (v *interfaceDisplay) AssociateTouch(flags dbus.Flags, outputName string, touchSerial string) error {
	return (<-v.GoAssociateTouch(flags, make(chan *dbus.Call, 1), outputName, touchSerial).Done).Err
}

// method CanRotate

func (v *interfaceDisplay) GoCanRotate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanRotate", flags, ch)
}

func (*interfaceDisplay) StoreCanRotate(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) CanRotate(flags dbus.Flags) (bool, error) {
	return v.StoreCanRotate(
		<-v.GoCanRotate(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSetBrightness

func (v *interfaceDisplay) GoCanSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSetBrightness", flags, ch, outputName)
}

func (*interfaceDisplay) StoreCanSetBrightness(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) CanSetBrightness(flags dbus.Flags, outputName string) (bool, error) {
	return v.StoreCanSetBrightness(
		<-v.GoCanSetBrightness(flags, make(chan *dbus.Call, 1), outputName).Done)
}

// method ChangeBrightness

func (v *interfaceDisplay) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, raised bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeBrightness", flags, ch, raised)
}

func (v *interfaceDisplay) ChangeBrightness(flags dbus.Flags, raised bool) error {
	return (<-v.GoChangeBrightness(flags, make(chan *dbus.Call, 1), raised).Done).Err
}

// method DeleteCustomMode

func (v *interfaceDisplay) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCustomMode", flags, ch, name)
}

func (v *interfaceDisplay) DeleteCustomMode(flags dbus.Flags, name string) error {
	return (<-v.GoDeleteCustomMode(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method GetBrightness

func (v *interfaceDisplay) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch)
}

func (*interfaceDisplay) StoreGetBrightness(call *dbus.Call) (outArg0 map[string]float64, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) GetBrightness(flags dbus.Flags) (map[string]float64, error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetBuiltinMonitor

func (v *interfaceDisplay) GoGetBuiltinMonitor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBuiltinMonitor", flags, ch)
}

func (*interfaceDisplay) StoreGetBuiltinMonitor(call *dbus.Call) (outArg0 string, outArg1 dbus.ObjectPath, err error) {
	err = call.Store(&outArg0, &outArg1)
	return
}

func (v *interfaceDisplay) GetBuiltinMonitor(flags dbus.Flags) (string, dbus.ObjectPath, error) {
	return v.StoreGetBuiltinMonitor(
		<-v.GoGetBuiltinMonitor(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetRealDisplayMode

func (v *interfaceDisplay) GoGetRealDisplayMode(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetRealDisplayMode", flags, ch)
}

func (*interfaceDisplay) StoreGetRealDisplayMode(call *dbus.Call) (outArg0 uint8, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) GetRealDisplayMode(flags dbus.Flags) (uint8, error) {
	return v.StoreGetRealDisplayMode(
		<-v.GoGetRealDisplayMode(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputNames

func (v *interfaceDisplay) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputNames", flags, ch)
}

func (*interfaceDisplay) StoreListOutputNames(call *dbus.Call) (outArg0 []string, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) ListOutputNames(flags dbus.Flags) ([]string, error) {
	return v.StoreListOutputNames(
		<-v.GoListOutputNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputsCommonModes

func (v *interfaceDisplay) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputsCommonModes", flags, ch)
}

func (*interfaceDisplay) StoreListOutputsCommonModes(call *dbus.Call) (outArg0 [][]interface{}, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceDisplay) ListOutputsCommonModes(flags dbus.Flags) ([][]interface{}, error) {
	return v.StoreListOutputsCommonModes(
		<-v.GoListOutputsCommonModes(flags, make(chan *dbus.Call, 1)).Done)
}

// method ModifyConfigName

func (v *interfaceDisplay) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, name string, newName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ModifyConfigName", flags, ch, name, newName)
}

func (v *interfaceDisplay) ModifyConfigName(flags dbus.Flags, name string, newName string) error {
	return (<-v.GoModifyConfigName(flags, make(chan *dbus.Call, 1), name, newName).Done).Err
}

// method RefreshBrightness

func (v *interfaceDisplay) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBrightness", flags, ch)
}

func (v *interfaceDisplay) RefreshBrightness(flags dbus.Flags) error {
	return (<-v.GoRefreshBrightness(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reset

func (v *interfaceDisplay) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *interfaceDisplay) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ResetChanges

func (v *interfaceDisplay) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetChanges", flags, ch)
}

func (v *interfaceDisplay) ResetChanges(flags dbus.Flags) error {
	return (<-v.GoResetChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *interfaceDisplay) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *interfaceDisplay) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAndSaveBrightness

func (v *interfaceDisplay) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAndSaveBrightness", flags, ch, outputName, value)
}

func (v *interfaceDisplay) SetAndSaveBrightness(flags dbus.Flags, outputName string, value float64) error {
	return (<-v.GoSetAndSaveBrightness(flags, make(chan *dbus.Call, 1), outputName, value).Done).Err
}

// method SetBrightness

func (v *interfaceDisplay) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, outputName, value)
}

func (v *interfaceDisplay) SetBrightness(flags dbus.Flags, outputName string, value float64) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), outputName, value).Done).Err
}

// method SetColorTemperature

func (v *interfaceDisplay) GoSetColorTemperature(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetColorTemperature", flags, ch, value)
}

func (v *interfaceDisplay) SetColorTemperature(flags dbus.Flags, value int32) error {
	return (<-v.GoSetColorTemperature(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetMethodAdjustCCT

func (v *interfaceDisplay) GoSetMethodAdjustCCT(flags dbus.Flags, ch chan *dbus.Call, adjustMethod int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMethodAdjustCCT", flags, ch, adjustMethod)
}

func (v *interfaceDisplay) SetMethodAdjustCCT(flags dbus.Flags, adjustMethod int32) error {
	return (<-v.GoSetMethodAdjustCCT(flags, make(chan *dbus.Call, 1), adjustMethod).Done).Err
}

// method SetPrimary

func (v *interfaceDisplay) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrimary", flags, ch, outputName)
}

func (v *interfaceDisplay) SetPrimary(flags dbus.Flags, outputName string) error {
	return (<-v.GoSetPrimary(flags, make(chan *dbus.Call, 1), outputName).Done).Err
}

// method SwitchMode

func (v *interfaceDisplay) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, mode uint8, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchMode", flags, ch, mode, name)
}

func (v *interfaceDisplay) SwitchMode(flags dbus.Flags, mode uint8, name string) error {
	return (<-v.GoSwitchMode(flags, make(chan *dbus.Call, 1), mode, name).Done).Err
}

// property ColorTemperatureMode i

func (v *interfaceDisplay) ColorTemperatureMode() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ColorTemperatureMode",
	}
}

// property HasChanged b

func (v *interfaceDisplay) HasChanged() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasChanged",
	}
}

type PropDisplayTouchMap interface {
	Get(flags dbus.Flags) (value map[string]string, err error)
	Set(flags dbus.Flags, value map[string]string) error
	ConnectChanged(cb func(hasValue bool, value map[string]string)) error
}

type implPropDisplayTouchMap struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayTouchMap) Set(flags dbus.Flags, value map[string]string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property TouchMap a{ss}

func (v *interfaceDisplay) TouchMap() PropDisplayTouchMap {
	return &implPropDisplayTouchMap{
		Impl: v,
		Name: "TouchMap",
	}
}

type PropDisplayPrimaryRect interface {
	Get(flags dbus.Flags) (value Rectangle, err error)
	Set(flags dbus.Flags, value Rectangle) error
	ConnectChanged(cb func(hasValue bool, value Rectangle)) error
}

type implPropDisplayPrimaryRect struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayPrimaryRect) Set(flags dbus.Flags, value Rectangle) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v Rectangle
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, Rectangle{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property PrimaryRect (nnqq)

func (v *interfaceDisplay) PrimaryRect() PropDisplayPrimaryRect {
	return &implPropDisplayPrimaryRect{
		Impl: v,
		Name: "PrimaryRect",
	}
}

// property ScreenWidth q

func (v *interfaceDisplay) ScreenWidth() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "ScreenWidth",
	}
}

// property MaxBacklightBrightness u

func (v *interfaceDisplay) MaxBacklightBrightness() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "MaxBacklightBrightness",
	}
}

// property ColorTemperatureManual i

func (v *interfaceDisplay) ColorTemperatureManual() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ColorTemperatureManual",
	}
}

// property DisplayMode y

func (v *interfaceDisplay) DisplayMode() proxy.PropByte {
	return &proxy.ImplPropByte{
		Impl: v,
		Name: "DisplayMode",
	}
}

// property CustomIdList as

func (v *interfaceDisplay) CustomIdList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "CustomIdList",
	}
}

// property CurrentCustomId s

func (v *interfaceDisplay) CurrentCustomId() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CurrentCustomId",
	}
}

// property ScreenHeight q

func (v *interfaceDisplay) ScreenHeight() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "ScreenHeight",
	}
}

// property Monitors ao

func (v *interfaceDisplay) Monitors() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Monitors",
	}
}

type PropTouchscreens interface {
	Get(flags dbus.Flags) (value []Touchscreen, err error)
	Set(flags dbus.Flags, value []Touchscreen) error
	ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error
}

type implPropTouchscreens struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropTouchscreens) Get(flags dbus.Flags) (value []Touchscreen, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropTouchscreens) Set(flags dbus.Flags, value []Touchscreen) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropTouchscreens) ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []Touchscreen
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Touchscreens a(isss)

func (v *interfaceDisplay) Touchscreens() PropTouchscreens {
	return &implPropTouchscreens{
		Impl: v,
		Name: "Touchscreens",
	}
}

// property Primary s

func (v *interfaceDisplay) Primary() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Primary",
	}
}

type PropDisplayBrightness interface {
	Get(flags dbus.Flags) (value map[string]float64, err error)
	Set(flags dbus.Flags, value map[string]float64) error
	ConnectChanged(cb func(hasValue bool, value map[string]float64)) error
}

type implPropDisplayBrightness struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDisplayBrightness) Set(flags dbus.Flags, value map[string]float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]float64
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Brightness a{sd}

func (v *interfaceDisplay) Brightness() PropDisplayBrightness {
	return &implPropDisplayBrightness{
		Impl: v,
		Name: "Brightness",
	}
}

type Monitor interface {
	monitor // interface com.deepin.daemon.Display.Monitor
	proxy.Object
}

type objectMonitor struct {
	interfaceMonitor // interface com.deepin.daemon.Display.Monitor
	proxy.ImplObject
}

func NewMonitor(conn *dbus.Conn, path dbus.ObjectPath) (Monitor, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectMonitor)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Display", path)
	return obj, nil
}

type monitor interface {
	GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	Enable(flags dbus.Flags, enabled bool) error
	GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode uint32) *dbus.Call
	SetMode(flags dbus.Flags, mode uint32) error
	GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, width uint16, height uint16) *dbus.Call
	SetModeBySize(flags dbus.Flags, width uint16, height uint16) error
	GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, X int16, y int16) *dbus.Call
	SetPosition(flags dbus.Flags, X int16, y int16) error
	GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call
	SetReflect(flags dbus.Flags, value uint16) error
	GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call
	SetRefreshRate(flags dbus.Flags, value float64) error
	GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call
	SetRotation(flags dbus.Flags, value uint16) error
	MmHeight() proxy.PropUint32
	Width() proxy.PropUint16
	Rotation() proxy.PropUint16
	RefreshRate() proxy.PropDouble
	CurrentRotateMode() proxy.PropByte
	AvailableFillModes() proxy.PropStringArray
	ID() proxy.PropUint32
	Rotations() proxy.PropUint16Array
	BestMode() PropModeInfo
	MmWidth() proxy.PropUint32
	Brightness() proxy.PropDouble
	CurrentMode() PropModeInfo
	Name() proxy.PropString
	Manufacturer() proxy.PropString
	X() proxy.PropInt16
	Y() proxy.PropInt16
	Height() proxy.PropUint16
	CurrentFillMode() proxy.PropString
	Connected() proxy.PropBool
	Model() proxy.PropString
	Reflects() proxy.PropUint16Array
	Modes() PropModeInfoSlice
	PreferredModes() PropModeInfoSlice
	Enabled() proxy.PropBool
	Reflect() proxy.PropUint16
}

type interfaceMonitor struct{}

func (v *interfaceMonitor) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceMonitor) GetInterfaceName_() string {
	return "com.deepin.daemon.Display.Monitor"
}

// method Enable

func (v *interfaceMonitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *interfaceMonitor) Enable(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method SetMode

func (v *interfaceMonitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMode", flags, ch, mode)
}

func (v *interfaceMonitor) SetMode(flags dbus.Flags, mode uint32) error {
	return (<-v.GoSetMode(flags, make(chan *dbus.Call, 1), mode).Done).Err
}

// method SetModeBySize

func (v *interfaceMonitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, width uint16, height uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetModeBySize", flags, ch, width, height)
}

func (v *interfaceMonitor) SetModeBySize(flags dbus.Flags, width uint16, height uint16) error {
	return (<-v.GoSetModeBySize(flags, make(chan *dbus.Call, 1), width, height).Done).Err
}

// method SetPosition

func (v *interfaceMonitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, X int16, y int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, X, y)
}

func (v *interfaceMonitor) SetPosition(flags dbus.Flags, X int16, y int16) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), X, y).Done).Err
}

// method SetReflect

func (v *interfaceMonitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetReflect", flags, ch, value)
}

func (v *interfaceMonitor) SetReflect(flags dbus.Flags, value uint16) error {
	return (<-v.GoSetReflect(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetRefreshRate

func (v *interfaceMonitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRefreshRate", flags, ch, value)
}

func (v *interfaceMonitor) SetRefreshRate(flags dbus.Flags, value float64) error {
	return (<-v.GoSetRefreshRate(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetRotation

func (v *interfaceMonitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRotation", flags, ch, value)
}

func (v *interfaceMonitor) SetRotation(flags dbus.Flags, value uint16) error {
	return (<-v.GoSetRotation(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// property MmHeight u

func (v *interfaceMonitor) MmHeight() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "MmHeight",
	}
}

// property Width q

func (v *interfaceMonitor) Width() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Width",
	}
}

// property Rotation q

func (v *interfaceMonitor) Rotation() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Rotation",
	}
}

// property RefreshRate d

func (v *interfaceMonitor) RefreshRate() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "RefreshRate",
	}
}

// property CurrentRotateMode y

func (v *interfaceMonitor) CurrentRotateMode() proxy.PropByte {
	return &proxy.ImplPropByte{
		Impl: v,
		Name: "CurrentRotateMode",
	}
}

// property AvailableFillModes as

func (v *interfaceMonitor) AvailableFillModes() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "AvailableFillModes",
	}
}

// property ID u

func (v *interfaceMonitor) ID() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "ID",
	}
}

// property Rotations aq

func (v *interfaceMonitor) Rotations() proxy.PropUint16Array {
	return &proxy.ImplPropUint16Array{
		Impl: v,
		Name: "Rotations",
	}
}

// property BestMode (uqqd)

func (v *interfaceMonitor) BestMode() PropModeInfo {
	return &implPropModeInfo{
		Impl: v,
		Name: "BestMode",
	}
}

// property MmWidth u

func (v *interfaceMonitor) MmWidth() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "MmWidth",
	}
}

// property Brightness d

func (v *interfaceMonitor) Brightness() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Brightness",
	}
}

// property CurrentMode (uqqd)

func (v *interfaceMonitor) CurrentMode() PropModeInfo {
	return &implPropModeInfo{
		Impl: v,
		Name: "CurrentMode",
	}
}

// property Name s

func (v *interfaceMonitor) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property Manufacturer s

func (v *interfaceMonitor) Manufacturer() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property X n

func (v *interfaceMonitor) X() proxy.PropInt16 {
	return &proxy.ImplPropInt16{
		Impl: v,
		Name: "X",
	}
}

// property Y n

func (v *interfaceMonitor) Y() proxy.PropInt16 {
	return &proxy.ImplPropInt16{
		Impl: v,
		Name: "Y",
	}
}

// property Height q

func (v *interfaceMonitor) Height() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Height",
	}
}

// property CurrentFillMode s

func (v *interfaceMonitor) CurrentFillMode() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CurrentFillMode",
	}
}

// property Connected b

func (v *interfaceMonitor) Connected() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property Model s

func (v *interfaceMonitor) Model() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Model",
	}
}

// property Reflects aq

func (v *interfaceMonitor) Reflects() proxy.PropUint16Array {
	return &proxy.ImplPropUint16Array{
		Impl: v,
		Name: "Reflects",
	}
}

// property Modes a(uqqd)

func (v *interfaceMonitor) Modes() PropModeInfoSlice {
	return &implPropModeInfoSlice{
		Impl: v,
		Name: "Modes",
	}
}

// property PreferredModes a(uqqd)

func (v *interfaceMonitor) PreferredModes() PropModeInfoSlice {
	return &implPropModeInfoSlice{
		Impl: v,
		Name: "PreferredModes",
	}
}

// property Enabled b

func (v *interfaceMonitor) Enabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property Reflect q

func (v *interfaceMonitor) Reflect() proxy.PropUint16 {
	return &proxy.ImplPropUint16{
		Impl: v,
		Name: "Reflect",
	}
}

type PropModeInfo interface {
	Get(flags dbus.Flags) (value ModeInfo, err error)
	Set(flags dbus.Flags, value ModeInfo) error
	ConnectChanged(cb func(hasValue bool, value ModeInfo)) error
}

type implPropModeInfo struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropModeInfo) Get(flags dbus.Flags) (value ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropModeInfo) Set(flags dbus.Flags, value ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropModeInfo) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ModeInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropModeInfoSlice interface {
	Get(flags dbus.Flags) (value []ModeInfo, err error)
	Set(flags dbus.Flags, value []ModeInfo) error
	ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error
}

type implPropModeInfoSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropModeInfoSlice) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropModeInfoSlice) Set(flags dbus.Flags, value []ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropModeInfoSlice) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ModeInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
