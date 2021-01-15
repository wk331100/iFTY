package helper

import "strconv"

type Map map[string]interface{}

func (m *Map) String(str string) string{
	mapStruct := *m
	if mapStruct[str] != nil {
		return mapStruct[str].(string)
	}
	return ""
}

func (m *Map) Int(str string) int{
	mapStruct := *m
	if mapStruct[str] != nil {
		int, _ :=strconv.Atoi(mapStruct[str].(string))
		return int
	}
	return 0
}

func (m *Map) Bool(str string) bool{
	mapStruct := *m
	if mapStruct[str] != nil {
		return true
	}
	return false
}

func (m *Map) Interface(str string) interface{}{
	mapStruct := *m
	if mapStruct[str] != nil {
		return mapStruct[str]
	}
	return nil
}
