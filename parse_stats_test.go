package main

import (
	"encoding/json"
	"os"
	"testing"
)

const data = `
<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8"/>
<title>Leaderboard - Advent of Code 2019</title>
<!--[if lt IE 9]><script src="/static/html5.js"></script><![endif]-->
<link href='//fonts.googleapis.com/css?family=Source+Code+Pro:300&subset=latin,latin-ext' rel='stylesheet' type='text/css'>
<link rel="stylesheet" type="text/css" href="/static/style.css?24"/>
<link rel="stylesheet alternate" type="text/css" href="/static/highcontrast.css?0" title="High Contrast"/>
<link rel="shortcut icon" href="/favicon.png"/>
</head><!--




Oh, hello!  Funny seeing you here.

I appreciate your enthusiasm, but you aren't going to find much down here.
There certainly aren't clues to any of the puzzles.  The best surprises don't
even appear in the source until you unlock them for real.

Please be careful with automated requests; I'm not a massive company, and I can
only take so much traffic.  Please be considerate so that everyone gets to play.

If you're curious about how Advent of Code works, it's running on some custom
Perl code. Other than a few integrations (auth, analytics, ads, social media),
I built the whole thing myself, including the design, animations, prose, and
all of the puzzles.

The puzzles are most of the work; preparing a new calendar and a new set of
puzzles each year takes all of my free time for 4-5 months. A lot of effort
went into building this thing - I hope you're enjoying playing it as much as I
enjoyed making it for you!

If you'd like to hang out, I'm @ericwastl on Twitter.

- Eric Wastl


















































-->
<body>
<header><div><h1 class="title-global"><a href="/">Advent of Code</a></h1><nav><ul><li><a href="/2019/about">[About]</a></li><li><a href="/2019/events">[Events]</a></li><li><a href="https://teespring.com/adventofcode-2019" target="_blank">[Shop]</a></li><li><a href="/2019/settings">[Settings]</a></li><li><a href="/2019/auth/logout">[Log Out]</a></li></ul></nav><div class="user">thechriswalker <span class="star-count">22*</span></div></div><div><h1 class="title-event">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="title-event-wrap">/^</span><a href="/2019">2019</a><span class="title-event-wrap">$/</span></h1><nav><ul><li><a href="/2019">[Calendar]</a></li><li><a href="/2019/support">[AoC++]</a></li><li><a href="/2019/sponsors">[Sponsors]</a></li><li><a href="/2019/leaderboard">[Leaderboard]</a></li><li><a href="/2019/stats">[Stats]</a></li></ul></nav></div></header>

<div id="sidebar">
<div id="sponsor"><div class="quiet">Our <a href="/2019/sponsors">sponsors</a> help make Advent of Code possible:</div><div class="sponsor"><a href="https://auricsystems.com/tokenize-almost-anything?utm_source=advent+of+code&amp;utm_medium=display&amp;utm_camapign=advent2019" target="_blank" onclick="if(ga)ga('send','event','sponsor','sidebar',this.href);" rel="noopener">AuricVault®</a> - Thieves can&apos;t steal -- what isn&apos;t there! Secure your sensitive data. Simple API. FREE test credentials. Start simplifying your compliance.</div></div>
</div><!--/sidebar-->

<main>
<article><p>These are your personal leaderboard statistics.  <em>Rank</em> is your position on that leaderboard: 1 means you were the first person to get that star, 2 means the second, 100 means the 100th, etc.  <em>Score</em> is the number of points you got for that rank: 100 for 1st, 99 for 2nd, ..., 1 for 100th, and 0 otherwise.</p><pre>      <span class="leaderboard-daydesc-first">--------Part 1--------</span>   <span class="leaderboard-daydesc-both">--------Part 2--------</span>
Day   <span class="leaderboard-daydesc-first">    Time   Rank  Score</span>   <span class="leaderboard-daydesc-both">    Time   Rank  Score</span>
 11   05:01:34   3994      0          -      -      -
 10   16:03:26   8886      0   18:16:28   6888      0
  9   12:19:42   7211      0   12:21:32   7124      0
  8       &gt;24h  15823      0       &gt;24h  15071      0
  7       &gt;24h  15788      0       &gt;24h  12076      0
  6   11:40:33  11897      0   11:55:24  10581      0
  5   04:31:26   5637      0   05:01:31   4872      0
  4   07:17:03  14203      0   07:25:09  12170      0
  3       &gt;24h  24018      0       &gt;24h  20528      0
  2       &gt;24h  42252      0       &gt;24h  38068      0
  1       &gt;24h  52685      0       &gt;24h  46406      0
</pre>
</article>
</main>

<!-- ga -->
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-69522494-1', 'auto');
ga('set', 'anonymizeIp', true);
ga('send', 'pageview');
</script>
<!-- /ga -->
</body>
</html>`

func TestParseHTML(t *testing.T) {
	p := YearProgress{
		Year: 2019,
	}
	err := parseStatsHTML(data, &p)
	if err != nil {
		t.Fatalf("Parse Fail %v", err)
	}
	json.NewEncoder(os.Stdout).Encode(p)
}
