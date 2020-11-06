package _528_ShuffleString

import "strings"

func RestoreString(s string, indices []int) string {
  newArray := make([]string, len(indices))
  str := []rune(s)

  for i := 0; i < len(str); i++ {
    newArray[indices[i]] = string(str[i])
  }

  return strings.Join(newArray, "")
}
