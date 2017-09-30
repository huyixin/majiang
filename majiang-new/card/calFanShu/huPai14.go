package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"reflect"
)

/*说明：
一色四同顺：一种花色4副序数相同的顺子，不计一色三节高、一般高、四归一。
*/

const (
	_HUPAI14_ID     = 14
	_HUPAI14_NAME   = "一色四同顺"
	_HUPAI14_FANSHU = 48
	_HUPAI14_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI14_CHECKID_ = []int{23, 24, 69, 64} //

func init() {
	fanShuMgr.registerHander(&huPai14{
		id:             _HUPAI14_ID,
		name:           _HUPAI14_NAME,
		fanShu:         _HUPAI14_FANSHU,
		setChcFanShuID: _HUPAI14_CHECKID_,
		huKind:         _HUPAI14_KIND,
	})
}

type huPai14 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai14) GetID() int {
	return h.id
}

func (h *huPai14) Name() string {
	return h.name
}

func (h *huPai14) GetFanShu() int {
	return h.fanShu
}

func (h *huPai14) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai14) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetPengCard() != nil {
		return false
	}
	if method.GetUnHiddenGangCard() != nil {
		return false
	}
	if method.GetHiddenGangCard() != nil {
		return false
	}
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
		if iTimes == 4 {
			return true
		}
	}
	return false
}
