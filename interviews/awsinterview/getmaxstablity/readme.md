Problem Statement

AWS provides a range of servers to meet the deployment needs of its clients. A client wants to choose a set of servers
to deploy their application. Each server is associated with an availability factor and a reliability factor.

The client defines the stability of a set of servers as:

stability = (minimum availability) * (sum of reliabilities)

Given two arrays of integers, availability and reliability, where:
• availability[i] represents the availability of the i-th server.
• reliability[i] represents the reliability of the i-th server.

Find the maximum possible stability of any subset of servers.

Since the answer can be large, report the answer modulo (10^9 + 7).

Example

Consider the set of servers where:
• reliability = [1, 2, 2]
• availability = [1, 1, 3]

Possible subsets and their stability:

Indices Stability Calculation Stability Value
[0]                    1 * 1 1
[1]                    1 * 2 2
[2]                    3 * 2 6
[0,1]                min(1,1) * (1+2)        3
[0,2]                min(1,3) * (1+2)        3
[1,2]                min(1,3) * (2+2)        4
[0,1,2]                min(1,1,3) * (1+2+2)    5

The maximum stability is 6, from subset {2}.
Thus, the answer is:

6 mod (10^9 + 7) = 6

Function Description

Complete the function getMaxStability in the editor below.

func getMaxStability(reliability []int, availability []int) int

Parameters
• int reliability[n] → Server reliability values
• int availability[n] → Server availability values

Returns
• int → The maximum stability among all possible subsets, modulo 10^9 + 7.

Constraints
• 1 <= n <= 10^5
• 1 <= reliability[i], availability[i] <= 10^6
• It is guaranteed that the lengths of reliability and availability are the same.
