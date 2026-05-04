package ascii

import (
	"fmt"
	"os"
	"strings"
)

func doItMyself() {

	inputFile := os.Args[1]
	inputText := strings.ReplaceAll(os.Args[2], "\\n", "\n")
	data := getInput("banners/" + inputFile + ".txt")
	splitted := splitter(data)
	mapping := mappingShii(splitted)

	result := renderer(inputText, mapping)
	fmt.Println(result)
	
}




func getInput(text string) string {
	data, _ := os.ReadFile(text)
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


func renderer(text string, font map[rune][]string) string {

    for row := 0; row < 8; row++ {
        for _, ch := range text {
            if art, ok := font[ch]; ok {
                fmt.Print(art[row])
            }
        }
        fmt.Println()
    }
	return ""
}

