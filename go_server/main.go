package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	Age        uint16
	Money      int16
	Avg_grades float64
	Happiness  float64
	Hobbies    []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is : %s. He is  %d and he has money equel: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"football", "Scate", "Dance"}}
	// bob.setNewName("alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "privet loshari")
}

func main() {

	http.HandleFunc("/", home_page)

	http.HandleFunc("/contacts/", contacts_page)

	fmt.Println("Server listening at port 80")

	http.ListenAndServe(":80", nil)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to my website!")
	// })

}
