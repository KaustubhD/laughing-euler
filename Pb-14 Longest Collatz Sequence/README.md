# Problem 14

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)

n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under the given **limit**, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million

Eg.

1. ```longestCollatzSequence(14)``` should return 9
2. ```longestCollatzSequence(5847)``` should return 3711
3. ```longestCollatzSequence(46500)``` should return 35655
4. ```longestCollatzSequence(54512)``` should return 52527
5. ```longestCollatzSequence(1000000)``` should return 837799
