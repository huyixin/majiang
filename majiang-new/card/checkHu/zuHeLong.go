package checkHu

import (
	"majiang-new/card/cardType"
)

var (
	zuHeLongPattern = []([]uint8){
		[]uint8{1, 4, 7, 12, 15, 18, 23, 26, 29},
		[]uint8{1, 4, 7, 13, 16, 19, 22, 25, 28},
		[]uint8{2, 5, 8, 11, 14, 17, 23, 26, 29},
		[]uint8{2, 5, 8, 13, 16, 19, 21, 24, 27},
		[]uint8{3, 6, 9, 12, 15, 18, 21, 24, 27},
		[]uint8{3, 6, 9, 11, 14, 17, 22, 25, 28},
	}
)

func getZuHeLongPattern() []([]uint8) {
	return zuHeLongPattern
}

func checkZuHelong(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod) {

	slHasZuHeLong := CheckHasZuheLongAndReturn(ownerCard)
	if slHasZuHeLong == nil {
		return false, nil
	}
	normalHander := huHandlerMgr.getHanderByName("NormalHu")
	if normalHander == nil {
		return false, nil
	}
	slOtherCard := cardType.GetElemDiffByCard(ownerCard, slHasZuHeLong)

	ok, slMethod := normalHander(slOtherCard, card, bZiMo, pengCard, hiddenGangCard, unHiddenGangCard, chiCard)
	if !ok {
		return false, nil
	}
	//把所有牌和组合龙放入进去
	for _, method := range slMethod {
		method.SetAllCard(ownerCard)
		method.SetHasZuHeLong(true)
	}
	return true, slMethod
}

func CheckHasZuheLongAndReturn(ownerCard cardType.OwnerCardType) []uint8 {
	handCard := ownerCard.GetAllCard()
	for _, slChkCards := range zuHeLongPattern {
		if cardType.CheckContain(handCard, slChkCards) {
			return slChkCards
		}
	}
	return nil
}

func init() {
	huHandlerMgr.registerHander("ZuHelong", checkZuHelong)
}
