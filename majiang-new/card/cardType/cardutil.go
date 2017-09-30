package cardType

import (
	"majiang-new/common"
	"reflect"
)

func GetSlCardDiff(slData []uint8, slSample []uint8) (slDiff []uint8) {
	for _, data := range slData {
		if !common.InUInt8Slace(slSample, data) {
			slDiff = append(slDiff, data)
		}
	}
	return slDiff
}

func CheckIsShunZi(cards OwnerCardType) bool {
	if cards.Len() != 3 {
		return false
	}
	if IsFengCard(cards[0]) {
		return false
	}
	if cards[1]-cards[0] != 1 || cards[2]-cards[1] != 1 {
		return false
	}
	return true
}

func CheckContain(slData1 []uint8, slData2 []uint8) bool {
	for _, card := range slData1 {
		if !common.InUInt8Slace(slData2, card) {
			continue
		}
		slData2 = common.RemoveUint8Slace(slData2, card)
	}
	if len(slData2) != 0 {
		return false
	}
	return true
}

func CheckContainShunZi(slAllShunZi [][3]uint8, slData2 [][3]uint8) bool {
	var slHased [][3]uint8
	for _, slShunZi := range slAllShunZi {
		for _, slShunZi1 := range slData2 {
			if slShunZi == slShunZi1 { //数组可以相等
				slHased = append(slHased, slShunZi1)
				break
			}
		}
	}
	return reflect.DeepEqual(slHased, slData2)
}

//得到元素集的不同的地方
func GetElemDiffByShunZi(slData1 [][3]uint8, slData2 [][3]uint8) [][3]uint8 {
	mapShunZi1 := make(map[[3]uint8]int)
	for _, slShunZi := range slData1 {
		if _, ok := mapShunZi1[slShunZi]; !ok {
			mapShunZi1[slShunZi] = 1
		} else {
			mapShunZi1[slShunZi] += 1
		}
	}

	for _, slShunZi := range slData2 {
		if num, ok := mapShunZi1[slShunZi]; ok && num > 0 {
			mapShunZi1[slShunZi] -= 1
		}
	}
	var rtnShunZi [][3]uint8
	for slShunZi, num := range mapShunZi1 {
		for i := 0; i < num; i++ {
			rtnShunZi = append(rtnShunZi, slShunZi)
		}
	}
	return rtnShunZi
}

//得到元素集的不同的地方
func GetElemDiffByCard(slData1 []uint8, slData2 []uint8) []uint8 {
	mapCard1 := make(map[uint8]int)
	for _, card := range slData1 {
		if _, ok := mapCard1[card]; !ok {
			mapCard1[card] = 1
		} else {
			mapCard1[card] += 1
		}
	}
	for _, card := range slData2 {
		if num, ok := mapCard1[card]; ok && num > 0 {
			mapCard1[card] -= 1
		}
	}

	var rtnData []uint8
	for card, num := range mapCard1 {
		for i := 0; i < num; i++ {
			rtnData = append(rtnData, card)
		}
	}
	return rtnData
}

func RemoveShunZi(slData [][3]uint8, rmShunZi [3]uint8) [][3]uint8 {
	removeIndex := -1
	for index, data := range slData {
		if data == rmShunZi {
			removeIndex = index
		}
	}
	if removeIndex == -1 {
		return slData
	}
	var slRemovedData [][3]uint8
	slRemovedData = append(slRemovedData, slData[:removeIndex]...)
	if len(slData)-1 != removeIndex {
		slRemovedData = append(slRemovedData, slData[removeIndex+1:]...)
	}
	return slRemovedData
}
