# Advent of Code: https://adventofcode.com/

My solutions to the advent of code problems.

I started in Nov 2017 with the 2016 problem set. I didn't finish before the
2017 set started and I only got so far through that. But these problems can
be solved anytime and I was never intending to top the scoreboard.

As I mostly use Go for these, there is a `main.go` script here in the repo.

Running `go run main.go -year 2018 -day 1` will generate a skeleton problem
for the year and day. `-year` is optional and defaults to this year. `-day` is
optional only in December.

If the problem already exists then it simply executes the `main.go` file in the
problem dir.

The template is simple because all the problems essentially boil down to: convert
this long string into a (certain) short string.
