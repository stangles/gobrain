package bf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func Run(program string, reader *bufio.Reader) (string, error) {
	var output bytes.Buffer

	// 30000 bytes according to brainfuck "spec"
	var data [30000]byte

	dataPointer := 0
	prevOpenLoopIdx := -1
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
				return "", errors.New("error encountered when reading input")
			}
			data[dataPointer] = b
		case '[':
			if data[dataPointer] != 0 {
				prevOpenLoopIdx = i
				continue
			} else {
				// set ip to instruction after next ']'
				afterClose := strings.IndexByte(string(program[i:]), ']')
				if afterClose == -1 {
					return "", errors.New(fmt.Sprintf("Unterminated loop caught beginning at idx: %v", i))
				}
				prevOpenLoopIdx = -1
				i = afterClose
			}
		case ']':
			if prevOpenLoopIdx == -1 {
				return "", errors.New(fmt.Sprintf("Encountered loop termination without opening bracket at idx: %v", i))
			}

			if data[dataPointer] != 0 {
				// set ip to instruction after previous '['
				i = prevOpenLoopIdx
			} else {
				prevOpenLoopIdx = strings.LastIndexByte(string(program[:i]), '[')
				continue
			}
		}
	}
	return output.String(), nil
}
