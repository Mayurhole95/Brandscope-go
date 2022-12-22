package csv_upload

import (
	"context"
	"fmt"

	"github.com/Mayurhole95/Brandscope-go/csv_validate"
	"github.com/Mayurhole95/Brandscope-go/db"
	"go.uber.org/zap"
)

func (cs *CsvService_upload) Upload(ctx context.Context, id string) (successmessage string, err error) {
	LogData, err = cs.store.FindLogID(ctx, id)
	if err == ErrEmptyData {
		fmt.Println("empty")
		return
	}
	ReturnError(err)
	exist, err := cs.store.FindID(ctx, LogData.BrandID, LogData.ReleaseID)
	csvData := csv_validate.CSVData
	fmt.Println(csvData)
	if !exist {
		return "Brand Doesn't exist", nil //Will change this later on
	}
	ReturnError(err)
	return "", nil
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &CsvService_upload{
		store: s,
	}
}

func ReturnError(err error) {
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}
}
