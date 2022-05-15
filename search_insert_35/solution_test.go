package search_insert_35

/* 二分法 查找插入的位置 */

func searchInsert(nums []int, target int) int {
  low, high := 0, len(nums)-1
  for low <= high {
    mid := (low + high) / 2
    if nums[mid] < target {
      low = mid + 1
    } else if nums[mid] > target {
      high = mid - 1
    } else {
      return mid
    }
  }
  return low
}
