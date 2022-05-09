package main

// 合并两个有序数组 虽然目标是递增 但是反向来看就是递减
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := m + n - 1; i >= 0; i-- {
		if n > 0 {
			if m > 0 && nums1[m-1] > nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	nums2 := []int{3, 6, 6}
	merge(nums1, 4, nums2, 3)
}
