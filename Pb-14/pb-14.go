//pb-14
package main

import(
  "fmt"
  "time"
)

func longestCollatzSequence(n int) int{
  var coll_lens = make([]int, n) // Array of collatz lengths of all numbers till 'n'
  max := [2]int{0, 0}

  coll_lens[1] = 1

  for i := 0; i < n; i++{ // checking every number till 'n'
    var ith, count int = i, 0
    intermediates := make([]int, 0)
    for ith > n - 1 || count < 3 || coll_lens[ith] < 1{
      intermediates = append(intermediates, ith)
      fmt.Println(intermediates)
      if ith % 2 == 0{
        ith /= 2
      }else{
        ith = 3 * ith  + 1
      }
      count += 1
    }

    for ind, intermed := range intermediates{
      intermed_coll_length := coll_lens[count] + len(intermediates) - ind
      coll_lens[intermed] = intermed_coll_length
      if intermed_coll_length > max[1]{
        max[0], max[1] = intermed, intermed_coll_length
      }
    }

  }
  return max[0]
}

func main(){
  start := time.Now()

  fmt.Println(longestCollatzSequence(14))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}
