# Conway's Game of Life
- zero-player game simulating life of cells
- [Wikipedia article](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
- this is my personal CLI implementation in Go

## Usage
```bash
go build conway.go
./conway width height file
```
- where width and height are unsigned integers that determine the size of field
- and file is a file with the starting seed

## Seed
- The only influence the user has on the Game of Life is the starting position determined by the seed
- The seed for my implementation has the following format:
```
0, 0
9, 3
```
- where the tuples of numbers determine the starting cells which are alive
- there are example seeds in the repository ``block``, ``blinker`` and ``glider``

## Example
```
./conway 10 10 glider
Generation  0 :

_X________
__X_______
XXX_______
__________
__________
__________
__________
__________
__________
__________

Pres return for next generation, Ctrl-c to stop.
```


