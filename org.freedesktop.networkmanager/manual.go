package networkmanager

type IP6Address struct {
	Address []byte
	Prefix  uint32
	Gateway []byte
}

type IP6Route struct {
	Route   []byte
	Prefix  uint32
	NextHop []byte
	Metric  uint32
}

type DeviceStateReason struct {
	State  uint32
	Reason uint32
}

func (obj *AccessPoint) AccessPoint() *accessPoint {
	return &obj.accessPoint
}
