package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
全双刻：由2、4、6、8序数牌的刻子、将牌组成的和牌。不计碰碰和、断幺。
*/

const (
	_HUPAI21_ID     = 21
	_HUPAI21_NAME   = "全双刻"
	_HUPAI21_FANSHU = 24
	_HUPAI21_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI21_CHECKID_ = []int{48, 68} //

func init() {
	fanShuMgr.registerHander(&huPai21{
		id:             _HUPAI21_ID,
		name:           _HUPAI21_NAME,
		fanShu:         _HUPAI21_FANSHU,
		setChcFanShuID: _HUPAI21_CHECKID_,
		huKind:         _HUPAI21_KIND,
	})
}

type huPai21 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai21) GetID() int {
	return h.id
}

func (h *huPai21) Name() string {
	return h.name
}

func (h *huPai21) GetFanShu() int {
	return h.fanShu
}

func (h *huPai21) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai21) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil || method.GetShunZi() != nil {
		return false
	}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetJiangCard())

	slChkNum := []int{2, 4, 6, 8}
	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			return false
		}
		if !common.InIntSlace(slChkNum, int(card%10)) {
			return false
		}
	}
	return true
}
