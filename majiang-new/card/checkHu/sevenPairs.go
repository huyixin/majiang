package checkHu

import (
	"majiang-new/card/cardType"
)

//七小对胡法, 包括有杠的胡法
func checkSevenPairs(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod) {
	//排序后，只要出现偶奇不一样的情况，就不是七对
	iLen := len(ownerCard)
	if iLen != 14 {
		return false, nil
	}
	for i := 0; i < iLen; i += 2 {
		if ownerCard[i] != ownerCard[i+1] {
			return false, nil
		}
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
	method.SetIsQiDui(true)
	return true, []*cardType.HuMethod{method}
}

func init() {
	huHandlerMgr.registerHander("SevenPairs", checkSevenPairs)
}
