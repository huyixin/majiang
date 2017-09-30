package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
双箭刻：2副箭刻（或杠）。
*/

const (
	_HUPAI54_ID     = 54
	_HUPAI54_NAME   = "双箭刻"
	_HUPAI54_FANSHU = 6
	_HUPAI54_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI54_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai54{
		id:             _HUPAI54_ID,
		name:           _HUPAI54_NAME,
		fanShu:         _HUPAI54_FANSHU,
		setChcFanShuID: _HUPAI54_CHECKID_,
		huKind:         _HUPAI54_KIND,
	})
}

type huPai54 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai54) GetID() int {
	return h.id
}

func (h *huPai54) Name() string {
	return h.name
}

func (h *huPai54) GetFanShu() int {
	return h.fanShu
}

func (h *huPai54) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai54) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := method.GetAllKeZi()
	var slJianKe []uint8
	for _, card := range slThreePairs {
		if cardType.IsJianCard(card) {
			slJianKe = append(slJianKe, card)
		}
	}
	if len(slJianKe) != 2 {
		return false
	}
	return true
}
