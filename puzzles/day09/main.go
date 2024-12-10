package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type DiskMap struct {
	Blocks     map[int]*BlockFile
	Filesystem []int
	FileBlocks []*BlockFile
	Checksum   int
}

type BlockFile struct {
	Id        int
	BlockSize int
	FreeSpace int
	Attempted bool
}

func main() {
	data := ParsePuzzleInput()
	// main area to display puzzle answers
	fmt.Printf("Part 1: %d\n", Part1(data))

	data = ParsePuzzleInput()
	fmt.Printf("Part 2: %d\n", Part2(data))
}

func Part1(data DiskMap) int {
	moveBlocksAndCalcChecksum(&data)
	return data.Checksum
}

func moveBlocksAndCalcChecksum(data *DiskMap) {
	filesystem := &data.Filesystem
	// fmt.Println(*filesystem)
	for lidx := findLeftmostSpace(*filesystem); lidx < len(*filesystem); lidx++ {
		// check if we are done swapping
		if stopSwapping((*filesystem)[lidx+1:]) {
			break
		}

		ridx := findRightmostSpace(*filesystem)
		if (*filesystem)[lidx] == -1 && (*filesystem)[ridx] != -1 {
			utils.Swap(*filesystem, lidx, ridx)
			// fmt.Println(*filesystem)
		}
	}

	data.Checksum = 0
	for i := range *filesystem {
		id := (*filesystem)[i]
		n, ok := data.Blocks[id]
		if !ok {
			continue
		}

		data.Checksum += i * n.Id
	}
}

func Part2(data DiskMap) int {
	moveFiles(&data)
	return data.Checksum
}

func moveFiles(data *DiskMap) {
	filesystem := data.Filesystem
	// start with the highest idx
	for ridx := findRightmostSpace(filesystem); ridx >= 0; ridx-- {
		c, ok := data.Blocks[filesystem[ridx]]
		if !ok || c.Attempted {
			continue
		}

		count := 0
		startingPoint := findLeftmostSpace(filesystem)
		found := false
		for lidx := startingPoint; lidx < ridx; lidx++ {
			if filesystem[lidx] == -1 {
				if count == 0 {
					startingPoint = lidx
				}
				count++
				if count == c.BlockSize {
					found = true
					break
				}
			} else {
				count = 0
			}
		}

		if found {
			for i := 0; i < c.BlockSize; i++ {
				utils.Swap(filesystem, ridx-i, startingPoint+i)
			}
		}

		c.Attempted = true
	}

	data.Checksum = 0
	for i := range filesystem {
		id := filesystem[i]
		if id == -1 {
			continue
		}
		data.Checksum += i * id
	}
}

func stopSwapping(filesystem []int) bool {
	for i := range filesystem {
		if filesystem[i] != -1 {
			return false
		}
	}

	return true
}

func findLeftmostSpace(filesystem []int) int {
	for i := range filesystem {
		if filesystem[i] == -1 {
			return i
		}
	}

	panic("NO SPACES FOUND")
}

func findRightmostSpace(filesystem []int) int {
	for i := len(filesystem) - 1; i >= 0; i-- {
		if filesystem[i] != -1 {
			return i
		}
	}

	panic("COULD NOT FIND A SPACE")
}

func ParsePuzzleInput() DiskMap {
	// The digits alternate between indicating the length of a file and the length of free space.
	file := utils.GetPuzzleInput("day9.txt", false)

	scn := bufio.NewScanner(file)
	scn.Scan()
	data := scn.Text()

	diskMap := DiskMap{
		Blocks: make(map[int]*BlockFile),
	}
	id := 0
	var filesystem []int
	var fileblocks []*BlockFile

	for i := 0; i < len(data); i += 2 {
		var bs, fs int

		bs, _ = strconv.Atoi(string(data[i]))
		if i+1 < len(data) {
			fs, _ = strconv.Atoi(string(data[i+1]))
		} else {
			fs = 0
		}

		c := BlockFile{
			Id:        id,
			BlockSize: bs,
			FreeSpace: fs,
			Attempted: false,
		}

		filesystem = append(filesystem, generatePair(id, bs, fs)...)

		diskMap.Blocks[id] = &c
		fileblocks = append(fileblocks, &c)
		id++
	}

	diskMap.Filesystem = filesystem
	diskMap.FileBlocks = fileblocks
	return diskMap
}

func generatePair(id int, bs int, fs int) []int {
	b := make([]int, bs)
	s := make([]int, fs)

	for i := range bs {
		b[i] = id
	}
	for i := range fs {
		s[i] = -1
	}

	return append(b, s...)
}
