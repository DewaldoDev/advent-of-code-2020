package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	code := []string{}
	readFile(func(line string) {
		code = append(code, line)
	})

	acc := 0
	line := 0
	visited := make(map[int]bool) // codePoints[234] true or false

	executeLine(&code, acc, line, &visited)
}

func executeLine(code *[]string, acc int, line int, visited *map[int]bool) {
	fmt.Println(code, acc, line, visited)
}

// readFile reads the input file line by line and passes each line to the readerFunc
func readFile(readerFunc func(string)) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readerFunc(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
