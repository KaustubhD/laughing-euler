//pb-19
package main

import(
  "fmt"
  "time"
)

var monthOff = [12]int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4};

func GetDay(day, month, year int) int{
  if month < 3{
    year -= 1
  }
  return ((year) + (year/4) - (year/100) + (year/400) + monthOff[month - 1] + day) % 7;
}

func countingSundays(st, en int) int{
  var count int = 0
  for y := st; y <= en; y++{
    for m := 1; m <= 12; m++{
      if(GetDay(1, m, y) == 0){
        count++
      }
    }
  }
  return count
}

func main(){
  start := time.Now()

  fmt.Println(countingSundays(1901, 2000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
