// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package display1

// ModeInfo type
type ModeInfo struct {
	Id     uint32
	Width  uint16
	Height uint16
	Rate   float64
}

// Rectangle type
type Rectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

// Touchscreen type
type Touchscreen struct {
	Id         int32
	Node       string
	DeviceNode string
	Serial     string
}

type TouchscreenV2 struct {
	Id         int32
	Node       string
	DeviceNode string
	Serial     string
	UUID       string
}
