package main

import(
  "fmt"
  "time"
)

func smallestMult(n int) int{
  lcm := 2
  for i := 3; i <= n; i++{
    lcm = (lcm * i) / gcd(i, lcm)
  }
  return lcm
}

func gcd(x, y int) int{
  for x!=y{
    if x > y{
      x -= y
    }else{
      y -= x
    }
  }
  return x
}

func main(){
  start := time.Now()

  fmt.Println(smallestMult(20))

  elapsed := time.Since(start)
  fmt.Printf("Time taken %s\n", elapsed)
}
