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
  var ith, count int = 0, 0

  for i := 1; i <= n; i++{ // checking every number till 'n'
    ith, count = i, 0
    intermediates := make([]int, 0)
    for ith > n - 1 || coll_lens[ith] < 1{ // making a collatz sequence till a known element is found
      intermediates = append(intermediates, ith) // if not known, add it to the intermediates array
      if ith % 2 == 0{
        ith /= 2
      }else{
        ith = 3 * ith  + 1
      }
      count += 1 // length of collatz sequence till a known is found
    }

    for ind, intermed := range intermediates{
      if intermed < n{ // only want elements in range till 'n'
        intermed_coll_length := coll_lens[ith] + count - ind
        coll_lens[intermed] = intermed_coll_length
        if intermed_coll_length > max[1]{
          max[0], max[1] = intermed, intermed_coll_length
        }
      }
    }
  }
  return max[0]
}

func main(){
  start := time.Now()

  fmt.Println(longestCollatzSequence(1000000))

  elapsed := time.Since(start)
  fmt.Printf("Time taken: %s\n", elapsed)
}

// 7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
// 9 -> 28 -> 14 -> 7 ->
