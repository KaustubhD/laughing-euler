//pb-12
package main

import (
  "fmt"
  "time"
)

func numDivisors(n int) int{
  if n % 2 == 0{
    n /= 2
  } // We skipped one 2
  var numDivs, count, chk int = 1, 0, 2

  for n % chk == 0{
    n /= chk
    count += 1
  }
  numDivs *= count + 1

  chk = 3
  for n != 1{
    count = 0
    for n % chk == 0{
      n /= chk
      count += 1
    }
    chk += 2 // all primes are odd
    numDivs *= count + 1
  }
  return numDivs
}

func divisibleTriangleNumber(n int) int{
  var nats int = 1
  var nDivs, nPlusDivs int = numDivisors(nats), numDivisors(nats + 1)

  for nDivs * nPlusDivs < n{
    nats++
    nDivs, nPlusDivs = nPlusDivs, numDivisors(nats + 1)
  }

  return (nats * (nats + 1)) / 2
}

func main(){
  start := time.Now()

  fmt.Println(divisibleTriangleNumber(500))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
