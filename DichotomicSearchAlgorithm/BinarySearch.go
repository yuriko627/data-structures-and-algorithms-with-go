package main

import "fmt"

func main() {
	var scores = []int{30, 40, 50, 70, 85, 90, 100}
	var length = len(scores)

	// value to find
	var searchValue = 40

	var position = binarySearch(scores, length, searchValue)
	fmt.Printf("%d position: %d", searchValue, position)

	fmt.Printf("\n-----------------------------\n")

	searchValue = 90
	position = binarySearch(scores, length, searchValue)
	fmt.Printf("%d position: %d", searchValue, position)
}

func binarySearch(arrays []int, length int, searchValue int) int {
	//initialize the smallest and largest index
	var low = 0
	var high = length
	var mid = 0

	for {
		//end the program when it does not find the search value in the array
		if low >= high {
			break
		}

		// set the middle index
		mid = (low + high) / 2

		if arrays[mid] == searchValue {
			return mid //found the position for the serachValue
		} else if arrays[mid] < searchValue {
			low = mid + 1 //starts searching in the right half
		} else if arrays[mid] > searchValue {
			high = mid - 1 //starts searching in the left half
		}

	}

	return -1
}
