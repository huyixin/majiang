package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
清龙：和牌时，有一种花1-9相连接的序数牌。又称“一条龙”“一气通贯”等。
*/

const (
	_HUPAI28_ID     = 28
	_HUPAI28_NAME   = "清龙"
	_HUPAI28_FANSHU = 16
	_HUPAI28_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI28_CHECKID_ = []int{22} //

func init() {
	fanShuMgr.registerHander(&huPai28{
		id:             _HUPAI28_ID,
		name:           _HUPAI28_NAME,
		fanShu:         _HUPAI28_FANSHU,
		setChcFanShuID: _HUPAI28_CHECKID_,
		huKind:         _HUPAI28_KIND,
	})
}

type huPai28 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai28) GetID() int {
	return h.id
}

func (h *huPai28) Name() string {
	return h.name
}

func (h *huPai28) GetFanShu() int {
	return h.fanShu
}

func (h *huPai28) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai28) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slSatisfyShunZi := getQingLongSatisfyShunZi(method)
	if len(slSatisfyShunZi) != 3 {
		return false
	}
	return true
}

func getQingLongSatisfyShunZi(method *cardType.HuMethod) (slSatisfyShunZi [][3]uint8) {
	slShunZi := method.GetShunZi()
	for _, slChiCard := range method.GetChiCard() {
		slShunZi = append(slShunZi, slChiCard)
	}
	if len(slShunZi) < 3 {
		return slSatisfyShunZi
	}
	//Cn3 n个数据取出3个
	for index, slCards := range slShunZi {
		if len(slShunZi)-1-index < 2 { //还需俩个
			break
		}
		slShunZiLayer2 := slShunZi[index+1:]
		for index1, slCards1 := range slShunZiLayer2 {
			if len(slShunZiLayer2)-1-index1 < 1 { //还需一个
				break
			}
			for _, slCards2 := range slShunZiLayer2[index1+1:] {
				onwerCard := cardType.OwnerCardType(slCards[:])
				onwerCard = onwerCard.Push(slCards1[:])
				onwerCard = onwerCard.Push(slCards2[:])
				if onwerCard.GetColorCnt() != 1 {
					continue
				}
				var slRemind []uint8
				//9个必须是1-9
				for _, card := range onwerCard.GetAllCard() {
					//有重复就一定不满足
					if common.InUInt8Slace(slRemind, card) {
						break
					}
					slRemind = append(slRemind, card)
				}
				if len(slRemind) == 9 {
					slSatisfyShunZi = append(slSatisfyShunZi, slCards, slCards1, slCards2)
					return slSatisfyShunZi
				}
			}
		}
	}
	return slSatisfyShunZi
}
