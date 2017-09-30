package checkHu

import (
	"majiang-new/card/cardType"
)

/*
	普通的胡法都有一对，然后加暗刻，顺子
	麻将普通胡的算法：
		1.找出是否牌中还有花，有则不能胡
		2.找出为3n + 2的牌色，在其中，找出除队之外的所有顺子（字一般不能成为顺子）
		3.其他的都为3n,只要找出顺子或者暗刻即可

*/

//普通胡法
func checkNormalHu(ownerCard cardType.OwnerCardType,
	card uint8,
	bZiMo bool,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) (bool, []*cardType.HuMethod) {
	slWan, slTong, slTiao, slWord, slFlower := ownerCard.GetTypeSet()
	if len(slFlower) != 0 {
		return false, nil
	}
	var mapHasJiang map[uint8]interface{}
	var slThreeInfo []map[uint8]interface{}

	slNormalTypeSet := [...]([]uint8){slWan, slTong, slTiao, slWord} //循环注意index为3的字牌情况

	bHasDouble := false
	for index, slCardSet := range slNormalTypeSet {
		iSetLen := len(slCardSet)
		if iSetLen == 0 {
			continue
		}
		//花色只存在 3n, 3n+2
		if (iSetLen+2)%3 == 0 {
			return false, nil
		}
		iCanSequence := true
		if index == 3 {
			iCanSequence = false
		}
		if bHasDouble {
			if iSetLen%3 != 0 {
				return false, nil
			}

			if ok, mapTmp := checkSatisfy(slCardSet, false, iCanSequence); ok {
				slThreeInfo = append(slThreeInfo, mapTmp)
			} else {
				return false, nil
			}
		} else {
			if (iSetLen+1)%3 == 0 { //(3n+2 + 1) %3 = 0
				bHasDouble = true
			}

			if ok, mapTmp := checkSatisfy(slCardSet, bHasDouble, iCanSequence); ok {
				if bHasDouble {
					mapHasJiang = mapTmp
				} else {
					slThreeInfo = append(slThreeInfo, mapTmp)
				}
			} else {
				return false, nil
			}
		}
	}

	//再把将牌对应的顺子，暗刻全部放入进入
	var slHuMthoed []*cardType.HuMethod
	for jiangPaiCard, threeInfo := range mapHasJiang {
		var slShunZi []([3]uint8)
		var slAnKe []uint8
		slRtnAnKe, slRtnShunZi := getPairs(threeInfo)
		slAnKe = append(slAnKe, slRtnAnKe...)
		slShunZi = append(slShunZi, slRtnShunZi...)

		for _, mapThreeOther := range slThreeInfo {
			for _, threeInfo := range mapThreeOther {
				slRtnAnKe, slRtnShunZi := getPairs(threeInfo)
				slAnKe = append(slAnKe, slRtnAnKe...)
				slShunZi = append(slShunZi, slRtnShunZi...)
			}
		}
		slHuMthoed = append(slHuMthoed, cardType.NewHuMethod(ownerCard,
			cardType.HUMETHOD_NORMAL,
			bZiMo,
			card,
			slShunZi,
			slAnKe,
			jiangPaiCard,
			pengCard,
			hiddenGangCard,
			unHiddenGangCard,
			chiCard))
	}
	return true, slHuMthoed
}

func getPairs(threeInfo interface{}) (slAnKe []uint8, slShunZi []([3]uint8)) {
	mapThreeInfo, ok := threeInfo.(map[string]interface{})
	if !ok {
		return
	}

	if childShunZi, ok := mapThreeInfo["ShunZi"]; ok {
		if slChildShunZi, ok1 := childShunZi.([]([3]uint8)); ok1 {
			slShunZi = append(slShunZi, slChildShunZi...)
		}
	}
	if childAnKe, ok := mapThreeInfo["AnKe"]; ok {
		if slChildAnKe, ok1 := childAnKe.([]uint8); ok1 {
			slAnKe = append(slAnKe, slChildAnKe...)
		}
	}
	return slAnKe, slShunZi
}

//返回值，如果ok, huInfo 为有将牌的，以奖牌为key, 否则都为255为key
func checkSatisfy(slCard []uint8, shouldDoublePair bool, canSuquence bool) (ok bool, huInfo map[uint8]interface{}) {
	iCardLen := len(slCard)
	if iCardLen == 0 {
		return
	}

	if shouldDoublePair && (len(slCard)+1)%3 != 0 {
		return
	}

	huInfo = make(map[uint8]interface{})
	if shouldDoublePair {
		checkHasDoublePair(slCard, canSuquence, huInfo)
		if len(huInfo) == 0 {
			return false, nil
		}
		return true, huInfo
	}
	ok1, mapThreeInfo, _ := CheckThreePairs(slCard, canSuquence)
	return ok1, mapThreeInfo
}

func checkHasDoublePair(slCard []uint8, canSuquence bool, huInfo map[uint8]interface{}) {
	mapCardCnt := make(map[uint8]uint8)
	for _, card := range slCard {
		if _, ok := mapCardCnt[card]; ok {
			mapCardCnt[card] += 1
		} else {
			mapCardCnt[card] = 1
		}
	}
	for card, cnt := range mapCardCnt {
		_, ok := huInfo[card]
		if !ok && cnt >= 2 {

			slRemovedCard := removeTwice(slCard, card)
			if ok1, mapThreeInfo, _ := CheckThreePairs(slRemovedCard, canSuquence); ok1 {
				if len(mapThreeInfo) != 1 {
					continue
				}
				//注意这里转换，map内存存储一个map[string]interface{} 这里面放置着顺子以及暗刻信息
				mapInfo := make(map[string]interface{})
				huInfo[card] = mapInfo
				for _, mapThreePairInfo := range mapThreeInfo {
					mapThree := mapThreePairInfo.(map[string]interface{})
					for key, data := range mapThree {
						mapInfo[key] = data
					}
				}
			}
		}
	}
}

//注意，进去的牌需要排序
func removeTwice(slCard []uint8, card uint8) []uint8 {
	slRtnCard := make([]uint8, len(slCard)-2)
	rtnCardIndex := 0
	for index, tempCard := range slCard {
		if tempCard == card {
			rtnCardIndex = index
			break
		}
	}
	iLen := len(slCard[:rtnCardIndex])
	copy(slRtnCard, slCard[:rtnCardIndex])
	copy(slRtnCard[iLen:], slCard[rtnCardIndex+2:])
	return slRtnCard
}

func CheckThreePairs(slCard []uint8, canSuquence bool) (ok bool, huInfo map[uint8]interface{}, slFeiPai []uint8) {
	huInfo = make(map[uint8]interface{})
	mapInfo := make(map[string]interface{})
	huInfo[uint8(255)] = mapInfo //注意是255
	if len(slCard) == 0 {
		return true, huInfo, nil
	}
	var tmpList = make([]uint8, len(slCard))
	copy(tmpList, slCard)
	//第一个牌固定，然后其他的两个表全扫一遍
	var slShunZi []([3]uint8)
	var slAnKe []uint8

	secondIndex := 1
	thirdIndex := 2
	for {
		//i ++
		if secondIndex >= len(tmpList) {
			break
		} else {
			for {
				if thirdIndex >= len(tmpList) {
					break
				} else {
					card1 := tmpList[0]
					card2 := tmpList[secondIndex]
					card3 := tmpList[thirdIndex]

					iFlag := isThreePair(card1, card2, card3, canSuquence)
					if iFlag != 0 {
						if iFlag == 1 {
							//顺子
							slChildShunZi := [...]uint8{card1, card2, card3}
							slShunZi = append(slShunZi, slChildShunZi)
						} else if iFlag == 2 {
							slAnKe = append(slAnKe, card1)
						}
						index1 := getIndex(tmpList, card1)
						tmpList = append(tmpList[:index1], tmpList[index1+1:]...)
						index2 := getIndex(tmpList, card2)
						tmpList = append(tmpList[:index2], tmpList[index2+1:]...)
						index3 := getIndex(tmpList, card3)
						tmpList = append(tmpList[:index3], tmpList[index3+1:]...)

						// 再从前三张开始
						secondIndex = 1
						thirdIndex = 2
					} else {
						thirdIndex += 1
					}
				}
			}
			secondIndex += 1
			thirdIndex = secondIndex + 1
		}
	}
	mapInfo["ShunZi"] = slShunZi
	mapInfo["AnKe"] = slAnKe

	if 0 == len(tmpList) {
		return true, huInfo, nil
	}
	return false, huInfo, tmpList
}

func getIndex(slCard []uint8, card uint8) int {
	for index, tmpCard := range slCard {
		if card == tmpCard {
			return index
		}
	}
	return 0
}

//0表示不能形成顺子或暗刻，1代表能形成顺子，2代表是暗刻
func isThreePair(card1, card2, card3 uint8, canSuquence bool) int {
	if card1 == card2 && card2 == card3 { //暗刻比顺子番数更高，先计算暗刻，不然番会出问题
		return 2
	}
	if (card2-card1 == 1) && (card3-card2 == 1) {
		if !canSuquence {
			return 0
		}
		return 1
	}
	return 0
}

func init() {
	huHandlerMgr.registerHander("NormalHu", checkNormalHu)
}
