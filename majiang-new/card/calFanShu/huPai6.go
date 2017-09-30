package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
由一种花色序数牌组成序数相连的7个对子的和牌。不计七对、清一色、门前清、单钓。
*/

const (
	_HUPAI6_ID     = 6
	_HUPAI6_NAME   = "连7对"
	_HUPAI6_FANSHU = 88
	_HUPAI6_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI6_CHECKID_ = []int{19, 22, 56, 62, 76, 79} //

func init() {
	fanShuMgr.registerHander(&huPai6{
		id:             _HUPAI6_ID,
		name:           _HUPAI6_NAME,
		fanShu:         _HUPAI6_FANSHU,
		setChcFanShuID: _HUPAI6_CHECKID_,
		huKind:         _HUPAI6_KIND,
	})
}

type huPai6 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai6) GetID() int {
	return h.id
}

func (h *huPai6) Name() string {
	return h.name
}

func (h *huPai6) GetFanShu() int {
	return h.fanShu
}

func (h *huPai6) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai6) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if !method.IsQiDui() {
		return false
	}
	allCard := method.GetAllCard()
	allKind := allCard.GetAllKind() //手上的

	if len(allKind) != 7 {
		return false
	}
	if allKind[len(allKind)-1] >= cardType.FENG_WORD_START {
		return false
	}
	//判断7对
	var slDoubleKind []uint8
	for _, kind := range allKind {
		if allCard.GetCardNum(kind) != 2 {
			return false
		}
		slDoubleKind = append(slDoubleKind, kind)
	}
	if len(slDoubleKind) != 7 {
		return false
	}

	baseKind := allKind[0]
	slChkDoubleKind := make([]uint8, 7)
	for i := 0; i < 7; i++ {
		slChkDoubleKind[i] = uint8(i) + baseKind
	}

	if cardType.GetSlCardDiff(slDoubleKind, slChkDoubleKind) != nil {
		return false
	}
	return true
}
