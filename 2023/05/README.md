# Advent of Code 2023 day 5

See [https://adventofcode.com/2023/day/5](https://adventofcode.com/2023/day/5)

## Notes

This was a classic AoC problem.

The naive solution for part 1 that passes the test is to use a `map[int]int` and literally map every possibility for each map range, then follow them from seed->location.

But of course in the real input, the ranges are HUGE, so this approach falls over... And I switched to representing the map as a series of ranges. This way we only store the data needed to map an input to an output, but not a list all non-sparse inputs. 

Then we move to part 2. The same process as part one, but this time expanding the input from a few seeds to the range worked fine in the test (small ranges), but the real input immediately proved way too slow (target is < 15 seconds on 10 year old hardware, so we are realistically looking a ~1 second). So we re-think _again_.

There is a good trick to the fast solution, in that we only consider "ranges", and instead of mapping values to values we map an input range to a list of output ranges. If we do this for the starting ranges and then apply the new list of ranges to the next map and so on, we keep the number of lookups and object manageable.
At the end, we just look for the location that is the minimum value of any of the final ranges.