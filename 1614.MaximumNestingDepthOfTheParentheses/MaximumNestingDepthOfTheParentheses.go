package _614_MaximumNestingDepthOfTheParentheses

func MaxDepth(s string) int {
  ans, nest := 0, 0

  for _, c := range s {
    if c == '(' {
      nest++
      if ans > nest {
        ans = ans
      } else {
        ans = nest
      }
    } else if c == ')' {
      nest--
    }
  }

  return ans
}
