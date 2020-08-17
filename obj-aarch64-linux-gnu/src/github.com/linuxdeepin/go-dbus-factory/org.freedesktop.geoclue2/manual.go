package geoclue2

type Timestamp struct {
	Seconds      uint64
	Microseconds uint64
}

const (
	AccuracyLevelNone         = 0
	AccuracyLevelCountry      = 1
	AccuracyLevelCity         = 4
	AccuracyLevelNeighborhood = 5
	AccuracyLevelStreet       = 6
	AccuracyLevelExact        = 8
)
