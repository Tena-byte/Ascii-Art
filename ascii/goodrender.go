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


	const (
		start = 32
		stop = 126
	)

	lines := strings.Split(inputext, "\\n")
	for _, line := range lines {

		if line == "" {
			fmt.Println()
			continue
		}


		makeMap := make(map[rune][]string)
		
		for i := 0; i < 8; i++ {

			for _, ch := range line {

				if ch < start || ch > stop {
					fmt.Println()
					return
				}
				startIndex := int(ch-32)*9 + 1

				makeMap[ch] = rows[startIndex : startIndex+8]
				fmt.Print(makeMap[ch][i])
			}
			fmt.Println()
		}
	}
}
