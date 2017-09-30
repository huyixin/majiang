package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/card/checkHu"
	"majiang-new/common"
)

/*说明：
七星不靠：必须有7个单张的东西南北中发白，加上3种花色，数位按147、258、369中的7张序数牌组成没有将牌的和牌。不计五门齐、门前清、单钓。
*/

const (
	_HUPAI20_ID     = 20
	_HUPAI20_NAME   = "七星不靠"
	_HUPAI20_FANSHU = 24
	_HUPAI20_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI20_CHECKID_ = []int{51, 56, 62, 79} //

func init() {
	fanShuMgr.registerHander(&huPai20{
		id:             _HUPAI20_ID,
		name:           _HUPAI20_NAME,
		fanShu:         _HUPAI20_FANSHU,
		setChcFanShuID: _HUPAI20_CHECKID_,
		huKind:         _HUPAI20_KIND,
	})
}

type huPai20 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai20) GetID() int {
	return h.id
}

func (h *huPai20) Name() string {
	return h.name
}

func (h *huPai20) GetFanShu() int {
	return h.fanShu
}

func (h *huPai20) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai20) CheckSatisfySelf(method *cardType.HuMethod) bool {
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
	if len(slWord) != 7 {
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
