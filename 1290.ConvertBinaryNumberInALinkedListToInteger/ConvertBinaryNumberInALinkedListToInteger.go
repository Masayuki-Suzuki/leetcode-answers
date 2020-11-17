package _290_ConvertBinaryNumberInALinkedListToInteger

import (
  "math"
)

// **
// * Definition for singly-linked list.
type ListNode struct {
  Val  int
  Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
  binaries := make([]int, 0)
  current := head
  result := 0.0

  for current != nil {
    binaries = append(binaries, current.Val)
    current = current.Next
  }

  length := len(binaries)

  for i, c := length-1, 0; i >= 0 && c < length; i, c = i-1, c+1 {
    if binaries[i] > 0 {
      result = result + math.Pow(2, float64(c))
    }
  }

  return int(result)
}

// Answers in discussion
//func getDecimalValue(head *ListNode) int {
//  curr, result := head, 0
//  for curr != nil {
//    result = result<<1
//    result += curr.Val
//    curr = curr.Next
//  }
//  return result
//}

//func getDecimalValue(head *ListNode) int {
//  sum:=0
//  for head!=nil{
//    fmt.Println(sum, head.Val, sum*2+head.Val)
//    sum = sum*2+head.Val
//    head = head.Next
//  }
//  return sum
//}
