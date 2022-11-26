package book

import "strconv"

func getFilterList(data map[string]string) (int, int, map[string]interface{}) {
	page := -1
	size := -1
	pageValue, err := strconv.ParseInt(data["page"], 10, 32)
	if err == nil {
		page = int(pageValue)
	}

	sizeValue, err := strconv.ParseInt(data["size"], 10, 32)
	if err == nil {
		size = int(sizeValue)
	}

	filter := make(map[string]interface{}, len(data))
	for k, v := range data {
		filter[k] = v
	}

	return page, size, filter
}
