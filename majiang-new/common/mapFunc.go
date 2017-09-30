package common

import (
	"reflect"
)

func GetUInt8MapKey(mapData map[uint8]bool) (slKey []uint8) {
	for key, _ := range mapData {
		slKey = append(slKey, key)
	}
	return slKey
}

//map或者切片深度拷贝，注意内部是指针(尚不支持指针)
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap

	} else if valueMap, ok := value.(map[int]interface{}); ok {
		newMap := make(map[int]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}

	return value
}

func GetMapKeys(maps interface{}) interface{} {
	value := reflect.ValueOf(maps)
	if value.Kind() != reflect.Map {
		return nil
	}
	setKeyValue := value.MapKeys()
	if len(setKeyValue) == 0 {
		return nil
	}
	kind := setKeyValue[0].Kind()

	var setKeys interface{}
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		setKeys = make([]int, len(setKeyValue))
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		setKeys = make([]uint, len(setKeyValue))
	case reflect.String:
		setKeys = make([]string, len(setKeyValue))
	default:
		return nil
	}

	for iIndex, key := range setKeyValue {
		switch key.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			setKeys.([]int)[iIndex] = int(key.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			setKeys.([]uint)[iIndex] = uint(key.Uint())
		case reflect.String:
			setKeys.([]string)[iIndex] = key.String()
		}
	}
	return setKeys
}
