package calFanShu

import (
	"majiang-new/card/cardType"
)

//根据吃，碰，暗杠，明杠，吃牌来确定番数
func GetFanShu(slMethod []*cardType.HuMethod) (int, []int) {
	return fanShuMgr.getFanShu(slMethod)
}
