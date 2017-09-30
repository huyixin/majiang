package checkHu

import (
	"majiang-new/card/cardType"
)

func CheckHuCard(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	isHuLastCard bool,
	isHuLastPlayCard bool,
	isHuAfterGang bool,
	isHuByOtherGang bool,
	isHuLastestOne bool,
	direction int,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
	flowerCard []uint8,
	iListenStat int,
	bTianHu, bDiHu bool,
	addFanCard uint8,
	allCanHuCard []uint8,
) (bool, []*cardType.HuMethod) {
	if card != 0 {
		ownerCard = ownerCard.GetCopy().Push([]uint8{card})
		ownerCard.Sort()
	}
	return huHandlerMgr.checkCanHu(ownerCard, card, bZiMo, isHuLastCard, isHuLastPlayCard, isHuAfterGang, isHuByOtherGang, isHuLastestOne,
		direction, pengCard, hiddenGangCard, unHiddenGangCard, chiCard, flowerCard, iListenStat, bTianHu, bDiHu, addFanCard, allCanHuCard)
}
