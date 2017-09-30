package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
箭刻：由中、发、白3张相同的牌组成的刻子。
*/

const (
	_HUPAI59_ID     = 59
	_HUPAI59_NAME   = "箭刻"
	_HUPAI59_FANSHU = 2
	_HUPAI59_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI59_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai59{
		id:             _HUPAI59_ID,
		name:           _HUPAI59_NAME,
		fanShu:         _HUPAI59_FANSHU,
		setChcFanShuID: _HUPAI59_CHECKID_,
		huKind:         _HUPAI59_KIND,
	})
}

type huPai59 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai59) GetID() int {
	return h.id
}

func (h *huPai59) Name() string {
	return h.name
}

func (h *huPai59) GetFanShu() int {
	return h.fanShu
}

func (h *huPai59) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai59) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	var slJianKe []uint8
	for _, card := range slThreePairs {
		if cardType.IsJianCard(card) {
			slJianKe = append(slJianKe, card)
		}
	}
	if len(slJianKe) != 1 {
		return false
	}
	return true
}
