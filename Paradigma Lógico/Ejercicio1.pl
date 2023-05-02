
/*
Primer ejercicio
Implemente un predicado que, a partir de una lista de cadenas de
parámetro, filtre aquellas que contengan una subcadena que el usuario
indique en otro argumento
*/

sub_cadenas(Subcadena, Lista, Filtradas) :-
    sub_cadenas_aux(Subcadena, Lista, Filtradas, []).

sub_cadenas_aux(_, [], Filtradas, Filtradas).
sub_cadenas_aux(Subcadena, [Cabeza | Cola], Filtradas, Acumulador) :-
    (   sub_atom(Cabeza, _, _, _, Subcadena) ->
        sub_cadenas_aux(Subcadena, Cola, Filtradas, [Cabeza | Acumulador])
    ;   sub_cadenas_aux(Subcadena, Cola, Filtradas, Acumulador)
    ).


/*Resultados
?- sub_cadenas("la", ["la tarea", "el examen", "el proyecto", "la asistencia"],Filtradas).
Filtradas = ["la asistencia", "la tarea"]
*/
