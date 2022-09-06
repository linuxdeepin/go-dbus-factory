// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package accounts

type LoginUtmpx struct {
	InittabID string
	Line      string
	Host      string
	Address   string
	Time      string
}

type LoginReminderInfo struct {
	Username string
	Spent    struct {
		LastChange int
		Min        int
		Max        int
		Warn       int
		Inactive   int
		Expire     int
	}
	CurrentLogin            LoginUtmpx
	LastLogin               LoginUtmpx
	FailCountSinceLastLogin int
}
