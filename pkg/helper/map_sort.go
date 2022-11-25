package helper

import "strings"

func CheckSortBy(sortBy string) string {
	switch strings.ToUpper(sortBy) {
	case "ASC":
		sortBy = "ASC"
	case "DESC":
		sortBy = "DESC"
	default:
		sortBy = "DESC"
	}

	return sortBy
}
