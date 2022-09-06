// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package networkmanager

type IP6Address struct {
	Address []byte
	Prefix  uint32
	Gateway []byte
}

type IP6Route struct {
	Route   []byte
	Prefix  uint32
	NextHop []byte
	Metric  uint32
}

type DeviceStateReason struct {
	State  uint32
	Reason uint32
}
