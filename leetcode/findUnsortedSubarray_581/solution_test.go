func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n==1 || n==0 {
		return 0
	}
	// 先把数组排个序 比较两个数组不同的部分
	sortNums := make([]int, n)
	copy(sortNums, nums)
	sort.Ints(sortNums)

	var left, right int
	for i:=0;i<n;i++ {
		if sortNums[i]!=nums[i] {
			// 核心记住第一个变化的下标的位置
			left = i
			break
		}
	}

	for j:=n-1;j>=0;j-- {
		if sortNums[j]!=nums[j] {
			right = j
			break
		}
	}

	if left==right{
		return 0
	}

	// 举例下标计算
	return right-left+1
}