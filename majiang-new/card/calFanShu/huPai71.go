package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"math"
)

/*说明：
连六：一种花色6张相连接的序数牌。
*/

const (
	_HUPAI71_ID     = 71
	_HUPAI71_NAME   = "连六"
	_HUPAI71_FANSHU = 1
	_HUPAI71_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI71_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai71{
		id:             _HUPAI71_ID,
		name:           _HUPAI71_NAME,
		fanShu:         _HUPAI71_FANSHU,
		setChcFanShuID: _HUPAI71_CHECKID_,
		huKind:         _HUPAI71_KIND,
	})
}

type huPai71 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai71) GetID() int {
	return h.id
}

func (h *huPai71) Name() string {
	return h.name
}

func (h *huPai71) GetFanShu() int {
	return h.fanShu
}

func (h *huPai71) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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
	//连六可能有俩个
	iTimes := h.CheckSatisfySelf(method, satisfyedID)
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

//连6比较特殊，可能有俩个
func (h *huPai71) CheckSatisfySelf(method *cardType.HuMethod, satisfyedID []int) int {
	return getLian6Times(method, satisfyedID)
}

func getLian6Times(method *cardType.HuMethod, satisfyedID []int) int {
	slAllShunZis := method.GetShunZi()
	slAllShunZis = append(slAllShunZis, method.GetChiCard()...)
	if len(slAllShunZis) < 2 {
		return 0
	}
	if common.InIntSlace(satisfyedID, 28) { //青龙
		return 0
	}
	iTimes := 0

	slOtherChkShunZi := slAllShunZis
	for _, slShunZi1 := range slAllShunZis {
		for _, slShunZi2 := range slOtherChkShunZi {
			if checkIsLian6(slShunZi1, slShunZi2) {
				iTimes += 1
			}
		}
	}
	return iTimes / 2
}

func checkIsLian6(slShunZi1 [3]uint8, slShunZi2 [3]uint8) bool {
	if uint8(math.Abs(float64(int(slShunZi1[0])-int(slShunZi2[0])))) == 3 {
		return true
	}
	return false
}
