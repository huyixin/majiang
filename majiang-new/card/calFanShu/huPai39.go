package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
花龙：3种花色的3副顺子连接成1-9的序数牌。
*/

const (
	_HUPAI39_ID     = 39
	_HUPAI39_NAME   = "花龙"
	_HUPAI39_FANSHU = 8
	_HUPAI39_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI39_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai39{
		id:             _HUPAI39_ID,
		name:           _HUPAI39_NAME,
		fanShu:         _HUPAI39_FANSHU,
		setChcFanShuID: _HUPAI39_CHECKID_,
		huKind:         _HUPAI39_KIND,
	})
}

type huPai39 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai39) GetID() int {
	return h.id
}

func (h *huPai39) Name() string {
	return h.name
}

func (h *huPai39) GetFanShu() int {
	return h.fanShu
}

func (h *huPai39) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai39) CheckSatisfySelf(method *cardType.HuMethod) bool {
	chkNum1 := []uint8{1, 2, 3, 14, 15, 16, 27, 28, 29}
	chkNum2 := []uint8{1, 2, 3, 17, 18, 19, 24, 25, 26}
	chkNum3 := []uint8{4, 5, 6, 11, 12, 13, 27, 28, 29}
	chkNum4 := []uint8{4, 5, 6, 17, 18, 19, 21, 22, 23}
	chkNum5 := []uint8{7, 8, 9, 11, 12, 13, 24, 25, 26}
	chkNum6 := []uint8{7, 8, 9, 14, 15, 16, 21, 22, 23}
	slCheck := []([]uint8){chkNum1, chkNum2, chkNum3, chkNum4, chkNum5, chkNum6}
	var slShunZi []uint8
	for _, slCard := range method.GetShunZi() {
		slShunZi = append(slShunZi, slCard[:]...)
	}
	for _, slCard := range method.GetChiCard() {
		slShunZi = append(slShunZi, slCard[:]...)
	}
	ownerCard := cardType.OwnerCardType(slShunZi)
	ownerCard.Sort()
	for _, slCheckCard := range slCheck {
		if cardType.CheckContain(ownerCard, slCheckCard) {
			return true
		}
	}
	return false
}
