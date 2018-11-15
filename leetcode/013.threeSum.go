package leetcode

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	var arrResults [][]int
	var am = make(map[string]string)
	// 先排序
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	// 分组
	positiveInteger, negtiveInteger, zero := sliceint(nums)
	if len(zero) >= 3 {
		var tmpArr []int
		tmpArr = append(tmpArr, 0)
		tmpArr = append(tmpArr, 0)
		tmpArr = append(tmpArr, 0)
		arrResults = append(arrResults, tmpArr)
	}
	var ps, pe = 0, len(positiveInteger) - 1
	var ns, ne = 0, len(negtiveInteger) - 1
	var a, b = 0, 0

	for ps <= pe && ns <= ne {
		if positiveInteger[pe]+negtiveInteger[ns] >= 0 {
			a = positiveInteger[pe]
			for i := ne; i >= ns; i-- {
				b = negtiveInteger[i]
				if a+b == 0 {
					if len(zero) >= 1 {
						a1, a2, a3 := sort3(a, b, 0)
						str := fmt.Sprintf("%d,%d,%d", a1, a2, a3)
						_, ok := am[str]
						if ok {
							continue
						}
						var tmpArr []int
						tmpArr = append(tmpArr, a)
						tmpArr = append(tmpArr, b)
						tmpArr = append(tmpArr, 0)
						arrResults = append(arrResults, tmpArr)
						am[str] = str
					}
					continue
				}
				for j := 0; j <= ne; j++ {
					if i == j {
						continue
					}
					if a+b+negtiveInteger[j] == 0 {
						a1, a2, a3 := sort3(a, b, negtiveInteger[j])
						str := fmt.Sprintf("%d,%d,%d", a1, a2, a3)
						_, ok := am[str]
						if ok {
							break
						}
						var tmpArr []int
						tmpArr = append(tmpArr, a)
						tmpArr = append(tmpArr, b)
						tmpArr = append(tmpArr, negtiveInteger[j])
						arrResults = append(arrResults, tmpArr)
						am[str] = str
						break
					} else if a+b+negtiveInteger[j] > 0 {
						break
					}
				}
			}

		}
		if positiveInteger[pe]+negtiveInteger[ns] <= 0 {
			a = negtiveInteger[ns]
			for i := ps; i <= pe; i++ {
				b = positiveInteger[i]
				if a+b == 0 {
					if len(zero) >= 1 {
						a1, a2, a3 := sort3(a, b, 0)
						str := fmt.Sprintf("%d,%d,%d", a1, a2, a3)
						_, ok := am[str]
						if ok {
							continue
						}
						var tmpArr []int
						tmpArr = append(tmpArr, a)
						tmpArr = append(tmpArr, b)
						tmpArr = append(tmpArr, 0)
						arrResults = append(arrResults, tmpArr)
						am[str] = str
					}
					continue
				}
				for j := 0; j <= pe; j++ {
					if i == j {
						continue
					}
					if a+b+positiveInteger[j] == 0 {
						a1, a2, a3 := sort3(a, b, positiveInteger[j])
						str := fmt.Sprintf("%d,%d,%d", a1, a2, a3)
						_, ok := am[str]
						if ok {
							break
						}
						var tmpArr []int
						tmpArr = append(tmpArr, a)
						tmpArr = append(tmpArr, b)
						tmpArr = append(tmpArr, positiveInteger[j])
						arrResults = append(arrResults, tmpArr)
						am[str] = str
						break
					} else if a+b+positiveInteger[j] > 0 {
						break
					}
				}
			}

		}
		if positiveInteger[pe]+negtiveInteger[ns] == 0 {
			ns++
			pe--
		} else if positiveInteger[pe]+negtiveInteger[ns] > 0 {
			pe--
		} else if positiveInteger[pe]+negtiveInteger[ns] < 0 {
			ns++
		}
	}
	return arrResults
}
func sort3(a, b, c int) (a1, a2, a3 int) {
	nums := []int{a, b, c}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] < nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums[0], nums[1], nums[2]
}

func sliceint(src []int) (positiveInteger, negtiveInteger, zero []int) {
	start0 := len(src)
	end := len(src)
	for i := 0; i < len(src); i++ {
		if src[i] >= 0 {
			start0 = i
			break
		}
	}
	negtiveInteger = src[0:start0]
	for i := start0; i < len(src); i++ {
		if src[i] > 0 {
			end = i
			break
		}
	}
	positiveInteger = src[end:len(src)]
	if start0 < end {
		zero = src[start0:end]
	}
	return
}
