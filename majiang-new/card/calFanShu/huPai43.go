package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
无番和：和牌后，数不出任何番种分（花牌不计算在内）。
*/

const (
	_HUPAI43_ID     = 43
	_HUPAI43_NAME   = "无番和"
	_HUPAI43_FANSHU = 8
	_HUPAI43_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI43_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai43{
		id:             _HUPAI43_ID,
		name:           _HUPAI43_NAME,
		fanShu:         _HUPAI43_FANSHU,
		setChcFanShuID: _HUPAI43_CHECKID_,
		huKind:         _HUPAI43_KIND,
	})
}

type huPai43 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai43) GetID() int {
	return h.id
}

func (h *huPai43) Name() string {
	return h.name
}

func (h *huPai43) GetFanShu() int {
	return h.fanShu
}

func (h *huPai43) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai43) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return false
}
