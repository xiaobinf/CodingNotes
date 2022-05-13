package str_str_28

func strStr(haystack string, needle string) int {
  if len(needle) == 0 {
    return 0
  }

  if len(haystack) == 0 || len(needle) > len(haystack) {
    return -1
  }

  for i := 0; i < len(haystack)-len(needle)+1; i++ {
    for j := 0; j < len(needle); j++ {
      if needle[j] != haystack[i+j] {
        break
      }
      if j == len(needle)-1 {
        return i
      }
    }

  }
  return -1
}
