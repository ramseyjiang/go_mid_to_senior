Two unordered slices of any length are given as input. It is necessary to write a function that returns their
intersection.

The usual way: it will use longer time, by sorting without allocating additional memory.

The shorter way: allocate additional memory and solve in linear way.

Example 1:
a := []int{23, 3, 1, 2} b := []int{6, 2, 4, 23}

result1: [2, 23]

Example 2:
a := []int{1, 1, 1} b := []int{1, 1, 1, 1}

result2: [1, 1, 1]