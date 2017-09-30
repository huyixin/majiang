package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
小四喜：和牌时有风牌的3副刻子及将牌。不计三风刻，混一色。
*/

const (
	_HUPAI9_ID     = 9
	_HUPAI9_NAME   = "小四喜"
	_HUPAI9_FANSHU = 64
	_HUPAI9_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI9_CHECKID_ = []int{38} //

func init() {
	fanShuMgr.registerHander(&huPai9{
		id:             _HUPAI9_ID,
		name:           _HUPAI9_NAME,
		fanShu:         _HUPAI9_FANSHU,
		setChcFanShuID: _HUPAI9_CHECKID_,
		huKind:         _HUPAI9_KIND,
	})
}

type huPai9 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai9) GetID() int {
	return h.id
}

func (h *huPai9) Name() string {
	return h.name
}

func (h *huPai9) GetFanShu() int {
	return h.fanShu
}

func (h *huPai9) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai9) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slCardAnKe := method.GetAnKe()
	slCardAnKe = append(slCardAnKe, method.GetPengCard()...)
	jiangCard := method.GetJiangCard()

	slAddedJiang := append(slCardAnKe, jiangCard)
	slCheckKind := []uint8{31, 32, 33, 34}
	//先判断将牌
	if !common.InUInt8Slace(slCheckKind, jiangCard) {
		return false
	}

	if cardType.GetSlCardDiff(slCheckKind, slAddedJiang) != nil {
		return false
	}
	return true
}
