package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle = 0
	SidesTriangle = 3
	SidesSquare = 4
)


func CalcSquare(sideLen float64, sidesNum uint8) float64 {
	var square float64
	switch sidesNum {
	case 0:
		square = math.Pi * sideLen * sideLen
	case 3:
		square = sideLen * sideLen * math.Sqrt(3) / 4
	case 4:
		square = sideLen * sideLen
	}
	return square
}
