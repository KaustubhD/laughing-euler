//pb-22
package main

import(
  "fmt"
  "time"
  "io/ioutil"
  "strings"
  "sort"
)
var test1 = []string{"THIS","IS","ONLY","A","TEST"}

var test2 = []string{"I","REPEAT","THIS","IS","ONLY","A","TEST"}

func calcScore(s string) int{
  var sum int = 0
  for index := range s{
    sum += int(s[index] - 64)
  }
  return sum
}

func nameScores(og []string) int{
  sort.Strings(og)
  var total int = 0
  for ind, str := range(og){
    total += (ind + 1) * calcScore(str)
  }
  return total
}

func main(){
  bytes, err := ioutil.ReadFile("names.txt")

  start := time.Now()
  if err != nil{
    fmt.Println(err)
  }else{
    strs := strings.Split(string(bytes), ",")
    fmt.Println(nameScores(strs))
  }

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
