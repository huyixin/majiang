package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
四暗刻：4个暗刻（暗杠）。不计门前清、碰碰和。
*/

const (
	_HUPAI12_ID     = 12
	_HUPAI12_NAME   = "四暗刻"
	_HUPAI12_FANSHU = 64
	_HUPAI12_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI12_CHECKID_ = []int{33, 48, 62} //

func init() {
	fanShuMgr.registerHander(&huPai12{
		id:             _HUPAI12_ID,
		name:           _HUPAI12_NAME,
		fanShu:         _HUPAI12_FANSHU,
		setChcFanShuID: _HUPAI12_CHECKID_,
		huKind:         _HUPAI12_KIND,
	})
}

type huPai12 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai12) GetID() int {
	return h.id
}

func (h *huPai12) Name() string {
	return h.name
}

func (h *huPai12) GetFanShu() int {
	return h.fanShu
}

func (h *huPai12) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai12) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if len(method.GetHiddenGangCard())+len(method.GetAnKe()) != 4 {
		return false
	}
	return true
}
