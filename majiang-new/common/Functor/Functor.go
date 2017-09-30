package Functor

import (
	log "github.com/alecthomas/log4go"
)

type Functor struct {
	function interface{}
	args     []interface{}
}

func checkFunc(function interface{}) bool {
	_, ok1 := function.(func(...interface{}))
	_, ok2 := function.(func(interface{}))
	_, ok3 := function.(func())
	if ok1 || ok2 || ok3 {
		return true
	}
	return false
}

func GetFunctor(function interface{}, args ...interface{}) *Functor {
	if !checkFunc(function) {
		log.Error("the first arg is not function")
		return nil
	}
	myFuncor := new(Functor)
	myFuncor.function = function
	myFuncor.args = args
	return myFuncor
}

//只有自己在函数里面解析了
func (functor *Functor) RunFunc() {
	iArgsLen := len(functor.args)
	if iArgsLen > 1 {
		functor.function.(func(...interface{}))(functor.args)
	} else if iArgsLen == 1 {
		functor.function.(func(interface{}))(functor.args[0])
	} else {
		functor.function.(func())()
	}
}
