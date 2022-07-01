package main

import "fmt"

func main() {
	var integers = []int{-8, 0, 1, 3, 7, 18, 74, 101, 223, 395, 683, 911}

	var key = 3
	var found bool = BinarySearchIteratively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 11
	found = BinarySearchIteratively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 395
	found = BinarySearchRecursively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 200
	found = BinarySearchRecursively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)
}

func BinarySearchIteratively(array []int, key int) bool {
	var found bool = false
	var length int = len(array)
	var start int = 0
	var end int = length - 1

	if key < array[start] {
		return false
	}
	if key > array[end] {
		return false
	}

	for found != true {
		var middle int = (end + start) / 2
		var element int = array[middle]
		if key == element {
			found = true
			break
		} else if key < element {
			end = middle - 1
		} else if key > element {
			start = middle + 1
		}

		if start > end {
			break
		}
	}
	return found
}

func BinarySearchRecursively(array []int, key int) bool {
	var found bool = false
	var length int = len(array)
	var start int = 0
	var end int = length - 1

	if key < array[start] {
		return false
	}
	if key > array[end] {
		return false
	}

	var middle int = (end + start) / 2
	var element int = array[middle]
	if key == element {
		return true
	} else {
		if key < element {
			end = middle - 1
		} else if key > element {
			start = middle + 1
		}
		slice := array[start : end+1]
		found = BinarySearchRecursively(slice, key)
	}
	return found
}
