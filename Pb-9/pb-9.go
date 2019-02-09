//pb-9
package main

import(
  "fmt"
  "math"
  "time"
)

func gcd(a, b int) int{
  for a != b{
    if a > b{
      a -= b
    } else{
      b -= a
    }
  }
  return a
}

func specialPythagoreanTriplet(n int) int{
  var k int = 0
  maxM := int(math.Sqrt(float64(n / 2)))
  for m := 2; m <= maxM; m++{
    if (n / 2) % m == 0{
      if m % 2 == 0{
        k = m + 1
      } else{
        k = m + 2
      }
      for k < 2 * m && k <= n / (2 * m){
        if n / (2 * m) % k == 0 && gcd(k, m) == 1{
          var d int = n / (2 * k * m)
          var temp int = k - m
          var a int = d * (m * m - temp * temp)
          var b int = d * (2 * m * temp)
          var c int = d * (m * m + temp * temp)
          return a * b * c
        }
        k += 2
      }
    }
  }
  return 0
}


func main(){
  start := time.Now()

  fmt.Println(specialPythagoreanTriplet(1000))

  elapsed := time.Since(start)
  fmt.Println("Time taken: ", elapsed)
}
