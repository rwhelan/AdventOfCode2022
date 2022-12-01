package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

func ReadInputFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return content
}

func sumElfCals(rawCals []byte) int {
	resp := 0
	for _, v := range bytes.Split(rawCals, []byte{'\n'}) {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Printf("%+v\n", v)
			panic(err)
		}

		resp += i
	}

	return resp
}

func GroupElfCals(input []byte) []int {
	groups := bytes.Split(input, []byte{'\n', '\n'})
	resp := make([]int, len(groups))

	for i, group := range groups {
		resp[i] = sumElfCals(group)
	}

	return resp
}

func main() {
	elfs := GroupElfCals(ReadInputFile("input"))
	sort.Ints(elfs)

	for i := 0; i < len(elfs)/2; i++ {
		elfs[i], elfs[len(elfs)-1-i] = elfs[len(elfs)-1-i], elfs[i]
	}

	fmt.Println("Problem One: ", elfs[0])
	fmt.Println("Problem Two: ", elfs[0]+elfs[1]+elfs[2])

}
