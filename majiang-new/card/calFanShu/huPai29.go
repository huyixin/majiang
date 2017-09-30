package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三色双龙会：2种花色2个老少副、另一种花色5作将的和牌。不计喜相逢、老少副、无字、平和。
*/

const (
	_HUPAI29_ID     = 29
	_HUPAI29_NAME   = "三色双龙会"
	_HUPAI29_FANSHU = 16
	_HUPAI29_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI29_CHECKID_ = []int{70, 72, 76, 63} //

func init() {
	fanShuMgr.registerHander(&huPai29{
		id:             _HUPAI29_ID,
		name:           _HUPAI29_NAME,
		fanShu:         _HUPAI29_FANSHU,
		setChcFanShuID: _HUPAI29_CHECKID_,
		huKind:         _HUPAI29_KIND,
	})
}

type huPai29 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai29) GetID() int {
	return h.id
}

func (h *huPai29) Name() string {
	return h.name
}

func (h *huPai29) GetFanShu() int {
	return h.fanShu
}

func (h *huPai29) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai29) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetPengCard() != nil || method.GetHiddenGangCard() != nil || method.GetUnHiddenGangCard() != nil {
		return false
	}
	jiangCard := method.GetJiangCard()
	if cardType.IsFengCard(jiangCard) {
		return false
	}
	if jiangCard%10 != 5 {
		return false
	}
	kind := jiangCard / 10
	handCard := method.GetAllCard()
	for _, slChiCard := range method.GetChiCard() {
		handCard = handCard.Push(slChiCard[:])
	}
	handCard.Sort()

	//check two laoshaofu
	slWan, slTong, slTiao, _, _ := handCard.GetTypeSet()
	chkNum := []uint8{1, 2, 3, 7, 8, 9}
	for index, slCards := range [...][]uint8{slWan, slTong, slTiao} {
		if uint8(index) == kind {
			continue
		}
		var slRemind []uint8
		for _, card := range slCards {
			slRemind = append(slRemind, card%10)
		}
		if cardType.GetSlCardDiff(chkNum, slRemind) != nil {
			return false
		}
	}
	return true
}
