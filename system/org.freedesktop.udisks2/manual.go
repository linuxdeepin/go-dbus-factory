// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package udisks2

import "github.com/godbus/dbus/v5"

type Attribute struct {
	Id         byte
	Name       string
	Flags      uint16
	Value      int32
	Worst      int32
	Threshold  int32
	Pretty     int64
	PrettyUnit int32
	Expansion  map[string]dbus.Variant
}

type ConfigurationItem struct {
	Type    string
	Details map[string]dbus.Variant
}
