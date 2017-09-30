package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
门风刻：与本门风相同的风刻。
*/

const (
	_HUPAI61_ID     = 61
	_HUPAI61_NAME   = "门风刻"
	_HUPAI61_FANSHU = 2
	_HUPAI61_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI61_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai61{
		id:             _HUPAI61_ID,
		name:           _HUPAI61_NAME,
		fanShu:         _HUPAI61_FANSHU,
		setChcFanShuID: _HUPAI61_CHECKID_,
		huKind:         _HUPAI61_KIND,
	})
}

type huPai61 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai61) GetID() int {
	return h.id
}

func (h *huPai61) Name() string {
	return h.name
}

func (h *huPai61) GetFanShu() int {
	return h.fanShu
}

func (h *huPai61) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai61) CheckSatisfySelf(method *cardType.HuMethod) bool {
	chkCard := method.GetMenFengCard()

	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	for _, card := range slThreePairs {
		if chkCard == card {
			return true
		}
	}
	return false
}
