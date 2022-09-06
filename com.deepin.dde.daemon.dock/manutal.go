// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dock

type FrontendWindowRect struct {
	X, Y          int32
	Width, Height uint32
}

type WindowInfo struct {
	Title string
	Flash bool
}
