package checkHu

import (
	"majiang-new/card/cardType"
)

func checkShiSanLan(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod) {
	if len(ownerCard) != 14 {
		return false, nil
	}
	//先检查重复性
	for _, card := range ownerCard.GetAllCard() {
		if ownerCard.GetCardNum(card) != 1 {
			return false, nil
		}
	}
	slWanZi, slTongZi, slTiaoZi, slWord, _ := ownerCard.GetTypeSet()
	if len(slWord) < 5 || len(slWord) > 7 {
		return false, nil
	}
	slChkNum := slWanZi[:]
	slChkNum = append(slChkNum, slTongZi...)
	slChkNum = append(slChkNum, slTiaoZi...)
	if !CheckShiSanYaoExceptWord(slChkNum) {
		return false, nil
	}
	method := cardType.NewHuMethod(ownerCard,
		cardType.HUMETHOD_NORMAL,
		bZiMo,
		card,
		nil,
		nil,
		0,
		nil,
		nil,
		nil,
		nil)
	//有组合龙的情况哟
	if CheckHasZuheLongAndReturn(ownerCard) != nil {
		method.SetHasZuHeLong(true)
	}

	return true, []*cardType.HuMethod{method}

}

func CheckShiSanYaoExceptWord(slCard []uint8) bool {
	for _, slZuHeLong := range getZuHeLongPattern() {
		if cardType.CheckContain(slZuHeLong, slCard) {
			return true
		}
	}
	return false
}

func init() {
	huHandlerMgr.registerHander("ShiSanLan", checkShiSanLan)
}
