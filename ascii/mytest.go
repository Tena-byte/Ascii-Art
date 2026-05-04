package ascii

import (
	"fmt"
	"os"
	"strings"
)

func doItMyself() {

	inputFile := os.Args[1]
	//inputText := os.Args[2]
	data := getInput("banners/" + inputFile + ".txt")
	splitted := splitter(data)

	mapwork := mappingShii(splitted)
	fmt.Println(mapwork)

}

func getInput(inputFile string) string {

	data, _ := os.ReadFile(inputFile)
	return string(data)
}

func splitter(text string) []string {
	spltt := strings.Split(text, "\n")
	return spltt
}

func mappingShii(text []string) map[rune][]string {

	result := make(map[rune][]string)

	for i := 32; i < 126; i++ {

		ch := rune(i)
		startIndex := int(ch-32)*9 + 1

		if startIndex+8 < len(text) {
			result[ch] = text[startIndex : startIndex+8]
		}
	}

	return result
}


