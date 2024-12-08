module Input

open System.IO

let processInput (filePath: string) =
    let mutable left = []
    let mutable right = []

    File.ReadLines(filePath) 
    |> Seq.iter (fun line -> 
            let parts = line.Split([| ' '; '\t'; '\n'; '\r' |], System.StringSplitOptions.RemoveEmptyEntries)
            left  <- (int parts.[0]) :: left
            right <- (int parts.[1]) :: right
        )

    (List.rev left, List.rev right)


