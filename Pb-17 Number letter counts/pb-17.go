//pb-17
package main

import(
  "fmt"
  "time"
  "strconv"
  "strings"
)
var try = []map[string]int{
  {"1": 3, "2": 3, "3": 5, "4": 4, "5": 4, "6": 3, "7": 5, "8": 5, "9": 4},
  {"2": 6, "3": 6, "4": 5, "5": 5, "6": 5, "7": 7, "8": 6, "9": 6, "10": 3, "11": 6, "12": 6, "13": 8, "14": 8, "15": 7, "16": 7, "17": 9, "18": 8, "19": 8},
  {"_": 7},
  {"_": 8},
}

func numberLetterCounts(n int) int{
  sum := 0
  for i := 1; i <= n; i++{
    sum += countThis(i)
  }
  return sum
}

func countThis(n int) int{
  var str string = strconv.Itoa(n)
  var totalLen int = len(str)
  var sum int = 0
  for i := totalLen - 1; i >= 0; i--{
    var revIdx = totalLen - i - 1
    if revIdx == 0 && totalLen >= 2 && str[i - 1] == '1'{
      sum += try[revIdx + 1][str[i - 1 : i + 1]]
      i = totalLen - 2
    }else if revIdx >= 2 && str[i: i + 1] != "0"{
      sum += try[revIdx]["_"] + try[0][str[i: i + 1]]
    }else{
      sum += try[revIdx][str[i : i + 1]]
    }
  }

  if totalLen > 2 && strings.Count(str, "0") != totalLen - 1{
    sum += 3
  }
  return sum
}

func main(){
  start := time.Now()

  fmt.Println(numberLetterCounts(1000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}

