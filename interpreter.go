package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ./gobrain filename.bf\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	if !strings.HasSuffix(filename, ".bf") {
		fmt.Fprintf(os.Stderr, "program filename must end with '.bf'\n")
		os.Exit(1)
	}

	program := getProgramFromFile(filename)
	output, err := run(program, bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encountered during program execution: %v", err)
		os.Exit(1)
	}

	fmt.Printf(output)
}

func getProgramFromFile(filename string) string {
	programBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open %v for reading\n", filename)
		os.Exit(1)
	}
	return string(programBytes)
}

func run(program string, reader *bufio.Reader) (string, error) {
	var output bytes.Buffer

	// 30000 bytes according to brainfuck "spec"
	var data [30000]byte

	dataPointer := 0
	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '>':
			dataPointer++
		case '<':
			dataPointer--
		case '+':
			data[dataPointer]++
		case '-':
			data[dataPointer]--
		case '.':
			output.WriteString(string(data[dataPointer]))
		case ',':
			b, err := reader.ReadByte()
			if err != nil {
				if err.Error() == "EOF" {
					continue
				}
				return "", errors.New(fmt.Sprintf("error encountered when reading input: %v\n", err))
			}
			data[dataPointer] = b
		case '[':
			if data[dataPointer] == 0 {
				// search forward for corresponding ']'
				loopStart := i
				i++
				for count := 1; count != 0; i++ {
					if i >= len(program) {
						return "", errors.New(fmt.Sprintf("Unterminated loop caught beginning at idx: %v\n", loopStart))
					}
					if program[i] == ']' {
						count--
					} else if program[i] == '[' {
						count++
					}
				}
			}
		case ']':
			if data[dataPointer] != 0 {
				// search backward for corresponding '['
				loopEnd := i
				i--
				for count := 1; count != 0; i-- {
					if i <= 0 {
						return "", errors.New(fmt.Sprintf("Encountered loop termination without opening bracket at idx: %v\n", loopEnd))
					}
					if program[i] == '[' {
						count--
					} else if program[i] == ']' {
						count++
					}
				}
			}
		}
	}
	return output.String(), nil
}
