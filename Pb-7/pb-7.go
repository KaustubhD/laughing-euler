//#pb-7
package main

import(
  "fmt"
  "time"
  "math"
)

func nthPrime(n int) int{
  const MAX int = 15000
  var arr [MAX]bool
  var count int = 0
  arr[0], arr[1] = true, true
  for ind, val := range arr{
    if !val{
      for i := ind * 2; i <= ((MAX - 1)/ind - 1); i += ind{
        count += 1
        arr[i] = true
      }
    }
    if count == n{
      return ind
    }
  }
  return n
}

func isPrime(n int) bool{
  if n % 2 == 0{
   return false
  }else{
    for i := 3; i <= int(math.Sqrt(float64(n))); i += 2{
      if n % i == 0{
        return false
      }
    }
  }
  return true
}


func main(){
  start := time.Now()

  fmt.Println(nthPrime(10001))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)

  // 18.74ms
  // 370.173µs -- this time
}
