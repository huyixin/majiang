package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
全带幺：和牌时，每副牌、将牌都有幺牌。
*/

const (
	_HUPAI55_ID     = 55
	_HUPAI55_NAME   = "全带幺"
	_HUPAI55_FANSHU = 4
	_HUPAI55_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI55_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai55{
		id:             _HUPAI55_ID,
		name:           _HUPAI55_NAME,
		fanShu:         _HUPAI55_FANSHU,
		setChcFanShuID: _HUPAI55_CHECKID_,
		huKind:         _HUPAI55_KIND,
	})
}

type huPai55 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai55) GetID() int {
	return h.id
}

func (h *huPai55) Name() string {
	return h.name
}

func (h *huPai55) GetFanShu() int {
	return h.fanShu
}

func (h *huPai55) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai55) CheckSatisfySelf(method *cardType.HuMethod) bool {
	if method.GetJiangCard() == 0 {
		return false
	}

	//判断将
	if !chkAllHasYao([]uint8{method.GetJiangCard()}) {
		return false
	}
	//顺子
	for _, slCard := range method.GetShunZi() {
		if !checkOneHasYao(slCard[:]) {
			return false
		}
	}
	//暗刻
	if !chkAllHasYao(method.GetAnKe()) {
		return false
	}

	//暗杠
	if !chkAllHasYao(method.GetHiddenGangCard()) {
		return false
	}

	//碰
	if !chkAllHasYao(method.GetPengCard()) {
		return false
	}

	//直杠
	if !chkAllHasYao(method.GetUnHiddenGangCard()) {
		return false
	}

	//吃牌
	for _, slCard := range method.GetChiCard() {
		if !checkOneHasYao(slCard[:]) {
			return false
		}
	}

	return true
}

func chkAllHasYao(slCard []uint8) bool {
	for _, card := range slCard {
		if cardType.IsFengCard(card) {
			continue
		}
		if card%10 != 1 && card%10 != 9 {
			return false
		}
	}
	return true
}

func checkOneHasYao(slCard []uint8) bool {
	if len(slCard) == 0 {
		return false
	}
	for _, card := range slCard {
		if card%10 == 1 || card%10 == 9 {
			return true
		}
	}
	return false
}
