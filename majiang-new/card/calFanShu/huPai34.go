package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/card/checkHu"
	"majiang-new/common"
)

/*说明：
全不靠：由单张3种花色147、258、369不能错位的序数牌及东南西北中发白中的任何14张牌组成的和牌。不计五门齐、门前清、单钓。若和牌时147 258 369都有，则加计组合龙.
*/

const (
	_HUPAI34_ID     = 34
	_HUPAI34_NAME   = "全不靠"
	_HUPAI34_FANSHU = 12
	_HUPAI34_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI34_CHECKID_ = []int{51, 56, 62, 79} //

func init() {
	fanShuMgr.registerHander(&huPai34{
		id:             _HUPAI34_ID,
		name:           _HUPAI34_NAME,
		fanShu:         _HUPAI34_FANSHU,
		setChcFanShuID: _HUPAI34_CHECKID_,
		huKind:         _HUPAI34_KIND,
	})
}

type huPai34 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai34) GetID() int {
	return h.id
}

func (h *huPai34) Name() string {
	return h.name
}

func (h *huPai34) GetFanShu() int {
	return h.fanShu
}

func (h *huPai34) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai34) CheckSatisfySelf(method *cardType.HuMethod) bool {
	handCard := method.GetAllCard()
	if len(handCard) != 14 {
		return false
	}
	for _, card := range handCard.GetAllCard() {
		if handCard.GetCardNum(card) != 1 {
			return false
		}
	}
	slWanZi, slTongZi, slTiaoZi, slWord, _ := handCard.GetTypeSet()
	//7个就是七星不靠了
	if len(slWord) < 5 || len(slWord) > 6 {
		return false
	}

	slChkNum := slWanZi[:]
	slChkNum = append(slChkNum, slTongZi...)
	slChkNum = append(slChkNum, slTiaoZi...)

	if !checkHu.CheckShiSanYaoExceptWord(slChkNum) {
		return false
	}
	return true
}
