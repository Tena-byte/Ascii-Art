package ascii

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func multiAscii() {

	if len(os.Args) != 2 {
		fmt.Println("failed")
		return
	}

	data, _ := os.ReadFile("banners/standard.txt")
	rowsInLine := strings.Split(string(data), "\n")

	text := os.Args[1]

	rowHeight := 0

	for i := 1; i < len(rowsInLine); i++ {

		if rowsInLine[i] == "" {
			rowHeight = i - 1
			break
		}
	}

	text = strings.NewReplacer("\\n", "\n").Replace(text)
	lines := strings.Split(text, "\n")

	for _, line := range lines {

		if line == "" {
			fmt.Println()
			return
		}

		art := make([]string, rowHeight)
		for _, ch := range line {

			startIndex := int(ch-32)*(rowHeight+1) + 1

			for j := 0; j < rowHeight; j++ {
				art[j] += rowsInLine[startIndex+j]
			}
		}
		//time.Sleep(time.Nanosecond)
		fmt.Println(strings.Join(art, "\n"))
	}

}
