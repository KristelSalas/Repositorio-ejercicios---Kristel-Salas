
/*
Segundo ejercicio
Construya una funci�n que se llame sub_conjunto. Esta funci�n recibe
dos listas y debe retornar True cuando el primer argumento es
subconjunto completo del segundo y false en caso contrario.
*/

sub_conjunto([], _).
sub_conjunto([Cabeza | Cola], Lista) :-
    pertenece(Cabeza, Lista),
    sub_conjunto(Cola, Lista).

pertenece(Elemento, [Elemento | _]).
pertenece(Elemento, [_ | Cola]) :-
    pertenece(Elemento, Cola).

/*Resultados
?- sub_conjunto([1,2,8],[1,2,3,4,5,6,7]).
false.

?- sub_conjunto([1,2,8],[1,2,3,4,5,6,7,8]).
true
*/
