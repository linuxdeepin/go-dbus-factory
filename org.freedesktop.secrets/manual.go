package secrets

import "pkg.deepin.io/lib/dbus1"

type Secret struct {
	Session     dbus.ObjectPath
	Parameters  []byte
	Value       []byte
	ContentType string
}
