package common

import (
	"math"
)

//保留后面几位小数函数, 不是四舍五入
func GetFloatByLimitDecimal(fNum float64, floatSite float64) float64 {
	powNum := math.Pow(float64(10), floatSite)
	return float64(int(powNum*fNum)) / powNum
}
