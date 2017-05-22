package service

import "strconv"

// service 通用方法

// 分页, 排序处理
func parsePageAndSort(pageNumber, pageSize int, sortField string, isAsc bool) (skipNum int, sortFieldR string) {
	skipNum = (pageNumber - 1) * pageSize
	if sortField == "" {
		sortField = "UpdatedTime"
	}
	if !isAsc {
		sortFieldR = "-" + sortField
	} else {
		sortFieldR = sortField
	}
	return
}

// 分页, 排序处理
func parsePage(pageNumber, pageSize int) (skipNum int) {
	skipNum = (pageNumber - 1) * pageSize

	return
}

// IsObjectId 判断id是否是数字
func IsObjectId(id string) bool {
	b, error := strconv.Atoi(id)
	return error == nil
}
