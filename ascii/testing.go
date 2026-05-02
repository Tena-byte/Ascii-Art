package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func testAscii() {

	input := os.Args[1]
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)
	}
}
