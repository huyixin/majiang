package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
五门齐：和牌时3种序数牌、风、箭牌齐全。
*/

const (
	_HUPAI51_ID     = 51
	_HUPAI51_NAME   = "五门齐"
	_HUPAI51_FANSHU = 6
	_HUPAI51_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI51_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai51{
		id:             _HUPAI51_ID,
		name:           _HUPAI51_NAME,
		fanShu:         _HUPAI51_FANSHU,
		setChcFanShuID: _HUPAI51_CHECKID_,
		huKind:         _HUPAI51_KIND,
	})
}

type huPai51 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai51) GetID() int {
	return h.id
}

func (h *huPai51) Name() string {
	return h.name
}

func (h *huPai51) GetFanShu() int {
	return h.fanShu
}

func (h *huPai51) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai51) CheckSatisfySelf(method *cardType.HuMethod) bool {
	owerIncAll := method.GetAllInclude()
	if owerIncAll.GetColorCnt() < 3 {
		return false
	}
	slChkNum := []uint8{0, 1, 2, 3, 4}
	for _, card := range owerIncAll {
		if cardType.IsJianCard(card) {
			slChkNum = common.RemoveUint8Slace(slChkNum, 4)
			continue
		}
		if cardType.IsFengWordCard(card) {
			slChkNum = common.RemoveUint8Slace(slChkNum, 3)
			continue
		}
		kind := card / 10
		slChkNum = common.RemoveUint8Slace(slChkNum, kind)
	}
	if len(slChkNum) != 0 {
		return false
	}
	return true
}
