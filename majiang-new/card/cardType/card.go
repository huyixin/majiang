package cardType

import (
	"errors"
	"fmt"
	"sort"
)

/*
	1~9		万子
	11~19	筒子
	21~29   条子
	31~37	东南西北中发白
	39~46	春夏秋冬梅兰竹菊
*/

const (
	CAN_SEQUENCE_BOUND = uint8(30) //可以成顺子的界限
	FENG_WORD_START    = uint8(31) //字牌开始的位置
	FLOWER_START_CARD  = uint8(39) //花牌的起始位置
)

func IsWanZi(card uint8) bool {
	if 1 <= card && card <= 9 {
		return true
	}
	return false
}

func IsTongZi(card uint8) bool {
	if 11 <= card && card <= 19 {
		return true
	}
	return false
}

func IsTiaoZi(card uint8) bool {
	if 21 <= card && card <= 29 {
		return true
	}
	return false
}

func IsFengCard(card uint8) bool {
	if 31 <= card && card <= 37 {
		return true
	}
	return false
}

func IsFengWordCard(card uint8) bool {
	if 31 <= card && card <= 34 {
		return true
	}
	return false
}

func IsJianCard(card uint8) bool {
	if 35 <= card && card <= 37 {
		return true
	}
	return false
}

func IsFlowerCard(card uint8) bool {
	if 39 <= card && card <= 46 {
		return true
	}
	return false
}

/*ID类型的牌，目前尚未使用，如果需要，可以把牌替换，胡牌计算和番数计算可以不替换*/

type Card struct {
	Id        int
	Value     uint8
	Direction int
}

func NewCard(id int, value uint8) Card {
	return Card{
		Id:    id,
		Value: value,
	}
}

func (card *Card) GetCardPackage() interface{} {
	mapInfo := make(map[string]interface{})
	mapInfo["Id"] = card.Id
	mapInfo["Value"] = card.Value
	mapInfo["Direction"] = card.Direction
	return mapInfo
}

func (card *Card) GetHdCardPackage() interface{} {
	mapInfo := make(map[string]interface{})
	mapInfo["Id"] = 0
	mapInfo["Value"] = 0
	mapInfo["Direction"] = card.Direction
	return mapInfo
}

type CardContainer []Card

func (c CardContainer) Len() int { return len(c) }
func (c CardContainer) Less(i, j int) bool {
	if c[i].Value < c[j].Value {
		return true
	}
	if c[i].Value == c[j].Value && c[i].Id < c[j].Id { //value第一优先级， Id第二优先级
		return true
	}
	return false
}
func (c CardContainer) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardContainer) Sort()         { sort.Sort(c) }
func (c CardContainer) Pop(id int) (CardContainer, Card, error) {
	iLen := c.Len()

	index, err := c.GetIndexByID(id)
	if err != nil {
		return c, Card{}, err
	}

	popCard := c[index]
	if index < iLen-1 {
		copy(c[index:iLen-1], c[index+1:])
	}
	return c[0 : iLen-1], popCard, nil
}

func (c CardContainer) PopValue(value Card) (CardContainer, Card, error) {
	return c.Pop(value.Id)
}

func (c CardContainer) GetIndexByID(id int) (int, error) {
	for index, card := range c {
		if card.Id == id {
			return index, nil
		}
	}
	return 0, errors.New("card not exist")
}

func (c CardContainer) GetCardByID(id int) (Card, error) {
	for _, card := range c {
		if card.Id == id {
			return card, nil
		}
	}
	return Card{}, errors.New("card not exist")
}

func (c CardContainer) Push(cards CardContainer) CardContainer {
	if len(cards) == 0 {
		return c
	}
	copyCard := make(CardContainer, c.Len()+len(cards))
	copy(copyCard[:c.Len()], c)
	copy(copyCard[c.Len():], cards)
	return copyCard[:]
}

func (c CardContainer) GetAllCard() CardContainer {
	return c
}

func (c CardContainer) GetAllCardValue() OwnerCardType {
	ownerCard := make(OwnerCardType, c.Len())
	for index, card := range c {
		ownerCard[index] = card.Value
	}
	return ownerCard
}

func (c CardContainer) GetCopy() CardContainer {
	slCard := make(CardContainer, c.Len())
	copy(slCard, c)
	return slCard
}

func (c CardContainer) GetCardPackage() []interface{} {
	slRtnCardInfo := make([]interface{}, c.Len())
	for index, card := range c {
		mapInfo := card.GetCardPackage()
		slRtnCardInfo[index] = mapInfo
	}
	return slRtnCardInfo
}

//不让其他人看到的数据
func (c CardContainer) GetHiddenCardPackage() []interface{} {
	slRtnCardInfo := make([]interface{}, c.Len())
	for index, card := range c {
		mapInfo := card.GetHdCardPackage()
		slRtnCardInfo[index] = mapInfo
	}
	return slRtnCardInfo
}

func (c CardContainer) GetCardNum(chkCard uint8) int {
	if chkCard >= FLOWER_START_CARD {
		return 0
	}
	slCard := c.GetAllCardValue()

	iTimes := 0
	for _, card := range slCard {
		if card == chkCard {
			iTimes++
		}
	}
	return iTimes
}

func GetArrCardContainPack(slContain []CardContainer) [][]interface{} {
	var slRtnCardInfo [][]interface{}
	for _, contain := range slContain {
		slRtnCardInfo = append(slRtnCardInfo, contain.GetCardPackage())
	}
	return slRtnCardInfo
}

func GetArrCardContainHdPack(slContain []CardContainer) [][]interface{} {
	var slRtnCardInfo [][]interface{}
	for _, contain := range slContain {
		slRtnCardInfo = append(slRtnCardInfo, contain.GetCardPackage())
	}
	return slRtnCardInfo
}

/***************以下为计算番数和胡牌的时候的使用类型*******************/

type OwnerCardType []uint8

func (p OwnerCardType) Len() int           { return len(p) }
func (p OwnerCardType) Less(i, j int) bool { return p[i] < p[j] }
func (p OwnerCardType) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p OwnerCardType) Sort()              { sort.Sort(p) }
func (p OwnerCardType) Pop(index int) (OwnerCardType, uint8, error) {
	iLen := p.Len()
	if index < 0 || index >= iLen {
		return nil, 0, errors.New("card index outside")
	}
	popCard := p[index]
	if index != iLen-1 {
		copy(p[index:iLen-1], p[index+1:])
	}
	return p[0 : iLen-1], popCard, nil
}

func (p OwnerCardType) PopValue(value uint8) (OwnerCardType, uint8, error) {
	index := -1
	for tmpIndex, tmpValue := range p.GetAllCard() {
		if value == tmpValue {
			index = tmpIndex
		}
	}

	if index == -1 {
		return p, 0, fmt.Errorf("p not have value")
	}
	return p.Pop(index)
}

func (p OwnerCardType) Push(card []uint8) OwnerCardType {
	if len(card) == 0 {
		return p
	}
	copyCard := make(OwnerCardType, p.Len()+len(card))
	copy(copyCard[:p.Len()], p)
	copy(copyCard[p.Len():], card)
	return copyCard[:]
}

func (p OwnerCardType) GetAllCard() []uint8 {
	return []uint8(p)
}

func (p OwnerCardType) GetCopy() OwnerCardType {
	copyCard := make(OwnerCardType, p.Len())
	copy(copyCard, p)
	return copyCard
}

//依次返回万子，筒子，条子，字， 花
func (p OwnerCardType) GetTypeSet() ([]uint8, []uint8, []uint8, []uint8, []uint8) {
	var slWan, slTong, slTiao, slWord, slFlower []uint8
	for _, card := range p {
		if card < 10 {
			slWan = append(slWan, card)
		} else if card < 20 {
			slTong = append(slTong, card)
		} else if card < 30 {
			slTiao = append(slTiao, card)
		} else if card < 38 {
			slWord = append(slWord, card)
		} else {
			slFlower = append(slFlower, card)
		}
	}
	return slWan, slTong, slTiao, slWord, slFlower
}

func (p OwnerCardType) GetAllKind() (slKind []uint8) {
	for _, card := range p.GetAllCard() {
		bRepeat := false
		for _, kind := range slKind {
			if kind == card {
				bRepeat = true
				break
			}
		}
		if bRepeat {
			continue
		}
		slKind = append(slKind, card)
	}
	return slKind
}

func (p OwnerCardType) GetColorCnt() int {
	mapColor := make(map[int]bool) // 1万， 2筒，3条
	for _, card := range p.GetAllCard() {
		if IsWanZi(card) {
			if _, ok := mapColor[1]; ok {
				continue
			} else {
				mapColor[1] = true
			}
		}

		if IsTongZi(card) {
			if _, ok := mapColor[2]; ok {
				continue
			} else {
				mapColor[2] = true
			}
		}

		if IsTiaoZi(card) {
			if _, ok := mapColor[3]; ok {
				continue
			} else {
				mapColor[3] = true
			}
		}
	}
	return len(mapColor)
}

func (p OwnerCardType) GetCardNum(chkCard uint8) int {
	if chkCard >= FLOWER_START_CARD {
		return 0
	}
	slCard := p.GetAllCard()

	iTimes := 0
	for _, card := range slCard {
		if card == chkCard {
			iTimes++
		}
	}
	return iTimes
}
