package routes

import (
	"github.com/marcos-nsantos/loja-go/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Index)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
