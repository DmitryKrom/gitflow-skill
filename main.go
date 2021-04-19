package main

import "fmt"

func main() {
	var listOfNotes []int = []int{1, 1, 5, 5, 3, 4, 3, 1, 1, 5, 5, 5, 5, 3, 3, 4, 3, 4, 4, 1}
	sortNotes(listOfNotes)
}

func sortNotes(list []int) {
	var statisticMap = make(map[int]int)
	for i := 0; i < len(list); i++ {
		var counter int = 1

		if _, ok := statisticMap[list[i]]; !ok {
			for j := i + 1; j < len(list); j++ {
				if list[j] == list[i] {
					counter++
				}
			}
			statisticMap[list[i]] = counter
		}

	}
	fmt.Println(statisticMap)
}
