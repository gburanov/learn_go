package main

import (
  "fmt"
  "strconv"
)

type ListElement struct {
    Value int
    Next *ListElement
}

func (e *ListElement) String() string {
  ret := strconv.Itoa(e.Value)
  if e.Next != nil {
    ret += " " + e.Next.String()
  }
  return ret
}

func New(arr []int) *ListElement {
  var previous *ListElement
  var next *ListElement
  var first *ListElement

  for _, val := range arr {
    if (previous == nil) {
      previous = &ListElement{Value: val}
      first = previous
    } else {
      next = &ListElement{Value: val}
      previous.Next = next
      previous = next
    }
  }
  return first
}

func removeDuplicatesRecursion(start *ListElement) {
  if start == nil {
      return
  }

  next_element := start.Next

  if next_element == nil {
      return
  }

  if start.Value == next_element.Value {
      new_next := next_element.Next
      if new_next == nil {
          start.Next = nil
          return
      }
      start.Next = new_next
      removeDuplicatesRecursion(start)
      return
  }
  removeDuplicatesRecursion(start.Next)
}

func removeDuplicatesIteration(start *ListElement) {
  if start == nil {
      return
  }

  for elem := start; elem != nil;  {
      next_element := elem.Next
      if next_element == nil {
          break
      }
      if elem.Value == next_element.Value {
          new_next := next_element.Next
          if (new_next == nil) {
              elem.Next = nil
              break
          }
          elem.Next = new_next
      } else {
          elem = elem.Next
      }
  }
}

func checkList(arr []int) {
  list := New(arr)
  fmt.Println(list)
  removeDuplicatesIteration(list)
  fmt.Println(list)

  list = New(arr)
  fmt.Println(list)
  removeDuplicatesRecursion(list)
  fmt.Println(list)
}

func main() {
  checkList(nil)
  checkList([]int{1, 1, 1, 1, 1, 1, 1})
  checkList([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2})
  checkList([]int{1, 1, 1, 3, 3, 4, 5})
  checkList([]int{1, 1, 1, 3, 3, 4, 5, 5, 5})
}
