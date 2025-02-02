package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type ResonantMap struct {
	City      [][]string
	Antennas  map[string][]Antenna
	AntiNodes map[string]AntiNode
	Wg        sync.WaitGroup
	Mu        *sync.Mutex
}

type Antenna struct {
	Frequency string
	X         int
	Y         int
}

type AntiNode struct {
	X int
	Y int
}

func main() {
	// main area to display puzzle answers
	data := ParsePuzzleInput()
	fmt.Printf("Part 1: %d\n", Part1(&data))
	fmt.Printf("Part 2: %d\n", Part2(&data))
}

func Part1(input *ResonantMap) int {
	for _, v := range input.Antennas {
		input.Wg.Add(1)
		go processAntennaGroup(input, v, false)
	}

	input.Wg.Wait()
	return len(input.AntiNodes)
}

func Part2(input *ResonantMap) int {
	for _, v := range input.Antennas {
		input.Wg.Add(1)
		go processAntennaGroup(input, v, true)
	}

	input.Wg.Wait()
	return len(input.AntiNodes)
}

func generateKey(x int, y int) string {
	nx := strconv.Itoa(x)
	ny := strconv.Itoa(y)
	return strings.Join([]string{nx, ny}, ";")
}

func processAntennaGroup(input *ResonantMap, antennaGroup []Antenna, expand bool) {
	defer input.Wg.Done()
	maxX := len(input.City)
	maxY := len(input.City[0])

	var antinodes []AntiNode
	for i, curr := range antennaGroup {
		for j, others := range antennaGroup {
			if i == j {
				continue
			}

			if expand {
				antinodes = append(antinodes, expandAntiNode(curr, others, maxX, maxY)...)
			} else {
				antinodes = append(antinodes, placeAntiNode(curr, others, maxX, maxY)...)
			}
		}

		for _, anti := range antinodes {
			key := generateKey(anti.X, anti.Y)
			input.Mu.Lock()
			input.AntiNodes[key] = anti
			input.Mu.Unlock()
		}
	}
}

func expandAntiNode(a1 Antenna, a2 Antenna, maxX int, maxY int) []AntiNode {
	var antinodes []AntiNode
	// calculate the positions
	xDiff := a2.X - a1.X
	yDiff := a2.Y - a1.Y

	positions := []struct{ nx, ny int }{
		{a1.X - xDiff, a1.Y - yDiff}, // a1 Opposite direction
		{a2.X, a2.Y},                 // a2 Same direction
	}

	for {
		finished := 0

		// check if both points are out of bounds
		for _, pos := range positions {
			if pos.nx < 0 || pos.nx >= maxX || pos.ny < 0 || pos.ny >= maxY {
				finished++
			}
		}
		// if both are, then stop processing antinodes
		if finished == 2 {
			break
		}

		for _, pos := range positions {
			if pos.nx >= 0 && pos.nx < maxX && pos.ny >= 0 && pos.ny < maxY {
				antinodes = append(antinodes, AntiNode{
					X: pos.nx,
					Y: pos.ny,
				})
			}
		}

		positions = []struct{ nx, ny int }{
			{positions[0].nx - xDiff, positions[0].ny - yDiff}, // a1 Opposite direction
			{positions[1].nx + xDiff, positions[1].ny + yDiff}, // a2 Same direction
		}
	}

	return antinodes
}

// Use Manhattan Distance
func placeAntiNode(a1 Antenna, a2 Antenna, maxX int, maxY int) []AntiNode {
	var antinodes []AntiNode
	// calculate the positions
	xDiff := a2.X - a1.X
	yDiff := a2.Y - a1.Y

	positions := []struct{ nx, ny int }{
		{a1.X - xDiff, a1.Y - yDiff}, // a1 Opposite direction
		{a2.X + xDiff, a2.Y + yDiff}, // a2 Same direction
	}

	for _, pos := range positions {
		if pos.nx >= 0 && pos.nx < maxX && pos.ny >= 0 && pos.ny < maxY {
			antinodes = append(antinodes, AntiNode{
				X: pos.nx,
				Y: pos.ny,
			})
		}
	}

	return antinodes
}

func ParsePuzzleInput() ResonantMap {
	file := utils.Get2DPuzzleInput("day8.txt", false)

	antennaMap := make(map[string][]Antenna)
	for i := range file {
		for j := range file[i] {
			if file[i][j] == "." {
				continue
			}

			nAntenna := Antenna{
				X:         j,
				Y:         i,
				Frequency: file[i][j],
			}

			list, ok := antennaMap[nAntenna.Frequency]
			if !ok {
				antennaMap[nAntenna.Frequency] = []Antenna{nAntenna}
			} else {
				antennaMap[nAntenna.Frequency] = append(list, nAntenna)
			}
		}
	}
	return ResonantMap{
		City:      file,
		Antennas:  antennaMap,
		AntiNodes: make(map[string]AntiNode),
		Wg:        sync.WaitGroup{},
		Mu:        &sync.Mutex{},
	}
}

func PrintCity(cityMap [][]string) {
	for i := range cityMap {
		for _, v := range cityMap[i] {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}
