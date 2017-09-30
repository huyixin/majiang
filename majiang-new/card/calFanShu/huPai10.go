package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
小三元：和牌时有箭牌的两副刻子和将牌。不计箭刻。
*/

const (
	_HUPAI10_ID     = 10
	_HUPAI10_NAME   = "小三元"
	_HUPAI10_FANSHU = 64
	_HUPAI10_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI10_CHECKID_ = []int{54, 59} //

func init() {
	fanShuMgr.registerHander(&huPai10{
		id:             _HUPAI10_ID,
		name:           _HUPAI10_NAME,
		fanShu:         _HUPAI10_FANSHU,
		setChcFanShuID: _HUPAI10_CHECKID_,
		huKind:         _HUPAI10_KIND,
	})
}

type huPai10 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai10) GetID() int {
	return h.id
}

func (h *huPai10) Name() string {
	return h.name
}

func (h *huPai10) GetFanShu() int {
	return h.fanShu
}

func (h *huPai10) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai10) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slCardAnKe := method.GetAnKe()
	slThreePairs := append(slCardAnKe, method.GetPengCard()...)
	jiangCard := method.GetJiangCard()

	slAddedJiang := append(slThreePairs, jiangCard)
	slCheckKind := []uint8{35, 36, 37}
	//先判断将牌
	if !common.InUInt8Slace(slCheckKind, jiangCard) {
		return false
	}

	if cardType.GetSlCardDiff(slCheckKind, slAddedJiang) != nil {
		return false
	}
	return true
}
