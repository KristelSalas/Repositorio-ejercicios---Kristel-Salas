
// Kristel Salas Villegas

(*Construya una función que se llame general-sec A B. Esta función genera una secuencia de números en
una lista de A hasta B, donde  A y B son números enteros.
Por ejemplo:
        generar-sec 0 5
        [0; 1; 2; 3; 4; 5]
*)

let rec general_sec a b =
    if a > b then []
    else a :: general_sec (a + 1) b
let secuencia = general_sec 2 15

printfn "%A" secuencia