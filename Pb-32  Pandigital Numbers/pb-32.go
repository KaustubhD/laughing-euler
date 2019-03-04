//pb-32
package main

import(
  "fmt"
  "time"
)

func sumDivisors(n int) int{
  var sum int = 0
  for i := 1; i < n; i++{
    if n % i == 0{
      sum += i
    }
  }
  return sum
}

func sumAmicableNum(n int) int{
  var sum, y int = 0, 0
  for x := 1; x <= n; x++{
    y = sumDivisors(x)
    if x < y && sumDivisors(y) == x{
      sum += x + y
    }
  }
  return sum
}

func main(){
  start := time.Now()

  fmt.Println(sumAmicableNum(10000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
