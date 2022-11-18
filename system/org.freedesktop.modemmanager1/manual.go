package modemmanager1

type ModemPort struct {
	Name string
	Type uint32
}

type ModemSignalQuality struct {
	Quality       uint32
	RecentlyTaken bool
}

type ModemModes struct {
	AllowedModes  uint32
	PreferredMode uint32
}

type OmaSession struct {
	Type uint32
	Id   uint32
}
