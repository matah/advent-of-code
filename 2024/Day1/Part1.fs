module Part1

let pairsDiff (left: List<int>) (right: List<int>) =
    List.map2 (fun x y -> abs(x - y)) (left |> List.sortDescending) (right |> List.sortDescending)


