package types

type String struct {
	String string
}

func (z *String) Add(x, y *String) *String {
	z = &String{x.String + y.String}
	return z
}
