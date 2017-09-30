package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
一色双龙会：一种花色的两个老少副，5为将牌。不计平和、七对、清一色、一般高、老少副。
*/

const (
	_HUPAI13_ID     = 13
	_HUPAI13_NAME   = "一色双龙会"
	_HUPAI13_FANSHU = 64
	_HUPAI13_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI13_CHECKID_ = []int{63, 19, 22, 69, 72} //

func init() {
	fanShuMgr.registerHander(&huPai13{
		id:             _HUPAI13_ID,
		name:           _HUPAI13_NAME,
		fanShu:         _HUPAI13_FANSHU,
		setChcFanShuID: _HUPAI13_CHECKID_,
		huKind:         _HUPAI13_KIND,
	})
}

type huPai13 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai13) GetID() int {
	return h.id
}

func (h *huPai13) Name() string {
	return h.name
}

func (h *huPai13) GetFanShu() int {
	return h.fanShu
}

func (h *huPai13) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai13) CheckSatisfySelf(method *cardType.HuMethod) bool {
	//一个颜色的俩个老少副，然后再加同色的将-5，也就是7对
	if method.GetPengCard() != nil {
		return false
	}
	if method.GetUnHiddenGangCard() != nil {
		return false
	}
	if method.GetHiddenGangCard() != nil {
		return false
	}
	//两幅老少副，碰，杠都不行 吃也只能1,2,3， 7,8,9
	handCard := method.GetAllCard()
	if handCard.GetColorCnt() != 1 { // check handcard color
		return false
	}

	slChiCard := method.GetChiCard()
	for _, slChildCard := range slChiCard {
		handCard = handCard.Push(slChildCard[:])
	}
	handCard.Sort()

	if handCard.GetColorCnt() != 1 {
		return false
	}

	slChkNum := []uint8{1, 2, 3, 5, 7, 8, 9}
	var slRemainder []uint8

	for _, card := range handCard.GetAllCard() {
		if cardType.IsFengCard(card) {
			return false
		}
		if handCard.GetCardNum(card) != 2 {
			return false
		}
		slRemainder = append(slRemainder, uint8(card%10))
	}
	if cardType.GetSlCardDiff(slRemainder, slChkNum) != nil {
		return false
	}
	return true
}
