package stats

import "sort"

func Sum(nums []float64) float64 {
	var total float64
	for _, n := range nums {
		total += n
	}
	return total
}

func Mean(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	return Sum(nums) / float64(len(nums))
}

func Median(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	cp := append([]float64(nil), nums...) // copy so we don’t mutate input
	sort.Float64s(cp)
	n := len(cp)
	mid := n / 2
	if n%2 == 1 {
		return cp[mid]
	}
	return (cp[mid-1] + cp[mid]) / 2
}

func Min(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func Max(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
