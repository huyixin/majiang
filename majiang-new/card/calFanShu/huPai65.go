package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
双同刻：2副序数相同的刻子。
*/

const (
	_HUPAI65_ID     = 65
	_HUPAI65_NAME   = "双同刻"
	_HUPAI65_FANSHU = 2
	_HUPAI65_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI65_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai65{
		id:             _HUPAI65_ID,
		name:           _HUPAI65_NAME,
		fanShu:         _HUPAI65_FANSHU,
		setChcFanShuID: _HUPAI65_CHECKID_,
		huKind:         _HUPAI65_KIND,
	})
}

type huPai65 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai65) GetID() int {
	return h.id
}

func (h *huPai65) Name() string {
	return h.name
}

func (h *huPai65) GetFanShu() int {
	return h.fanShu
}

func (h *huPai65) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

	iTimes := h.CheckSatisfySelf(method)
	if iTimes <= 0 {
		slBanID = append(slBanID, h.GetID())
		return false, 0, satisfyedID, slBanID
	}
	//满足后把自己自己要ban的id加入进去
	for _, id := range h.setChcFanShuID {
		if !common.InIntSlace(slBanID, id) {
			slBanID = append(slBanID, id)
		}
	}

	fanShu := h.GetFanShu() * iTimes
	for i := 0; i < iTimes; i++ {
		satisfyedID = append(satisfyedID, h.GetID())
	}
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
func (h *huPai65) CheckSatisfySelf(method *cardType.HuMethod) int {
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	if len(slThreePairs) < 2 {
		return 0
	}
	mapMem := make(map[uint8]bool)
	iSumTimes := 0
	for _, card := range slThreePairs {
		if cardType.IsFengCard(card) {
			continue
		}
		base := card % 10
		if mapMem[base] {
			continue
		}
		mapMem[base] = true
		iTimes := 1
		for _, card1 := range slThreePairs {
			if cardType.IsFengCard(card1) || card1 == card {
				continue
			}
			if card1%10 == base {
				iTimes += 1
			}
		}
		if iTimes == 2 {
			iSumTimes += 1
		}
	}
	return iSumTimes
}
