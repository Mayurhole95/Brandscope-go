package csv_validate

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Mayurhole95/Brandscope-go/db"
	"github.com/Mayurhole95/Brandscope-go/utils"
	"go.uber.org/zap"
)

type CsvService struct {
	store db.Storer
}

func (cs *CsvService) Validate(ctx context.Context, id string) (successMessage string, err error) {
	logdata, err = cs.store.FindLogID(ctx, id)
	var success utils.Success
	utils.ReturnError(err)

	exist, err := cs.store.FindID(ctx, logdata.BrandID, logdata.ReleaseID)
	utils.ReturnError(err)
	if !exist {
		success = utils.SuccessMessage(false, errBrandIDExists, file_name_errors)
		successMessage = utils.Marshal(success)
		fmt.Println(successMessage)
		return successMessage, nil
	}

	missingHeader, err := ValidateHeader()
	if err != nil {
		csvFile, err := os.Create(file_name_errors)
		utils.ReturnError(err)
		csvWriter := csv.NewWriter(csvFile)
		_ = csvWriter.WriteAll(missingHeader)
		csvWriter.Flush()
		csvFile.Close()
		success = utils.SuccessMessage(false, errHeadersMissing, file_name_errors)
		successMessage := utils.Marshal(success)
		return successMessage, errNoData
	}
	csvDataMap, err := cs.store.ListData(logdata.BrandID)
	utils.ReturnError(err)
	months, err := cs.store.ListMonths(logdata.ReleaseID)
	utils.ReturnError(err)
	dbMonths, err = ChangeDateFormat(months)
	utils.ReturnError(err)
	errorstring, err := ValidateCSVData(csvDataMap)
	utils.ReturnError(err)
	if errorstring == "" {
		success = utils.SuccessMessage(true, perfectEntry, "")

	} else {
		success = utils.SuccessMessage(false, errorstring, file_name_errors)
	}
	successMessage = utils.Marshal(success)
	fmt.Println(successMessage)

	return successMessage, nil

}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &CsvService{
		store: s,
	}
}
