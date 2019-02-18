//pb-16
package main

import(
  "fmt"
  "strings"
  "time"
  "math/big"
)

func powerDigitSum(n int) int64{
  var st strings.Builder
  st.WriteString("1")
  for i := 0; i < n; i++{
    st.WriteString("0")
  }
  intt := new(big.Int)
  intt.SetString(st.String(), 2)

  var sum int64 = 0
  runeFunc := func(r rune) rune{
    sum += int64(r - '0')
    return r
  }

  strings.Map(runeFunc, intt.String())
  return sum
}

func main(){
  start:= time.Now()

  fmt.Println(powerDigitSum(1000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
