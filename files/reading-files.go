package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// the most basic file reading task is slurping a file’s entire contents into memory.
	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	fmt.Println(string(dat))

	// more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.
	f, err := os.Open("./test.txt")
	check(err)

	// defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("b1 read %s\n", string(b1))
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes read @%d: %s\n", n2, o2, string(b2))
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but Seek(0, 0) accomplishes this.
	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package implements a buffered reader that may be useful both for its efficiency with
	// many small reads and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes read: %s\n", string(b4))

	f.Close()
}
