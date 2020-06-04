package types

type Float struct {
	Float float64
}

func (z *Float) Add(x, y *Float) *Float {
	z = &Float{x.Float + y.Float}
	return z
}

func (z *Float) Sub(x, y *Float) *Float {
	z = &Float{x.Float - y.Float}
	return z
}

func (z *Float) Mul(x, y *Float) *Float {
	z = &Float{x.Float * y.Float}
	return z
}

func (z *Float) Div(x, y *Float) *Float {
	z = &Float{x.Float / y.Float}
	return z
}
