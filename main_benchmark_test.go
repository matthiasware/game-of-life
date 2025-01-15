package main

import (
	"testing"
)

const (
	HEIGHT = 100
	WIDTH  = 100
	STEPS  = 100
	SEED   = 1023
)

func BenchmarkUpdateWorldSerial(b *testing.B) {
	world := NewWorld(HEIGHT, WIDTH)
	InitializeWorld(world, SEED)
	b.ResetTimer()
	for range b.N {
		world = UpdateWorldSerial(world)
	}
}

func BenchmarkUpdateWorldParallel(b *testing.B) {
	world := NewWorld(HEIGHT, WIDTH)
	InitializeWorld(world, SEED)
	b.ResetTimer()
	for range b.N {
		world = UpdateWorldParallel(world)
	}
}

//func BenchmarkEvolveWorldSerial(b *testing.B) {
//	world := NewWorld(HEIGHT, WIDTH)
//	InitializeWorld(world, SEED)
//	b.ResetTimer()
//
//	for range b.N {
//		world = EvolveWorldSerial(world, STEPS)
//	}
//}
