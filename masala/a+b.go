package main

import (
	"container/list"
	"fmt"
)

func main() {
// 	var a, b string
// 	fmt.Scan(&a, &b)
// 	fmt.Println(a+b)
// }
// func sum(a, b string)(res int64){
// 	var s []string = a[:]
// 	var c []string = b[:]
// 	var n, m int = len(a), len(b)
// 	if n < m {
// 		n = m
// 	}
// 	var d, q int 0, 0
// 	for i := n-1; i >= 0; i --{

// 	}
var ele *list.Element
 
  // Create list and insert elements in it.
  list_1 := list.New()
 
  list_1.PushBack(1) // 1
  list_1.PushBack(2) // 1 -> 2
  list_1.PushBack(3) // 1 -> 2 -> 3
  list_1.PushBack(4) // 1 -> 2 -> 3 -> 4
 
 
  fmt.Println("Print list before inserting element")
  for ele = list_1.Front(); ele != nil; ele = ele.Next() {
 
    fmt.Println(ele.Value)
     
  }
 
  list_1.PushFront(5)
  // Now list: 5 -> 1 -> 2 -> 3 -> 4
 
  fmt.Println("Print list after inserting element")
  for ele = list_1.Front(); ele != nil; ele = ele.Next() {
 
    fmt.Println(ele.Value)
     
  }
}