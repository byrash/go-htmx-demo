package main

import (
	"html/template"
	"log"
	"net/http"
)

type TodoItem struct {
	Item  string
	Store string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	todoItems := map[string][]TodoItem{
		"TodoItems": {
			{Item: "Apples", Store: "Fruit Market"},
			{Item: "Oranges", Store: "Fruit Market"},
			{Item: "Chicken", Store: "Prime Cuts"},
		},
	}

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, todoItems)
	}

	addTodoItemHandler := func(w http.ResponseWriter, r *http.Request) {
		item := r.PostFormValue("item")
		store := r.PostFormValue("store")
		tmpl.ExecuteTemplate(w, "todoItems", TodoItem{Item: item, Store: store})
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/add-todo-item", addTodoItemHandler)
	log.Println("Server listening at 8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
