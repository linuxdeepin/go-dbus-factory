package display

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

type Display struct {
	display // interface com.deepin.daemon.Display
	proxy.Object
}

func NewDisplay(conn *dbus.Conn) *Display {
	obj := new(Display)
	obj.Object.Init_(conn, "com.deepin.daemon.Display", "/com/deepin/daemon/Display")
	return obj
}

type display struct{}

func (v *display) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*display) GetInterfaceName_() string {
	return "com.deepin.daemon.Display"
}

// method ApplyChanges

func (v *display) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplyChanges", flags, ch)
}

func (v *display) ApplyChanges(flags dbus.Flags) error {
	return (<-v.GoApplyChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method AssociateTouch

func (v *display) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AssociateTouch", flags, ch, arg0, arg1)
}

func (v *display) AssociateTouch(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoAssociateTouch(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method ChangeBrightness

func (v *display) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeBrightness", flags, ch, arg0)
}

func (v *display) ChangeBrightness(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoChangeBrightness(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DeleteCustomMode

func (v *display) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteCustomMode", flags, ch, arg0)
}

func (v *display) DeleteCustomMode(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDeleteCustomMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetBrightness

func (v *display) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch)
}

func (*display) StoreGetBrightness(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) GetBrightness(flags dbus.Flags) (arg0 map[string]float64, err error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputNames

func (v *display) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputNames", flags, ch)
}

func (*display) StoreListOutputNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) ListOutputNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListOutputNames(
		<-v.GoListOutputNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListOutputsCommonModes

func (v *display) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListOutputsCommonModes", flags, ch)
}

func (*display) StoreListOutputsCommonModes(call *dbus.Call) (arg0 []ModeInfo, err error) {
	err = call.Store(&arg0)
	return
}

func (v *display) ListOutputsCommonModes(flags dbus.Flags) (arg0 []ModeInfo, err error) {
	return v.StoreListOutputsCommonModes(
		<-v.GoListOutputsCommonModes(flags, make(chan *dbus.Call, 1)).Done)
}

// method ModifyConfigName

func (v *display) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ModifyConfigName", flags, ch, arg0, arg1)
}

func (v *display) ModifyConfigName(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoModifyConfigName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method RefreshBrightness

func (v *display) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBrightness", flags, ch)
}

func (v *display) RefreshBrightness(flags dbus.Flags) error {
	return (<-v.GoRefreshBrightness(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reset

func (v *display) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *display) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ResetChanges

func (v *display) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ResetChanges", flags, ch)
}

func (v *display) ResetChanges(flags dbus.Flags) error {
	return (<-v.GoResetChanges(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Save

func (v *display) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Save", flags, ch)
}

func (v *display) Save(flags dbus.Flags) error {
	return (<-v.GoSave(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAndSaveBrightness

func (v *display) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAndSaveBrightness", flags, ch, arg0, arg1)
}

func (v *display) SetAndSaveBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetAndSaveBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetBrightness

func (v *display) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, arg0, arg1)
}

func (v *display) SetBrightness(flags dbus.Flags, arg0 string, arg1 float64) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPrimary

func (v *display) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrimary", flags, ch, arg0)
}

func (v *display) SetPrimary(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetPrimary(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SwitchMode

func (v *display) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint8, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchMode", flags, ch, arg0, arg1)
}

func (v *display) SwitchMode(flags dbus.Flags, arg0 uint8, arg1 string) error {
	return (<-v.GoSwitchMode(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// property HasChanged b

func (v *display) HasChanged() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasChanged",
	}
}

// property DisplayMode y

func (v *display) DisplayMode() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "DisplayMode",
	}
}

// property ScreenWidth q

func (v *display) ScreenWidth() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "ScreenWidth",
	}
}

// property ScreenHeight q

func (v *display) ScreenHeight() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "ScreenHeight",
	}
}

// property Primary s

func (v *display) Primary() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Primary",
	}
}

// property CurrentCustomId s

func (v *display) CurrentCustomId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CurrentCustomId",
	}
}

// property CustomIdList as

func (v *display) CustomIdList() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "CustomIdList",
	}
}

// property PrimaryRect (nnqq)

func (v *display) PrimaryRect() PropDisplayPrimaryRect {
	return PropDisplayPrimaryRect{
		Impl: v,
	}
}

type PropDisplayPrimaryRect struct {
	Impl proxy.Implementer
}

func (p PropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"PrimaryRect", &value)
	return
}

func (p PropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
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
		"PrimaryRect", cb0)
}

// property Monitors ao

func (v *display) Monitors() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Monitors",
	}
}

// property Brightness a{sd}

func (v *display) Brightness() PropDisplayBrightness {
	return PropDisplayBrightness{
		Impl: v,
	}
}

type PropDisplayBrightness struct {
	Impl proxy.Implementer
}

func (p PropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Brightness", &value)
	return
}

func (p PropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
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
		"Brightness", cb0)
}

// property TouchMap a{ss}

func (v *display) TouchMap() PropDisplayTouchMap {
	return PropDisplayTouchMap{
		Impl: v,
	}
}

type PropDisplayTouchMap struct {
	Impl proxy.Implementer
}

func (p PropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"TouchMap", &value)
	return
}

func (p PropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
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
		"TouchMap", cb0)
}

type Monitor struct {
	monitor // interface com.deepin.daemon.Display.Monitor
	proxy.Object
}

func NewMonitor(conn *dbus.Conn, path dbus.ObjectPath) (*Monitor, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Monitor)
	obj.Object.Init_(conn, "com.deepin.daemon.Display", path)
	return obj, nil
}

type monitor struct{}

func (v *monitor) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*monitor) GetInterfaceName_() string {
	return "com.deepin.daemon.Display.Monitor"
}

// method Enable

func (v *monitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, arg0)
}

func (v *monitor) Enable(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetMode

func (v *monitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMode", flags, ch, arg0)
}

func (v *monitor) SetMode(flags dbus.Flags, arg0 uint32) error {
	return (<-v.GoSetMode(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetModeBySize

func (v *monitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16, arg1 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetModeBySize", flags, ch, arg0, arg1)
}

func (v *monitor) SetModeBySize(flags dbus.Flags, arg0 uint16, arg1 uint16) error {
	return (<-v.GoSetModeBySize(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPosition

func (v *monitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, arg0 int16, arg1 int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, arg0, arg1)
}

func (v *monitor) SetPosition(flags dbus.Flags, arg0 int16, arg1 int16) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetReflect

func (v *monitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetReflect", flags, ch, arg0)
}

func (v *monitor) SetReflect(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetReflect(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRefreshRate

func (v *monitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRefreshRate", flags, ch, arg0)
}

func (v *monitor) SetRefreshRate(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetRefreshRate(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetRotation

func (v *monitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, arg0 uint16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRotation", flags, ch, arg0)
}

func (v *monitor) SetRotation(flags dbus.Flags, arg0 uint16) error {
	return (<-v.GoSetRotation(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// property Name s

func (v *monitor) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Enabled b

func (v *monitor) Enabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property Connected b

func (v *monitor) Connected() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Connected",
	}
}

// property X n

func (v *monitor) X() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "X",
	}
}

// property Y n

func (v *monitor) Y() proxy.PropInt16 {
	return proxy.PropInt16{
		Impl: v,
		Name: "Y",
	}
}

// property Width q

func (v *monitor) Width() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Width",
	}
}

// property Height q

func (v *monitor) Height() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Height",
	}
}

// property Rotation q

func (v *monitor) Rotation() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Rotation",
	}
}

// property Reflect q

func (v *monitor) Reflect() proxy.PropUint16 {
	return proxy.PropUint16{
		Impl: v,
		Name: "Reflect",
	}
}

// property RefreshRate d

func (v *monitor) RefreshRate() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "RefreshRate",
	}
}

// property Rotations aq

func (v *monitor) Rotations() proxy.PropUint16Array {
	return proxy.PropUint16Array{
		Impl: v,
		Name: "Rotations",
	}
}

// property Reflects aq

func (v *monitor) Reflects() proxy.PropUint16Array {
	return proxy.PropUint16Array{
		Impl: v,
		Name: "Reflects",
	}
}

// property BestMode (uqqd)

func (v *monitor) BestMode() PropModeInfo {
	return PropModeInfo{
		Impl: v,
		Name: "BestMode",
	}
}

// property CurrentMode (uqqd)

func (v *monitor) CurrentMode() PropModeInfo {
	return PropModeInfo{
		Impl: v,
		Name: "CurrentMode",
	}
}

// property Modes a(uqqd)

func (v *monitor) Modes() PropModeInfoSlice {
	return PropModeInfoSlice{
		Impl: v,
		Name: "Modes",
	}
}

// property PreferredModes a(uqqd)

func (v *monitor) PreferredModes() PropModeInfoSlice {
	return PropModeInfoSlice{
		Impl: v,
		Name: "PreferredModes",
	}
}

type PropModeInfo struct {
	Impl proxy.Implementer
	Name string
}

func (p PropModeInfo) Get(flags dbus.Flags) (value ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropModeInfo) Set(flags dbus.Flags, value ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropModeInfo) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
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

type PropModeInfoSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p PropModeInfoSlice) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropModeInfoSlice) Set(flags dbus.Flags, value []ModeInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropModeInfoSlice) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
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
