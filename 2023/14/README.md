# Advent of Code 2023 day 14

See [https://adventofcode.com/2023/day/14](https://adventofcode.com/2023/day/14)

## notes

This was fun. Part 1 is simple. Part 2 was more tricky as I immediately noticed that we didn't need to consider the
orientation of the rocks, but instead simply the pattern of "load"s.

So I turned it into a repeating-pattern-spotting exercise instead of a "have we seen this state before" exercise.

The repeating pattern spotting code was interesting and fun to write, but unnecessary and slower than the "cache" 
all states and find the first repetition.

But the essence is the same either way. Find the cycle and use modulo arithmetic to find the index of the 1 billionth
iteration without actually doing them all.

- recording weights and finding repeating patterns took about 85ms
- caching states and finding first re-visited state took about 60ms

not much in it.