package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

type World struct {
	height, width int
	grid          [][]bool
}

func NewWorld(width, height int) World {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	w := World{
		height,
		width,
		grid,
	}
	return w
}

func InitializeWorld(w World, seed int64) {
	source := rand.NewSource(seed)
	rng := rand.New(source)
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			w.grid[i][j] = rng.Intn(2) == 0
		}
	}
}

func DisplayWorld(w World) {
	// Clear screen
	ClearConsole()
	//fmt.Print("\033[H\033[2J")
	for _, row := range w.grid {
		for _, cell := range row {
			if cell {
				fmt.Print("⬜")
			} else {
				fmt.Print("⬛")
			}
		}
		fmt.Println()
	}
}

func CountAliveNeighbors(w World, x, y int) int {
	count := 0
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < w.width && ny >= 0 && ny < w.height && w.grid[ny][nx] == true {
			count++
		}
	}
	return count
}

func UpdateWorldSerial(w World) World {
	newWorld := NewWorld(w.width, w.height)
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			n := CountAliveNeighbors(w, x, y)
			if w.grid[y][x] {
				newWorld.grid[y][x] = n == 2 || n == 3
			} else {
				newWorld.grid[y][x] = n == 3
			}
		}
	}
	return newWorld
}

func UpdateRow(w World, newWorld World, y int) {
	for x := 0; x < w.width; x++ {
		n := CountAliveNeighbors(w, x, y)
		if w.grid[y][x] {
			newWorld.grid[y][x] = n == 2 || n == 3
		} else {
			newWorld.grid[y][x] = n == 3
		}
	}
}

func UpdateWorldParallel(w World) World {
	newWorld := NewWorld(w.width, w.height)
	var wg sync.WaitGroup
	for y := 0; y < w.height; y++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			UpdateRow(w, newWorld, y)
		}()
	}
	wg.Wait()
	return newWorld
}

func EvolveWorldParallel(w World, steps int) World {
	for range steps {
		w = UpdateWorldParallel(w)
	}
	return w
}

func EvolveWorldSerial(w World, steps int) World {
	for i := 0; i < steps; i++ {
		w = UpdateWorldSerial(w)
	}
	return w
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	height, width := 50, 50
	steps := 1000
	sleep := 10
	world := NewWorld(height, width)
	InitializeWorld(world, time.Now().UnixNano())
	for _ = range steps {
		DisplayWorld(world)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		world = UpdateWorldParallel(world)
	}
}
