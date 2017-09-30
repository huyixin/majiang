package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
清一色：由一种花色的序数牌组成的和牌。不计无字。
*/

const (
	_HUPAI22_ID     = 22
	_HUPAI22_NAME   = "清一色"
	_HUPAI22_FANSHU = 24
	_HUPAI22_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI22_CHECKID_ = []int{76} //

func init() {
	fanShuMgr.registerHander(&huPai22{
		id:             _HUPAI22_ID,
		name:           _HUPAI22_NAME,
		fanShu:         _HUPAI22_FANSHU,
		setChcFanShuID: _HUPAI22_CHECKID_,
		huKind:         _HUPAI22_KIND,
	})
}

type huPai22 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai22) GetID() int {
	return h.id
}

func (h *huPai22) Name() string {
	return h.name
}

func (h *huPai22) GetFanShu() int {
	return h.fanShu
}

func (h *huPai22) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai22) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slAllKind := method.GetAllCardKind()
	for _, card := range slAllKind {
		if cardType.IsFengCard(card) {
			return false
		}
	}

	return (cardType.OwnerCardType(slAllKind)).GetColorCnt() == 1
}
