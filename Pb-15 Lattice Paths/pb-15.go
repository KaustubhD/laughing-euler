//pb-15
package main

import(
  "fmt"
  "time"
)

func latticPaths(n int) int{
  arr := make([]int, n + 1)
  arr[0] = 1
  for i := 1; i <= n; i++{
    arr[i] = 1
    for j := 1; j < i; j++{
      arr[j] += arr[j- 1]
    }
    arr[i] = arr[i - 1] * 2
  }

  return arr[n]
}


func main(){
  start := time.Now()

  fmt.Println(latticPaths(20))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
