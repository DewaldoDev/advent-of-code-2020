package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile(func(line string) {
		fmt.Println(line)
	})
}

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
