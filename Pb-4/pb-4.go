package main
import(
  "fmt"
  "math"
  "strconv"
)

func getPalindrome(n float64) int64{
  MAX := math.Pow(10, n) - 1
  // fmt.Println(MAX)
  big, bdiff := 0.0, 0.0
  prod, diff := 0.0, 0.0
  if int(n) % 2 == 0{
    return int64((math.Pow(10, n / 2) - 1) * math.Pow(10, (3 * n) / 2) + math.Pow(10, n / 2) - 1)
  }else{
    s := ""
    for i := MAX; i >= 1; i--{
      // skip := 1.0
      // if int(i) % 11 != 0{
      //   skip = 11
      // }
      for j:= i; j >= 1; j --{
        prod = i * j
        diff = math.Abs(i - j)
        if prod > big && diff > bdiff{
          // fmt.Println("Doing >",i,j,"<")
          s = strconv.Itoa(int(prod))
          if s[0] == '9' && s[len(s) - 1] == '9' && isPal(int64(prod)){
            big = prod
          }
        }
      }
    }
  }
  return int64(big)
}

func isPal(n int64) bool{
  var rev , og int64 = 0, n
  for n > 0 {
    remainder := n % 10
    rev *= 10
    rev += remainder
    n /= 10
  }
  return og == rev
}

func main(){
  fmt.Println(getPalindrome(5))
}
