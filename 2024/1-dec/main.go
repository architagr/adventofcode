package main

import (
	readinput "aoc_1dec_2024/read_input"
	"flag"
	"fmt"
	"sort"
)

var (
	inputFileName = flag.String("input-file", "input.txt", "this will be used as input")
)

func main() {
	flag.Parse()
	inputData1, inputData2 := readinput.ReadFile(*inputFileName)
	sort.Ints(inputData1)
	sort.Ints(inputData2)
	sum := int64(0)

	for i := 0; i < len(inputData1); i++ {
		sum += int64(absVal(inputData1[i] - inputData2[i]))
	}
	fmt.Println("distance-", sum)
	frequencyMap := getFrequencyMap(inputData2)
	similarityScore := int64(0)
	for i := 0; i < len(inputData1); i++ {
		similarityScore += int64(inputData1[i]) * int64(frequencyMap[inputData1[i]])
	}
	fmt.Println("similarityScore-", similarityScore)
}
func getFrequencyMap(right []int) map[int]int {
	m := make(map[int]int)
	for _, v := range right {
		m[v]++
	}
	return m
}
func absVal(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
