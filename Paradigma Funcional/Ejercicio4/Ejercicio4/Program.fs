open System
// Kristel Salas Villegas

// Ejercicios de rutas en Grafos con búsqueda en profundidad

//        a ---- c ---- x
//      /   \  /
//     /     \/
//   i       X
//     \    / \
//      \  /   \
//        b ---- d ---- f

let grafo = [
    [('i', 0);('a', 1);('b', 2)];
    [('a', 1);('i', 0);('c', 3);('d', 2)];
    [('b', 2);('i', 0);('c', 1);('d', 4)];
    [('c', 3);('a', 1);('b', 1);('x', 2)];
    [('d', 2);('a', 4);('b', 1);('f', 3)];
    [('f', 3);('d', 2)];
    [('x', 2);('c', 3)]
]

let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, peso)::adyacencias)::tail when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail  

let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> [])

let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        //List.rev rutas.Head
        List.append
            ([List.rev rutas.Head])
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
        
let prof ini fin grafo =
    prof_aux fin [[(ini,0)]] grafo

printf "Todas las rutas de i a f\n"
printf "%A" (prof 'i' 'f' grafo)

//EJERCICIO DEL REPO 
(* Modifique el ejercicio de búsqueda en profundidad visto en clase para cálculo de rutas en un grafo,
para ingresar pesos a los vértices y con esta nueva información, calcular el camino más corto en la búsqueda
en profundidad basándose en la sumatoria de pesos de los vértices de la ruta*)

let rec rutaMenor fin (rutas: (char * int) list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        let rutaActual = List.rev rutas.Head
        let pesoTotal = rutaActual |> List.map snd |> List.sum
        let rutasRestantes = rutaMenor fin rutas.Tail grafo
        let rutasConPesoTotal = (rutaActual, pesoTotal)::rutasRestantes
        rutasConPesoTotal
    else
        rutaMenor fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo

let prof2 ini fin grafo =
    rutaMenor fin [[(ini,0)]] grafo
    |> List.sortBy (fun (_, peso) -> peso)
    |> List.head
    |> fst

printf "\nLa ruta de menor peso de i a f:\n"
printf "%A" (prof2 'i' 'f' grafo)

