package policykit1

import "github.com/godbus/dbus"

type ActionDescription struct {
	ActionId         string
	Description      string
	Message          string
	VendorName       string
	VendorURL        string
	IconName         string
	ImplicitAny      ImplicitAuthorization
	ImplicitInactive ImplicitAuthorization
	ImplicitActive   ImplicitAuthorization
}

type ImplicitAuthorization uint32

type Subject struct {
	Kind    string
	Details map[string]dbus.Variant
}

type Identity struct {
	Kind    string
	Details map[string]dbus.Variant
}

type AuthorizationResult struct {
	IsAuthorized bool
	IsChallenge  bool
	Details      map[string]dbus.Variant
}

type TemporaryAuthorization struct {
	Id           string
	ActionId     string
	Subject      Subject
	TimeObtained uint64
	TimeExpires  uint64
}

// SubjectKind
const (
	SubjectKindUnixProcess   = "unix-process"
	SubjectKindUnixSession   = "unix-session"
	SubjectKindSystemBusName = "system-bus-name"
)

func MakeSubject(kind string) Subject {
	return Subject{
		Kind:    kind,
		Details: make(map[string]dbus.Variant),
	}
}

func (s *Subject) SetDetail(key string, value interface{}) {
	s.Details[key] = dbus.MakeVariant(value)
}

// CheckAuthorizationFlags
const (
	CheckAuthorizationFlagsNone                 = 0
	CheckAuthorizationFlagsAllowUserInteraction = 1
)
