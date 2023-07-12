package findRepeatNumber_offer_03
// 从下标寻找规律
func findRepeatNumber(nums []int) int {
	var ret int
    n := len(nums)
	for i:=0;i<len(nums); i++ {
		nums[nums[i]%n] += n
	}

    for i:=0;i<len(nums); i++ {
        if nums[i] >=2*n {
            ret = i
            break
        }
	}
    
	return ret
}