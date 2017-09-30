package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
清幺九：由序数牌一、九组成的和牌。不计碰碰和、全带幺、幺九刻、无字。可计七对、两组双同刻或一组三同刻。又称“清老头”。
*/

const (
	_HUPAI8_ID     = 8
	_HUPAI8_NAME   = "清九幺"
	_HUPAI8_FANSHU = 64
	_HUPAI8_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI8_CHECKID_ = []int{48, 55, 73, 76} //

func init() {
	fanShuMgr.registerHander(&huPai8{
		id:             _HUPAI8_ID,
		name:           _HUPAI8_NAME,
		fanShu:         _HUPAI8_FANSHU,
		setChcFanShuID: _HUPAI8_CHECKID_,
		huKind:         _HUPAI8_KIND,
	})
}

type huPai8 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai8) GetID() int {
	return h.id
}

func (h *huPai8) Name() string {
	return h.name
}

func (h *huPai8) GetFanShu() int {
	return h.fanShu
}

func (h *huPai8) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai8) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetChiCard() != nil {
		return false
	}
	handCard := method.GetAllCard()
	handKind := handCard.GetAllKind()

	allKind := append(handKind, method.GetHiddenGangCard()...)
	allKind = append(allKind, method.GetPengCard()...)
	allKind = append(allKind, method.GetUnHiddenGangCard()...)

	//只存在1,9
	for _, kind := range allKind {
		if kind >= cardType.FENG_WORD_START {
			return false
		}
		remainValue := kind % 10
		if remainValue != 1 && remainValue != 9 {
			return false
		}
	}
	return true
}
