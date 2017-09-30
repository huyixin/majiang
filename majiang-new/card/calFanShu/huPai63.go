package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
平和：由4副顺子及序数牌作将组成的和牌，边、坎、钓不影响平和。
*/

const (
	_HUPAI63_ID     = 63
	_HUPAI63_NAME   = "平和"
	_HUPAI63_FANSHU = 2
	_HUPAI63_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI63_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai63{
		id:             _HUPAI63_ID,
		name:           _HUPAI63_NAME,
		fanShu:         _HUPAI63_FANSHU,
		setChcFanShuID: _HUPAI63_CHECKID_,
		huKind:         _HUPAI63_KIND,
	})
}

type huPai63 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai63) GetID() int {
	return h.id
}

func (h *huPai63) Name() string {
	return h.name
}

func (h *huPai63) GetFanShu() int {
	return h.fanShu
}

func (h *huPai63) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai63) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetPengCard() != nil || method.GetHiddenGangCard() != nil || method.GetUnHiddenGangCard() != nil || method.GetAnKe() != nil {
		return false
	}
	allShunZi := method.GetChiCard()
	allShunZi = append(allShunZi, method.GetShunZi()...)
	if len(allShunZi) != 4 {
		return false
	}
	return true
}
