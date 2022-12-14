package server

import (
	"net/http"

	"github.com/Mayurhole95/Brandscope-go/api"
	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//validate
	router.HandleFunc("/validate", CSV.validateCSV()).Methods(http.MethodPost)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
