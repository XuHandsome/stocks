package stocksUntil

import "fmt"

func Calculator(nowPrice float64, holdPrice float64, HoldNumber int) string {
	profit := float64(HoldNumber) * (nowPrice - holdPrice)
	prefix := ""
	if profit >= 0 {
		prefix = "+"
	}
	return prefix + fmt.Sprintf("%.3f", profit)
}
