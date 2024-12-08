module Part2

let similarity x c = x * c

let occurrences x list = List.filter ((=) x) list |> List.length

let findSimilarities (left: List<int>) (right: List<int>) =
    let rec findSimilarities' (left: List<int>) (right: List<int>) (acc: List<int>) =
        match left with
        | [] -> acc
        | x::xs -> 
            let c = occurrences x right
            findSimilarities' xs right (acc @ [similarity x c])
            
    findSimilarities' left right []


