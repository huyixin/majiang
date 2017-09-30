package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
七对：由7个对子组成的和牌。不计门前清、单钓，自摸计不求人。又称“七对子”“七小对”等。
*/

const (
	_HUPAI19_ID     = 19
	_HUPAI19_NAME   = "七对"
	_HUPAI19_FANSHU = 24
	_HUPAI19_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI19_CHECKID_ = []int{56, 62, 79} //

func init() {
	fanShuMgr.registerHander(&huPai19{
		id:             _HUPAI19_ID,
		name:           _HUPAI19_NAME,
		fanShu:         _HUPAI19_FANSHU,
		setChcFanShuID: _HUPAI19_CHECKID_,
		huKind:         _HUPAI19_KIND,
	})
}

type huPai19 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai19) GetID() int {
	return h.id
}

func (h *huPai19) Name() string {
	return h.name
}

func (h *huPai19) GetFanShu() int {
	return h.fanShu
}

func (h *huPai19) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai19) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil || method.GetShunZi() != nil || method.GetPengCard() != nil {
		return false
	}
	if method.GetUnHiddenGangCard() != nil || method.GetUnHiddenGangCard() != nil {
		return false
	}
	handCard := method.GetAllCard()

	if len(handCard.GetAllKind()) != 7 {
		return false
	}

	for _, kind := range handCard.GetAllKind() {
		if handCard.GetCardNum(kind) != 2 {
			return false
		}
	}
	return true
}
