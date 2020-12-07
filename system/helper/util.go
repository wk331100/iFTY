package helper

import (
	"bytes"
	"strconv"
	"strings"
)

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


func Implode(glue string, pieces []interface{}) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		if val, ok := str.(int); ok{
			str = strconv.Itoa(val)
		}
		buf.WriteString(str.(string))
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}


func ImplodeString(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}


func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}


func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}