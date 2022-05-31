package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleGetExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An example of a get request."))
}

func handlePostExample(w http.ResponseWriter, r *http.Request) {
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
