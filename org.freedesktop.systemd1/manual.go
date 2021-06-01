package systemd1

import (
	"github.com/godbus/dbus"
)

type Property struct {
	Name  string
	Value dbus.Variant
}

type PropertyCollection struct {
	Name       string
	Properties []Property
}

type UnitInfo struct {
	Name        string
	Description string
	LoadState   string // i.e. whether the unit file has been loaded successfully
	ActiveState string // i.e. whether the unit is currently started or not
	SubState    string // a more fine-grained version of the active state that is specific to the unit type, which the active state is not
	FollowUnit  string // A unit that is being followed in its state by this unit, if there is any, otherwise the empty string.
	Path        dbus.ObjectPath
	JobID       uint32 // If there is a job queued for the job unit the numeric job id, 0 otherwise
	JobType     string
	JobPath     dbus.ObjectPath
}

type JobInfo struct {
	ID       uint32
	Name     string
	Type     string
	State    string
	Path     dbus.ObjectPath
	UnitPath dbus.ObjectPath
}

type JobIdPath struct {
	Id   uint32
	Path dbus.ObjectPath
}

type UnitFile struct {
	Path  string
	State string
}

type UnitFileChange struct {
	Type        string // the type of the change (one of symlink or unlink)
	Filename    string // the file name of the symlink
	Destination string // the destination of the symlink
}

type DynamicUser struct {
	ID   uint32
	Name string
}

type UnitProcess struct {
	ControlGroup string
	PID          uint32
	Command      string
}

type Condition struct {
	Type       string
	IsTrigger  bool
	IsReversed bool
	Right      string

	// The status can be 0, in which case the condition hasn't been checked yet, a positive value, in which case the condition passed, or a negative value, in which case the condition failed. Currently only 0, +1, and -1 are used, but additional values may be used in the future, retaining the meaning of zero/positive/negative values.
	Status int32
}

type Assert Condition

type LoadError struct {
	Type    string
	Message string
}

type ExitStatus struct {
	A []int32
	B []int32
}

type ExecInfo struct {
	Path                string
	Arguments           []string
	ExitFailure         bool
	BeginClockRealtime  uint64
	BeginClockMonotonic uint64
	EndClockRealtime    uint64
	EndClockMonotonic   uint64
	PID                 uint32
	ExistCode           int32
	Status              int32
}

type IOParam struct {
	Name  string
	Value uint64
}

type DeviceAllowItem struct {
	Name  string
	Value string
}

type IPAddressAllowItem struct {
	I  uint32
	AV []dbus.Variant
	U  uint32
}

type IPAddressDenyItem IPAddressAllowItem

type EnvironmentFile struct {
	S string
	B bool
}

type BS struct {
	B bool
	S string
}

type SystemCallFilter struct {
	B  bool
	AS string
}

type RestrictAddressFamilies struct {
	B  bool
	AS string
}

type BindPath struct {
	S1 string
	S2 string
	B  bool
	T  uint64
}

type TemporaryFileSystemItem struct {
	S1 string
	S2 string
}
