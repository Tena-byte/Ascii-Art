package ascii

import (
	"fmt"
	"os"
	"strings"
)

func testAscii() {

	inputfile := os.Args[1]
	inputext := os.Args[2]

	input := Fileloder(inputfile)

	splited := spliter(input)
	mapped := Mapper(splited)
	rendered := Renderer(inputext, mapped)

	fmt.Println(rendered)

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

	for index < len(lines) {
		if index+8 > len(lines) {
			break
		}
		block := lines[index : index+8]
		cMap[rune(asciicode)] = block
		index += 9
		asciicode++
	}
	return cMap

}

func Renderer(inputext string, cMap map[rune][]string) string {
	words := strings.Split(inputext, "\\n")

	for _, line := range words{
		for i := range 8{
			if line == " " {
				fmt.Println("Empty words not allowed! Enter a word")
				continue
			}
			for _, char := range line {
				if block, ok := cMap[rune(char)]; ok {
					fmt.Print(block[i])
				}else{
					fmt.Print(cMap[' '][i])
				}
			}
			fmt.Println()
		}
	}
	return ""

}
