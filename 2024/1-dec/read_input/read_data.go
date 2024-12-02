package readinput

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(inputFileName string) ([]int, []int) {
	// read input.txt
	fi, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	input1 := make([]int, 0)
	input2 := make([]int, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		a := ReadArrInt(scanner)
		input1 = append(input1, a[0])
		input2 = append(input2, a[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input1, input2
}

func ReadString(in *bufio.Scanner) string {
	nStr := in.Text()
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	return nStr
}

func ReadStrArr(in *bufio.Scanner) []string {
	line := ReadString(in)
	numbs := strings.Split(line, "   ")
	return numbs
}

func ReadArrInt(in *bufio.Scanner) []int {
	numbs := ReadStrArr(in)
	arr := make([]int, 0)
	for _, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr = append(arr, val)
	}
	return arr
}
