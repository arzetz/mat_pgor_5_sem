package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

const (
	width  = 10
	height = 9
)

var (
	maze = [height][width]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 0, 3},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	posX, posY = 1, 1
)

func drawMaze() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == posX && y == posY {
				fmt.Print("O ")
			} else {
				switch maze[y][x] {
				case 0:
					fmt.Print(". ")
				case 1:
					fmt.Print("# ")
				case 3:
					fmt.Print("X ")
				}
			}
		}
		fmt.Println()
	}
}

func moveObject(dx, dy int) {
	newX, newY := posX+dx, posY+dy
	//Коллизия
	if newX >= 0 && newX < width && newY >= 0 && newY < height && (maze[newY][newX] == 0 || maze[newY][newX] == 3) {
		posX, posY = newX, newY
	}
	if maze[newY][newX] == 3 {
		posX, posY = 1, 1
	}
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Используйте стрелки. Чтобы выйти, нажмите ESC")
	for {
		drawMaze()
		_, key, _ := keyboard.GetKey()
		if key == keyboard.KeyEsc {
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			moveObject(0, -1)
		case keyboard.KeyArrowDown:
			moveObject(0, 1)
		case keyboard.KeyArrowLeft:
			moveObject(-1, 0)
		case keyboard.KeyArrowRight:
			moveObject(1, 0)
		}
		fmt.Print("\033[H\033[2J") //Очистка экрана
	}
}
