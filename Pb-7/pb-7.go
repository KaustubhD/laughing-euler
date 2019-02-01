//#pb-7
package main

import(
  "fmt"
  "time"
)

func nthPrime(n int) int{
  const MAX int = 150000
  var arr [MAX]bool
  var count int = 0
  arr[0], arr[1] = true, true
  for i := 0; i < MAX; i++{
    if !arr[i]{
      fmt.Println("Found false", i, arr[i], count)
      for j := i * 2; j < ((MAX - 1)/i) - 1; j += i{
//        fmt.Println("j is",j)
        arr[j] = true
      }
      count += 1
    }
//    fmt.Println(arr)
    if count == n{
      fmt.Println("count", count, "index", i)
      return i
    }
  }
//  fmt.Println(count)
  return 0
}

func main(){
  start := time.Now()

  fmt.Println(nthPrime(1000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)

  // 18.74ms
  // 370.173µs -- this time
}
