package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		
		log.Fatalln("Error listening and serving!")
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am root"))
	log.Println("log for root")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am hello"))
	log.Println("log for hello")
}
