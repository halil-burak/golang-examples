package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("created.")
	return f
}

func closeFile(file *os.File) {
	fmt.Println("closing file...")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("closed file.")
	}
}

func writeFile(file *os.File) {
	fmt.Println("writing to file...")
	fmt.Fprintln(file, "hello world!")
	fmt.Println("written to file.")
}
