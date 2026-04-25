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

	ch := '&'
	var art []string

	data, _ := os.ReadFile("banners/standard.txt")

	rows := strings.Split(string(data), "\n")

	
	rowHeight := 0

	for i := 1; i < len(rows); i++{
		if rows[i] == ""{
			rowHeight = i - 1
			break
		}
	}

	start := int(ch - 32) * (rowHeight + 1) + 1

	for j := 0; j < rowHeight; j++{

		startIndex := start + j
		art = append(art, rows[startIndex])
	}

	fmt.Println(strings.Join(art, "\n"))
}