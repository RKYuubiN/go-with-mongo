package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rkyuubin/gowithmongo/router"
)

func main() {
	fmt.Println("Mongo with Go")
	router := router.Router()
	fmt.Println("Listening At Port 8086...")
	log.Fatal(http.ListenAndServe(":8086", router))

	// ? using log.fatal so that if http.ListenAndServe doesn't work we get logs???
}
