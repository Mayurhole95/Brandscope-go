package csv

import (
	"net/http"

	"github.com/Mayurhole95/Brandscope-go/api"
)

type empData struct {
	name string
	age  string
	city string
}

func validateCSV(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.validate(req.Context())
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
