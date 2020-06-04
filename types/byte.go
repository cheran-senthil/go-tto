package types

type Byte struct {
	Byte byte
}

func (z *Byte) Add(x, y *Byte) *Byte {
	z = &Byte{x.Byte + y.Byte}
	return z
}

func (z *Byte) Sub(x, y *Byte) *Byte {
	z = &Byte{x.Byte - y.Byte}
	return z
}

func (z *Byte) Mul(x, y *Byte) *Byte {
	z = &Byte{x.Byte * y.Byte}
	return z
}

func (z *Byte) Div(x, y *Byte) *Byte {
	z = &Byte{x.Byte / y.Byte}
	return z
}

func (z *Byte) Mod(x, y *Byte) *Byte {
	z = &Byte{x.Byte % y.Byte}
	return z
}

func (z *Byte) And(x, y *Byte) *Byte {
	z = &Byte{x.Byte & y.Byte}
	return z
}

func (z *Byte) Or(x, y *Byte) *Byte {
	z = &Byte{x.Byte | y.Byte}
	return z
}

func (z *Byte) Xor(x, y *Byte) *Byte {
	z = &Byte{x.Byte ^ y.Byte}
	return z
}

func (z *Byte) AndNot(x, y *Byte) *Byte {
	z = &Byte{x.Byte &^ y.Byte}
	return z
}

func (z *Byte) Lsh(x *Byte, n uint) *Byte {
	z = &Byte{x.Byte << n}
	return z
}

func (z *Byte) Rsh(x *Byte, n uint) *Byte {
	z = &Byte{x.Byte >> n}
	return z
}
