package udisks2

import "pkg.deepin.io/lib/dbus1"

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
