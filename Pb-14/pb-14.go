//pb-14
package main

import(
  "fmt"
  "time"
)

func longestCollatzSequence(n int) int{
  var coll_lens = make([]int, n + 1) // Array of collatz lengths of all numbers till 'n'
  max := [2]int{0, 0}

  coll_lens[1] = 1

  for i := 1; i <= n; i++{ // checking every number till 'n'
    var ith, count int = i, 0
    intermediates := make([]int, 0)
    for ith > n - 1 || coll_lens[ith] >= 1{
      if ith < n{
        intermediates = append(intermediates, ith)
      }
      fmt.Println(ith, n - 1)
      // fmt.Println(intermediates)
      if ith % 2 == 0{
        ith /= 2
      }else{
        ith = 3 * ith  + 1
      }
      count += 1
    }

    for _, intermed := range intermediates{
      fmt.Println("intermed : ", intermed, " and its len", coll_lens[intermed])
      intermed_coll_length := coll_lens[count] + len(intermediates)
      coll_lens[intermed] = intermed_coll_length
      if intermed_coll_length > max[1]{
        max[0], max[1] = intermed, intermed_coll_length
      }
    }
    fmt.Println("Length ", i,":", coll_lens[i])
  }
  return max[0]
}

func main(){
  start := time.Now()

  fmt.Println(longestCollatzSequence(14))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}

// 7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
