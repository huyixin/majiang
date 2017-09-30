package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三同刻:3个序数相同的刻子（杠）
*/

const (
	_HUPAI32_ID     = 32
	_HUPAI32_NAME   = "三同刻"
	_HUPAI32_FANSHU = 16
	_HUPAI32_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI32_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai32{
		id:             _HUPAI32_ID,
		name:           _HUPAI32_NAME,
		fanShu:         _HUPAI32_FANSHU,
		setChcFanShuID: _HUPAI32_CHECKID_,
		huKind:         _HUPAI32_KIND,
	})
}

type huPai32 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai32) GetID() int {
	return h.id
}

func (h *huPai32) Name() string {
	return h.name
}

func (h *huPai32) GetFanShu() int {
	return h.fanShu
}

func (h *huPai32) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai32) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	if len(slThreePairs) < 3 {
		return false
	}
	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			continue
		}
		base := card % 10
		chkKind := []uint8{0, 1, 2}
		chkKind = common.RemoveUint8Slace(chkKind, card/10)
		for _, card := range slThreePairs {
			num := card % 10
			kind := card / 10
			if common.InUInt8Slace(chkKind, kind) && num == base {
				chkKind = common.RemoveUint8Slace(chkKind, kind)
			}
		}
		if len(chkKind) == 0 {
			return true
		}
	}
	return false
}
