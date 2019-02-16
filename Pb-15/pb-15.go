//pb-15
package main

import(
  "fmt"
  "time"
)

func latticPaths(n int) int64{
  var x int64 = 0
  x += callThis(0, 0, n)
  return x
}

func callThis(i, j, n int) int64{
  // fmt.Println(i, j)
  var nums int64 = 0
  if i == n && j == n{
    return 1
  }
  if i < n{
    nums += callThis(i + 1, j, n)
  }
  if j < n{
    nums += callThis(i, j + 1, n)
  }
  return nums
}

func main(){
  start := time.Now()

  fmt.Println(latticPaths(18))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
