package main

import (
	"flag"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)

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
	for scanner.Scan() {
		currentLine := scanner.Text()

		data = append(data, currentLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	ioutil.WriteFile(outPath, []byte(strings.Join(data, ",")), 0644)
}
