package sort

import "fmt"

func quickSort(arr []int, left, right int) {
    if left < right {
        pivot := partition(arr, left, right)
        quickSort(arr, left, pivot-1)
        quickSort(arr, pivot+1, right)
    }
}

func partition(arr []int, left, right int) int {
    pivot := arr[right]
    i := left - 1
    // 把所有比pivot小的数放在左侧，涉及到交换
    for j := left; j < right; j++ {
        if arr[j] < pivot {
            // i的步数是小于pivot的元素个数
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[right] = arr[right], arr[i+1]
    return i + 1
}

func main() {
    arr := []int{5, 2, 6, 1, 3, 9, 4, 8, 7}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}

