package main

import (
	"bufio"
	"bytes"
	"testing"
)

var bfTests = []struct {
	program  string
	input    string
	expected string
}{
	{"", "", ""},
	{",", "", ""},
	{"asdfasdfasliy a479489t 2084t ;a;sodif jlasuh ", "", ""},
	{"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.", "", "Hello World!\n"},
	{",.", "!", "!"},
	{"[]", "", ""},
	{"+>,>,<<[>[>.<-]<-]", "!!", "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"},
	// rot13: https://copy.sh/brainfuck/prog/rot13.b
	{",[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>++++++++++++++<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>>+++++[<----->-]<<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>++++++++++++++<-[>+<-[>+<-[>+<-[>+<-[>+<-[>++++++++++++++<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>>+++++[<----->-]<<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>+<-[>++++++++++++++<-[>+<-]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]>.[-]<,]", "hello world 123", "uryyb jbeyq 123"},
}

func TestInterpreter(t *testing.T) {
	for _, tt := range bfTests {
		actual, err := Run(tt.program, bufio.NewReader(bytes.NewReader([]byte(tt.input))))
		if err != nil {
			t.Errorf("unexpected error: %v with program: %v", err, tt.program)
		}
		if actual != tt.expected {
			t.Errorf("Run(%s): expected %s, actual %s", tt.program, tt.expected, actual)
		}
	}
}

func TestUnterminatedLoop(t *testing.T) {
	_, err := Run("[.", bufio.NewReader(bytes.NewReader([]byte(""))))
	if err == nil {
		t.Errorf("expected error, instead err was nil")
	}

	expectedErr := "Unterminated loop caught beginning at idx: 0"
	if err.Error() != expectedErr {
		t.Errorf("expected err: %s, actual err: %s", expectedErr, err.Error())
	}
}

func TestPrematureLoopTermination(t *testing.T) {
	_, err := Run("+]", bufio.NewReader(bytes.NewReader([]byte(""))))
	if err == nil {
		t.Errorf("expected error, instead err was nil")
	}

	expectedErr := "Encountered loop termination without opening bracket at idx: 1"
	if err.Error() != expectedErr {
		t.Errorf("expected err: %s, actual err: %s", expectedErr, err.Error())
	}
}
