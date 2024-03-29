package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func main() {
	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	sum := 0
	for _, line := range lines {
		lnum, err := strconv.Atoi(line)
		if err == nil {
			sum += lnum
		}
	}
	fmt.Println(sum)

}
