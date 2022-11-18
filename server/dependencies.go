package server

import (
	"github.com/Mayurhole95/Brandscope-go/app"
	csv "github.com/Mayurhole95/Brandscope-go/csv_validate"
	"github.com/Mayurhole95/Brandscope-go/db"
)

type dependencies struct {
	csvService csv.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)
	CsvService := csv.NewService(dbStore, logger)
	return dependencies{
		csvService: CsvService,
	}, nil
}
