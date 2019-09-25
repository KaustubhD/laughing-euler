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
      for j, k := i * 2, 0; k < ((MAX - 1)/i) - 1; k++{
        arr[j] = true
        j += i
      }
      count += 1
    }
    if count == n{
      return i
    }
  }
  return 0
}

func main(){
  start := time.Now()

  fmt.Println(nthPrime(10001))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)

  // 18.74ms
  // 5.5ms --now
}
