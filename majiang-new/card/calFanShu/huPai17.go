package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三杠：整副牌中有3个杠子，暗杠加计。
*/

const (
	_HUPAI17_ID     = 17
	_HUPAI17_NAME   = "三杠"
	_HUPAI17_FANSHU = 32
	_HUPAI17_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI17_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai17{
		id:             _HUPAI17_ID,
		name:           _HUPAI17_NAME,
		fanShu:         _HUPAI17_FANSHU,
		setChcFanShuID: _HUPAI17_CHECKID_,
		huKind:         _HUPAI17_KIND,
	})
}

type huPai17 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai17) GetID() int {
	return h.id
}

func (h *huPai17) Name() string {
	return h.name
}

func (h *huPai17) GetFanShu() int {
	return h.fanShu
}

func (h *huPai17) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai17) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetHiddenGangCard())+len(method.GetUnHiddenGangCard()) != 3 {
		return false
	}
	return true
}
