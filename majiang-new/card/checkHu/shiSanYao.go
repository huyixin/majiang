package checkHu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
)

func checkShiSanYao(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod) {
	if !CheckSanYaoHuPai(ownerCard) {
		return false, nil
	}
	return true, []*cardType.HuMethod{cardType.NewHuMethod(ownerCard,
		cardType.HUMETHOD_NORMAL,
		bZiMo,
		card,
		nil,
		nil,
		0,
		nil,
		nil,
		nil,
		nil)}
}

func CheckSanYaoHuPai(ownerCard cardType.OwnerCardType) bool {
	if len(ownerCard) != 14 {
		return false
	}
	chkNum := []uint8{1, 9, 11, 19, 21, 29, 31, 32, 33, 34, 35, 36, 37}
	handCard := ownerCard.GetAllCard()
	slDif := cardType.GetElemDiffByCard(handCard, chkNum)
	if len(slDif) != 1 {
		return false
	}
	if !common.InUInt8Slace(chkNum, slDif[0]) {
		return false
	}
	return true
}

func init() {
	huHandlerMgr.registerHander("ShiSanYao", checkShiSanYao)
}
