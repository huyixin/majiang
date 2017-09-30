package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
圈风刻：与圈风相同的风刻。
*/

const (
	_HUPAI60_ID     = 60
	_HUPAI60_NAME   = "圈风刻"
	_HUPAI60_FANSHU = 2
	_HUPAI60_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI60_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai60{
		id:             _HUPAI60_ID,
		name:           _HUPAI60_NAME,
		fanShu:         _HUPAI60_FANSHU,
		setChcFanShuID: _HUPAI60_CHECKID_,
		huKind:         _HUPAI60_KIND,
	})
}

type huPai60 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai60) GetID() int {
	return h.id
}

func (h *huPai60) Name() string {
	return h.name
}

func (h *huPai60) GetFanShu() int {
	return h.fanShu
}

func (h *huPai60) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai60) CheckSatisfySelf(method *cardType.HuMethod) bool {
	chkCard := method.GetQuanFengCard()

	slThreePairs := method.GetAllKeZi()
	for _, card := range slThreePairs {
		if chkCard == card {
			return true
		}
	}
	return false
}
