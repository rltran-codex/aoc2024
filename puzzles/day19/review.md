# Day 19

## Analysis

We are given a stack of towels that we can use to combine together in hopes to generate a match
of a design. There are no limitations on the number of towels we use, but we cannot rearrange or reverse the proposed design.
From the sample:

```
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
```

-- brwrr can be made with a br towel, then a wr towel, and then finally an r towel.
-- bggr can be made with a b towel, two g towels, and then an r towel.
-- gbbr can be made with a gb towel and then a br towel.
-- rrbgbr can be made with r, rb, g, and br.
-- ubwu is impossible.
-- bwurrg can be made with bwu, r, r, and g.
-- brgr can be made with br, g, and r.
-- bbrgwb is impossible.

## Approach and Solution Debrief

Initially, I approached this problem as a divide-and-conquer puzzle for combinatorial pattern-matching. My idea was to split the string into parts wherever a match was found using `regex.FindAllIndex`, without focusing on whether to process from left to right. I would then recursively handle the left and right parts of the string separately.

This approach worked for my initial test cases, but while debugging an actual design from the puzzle input, I noticed some critical flaws. Specifically, when the left side of the string couldn’t find a valid match, my algorithm got stuck in an infinite loop, repeatedly trying to solve the leftmost part.

To address this issue, I revised my strategy to always process the string from left to right. By doing so, I ensured that each split would result in two separate substrings—one that I could confidently process (the left) and the other to be tackled afterward (the right). This method allowed me to maintain control and avoid infinite loops.

The key condition, `lValid && rValid`, determined whether the left and right parts both matched successfully. If true, the solution would proceed; otherwise, it would backtrack and try the next available towel. This adjustment made the algorithm more robust and reliable for handling all input designs.

In part 2, I had to get redo my approach. Instead of keeping track of the combo that is built, even though it was great to see the exact make up of the combo, this was causing the problem to become confusing for me.
I needed a way to memoize the problem, so I went ahead and brushed up on the topic and used
which would store the slice of the design with the combination count. In turn, this was reduced any unneccessary reiterations.

```go
var cache map[string]int = make(map[string]int)
```

## Self-Reflect

Need to get better at keeping the problem more simple and brush up on Dynamic Programming.
