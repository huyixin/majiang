package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
4个杠，因和牌时总共18张牌，又称“十八罗汉”。
*/

const (
	_HUPAI5_ID     = 5
	_HUPAI5_NAME   = "四杠"
	_HUPAI5_FANSHU = 88
	_HUPAI5_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI5_CHECKID_ = []int{1, 2} //

func init() {
	fanShuMgr.registerHander(&huPai5{
		id:             _HUPAI5_ID,
		name:           _HUPAI5_NAME,
		fanShu:         _HUPAI5_FANSHU,
		setChcFanShuID: _HUPAI5_CHECKID_,
		huKind:         _HUPAI5_KIND,
	})
}

type huPai5 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai5) GetID() int {
	return h.id
}

func (h *huPai5) Name() string {
	return h.name
}

func (h *huPai5) GetFanShu() int {
	return h.fanShu
}

func (h *huPai5) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai5) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetHiddenGangCard())+len(method.GetUnHiddenGangCard()) < 4 {
		return false
	}
	return true
}
