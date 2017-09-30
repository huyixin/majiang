package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
三色三步高：3种花色3副依次递增一位序数的顺子。
*/

const (
	_HUPAI50_ID     = 50
	_HUPAI50_NAME   = "三色三步高"
	_HUPAI50_FANSHU = 6
	_HUPAI50_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI50_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai50{
		id:             _HUPAI50_ID,
		name:           _HUPAI50_NAME,
		fanShu:         _HUPAI50_FANSHU,
		setChcFanShuID: _HUPAI50_CHECKID_,
		huKind:         _HUPAI50_KIND,
	})
}

type huPai50 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai50) GetID() int {
	return h.id
}

func (h *huPai50) Name() string {
	return h.name
}

func (h *huPai50) GetFanShu() int {
	return h.fanShu
}

func (h *huPai50) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai50) CheckSatisfySelf(method *cardType.HuMethod) bool {
	var slShunZi [][3]uint8
	for _, slChiCard := range method.GetChiCard() {
		slShunZi = append(slShunZi, slChiCard)
	}
	for _, slShunCard := range method.GetShunZi() {
		slShunZi = append(slShunZi, slShunCard)
	}

	if len(slShunZi) < 3 {
		return false
	}

	for _, slTmpCard := range slShunZi {
		slTmpKind := []uint8{0, 1, 2}
		slTmpKind = common.RemoveUint8Slace(slTmpKind, slTmpCard[0]/10)
		base := slTmpCard[0] % 10
		kindbase11 := slTmpKind[0]*10 + base + 1
		kindbase12 := slTmpKind[1]*10 + base + 2
		//第一种
		slChkShunZi1 := [][3]uint8{[...]uint8{kindbase11, kindbase11 + 1, kindbase11 + 2}, [...]uint8{kindbase12, kindbase12 + 1, kindbase12 + 2}}

		//第二种
		kindbase21 := slTmpKind[0]*10 + base + 2
		kindbase22 := slTmpKind[1]*10 + base + 1
		slChkShunZi2 := [][3]uint8{[...]uint8{kindbase21, kindbase21 + 1, kindbase21 + 2}, [...]uint8{kindbase22, kindbase22 + 1, kindbase22 + 2}}

		if cardType.CheckContainShunZi(slShunZi, slChkShunZi1) || cardType.CheckContainShunZi(slShunZi, slChkShunZi2) {
			return true
		}
	}

	return false
}
