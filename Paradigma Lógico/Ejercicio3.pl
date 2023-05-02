
/*
Tercer ejercicio
Implemente la funci�n aplanar. Esta funci�n recibe una lista con
m�ltiples listas anidadas como elementos y devuelve una lista con los
mismos elementos de manera lineal (sin listas).
*/

aplanar([], []).
aplanar([Cabeza | Cola], Aplanado) :-
    es_lista(Cabeza),
    aplanar(Cabeza, AplanadoCabeza),
    aplanar(Cola, AplanadoCola),
    concatenar(AplanadoCabeza, AplanadoCola, Aplanado).
aplanar([Cabeza | Cola], [Cabeza | AplanadoCola]) :-
    not(es_lista(Cabeza)),
    aplanar(Cola, AplanadoCola).

es_lista([]).
es_lista([_ | _]).

concatenar([], L, L).
concatenar([H|T], L2, [H|L3]) :- concatenar(T, L2, L3).

/*Resultados
?- aplanar([1,[8],9,[4,[3,2],6]], X).
X = [1, 8, 9, 4, 3, 2, 6] .
*/
