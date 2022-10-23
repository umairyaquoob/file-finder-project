package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1st B: To be exact, find the total size of all the empty files
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			// +1 for the newline character
			// when printing the filename afterward
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required space: %d bytes.\n", total)

	// 2nd: allocate a large enough byte slice in one go
	names := make([]byte, 0, total)

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()

			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
