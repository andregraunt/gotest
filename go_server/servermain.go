// package main
//
// import (
// 	"fmt"
// 	"net/http"
// )
//
// func main() {
// 	http.FileServer()
//
// 	fmt.Println("Server listening at port 80")
// 	http.ListenAndServe(":80", nil)
// }

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	go http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("Server listening at port 80")
	http.ListenAndServe(":80", nil)
	os.Exit(0)
}
