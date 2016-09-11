package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 30000 bytes according to brainfuck "spec"
var data [30000]byte

func main() {
	// input := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	input := "+[+[-]]"
	dataPointer := 0

	reader := bufio.NewReader(os.Stdin)

	prevOpenLoopIdx := -1
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '>':
			dataPointer++
		case '<':
			dataPointer--
		case '+':
			data[dataPointer]++
		case '-':
			data[dataPointer]--
		case '.':
			fmt.Printf("%v", string(data[dataPointer]))
		case ',':
			b, err := reader.ReadByte()
			if err != nil {
				fmt.Println(err)
				return
			}
			data[dataPointer] = b
		case '[':
			if data[dataPointer] != 0 {
				prevOpenLoopIdx = i
				continue
			} else {
				// set ip to instruction after next ']'
				afterClose := strings.IndexByte(string(input[i:]), ']')
				if afterClose == -1 {
					panic(fmt.Sprintf("Unterminated loop caught beginning at idx: %v", i))
				}
				prevOpenLoopIdx = -1
				i = afterClose
			}
		case ']':
			if data[dataPointer] != 0 {
				// set ip to instruction after previous '['
				if prevOpenLoopIdx != -1 {
					i = prevOpenLoopIdx
				} else {
					panic(fmt.Sprintf("Encountered loop termination without opening bracket at idx: %v", i))
				}
			} else {
				prevOpenLoopIdx = strings.IndexByte(string(input[:i]), '[')
				continue
			}
		}
	}
}
