package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
混幺九：由字牌和序数牌一、九组成的和牌。不计碰碰和、全带幺。可计七对。又称“混老头”。
*/

const (
	_HUPAI18_ID     = 18
	_HUPAI18_NAME   = "混幺九"
	_HUPAI18_FANSHU = 32
	_HUPAI18_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI18_CHECKID_ = []int{48, 55, 73} //

func init() {
	fanShuMgr.registerHander(&huPai18{
		id:             _HUPAI18_ID,
		name:           _HUPAI18_NAME,
		fanShu:         _HUPAI18_FANSHU,
		setChcFanShuID: _HUPAI18_CHECKID_,
		huKind:         _HUPAI18_KIND,
	})
}

type huPai18 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai18) GetID() int {
	return h.id
}

func (h *huPai18) Name() string {
	return h.name
}

func (h *huPai18) GetFanShu() int {
	return h.fanShu
}

func (h *huPai18) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

//混幺九
func (h *huPai18) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil || method.GetShunZi() != nil {
		return false
	}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetJiangCard())
	bHasFengCard := false
	bHasYaoJiu := false
	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			bHasFengCard = true
			continue
		}
		remined := card % 10
		if remined == 1 || remined == 9 {
			bHasYaoJiu = true
			continue
		} else {
			return false
		}
	}

	return bHasFengCard && bHasYaoJiu
}
