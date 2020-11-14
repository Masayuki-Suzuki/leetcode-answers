package _221_SplitAStringInBalancedStrings

func BalancedStringSplit(s string) int {
  count, pos := 0, 0

  for i, chr := range s {

    if chr == 'L' {
      pos -= 1
    } else if chr == 'R' {
      pos += 1
    }

    if i > 0 && pos == 0 {
      count++
    }
  }

  return count
}
