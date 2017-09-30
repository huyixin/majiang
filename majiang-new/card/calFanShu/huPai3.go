package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
绿一色：由23468条及发字中的任何牌组成的和牌。按新规定，无“发”时可计清一色、断幺，有“发”时可计混一色。
*/

const (
	_HUPAI3_ID     = 3
	_HUPAI3_NAME   = "绿一色"
	_HUPAI3_FANSHU = 88
	_HUPAI3_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI3_CHECKID_ = []int{49} //

func init() {
	fanShuMgr.registerHander(&huPai3{
		id:             _HUPAI3_ID,
		name:           _HUPAI3_NAME,
		fanShu:         _HUPAI3_FANSHU,
		setChcFanShuID: _HUPAI3_CHECKID_,
		huKind:         _HUPAI3_KIND,
	})
}

type huPai3 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai3) GetID() int {
	return h.id
}

func (h *huPai3) Name() string {
	return h.name
}

func (h *huPai3) GetFanShu() int {
	return h.fanShu
}

func (h *huPai3) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai3) CheckSatisfySelf(method *cardType.HuMethod) bool {
	allKind := method.GetAllCardKind()
	slOnlyKind := []uint8{22, 23, 24, 26, 28, 36}
	if cardType.GetSlCardDiff(allKind, slOnlyKind) != nil {
		return false
	}
	return true
}
