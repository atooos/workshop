package myint

import (
	"errors"
)

var ErrTooLarge = errors.New("values too large")

type MyInt int32

func (i MyInt) Sub(n int) (MyInt, error) {
	res := int64(i) - int64(n)
	if res < -2147483648 {
		return 0, ErrTooLarge
	}
	return MyInt(res), nil
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
