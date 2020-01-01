package main

import "fmt"

// 插入排序
func insertSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	for i := 1; i < len(s); i++ {
		v := s[i]
		j := i - 1
		for ; j >= 0; j-- {
			if s[j] > v {
				s[j+1] = s[j]
			} else {
				break
			}
		}
		s[j+1] = v
	}
	return s
}

// 选择排序
func selectSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}

	return s
}

// 冒泡排序
func bubbleSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		flag := false

		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return s
}

//归并排序
func merge(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]
	return mergeSort(merge(left), merge(right))
}

func mergeSort(s1, s2 []int) []int {
	var s []int

	for len(s1) > 0 && len(s2) > 0 {
		if s1[0] > s2[0] {
			s = append(s, s2[0])
			s2 = s2[1:]
		} else {
			s = append(s, s1[0])
			s1 = s1[1:]
		}
	}

	for len(s1) > 0 {
		s = append(s, s1[0])
		s1 = s1[1:]
	}

	for len(s2) > 0 {
		s = append(s, s2[0])
		s2 = s2[1:]
	}
	return s
}

//快速排序
func quickSort(s []int) []int {
	quick(s, 0, len(s)-1)
	return s
}

func quick(s []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(s, start, end)
	quick(s, start, i-1)
	quick(s, i+1, end)
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	s := []int{2, 0, 14, 5, 6, 3, 4, 11, 1, 8, 2,2}
	//fmt.Println(insertSort(s))
	//fmt.Println(selectSort(s))
	//fmt.Println(bubbleSort(s))
	//fmt.Println(merge(s))
	fmt.Println(quickSort(s))
}
