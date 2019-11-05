package structsmethodsintrerfaces

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height / 2.0
}
