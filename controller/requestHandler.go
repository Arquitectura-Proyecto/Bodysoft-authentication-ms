package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authentication-MS: State Up")
}

// HandleRequest ..
func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", apiStatus)
	myRouter.HandleFunc("/create-user", createUserController).Methods("POST")
	myRouter.HandleFunc("/get-types", getTypesController).Methods("GET")
	myRouter.HandleFunc("/recover-password/{email}", recoverPasswordWithEmail).Methods("GET")
	myRouter.HandleFunc("/authentication/{email}/{password}", authenticationController).Methods("GET")
	myRouter.HandleFunc("/change-password", chagePasswordController).Methods("PUT")

	fmt.Println("Port 4000 is listening")
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}
