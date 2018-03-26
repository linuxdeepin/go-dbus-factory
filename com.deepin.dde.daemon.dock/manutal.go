package libdock

type FrontendWindowRect struct {
	X, Y          int32
	Width, Height uint32
}

type WindowInfo struct {
	Title string
	Flash bool
}
