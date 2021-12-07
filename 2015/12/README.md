# Advent of Code 2015 day 12: JSAbacusFramework.io ---

## Comments

This was fun. Golang's JSON package has a decoder which you can use in a streaming fashion, token by token. This made the first part easy, just ignore all tokens but numbers and add them up. However, the second was a bit more difficult, because I had to now stream arrays and object separately in order to check the objects for a key with the value "red". In JS this would have been quite simple with a "reviver" function in `JSON.parse` (which is highly optimised), but took a bit of fudging in go until I got each of the sections correctly handling each sub-object/sub-array.

## Problem 1

Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their accounting software uses a peculiar storage format. That's where you come in.

They have a JSON document which contains a variety of things: arrays (`[1,2,3]`), objects (`{"a":1, "b":2}`), numbers, and strings. Your first job is to simply find all of the **numbers** throughout the document and add them together.

For example:

 - `[1,2,3]` and `{"a":2,"b":4}` both have a sum of `6`.
 - `[[[3]]]` and `{"a":{"b":4},"c":-1}` both have a sum of `3`.
 - `{"a":[-1,1]}` and `[-1,{"a":1}]` both have a sum of `0`.
 - `[]` and `{}` both have a sum of 0.

You will not encounter any strings containing numbers.

What is the **sum of all numbers in the document?**

## Problem 2

Uh oh - the Accounting-Elves have realized that they double-counted everything **red**.

Ignore any object (and all of its children) which has any property with the value `"red"`. Do this only for objects (`{...}`), not arrays (`[...]`).

 - `[1,2,3]` still has a sum of `6`.
 - `[1,{"c":"red","b":2},3]` now has a sum of `4`, because the middle object is ignored.
 - `{"d":"red","e":[1,2,3,4],"f":5}` now has a sum of `0`, because the entire structure is ignored.
 - `[1,"red",5]` has a sum of `6`, because "red" in an array has no effect.


...
