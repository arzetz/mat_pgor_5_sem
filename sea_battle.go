package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func insertion() ([][]string, map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	sea_map := make([][]string, 0, 5)
	sea_map_hash := make(map[string]string)
	for i := 0; i < 5; i++ {
		scanner.Scan()
		expression := strings.Fields(scanner.Text())
		sea_map = append(sea_map, expression)
		for j, num := range expression {
			key := fmt.Sprintf("%c%d", 'A'+j, i+1)
			sea_map_hash[key] = num
		}
	}
	return sea_map, sea_map_hash
}

func shot(sea_map [][]string, sea_map_hash map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for godem := 0; godem < 3; {
		fmt.Print("\nВведите точку выстрела: ")
		scanner.Scan()
		shot_key := scanner.Text()

		if sea_map_hash[shot_key] == "1" {
			fmt.Print("\nПопал!\n")
			sea_map_hash[shot_key] = "X"
			godem++
		} else {
			fmt.Print("\nНе попал!\n")
			sea_map_hash[shot_key] = "L"
		}
		fmt.Print("  A B C D E")

		for i, col := range sea_map {
			fmt.Printf("\n%d", i+1)
			for j, _ := range col {
				key := fmt.Sprintf("%c%d", 'A'+j, i+1)
				if key == shot_key {
					sea_map[i][j] = sea_map_hash[shot_key]
				}
				fmt.Printf(" %s", sea_map[i][j])
			}
		}
	}
	fmt.Println("\n\nWinner-winner chicken dinner!")
	return
}

func main() {
	fmt.Println("\nРазместите корабли на вашей карте морского боя:")
	letter := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		letter = append(letter, string('A'+i))
		print(letter[i] + " ")
	}
	fmt.Println()
	shot(insertion())
	return
}
