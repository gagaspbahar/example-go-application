package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/model"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/service"
)

// main driver function and entry point
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	configService := &service.FileConfigService{Scanner: scanner}

	if err := run(configService, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
}

// run function to calculate the total distance
func run(configService service.ConfigService, output, errorOutput io.Writer) error {
	// get inputs
	field, err := configService.GetFieldDimensions()
	if err != nil {
		return err
	}

	treesData, err := configService.GetTrees(field.NumberOfTrees)
	if err != nil {
		return err
	}

	// create 2D array to store tree heights
	trees := make([][]int, field.Width+1)
	for i := range trees {
		trees[i] = make([]int, field.Length+1)
	}

	for _, tree := range treesData {
		trees[tree.X][tree.Y] = tree.Height
	}

	// calculate total distance
	totalDistance := calculateTotalDistance(field, trees)
	fmt.Fprintln(output, totalDistance)

	return nil
}

func calculateTotalDistance(field model.Field, trees [][]int) int {
	// initialize variables
	currentHeight := 1
	totalDistance := 1

	// iterate through the field zigzagging
	for i := 1; i <= field.Width; i++ {
		if i%2 != 0 {
			// case for odd rows
			for j := 1; j <= field.Length; j++ {
				totalDistance += 10
				if trees[i][j] == 0 {
					totalDistance += currentHeight - 1
					currentHeight = 1
				} else if trees[i][j] != (currentHeight - 1) {
					totalDistance += abs(currentHeight - (trees[i][j] + 1))
					currentHeight = trees[i][j] + 1
				}
			}
		} else {
			// case for even rows
			for j := field.Length; j > 0; j-- {
				totalDistance += 10
				if trees[i][j] == 0 {
					totalDistance += currentHeight - 1
					currentHeight = 1
				} else if trees[i][j] != (currentHeight - 1) {
					totalDistance += abs(currentHeight - (trees[i][j] + 1))
					currentHeight = trees[i][j] + 1
				}
			}
		}
	}

	// add the last height to the total distance
	totalDistance += currentHeight
	// remove the last 10 distance because of redundant addition in the last iteration
	totalDistance -= 10

	return totalDistance
}

// abs function to get the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
