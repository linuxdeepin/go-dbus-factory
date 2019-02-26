package soundthemeplayer

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

type SoundThemePlayer struct {
	soundThemePlayer // interface com.deepin.api.SoundThemePlayer
	proxy.Object
}

func NewSoundThemePlayer(conn *dbus.Conn) *SoundThemePlayer {
	obj := new(SoundThemePlayer)
	obj.Object.Init_(conn, "com.deepin.api.SoundThemePlayer", "/com/deepin/api/SoundThemePlayer")
	return obj
}

type soundThemePlayer struct{}

func (v *soundThemePlayer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*soundThemePlayer) GetInterfaceName_() string {
	return "com.deepin.api.SoundThemePlayer"
}

// method Play

func (v *soundThemePlayer) GoPlay(flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Play", flags, ch, theme, event, device)
}

func (v *soundThemePlayer) Play(flags dbus.Flags, theme string, event string, device string) error {
	return (<-v.GoPlay(flags, make(chan *dbus.Call, 1), theme, event, device).Done).Err
}
