/*Detalles: profesor yo elimine unas conexiones porque según la imagen que utilizó en el video como referencia no deben estar conectados los puntos (22,28) y el (1,7) se encontraba repetido */

conectado(i,2).
conectado(2,3).
conectado(2,8).
conectado(8,9).
conectado(9,3).
conectado(3,4).
conectado(4,10).
conectado(10,16).
conectado(16,22).
conectado(22,21).
conectado(21,15).
conectado(15,14).
conectado(14,13).
conectado(13,7).
conectado(7,1).
conectado(14,20).
conectado(20,26).
conectado(26,27).
conectado(27,28).
conectado(28,29).
conectado(29,23).
conectado(23,17).
conectado(17,11).
conectado(11,5).
conectado(5,6).
conectado(28,34).
conectado(34,35).
conectado(35,36).
conectado(36,30).
conectado(30,24).
conectado(24,18).
conectado(18,12).
conectado(32,31).
conectado(31,25).
conectado(25,19).
conectado(34,33).
conectado(33,32).
conectado(32,f).


conectados(X,Y):- conectado(X,Y).
conectados(X,Y):- conectado(Y,X).

ruta(Fin,Fin,LRuta,Res):-
    reverse([Fin|LRuta],Res).
ruta(Ini,Fin,LRuta,Res):-
    conectados(Ini,Z),
    not(member(Z,LRuta)),
    ruta(Z,Fin,[Ini|LRuta],Res).

todasLasRutas(Ini,Fin,Rutas):-
    findall(Ruta,ruta(Ini,Fin,[],Ruta),Rutas).

%para imprimir las rutas sin que salgan cortadas
imprimir_rutas([]).
imprimir_rutas([Ruta|Resto]):-
    write(Ruta), nl,
    imprimir_rutas(Resto).


/*Resultados
?- todasLasRutas(i,f,Rutas), imprimir_rutas(Rutas).
[i,2,3,4,10,16,22,21,15,14,20,26,27,28,34,33,32,f]
[i,2,8,9,3,4,10,16,22,21,15,14,20,26,27,28,34,33,32,f]
Rutas = [[i, 2, 3, 4, 10, 16, 22, 21|...], [i, 2, 8, 9, 3, 4, 10|...]].
*/

