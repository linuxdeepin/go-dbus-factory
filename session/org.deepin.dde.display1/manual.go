package display1

// ModeInfo type
type ModeInfo struct {
	Id     uint32
	Width  uint16
	Height uint16
	Rate   float64
}

// Rectangle type
type Rectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

// Touchscreen type
type Touchscreen struct {
	Id         int32
	Node       string
	DeviceNode string
	Serial     string
}
