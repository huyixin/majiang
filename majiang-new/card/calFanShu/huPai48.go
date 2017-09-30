package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
碰碰和：由4副刻子（或杠）、将牌组成的和牌。又称“对对和(胡)”
*/

const (
	_HUPAI48_ID     = 48
	_HUPAI48_NAME   = "碰碰和"
	_HUPAI48_FANSHU = 6
	_HUPAI48_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI48_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai48{
		id:             _HUPAI48_ID,
		name:           _HUPAI48_NAME,
		fanShu:         _HUPAI48_FANSHU,
		setChcFanShuID: _HUPAI48_CHECKID_,
		huKind:         _HUPAI48_KIND,
	})
}

type huPai48 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai48) GetID() int {
	return h.id
}

func (h *huPai48) Name() string {
	return h.name
}

func (h *huPai48) GetFanShu() int {
	return h.fanShu
}

func (h *huPai48) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai48) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetPengCard(), method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetAnKe()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)

	if len(slThreePairs) != 4 {
		return false
	}
	return true
}
