package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
大四喜：由4副风刻（杠）组成的和牌。不计圈风刻、门风刻、三风刻、碰碰和、幺九刻。
*/

const (
	_HUPAI1_ID     = 1
	_HUPAI1_NAME   = "大四喜"
	_HUPAI1_FANSHU = 88
	_HUPAI1_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI1_CHECKID_ = []int{60, 61, 38, 48, 73}

func init() {
	fanShuMgr.registerHander(&huPai1{
		id:             _HUPAI1_ID,
		name:           _HUPAI1_NAME,
		fanShu:         _HUPAI1_FANSHU,
		setChcFanShuID: _HUPAI1_CHECKID_,
		huKind:         _HUPAI1_KIND,
	})
}

type huPai1 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai1) GetID() int {
	return h.id
}

func (h *huPai1) Name() string {
	return h.name
}

func (h *huPai1) GetFanShu() int {
	return h.fanShu
}

func (h *huPai1) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

//东南西北风+任何一对
func (h *huPai1) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method == nil {
		return false
	}

	if method.GetChiCard() != nil {
		return false
	}

	slChkThreePairs := []uint8{31, 32, 33, 34}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	if len(slThreePairs) == 0 {
		return false
	}

	if cardType.GetSlCardDiff(slChkThreePairs, slThreePairs) != nil {
		return false
	}
	return true
}
