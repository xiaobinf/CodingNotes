package longest_common_prefix_14

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
  var commonStr = ""
  for i := 0; i < len(strs[0]); i++ {
    commonStr += string(strs[0][i])
    for j := 1; j < len(strs); j++ {
      if i > len(strs[j])-1 || strs[j][i] != strs[0][i] {
        return commonStr[:i]
      }
    }
  }
  return commonStr
}

func longestCommonPrefix1(strs []string) string {
  if len(strs) == 0 {
    return ""
  }
  if len(strs) == 1 {
    return strs[0]
  }
  for i := 0; i < len(strs[0]); i++ {
    for j := 1; j < len(strs); j++ {
      if i > len(strs[j])-1 || strs[j][i] != strs[0][i] {
        return strs[0][:i]
      }
    }
  }
  return strs[0]
}
