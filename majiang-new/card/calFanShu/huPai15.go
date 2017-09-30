package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
一色四节高：一种花色4副依次递增一位数的刻子，不计一色三同顺、碰碰和。
*/

const (
	_HUPAI15_ID     = 15
	_HUPAI15_NAME   = "一色四节高"
	_HUPAI15_FANSHU = 48
	_HUPAI15_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI15_CHECKID_ = []int{23, 24, 48} //

func init() {
	fanShuMgr.registerHander(&huPai15{
		id:             _HUPAI15_ID,
		name:           _HUPAI15_NAME,
		fanShu:         _HUPAI15_FANSHU,
		setChcFanShuID: _HUPAI15_CHECKID_,
		huKind:         _HUPAI15_KIND,
	})
}

type huPai15 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai15) GetID() int {
	return h.id
}

func (h *huPai15) Name() string {
	return h.name
}

func (h *huPai15) GetFanShu() int {
	return h.fanShu
}

func (h *huPai15) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
	if method.GetHuPaiKind() != h.huKind {
		return false, 0, satisfyedID, slBanID
	}

	if common.InIntSlace(satisfyedID, h.GetID()) {
		return false, 0, satisfyedID, slBanID
	}

	//不能计算的直接退出
	if common.InIntSlace(slBanID, h.GetID()) {
		return false, 0, satisfyedID, slBanID
	}

	if !h.CheckSatisfySelf(method) {
		slBanID = append(slBanID, h.GetID())
		return false, 0, satisfyedID, slBanID
	}
	//满足后把自己自己要ban的id加入进去
	for _, id := range h.setChcFanShuID {
		if !common.InIntSlace(slBanID, id) {
			slBanID = append(slBanID, id)
		}
	}

	fanShu := h.GetFanShu()
	satisfyedID = append(satisfyedID, h.GetID())
	//再把其他的所有的id全部遍历，有就加上去
	otherChkHander := fanShuMgr.getHanderExcept(append(satisfyedID, slBanID...))
	for _, hander := range otherChkHander {
		ok, tmpFanShu, tmpSatisfyID, slTmpBanID := hander.Satisfy(method, satisfyedID, slBanID)
		slBanID = slTmpBanID
		if ok {
			fanShu += tmpFanShu
			satisfyedID = tmpSatisfyID
		}
	}

	return true, fanShu, satisfyedID, slBanID
}

func (h *huPai15) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil {
		return false
	}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	if len(slThreePairs) != 4 {
		return false
	}

	minCard := cardType.FLOWER_START_CARD
	//找出最小的基准
	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			return false
		}
		if card < minCard {
			minCard = card
		}
	}
	var slThreeCard []uint8
	for _, card := range slThreePairs {
		slThreeCard = append(slThreeCard, card-minCard)
	}

	slChkNum := []uint8{0, 1, 2, 3}
	if cardType.GetSlCardDiff(slThreeCard, slChkNum) != nil {
		return false
	}
	return true
}
