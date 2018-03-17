package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

var addr = flag.String("localhost", ":8000", "hostname of server")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloPage)
	mux.HandleFunc("/", formPage)
	http.ListenAndServe(*addr, mux)
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home").ParseFiles("./tmpl/home.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, map[string]interface{}{
		"name":    "Gopher",
		"version": "1.10",
		"value":   "",
	})

	if err != nil {
		fmt.Println(err)
	}
}

func sayHello(name string, w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home").ParseFiles("./tmpl/home.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		fmt.Println(err)
	}

}

func formPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		tmpl, _ := template.New("form").ParseFiles("./tmpl/form.html", "./tmpl/header.html")
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		name := r.FormValue("username")
		sayHello(name, w, r)
	}

}
