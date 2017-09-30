package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
天胡
*/

const (
	_HUPAI84_ID     = 84
	_HUPAI84_NAME   = "天胡"
	_HUPAI84_FANSHU = 32
	_HUPAI84_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI84_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai84{
		id:             _HUPAI84_ID,
		name:           _HUPAI84_NAME,
		fanShu:         _HUPAI84_FANSHU,
		setChcFanShuID: _HUPAI84_CHECKID_,
		huKind:         _HUPAI84_KIND,
	})
}

type huPai84 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai84) GetID() int {
	return h.id
}

func (h *huPai84) Name() string {
	return h.name
}

func (h *huPai84) GetFanShu() int {
	return h.fanShu
}

func (h *huPai84) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai84) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.IsTianHu() {
		return true
	}
	return false
}
