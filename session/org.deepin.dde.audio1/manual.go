// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package audio1

type PortInfo struct {
	Name        string
	Description string
	Priority    uint32
	Available   int // 0: Unknown, 1: No, 2: Yes
}
