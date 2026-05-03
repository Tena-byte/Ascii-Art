package ascii

import (
	"fmt"
	"os"
	"strings"
)

func multiAscii() {

	text := os.Args[1]
	data, _ := os.ReadFile("banners/standard.txt")
	rowsInLine := strings.Split(string(data), "\n")

	rowHeight := 0
	for i := 1; i < len(rowsInLine); i++ {
		if rowsInLine[i] == "" {
			rowHeight = i - 1
			break
		}
	}
	
	lines := strings.Split(text, "\\n")

	const (
		start = 32
		stop  = 127
	)

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			return
		}

		art := make([]string, rowHeight)
		for _, ch := range line {

			if ch < start || ch > stop {
				fmt.Printf("Invalid Ascii value found : %q \n", ch)
				return 
			}

			startIndex := int(ch-32)*(rowHeight+1) + 1

			for j := 0; j < rowHeight; j++ {
				art[j] += rowsInLine[startIndex+j]
			}
		}
		//time.Sleep(time.Nanosecond)
		fmt.Println(strings.Join(art, "\n"))
	}

}
