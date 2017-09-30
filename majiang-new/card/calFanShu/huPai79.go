package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
单钓将：钓单张牌作将成和。
*/

const (
	_HUPAI79_ID     = 79
	_HUPAI79_NAME   = "单钓将"
	_HUPAI79_FANSHU = 1
	_HUPAI79_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI79_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai79{
		id:             _HUPAI79_ID,
		name:           _HUPAI79_NAME,
		fanShu:         _HUPAI79_FANSHU,
		setChcFanShuID: _HUPAI79_CHECKID_,
		huKind:         _HUPAI79_KIND,
	})
}

type huPai79 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai79) GetID() int {
	return h.id
}

func (h *huPai79) Name() string {
	return h.name
}

func (h *huPai79) GetFanShu() int {
	return h.fanShu
}

func (h *huPai79) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai79) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetAllHuCard()) != 1 {
		return false
	}
	huCard := method.GetHuPai()
	if huCard == method.GetJiangCard() {
		return true
	}
	return false
}
