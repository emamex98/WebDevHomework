// forms.go
package main

import (
    "html/template"
	"net/http"
	"fmt"
)

func main() {

	tmpl := template.Must(template.ParseFiles("form.html"))


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        var name string =   r.FormValue("name")
        var email string = r.FormValue("email")
        var message string = r.FormValue("message")

		fmt.Println("*** ¡Nuevo Mensaje! ***")
		fmt.Println("Nombre: ", name)
		fmt.Println("Correo: ", email)
		fmt.Println("Mensaje: ", message)
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}