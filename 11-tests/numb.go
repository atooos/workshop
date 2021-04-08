package myint

type MyInt int32

func (i MyInt) Sub(n int) MyInt {
	return i - MyInt(n)
}

func (i MyInt) Add(n int) MyInt {
	return i + MyInt(n)
}

func (i MyInt) Multiply(n int) MyInt {
	return i * MyInt(n)
}

func (i MyInt) Divide(n int) MyInt {
	return i / MyInt(n)
}
