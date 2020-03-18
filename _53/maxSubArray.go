package _53

/*
	对序列进行遍历，如果当前的和为负数或0，则当前序列再延长就没有意义了，可将当前序列值的和置0
	当目前序列和为0时，在当前序列后加上数字，都与不加上当前序列所得的值一样
	当目前力学和为负数时，在当前序列后加上数字，都比加上当前序列所得的值大
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	BIGGEST := nums[0]
	var now int
	for _, i := range nums {
		now += i
		if now > BIGGEST {
			BIGGEST = now
		}
		if now <= 0 {
			now = 0
			continue
		}
	}
	return BIGGEST
}
