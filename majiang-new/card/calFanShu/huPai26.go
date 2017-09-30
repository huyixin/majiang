package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
全中：由序数牌456组成的顺子、刻子（杠）、将牌的和牌。不计断幺。
*/

const (
	_HUPAI26_ID     = 26
	_HUPAI26_NAME   = "全中"
	_HUPAI26_FANSHU = 24
	_HUPAI26_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI26_CHECKID_ = []int{68, 76} //

func init() {
	fanShuMgr.registerHander(&huPai26{
		id:             _HUPAI26_ID,
		name:           _HUPAI26_NAME,
		fanShu:         _HUPAI26_FANSHU,
		setChcFanShuID: _HUPAI26_CHECKID_,
		huKind:         _HUPAI26_KIND,
	})
}

type huPai26 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai26) GetID() int {
	return h.id
}

func (h *huPai26) Name() string {
	return h.name
}

func (h *huPai26) GetFanShu() int {
	return h.fanShu
}

func (h *huPai26) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai26) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetPengCard(), method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	handCard := method.GetAllCard()
	handCard = handCard.Push(slThreePairs)

	for _, slChiCard := range method.GetChiCard() {
		handCard = handCard.Push(slChiCard[:])
	}

	slChkNum := []uint8{4, 5, 6}
	for _, card := range handCard.GetAllKind() {
		if cardType.IsFengCard(card) {
			return false
		}
		if !common.InUInt8Slace(slChkNum, uint8(card%10)) {
			return false
		}
	}
	return true
}
