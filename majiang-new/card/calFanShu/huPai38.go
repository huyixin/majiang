package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三风刻：3个风刻。
*/

const (
	_HUPAI38_ID     = 38
	_HUPAI38_NAME   = "三风刻"
	_HUPAI38_FANSHU = 12
	_HUPAI38_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI38_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai38{
		id:             _HUPAI38_ID,
		name:           _HUPAI38_NAME,
		fanShu:         _HUPAI38_FANSHU,
		setChcFanShuID: _HUPAI38_CHECKID_,
		huKind:         _HUPAI38_KIND,
	})
}

type huPai38 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai38) GetID() int {
	return h.id
}

func (h *huPai38) Name() string {
	return h.name
}

func (h *huPai38) GetFanShu() int {
	return h.fanShu
}

func (h *huPai38) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai38) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	if len(slThreePairs) < 3 {
		return false
	}
	iTimes := 0
	for _, card := range slThreePairs {
		if cardType.IsFengWordCard(card) {
			iTimes += 1
		}
	}
	if iTimes != 3 {
		return false
	}
	return true
}
