package HW_3

import (
	"fmt"
	"math"
)

const fieldSize = 8
const sharpSymbol = "#"

func RunChessGame() {
	for i := 0; i < fieldSize; i++ {
		for j := 0; j < fieldSize; j++ {
			if math.Mod(float64(i+j), 2) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(sharpSymbol)
			}
		}
		fmt.Println()
	}
}
