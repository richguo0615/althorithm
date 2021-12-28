package main

import "math/rand"

//排序法 http://notepad.yehyeh.net/Content/Algorithm/Sort/Sort.php

func bubbleSort(arr []int) []int {
	end := len(arr)
	for end != 0 {
		for j := 0; j < end-1; j++ {
			k := j+1
			if arr[j] > arr[k] {
				swap(arr, j, k)
			}
		}
		end--
	}
	return arr
}

func insertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		curr := arr[i]

		insert := false
		for j, n := range res {
			if curr < n {
				tmp := make([]int, len(res[j:]))
				copy(tmp, res[j:])
				res = append(res[:j], curr)
				res = append(res, tmp...)
				insert = true
				break
			}
		}

		if !insert {
			res = append(res, curr)
		}
	}

	return res
}

func selectionSort(arr []int) []int {
	idx := 0
	for idx < len(arr) {
		minIdx := idx
		for i := idx; i < len(arr); i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
		}
		swap(arr, idx, minIdx)
		idx++
	}

	return arr
}

//https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
//func quickSort(arr []int) []int {
//	if len(arr) < 2 {
//		return arr
//	}
//
//	pivot := rand.Int() % len(arr)
//
//	left, right := make([]int, 0), make([]int, 0)
//
//	for i := 0; i < len(arr); i++ {
//		if i == pivot {
//			continue
//		}
//		if arr[i] < arr[pivot] {
//			left = append(left, arr[i])
//		} else {
//			right = append(right, arr[i])
//		}
//	}
//
//	tmp := append(left, arr[pivot])
//	tmp = append(tmp, right...)
//
//	res := append(quickSort(tmp[:len(left)]), arr[pivot])
//	res = append(res, quickSort(tmp[len(left)+1:])...)
//
//	return res
//}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := rand.Int() % len(arr)
	left, right := 0, len(arr)-1

	arr[pivot], arr[right] = arr[pivot], arr[right]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr)/2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	arr := make([]int, len(left)+len(right))

	var i, j, idx int
	for {
		if i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr[idx] = left[i]
				i++
			} else {
				arr[idx] = right[j]
				j++
			}
		} else if i < len(left) {
			arr[idx] = left[i]
			i++
		} else if j < len(right) {
			arr[idx] = right[j]
			j++
		} else {
			break
		}
		idx++
	}

	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
