package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
幺九刻：3张相同的一、九序数牌及字牌组成的刻子（或杠）。
*/

const (
	_HUPAI73_ID     = 73
	_HUPAI73_NAME   = "幺九刻"
	_HUPAI73_FANSHU = 1
	_HUPAI73_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI73_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai73{
		id:             _HUPAI73_ID,
		name:           _HUPAI73_NAME,
		fanShu:         _HUPAI73_FANSHU,
		setChcFanShuID: _HUPAI73_CHECKID_,
		huKind:         _HUPAI73_KIND,
	})
}

type huPai73 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai73) GetID() int {
	return h.id
}

func (h *huPai73) Name() string {
	return h.name
}

func (h *huPai73) GetFanShu() int {
	return h.fanShu
}

func (h *huPai73) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
	iTimes := h.CheckSatisfySelf(method)
	if iTimes == 0 {
		slBanID = append(slBanID, h.GetID())
		return false, 0, satisfyedID, slBanID
	}
	//满足后把自己自己要ban的id加入进去
	for _, id := range h.setChcFanShuID {
		if !common.InIntSlace(slBanID, id) {
			slBanID = append(slBanID, id)
		}
	}

	fanShu := h.GetFanShu() * iTimes
	for i := 0; i < iTimes; i++ {
		satisfyedID = append(satisfyedID, h.GetID())
	}
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

func (h *huPai73) CheckSatisfySelf(method *cardType.HuMethod) int {
	slThreePairs := method.GetAllKeZi()

	iTimes := 0
	for _, card := range slThreePairs {
		if cardType.IsJianCard(card) {
			continue
		}
		if cardType.IsFengWordCard(card) && card != method.GetQuanFengCard() && card != method.GetMenFengCard() {
			iTimes += 1
			continue
		}
		remind := card % 10
		if remind%10 == 1 || remind%10 == 9 {
			iTimes += 1
		}

	}
	return iTimes
}
