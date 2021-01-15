package helper

import (
	"bytes"
	"strconv"
	"strings"
)

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

// 取两个切片的交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

//取要校验的和已经校验过的差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string,0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			n = append(n, v)
		}
	}
	return n
}

//取出slice1中有，slice2中没有的
func DifferenceDel(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string,0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}
	return n
}