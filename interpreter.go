package main

import (
	"fmt"
	"strings"
)

var data [30000]byte

func main() {
	input := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	dataPointer := 0

	prevOpenLoopIdx := -1
	// prevCloseLoopIdx := -1
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '>':
			// fmt.Printf("data pointer was: %v, now: %v\n", dataPointer, dataPointer+1)
			dataPointer++
		case '<':
			// fmt.Printf("data pointer was: %v, now: %v\n", dataPointer, dataPointer-1)
			dataPointer--
		case '+':
			data[dataPointer]++
			// fmt.Printf("incrementing cell at %v, value now %v\n", dataPointer, data[dataPointer])
		case '-':
			data[dataPointer]--
			// fmt.Printf("decrementing cell at %v, value now %v\n", dataPointer, data[dataPointer])
		case '.':
			// fmt.Printf("printing cell at %v\n", dataPointer)
			fmt.Printf("%v", string(data[dataPointer]))
		case ',':
			// TODO accept single byte of input
		case '[':
			if data[dataPointer] != 0 {
				// fmt.Printf("entering loop, data pointer: %v\n", dataPointer)
				prevOpenLoopIdx = i
				continue
			} else {
				// set ip to instruction after next ']'
				// fmt.Printf("exiting loop, data pointer: %v\n", dataPointer)
				afterClose := strings.IndexByte(string(input[i:]), ']') + 1
				if afterClose == 0 {
					panic(fmt.Sprintf("Unterminated loop caught beginning at idx: %v", i))
				}
				prevOpenLoopIdx = -1
				i = afterClose - 1
			}
		case ']':
			if data[dataPointer] != 0 {
				// set ip to instruction after previous '['
				if prevOpenLoopIdx != -1 {
					// fmt.Printf("looping, data pointer: %v\n", dataPointer)
					i = prevOpenLoopIdx
				} else {
					panic(fmt.Sprintf("Encountered loop termination without opening bracket at idx: %v", i))
				}
			} else {
				// fmt.Printf("exiting loop, data pointer: %v\n", dataPointer)
				prevOpenLoopIdx = -1
				continue
			}
		}
	}
}
