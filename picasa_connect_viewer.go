package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

const (
	DEFAULT_PORT = "9000"
)

type Model struct {
	Title    string
	UserId   string
	PageSize int64
}

func handler(w http.ResponseWriter, req *http.Request) {
	var title string
	var userId string
	var pageSize int64

	if title = os.Getenv("TITLE"); len(title) == 0 {
		title = "Picasa Connect Viewer"
	}
	if userId = os.Getenv("USER_ID"); len(userId) == 0 {
		userId = "100851576803920751047"
	}
	if pageSize, _ = strconv.ParseInt(os.Getenv("PAGE_SIZE"), 10, 32); pageSize == 0 {
		pageSize = 32
	}

	model := Model{title, userId, pageSize}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, model)
	if err != nil {
		panic(err)
	}
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: ", err)
	}
}
