package main

import ( 
    "fmt"
)

func main() {
    gameBoard := GameBoard{}
    initialize(&gameBoard)
    (&gameBoard).tokenOne = 'X'
    (&gameBoard).tokenTwo = '0'
    var (
        col, turn   int
        gameOver    bool = false
        token       byte

    )
    for !gameOver {
        if turn % 2 == 0 {
            token = (&gameBoard).tokenOne
        } else {
            token = (&gameBoard).tokenTwo
        }
        turn = (turn + 1) % 2
        fmt.Println(string(token) + "'s turn. enter column:")
        n, err := fmt.Scanf("%d\n", &col)
        if n < 1 {
            fmt.Println(err)
        }
        insertToken(&gameBoard, token, col)
        if checkWinner(&gameBoard, token){
            gameOver = true
        }
        fmt.Println(toString(&gameBoard))
    }
}