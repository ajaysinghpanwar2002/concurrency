## Channels
* A means of allowing communication to and from a GoRoutine
* Channels can be buffered, or unbuffered
* Once you're done with a channel, you must close it
* Channels typically only accept a given type or interface.

## The sleeping Barber
* A classic CS problem introduced by Dijkstra in 1965
* A barber goes to work in a barbershop with a waiting room with a fixed number of seats.
* If no one is in the waiting room, the barber goes to sleep.
* When a client shows up, if there are no seats available, he or she leaves.