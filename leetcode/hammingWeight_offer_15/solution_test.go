func hammingWeight(num uint32) int {
	count := 0
	// 位运算 移位 逐个比较
	for i:=0;i<32;i++{
		if 1<<i&num > 0 {
			count ++
		}
	}

	return count
}

&运算符是位运算符之一，表示按位与运算。它对两个二进制数的每个位进行比较，只有当两个二进制数的对应位都为1时，结果的对应位才为1，否则为0。