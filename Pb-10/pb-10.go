//#pb-7
package main

import(
  "fmt"
  "time"
)

func primeSummation(n int) int{
  var MAX = n
  var sum int = 0
  var arr  = make([]bool, n)
  var count int = 0
  arr[0], arr[1] = true, true
  for i := 0; i < MAX; i++{
    //fmt.Println(i)
    if !arr[i]{
      sum += i
      for j, k := i * 2, 0; k < ((MAX - 1)/i) - 1; k++{
        arr[j] = true
        j += i
      }
      count += 1
    }
//    if count == n{
//      return sum
//    }
  }
  return sum
}

func main(){
  start := time.Now()

  fmt.Println(primeSummation(2000000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)

  // 18.74ms
  // 5.5ms --now
}

