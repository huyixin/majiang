package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
一色三步高：和牌时，有一种花色3副依次递增一位或依次递增二位数字的顺子。
*/

const (
	_HUPAI30_ID     = 30
	_HUPAI30_NAME   = "一色三步高"
	_HUPAI30_FANSHU = 16
	_HUPAI30_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI30_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai30{
		id:             _HUPAI30_ID,
		name:           _HUPAI30_NAME,
		fanShu:         _HUPAI30_FANSHU,
		setChcFanShuID: _HUPAI30_CHECKID_,
		huKind:         _HUPAI30_KIND,
	})
}

type huPai30 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai30) GetID() int {
	return h.id
}

func (h *huPai30) Name() string {
	return h.name
}

func (h *huPai30) GetFanShu() int {
	return h.fanShu
}

func (h *huPai30) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai30) CheckSatisfySelf(method *cardType.HuMethod) bool {
	var slShunZiCard [][3]uint8
	for _, slShunZi := range method.GetShunZi() {
		slShunZiCard = append(slShunZiCard, slShunZi)
	}
	for _, slShunZi := range method.GetChiCard() {
		slShunZiCard = append(slShunZiCard, slShunZi)
	}
	if len(slShunZiCard) < 3 {
		return false
	}
	if len(slShunZiCard) == 3 {
		if checkHuPai32(getAllShunZiCards32(slShunZiCard)) {
			return true
		}
	} else {
		for i := 0; i < len(slShunZiCard); i++ {
			var slCheckShunZi [][3]uint8
			slCheckShunZi = append(slCheckShunZi, slShunZiCard[:i]...)
			if i != len(slCheckShunZi)-1 {
				slCheckShunZi = append(slCheckShunZi, slShunZiCard[i+1:]...)
			}
			if checkHuPai32(getAllShunZiCards32(slCheckShunZi)) {
				return true
			}
		}
	}
	return false
}

func getAllShunZiCards32(slShunZiCard [][3]uint8) cardType.OwnerCardType {
	var allCards cardType.OwnerCardType
	for _, slShunZi := range slShunZiCard {
		allCards = allCards.Push(slShunZi[:])
	}
	allCards.Sort()
	return allCards
}

//1, 2, 2, 3, 3, 3, 4, 5 或者 1, 2, 3, 3, 4, 5, 5, 6, 7
func checkHuPai32(ownCard cardType.OwnerCardType) bool {
	base := ownCard[0] - 1
	slHasNum := make([]uint8, ownCard.Len())
	for index, card := range ownCard.GetAllCard() {
		slHasNum[index] = card - base
	}
	slChkNum := []uint8{1, 2, 2, 3, 3, 3, 4, 5}
	slChkNum1 := []uint8{1, 2, 3, 3, 4, 5, 5, 6, 7}
	if cardType.CheckContain(ownCard.GetAllCard(), slChkNum) || cardType.CheckContain(ownCard.GetAllCard(), slChkNum1) {
		return true
	}
	return false
}
