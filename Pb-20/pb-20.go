//pb-20
package main

import(
  "fmt"
  "time"
  "math/big"
)

var fact = big.NewInt(1)

func sumFactorialDigits(n int) int{

  fact.Set(n)

  for i := 2; i < n; i++{
    fact.Mul(fact, i)
  }
  var sum int = 0
  for _, i := range(str){
    sum += int(i - '0')
  }
  return sum
}

func main(){
  start := time.Now()

  fmt.Println(sumFactorialDigits(25))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
