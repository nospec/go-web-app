package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// create a struct that holds information to be displayed in our html file
type Welcome struct {
	Name string
	Time string
}

// go app entrypoint
func main() {
	welcome := Welcome{"Annonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening")
	//fmt.Println(http.ListenAndServe(":80", nil))
	fmt.Println(http.ListenAndServeTLS(":443", "cert.pem", "server.key", nil))
}
