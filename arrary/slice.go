package arrary

import "reflect"

// InMap check index in Slice
// 判断是否在切片中存在，开辟一个Map 速度更快点
func InMap(needle string, stack []string) bool {
	newStack := map[string]bool{}

	for _, value := range stack {
		newStack[value] = true
	}

	if _, ok := newStack[needle]; ok {
		return true
	}
	return false
}

// InSlice check index in slice
// 判断是否在切片中存在，直接遍历寻找
func InSlice(needle string, stack []string) bool {
	for _, value := range stack {
		if value == needle {
			return true
		}
	}
	return false
}

// InSliceInt64 check index in slice int64
// 判断是否在切片中存在，直接遍历寻找
func InSliceInt64(needle int64, stack []int64) bool {
	for _, value := range stack {
		if value == needle {
			return true
		}
	}
	return false
}

// InSliceInt check index in slice int
// 判断是否在切片中存在，直接遍历寻找
func InSliceInt(needle int, stack []int) bool {
	for _, value := range stack {
		if value == needle {
			return true
		}
	}
	return false
}

// InArray index in array for interface{}
// 判断是否存在并返回下标地址
// 不存在返回-1
func InArray(needle interface{}, array []interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// Diff show difference in two array.
// 两个切片的不同的元素
func Diff(s, t []string) []string {
	slice1 := make([]string, len(s))
	slice2 := make([]string, len(t))
	copy(slice1, s)
	copy(slice2, t)
	v := []string{}
	if len(slice1) == 0 && len(slice2) == 0 {
		return []string{}
	}

	if len(slice1) == 0 {
		return slice2
	}

	if len(slice2) == 0 {
		return slice1
	}

	if len(slice1) > len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	for _, val := range slice1 {
		if newT, ok := In(val, slice2); ok {
			slice2 = newT
			continue
		}

		v = append(v, val)
	}

	if len(slice2) > 0 {
		v = append(v, slice2...)
	}

	return v
}

// In check string in array.
func In(needle string, haystack []string) ([]string, bool) {
	newHaystack := make([]string, len(haystack))
	copy(newHaystack, haystack)

	if len(newHaystack) == 0 {
		return newHaystack, false
	}

	for i, val := range newHaystack {
		if val == needle {
			newHaystack = append(newHaystack[:i], newHaystack[i+1:]...)
			return newHaystack, true
		}
	}

	return newHaystack, false
}
