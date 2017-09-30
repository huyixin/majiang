package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
组合龙：3种花色的147、258、369不能错位的序数牌。若和牌时另两组牌为顺子+将，可计平和。
*/

const (
	_HUPAI35_ID     = 35
	_HUPAI35_NAME   = "组合龙"
	_HUPAI35_FANSHU = 12
	_HUPAI35_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI35_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai35{
		id:             _HUPAI35_ID,
		name:           _HUPAI35_NAME,
		fanShu:         _HUPAI35_FANSHU,
		setChcFanShuID: _HUPAI35_CHECKID_,
		huKind:         _HUPAI35_KIND,
	})
}

type huPai35 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai35) GetID() int {
	return h.id
}

func (h *huPai35) Name() string {
	return h.name
}

func (h *huPai35) GetFanShu() int {
	return h.fanShu
}

func (h *huPai35) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai35) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return method.HasZuHeLong()
}
