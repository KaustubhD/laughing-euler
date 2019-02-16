//pb-15
package main

import(
  "fmt"
  "time"
)

func latticPaths(n int) int{
//  var x int64 = 0
  arr := make([][]int, n + 1)
  for i := 0; i <= n; i++{
    arr[i] = make([]int, n + 1)
  }
  fmt.Println(arr)
  

  for i := 0; i <= n; i++{
    for j := 0; j <= i; j++{
      fmt.Println(i, j)
      if i == 0 || j == 0{
        arr[i][i] = 1
      }else{
        arr[i][j] = arr[i + 1][j] + arr[i][j + 1]
      }
    }
  }
  return arr[n][n]
}


func main(){
  start := time.Now()

  fmt.Println(latticPaths(2))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
