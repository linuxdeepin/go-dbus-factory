package lunarcalendar

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type LunarCalendar struct {
	lunarCalendar // interface com.deepin.api.LunarCalendar
	proxy.Object
}

func NewLunarCalendar(conn *dbus.Conn) *LunarCalendar {
	obj := new(LunarCalendar)
	obj.Object.Init_(conn, "com.deepin.api.LunarCalendar", "/com/deepin/api/LunarCalendar")
	return obj
}

type lunarCalendar struct{}

func (v *lunarCalendar) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*lunarCalendar) GetInterfaceName_() string {
	return "com.deepin.api.LunarCalendar"
}

// method GetFestivalMonth

func (v *lunarCalendar) GoGetFestivalMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetFestivalMonth", flags, ch, year, month)
}

func (v *lunarCalendar) GoGetFestivalMonthWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, year int32, month int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetFestivalMonth", flags, ch, year, month)
}

func (*lunarCalendar) StoreGetFestivalMonth(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *lunarCalendar) GetFestivalMonth(flags dbus.Flags, year int32, month int32) (json string, err error) {
	return v.StoreGetFestivalMonth(
		<-v.GoGetFestivalMonth(flags, make(chan *dbus.Call, 1), year, month).Done)
}

func (v *lunarCalendar) GetFestivalMonthWithTimeout(timeout time.Duration, flags dbus.Flags, year int32, month int32) (json string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetFestivalMonthWithContext(ctx, flags, make(chan *dbus.Call, 1), year, month).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetFestivalMonth(call)
}

// method GetFestivalsInRange

func (v *lunarCalendar) GoGetFestivalsInRange(flags dbus.Flags, ch chan *dbus.Call, startDate string, endDate string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetFestivalsInRange", flags, ch, startDate, endDate)
}

func (v *lunarCalendar) GoGetFestivalsInRangeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, startDate string, endDate string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetFestivalsInRange", flags, ch, startDate, endDate)
}

func (*lunarCalendar) StoreGetFestivalsInRange(call *dbus.Call) (result []DayFestival, err error) {
	err = call.Store(&result)
	return
}

func (v *lunarCalendar) GetFestivalsInRange(flags dbus.Flags, startDate string, endDate string) (result []DayFestival, err error) {
	return v.StoreGetFestivalsInRange(
		<-v.GoGetFestivalsInRange(flags, make(chan *dbus.Call, 1), startDate, endDate).Done)
}

func (v *lunarCalendar) GetFestivalsInRangeWithTimeout(timeout time.Duration, flags dbus.Flags, startDate string, endDate string) (result []DayFestival, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetFestivalsInRangeWithContext(ctx, flags, make(chan *dbus.Call, 1), startDate, endDate).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetFestivalsInRange(call)
}

// method GetHuangLiDay

func (v *lunarCalendar) GoGetHuangLiDay(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, day int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetHuangLiDay", flags, ch, year, month, day)
}

func (v *lunarCalendar) GoGetHuangLiDayWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, day int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetHuangLiDay", flags, ch, year, month, day)
}

func (*lunarCalendar) StoreGetHuangLiDay(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *lunarCalendar) GetHuangLiDay(flags dbus.Flags, year int32, month int32, day int32) (json string, err error) {
	return v.StoreGetHuangLiDay(
		<-v.GoGetHuangLiDay(flags, make(chan *dbus.Call, 1), year, month, day).Done)
}

func (v *lunarCalendar) GetHuangLiDayWithTimeout(timeout time.Duration, flags dbus.Flags, year int32, month int32, day int32) (json string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetHuangLiDayWithContext(ctx, flags, make(chan *dbus.Call, 1), year, month, day).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetHuangLiDay(call)
}

// method GetHuangLiMonth

func (v *lunarCalendar) GoGetHuangLiMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, fill bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetHuangLiMonth", flags, ch, year, month, fill)
}

func (v *lunarCalendar) GoGetHuangLiMonthWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, fill bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetHuangLiMonth", flags, ch, year, month, fill)
}

func (*lunarCalendar) StoreGetHuangLiMonth(call *dbus.Call) (json string, err error) {
	err = call.Store(&json)
	return
}

func (v *lunarCalendar) GetHuangLiMonth(flags dbus.Flags, year int32, month int32, fill bool) (json string, err error) {
	return v.StoreGetHuangLiMonth(
		<-v.GoGetHuangLiMonth(flags, make(chan *dbus.Call, 1), year, month, fill).Done)
}

func (v *lunarCalendar) GetHuangLiMonthWithTimeout(timeout time.Duration, flags dbus.Flags, year int32, month int32, fill bool) (json string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetHuangLiMonthWithContext(ctx, flags, make(chan *dbus.Call, 1), year, month, fill).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetHuangLiMonth(call)
}
