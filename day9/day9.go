package day9

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {
	diskMap, _ := parseInput(input)

	var cSum = 0
	for true {
		// This is too slow!
		newSum := checkSum(diskMap)
		if cSum == newSum {
			break
		}
		cSum = newSum

		back := diskMap.Back()
		for front := diskMap.Front(); front != nil; front = front.Next() {
			if front == back {
				break
			}
			frontVal, backVal := front.Value.(DiskMapEntry), back.Value.(DiskMapEntry)
			if front.Value.(DiskMapEntry).Empty > 0 {
				// Prepare block from back
				backVal.Size = backVal.Size - 1
				backVal.Empty = backVal.Empty + 1
				// Prepare new block
				newBlock := DiskMapEntry{
					ID:    backVal.ID,
					Size:  1,
					Empty: frontVal.Empty - 1,
				}
				// Prepare front block
				frontVal.Empty = 0
				// Update front block
				front.Value = frontVal
				// Update back block or delete if size zero
				if backVal.Size == 0 {
					// TODO: When removing the back block
					// Carry over empty space
					diskMap.Remove(back)
				} else {
					back.Value = backVal
				}
				// Insert new block
				diskMap.InsertAfter(newBlock, front)
				break
			}
		}
	}

	return checkSum(diskMap), err
}

func Part2(input string) (result int, err error) {
	diskMap, _ := parseInput(input)

	fragged := make(map[int]bool)
	back := diskMap.Back()
	for true {
		if back == nil {
			break
		}
		backBlock := back.Value.(DiskMapEntry)

		// Skip if fragged
		if _, ok := fragged[backBlock.ID]; ok {

			back = back.Prev()
			continue
		}
		var updatedCursor bool
		// Compare with blocks in the front for a space
		for front := diskMap.Front(); front != nil; front = front.Next() {
			if front == back {
				break
			}
			frontBlock := front.Value.(DiskMapEntry)
			// is the front space enough for back block
			if frontBlock.Empty >= backBlock.Size {

				fragged[backBlock.ID] = true
				// Is next to each other
				if front.Next() == back {
					// Move front spaces to back
					backBlock.Empty += frontBlock.Empty
					back.Value = backBlock
					// Clear front spaces
					frontBlock.Empty = 0
					front.Value = frontBlock
				} else {
					// Modify back block's parent with new space
					parentBlock := back.Prev().Value.(DiskMapEntry)
					parentBlock.Empty += backBlock.Size + backBlock.Empty
					back.Prev().Value = parentBlock
					// Modify back block with fronts space
					backBlock.Empty = frontBlock.Empty - backBlock.Size
					back.Value = backBlock
					// Modify front block
					frontBlock.Empty = 0
					front.Value = frontBlock
					// Move to front & update cursor
					newBack := back.Prev()
					diskMap.MoveAfter(back, front)
					back = newBack
					updatedCursor = true
				}
				break
			}
		}
		// Update cursor
		if !updatedCursor {
			back = back.Prev()
		}
	}

	return checkSum(diskMap), err
}

func checkSum(diskMap *list.List) (checkSum int) {
	var blockPos int
	for e := diskMap.Front(); e != nil; e = e.Next() {
		cur := e.Value.(DiskMapEntry)
		for blockIdx := 0; blockIdx < cur.Size; blockIdx++ {
			checkSum += blockPos * cur.ID
			blockPos++
		}
		blockPos += cur.Empty
	}
	return
}

func viz(diskMap *list.List) {
	for e := diskMap.Front(); e != nil; e = e.Next() {
		cur := e.Value.(DiskMapEntry)
		for i := 0; i < cur.Size; i++ {
			fmt.Print(cur.ID)
		}
		for i := 0; i < cur.Empty; i++ {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

type DiskMapEntry struct {
	ID    int
	Size  int
	Empty int
}

func parseInput(input string) (diskMap *list.List, count int) {
	diskMap = list.New()
	input = strings.TrimSpace(input)
	for i := 0; i < len(input); i += 2 {
		var entry DiskMapEntry
		entry.ID = count
		entry.Size, _ = strconv.Atoi(string(input[i]))
		if i+1 < len(input) {
			entry.Empty, _ = strconv.Atoi(string(input[i+1]))
		}
		diskMap.PushBack(entry)
		count += 1
	}
	return
}
