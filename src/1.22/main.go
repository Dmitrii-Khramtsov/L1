package main

import (
	"fmt"
	"math/big"
)

type BigFloat struct {
	x *big.Float
	y *big.Float
}

func newBigFloat(a, b string) *BigFloat {
	x, y := new(big.Float), new(big.Float)
	x.SetString(a)
	y.SetString(b)
	return &BigFloat{
		x: x,
		y: y,
	}
}

func (b *BigFloat)Sum() *big.Float {
	res := new(big.Float)
	res.Add(b.x, b.y)
	return res
}

func (b *BigFloat)Sub() *big.Float {
	res := new(big.Float)
	res.Sub(b.x, b.y)
	return res
}

func (b *BigFloat) Mul() *big.Float {
	res := new(big.Float)
	return res.Mul(b.x, b.y)
}

func (b *BigFloat) Div() *big.Float {
	res := new(big.Float)
	return res.Quo(b.x, b.y)
}

func main() {
	x, y := "123456789012345678901234567895", "987654321098765432109876543210"

	nums := newBigFloat(x, y)

	fmt.Println("Сумма:       ", nums.Sum())
	fmt.Println("Разность:    ", nums.Sub())
	fmt.Println("Произведение:", nums.Mul())
	fmt.Println("Частное:     ", nums.Div())
}
