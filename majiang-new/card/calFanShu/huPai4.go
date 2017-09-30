package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"reflect"
)

/*说明：
九莲宝灯：由一种花色序数牌按①①①②③④⑤⑥⑦⑧⑨⑨⑨组成的特定牌型，见同花色任何1张序数牌即成和牌。
不计清一色、门前清、幺九刻，自摸计不求人。因听牌时听同花色所有9种牌而得名。
如不是听九种牌的情况但和牌后牌型符合九莲宝灯牌型的，一般不算九莲宝灯，但有的场合也算(如QQ麻将)。
*/

const (
	_HUPAI4_ID     = 4
	_HUPAI4_NAME   = "九莲宝灯"
	_HUPAI4_FANSHU = 88
	_HUPAI4_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI4_CHECKID_ = []int{22, 62, 72, 76} //

func init() {
	fanShuMgr.registerHander(&huPai4{
		id:             _HUPAI4_ID,
		name:           _HUPAI4_NAME,
		fanShu:         _HUPAI4_FANSHU,
		setChcFanShuID: _HUPAI4_CHECKID_,
		huKind:         _HUPAI4_KIND,
	})
}

type huPai4 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai4) GetID() int {
	return h.id
}

func (h *huPai4) Name() string {
	return h.name
}

func (h *huPai4) GetFanShu() int {
	return h.fanShu
}

func (h *huPai4) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai4) CheckSatisfySelf(method *cardType.HuMethod) bool {
	allCard := method.GetHandCard()
	allKind := allCard.GetAllKind()
	//花色第一个都是1
	if allKind[0]%10 != 1 {
		return false
	}

	var slNeedKind []uint8
	for i := uint8(1); i < 10; i++ {
		if i == 1 || i == 9 {
			slNeedKind = append(slNeedKind, []uint8{i, i, i}...)
			continue
		}
		slNeedKind = append(slNeedKind, i)
	}
	var slChk []uint8
	for _, card := range allCard {
		if cardType.IsFengCard(card) {
			return false
		}
		slChk = append(slChk, card%10)
	}
	return reflect.DeepEqual(slChk, slNeedKind)
}
