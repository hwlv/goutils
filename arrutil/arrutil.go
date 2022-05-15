package arrutil

func Contains[T comparable](arr []T, e T) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}

func RemoveDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		if Contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}
