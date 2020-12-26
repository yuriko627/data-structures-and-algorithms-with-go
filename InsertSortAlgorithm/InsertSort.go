package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length; i++ {
		var insertElement = arrays[i]
		var insertPosition = i
		for j := insertPosition - 1; j >= 0; j-- {
			if insertElement < arrays[j] {
				//shift the insertElement to the right
				arrays[j+1] = arrays[j]
				insertPosition--
			}
		}
		//inser a new element
		arrays[insertPosition] = insertElement
	}
}
