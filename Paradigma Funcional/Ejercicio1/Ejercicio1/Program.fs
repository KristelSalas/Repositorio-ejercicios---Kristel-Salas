
//Kristel Salas Villegas

(* Haciendo uso de la función filter, implemente una función que a partir de una lista de cadenas representando 
frases de parámetro, filtre aquellas que contengan una palabra que el usuario indique en otro argumento 
(palabra completa contenida). 
Ej sub_cadenas “la” [“la casa; “el perro”; “pintando la cerca”]
    [“la casa, “pintando la cerca”]
*)

let sub_cadenas (palabra: string) (frases: string list) =
    frases |> List.filter (fun frase -> frase.Contains(palabra))

let listaFrases = ["el proyecto"; "la tarea"; "el quiz"; "la exposición"]

printfn "%A" (sub_cadenas "el" listaFrases)
