package audio

type PortInfo struct {
	Name        string
	Description string
	Priority    uint32
	Available   int // 0: Unknown, 1: No, 2: Yes
}
