package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
双暗杠：2个暗杠。
*/

const (
	_HUPAI53_ID     = 53
	_HUPAI53_NAME   = "双暗杠"
	_HUPAI53_FANSHU = 6
	_HUPAI53_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI53_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai53{
		id:             _HUPAI53_ID,
		name:           _HUPAI53_NAME,
		fanShu:         _HUPAI53_FANSHU,
		setChcFanShuID: _HUPAI53_CHECKID_,
		huKind:         _HUPAI53_KIND,
	})
}

type huPai53 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai53) GetID() int {
	return h.id
}

func (h *huPai53) Name() string {
	return h.name
}

func (h *huPai53) GetFanShu() int {
	return h.fanShu
}

func (h *huPai53) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai53) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetHiddenGangCard()) == 2 {
		return true
	}
	return false
}
