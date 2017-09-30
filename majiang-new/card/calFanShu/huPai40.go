package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
推不倒：由牌面图形没有上下区别的牌组成的和牌，包括1234589饼、245689条、白板。不计缺一门。
*/

const (
	_HUPAI40_ID     = 40
	_HUPAI40_NAME   = "推不倒"
	_HUPAI40_FANSHU = 8
	_HUPAI40_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI40_CHECKID_ = []int{75} //

func init() {
	fanShuMgr.registerHander(&huPai40{
		id:             _HUPAI40_ID,
		name:           _HUPAI40_NAME,
		fanShu:         _HUPAI40_FANSHU,
		setChcFanShuID: _HUPAI40_CHECKID_,
		huKind:         _HUPAI40_KIND,
	})
}

type huPai40 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai40) GetID() int {
	return h.id
}

func (h *huPai40) Name() string {
	return h.name
}

func (h *huPai40) GetFanShu() int {
	return h.fanShu
}

func (h *huPai40) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai40) CheckSatisfySelf(method *cardType.HuMethod) bool {
	pCard := method.GetAllInclude()
	slChk := []uint8{11, 12, 13, 14, 15, 18, 19, 22, 24, 25, 26, 28, 29, 37}
	for _, card := range pCard.GetAllCard() {
		if !common.InUInt8Slace(slChk, card) {
			return false
		}
	}
	return true
}
