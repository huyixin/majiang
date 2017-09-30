package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
四归一：和牌中，有4张相同的牌归于一家的顺子、刻子、对子、将牌中（不包括杠牌）。
*/

const (
	_HUPAI64_ID     = 64
	_HUPAI64_NAME   = "四归一"
	_HUPAI64_FANSHU = 2
	_HUPAI64_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI64_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai64{
		id:             _HUPAI64_ID,
		name:           _HUPAI64_NAME,
		fanShu:         _HUPAI64_FANSHU,
		setChcFanShuID: _HUPAI64_CHECKID_,
		huKind:         _HUPAI64_KIND,
	})
}

type huPai64 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai64) GetID() int {
	return h.id
}

func (h *huPai64) Name() string {
	return h.name
}

func (h *huPai64) GetFanShu() int {
	return h.fanShu
}

func (h *huPai64) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai64) CheckSatisfySelf(method *cardType.HuMethod) int {
	handCard := method.GetAllCard()
	for _, slChiCard := range method.GetChiCard() {
		handCard = handCard.Push(slChiCard[:])
	}
	for _, pengCard := range method.GetPengCard() {
		handCard = handCard.Push([]uint8{pengCard, pengCard, pengCard}) //碰算三个
	}
	iTimes := 0
	mapChk := make(map[uint8]bool)
	for _, card := range handCard.GetAllCard() {
		if mapChk[card] {
			continue
		}
		mapChk[card] = true
		if handCard.GetCardNum(card) == 4 {
			iTimes += 1
		}
	}
	return iTimes
}
