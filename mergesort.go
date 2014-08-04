package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("sorry fellers, need a file name")
		return
	}
	filename := os.Args[1]
	valuesToSort := extractSliceFromCsv(filename)
	sorted := MergeSort(valuesToSort)
	for _, v := range sorted {
		fmt.Print(v)
		fmt.Print(",")
	}
}

func MergeSort(valuesToSort []int) []int {
	//base case
	if len(valuesToSort) == 1 {
		return valuesToSort
	}
	//try with making a goroutine, and without
	left := valuesToSort[0 : len(valuesToSort)/2]
	right := valuesToSort[len(valuesToSort)/2 : len(valuesToSort)]
	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)
	err, merged := Merge(sortedLeft, sortedRight)
	if err != nil {
		panic(err)
	}
	return merged
}

func Merge(left []int, right []int) (error, []int) {
	result := []int{}
	i, j := 0, 0
	eltCount := len(left) + len(right)
	k := 0
	for k < eltCount {
		if left[i] < right[j] {
			result = append(result, left[i])
			k++
			i++
			if i >= len(left) {
				result = append(result, right[j:len(right)]...)
				return nil, result
			}
		} else {
			result = append(result, right[j])
			k++
			j++
			if j >= len(right) {
				result = append(result, left[i:len(left)]...)
				return nil, result
			}
		}
	}
	return errors.New("whoops,this merging is probably written wrong"), nil
}

func extractSliceFromCsv(filename string) []int {
	//open the file
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("problem opening file")
		panic(err)
	}

	defer func() {
		if err := inputFile.Close(); err != nil {
			fmt.Println("problem closing file")
			panic(err)
		}
	}()

	reader := bufio.NewReader(inputFile)
	extracted := []int{}
	token, readerr := reader.ReadString(',')
	counter := 0

	for readerr == nil {
		counter++
		extracted = ConvertAndAppend(token,extracted)
		token, readerr = reader.ReadString(',')
	}
	if readerr == io.EOF {
		extracted = ConvertAndAppend(token,extracted)

	}
	return extracted
}

func ConvertAndAppend(token string, appendee []int) []int{
	num, err := strconv.Atoi(TrimPunc(token))
		if err != nil {
			fmt.Println("problem converting")
			panic(err)
		}
		return append(appendee, num)
}

func TrimPunc(word string) string {
	return strings.Trim(word, ".,\t,\n")
}
