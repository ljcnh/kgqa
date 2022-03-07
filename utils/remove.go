package utils

import "regexp"

// RemoveRep 元素去重
func RemoveRep(slc []string) []string {
	// 切片长度小于256的时候，循环来过滤
	if len(slc) < 256 {
		return removeRepByLoop(slc)
	} else {
		return removeRepByMap(slc)
	}
}

// RemoveSpaceAndTab 去除空白和制表符
func RemoveSpaceAndTab(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

// 通过两重循环过滤重复元素
func removeRepByLoop(slc []string) []string {
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []string) []string {
	var result []string
	tempMap := make(map[string]struct{})
	for _, e := range slc {
		if _, ok := tempMap[e]; !ok {
			tempMap[e] = struct{}{}
			result = append(result, e)
		}
	}
	return result
}
