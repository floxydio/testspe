package blueocean

func BlueOcean(list1 []int, list2 []int) []int {
    filterMap := make(map[int]bool)
    for _, num := range list2 {
        filterMap[num] = true
    }
    filteredList := []int{}
    for _, num := range list1 {
        if !filterMap[num] {
            filteredList = append(filteredList, num)
        }
    }

    return filteredList
}