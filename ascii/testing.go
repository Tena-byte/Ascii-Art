package ascii

import (
	"fmt"
	"os"
	"strings"
)

func testAscii() {

	inputfile := os.Args[1]
	//inputext := os.Args[2]

	input := Fileloder(inputfile)

	splited := spliter(input)
	mapped := Mapper(splited)
	//rendered := Renderer(inputext, mapped)

	fmt.Println(mapped)

}

func Fileloder(inputfile string) string {
	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func spliter(data string) []string {
	lines := strings.Split(data, "\n")

	return lines
}

func Mapper(lines []string) map[rune][]string {
	cMap := make(map[rune][]string)

	index := 0
	asciicode := 32

	for _, line := range lines {
		if index+8 < len(line) {
			break
		}
		block := lines[index : index+8]
		fmt.Println(block)
		cMap[rune(asciicode)] = block
		index += 9
		asciicode++
	}
	return cMap

}

// func Renderer(inputext byte, cMap map[rune][]string) string {
// 	words := string(inputext)
// 	word := strings.Split(words, "\\n")

// }

