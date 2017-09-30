package checkHu

import (
	"majiang-new/card/cardType"
	"sort"
)

//这里和番数无关，只用来判定是否能胡
var huHandlerMgr *checkHuManager

func init() {
	huHandlerMgr = newCheckHuManager()
}

type checkHuHander func(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod)

type checkHuManager struct {
	mapHander map[string]checkHuHander
}

func newCheckHuManager() *checkHuManager {
	return &checkHuManager{
		mapHander: make(map[string]checkHuHander),
	}
}

func (mgr *checkHuManager) registerHander(name string, hander checkHuHander) {
	mgr.mapHander[name] = hander
}

func (mgr *checkHuManager) unRegisterHander(name string) {
	delete(mgr.mapHander, name)
}

func (mgr *checkHuManager) getHanderByName(name string) checkHuHander {
	hander, ok := mgr.mapHander[name]
	if !ok {
		return nil
	}
	return hander
}

func (mgr *checkHuManager) checkCanHu(ownerCard cardType.OwnerCardType,
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
	if !sort.IsSorted(ownerCard) {
		return false, nil
	}
	var slRtnMethod []*cardType.HuMethod
	for _, huHander := range mgr.mapHander {
		if ok, slMethod := huHander(ownerCard.GetCopy(),
			card,
			bZiMo,
			pengCard,
			hiddenGangCard,
			unHiddenGangCard,
			chiCard,
		); ok {
			slRtnMethod = append(slRtnMethod, slMethod...)
		}
	}
	if len(slRtnMethod) == 0 {
		return false, nil
	}
	for _, method := range slRtnMethod {
		method.SetFlowerCard(flowerCard)
		method.SetAddFanCard(addFanCard)
		method.SetAllCanHuCard(allCanHuCard)
		method.SetHuLastCard(isHuLastCard)
		method.SetHuLastPlayCard(isHuLastPlayCard)
		method.SetHuAfterGang(isHuAfterGang)
		method.SetHuByOtherGang(isHuByOtherGang)
		method.SetHuLastestOne(isHuLastestOne)
		method.SetDirection(direction)
		method.SetListenStat(iListenStat)
		method.SetTianHu(bTianHu)
		method.SetDiHu(bDiHu)
	}
	return true, slRtnMethod
}
