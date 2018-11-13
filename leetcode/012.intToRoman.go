package leetcode

func intToRoman(num int) string {
	t := 0
	r := ""
	for rate := 100; rate > 0; rate /= 10 {
		t = int(num / int(10*rate))
		if t > 0 {
			for i := 0; i < t; i++ {
				switch rate {
				case 100:
					r += "M"
				case 10:
					r += "C"
				case 1:
					r += "X"
				default:
				}
			}
		}
		num -= t * int(10*rate)
		t = int(num / int(9*rate))
		if t > 0 {
			switch rate {
			case 100:
				r += "CM"
			case 10:
				r += "XC"
			case 1:
				r += "IX"
			default:
			}
		}
		num -= t * int(9*rate)
		t = int(num / int(5*rate))
		if t > 0 {
			switch rate {
			case 100:
				r += "D"
			case 10:
				r += "L"
			case 1:
				r += "V"
			default:
			}
		}
		num -= t * int(5*rate)
		t = int(num / int(4*rate))
		if t > 0 {
			switch rate {
			case 100:
				r += "CD"
			case 10:
				r += "XL"
			case 1:
				r += "IV"
			default:
			}
		}
		num -= t * int(4*rate)
	}
	for i := 0; i < num; i++ {
		r += "I"
	}
	return r
}
