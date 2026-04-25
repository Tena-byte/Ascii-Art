package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	asciiArt()
}

func asciiArt() {

	data, _ := os.ReadFile("banners/testing.txt")

	//fmt.Print(data)

	//rowInline := strings.Fields(string(data))
	//fmt.Println(len(rowInline))

	rowsInline := strings.Split(string(data), "")
	//fmt.Print(  len(rowsInline))

	for i, r := range rowsInline{

		fmt.Println(i, r)
	}

	// rowsHeight := 0

	// for i := 0; i < len(rowsInline); i++ {

	// 	if rowsInline[i] == "" {

	// 		rowsHeight   = i - 1

	// 	}
	// }

	//fmt.Println(rowsHeight)

}
