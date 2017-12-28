// https://github.com/samuell/gormula/blob/master/main.go
package main

import (
	"fmt"
)

func Seq(start float64, step float64, end float64) valstream {
	res := make(valstream)
	go func() {
		defer close(res)
		val := start
		for (end + step - val) > 0.0001 { // Same as val <= end, but take care of propagating float errors
			res <- val
			val = val + step
		}
	}()
	return res
}

func main() {
	// from 1k to 10k  USD
	borrowedAmount := Seq(1000.0, 1000.0, 10000.0)
	rate := 2.0    // p -> percent
	months := 24.0 // N

	// The formula
	monthlyPaymentUSD := Mul(
		Div(
			Mul( // p/1200
				Div(Val(rate),
					Val(1200.0),
				),
				Exp( // (1 + p/1200)**N
					Add(
						Val(1.0),
						Div(
							Val(rate),
							Val(1200.0)),
					),
					Val(months)),
			),
			Sub(
				Exp( // (1 + p/1200)**N
					Add(
						Val(1.0),
						Div(
							Val(rate),
							Val(1200.0)),
					),
					Val(months),
				),
				Val(1.0)),
		),
		borrowedAmount,
	)

	// Print out all the resulting monthly payments:
	borrowedAmountForPrint := Seq(1000.0, 100.0, 10000.0)
	for monthPay := range monthlyPaymentUSD {
		borrowed := <-borrowedAmountForPrint
		fmt.Printf("Monthly payment for 24 months, when borrowing %.2f USD: %.2f USD\n", borrowed, monthPay)
	}
}
