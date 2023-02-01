package Math

type Square struct {
	AC int
	BD int
}

func GetArea(ac int, bc int) float32 {
	return float32(ac * bc)
}

func (s *Square) getArea() int {
	return s.BD * s.AC
}

func (s *Square) Info() string {
	if s.AC == s.BD {
		return "Square"
	}
	return "Rectangle"
}
