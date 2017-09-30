package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
大于五：由序数牌6-9的顺子、刻子、将牌组成的和牌。不计无字。
*/

const (
	_HUPAI36_ID     = 36
	_HUPAI36_NAME   = "大于五"
	_HUPAI36_FANSHU = 12
	_HUPAI36_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI36_CHECKID_ = []int{76} //

func init() {
	fanShuMgr.registerHander(&huPai36{
		id:             _HUPAI36_ID,
		name:           _HUPAI36_NAME,
		fanShu:         _HUPAI36_FANSHU,
		setChcFanShuID: _HUPAI36_CHECKID_,
		huKind:         _HUPAI36_KIND,
	})
}

type huPai36 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai36) GetID() int {
	return h.id
}

func (h *huPai36) Name() string {
	return h.name
}

func (h *huPai36) GetFanShu() int {
	return h.fanShu
}

func (h *huPai36) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai36) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetPengCard(), method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	handKind := method.GetAllCard()
	for _, slChiCard := range method.GetChiCard() {
		handKind = handKind.Push(slChiCard[:])
	}
	for _, card := range handKind.GetAllCard() {
		if cardType.IsFengCard(card) {
			return false
		}
		if card%10 < 6 {
			return false
		}
	}
	return true
}
