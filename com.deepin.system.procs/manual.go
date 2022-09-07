// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package procs

type ProcMessage struct {
	ExecPath   string
	CGroupPath string
	Pid        string
	PPid       string
}
