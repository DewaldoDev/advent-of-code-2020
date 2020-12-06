package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := []int{}
	readFile(func(line string) {
		num, _ := strconv.Atoi(line)
		inputs = append(inputs, num)
	})

	fmt.Println(inputs)
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
