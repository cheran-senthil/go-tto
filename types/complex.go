package types

type Complex struct {
	Complex complex128
}

func (z *Complex) Add(x, y *Complex) *Complex {
	z = &Complex{x.Complex + y.Complex}
	return z
}

func (z *Complex) Sub(x, y *Complex) *Complex {
	z = &Complex{x.Complex - y.Complex}
	return z
}

func (z *Complex) Mul(x, y *Complex) *Complex {
	z = &Complex{x.Complex * y.Complex}
	return z
}

func (z *Complex) Div(x, y *Complex) *Complex {
	z = &Complex{x.Complex / y.Complex}
	return z
}
