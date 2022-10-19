package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ReturnPersonality).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonality).Methods("Put")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":666", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
