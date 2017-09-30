package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
不求人：4副牌及将中没有吃牌、碰牌（包括明杠），所有的牌包括所和的牌全部是自己摸到的。又称“门清自摸”。
*/

const (
	_HUPAI56_ID     = 56
	_HUPAI56_NAME   = "不求人"
	_HUPAI56_FANSHU = 4
	_HUPAI56_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI56_CHECKID_ = []int{80} //

func init() {
	fanShuMgr.registerHander(&huPai56{
		id:             _HUPAI56_ID,
		name:           _HUPAI56_NAME,
		fanShu:         _HUPAI56_FANSHU,
		setChcFanShuID: _HUPAI56_CHECKID_,
		huKind:         _HUPAI56_KIND,
	})
}

type huPai56 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai56) GetID() int {
	return h.id
}

func (h *huPai56) Name() string {
	return h.name
}

func (h *huPai56) GetFanShu() int {
	return h.fanShu
}

func (h *huPai56) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai56) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if !method.IsZiMo() {
		return false
	}
	if method.GetPengCard() != nil || method.GetChiCard() != nil || method.GetUnHiddenGangCard() != nil {
		return false
	}
	return true
}
