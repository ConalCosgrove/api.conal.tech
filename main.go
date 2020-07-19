package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Articles []Article

type Article struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Path  string `json:"path"`
	Slug  string `json:"slug"`
}

func articles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest(port string) {
	http.HandleFunc("/articles", articles)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT not specified")
	}
	Articles = []Article{
		Article{Title: "Javascript", Img: "https://upload.wikimedia.org/wikipedia/commons/9/99/Unofficial_JavaScript_logo_2.svg", Path: "/articles/js", Slug: "Everything Javascript, front and backend."},
		Article{Title: "Go", Img: "https://camo.githubusercontent.com/98ed65187a84ecf897273d9fa18118ce45845057/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67", Path: "/articles/golang", Slug: "Stuff about the Go programming language."},
		Article{Title: "Rust", Img: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/Rust_programming_language_black_logo.svg/768px-Rust_programming_language_black_logo.svg.png", Path: "/articles/rust", Slug: "My adventures in Rust."}}
	handleRequest(port)
}
