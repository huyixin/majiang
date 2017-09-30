package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/card/checkHu"
	"majiang-new/common"
)

/*说明：
十三幺：由3种序数牌的一、九牌，7种字牌及其中一对作将组成的和牌。不计五门齐、门前清、单钓，自摸加计不求人。又称“国士无双”。
*/

const (
	_HUPAI7_ID     = 7
	_HUPAI7_NAME   = "十三幺"
	_HUPAI7_FANSHU = 88
	_HUPAI7_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI7_CHECKID_ = []int{51, 56, 62, 79} //

func init() {
	fanShuMgr.registerHander(&huPai7{
		id:             _HUPAI7_ID,
		name:           _HUPAI7_NAME,
		fanShu:         _HUPAI7_FANSHU,
		setChcFanShuID: _HUPAI7_CHECKID_,
		huKind:         _HUPAI7_KIND,
	})
}

type huPai7 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai7) GetID() int {
	return h.id
}

func (h *huPai7) Name() string {
	return h.name
}

func (h *huPai7) GetFanShu() int {
	return h.fanShu
}

func (h *huPai7) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai7) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return checkHu.CheckSanYaoHuPai(method.GetAllCard())
}
