package tool

import "strconv"

// GetAllTag 拼接一个or条件的字符串去查询标签的信息
func GetAllTag(tagList []int) string {
	tagStrList := ""
	for i := 0; i < len(tagList); i++ {
		tag := strconv.Itoa(tagList[i])
		if i != len(tagList)-1 {
			tagStrList = tagStrList + "id = " + tag + " or "
		} else {
			tagStrList = tagStrList + "id = " + tag
		}
	}
	return tagStrList
}
