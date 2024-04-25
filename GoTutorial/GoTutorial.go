package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func scanForInputSuccessful() chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)

		file, err := os.Open("testdocumentsuccessful.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	return lines
}

func scanForInputFail() chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)

		file, err := os.Open("testdocumentfailed.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	return lines
}

func main() {
	linesSuccessful := scanForInputSuccessful()
	linesFail := scanForInputFail()

	line, ok := <-linesSuccessful

	if !ok {
		fmt.Println(line)
	} else {
		time.After(1 * time.Second)
		printHelpMessage()
	}

	line2, ok := <-linesFail

	if !ok {
		fmt.Println(line2)
	} else {
		time.After(1 * time.Second)
	}
	printHelpMessage()

}

func printHelpMessage() {
	fmt.Println("Help message")
}
