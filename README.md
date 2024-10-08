# command-line go-sudoku-solver

A command line sudoku solver built from scratch using go

![sudoku solver](./sudoku-solver.png)

---

## run locally

### prerequisites

- git
- go

### directions

- clone project

```git
git clone https://github.com/elameendaiyabu/go-sudoku-solver.git
```

- cd into project

```bash
cd go-sudoku-solver
```

- install required go packages

```go
go mod tidy
```

- go run

```go
go run .
```

## tech stack

- Golang
- Bubble Tea

## rules of the game

- a 9x9 grid must be filled with numbers between 1-9
- a number shouldnt be repeated in a row
- a number shouldnt be repeated in a column
- a number shouldnt be repeated in a sub-grid
- a subgrid is a 3x3 grid inside the original 9x9
- there a 9 subgrids on a board
- a board should only have one unique solution

## goals of this project

- Goal 1: represent the sudoku board (achieved)
- Goal 2: create algorithm to solve the board (achieved - used backtracking)
- Goal 3: create a tui using bubble tea (achieved)

## quick pseudocode

- use arrays to represent the board
- cells with no number will be represented with 0
- i will use a backtracking algorithm to solve the board
- i will have three main functions, checkColumn(), checkRow(), checkSubGrid()
- checkColumn : checks column of the current cell it is on whether the number exists or not
- checkRow : checks row of the current cell it is on whether the number exists or not
- checkSubGrid : checks subgrid of the current cell it is on whether the number exists or not
- if the number doesnt exist in either the column, row, or subgrid, it moves on to the next empty cell.
- if it exists, move back one cell, increment it by one, then try the functions again.
- this is how it will work until all cells are filled and a solution is gotten.
