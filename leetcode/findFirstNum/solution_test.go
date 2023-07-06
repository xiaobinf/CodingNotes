package findFirstNum

func findFirst(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid]-nums[mid-1] != 1 {
            return nums[mid]
        }
        if nums[mid]-nums[0] == mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return nums[left]
}