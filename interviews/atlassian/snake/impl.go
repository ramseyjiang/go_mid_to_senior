package snake

type SnakeGame struct {
	width     int
	height    int
	score     int
	foodIndex int
	food      [][]int
	snake     []Coord
}

type Coord struct {
	x int
	y int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		width:     width,
		height:    height,
		score:     0,
		foodIndex: 0,
		snake:     []Coord{{0, 0}},
		food:      food,
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

	// Self-Collision Check
	for i := 0; i < len(sg.snake)-1; i++ {
		if head.y == sg.snake[i].y && head.x == sg.snake[i].x {
			return -1
		}
	}

	// Please notice: the food coord, y is the first, x is the second.
	// Use the snake head position to check whether the snake ate or not.
	ateFood := sg.foodIndex < len(sg.food) && head.y == sg.food[sg.foodIndex][0] && head.x == sg.food[sg.foodIndex][1]
	if ateFood {
		sg.score++
		sg.foodIndex++
	} else {
		// update snake coord, at the beginning, the snake is no body.
		// If no food is eaten, the last segment of the snake (sg.snake[:len(sg.snake)-1]) is removed to simulate the snake's movement.
		sg.snake = sg.snake[:len(sg.snake)-1]
	}

	// update the snake length and position after each move is end.
	// It adds previous snake as a part of body and current snake head
	sg.snake = append([]Coord{head}, sg.snake...)

	return sg.score
}
