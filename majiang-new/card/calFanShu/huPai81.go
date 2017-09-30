package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
花牌，可能有多个
*/

const (
	_HUPAI81_ID     = 81
	_HUPAI81_NAME   = "花牌"
	_HUPAI81_FANSHU = 1
	_HUPAI81_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI81_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai81{
		id:             _HUPAI81_ID,
		name:           _HUPAI81_NAME,
		fanShu:         _HUPAI81_FANSHU,
		setChcFanShuID: _HUPAI81_CHECKID_,
		huKind:         _HUPAI81_KIND,
	})
}

type huPai81 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai81) GetID() int {
	return h.id
}

func (h *huPai81) Name() string {
	return h.name
}

func (h *huPai81) GetFanShu() int {
	return h.fanShu
}

func (h *huPai81) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

	iTimes := h.CheckSatisfySelf(method)
	if iTimes == 0 {
		slBanID = append(slBanID, h.GetID())
		return false, 0, satisfyedID, slBanID
	}
	//满足后把自己自己要ban的id加入进去
	for _, id := range h.setChcFanShuID {
		if !common.InIntSlace(slBanID, id) {
			slBanID = append(slBanID, id)
		}
	}

	fanShu := h.GetFanShu() * iTimes
	for i := 0; i < iTimes; i++ {
		satisfyedID = append(satisfyedID, h.GetID())
	}
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

func (h *huPai81) CheckSatisfySelf(method *cardType.HuMethod) int {
	return len(method.GetFlowerCard())
}
