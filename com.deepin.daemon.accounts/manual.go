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
		Warn       int
		Expire     int
	}
	CurrentLogin            LoginUtmpx
	LastLogin               LoginUtmpx
	FailCountSinceLastLogin int
}
