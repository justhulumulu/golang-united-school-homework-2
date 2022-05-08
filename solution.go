package square

import "math"

type sides int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	switch sidesNum {
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3.0) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	}
	return 0
}
