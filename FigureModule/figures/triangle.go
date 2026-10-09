package figures

type Triangle struct {
	height int
	width  int
	b      int
	c      int
}

func NewTriangle(height int, width int, b int, c int) Triangle {
	return Triangle{
		height: height,
		width:  width,
		b:      b,
		c:      c,
	}
}

func (t Triangle) Area() int {
	return t.height * t.width / 2
}

func (t Triangle) Perimeter() int {
	return t.b + t.width + t.c
}
