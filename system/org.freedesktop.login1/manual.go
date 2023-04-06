// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package login1

import "github.com/godbus/dbus/v5"

type ScheduledShutdown struct {
	Type string
	USec uint64
}

type SessionInfo struct {
	Id   string
	Path dbus.ObjectPath
}

type UserInfo struct {
	UID  uint32
	Path dbus.ObjectPath
}

type SeatInfo struct {
	Id   string
	Path dbus.ObjectPath
}

type SessionDetail struct {
	SessionId string
	UID       uint32
	UserName  string
	SeatId    string
	Path      dbus.ObjectPath
}

type UserDetail struct {
	UID  uint32
	Name string
	Path dbus.ObjectPath
}

type InhibitorInfo struct {
	What string
	Who  string
	Why  string
	Mode string
	UID  uint32
	PID  uint32
}
