package policykit1

import "pkg.deepin.io/lib/dbus1"

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
	annotations      map[string]string
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

// Identity

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
