package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
混一色：由一种花色序数牌及字牌组成的和牌。
*/

const (
	_HUPAI49_ID     = 49
	_HUPAI49_NAME   = "混一色"
	_HUPAI49_FANSHU = 6
	_HUPAI49_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI49_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai49{
		id:             _HUPAI49_ID,
		name:           _HUPAI49_NAME,
		fanShu:         _HUPAI49_FANSHU,
		setChcFanShuID: _HUPAI49_CHECKID_,
		huKind:         _HUPAI49_KIND,
	})
}

type huPai49 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai49) GetID() int {
	return h.id
}

func (h *huPai49) Name() string {
	return h.name
}

func (h *huPai49) GetFanShu() int {
	return h.fanShu
}

func (h *huPai49) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai49) CheckSatisfySelf(method *cardType.HuMethod) bool {
	allIncCard := method.GetAllInclude()
	if allIncCard.GetColorCnt() != 1 {
		return false
	}
	for _, card := range allIncCard {
		if cardType.IsFengCard(card) {
			return true
		}
	}
	return false
}
