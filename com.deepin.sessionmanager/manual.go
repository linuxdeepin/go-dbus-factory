package sessionmanager

type ModeInfo struct {
	Id     uint32
	Width  uint16
	Height uint16
	Rate   float64
}

type Rectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}
