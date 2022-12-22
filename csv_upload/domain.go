package csv_upload

import (
	"context"

	"github.com/Mayurhole95/Brandscope-go/db"
)

type Service interface {
	Upload(ctx context.Context, id string) (successmessage string, err error)
}

type CsvService_upload struct {
	store db.Storer
}

var LogData db.LogID

type Success struct {
	Success  bool   `json:"Success"`
	Message  string `json:"Message"`
	Filepath string `json:"Filepath"`
}

type LogID struct {
	Original_file_location string `db:"original_file_location"`
	ReleaseID              string `db:"release_id"`
	BrandID                string `db:"brand_id"`
}
