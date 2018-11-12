package leetcode

func maxArea(height []int) int {
	var max = 0
	var l, h = 0, 0
	for i, _ := range height {
		for j := len(height) - 1; j > i; j-- {
			l = j - i
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			if l*h > max {
				max = l * h
			}
		}
	}
	return max
}
