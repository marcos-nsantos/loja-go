package routes

import (
	"github.com/marcos-nsantos/loja-go/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
