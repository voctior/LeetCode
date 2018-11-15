package leetcode

func threeSum(nums []int) [][]int {
	var arrResults [][]int
	// 先排序
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	// 分组
	var positiveInteger []int
	var negtiveInteger []int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			positiveInteger = nums[i:len(nums)]
			negtiveInteger = nums[0:i]
			break
		}
	}
	var pe = len(positiveInteger) - 1
	var ns = 0
	var a, b = 0, 0
	for pe >= 0 && ns < len(negtiveInteger) {
		var tmpArr []int
		a = positiveInteger[pe]
		b = negtiveInteger[ns]
		if a+b > 0 {
			for i := range negtiveInteger {
				if a+b+negtiveInteger[i] == 0 {
					// TODO
					tmpArr = append(tmpArr, a)
					tmpArr = append(tmpArr, b)
					tmpArr = append(tmpArr, negtiveInteger[i])
					arrResults = append(arrResults, tmpArr)
					pe--
				} else if a+b+negtiveInteger[i] < 0 {
					pe--
					break
				}
			}
		} else if a+b < 0 {
			for i := range positiveInteger {
				if a+b+positiveInteger[i] == 0 {
					// TODO
					tmpArr = append(tmpArr, a)
					tmpArr = append(tmpArr, b)
					tmpArr = append(tmpArr, positiveInteger[i])
					arrResults = append(arrResults, tmpArr)
					ns++
				} else if a+b+positiveInteger[i] > 0 {
					ns++
					break
				}
			}
		} else if a+b == 0 {
			ns++
			pe--
		}
	}
	return arrResults
}
