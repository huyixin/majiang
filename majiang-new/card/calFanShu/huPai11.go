package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
字一色：由字牌组成的和牌。不计碰碰和，可计七对。
*/

const (
	_HUPAI11_ID     = 11
	_HUPAI11_NAME   = "字一色"
	_HUPAI11_FANSHU = 64
	_HUPAI11_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI11_CHECKID_ = []int{48, 55} //

func init() {
	fanShuMgr.registerHander(&huPai11{
		id:             _HUPAI11_ID,
		name:           _HUPAI11_NAME,
		fanShu:         _HUPAI11_FANSHU,
		setChcFanShuID: _HUPAI11_CHECKID_,
		huKind:         _HUPAI11_KIND,
	})
}

type huPai11 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai11) GetID() int {
	return h.id
}

func (h *huPai11) Name() string {
	return h.name
}

func (h *huPai11) GetFanShu() int {
	return h.fanShu
}

func (h *huPai11) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai11) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil {
		return false
	}
	allKind := method.GetAllCardKind()
	slCheckKind := []uint8{31, 32, 33, 34, 35, 36, 37}
	if cardType.GetSlCardDiff(allKind, slCheckKind) != nil {
		return false
	}
	return true
}
