package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
小于五：由序数牌1-4的顺子、刻子、将牌组成的和牌。不计无字。
*/

const (
	_HUPAI37_ID     = 37
	_HUPAI37_NAME   = "小于五"
	_HUPAI37_FANSHU = 12
	_HUPAI37_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI37_CHECKID_ = []int{76} //

func init() {
	fanShuMgr.registerHander(&huPai37{
		id:             _HUPAI37_ID,
		name:           _HUPAI37_NAME,
		fanShu:         _HUPAI37_FANSHU,
		setChcFanShuID: _HUPAI37_CHECKID_,
		huKind:         _HUPAI37_KIND,
	})
}

type huPai37 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai37) GetID() int {
	return h.id
}

func (h *huPai37) Name() string {
	return h.name
}

func (h *huPai37) GetFanShu() int {
	return h.fanShu
}

func (h *huPai37) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai37) CheckSatisfySelf(method *cardType.HuMethod) bool {
	for _, card := range method.GetAllCardKind() {
		if cardType.IsFengCard(card) {
			return false
		}
		if card%10 > 4 {
			return false
		}
	}
	return true
}
