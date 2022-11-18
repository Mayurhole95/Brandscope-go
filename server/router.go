package server

import (
	"net/http"

	"github.com/Mayurhole95/Brandscope-go/api"
	csv "github.com/Mayurhole95/Brandscope-go/csv_validate"
	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//validate
	router.HandleFunc("/validate", csv.ValidateCSV(dep.csvService)).Methods(http.MethodGet)
	// router.HandleFunc("/list", csv.List(dep.csvService)).Methods(http.MethodGet)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
