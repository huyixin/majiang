package calFanShu

import (
	"majiang-new/card/cardType"
	"majiang-new/common"
	"sort"

	"log"
)

var (
	fanShuMgr *fanShuManager
)

func init() {
	fanShuMgr = newFanShuManager()
}

//是否满足，并且返回番数
type calFSInterface interface {
	Satisfy(method *cardType.HuMethod,
		satisfyedID []int, slBanID []int, //已经检测的胡牌id
	) (bool, int, []int, []int) //返回是否满足，番数， 以及满足番数hander的id集合, 已经ban检测过不满足，或者被检测过程中ban的ID
	GetID() int
	GetFanShu() int
	Name() string
}

type fanShuManager struct {
	fanShuKind sort.IntSlice              //根据番数区分
	allHander  map[int]([]calFSInterface) //按照番数分类
	handerLen  int

	mapID2Hander map[int]calFSInterface
}

func newFanShuManager() *fanShuManager {
	return &fanShuManager{
		allHander:    make(map[int]([]calFSInterface)),
		mapID2Hander: make(map[int]calFSInterface),
	}
}

func (mgr *fanShuManager) registerHander(hander calFSInterface) {
	bFind := false
	fanShu := hander.GetFanShu()
	for _, kind := range mgr.fanShuKind {
		if kind == fanShu {
			break
		} else {
			bFind = true
		}
	}
	if !bFind {
		mgr.fanShuKind = append(mgr.fanShuKind, fanShu)
		mgr.fanShuKind.Sort()
	}

	id := hander.GetID()
	if _, ok := mgr.mapID2Hander[id]; ok {
		log.Fatal("repeated fanshuHander:", id)
		return
	}

	mgr.mapID2Hander[id] = hander
	mgr.allHander[fanShu] = append(mgr.allHander[fanShu], hander)
	mgr.handerLen += 1
}

func (mgr *fanShuManager) unRegisterHander(id int) {
	hander, ok := mgr.mapID2Hander[id]
	if !ok {
		return
	}
	delete(mgr.mapID2Hander, hander.GetID())

	fanShu := hander.GetFanShu()
	if slHander, ok := mgr.allHander[fanShu]; ok {
		for index, tmpHander := range slHander {
			if tmpHander == hander {
				var deledHander []calFSInterface
				deledHander = append(deledHander, slHander[:index]...)
				deledHander = append(deledHander, slHander[index:len(slHander)]...)
				mgr.handerLen -= 1
				mgr.allHander[fanShu] = deledHander
			}
		}
	}
}

func (mgr *fanShuManager) getFanShu(slMethod []*cardType.HuMethod) (int, []int) {
	allHander := make([]calFSInterface, mgr.handerLen)
	var slFanShuKey sort.IntSlice
	for key, _ := range mgr.allHander {
		slFanShuKey = append(slFanShuKey, key)
	}

	slFanShuKey.Sort()

	index := 0
	for i := len(slFanShuKey) - 1; i >= 0; i-- {
		key := slFanShuKey[i]
		slHander := mgr.allHander[key]
		for _, hander := range slHander {
			allHander[index] = hander
			index += 1
		}
	}
	var setMaxFanShuID []int
	maxFanShu := 0
	//找到所有胡法的最大番数
	for _, method := range slMethod {
		var satisfyID []int
		var slBanID []int
		//从高到低，满足最高的就把所有的都遍历了
		for _, hander := range allHander {
			ok, tmpFanShu, tmpSatisfyID, slTmpBanID := hander.Satisfy(method, satisfyID, slBanID)
			//成功，则会返回ban, 没成功，会把其id加入进去，内部不用再继续检查
			slBanID = slTmpBanID
			if ok && maxFanShu < tmpFanShu {
				setMaxFanShuID = tmpSatisfyID
				maxFanShu = tmpFanShu
				break
			}
		}
	}
	return maxFanShu, setMaxFanShuID
}

func (mgr *fanShuManager) getHanderByID(id int) calFSInterface {
	hander, ok := mgr.mapID2Hander[id]
	if !ok {
		return nil
	}
	return hander
}

func (mgr *fanShuManager) getHanderExcept(slExceptChkID []int) []calFSInterface {
	var allHander []calFSInterface
	var slFanShuKey sort.IntSlice
	for key, _ := range mgr.allHander {
		slFanShuKey = append(slFanShuKey, key)
	}

	slFanShuKey.Sort()

	for i := len(slFanShuKey) - 1; i >= 0; i-- {
		key := slFanShuKey[i]
		slHander := mgr.allHander[key]
		for _, hander := range slHander {
			if common.InIntSlace(slExceptChkID, hander.GetID()) {
				continue
			}
			allHander = append(allHander, hander)
		}
	}
	return allHander
}
