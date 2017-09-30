package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"reflect"
)

/*说明：
一色三同顺：和牌时有一种花色3副序数相同的顺子。不计一色三节高。
*/

const (
	_HUPAI23_ID     = 23
	_HUPAI23_NAME   = "一色三同顺"
	_HUPAI23_FANSHU = 24
	_HUPAI23_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI23_CHECKID_ = []int{24} //

func init() {
	fanShuMgr.registerHander(&huPai23{
		id:             _HUPAI23_ID,
		name:           _HUPAI23_NAME,
		fanShu:         _HUPAI23_FANSHU,
		setChcFanShuID: _HUPAI23_CHECKID_,
		huKind:         _HUPAI23_KIND,
	})
}

type huPai23 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai23) GetID() int {
	return h.id
}

func (h *huPai23) Name() string {
	return h.name
}

func (h *huPai23) GetFanShu() int {
	return h.fanShu
}

func (h *huPai23) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai23) CheckSatisfySelf(method *cardType.HuMethod) bool {
	var slShunZiCard []([3]uint8)
	for _, slShunZi := range method.GetChiCard() {
		slShunZiCard = append(slShunZiCard, slShunZi)
	}
	for _, slShunZi := range method.GetShunZi() {
		slShunZiCard = append(slShunZiCard, slShunZi)
	}
	for index, slShunZi := range slShunZiCard {
		iTimes := 1
		if len(slShunZiCard)-1 == index {
			return false
		}
		for _, slSecondShunZi := range slShunZiCard[index+1:] {
			if reflect.DeepEqual(slShunZi, slSecondShunZi) {
				iTimes += 1
			}
		}
		if iTimes == 3 {
			return true
		}
	}
	return false
}
