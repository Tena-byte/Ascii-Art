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

		
		for i := 0; i < 8; i++ {

			for _, ch := range line {

				startIndex := int(ch-32)*9 + 1
				
				cMap[ch] = rows[startIndex : startIndex+8]

				fmt.Print(cMap[ch][i])
			}
			fmt.Println()
		}

		

	}
}


// package ascii

// import (
//     "fmt"
//     "os"
//     "strings"
// )

// func GoodRender() {
//     if len(os.Args) < 3 {
//         fmt.Println("Usage: [banner file] [text]")
//         return
//     }

//     bannerPath := "banners/" + os.Args[1] + ".txt"
//     inputext := os.Args[2]

   
//     data, _ := os.ReadFile(bannerPath)
//     rows := strings.Split(string(data), "\n")

//     cMap := make(map[rune][]string)

//     for i := 32; i <= 126; i++ {

//         startIndex := (i-32)*9 + 1

//         if startIndex+8 <= len(rows) {

//             cMap[rune(i)] = rows[startIndex : startIndex+8]
//         }
//     }

   
//     lines := strings.Split(inputext, "\\n")
//     for _, line := range lines {
//         if line == "" {
//             fmt.Println()
//             continue
//         }

//         // 4. THE HORIZONTAL FIX:
//         // We must print row-by-row (8 rows total)
//         for i := 0; i < 8; i++ {
//             for _, ch := range line {

//                 if block, ok := cMap[ch]; ok {
//                     // Print the i-th line of the current character
//                     fmt.Print(block[i])
//                 }
//             }
//             // After printing the i-th row for ALL characters, move to the next row
//             fmt.Println()
//         }
//     }
// }