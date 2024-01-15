package snake

type SnakeGame struct {
	width, height int
	food          [][]int
	score         int
	snake         []Coord
	foodIndex     int
}

type Coord struct {
	x, y int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		width:     width,
		height:    height,
		food:      food,
		snake:     []Coord{{0, 0}},
		foodIndex: 0,
	}
}

func (sg *SnakeGame) Move(direction string) int {
	head := sg.snake[0]

	switch direction {
	case "U":
		head.y--
	case "D":
		head.y++
	case "L":
		head.x--
	case "R":
		head.x++
	}

	if head.y < 0 || head.y >= sg.height || head.x < 0 || head.x >= sg.width {
		return -1
	}

	ateFood := sg.foodIndex < len(sg.food) && head.y == sg.food[sg.foodIndex][0] && head.x == sg.food[sg.foodIndex][1]
	if ateFood {
		sg.score++
		sg.foodIndex++
	}

	for i := 0; i < len(sg.snake)-1; i++ {
		if head.y == sg.snake[i].y && head.x == sg.snake[i].x {
			return -1
		}
	}

	sg.snake = append([]Coord{head}, sg.snake...)

	if !ateFood {
		sg.snake = sg.snake[:len(sg.snake)-1]
	}

	return sg.score
}
