package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fmt.Println("Starting server...")

	router, err := Routes()
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", router)
}
