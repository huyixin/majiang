package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/card/checkHu"
	"testing"
)

/*
	1~9		万子
	11~19	筒子
	21~29   条子
	31~37	东南西北中发白
	39~46	春夏秋冬梅兰竹菊
*/

//大四喜
func TestHupai1(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 31, 31, 31, 32, 32, 32, 34, 34}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 34, true, false, false, false, false, false, 1, nil, nil, []uint8{33}, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//大三元
func TestHupai2(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 34, 34, 34, 35, 35, 35, 36, 36, 36, 37, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//绿一色
func TestHupai3(t *testing.T) {
	cardSet := cardType.OwnerCardType{22, 23, 23, 24, 24, 24, 24, 28, 28, 28, 36, 36, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//九莲宝灯
func TestHupai4(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 2, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//四杠 + 大四喜
func TestHupai5(t *testing.T) {
	cardSet := cardType.OwnerCardType{37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, []uint8{31, 32, 33, 34}, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//连7对
func TestHupai6(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//十三幺
func TestHupai7(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 9, 11, 19, 21, 29, 31, 32, 33, 34, 35, 36, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//清九幺
func TestHupai8(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 9, 9, 9, 11, 11, 11, 19, 19, 19, 21, 21, 21}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, []uint8{1})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//小四喜
func TestHupai9(t *testing.T) {
	cardSet := cardType.OwnerCardType{2, 3, 31, 31, 31, 32, 32, 32, 33, 33, 34, 34, 34}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//小三元
func TestHupai10(t *testing.T) {
	cardSet := cardType.OwnerCardType{2, 3, 7, 8, 9, 35, 35, 35, 36, 36, 37, 37, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//字一色
func TestHupai11(t *testing.T) {
	cardSet := cardType.OwnerCardType{31, 31, 32, 32, 32, 33, 33, 33, 35, 35, 35, 36, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 31, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//四暗刻
func TestHupai12(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 2, 33, 33, 33, 34, 34, 34, 36, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 2, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色双龙会
func TestHupai13(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 2, 3, 3, 5, 5, 7, 7, 8, 8, 9, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 1, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色四同顺
func TestHupai14(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 2, 2, 3, 3, 5}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, []([3]uint8){[...]uint8{1, 2, 3}, [...]uint8{1, 2, 3}}, nil, 0, false, false, 0, []uint8{5})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色四节高
func TestHupai15(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色四步高
func TestHupai16(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 3, 4, 5, 5, 6, 7, 7, 8, 9, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 9, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三杠
func TestHupai17(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 5}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, []uint8{31, 32, 33}, nil, nil, 0, false, false, 0, []uint8{5})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//混幺九
func TestHupai18(t *testing.T) {
	cardSet := cardType.OwnerCardType{9, 9, 9, 11, 11, 11, 19, 19, 19, 35}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 35, true, false, false, false, false, false, 1, []uint8{1}, nil, nil, nil, nil, 0, false, false, 0, []uint8{35})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//七对
func TestHupai19(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 3, 3, 4, 4, 5, 5, 7, 7, 8, 8, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 9, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//七星不靠
func TestHupai20(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 4, 13, 16, 19, 22, 25, 31, 32, 33, 34, 35, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全双刻
func TestHupai21(t *testing.T) {
	cardSet := cardType.OwnerCardType{2, 2, 12, 12, 12, 14, 14}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 14, true, false, false, false, false, false, 1, []uint8{4, 8}, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//清一色
func TestHupai22(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 3, 3, 4, 4, 5, 5, 7, 7, 8, 8, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 9, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色三同顺
func TestHupai23(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 2, 2, 3, 3, 5, 15, 16, 17}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, []([3]uint8){[...]uint8{1, 2, 3}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色三节高
func TestHupai24(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 2, 2, 3, 3, 3, 5, 7, 8, 9}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全大
func TestHupai25(t *testing.T) {
	cardSet := cardType.OwnerCardType{7, 7, 8, 8, 9, 9, 17, 17, 18, 19, 27, 28, 29}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 17, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全中
func TestHupai26(t *testing.T) {
	cardSet := cardType.OwnerCardType{4, 4, 5, 5, 6, 6, 14, 15, 16, 16, 24, 25, 26}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 16, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全小
func TestHupai27(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 2, 2, 3, 3, 11, 12, 13, 13, 21, 22, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 13, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//清龙
func TestHupai28(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 11, 12}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 13, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三色双龙会
func TestHupai29(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 7, 8, 9, 11, 12, 13, 17, 18, 19, 25}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 25, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一色三步高
func TestHupai30(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 3, 4, 5, 15, 17, 18, 19}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 15, true, false, false, false, false, false, 1, nil, nil, nil, []([3]uint8){[...]uint8{2, 3, 4}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全带五
func TestHupai31(t *testing.T) {
	cardSet := cardType.OwnerCardType{5, 13, 14, 15, 25, 26, 27}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, []([3]uint8){[...]uint8{3, 4, 5}, [...]uint8{15, 16, 17}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三同刻
func TestHupai32(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 2, 2, 11, 11, 11, 21, 21, 21, 15}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 15, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, []uint8{15})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三暗刻
func TestHupai33(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 2, 2, 11, 11, 11, 21, 22, 23, 15}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 15, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全不靠+组合龙
func TestHupai34(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 4, 7, 12, 15, 18, 23, 26, 29, 31, 32, 33, 34}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 35, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//大于五
func TestHupai36(t *testing.T) {
	cardSet := cardType.OwnerCardType{6, 7, 8, 17, 18, 19, 17, 18, 19, 27, 28, 29, 26}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 26, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三风刻
func TestHupai38(t *testing.T) {
	cardSet := cardType.OwnerCardType{31, 31, 31, 32, 32, 32, 33, 33, 33, 1, 1, 1, 5}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 5, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//花龙
func TestHupai39(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 17, 18, 19, 24, 25, 26, 35, 35, 35, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//推不倒
func TestHupai40(t *testing.T) {
	cardSet := cardType.OwnerCardType{22, 22, 22, 24, 25, 26, 12, 13, 14, 18, 18, 37, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三色三同顺
func TestHupai41(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 14, 15, 16, 24, 25, 26, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三色三节高
func TestHupai42(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 12, 12, 12, 4, 5, 6, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, []uint8{23}, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//碰碰和
func TestHupai48(t *testing.T) {
	cardSet := cardType.OwnerCardType{37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, []uint8{1, 3, 7, 9}, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//混一色
func TestHupai49(t *testing.T) {
	cardSet := cardType.OwnerCardType{37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, []uint8{1, 3, 7, 9}, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//三色三步高
func TestHupai50(t *testing.T) {
	cardSet := cardType.OwnerCardType{4, 5, 6, 13, 14, 15, 25, 26, 27, 28, 28, 28, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//五门齐
func TestHupai51(t *testing.T) {
	cardSet := cardType.OwnerCardType{3, 4, 5, 16, 17, 18, 21, 21, 21, 34, 34, 34, 37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全求人
func TestHupai52(t *testing.T) {
	cardSet := cardType.OwnerCardType{37}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 37, false, false, false, false, false, false, 1, []uint8{1, 11, 33}, nil, nil, [][3]uint8{[...]uint8{21, 22, 23}}, nil, 0, false, false, 0, []uint8{37})
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//双箭刻
func TestHupai54(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 15, 16, 17, 35, 35, 35, 36, 36, 25, 25}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//全带幺
func TestHupai55(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 11, 11, 11, 17, 18, 19, 31, 31, 31, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//不求人
func TestHupai56(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 17, 18, 19, 31, 31, 31, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//圈风刻
func TestHupai60(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 17, 18, 19, 31, 31, 31, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 3, []uint8{5}, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//门清
func TestHupai62(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 17, 18, 19, 22, 22, 22, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, false, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//平和
func TestHupai63(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 17, 18, 19, 22, 23, 24, 36}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 36, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//四归一
func TestHupai64(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 1, 2, 3, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, [][3]uint8{[...]uint8{21, 22, 23}, [...]uint8{24, 25, 26}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//双同刻
func TestHupai65(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 2, 3, 4, 5, 6, 7, 11, 11, 11, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//断幺：和牌中没有一、九及字牌。
func TestHupai68(t *testing.T) {
	cardSet := cardType.OwnerCardType{8, 8, 8, 2, 3, 4, 5, 6, 7, 12, 12, 12, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, nil, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//一般高
func TestHupai69(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 6, 7, 8, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, [][3]uint8{[...]uint8{1, 2, 3}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//喜相逢
func TestHupai70(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 6, 7, 8, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, [][3]uint8{[...]uint8{14, 15, 16}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//连六
func TestHupai71(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 2, 3, 4, 5, 6, 11, 12, 13, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, [][3]uint8{[...]uint8{14, 15, 16}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}

//幺九刻
func TestHupai73(t *testing.T) {
	cardSet := cardType.OwnerCardType{1, 1, 1, 4, 5, 6, 6, 7, 8, 23}
	ok, slHuMethod := checkHu.CheckHuCard(cardSet, 23, true, false, false, false, false, false, 1, nil, nil, nil, [][3]uint8{[...]uint8{14, 15, 16}}, nil, 0, false, false, 0, nil)
	if !ok {
		t.Fatal("can not hupai")
	}
	fanShu, satisfyedID := GetFanShu(slHuMethod)
	t.Log("番数是", fanShu, satisfyedID)
}
