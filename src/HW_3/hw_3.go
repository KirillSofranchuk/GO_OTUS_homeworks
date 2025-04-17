package HW_3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const sharpSymbol = "#"
const maxFieldSize = 10
const minFieldSize = 2
const defaultFieldSize = 8

func RunChessGame() {
	fieldSize := getFieldSize()

	fmt.Println("Chess field")

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

func getFieldSize() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите размер поля. От %d до %d\n", minFieldSize, maxFieldSize)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fieldSize, err := strconv.Atoi(input)
	if err != nil || fieldSize > maxFieldSize || fieldSize < minFieldSize {
		fmt.Printf("Ошибка: введен неверный размер поля. Будет задано стандартное значение: %d \n", defaultFieldSize)
		return defaultFieldSize
	}

	return fieldSize
}
