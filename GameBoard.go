package main
import (
	"strings"
)

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
//Inserts token into specified column. 
//Returns true on success and false otherwise
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

//Returns true if a winstate is found (four in a row)
func checkWinner (g *GameBoard, token byte) bool{	
	var row int = len(g.game) - 1
	var goalStr string = string([]byte{token, token, token, token})
	//check horizontal goal
	for row >= 0 {
		str := string(g.game[row][:])
			if strings.Contains(str, goalStr){
				return true
			}
		row--
	}

	var column int = len(g.game[0]) - 1
	var sb strings.Builder
	//check vertical goal
	for column >= 0 {
		row = len(g.game) - 1
		for row >= 0 {
			sb.WriteString(string(g.game[row][column]))
			row--
		}
		str := sb.String()
		sb.Reset()
		if strings.Contains(str, goalStr){
			return true
		}
		column--
	}	
	column = 0
	//check horizontal goal (\)
	for column < len(g.game[0]) - 3 {
		row = 0;
		for row < len(g.game) - 3{
			var i int = 0 
			for row + i < len(g.game) && column + i < len(g.game[0]){
				sb.WriteString(string(g.game[row + i][column + i]))
				i++
			}
			row++
			str := sb.String()
			sb.Reset()
			if strings.Contains(str, goalStr){
				return true
			}
		}
		column++
	}

	column = len(g.game[0]) - 1 
	//check horizontal goal (/)
	for column > 2 {
		row = 0;
		for row < len(g.game) - 3 {
			var i int = 0 
			for row + i < len(g.game) && column - i >= 0{
				sb.WriteString(string(g.game[row + i][column - i]))
				i++
			}
			row++
			str := sb.String()
			sb.Reset()
			if strings.Contains(str, goalStr){
				return true
			}
		}
		column--
	}
	return false
}
					   
