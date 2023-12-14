# Advent of Code 2023 day 12

See [https://adventofcode.com/2023/day/12](https://adventofcode.com/2023/day/12)

## notes

This was super difficult. The brute force method was way too slow for the longer
strings, the possibility space exploded.

Took me ages to come up with the dynamic programming solution, where we can cache states
easily and iterate over the string the minimum amount.
