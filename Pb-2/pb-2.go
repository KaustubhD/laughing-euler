package main
import "fmt"

func fiboEvenSum(num int) int{
  a, b, sum := 0, 2, 0
  for b < num{
    sum += b
    a, b = b, a + 4 * b
  }
  return sum
}

func main(){
  fmt.Println(fiboEvenSum(43))
}
