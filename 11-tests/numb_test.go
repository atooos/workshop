package myint

import (
	"testing"
)

func TestMyIntSub(t *testing.T) {
	data := []struct {
		title  string
		init   MyInt
		param  int
		output int
		err    error
	}{
		{"A", 1, 1, 0, nil},
		{"B", -2147483648, 2, 0, ErrTooLarge},
	}
	for _, d := range data {
		res, err := d.init.Sub(d.param)
		t.Logf("res of sub is %v", res)
		if res != MyInt(d.output) || d.err != err {
			t.Errorf("for test %v got %v but is waiting for %v", d.title, res, d.output)
		}
	}

}
