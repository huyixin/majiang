package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
加番牌
*/

const (
	_HUPAI86_ID     = 86
	_HUPAI86_NAME   = "加番牌"
	_HUPAI86_FANSHU = 10
	_HUPAI86_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI86_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai86{
		id:             _HUPAI86_ID,
		name:           _HUPAI86_NAME,
		fanShu:         _HUPAI86_FANSHU,
		setChcFanShuID: _HUPAI86_CHECKID_,
		huKind:         _HUPAI86_KIND,
	})
}

type huPai86 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai86) GetID() int {
	return h.id
}

func (h *huPai86) Name() string {
	return h.name
}

func (h *huPai86) GetFanShu() int {
	return h.fanShu
}

func (h *huPai86) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
	if iTimes == 0 {
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

func (h *huPai86) CheckSatisfySelf(method *cardType.HuMethod) int {
	cardValue := method.GetAddFanCard()
	iTime := 0
	allCard := method.GetAllCard()
	for _, slChiCard := range method.GetChiCard() {
		allCard = allCard.Push(slChiCard[:])
	}
	//暗杠2个
	for _, pengCard := range method.GetPengCard() {
		allCard = allCard.Push([]uint8{pengCard, pengCard, pengCard})
	}
	//暗杠 明杠4个
	for _, gangCard := range method.GetHiddenGangCard() {
		allCard = allCard.Push([]uint8{gangCard, gangCard, gangCard, gangCard})
	}
	for _, gangCard := range method.GetUnHiddenGangCard() {
		allCard = allCard.Push([]uint8{gangCard, gangCard, gangCard, gangCard})
	}
	for _, card := range allCard {
		if card == cardValue {
			iTime += 1
		}
	}
	if iTime > 4 {
		iTime = 4
	}
	return iTime
}
