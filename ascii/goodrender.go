package ascii

import (
	"fmt"
	"os"
	"strings"
)

func goodRender() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: [banner file] [text]")
		return
	}

	bannerPath := "banners/" + os.Args[1] + ".txt"
	inputext := os.Args[2]

	data, _ := os.ReadFile(bannerPath)
	rows := strings.Split(string(data), "\n")

	lines := strings.Split(inputext, "\\n")
	
	for _, line := range lines {

		
		if line == "" {
			fmt.Println()
			continue
		}
		cMap := make(map[rune][]string)

		for _, ch := range line {

			startIndex := int(ch-32)*9 + 1
			fmt.Println(startIndex)

			for i := 0; i < 8; i++ {

				cMap[ch] = rows[startIndex : startIndex+8]
				fmt.Print(cMap[ch][i])
			}

		}
		fmt.Println()

	}
}
