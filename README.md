# conc

UNIVERSIDAD PERUANA DE CIENCIAS APLICADAS

FACULTAD DE INGENIERÍA

CARRERA DE CIENCIAS DE LA COMPUTACIÓN





PROGRAMACIÓN CONCURRENTE Y DISTRIBUIDA

TAREA ACADEMICA 04




Integrantes :						Codigo:
Cesar Osorio	Guerra					U201314396
Enzo Camargo Ramírez           U202010122
Henry Josue Diaz Huarcaya				U20201C579






Lima, 2023-02

Introduccion

En este proyecto realizamos un algoritmo para simular el juego de Ludo modificado. El juego involucra a múltiples jugadores compitiendo para llevar a sus personajes a través de un laberinto lleno de obstáculos. Utilizaremos programación concurrente y canales para gestionar la comunicación entre jugadores y el tablero del juego.
El algoritmo se centrará en implementar canales bidireccionales para que los jugadores envíen resultados de lanzamientos de dados y reciban instrucciones de movimiento. La lógica del movimiento garantizará que los personajes avancen o retrocedan según las reglas del juego, teniendo en cuenta obstáculos y operaciones matemáticas. También implementaremos una gestión de turnos para que los movimientos se realicen de manera sincronizada.
El objetivo final del algoritmo es detectar cuándo un jugador ha llevado a todos sus personajes a la meta y declararlo como ganador. Para esto tendrá que pasar por varios obstáculos y hacerlo en el menor tiempo posible. La simulación mostrará el progreso del juego en tiempo real.



Planeamiento

Este algoritmo utiliza la programación concurrente para habilitar la competencia sincronizada y justa entre múltiples jugadores en el juego de Laberinto modificado. Los canales de comunicación aseguran una interacción eficiente entre los jugadores y el tablero. La lógica del movimiento y la gestión de obstáculos se implementan de manera coherente, y la detección del ganador determina el punto de finalización del juego.
 En el código, se establecen constantes que definen parámetros esenciales como el tamaño del tablero, el número de jugadores, la cantidad máxima de personajes por jugador y la posición de victoria. Variables globales representan el estado del tablero, los canales de comunicación entre los jugadores, un indicador de juego, el jugador actual y un grupo de espera para sincronizar las goroutines.
 La función principal inicializa la semilla del generador de números aleatorios, crea el tablero y establece los canales de comunicación. Posteriormente, se inicia un goroutine para cada jugador mediante la función jugarLudo(), lo que posibilita que los jugadores participen simultáneamente en el juego.
 La función jugarLudo() encapsula la lógica del juego para un jugador, continuando hasta que se cumple la condición de fin de juego. Un mutex (mu) se emplea para asegurar la sincronización entre jugadores, garantizando que un jugador complete su turno antes de que otro comience el suyo.
 Dentro de la función jugarLudo(), se verifica si es el turno del jugador actual, y se presenta un mensaje asociado a su turno. Luego, se lanzan los dados utilizando la función tirarDados() y se muestra el resultado obtenido. Se determina si el jugador ha ganado, verificando si todas sus fichas han alcanzado la posición de victoria. En caso de victoria, se emite un mensaje correspondiente y se establece el indicador de fin de juego. Si el jugador no ha ganado, se lleva a cabo el movimiento de fichas de acuerdo con las reglas del juego, y se actualiza su posición en el tablero. Finalmente, el turno se transfiere al siguiente jugador.
 La función tirarDados() genera un número aleatorio entre 1 y 6 para simular el lanzamiento de dados. La función Ganador verifica si un jugador ha ganado, comprobando si todas sus fichas han alcanzado la posición de victoria.



Conclusiones
La implementación del algoritmo para el juego de Laberinto modificado es eficiente en el uso de la programación concurrente y canales de comunicación. Los jugadores compiten de manera sincronizada gracias a canales bidireccionales que facilitan la interacción segura con el tablero.
 La lógica del movimiento de los personajes y la gestión de obstáculos se aplican de manera coherente. El uso de un mutex garantiza una sincronización adecuada entre los jugadores, evitando conflictos y asegurando que un jugador complete su turno antes de que otro comience.
 La detección del ganador y la conclusión del juego se maneja de manera eficaz. La verificación precisa de si un jugador ha ganado y la finalización inmediata del juego al declarar al ganador son aspectos importantes del algoritmo.
 La flexibilidad y escalabilidad se evidencian en la capacidad del algoritmo para adaptarse a cambios en el tamaño del tablero, el número de jugadores o reglas adicionales. La introducción de la generación aleatoria de dados añade un elemento de aleatoriedad, haciendo cada partida única y emocionante.

Github
https://github.com/cesar6793/conc
