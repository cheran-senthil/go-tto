package types

type Int struct {
	Int int
}

func (z *Int) Add(x, y *Int) *Int {
	z = &Int{x.Int + y.Int}
	return z
}

func (z *Int) Sub(x, y *Int) *Int {
	z = &Int{x.Int - y.Int}
	return z
}

func (z *Int) Mul(x, y *Int) *Int {
	z = &Int{x.Int * y.Int}
	return z
}

func (z *Int) Div(x, y *Int) *Int {
	z = &Int{x.Int / y.Int}
	return z
}

func (z *Int) Mod(x, y *Int) *Int {
	z = &Int{x.Int % y.Int}
	return z
}

func (z *Int) And(x, y *Int) *Int {
	z = &Int{x.Int & y.Int}
	return z
}

func (z *Int) Or(x, y *Int) *Int {
	z = &Int{x.Int | y.Int}
	return z
}

func (z *Int) Xor(x, y *Int) *Int {
	z = &Int{x.Int ^ y.Int}
	return z
}

func (z *Int) AndNot(x, y *Int) *Int {
	z = &Int{x.Int &^ y.Int}
	return z
}

func (z *Int) Lsh(x *Int, n uint) *Int {
	z = &Int{x.Int << n}
	return z
}

func (z *Int) Rsh(x *Int, n uint) *Int {
	z = &Int{x.Int >> n}
	return z
}
