package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三色三同顺：和牌时，有3种花色3副序数相同的顺子。
*/

const (
	_HUPAI41_ID     = 41
	_HUPAI41_NAME   = "三色三同顺"
	_HUPAI41_FANSHU = 8
	_HUPAI41_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI41_CHECKID_ = []int{70} //

func init() {
	fanShuMgr.registerHander(&huPai41{
		id:             _HUPAI41_ID,
		name:           _HUPAI41_NAME,
		fanShu:         _HUPAI41_FANSHU,
		setChcFanShuID: _HUPAI41_CHECKID_,
		huKind:         _HUPAI41_KIND,
	})
}

type huPai41 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai41) GetID() int {
	return h.id
}

func (h *huPai41) Name() string {
	return h.name
}

func (h *huPai41) GetFanShu() int {
	return h.fanShu
}

func (h *huPai41) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai41) CheckSatisfySelf(method *cardType.HuMethod) bool {
	var slShunZi [][3]uint8
	for _, slChiCard := range method.GetChiCard() {
		slShunZi = append(slShunZi, slChiCard)
	}
	for _, slShunCard := range method.GetShunZi() {
		slShunZi = append(slShunZi, slShunCard)
	}

	if len(slShunZi) < 3 {
		return false
	}

	for _, slTmpCard := range slShunZi {
		slTmpKind := []uint8{0, 1, 2}
		slTmpKind = common.RemoveUint8Slace(slTmpKind, slTmpCard[0]/10)
		base := slTmpCard[0] % 10
		slChkBase := [...]uint8{base, base + 1, base + 2}

		for _, slTmpCard1 := range slShunZi {
			kind1 := slTmpCard1[0] / 10
			base1 := slTmpCard1[0] % 10
			slChkBase1 := [...]uint8{base1, base1 + 1, base1 + 2}
			if common.InUInt8Slace(slTmpKind, kind1) && slChkBase == slChkBase1 {
				slTmpKind = common.RemoveUint8Slace(slTmpKind, kind1)
			}
		}
		if len(slTmpKind) == 0 {
			return true
		}
	}

	return false
}
