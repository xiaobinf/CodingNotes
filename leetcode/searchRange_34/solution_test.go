func searchRange(nums []int, target int) []int {
	start,end := -1,-1
	n:=len(nums)
	left,right := 0,n-1
	var mid int
	for left<=right{
		mid = (left+right)/2
		if nums[mid] == target {
			// 左右扩展
			start, end = mid, mid
			// 注意边界 -1可能越界
			for start>0 && nums[start-1] == target {
				start --
			}

			for end<n-1 && nums[end+1] == target {
				end ++
			}

			break
		} else if nums[mid] < target {
			left ++
		} else {
			right --
		}
	}

	return []int{start, end}
}