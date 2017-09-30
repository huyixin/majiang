package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
一色三节高：和牌时有一种花色3副依次递增一位数字的刻子。不计一色三同顺。
*/

const (
	_HUPAI24_ID     = 24
	_HUPAI24_NAME   = "一色三同顺"
	_HUPAI24_FANSHU = 24
	_HUPAI24_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI24_CHECKID_ = []int{23} //

func init() {
	fanShuMgr.registerHander(&huPai24{
		id:             _HUPAI24_ID,
		name:           _HUPAI24_NAME,
		fanShu:         _HUPAI24_FANSHU,
		setChcFanShuID: _HUPAI24_CHECKID_,
		huKind:         _HUPAI24_KIND,
	})
}

type huPai24 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai24) GetID() int {
	return h.id
}

func (h *huPai24) Name() string {
	return h.name
}

func (h *huPai24) GetFanShu() int {
	return h.fanShu
}

func (h *huPai24) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
func (h *huPai24) CheckSatisfySelf(method *cardType.HuMethod) bool {
	slThreePairs := append(method.GetAnKe(), method.GetPengCard()...)
	cards := cardType.OwnerCardType(slThreePairs)
	iLen := cards.Len()
	if iLen < 3 || iLen > 4 {
		return false
	}
	cards.Sort()
	if iLen == 3 {
		return cardType.CheckIsShunZi(cards)
	}
	for index, _ := range cards {
		slPopCard, _, err := cards.GetCopy().Pop(index)
		if err != nil {
			return false
		}
		if cardType.CheckIsShunZi(slPopCard) {
			return true
		}
	}
	return false
}
