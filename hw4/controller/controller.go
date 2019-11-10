package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func UnknownHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "Unknown router")
}

func PostFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	passwd := r.FormValue("passwd")
	t := template.Must(template.ParseFiles("template/form.html"))
	m := map[string]string{
		"name":   name,
		"passwd": passwd,
	}
	t.Execute(w, m)
}

func TestJs(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(&struct {
		Info string
	}{
		Info: "hello world",
	})
	fmt.Fprintf(w, string(data))
}
