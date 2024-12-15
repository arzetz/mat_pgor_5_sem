package main

import (
	"fmt"
	"time"

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
	posX, posY  = 1, 1
	destination = [2]int{9, 6}
)

type Point struct {
	x, y int
}

func waveAlgorithm(startX, startY, endX, endY int) [][]int {
	distance := make([][]int, height)
	for i := range distance {
		distance[i] = make([]int, width)
		for j := range distance[i] {
			distance[i][j] = -1
		}
	}

	distance[startY][startX] = 0
	queue := []Point{{x: startX, y: startY}}
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			nx, ny := current.x+d[0], current.y+d[1]
			if nx >= 0 && nx < width && ny >= 0 && ny < height && maze[ny][nx] != 1 && distance[ny][nx] == -1 {
				distance[ny][nx] = distance[current.y][current.x] + 1
				queue = append(queue, Point{x: nx, y: ny})
			}
		}
	}

	return distance
}

func reconstructPath(distance [][]int, endX, endY int) []Point {
	if distance[endY][endX] == -1 {
		return nil // Путь не найден
	}

	path := []Point{{x: endX, y: endY}}
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	current := Point{x: endX, y: endY}
	for distance[current.y][current.x] != 0 {
		for _, d := range directions {
			nx, ny := current.x+d[0], current.y+d[1]
			if nx >= 0 && nx < width && ny >= 0 && ny < height && distance[ny][nx] == distance[current.y][current.x]-1 {
				current = Point{x: nx, y: ny}
				path = append([]Point{current}, path...)
				break
			}
		}
	}

	return path
}

func drawMaze(path []Point) {
	pathMap := make(map[Point]bool)
	for _, p := range path {
		pathMap[p] = true
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == posX && y == posY {
				fmt.Print("O ")
			} else if pathMap[Point{x: x, y: y}] {
				fmt.Print("* ")
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

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Используйте стрелки. Чтобы выйти, нажмите ESC")
	startTime := time.Now()
	distance := waveAlgorithm(posX, posY, destination[0], destination[1])
	path := reconstructPath(distance, destination[0], destination[1])
	steps := new(int)
	*steps = 0
	for {
		drawMaze(path)
		_, key, _ := keyboard.GetKey()
		if key == keyboard.KeyEsc {
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			moveObject(0, -1, steps)
		case keyboard.KeyArrowDown:
			moveObject(0, 1, steps)
		case keyboard.KeyArrowLeft:
			moveObject(-1, 0, steps)
		case keyboard.KeyArrowRight:
			moveObject(1, 0, steps)
		}
		fmt.Print("\033[H\033[2J") // Очистка экрана
	}
	duration := time.Since(startTime)
	fmt.Printf("\nВремя прохождения: %v\nДлина пути: %d\nКоличество шагов алгоритма: %d\nКоличество шагов пользователя: %d\n", duration, len(path)-1, len(path)-1, *steps)
}

func moveObject(dx, dy int, steps *int) {
	newX, newY := posX+dx, posY+dy
	if newX >= 0 && newX < width && newY >= 0 && newY < height && maze[newY][newX] != 1 {
		posX, posY = newX, newY
		*steps++
	}
}
