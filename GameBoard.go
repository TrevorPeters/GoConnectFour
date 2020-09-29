package main
import "strings"
type GameBoard struct {
	tokenOne byte
	tokenTwo byte
	game [6][7]byte
	
}

func initialize(g *GameBoard){
	g.game = [6][7]byte{{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '}} 
}

func toString(g *GameBoard) string {
	var sb strings.Builder
	var row, column int
	sb.WriteString("+---------------+\n")
	for row = 0; row < 6; row++ {
		sb.WriteString("| ")
		for column = 0; column < 7; column++ {
			sb.WriteString(string(g.game[row][column]))
			sb.WriteString(" ")
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("+---------------+\n")
	sb.WriteString("  0 1 2 3 4 5 6")
	return sb.String()
}

func insertToken (g *GameBoard, token byte, column int) bool {
	var success bool = false
	var row int = len(g.game) - 1
	for row >= 0 {
		if g.game[row][column] == ' ' {
			success = true
			g.game[row][column] = token
			break
		}
		row--
	}
	return success
}
					   
