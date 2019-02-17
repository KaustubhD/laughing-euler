//pb-16
package main

import(
  "fmt"
  "strconv"
  "strings"
  "time"
)



func powerDigitSum(n int) int64{

  var st strings.Builder
  st.WriteString("1")

  for i := 0; i < n; i++{
    st.WriteString("0")
  }
  intt, _ := strconv.ParseInt(st.String(), 2, 64)
  fmt.Println(st.String())
  var str string = strconv.FormatInt(intt, 10)
  var sum int64 = 0
  runeFunc := func(r rune) rune{
    //fmt.Println(r)
    sum += int64(r - '0')
    return r
  }
  fmt.Println(str)
  str = strings.Map(runeFunc, str)
  //fmt.Println(sum)
  return sum
}

func main(){
  start:= time.Now()

  fmt.Println(powerDigitSum(200))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
