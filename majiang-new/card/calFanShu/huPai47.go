package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
抢杠和：和别人自抓开明杠的牌。不计和绝张。
*/

const (
	_HUPAI47_ID     = 47
	_HUPAI47_NAME   = "抢杠和"
	_HUPAI47_FANSHU = 8
	_HUPAI47_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI47_CHECKID_ = []int{47} //

func init() {
	fanShuMgr.registerHander(&huPai47{
		id:             _HUPAI47_ID,
		name:           _HUPAI47_NAME,
		fanShu:         _HUPAI47_FANSHU,
		setChcFanShuID: _HUPAI47_CHECKID_,
		huKind:         _HUPAI47_KIND,
	})
}

type huPai47 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai47) GetID() int {
	return h.id
}

func (h *huPai47) Name() string {
	return h.name
}

func (h *huPai47) GetFanShu() int {
	return h.fanShu
}

func (h *huPai47) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai47) CheckSatisfySelf(method *cardType.HuMethod) bool {
	return method.IsHuByOtherGang()
}
