package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
双暗刻：2个暗刻。
*/

const (
	_HUPAI66_ID     = 66
	_HUPAI66_NAME   = "双暗刻"
	_HUPAI66_FANSHU = 2
	_HUPAI66_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI66_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai66{
		id:             _HUPAI66_ID,
		name:           _HUPAI66_NAME,
		fanShu:         _HUPAI66_FANSHU,
		setChcFanShuID: _HUPAI66_CHECKID_,
		huKind:         _HUPAI66_KIND,
	})
}

type huPai66 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai66) GetID() int {
	return h.id
}

func (h *huPai66) Name() string {
	return h.name
}

func (h *huPai66) GetFanShu() int {
	return h.fanShu
}

func (h *huPai66) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai66) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetAnKe()) == 2 {
		return true
	}
	return false
}
