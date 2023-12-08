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

(2023/12/08): I also just read this bit in the "about" section of the AoC website:

> **Can I copy/redistribute part of Advent of Code?**
>
> Please don't. Advent of Code is free to use, not free to copy. If you're posting a 
> code repository somewhere, **please don't include parts of Advent of Code like the 
> puzzle text or your inputs.** If you're making a website, please don't make it look 
> like Advent of Code or name it something similar.

I want to keep the puzzle text and input (for me) so I can easily see what the problems where 
when looking back over this. Occasionally I refactor and want to run the whole thing again.
While I could make this whole repo private, I feel like that takes all the fun out of sharing
solutions.

I guess I need to keep the "puzzle text" and "input" separately, and privately. I have moved them
to a `not_public/` folder - and I now just have to remember to keep that backed up!

It's not perfect - the data are still in git history, and there are a few exceptions like when
I have done work in specific languages other than go.