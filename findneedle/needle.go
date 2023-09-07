package findneedle

func FindNeedle(list []string, search string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == search {
			return i
		}
	}
	return -1
}
