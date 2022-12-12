package csv

import (
	"net/http"

	"github.com/Mayurhole95/Brandscope-go/api"
	"github.com/gorilla/mux"
)

func ValidateCSV(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		resp, err := service.Validate(req.Context(), vars["id"])
		if err == errNoData {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

// func List(service Service) http.HandlerFunc {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		resp, err := service.ShowTables(req.Context())
// 		if err == errNoData {
// 			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
// 			return
// 		}
// 		if err != nil {
// 			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
// 			return
// 		}

// 		api.Success(rw, http.StatusOK, resp)
// 	})
// }
