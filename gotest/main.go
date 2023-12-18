package main

import "fmt"

// "strings"
// "time"
func main() {
	total := sum(10)
	// fmt.Println(sum(4))

	fmt.Println(total)

	// const col = 30
	// Clear the screen by printing \x0c.
	// bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	// for i := 0; i < col; i++ {
	// fmt.Printf(bar, strings.Repeat("=", i)+">")
	// time.Sleep(500 * time.Millisecond)
	// }
	// fmt.Printf(bar+" Done!", strings.Repeat("=", col))

}

func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)

}
