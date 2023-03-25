
// Kristel Salas Villegas

(* Migrar el ejercicio realizado en Go sobre facturas con listas de productos haciendo énfasis en la
implementación de las funciones para calcular el monto a pagar por la factura completa y el monto a pagar
 por concepto del 13% de impuesto de venta para aquellos productos que deban pagarlo según criterio de selección.*)

type producto = {
    nombre: string
    descripcion: string
    montoVenta: int
}

type Productos = producto list

let factura : Productos = []

let rangoPagoImpuestos = 20000
let porcentajeImpuesto = 0.13

let agregarProducto (f: Productos) (nom: string) (desc: string) (pre: int) =
    let nuevoProducto = {nombre=nom; descripcion=desc; montoVenta=pre}
    match List.tryFind (fun p -> p.nombre = nom) f with
    | None -> nuevoProducto::f // El producto no existe, así que lo agregamos al inicio de la lista
    | Some _ -> printfn "No se puede agregar un producto existente"; f

let calcularImpuestoFactura (f: Productos) =
    let lista = f |> List.filter (fun p -> p.montoVenta > rangoPagoImpuestos)
    let lista2 = lista |> List.map (fun p -> int(float p.montoVenta * porcentajeImpuesto))
    lista2 |> List.reduce (+)

let calcularMontoFactura (f: Productos) =
    let lista = f |> List.map (fun p -> p.montoVenta)
    lista |> List.reduce (+)


let facturaConProducto1 = agregarProducto factura "tarjeta madre" "Asus" 54200
let facturaConProducto2 = agregarProducto facturaConProducto1 "mouse" "alámbrico" 15000
let facturaConProducto3 = agregarProducto facturaConProducto2 "teclado" "gamer con luces" 30000
let facturaConProducto4 = agregarProducto facturaConProducto3 "memoria ssd" "2 gb" 65000
let facturaConProducto5 = agregarProducto facturaConProducto4 "cable video" "display port 1m" 18000

printfn "El monto total de la factura es: %d" (calcularMontoFactura facturaConProducto5)
printfn "El impuesto total a pagar es: %d" (calcularImpuestoFactura facturaConProducto5)
