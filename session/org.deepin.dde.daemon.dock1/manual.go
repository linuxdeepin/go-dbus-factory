// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dock1

// FrontendWindowRect type
type FrontendWindowRect struct {
	X, Y          int32
	Width, Height uint32
}

// WindowInfo type
type WindowInfo struct {
	Title string
	Flash bool
}
