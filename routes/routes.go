package routes

import (
	"log"
	"net/http"

	"github.com/alexsantossilva/go-api-rest/controllers"
	"github.com/alexsantossilva/go-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/v1/personas", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/api/v1/personas/{id}", controllers.GetPerson).Methods("Get")
	r.HandleFunc("/api/v1/personas", controllers.NewPerson).Methods("Post")
	r.HandleFunc("/api/v1/personas/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/v1/personas/{id}", controllers.UpdatePerson).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
