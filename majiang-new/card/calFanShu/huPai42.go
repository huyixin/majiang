package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三色三节高：和牌时，有3种花色3副依次递增一位数的刻子。
*/

const (
	_HUPAI42_ID     = 42
	_HUPAI42_NAME   = "三色三节高"
	_HUPAI42_FANSHU = 8
	_HUPAI42_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI42_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai42{
		id:             _HUPAI42_ID,
		name:           _HUPAI42_NAME,
		fanShu:         _HUPAI42_FANSHU,
		setChcFanShuID: _HUPAI42_CHECKID_,
		huKind:         _HUPAI42_KIND,
	})
}

type huPai42 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai42) GetID() int {
	return h.id
}

func (h *huPai42) Name() string {
	return h.name
}

func (h *huPai42) GetFanShu() int {
	return h.fanShu
}

func (h *huPai42) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai42) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := method.GetAllKeZi()
	if len(slThreePairs) < 3 {
		return false
	}

	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			continue
		}
		slTmpKind := []uint8{0, 1, 2}
		base := card % 10
		slTmpKind = common.RemoveUint8Slace(slTmpKind, card/10)
		//三种情况下的暗刻
		slChkCard := []([]uint8){[]uint8{slTmpKind[0]*10 + base - 2, slTmpKind[1]*10 + base - 1},
			[]uint8{slTmpKind[0]*10 + base - 1, slTmpKind[1]*10 + base + 1},
			[]uint8{slTmpKind[0]*10 + base + 1, slTmpKind[1]*10 + base + 2},
		}
		for _, slChkCards := range slChkCard {
			if cardType.CheckContain(slThreePairs, slChkCards) {
				return true
			}
		}
	}
	return false
}
