package ascii

import (
	"fmt"
	"os"
	"strings"
)

func asciiArt() {

	ch := 'T'
	var art []string

	data, _ := os.ReadFile("banners/standard.txt")

	rowsInline := strings.Split(string(data), "\n")

	rowsHeight := 0

	for i := 1; i < len(rowsInline); i++ {
		if rowsInline[i] == "" {
			rowsHeight = i - 1
			break
		}
	}

	startIndex := int(ch-32)*(rowsHeight+1) + 1

	for j := 0; j < rowsHeight; j++ {

		art = append(art, rowsInline[startIndex+j])

	}

	fmt.Println(strings.Join(art, "\n"))
}
