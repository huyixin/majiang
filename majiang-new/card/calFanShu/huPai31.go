package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
全带五：每副牌及将牌必须有5的序数牌。不计断幺。
*/

const (
	_HUPAI31_ID     = 31
	_HUPAI31_NAME   = "全带五"
	_HUPAI31_FANSHU = 16
	_HUPAI31_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI31_CHECKID_ = []int{68} //

func init() {
	fanShuMgr.registerHander(&huPai31{
		id:             _HUPAI31_ID,
		name:           _HUPAI31_NAME,
		fanShu:         _HUPAI31_FANSHU,
		setChcFanShuID: _HUPAI31_CHECKID_,
		huKind:         _HUPAI31_KIND,
	})
}

type huPai31 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai31) GetID() int {
	return h.id
}

func (h *huPai31) Name() string {
	return h.name
}

func (h *huPai31) GetFanShu() int {
	return h.fanShu
}

func (h *huPai31) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai31) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetJiangCard()%10 != 5 {
		return false
	}
	for _, slChiCard := range method.GetChiCard() {
		findCard := slChiCard[0] / 10
		if !common.InUInt8Slace(slChiCard[:], findCard*10+5) {
			return false
		}
	}
	for _, slShunZiCard := range method.GetShunZi() {
		findCard := slShunZiCard[0] / 10
		if !common.InUInt8Slace(slShunZiCard[:], findCard*10+5) {
			return false
		}
	}
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	slThreePairs = append(slThreePairs, method.GetHiddenGangCard()...)
	slThreePairs = append(slThreePairs, method.GetUnHiddenGangCard()...)
	for _, card := range slThreePairs {
		if card%10 != 5 {
			return false
		}
	}
	return true
}
