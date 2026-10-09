package figures

type Square struct {
	side int
}

func NewSquare(side int) Square {
	return Square{
		side: side,
	}
}

func (s Square) Area() int {
	return s.side * s.side
}

func (s Square) Perimeter() int {
	return s.side * 4
}
