# Day 13

## Analysis

The puzzle gives us two buttons and for each button press it moves the claw X spaces on the x-axis and Y spaces on the y-axis.
The goal is to position the claw over the prize, which the puzzle gives its exact xy coordinate with the least amount of tokens used.
It takes 3 tokens to press Button A and 1 token to press Button B. Additionally, the limit to press a button is 100.

## Approach

Parsing and setting up the data was simple and straight forward. I used a regex to parse out each claw machine's
configuration, store it in the struct `ClawMachine`, and return an array of `ClawMachine`

```go
type ClawMachine struct {
	A struct {
		X int
		Y int
	}

	B struct {
		X int
		Y int
	}

	Prize struct {
		X int
		Y int
	}
}
```

### Part 1

To solve part 1,
I could've looped from 0 to 100 for each button until the lowest number of N and K, the amount of times to press button A and button B respectively, were found. I.E

```go
for n := 0; n <= 100; n++ {
  for k := 0; k <= 100; k++ {
    // check if the current n and k satisfy the equation
  }
}
```

However, this is not optimal and rather brute force solution.
Carefully inspecting the puzzle and writing out the problem's formula, I can see that it is a system of equations.

$$
  N(A.x, A.y) + K(B.x, B.y) = (P.x, p.y)
$$

where (N) and (K) are the number of presses for buttons A and B, respectively.

The function `findCombinations()` extracts the coordinates for A, B, and Prize. Calculate the determinant and check
whether the claw machine has a combination to win the prize. This efficiently determines whether it is worth looking
for a solution. Then, solve for N and K. If N or K are not integers, we return [0,0] since valid button presses are whole numbers.

$$
{\text{determinant}} = A.x \cdot B.y - A.y \cdot B.x
$$

$$
N = \frac{(P.x \cdot B.y) - (P.y \cdot B.x)}{\text{determinant}}
$$

$$
K = \frac{P.x - (N \cdot A.x)}{B.x}
$$

## Part 2

Solving part 2 is fairly simple since 10000000000000 is added to the prize's X and Y before using `findCombinations()`. Thus, changing the formula to:

$$
  N(A.x, A.y) + K(B.x, B.y) = (P.x + 10000000000000, p.y + 10000000000000)
$$

## Solution Debrief

Day 13 was fairly trick in terms of wording. However, I am glad I ignored the 100 button rule and was lucky enough with my implementation that it worked for both part 1 and part 2. Only difference was to add the 10000000000000 to the Prize's X and Y and run the same process.

Stats

| **Part** | **Average Time** | **Average Bytes Allocated** | **Average Memory Allocations** |
| -------- | ---------------- | --------------------------- | ------------------------------ |
| 1        | 29413 ns/op      | 0 B/op                      | 0 allocs/op                    |
| 2        | 104562 ns/op     | 0 B/op                      | 0 allocs/op                    |

## Self-Reflect

Glad to remember how to do linear algebra...
