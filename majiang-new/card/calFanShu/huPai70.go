package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"math"
)

/*说明：
喜相逢：2种花色2副序数相同的顺子。
*/

const (
	_HUPAI70_ID     = 70
	_HUPAI70_NAME   = "喜相逢"
	_HUPAI70_FANSHU = 1
	_HUPAI70_KIND   = cardType.HUMETHOD_NORMAL
)

var _HUPAI70_CHECKID_ = []int{} //

func init() {
	fanShuMgr.registerHander(&huPai70{
		id:             _HUPAI70_ID,
		name:           _HUPAI70_NAME,
		fanShu:         _HUPAI70_FANSHU,
		setChcFanShuID: _HUPAI70_CHECKID_,
		huKind:         _HUPAI70_KIND,
	})
}

type huPai70 struct {
	id             int
	name           string
	fanShu         int
	setChcFanShuID []int
	huKind         int
}

func (h *huPai70) GetID() int {
	return h.id
}

func (h *huPai70) Name() string {
	return h.name
}

func (h *huPai70) GetFanShu() int {
	return h.fanShu
}

func (h *huPai70) Satisfy(method *cardType.HuMethod, satisfyedID []int, slBanID []int) (bool, int, []int, []int) {
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

func (h *huPai70) CheckSatisfySelf(method *cardType.HuMethod, satisfyedID []int) int {
	//连6俩个
	if getLian6Times(method, satisfyedID) == 2 && common.InIntSlace(satisfyedID, 71) {
		return 1
	}
	//老少副俩个,不管是同色不同色
	if getLaoShaoFuTimes(method) == 2 && common.InIntSlace(satisfyedID, 72) {
		return 1
	}
	return getSameShunZiTimes(method)
}

func getSameShunZiTimes(method *cardType.HuMethod) int {
	slAllShunZis := method.GetShunZi()
	slAllShunZis = append(slAllShunZis, method.GetChiCard()...)
	if len(slAllShunZis) < 2 {
		return 0
	}

	iTimes := 0
	slCheckShunZi := slAllShunZis
	for {
	contiFlag:
		for _, slShunZi1 := range slCheckShunZi {
			for _, slShunZi2 := range slCheckShunZi {
				if chkSameSZDiffColor(slShunZi1, slShunZi2) {
					iTimes += 1
					slCheckShunZi = cardType.RemoveShunZi(slCheckShunZi, slShunZi1)
					slCheckShunZi = cardType.RemoveShunZi(slCheckShunZi, slShunZi2)
					goto contiFlag
				}
			}
		}
		break
	}
	return iTimes
}

func chkSameSZDiffColor(slShunZi1 [3]uint8, slShunZi2 [3]uint8) bool {
	iAbs := uint8(math.Abs(float64(int(slShunZi1[0]) - int(slShunZi2[0]))))
	if iAbs == 10 || iAbs == 20 {
		return true
	}
	return false
}
