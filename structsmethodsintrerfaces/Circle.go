package structsmethodsintrerfaces

import "math"

type Circle struct {
	Redius float64
}

func (circle Circle) Area() float64 {
	return circle.Redius * circle.Redius * math.Pi
}
