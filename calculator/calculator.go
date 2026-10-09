package calculator

import "errors"

type Module struct {
	A float64
	B float64
}

func (m Module) Add() (float64, error) {
	if m.A < -1000 || m.A > 1000 || m.B < -1000 || m.B > 1000 {
		return 0.0, errors.New("Аргументы должны быть в пределах [-1000;1000]")
	}

	return m.A + m.B, nil
}

func (m Module) Subtraction() (float64, error) {
	if m.A < -1000 || m.A > 1000 || m.B < -1000 || m.B > 1000 {
		return 0.0, errors.New("Аргументы должны быть в пределах [-1000;1000]")
	}

	return m.A - m.B, nil

}

func (m Module) Multiplication() (float64, error) {
	if m.A < -1000 || m.A > 1000 || m.B < -1000 || m.B > 1000 {
		return 0.0, errors.New("Аргументы должны быть в пределах [-1000;1000]")
	}

	return m.A * m.B, nil
}

func (m Module) Division() (float64, error) {
	if m.A < -1000 || m.A > 1000 || m.B < -1000 || m.B > 1000 {
		return 0.0, errors.New("Аргументы должны быть в пределах [-1000;1000]")
	}

	if m.B == 0 {
		return 0.0, errors.New("Деление на 0")
	}

	return m.A / m.B, nil
}
