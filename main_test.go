package main

import (
	"fmt"
	"testing"
)

func TestCountAliveNeighbors(t *testing.T) {
	var w World
	var nNeighbors [][]int
	w = World{
		3, 3,
		[][]bool{
			[]bool{false, false, false},
			[]bool{false, true, true},
			[]bool{false, true, true}},
	}
	nNeighbors = [][]int{
		[]int{1, 2, 2},
		[]int{2, 3, 3},
		[]int{2, 3, 3},
	}

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			nExp := nNeighbors[y][x]
			nAct := CountAliveNeighbors(w, x, y)
			if nAct != nExp {
				t.Fatalf("%d,%d: expected %d, got %d", x, y, nExp, nAct)
			}
		}
	}
}

func compareWorlds(a, b World) (bool, string) {
	if a.height != b.height {
		return false, "Worlds do not have same height!"
	}
	if a.width != b.width {
		return false, "Worlds do not have same width!"
	}
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			if a.grid[y][x] != b.grid[y][x] {
				return false, fmt.Sprintf("Cell (%d, %d) not equal", x, y)
			}
		}
	}
	return true, ""
}

func TestUpdateWorld(t *testing.T) {
	now_exp := [][]World{
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
		},
		// rule 1: any live cell with fewer than two live neighbors dies due to underpopulation
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{true, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, true, true},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
		},
		// rule 2: any live cell with two or three live neighbours lives on the next generation
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, false},
					[]bool{true, true, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, false},
					[]bool{true, true, false},
					[]bool{false, false, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, false},
					[]bool{true, true, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, false},
					[]bool{true, true, false},
					[]bool{false, false, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, true},
					[]bool{false, true, false},
					[]bool{true, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, true, false},
					[]bool{false, false, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, true, false},
					[]bool{true, false, true}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, false, false},
					[]bool{false, true, false},
					[]bool{false, true, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, true},
					[]bool{false, false, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{false, true, false},
					[]bool{false, true, false},
					[]bool{false, false, false}},
			},
		},
		[]World{
			World{
				3, 3,
				[][]bool{
					[]bool{true, true, true},
					[]bool{true, true, false},
					[]bool{false, false, false}},
			},
			World{
				3, 3,
				[][]bool{
					[]bool{true, false, true},
					[]bool{true, false, true},
					[]bool{false, false, false}},
			},
		},
	}

	for _, worlds := range now_exp {
		wNow := worlds[0]
		wExp := worlds[1]
		wAct := UpdateWorldSerial(wNow)
		if equal, msg := compareWorlds(wAct, wExp); !equal {
			fmt.Println("Initial")
			DisplayWorld(wNow)
			fmt.Println("Expected")
			DisplayWorld(wExp)
			fmt.Println("Got")
			DisplayWorld(wAct)
			t.Fatal(msg)

		}
	}
}
