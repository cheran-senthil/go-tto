package types

type Int64 struct {
	Int64 int64
}

func (z *Int64) Add(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 + y.Int64}
	return z
}

func (z *Int64) Sub(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 - y.Int64}
	return z
}

func (z *Int64) Mul(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 * y.Int64}
	return z
}

func (z *Int64) Div(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 / y.Int64}
	return z
}

func (z *Int64) Mod(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 % y.Int64}
	return z
}

func (z *Int64) And(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 & y.Int64}
	return z
}

func (z *Int64) Or(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 | y.Int64}
	return z
}

func (z *Int64) Xor(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 ^ y.Int64}
	return z
}

func (z *Int64) AndNot(x, y *Int64) *Int64 {
	z = &Int64{x.Int64 &^ y.Int64}
	return z
}

func (z *Int64) Lsh(x *Int64, n uint) *Int64 {
	z = &Int64{x.Int64 << n}
	return z
}

func (z *Int64) Rsh(x *Int64, n uint) *Int64 {
	z = &Int64{x.Int64 >> n}
	return z
}
