package main

import (
  "fmt"
  "time"
)

func sumSquareDifference(n int) int{
  var squaredSum int = n * (n + 1) / 2
  var sumOfSquared int= n * (n + 1) * (2 * n + 1) / 6
  return sumOfSquared - squaredSum
}

func main(){
  start := time.Now()

  fmt.Println(sumSquareDifference(10))

  elapsed := time.Since(start)
  fmt.Println("Time taken %s\n", elapsed)
}
