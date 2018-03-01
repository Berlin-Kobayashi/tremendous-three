package main

import (
	"flag"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type Simulation struct {
	Rows, Columns, Vehicles, Rides, Bonus, Steps int
}

type Ride struct {
	StartX, StartY, EndX, EndY, Earliest, Latest int
}

func main() {
	var inputPath string
	flag.StringVar(&inputPath, "in", "", "The path to the input file")

	var outPath string
	flag.StringVar(&outPath, "out", "", "The path to the output file")

	flag.Parse()

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []string{}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	simulation := createSimulation(firstLine)

	fmt.Println(simulation)

	rides := []Ride{}

	for scanner.Scan() {
		currentLine := scanner.Text()

		rides = append(rides, createRide(currentLine))
	}

	fmt.Println(rides)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	ioutil.WriteFile(outPath, []byte(strings.Join(data, ",")), 0644)
}

func createSimulation(line string) Simulation {
	intSlice := toIntSlice(line)

	return Simulation{
		intSlice[0],
		intSlice[1],
		intSlice[2],
		intSlice[3],
		intSlice[4],
		intSlice[5],
	}
}

func createRide(line string) Ride {
	intSlice := toIntSlice(line)

	return Ride{
		intSlice[0],
		intSlice[1],
		intSlice[2],
		intSlice[3],
		intSlice[4],
		intSlice[5],
	}
}

func toIntSlice(s string) []int {
	parts := strings.Split(s, " ")
	intSlice := make([]int, len(parts))

	for i, firstLinePart := range parts {
		firstLineInt, err := strconv.Atoi(firstLinePart)
		if err != nil {
			panic(err)
		}

		intSlice[i] = firstLineInt
	}

	return intSlice
}
