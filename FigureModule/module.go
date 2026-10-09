package figure

type Figure interface {
	Area() int
	Perimeter() int
}

type Module struct {
	figure Figure
}

func NewModule(figure Figure) Module {
	return Module{
		figure: figure,
	}
}

func (m Module) GetArea() int {
	return m.figure.Area()
}

func (m Module) GetPerimeter() int {
	return m.figure.Perimeter()
}
