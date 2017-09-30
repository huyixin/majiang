package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
明杠：自己有暗刻，碰别人打出的一张相同的牌开杠；或自己抓进一张与碰的明刻相同的牌开杠。
*/

const (
	_HUPAI74_ID     = 74
	_HUPAI74_NAME   = "明杠"
	_HUPAI74_FANSHU = 1
	_HUPAI74_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI74_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai74{
		id:             _HUPAI74_ID,
		name:           _HUPAI74_NAME,
		fanShu:         _HUPAI74_FANSHU,
		setChcFanShuID: _HUPAI74_CHECKID_,
		huKind:         _HUPAI74_KIND,
	})
}

type huPai74 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai74) GetID() int {
	return h.id
}

func (h *huPai74) Name() string {
	return h.name
}

func (h *huPai74) GetFanShu() int {
	return h.fanShu
}

func (h *huPai74) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai74) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetUnHiddenGangCard()) == 1 {
		return true
	}
	return false
}
