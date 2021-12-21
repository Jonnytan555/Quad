package main

import (
	"Quad"
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Package contains the functions required to compare the command line output
// with the expected result stored in a file.

type testData struct {
	q        string
	x        int
	y        int
	expected string
}

var testCases = []testData{}

// All additional test cases will be read from this file
const fileName string = "test-cases.txt"

// TestQuad will test the output of the main.go program to match the desired result
func TestQuad(t *testing.T) {

	// add additional testcases from the file into the testCases
	AddTestCases(fileName)

	for _, test := range testCases {
		out := RunTest(test.q, test.x, test.y, test.expected)

		if out != test.expected {
			t.Errorf("For argument Quad%v (%v,%v) \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", test.q, test.x, test.y, test.expected, string(out))
			panic("err")
		} else {
			t.Logf("output for Quad%s (%v,%v) is \n\033[1;32m%s\033[0m", test.q, test.x, test.y, out)
			t.Log("\n")
			t.Logf("expected output for Quad%s (%v,%v) is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", test.q, test.x, test.y, test.expected)

		}
	}
}

// addTestCases will read the contents of the specified file, extract and add the desired output into the testCases
func AddTestCases(fileName string) {

	args := []string{}
	tempStr := ""

	lines, err := ReadLines(fileName)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		if line[0] == '$' {
			if tempStr != "" {
				x, _ := strconv.Atoi(args[1])
				y, _ := strconv.Atoi(args[2])
				testCases = append(testCases, testData{args[0], x, y, tempStr})
				tempStr = ""
			}
			args = strings.Split(string(line[1:len(line)-1]), ",")

		} else if line[0] == '£' {
			x, _ := strconv.Atoi(args[1])
			y, _ := strconv.Atoi(args[2])
			testCases = append(testCases, testData{args[0], x, y, ""})

		} else {
			tempStr += line + "\n"
		}
		// Last test case(No more testcases in the file)
		if i == len(lines)-1 {
			x, _ := strconv.Atoi(args[1])
			y, _ := strconv.Atoi(args[2])
			testCases = append(testCases, testData{args[0], x, y, tempStr})
		}
	}
}

func RunTest(question string, x, y int, exp string) string {
	original := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	switch question {
	case "A":
		Quad.QuadA(x, y)
	case "B":
		Quad.QuadB(x, y)
	case "C":
		Quad.QuadC(x, y)
	case "D":
		Quad.QuadD(x, y)
	case "E":
		Quad.QuadE(x, y)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = original

	return string(out)
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
