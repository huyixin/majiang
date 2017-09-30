package cardType

import (
	"majiang-new/common"
)

const (
	HUMETHOD_NORMAL = 1 //普通胡法
)

type HuMethod struct {
	direction int //方向

	huPaiKind        int           //胡牌的类型
	allCard          OwnerCardType //所有牌，特殊的番数判断，自行检查
	bZiMo            bool          //是否是自摸
	huCard           uint8         //当不是自摸的时候有效
	shunZi           []([3]uint8)  //手头上的顺子
	anKe             []uint8       //手头上的暗刻
	jiangCard        uint8         //奖牌，仅仅当普通胡法有效
	pengCard         []uint8
	hiddenGangCard   []uint8
	unHiddenGangCard []uint8
	chiCard          []([3]uint8)
	flowerCard       []uint8
	addFanCard       uint8
	allCanHuCard     []uint8

	isQiDui     bool
	hasZuHeLong bool //当有为组合龙胡牌的情况下有效

	isHuLastCard     bool //最后一张牌自摸
	isHuLastPlayCard bool //胡别人最后打出的牌
	isHuAfterGang    bool //杠上开花
	isHuByOtherGang  bool //抢杠和
	isHuLastestOne   bool //是否是胡一个牌的最后一张
	iListenStat      int  //听的状态
	isTianHu         bool //是不是天胡
	isDiHu           bool //是不是地胡
}

func NewHuMethod(allCard OwnerCardType,
	huPaiKind int,
	bZiMo bool,
	huCard uint8,
	shunZi []([3]uint8),
	anKe []uint8,
	jiangCard uint8,
	pengCard []uint8,
	hiddenGangCard []uint8,
	unHiddenGangCard []uint8,
	chiCard []([3]uint8),
) *HuMethod {
	if !bZiMo {
		if common.InUInt8Slace(anKe, huCard) {
			anKe = common.RemoveUint8Slace(anKe, huCard)
			pengCard = append(pengCard, huCard)
		}
	}
	return &HuMethod{
		allCard:          allCard,
		huPaiKind:        huPaiKind,
		bZiMo:            bZiMo,
		huCard:           huCard,
		shunZi:           shunZi,
		anKe:             anKe,
		jiangCard:        jiangCard,
		pengCard:         pengCard,
		hiddenGangCard:   hiddenGangCard,
		unHiddenGangCard: unHiddenGangCard,
		chiCard:          chiCard,
	}
}

func (method *HuMethod) SetAllCard(ownerCard OwnerCardType) {
	method.allCard = ownerCard
}

func (method *HuMethod) GetAllCard() OwnerCardType {
	if method.allCard == nil {
		return nil
	}
	slAllCard := method.allCard.GetCopy()
	return slAllCard
}

func (method *HuMethod) GetHandCard() OwnerCardType {
	slHandCards, _, _ := method.allCard.GetCopy().PopValue(method.huCard)
	return slHandCards
}

func (method *HuMethod) GetHuPaiKind() int {
	return method.huPaiKind
}

func (method *HuMethod) IsZiMo() bool {
	return method.bZiMo
}

func (method *HuMethod) GetHuPai() uint8 {
	return method.huCard
}

func (method *HuMethod) GetShunZi() []([3]uint8) {
	if method.shunZi == nil {
		return nil
	}
	slShunZi := make([]([3]uint8), len(method.shunZi))
	copy(slShunZi, method.shunZi)
	return slShunZi
}

func (method *HuMethod) GetAnKe() []uint8 {
	if method.anKe == nil {
		return nil
	}
	slAnke := make([]uint8, len(method.anKe))
	copy(slAnke, method.anKe)
	return slAnke
}

func (method *HuMethod) GetJiangCard() uint8 {
	return method.jiangCard
}

func (method *HuMethod) GetPengCard() []uint8 {
	if method.pengCard == nil {
		return nil
	}
	slPengCard := make([]uint8, len(method.pengCard))
	copy(slPengCard, method.pengCard)
	return slPengCard
}

func (method *HuMethod) GetHiddenGangCard() []uint8 {
	if method.hiddenGangCard == nil {
		return nil
	}
	slHiddenGangCard := make([]uint8, len(method.hiddenGangCard))
	copy(slHiddenGangCard, method.hiddenGangCard)
	return slHiddenGangCard
}

func (method *HuMethod) GetUnHiddenGangCard() []uint8 {
	if method.unHiddenGangCard == nil {
		return nil
	}
	slUnHiddenGangCard := make([]uint8, len(method.unHiddenGangCard))
	copy(slUnHiddenGangCard, method.unHiddenGangCard)
	return slUnHiddenGangCard
}

func (method *HuMethod) GetChiCard() []([3]uint8) {
	if method.chiCard == nil {
		return nil
	}
	slChiCard := make([]([3]uint8), len(method.chiCard))
	copy(slChiCard, method.chiCard)
	return slChiCard
}

func (method *HuMethod) SetFlowerCard(slCard []uint8) {
	method.flowerCard = slCard
}

func (method *HuMethod) GetFlowerCard() []uint8 {
	return method.flowerCard
}

func (method *HuMethod) SetAddFanCard(cardValue uint8) {
	method.addFanCard = cardValue
}

func (method *HuMethod) GetAddFanCard() uint8 {
	return method.addFanCard
}

func (method *HuMethod) SetAllCanHuCard(slHuCards []uint8) {
	method.allCanHuCard = slHuCards
}

func (method *HuMethod) GetAllHuCard() []uint8 {
	return method.allCanHuCard
}

func (method *HuMethod) GetAllCardKind() []uint8 {
	slHandKind := method.allCard.GetAllKind()
	slHandKind = append(slHandKind, method.pengCard...)
	slHandKind = append(slHandKind, method.hiddenGangCard...)
	slHandKind = append(slHandKind, method.unHiddenGangCard...)

	for _, slChiCard := range method.chiCard {
		for _, card := range slChiCard {
			if common.InUInt8Slace(slHandKind, card) {
				continue
			}
			slHandKind = append(slHandKind, card)
		}
	}
	return slHandKind
}

func (method *HuMethod) GetHandleCardNum() int {
	return len(method.allCard)
}

func (method *HuMethod) GetAllChiKind() []uint8 {
	var slCard []uint8
	for _, slChiData := range method.chiCard {
		for _, card := range slChiData {
			if common.InUInt8Slace(slCard, card) {
				continue
			}
			slCard = append(slCard, card)
		}
	}
	return slCard
}

func (method *HuMethod) GetAllInclude() OwnerCardType {
	p := method.GetAllCard()
	p = p.Push(method.GetUnHiddenGangCard())
	p = p.Push(method.GetPengCard())
	p = p.Push(method.GetHiddenGangCard())
	for _, slChiCard := range method.GetChiCard() {
		p = p.Push(slChiCard[:])
	}
	p.Sort()
	return p
}

func (method *HuMethod) GetAllKeZi() []uint8 {
	var allKeZi []uint8
	allKeZi = append(allKeZi, method.pengCard...)
	allKeZi = append(allKeZi, method.hiddenGangCard...)
	allKeZi = append(allKeZi, method.unHiddenGangCard...)
	allKeZi = append(allKeZi, method.anKe...)
	return allKeZi
}

func (method *HuMethod) SetIsQiDui(bQiDui bool) {
	method.isQiDui = bQiDui
}

func (method *HuMethod) IsQiDui() bool {
	return method.isQiDui
}

func (method *HuMethod) SetHasZuHeLong(hasZuHeLong bool) {
	method.hasZuHeLong = hasZuHeLong
}

func (method *HuMethod) HasZuHeLong() bool {
	return method.hasZuHeLong
}

func (method *HuMethod) IsHuLastMoCard() bool {
	return method.isHuLastCard
}

func (method *HuMethod) IsHuLastPlayCard() bool {
	return method.isHuLastPlayCard
}

func (method *HuMethod) IsHuAfterGang() bool {
	return method.isHuAfterGang
}

func (method *HuMethod) IsHuByOtherGang() bool {
	return method.isHuByOtherGang
}

func (method *HuMethod) IsHuLastestOne() bool {
	return method.isHuLastestOne
}

func (method *HuMethod) SetHuLastCard(bValue bool) {
	method.isHuLastCard = bValue
}

func (method *HuMethod) SetHuLastPlayCard(bValue bool) {
	method.isHuLastPlayCard = bValue
}

func (method *HuMethod) SetHuAfterGang(bValue bool) {
	method.isHuAfterGang = bValue
}

func (method *HuMethod) SetHuByOtherGang(bValue bool) {
	method.isHuByOtherGang = bValue
}

func (method *HuMethod) SetHuLastestOne(bValue bool) {
	method.isHuLastestOne = bValue
}

func (method *HuMethod) SetDirection(direction int) {
	method.direction = direction
}

func (method *HuMethod) GetDirection() int {
	return method.direction
}

func (method *HuMethod) GetMenFengCard() uint8 {
	return FENG_WORD_START - 1 + uint8(method.direction)
}

func (method *HuMethod) GetQuanFengCard() uint8 {
	return FENG_WORD_START //始终是东圈东局
}

func (method *HuMethod) SetListenStat(iListenStat int) {
	method.iListenStat = iListenStat
}

func (method *HuMethod) GetListenStat() int {
	return method.iListenStat
}

func (method *HuMethod) SetTianHu(isTianHu bool) {
	method.isTianHu = isTianHu
}

func (method *HuMethod) IsTianHu() bool {
	return method.isTianHu
}

func (method *HuMethod) SetDiHu(isTianHu bool) {
	method.isTianHu = isTianHu
}

func (method *HuMethod) IsDiHu() bool {
	return method.isTianHu
}
