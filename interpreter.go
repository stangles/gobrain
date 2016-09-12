package bf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

func Run(program string, reader *bufio.Reader) (string, error) {
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
					// do nothing if we encounter EOF when reading input
					continue
				}
				return "", errors.New(fmt.Sprintf("error encountered when reading input: %v", err))
			}
			data[dataPointer] = b
		case '[':
			if data[dataPointer] == 0 {
				// search forward for corresponding ']'
				loopStart := i
				i++
				for count := 1; count != 0; i++ {
					if i >= len(program) {
						return "", errors.New(fmt.Sprintf("Unterminated loop caught beginning at idx: %v", loopStart))
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
						return "", errors.New(fmt.Sprintf("Encountered loop termination without opening bracket at idx: %v", loopEnd))
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
