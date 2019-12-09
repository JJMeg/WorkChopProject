package math_cal

import "math"

// 保留三位小数
func floatRound(flow float64) float64 {
	precision := math.Pow(10, float64(3))
	input := int(flow*precision + math.Copysign(0.5, flow*precision))
	return float64(input) / precision

}
