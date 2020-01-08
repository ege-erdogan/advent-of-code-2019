package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		string := scanner.Text()
		code := mapToInt(strings.Split(string, ","))
		noun, verb := findNounVerb(code, 19690720)
		fmt.Println("Result: ", 100*noun+verb)
	}
}

func findNounVerb(code []int, target int) (noun int, verb int) {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			temp := make([]int, len(code))
			copy(temp, code)
			temp[1] = n
			temp[2] = v
			runProgram(temp)
			if temp[0] == target {
				return n, v
			}
		}
	}
	return 0, 0
}

// runs the given IntCode program
// returns the final state of the memory
func runProgram(code []int) {
	opIndex := 0
	nextOp := code[0]
	for nextOp != 99 && opIndex < len(code) {
		op := getOp(nextOp)
		loc1 := code[opIndex+1]
		loc2 := code[opIndex+2]
		resultLoc := code[opIndex+3]
		result := op(code[loc1], code[loc2])
		code[resultLoc] = result
		opIndex += 4
		nextOp = code[opIndex]
	}
}

// this is somewhat overkill, wanted to see how this is done in go.
func getOp(opcode int) func(int, int) int {
	if opcode == 1 {
		return func(o1 int, o2 int) int {
			return o1 + o2
		}
	} else if opcode == 2 {
		return func(o1 int, o2 int) int {
			return o1 * o2
		}
	}
	return nil
}

func mapToInt(arr []string) []int {
	var result = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i], _ = strconv.Atoi(arr[i])
	}
	return result
}
