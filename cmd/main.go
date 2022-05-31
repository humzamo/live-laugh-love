package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting server...")

	r := chi.NewRouter()
	r.Get("/api/getExample", handleGet)
	r.Post("/api/postExample", handlePost)
	http.ListenAndServe(":8080", r)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An example of a get request."))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Print("Message from post request: ", bodyString)

	message := fmt.Sprintf("You just sent: %s", bodyString)

	if bodyString == "" {
		message = "You didn't send anything!"
	}

	w.Write([]byte(fmt.Sprintf("An example of a post request.\n%s", message)))
}
