package router

import (
	"fmt"
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Initialize() {
	router := mux.NewRouter()

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid := vars["uuid"]
		user := models.GetUser(uuid)
		fmt.Println(user)
	}).Methods("GET")

	err := http.ListenAndServe(os.Getenv("API_PORT"), router)
	if err != nil {
		fmt.Println("Error serving API")
		panic(err)
	}

}
