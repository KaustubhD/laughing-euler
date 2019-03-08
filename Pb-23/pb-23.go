import (
	"fmt"
	"math"
)

func getSum(p, a float64) int {
	fmt.Println("Factor", p, "is", a, "times")
	dobP := (math.Pow(p, a+1) - 1) / (p - 1)
	return int(dobP)
}

func primeSummation(n int) int {
	var i, count, sum int = 3, 0, 0
	for n%2 == 0 {
		n /= 2
		count++
	}
	sum += getSum(float64(2), float64(count))
	for n != 1 {
		count = 0
		for n%i == 0 {
			n /= i
			count++
		}
		if count > 0 {
			sum += getSum(float64(i), float64(count))
		}
		i += 2
	}
	// var MAX = n
	// var sum int = 0
	// var arr  = make([]bool, n)
	// var count int = 0
	// arr[0], arr[1] = true, true
	// for i := 0; i < MAX; i++{
	//   //fmt.Println(i)
	//   if !arr[i]{
	//     sum += i
	//     for j, k := i * 2, 0; k < ((MAX - 1)/i) - 1; k++{
	//       arr[j] = true
	//       j += i
	//     }
	//     count += 1
	//   }
	//    if count == n{
	//      return sum
	//    }
	// }
	return sum
}
