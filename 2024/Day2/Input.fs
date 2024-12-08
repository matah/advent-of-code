module Input

let processLine (input: string) =
    input.Split([|' '|])
    |> Array.map int

let readFile (path: string) =
    System.IO.File.ReadAllLines(path) |> Array.map processLine |> Array.toList

