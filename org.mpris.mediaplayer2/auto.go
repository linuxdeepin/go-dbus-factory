package mediaplayer2

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

type MediaPlayer struct {
	mediaPlayer // interface org.mpris.MediaPlayer2
	player      // interface org.mpris.MediaPlayer2.Player
	proxy.Object
}

func NewMediaPlayer(conn *dbus.Conn, serviceName string) *MediaPlayer {
	obj := new(MediaPlayer)
	obj.Object.Init_(conn, serviceName, "/org/mpris/MediaPlayer2")
	return obj
}

type mediaPlayer struct{}

func (v *mediaPlayer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*mediaPlayer) GetInterfaceName_() string {
	return "org.mpris.MediaPlayer2"
}

// method Quit

func (v *mediaPlayer) GoQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Quit", flags, ch)
}

func (v *mediaPlayer) Quit(flags dbus.Flags) error {
	return (<-v.GoQuit(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Raise

func (v *mediaPlayer) GoRaise(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Raise", flags, ch)
}

func (v *mediaPlayer) Raise(flags dbus.Flags) error {
	return (<-v.GoRaise(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property CanQuit b

func (v *mediaPlayer) CanQuit() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanQuit",
	}
}

// property CanRaise b

func (v *mediaPlayer) CanRaise() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanRaise",
	}
}

// property DesktopEntry s

func (v *mediaPlayer) DesktopEntry() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DesktopEntry",
	}
}

// property Identity s

func (v *mediaPlayer) Identity() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Identity",
	}
}

type player struct{}

func (v *player) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*player) GetInterfaceName_() string {
	return "org.mpris.MediaPlayer2.Player"
}

// method Next

func (v *player) GoNext(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Next", flags, ch)
}

func (v *player) Next(flags dbus.Flags) error {
	return (<-v.GoNext(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Pause

func (v *player) GoPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pause", flags, ch)
}

func (v *player) Pause(flags dbus.Flags) error {
	return (<-v.GoPause(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Play

func (v *player) GoPlay(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Play", flags, ch)
}

func (v *player) Play(flags dbus.Flags) error {
	return (<-v.GoPlay(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PlayPause

func (v *player) GoPlayPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PlayPause", flags, ch)
}

func (v *player) PlayPause(flags dbus.Flags) error {
	return (<-v.GoPlayPause(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Previous

func (v *player) GoPrevious(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Previous", flags, ch)
}

func (v *player) Previous(flags dbus.Flags) error {
	return (<-v.GoPrevious(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Seek

func (v *player) GoSeek(flags dbus.Flags, ch chan *dbus.Call, Offset int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Seek", flags, ch, Offset)
}

func (v *player) Seek(flags dbus.Flags, Offset int64) error {
	return (<-v.GoSeek(flags, make(chan *dbus.Call, 1), Offset).Done).Err
}

// method SetPosition

func (v *player) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, TrackId dbus.ObjectPath, Position int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, TrackId, Position)
}

func (v *player) SetPosition(flags dbus.Flags, TrackId dbus.ObjectPath, Position int64) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), TrackId, Position).Done).Err
}

// method Stop

func (v *player) GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Stop", flags, ch)
}

func (v *player) Stop(flags dbus.Flags) error {
	return (<-v.GoStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Seeked

func (v *player) ConnectSeeked(cb func(Position int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Seeked", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Seeked",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var Position int64
		err := dbus.Store(sig.Body, &Position)
		if err == nil {
			cb(Position)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property CanControl b

func (v *player) CanControl() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanControl",
	}
}

// property CanGoNext b

func (v *player) CanGoNext() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanGoNext",
	}
}

// property CanGoPrevious b

func (v *player) CanGoPrevious() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanGoPrevious",
	}
}

// property CanPause b

func (v *player) CanPause() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanPause",
	}
}

// property CanPlay b

func (v *player) CanPlay() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanPlay",
	}
}

// property CanSeek b

func (v *player) CanSeek() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanSeek",
	}
}

// property Metadata a{sv}

func (v *player) Metadata() PropPlayerMetadata {
	return PropPlayerMetadata{
		Impl: v,
	}
}

type PropPlayerMetadata struct {
	Impl proxy.Implementer
}

func (p PropPlayerMetadata) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Metadata", &value)
	return
}

func (p PropPlayerMetadata) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]dbus.Variant
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
		"Metadata", cb0)
}

// property PlaybackStatus s

func (v *player) PlaybackStatus() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "PlaybackStatus",
	}
}

// property Position x

func (v *player) Position() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Position",
	}
}

// property Volume d

func (v *player) Volume() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Volume",
	}
}
