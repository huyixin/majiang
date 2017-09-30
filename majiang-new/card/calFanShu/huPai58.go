package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
和绝张：和牌池、桌面已亮明的3张牌所剩的第4张牌（抢杠和不计和绝张）。
*/

const (
	_HUPAI58_ID     = 58
	_HUPAI58_NAME   = "和绝张"
	_HUPAI58_FANSHU = 4
	_HUPAI58_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI58_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai58{
		id:             _HUPAI58_ID,
		name:           _HUPAI58_NAME,
		fanShu:         _HUPAI58_FANSHU,
		setChcFanShuID: _HUPAI58_CHECKID_,
		huKind:         _HUPAI58_KIND,
	})
}

type huPai58 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai58) GetID() int {
	return h.id
}

func (h *huPai58) Name() string {
	return h.name
}

func (h *huPai58) GetFanShu() int {
	return h.fanShu
}

func (h *huPai58) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai58) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return method.IsHuLastestOne()
}
