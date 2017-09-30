package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

/*说明：
老少副：一种花色牌的123、789两副顺子。
*/

const (
	_HUPAI72_ID     = 72
	_HUPAI72_NAME   = "老少副"
	_HUPAI72_FANSHU = 1
	_HUPAI72_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI72_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai72{
		id:             _HUPAI72_ID,
		name:           _HUPAI72_NAME,
		fanShu:         _HUPAI72_FANSHU,
		setChcFanShuID: _HUPAI72_CHECKID_,
		huKind:         _HUPAI72_KIND,
	})
}

type huPai72 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai72) GetID() int {
	return h.id
}

func (h *huPai72) Name() string {
	return h.name
}

func (h *huPai72) GetFanShu() int {
	return h.fanShu
}

func (h *huPai72) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
	iTimes := h.CheckSatisfySelf(method)
	if iTimes == 0 {
		slBanID = append(slBanID, h.GetID())
		return false, 0, satisfyedID, slBanID
	}
	//满足后把自己自己要ban的id加入进去
	for _, id := range h.setChcFanShuID {
		if !common.InIntSlace(slBanID, id) {
			slBanID = append(slBanID, id)
		}
	}

	fanShu := h.GetFanShu() * iTimes
	for i := 0; i < iTimes; i++ {
		satisfyedID = append(satisfyedID, h.GetID())
	}

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

func (h *huPai72) CheckSatisfySelf(method *cardType.HuMethod) int {
	return getLaoShaoFuTimes(method)
}

func getLaoShaoFuTimes(method *cardType.HuMethod) int {
	slAllShunZis := method.GetShunZi()
	slAllShunZis = append(slAllShunZis, method.GetChiCard()...)
	if len(slAllShunZis) < 2 {
		return 0
	}

	iTimes := 0
	slOtherChkShunZi := slAllShunZis
	for _, slShunZi1 := range slAllShunZis {
		for _, slShunZi2 := range slOtherChkShunZi {
			if checkLaoShao(slShunZi1, slShunZi2) {
				iTimes += 1
			}
		}
	}
	return iTimes / 2
}

func checkLaoShao(slShunZi1 [3]uint8, slShunZi2 [3]uint8) bool {
	onwerCard := cardType.OwnerCardType(slShunZi1[:])
	onwerCard = onwerCard.Push(slShunZi2[:])

	if onwerCard.GetColorCnt() != 1 {
		return false
	}
	slChkNum := []uint8{1, 2, 3, 7, 8, 9}
	for _, card := range onwerCard.GetAllCard() {
		remindCard := card % 10
		if common.InUInt8Slace(slChkNum, remindCard) {
			slChkNum = common.RemoveUint8Slace(slChkNum, remindCard)
		} else {
			return false
		}
	}
	return true
}
