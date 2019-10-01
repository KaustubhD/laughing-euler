//pb-20
package main

import(
  "fmt"
  "time"
  "math/big"
)

var fact = big.NewInt(1)

func sumFactorialDigits(n int) int{
  var sum int = 0

  fact.MulRange(1, int64(n))
  for _, i := range(fact.String()){
    sum += int(i - '0')
  }
  return sum
}

func main(){
  start := time.Now()

  fmt.Println(sumFactorialDigits(100))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
