// Code generated by "./generator ./com.deepin.api.lunarcalendar"; DO NOT EDIT.

package lunarcalendar

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type LunarCalendar interface {
	lunarCalendar // interface com.deepin.api.LunarCalendar
	proxy.Object
}

type objectLunarCalendar struct {
	interfaceLunarCalendar // interface com.deepin.api.LunarCalendar
	proxy.ImplObject
}

func NewLunarCalendar(conn *dbus.Conn) LunarCalendar {
	obj := new(objectLunarCalendar)
	obj.ImplObject.Init_(conn, "com.deepin.api.LunarCalendar", "/com/deepin/api/LunarCalendar")
	return obj
}

type lunarCalendar interface {
	GoGetFestivalMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32) *dbus.Call
	GetFestivalMonth(flags dbus.Flags, year int32, month int32) (string, error)
	GoGetFestivalsInRange(flags dbus.Flags, ch chan *dbus.Call, startDate string, endDate string) *dbus.Call
	GetFestivalsInRange(flags dbus.Flags, startDate string, endDate string) ([]DayFestival, error)
	GoGetHuangLiDay(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, day int32) *dbus.Call
	GetHuangLiDay(flags dbus.Flags, year int32, month int32, day int32) (string, error)
	GoGetHuangLiMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, fill bool) *dbus.Call
	GetHuangLiMonth(flags dbus.Flags, year int32, month int32, fill bool) (string, error)
}

type interfaceLunarCalendar struct{}

func (v *interfaceLunarCalendar) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLunarCalendar) GetInterfaceName_() string {
	return "com.deepin.api.LunarCalendar"
}

// method GetFestivalMonth

func (v *interfaceLunarCalendar) GoGetFestivalMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetFestivalMonth", flags, ch, year, month)
}

func (*interfaceLunarCalendar) StoreGetFestivalMonth(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *interfaceLunarCalendar) GetFestivalMonth(flags dbus.Flags, year int32, month int32) (string, error) {
	return v.StoreGetFestivalMonth(
		<-v.GoGetFestivalMonth(flags, make(chan *dbus.Call, 1), year, month).Done)
}

// method GetFestivalsInRange

func (v *interfaceLunarCalendar) GoGetFestivalsInRange(flags dbus.Flags, ch chan *dbus.Call, startDate string, endDate string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetFestivalsInRange", flags, ch, startDate, endDate)
}

func (*interfaceLunarCalendar) StoreGetFestivalsInRange(call *dbus.Call) (result []DayFestival, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceLunarCalendar) GetFestivalsInRange(flags dbus.Flags, startDate string, endDate string) ([]DayFestival, error) {
	return v.StoreGetFestivalsInRange(
		<-v.GoGetFestivalsInRange(flags, make(chan *dbus.Call, 1), startDate, endDate).Done)
}

// method GetHuangLiDay

func (v *interfaceLunarCalendar) GoGetHuangLiDay(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, day int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetHuangLiDay", flags, ch, year, month, day)
}

func (*interfaceLunarCalendar) StoreGetHuangLiDay(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *interfaceLunarCalendar) GetHuangLiDay(flags dbus.Flags, year int32, month int32, day int32) (string, error) {
	return v.StoreGetHuangLiDay(
		<-v.GoGetHuangLiDay(flags, make(chan *dbus.Call, 1), year, month, day).Done)
}

// method GetHuangLiMonth

func (v *interfaceLunarCalendar) GoGetHuangLiMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, fill bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetHuangLiMonth", flags, ch, year, month, fill)
}

func (*interfaceLunarCalendar) StoreGetHuangLiMonth(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *interfaceLunarCalendar) GetHuangLiMonth(flags dbus.Flags, year int32, month int32, fill bool) (string, error) {
	return v.StoreGetHuangLiMonth(
		<-v.GoGetHuangLiMonth(flags, make(chan *dbus.Call, 1), year, month, fill).Done)
}
