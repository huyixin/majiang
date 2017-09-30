package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
大三元：和牌中，有中发白3副刻子。不计双箭刻、箭刻。
*/

const (
	_HUPAI2_ID     = 2
	_HUPAI2_NAME   = "大三元"
	_HUPAI2_FANSHU = 88
	_HUPAI2_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI2_CHECKID_ = []int{54, 59} //不计箭刻

func init() {
	fanShuMgr.registerHander(&huPai2{
		id:             _HUPAI2_ID,
		name:           _HUPAI2_NAME,
		fanShu:         _HUPAI2_FANSHU,
		setChcFanShuID: _HUPAI2_CHECKID_,
		huKind:         _HUPAI2_KIND,
	})
}

type huPai2 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai2) GetID() int {
	return h.id
}

func (h *huPai2) Name() string {
	return h.name
}

func (h *huPai2) GetFanShu() int {
	return h.fanShu
}

func (h *huPai2) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai2) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slChkThreePairs := []uint8{35, 36, 37}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	if len(slThreePairs) == 0 {
		return false
	}

	if cardType.GetSlCardDiff(slChkThreePairs, slThreePairs) != nil {
		return false
	}
	return true
}
