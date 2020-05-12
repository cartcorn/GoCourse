package main

import (
	"fmt"
)

type CashNds struct {
	nds float64
}

func (cs CashNds) CashValue() float64 {
	return cs.nds * 0.13
}

type Cash struct {
	cash float64
}

func (c Cash) CashValue() float64 {
	return c.cash * c.cash / 2
}

type CashInt interface {
	CashValue() float64
}

func SummCash(cashint ...CashInt) float64 {
	res := 0.0
	for _, ci := range cashint {
		res += ci.CashValue()
	}
	return res
}
func main() {
	firstCash := Cash{10000}
	secondCash := Cash{15000}
	firstCashnd := CashNds{500}
	secondCashnd := CashNds{700}
	total := SummCash(firstCash, secondCash, firstCashnd, secondCashnd)
	fmt.Println(total)
}
