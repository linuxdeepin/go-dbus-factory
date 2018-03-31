package audio

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

type Audio struct {
	audio // interface com.deepin.daemon.Audio
	proxy.Object
}

func NewAudio(conn *dbus.Conn) *Audio {
	obj := new(Audio)
	obj.Object.Init_(conn, "com.deepin.daemon.Audio", "/com/deepin/daemon/Audio")
	return obj
}

type audio struct{}

func (v *audio) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*audio) GetInterfaceName_() string {
	return "com.deepin.daemon.Audio"
}

// method SetDefaultSink

func (v *audio) GoSetDefaultSink(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultSink", flags, ch, name)
}

func (v *audio) SetDefaultSink(flags dbus.Flags, name string) error {
	return (<-v.GoSetDefaultSink(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetDefaultSource

func (v *audio) GoSetDefaultSource(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultSource", flags, ch, name)
}

func (v *audio) SetDefaultSource(flags dbus.Flags, name string) error {
	return (<-v.GoSetDefaultSource(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetPort

func (v *audio) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, cardId uint32, portName string, direction int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPort", flags, ch, cardId, portName, direction)
}

func (v *audio) SetPort(flags dbus.Flags, cardId uint32, portName string, direction int32) error {
	return (<-v.GoSetPort(flags, make(chan *dbus.Call, 1), cardId, portName, direction).Done).Err
}

// property MaxUIVolume d

func (v *audio) MaxUIVolume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "MaxUIVolume",
	}
}

// property SinkInputs ao

func (v *audio) SinkInputs() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "SinkInputs",
	}
}

// property DefaultSink o

func (v *audio) DefaultSink() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "DefaultSink",
	}
}

// property DefaultSource o

func (v *audio) DefaultSource() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "DefaultSource",
	}
}

// property Cards s

func (v *audio) Cards() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Cards",
	}
}

type Sink struct {
	sink // interface com.deepin.daemon.Audio.Sink
	proxy.Object
}

func NewSink(conn *dbus.Conn, path dbus.ObjectPath) (*Sink, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Sink)
	obj.Object.Init_(conn, "com.deepin.daemon.Audio", path)
	return obj, nil
}

type sink struct{}

func (v *sink) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sink) GetInterfaceName_() string {
	return "com.deepin.daemon.Audio.Sink"
}

// method GetMeter

func (v *sink) GoGetMeter(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetMeter", flags, ch)
}

func (*sink) StoreGetMeter(call *dbus.Call) (meter dbus.ObjectPath, err error) {
	err = call.Store(&meter)
	return
}

func (v *sink) GetMeter(flags dbus.Flags) (meter dbus.ObjectPath, err error) {
	return v.StoreGetMeter(
		<-v.GoGetMeter(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetBalance

func (v *sink) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBalance", flags, ch, value, isPlay)
}

func (v *sink) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetBalance(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// method SetFade

func (v *sink) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFade", flags, ch, value)
}

func (v *sink) SetFade(flags dbus.Flags, value float64) error {
	return (<-v.GoSetFade(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetMute

func (v *sink) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMute", flags, ch, value)
}

func (v *sink) SetMute(flags dbus.Flags, value bool) error {
	return (<-v.GoSetMute(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetPort

func (v *sink) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPort", flags, ch, name)
}

func (v *sink) SetPort(flags dbus.Flags, name string) error {
	return (<-v.GoSetPort(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetVolume

func (v *sink) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetVolume", flags, ch, value, isPlay)
}

func (v *sink) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetVolume(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// property SupportBalance b

func (v *sink) SupportBalance() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportBalance",
	}
}

// property Ports a(ssy)

func (v *sink) Ports() PropSinkPorts {
	return PropSinkPorts{
		Impl: v,
	}
}

type PropSinkPorts struct {
	Impl proxy.Implementer
}

func (p PropSinkPorts) Get(flags dbus.Flags) (value []PortInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Ports", &value)
	return
}

func (p PropSinkPorts) ConnectChanged(cb func(hasValue bool, value []PortInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []PortInfo
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
		"Ports", cb0)
}

// property Name s

func (v *sink) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Mute b

func (v *sink) Mute() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Mute",
	}
}

// property Volume d

func (v *sink) Volume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Volume",
	}
}

// property Balance d

func (v *sink) Balance() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Balance",
	}
}

// property ActivePort (ssy)

func (v *sink) ActivePort() PropSinkActivePort {
	return PropSinkActivePort{
		Impl: v,
	}
}

type PropSinkActivePort struct {
	Impl proxy.Implementer
}

func (p PropSinkActivePort) Get(flags dbus.Flags) (value PortInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"ActivePort", &value)
	return
}

func (p PropSinkActivePort) ConnectChanged(cb func(hasValue bool, value PortInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v PortInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, PortInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"ActivePort", cb0)
}

// property Card u

func (v *sink) Card() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Card",
	}
}

// property Description s

func (v *sink) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property BaseVolume d

func (v *sink) BaseVolume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "BaseVolume",
	}
}

// property Fade d

func (v *sink) Fade() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Fade",
	}
}

// property SupportFade b

func (v *sink) SupportFade() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportFade",
	}
}

type Source struct {
	source // interface com.deepin.daemon.Audio.Source
	proxy.Object
}

func NewSource(conn *dbus.Conn, path dbus.ObjectPath) (*Source, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Source)
	obj.Object.Init_(conn, "com.deepin.daemon.Audio", path)
	return obj, nil
}

type source struct{}

func (v *source) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*source) GetInterfaceName_() string {
	return "com.deepin.daemon.Audio.Source"
}

// method GetMeter

func (v *source) GoGetMeter(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetMeter", flags, ch)
}

func (*source) StoreGetMeter(call *dbus.Call) (meter dbus.ObjectPath, err error) {
	err = call.Store(&meter)
	return
}

func (v *source) GetMeter(flags dbus.Flags) (meter dbus.ObjectPath, err error) {
	return v.StoreGetMeter(
		<-v.GoGetMeter(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetBalance

func (v *source) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBalance", flags, ch, value, isPlay)
}

func (v *source) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetBalance(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// method SetFade

func (v *source) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFade", flags, ch, value)
}

func (v *source) SetFade(flags dbus.Flags, value float64) error {
	return (<-v.GoSetFade(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetMute

func (v *source) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMute", flags, ch, value)
}

func (v *source) SetMute(flags dbus.Flags, value bool) error {
	return (<-v.GoSetMute(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetPort

func (v *source) GoSetPort(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPort", flags, ch, name)
}

func (v *source) SetPort(flags dbus.Flags, name string) error {
	return (<-v.GoSetPort(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetVolume

func (v *source) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetVolume", flags, ch, value, isPlay)
}

func (v *source) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetVolume(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// property Mute b

func (v *source) Mute() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Mute",
	}
}

// property Balance d

func (v *source) Balance() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Balance",
	}
}

// property SupportBalance b

func (v *source) SupportBalance() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportBalance",
	}
}

// property Fade d

func (v *source) Fade() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Fade",
	}
}

// property Ports a(ssy)

func (v *source) Ports() PropSourcePorts {
	return PropSourcePorts{
		Impl: v,
	}
}

type PropSourcePorts struct {
	Impl proxy.Implementer
}

func (p PropSourcePorts) Get(flags dbus.Flags) (value []PortInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Ports", &value)
	return
}

func (p PropSourcePorts) ConnectChanged(cb func(hasValue bool, value []PortInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []PortInfo
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
		"Ports", cb0)
}

// property Card u

func (v *source) Card() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Card",
	}
}

// property BaseVolume d

func (v *source) BaseVolume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "BaseVolume",
	}
}

// property Description s

func (v *source) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property Volume d

func (v *source) Volume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Volume",
	}
}

// property SupportFade b

func (v *source) SupportFade() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportFade",
	}
}

// property ActivePort (ssy)

func (v *source) ActivePort() PropSourceActivePort {
	return PropSourceActivePort{
		Impl: v,
	}
}

type PropSourceActivePort struct {
	Impl proxy.Implementer
}

func (p PropSourceActivePort) Get(flags dbus.Flags) (value PortInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"ActivePort", &value)
	return
}

func (p PropSourceActivePort) ConnectChanged(cb func(hasValue bool, value PortInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v PortInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, PortInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"ActivePort", cb0)
}

// property Name s

func (v *source) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

type SinkInput struct {
	sinkInput // interface com.deepin.daemon.Audio.SinkInput
	proxy.Object
}

func NewSinkInput(conn *dbus.Conn, path dbus.ObjectPath) (*SinkInput, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(SinkInput)
	obj.Object.Init_(conn, "com.deepin.daemon.Audio", path)
	return obj, nil
}

type sinkInput struct{}

func (v *sinkInput) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sinkInput) GetInterfaceName_() string {
	return "com.deepin.daemon.Audio.SinkInput"
}

// method SetBalance

func (v *sinkInput) GoSetBalance(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBalance", flags, ch, value, isPlay)
}

func (v *sinkInput) SetBalance(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetBalance(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// method SetFade

func (v *sinkInput) GoSetFade(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFade", flags, ch, value)
}

func (v *sinkInput) SetFade(flags dbus.Flags, value float64) error {
	return (<-v.GoSetFade(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetMute

func (v *sinkInput) GoSetMute(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMute", flags, ch, value)
}

func (v *sinkInput) SetMute(flags dbus.Flags, value bool) error {
	return (<-v.GoSetMute(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetVolume

func (v *sinkInput) GoSetVolume(flags dbus.Flags, ch chan *dbus.Call, value float64, isPlay bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetVolume", flags, ch, value, isPlay)
}

func (v *sinkInput) SetVolume(flags dbus.Flags, value float64, isPlay bool) error {
	return (<-v.GoSetVolume(flags, make(chan *dbus.Call, 1), value, isPlay).Done).Err
}

// property Volume d

func (v *sinkInput) Volume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Volume",
	}
}

// property Balance d

func (v *sinkInput) Balance() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Balance",
	}
}

// property SupportBalance b

func (v *sinkInput) SupportBalance() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportBalance",
	}
}

// property Fade d

func (v *sinkInput) Fade() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Fade",
	}
}

// property SupportFade b

func (v *sinkInput) SupportFade() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SupportFade",
	}
}

// property Name s

func (v *sinkInput) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Icon s

func (v *sinkInput) Icon() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Icon",
	}
}

// property Mute b

func (v *sinkInput) Mute() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Mute",
	}
}
