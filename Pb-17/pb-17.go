//pb-17
package main

import(
  "fmt"
  "time"
  "strconv"
)
var try = []map[string]int{
  {"1": 3, "2": 3, "3": 5, "4": 4, "5": 4, "6": 3, "7": 5, "8": 5, "9": 4},
  {"2": 6, "3": 6, "4": 6, "5": 5, "6": 5, "7": 7, "8": 6, "9": 6, "10": 3, "11": 6, "12": 6, "13": 8, "14": 8, "15": 7, "16": 7, "17": 9, "18": 8, "19": 8},
  {}
}

var PRIMITIVES = map[string]int{
  "1": 3, "2": 3, "3": 5, "4": 4, "5": 4, "6": 3, "7": 5, "8": 5, "9": 4,
}

var TEENS = map[string]int{
  "10": 3, "11": 6, "12": 6, "13": 8, "14": 8, "15": 7, "16": 7, "17": 9, "18": 8, "19": 8,
}

var TENS = map[string]int{
  "2": 6, "3": 6, "4": 6, "5": 5, "6": 5, "7": 7, "8": 6, "9": 6,
}

func numberLetterCounts(n int) int{
  var str string = strconv.Itoa(n)
  var totalLen int = len(str)
  var sum int = 0
  for i := 0; i < totalLen; i++{
    switch i {
      case 1:
        if str[i] == '1'{
          sum += try[i][str[i:2]
    }
  }


  return 44
}

func main(){
  start := time.Now()

  fmt.Println(numberLetterCounts(5))
  //fmt.Println(try[0]["1"])

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}

