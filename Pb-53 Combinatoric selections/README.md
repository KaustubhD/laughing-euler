# Problem 53
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, <sup>5</sup>C<sub>3</sub> = 10.

In general,

<sup>n</sup>C<sub>r</sub> =

n!r!(n−r)!

,where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.

It is not until n = 23, that a value exceeds one-million: <sup>23</sup>C<sub>10</sub> = 1144066.

How many, not necessarily distinct, values of  <sup>n</sup>C<sub>r</sub>, for 1 ≤ n ≤ 100, are greater than one-million?

Eg.
1. `combinatoricSelections(1000)`should return 4626.
2. `combinatoricSelections(10000)`should return 4431.
3. `combinatoricSelections(100000)`should return 4255.
4. `combinatoricSelections(1000000)`should return 4075.
