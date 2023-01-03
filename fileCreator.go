package main

import (
	"bufio"
	"fmt"
	"os"	
)

func createFile(filename string) File {
	file, err := os.Open(filename)
	var f File
	f.name = filename
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {		
		f.lines = append(f.lines, scanner.Text())
	}
	return f
}

type File struct {
	name  string
	lines []string
}