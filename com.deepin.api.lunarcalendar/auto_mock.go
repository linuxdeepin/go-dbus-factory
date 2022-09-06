// Code generated by "./generator ./com.deepin.api.lunarcalendar"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package lunarcalendar

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockLunarCalendar struct {
	MockInterfaceLunarCalendar // interface com.deepin.api.LunarCalendar
	proxy.MockObject
}

type MockInterfaceLunarCalendar struct {
	mock.Mock
}

// method GetFestivalMonth

func (v *MockInterfaceLunarCalendar) GoGetFestivalMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, year, month)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLunarCalendar) GetFestivalMonth(flags dbus.Flags, year int32, month int32) (string, error) {
	mockArgs := v.Called(flags, year, month)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetFestivalsInRange

func (v *MockInterfaceLunarCalendar) GoGetFestivalsInRange(flags dbus.Flags, ch chan *dbus.Call, startDate string, endDate string) *dbus.Call {
	mockArgs := v.Called(flags, ch, startDate, endDate)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLunarCalendar) GetFestivalsInRange(flags dbus.Flags, startDate string, endDate string) ([]DayFestival, error) {
	mockArgs := v.Called(flags, startDate, endDate)

	ret0, ok := mockArgs.Get(0).([]DayFestival)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetHuangLiDay

func (v *MockInterfaceLunarCalendar) GoGetHuangLiDay(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, day int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, year, month, day)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLunarCalendar) GetHuangLiDay(flags dbus.Flags, year int32, month int32, day int32) (string, error) {
	mockArgs := v.Called(flags, year, month, day)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetHuangLiMonth

func (v *MockInterfaceLunarCalendar) GoGetHuangLiMonth(flags dbus.Flags, ch chan *dbus.Call, year int32, month int32, fill bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, year, month, fill)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceLunarCalendar) GetHuangLiMonth(flags dbus.Flags, year int32, month int32, fill bool) (string, error) {
	mockArgs := v.Called(flags, year, month, fill)

	return mockArgs.String(0), mockArgs.Error(1)
}
