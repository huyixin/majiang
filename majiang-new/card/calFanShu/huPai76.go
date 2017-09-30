package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
无字：和牌中没有风、箭牌。
*/

const (
	_HUPAI76_ID     = 76
	_HUPAI76_NAME   = "无字"
	_HUPAI76_FANSHU = 1
	_HUPAI76_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI76_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai76{
		id:             _HUPAI76_ID,
		name:           _HUPAI76_NAME,
		fanShu:         _HUPAI76_FANSHU,
		setChcFanShuID: _HUPAI76_CHECKID_,
		huKind:         _HUPAI76_KIND,
	})
}

type huPai76 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai76) GetID() int {
	return h.id
}

func (h *huPai76) Name() string {
	return h.name
}

func (h *huPai76) GetFanShu() int {
	return h.fanShu
}

func (h *huPai76) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai76) CheckSatisfySelf(method *cardType.HuMethod) bool {
	ownerIncAll := method.GetAllInclude()
	for _, card := range ownerIncAll.GetAllCard() {
		if cardType.IsFengCard(card) {
			return false
		}
	}
	return true
}
