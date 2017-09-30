package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
门前清：没有吃、碰、明杠而听牌，和别人打出的牌。又称“门清”。
*/

const (
	_HUPAI62_ID     = 62
	_HUPAI62_NAME   = "门前清"
	_HUPAI62_FANSHU = 2
	_HUPAI62_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI62_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai62{
		id:             _HUPAI62_ID,
		name:           _HUPAI62_NAME,
		fanShu:         _HUPAI62_FANSHU,
		setChcFanShuID: _HUPAI62_CHECKID_,
		huKind:         _HUPAI62_KIND,
	})
}

type huPai62 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai62) GetID() int {
	return h.id
}

func (h *huPai62) Name() string {
	return h.name
}

func (h *huPai62) GetFanShu() int {
	return h.fanShu
}

func (h *huPai62) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai62) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.IsZiMo() {
		return false
	}
	if method.GetPengCard() != nil || method.GetChiCard() != nil || method.GetUnHiddenGangCard() != nil {
		return false
	}
	return true
}
