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
	printMostResult(statisticMap)
	printLeastResult(statisticMap)
}
func printMostResult(statisticMap map[int]int) {
	var maxRank int
	var bestNote int
	for key, val := range statisticMap {
		if val > maxRank {
			maxRank = val
			bestNote = key
		}
	}
	fmt.Printf("Оценка %d получена %d раз(а)\n", bestNote, maxRank)
}
func printLeastResult(statisticMap map[int]int) {
	var maxRank int = 1000
	var worstNote int
	for key, val := range statisticMap {
		if val < maxRank {
			maxRank = val
			worstNote = key
		}
	}
	fmt.Printf("Оценка %d получена %d раз(а)\n", worstNote, maxRank)
}
