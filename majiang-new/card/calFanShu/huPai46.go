package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
杠上开花：开杠抓进的牌成和牌（不包括补花）。不计自摸。
*/

const (
	_HUPAI46_ID     = 46
	_HUPAI46_NAME   = "杠上开花"
	_HUPAI46_FANSHU = 8
	_HUPAI46_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI46_CHECKID_ = []int{80} //

func init() {
	fanShuMgr.registerHander(&huPai46{
		id:             _HUPAI46_ID,
		name:           _HUPAI46_NAME,
		fanShu:         _HUPAI46_FANSHU,
		setChcFanShuID: _HUPAI46_CHECKID_,
		huKind:         _HUPAI46_KIND,
	})
}

type huPai46 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai46) GetID() int {
	return h.id
}

func (h *huPai46) Name() string {
	return h.name
}

func (h *huPai46) GetFanShu() int {
	return h.fanShu
}

func (h *huPai46) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai46) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return method.IsHuAfterGang()
}
