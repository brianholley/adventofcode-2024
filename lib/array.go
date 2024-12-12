package lib

import (
	"sort"
)

func ArrayMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func ArrayMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func ArraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func ArrayLast[K comparable](arr []K) K {
	return arr[len(arr)-1]
}

func ArrayRemoveIndex[K comparable](arr []K, index int) []K {
	if index == -1 {
		return arr
	} else if index == len(arr)-1 {
		return arr[:len(arr)-1]
	} else {
		return append(arr[:index], arr[index+1:]...)
	}
}

func ArrayRemoveItem[K comparable](arr []K, value K) []K {
	index := ArrayIndexOf(arr, value)
	if index == -1 {
		return arr
	} else if index == len(arr)-1 {
		return arr[:len(arr)-1]
	} else {
		return append(arr[:index], arr[index+1:]...)
	}
}

func ArrayContains[K comparable](arr []K, value K) bool {
	return ArrayIndexOf(arr, value) != -1
}

func ArrayIndexOf[K comparable](arr []K, value K) int {
	for i := range arr {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

func ArrayCopy[K comparable](arr []K) []K {
	c := make([]K, len(arr))
	copy(c, arr)
	return c
}

func Array2dCopy[K comparable](arr [][]K) [][]K {
	c := make([][]K, len(arr))
	for i := range arr {
		c[i] = make([]K, len(arr[i]))
		copy(c[i], arr[i])
	}
	return c
}

// Should be able to use [K Ordered] constraint here
func ArraySortAscending(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
func ArraySortDescending(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}
