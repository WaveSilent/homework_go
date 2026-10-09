package figures

type Rectangle struct {
	height int
	width  int
}

func NewRectangle(height int, width int) Rectangle {
	return Rectangle{
		height: height,
		width:  width,
	}
}

func (r Rectangle) Area() int {
	return r.height * r.width
}

func (r Rectangle) Perimeter() int {
	return (r.height + r.width) * 2
}
