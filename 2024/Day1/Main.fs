module Main

open Input
open Part1
open Part2

let input = "input.txt"
let (first, second) = processInput input

let differences = pairsDiff first second
let part1 = List.sum differences

printfn "Part 1: %d" part1

let similarities = findSimilarities first second
let part2 = List.sum similarities

printfn "Part 2: %d" part2


