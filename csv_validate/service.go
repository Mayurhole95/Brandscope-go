package csv_validate

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Mayurhole95/Brandscope-go/db"
	csvtag "github.com/artonge/go-csv-tag/v2"
	"go.uber.org/zap"
)

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &CsvSerive_validate{
		store: s,
	}
}

func (cs *CsvSerive_validate) Validate(ctx context.Context, id string) (successmessage string, err error) {
	LogData, err = cs.store.FindLogID(ctx, id)
	if err == ErrEmptyData {
		fmt.Println("empty")
		return
	}
	var success Success
	// var successmessage string
	ReturnError(err)
	// fmt.Println(LogData)

	Exist, err := cs.store.FindID(ctx, LogData.BrandID, LogData.ReleaseID)
	ReturnError(err)

	if Exist {
		missingheader, err := HeaderCheck()
		if err != nil {
			csvFile, err := os.Create(file_name_errors)
			ReturnError(err)
			csvwriter := csv.NewWriter(csvFile)
			// for i := 0; i < len(missingheader); i++ {
			// 	_ = csvwriter.Write(missingheader[i])
			// }
			for _, j := range missingheader {
				_ = csvwriter.Write(j)
			}
			csvwriter.Flush()
			csvFile.Close()
			success = SuccessMessage(false, errHeadersMissing, file_name_errors)
			status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
			statusstring, err := json.Marshal(status)
			successmessage := string(statusstring)
			return successmessage, errNoData
		}
		success = SuccessMessage(true, errHeadersFound, "No error")
		csvDataMap, err := cs.store.ListData(LogData.BrandID)
		ReturnError(err)
		months, err := cs.store.ListMonths(LogData.ReleaseID)
		dbMonths, err = ChangeDateFormat(months)
		// fmt.Println(dbMonths)
		ReturnError(err)
		errorstring, err := readCSVData(csvDataMap)
		ReturnError(err)
		if errorstring == "" {
			success = SuccessMessage(true, perfectEntry, "")
			status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
			statusstring, err := json.Marshal(status)
			ReturnError(err)
			successmessage = string(statusstring)
		} else {
			success = SuccessMessage(false, errorstring, file_name_errors)
			success = SuccessMessage(true, perfectEntry, "")
			status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
			statusstring, err := json.Marshal(status)
			ReturnError(err)
			successmessage = string(statusstring)
		}
	} else {
		success = SuccessMessage(false, errBrandIDExists, file_name_errors)
		success = SuccessMessage(true, perfectEntry, "")
		status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
		statusstring, err := json.Marshal(status)
		ReturnError(err)
		successmessage = string(statusstring)
	}
	status := &Success{Success: success.Success, Message: success.Message, Filepath: success.Filepath}
	statusstring, err := json.Marshal(status)
	successmessage = string(statusstring)
	// fmt.Println(successmessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(statusstring))
	// fmt.Println("Success : ", success.Success)
	// fmt.Println("Message : ", success.Message)
	// fmt.Println("Filepath : ", success.Filepath)
	return successmessage, nil
}

func ReturnError(err error) {
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}
}

func SuccessMessage(status bool, message string, filepath string) (success Success) {

	success.Success = status
	success.Message = message
	success.Filepath = filepath

	// on:"Success"`
	// Message  string `json:"Message"`
	// Filepath string `json:"Filepath"`

	return success

}

func HeaderCheck() (missingheader [][]string, err error) {
	var map_headers = make(map[string]int)
	headers := [23]string{"CatalogueOrder", "BrandscopeCarryOver", "Integration_ID", "Barcode", "SKU", "ProductName", "ProductColourCode", "ProductDisplayColour", "GenericColour", "SizeBreak", "AttributeValue", "AttributeType", "AttributeSequence", "DisplayWholesaleRange", "DisplayWholesale", "DisplayRetail", "PackUnits", "AvailableMonths", "AgeGroup", "Gender", "State", "PreOrderLeadTimeDays", "PreOrderMessage"}
	missingheaders := make([]string, 0)
	missingheaders2d := make([][]string, 0)
	readCsvFile, err := os.Open(LogData.Original_file_location)
	ReturnError(err)
	defer readCsvFile.Close()

	csvReader := csv.NewReader(readCsvFile)
	records, err := csvReader.Read()

	ReturnError(err)
	lenarr := len(records)
	count := 0
	for column_headers := range Iterate(len(headers)) {
		header_present := "false"
		for column_csv := range Iterate(lenarr) {
			if headers[column_headers] == records[column_csv] {
				map_headers[headers[column_headers]] = column_csv
				header_present = "true"
				count = count + 1
				break
			}

		}

		if header_present == "false" {
			missingheaders = append(missingheaders, headers[column_headers])
			missingheaders2d = append(missingheaders2d, [][]string{missingheaders}...)
			missingheaders = missingheaders[1:]
		}

	}

	if count < len(headers) {
		err = errors.New(errHeadersMissing)
		return missingheaders2d, err
	}
	err = nil
	return missingheaders2d, nil
}

func readCSVData(dbData map[string]db.Verify) (errorMessage string, err error) {
	// var CsvData = []BrandHeader{}
	err = csvtag.LoadFromPath(LogData.Original_file_location,
		&CSVData,
		csvtag.CsvOptions{ // Load your csv with optional options
			Separator: ',', // changes the values separator, default to ','
		})
	if err != nil {
		return "", err
	}
	var file [][]string
	readCsvFile, err := os.Open(LogData.Original_file_location)
	ReturnError(err)

	// remember to close the file at the end of the program
	defer readCsvFile.Close()
	csvReader := csv.NewReader(readCsvFile)
	csvReader.FieldsPerRecord = -1
	file, err = csvReader.ReadAll()
	ReturnError(err)

	errorMessage = ""
	str := ""
	file[0] = append(file[0], "status")
	count := 0
	verifiedFields := make(map[string]bool)

	for i, j := range CSVData {
		count = 0
		str = ""
		if _, find := verifiedFields[j.Integration_ID]; find {
			str = errCarryOverNot
			count = 1
		}
		if count == 1 {
			continue
		}
		if _, ok := dbData[j.Integration_ID]; ok {
			if j.BrandscopeCarryOver == "N" || j.BrandscopeCarryOver == "n" {
				str = errCarryOverNot
				errorMessage += str
				file[i] = append(file[i], str)
				count = 1
			} else {
				if dbData[j.Integration_ID].SKU != j.SKU || dbData[j.Integration_ID].Colour_code != j.ProductColourCode || dbData[j.Integration_ID].Size != j.SizeBreak {
					str = errCarryOverYes
					errorMessage += str
					file[i] = append(file[i], str)
					count = 1
				}

			}
		} else {
			if j.BrandscopeCarryOver == "Y" || j.BrandscopeCarryOver == "y" {
				str = errCarryOverYes
				errorMessage += str
				file[i] = append(file[i], str)
				count = 1
			}
		}
		if count == 1 {
			continue
		}

		str, err = CheckValidations(j, i)

		if str == perfectEntry {
			verifiedFields[j.Integration_ID] = true
		}
		file[i] = append(file[i], str)
		if str != "" {
			// errorMessage += strconv.Itoa(i)
		}
		errorMessage += str
	}

	csvFile, err := os.Create(file_name_errors)
	ReturnError(err)
	csvwriter := csv.NewWriter(csvFile)

	for i := range Iterate(len(file)) {

		_ = csvwriter.Write(file[i])
	}
	csvwriter.Flush()
	csvFile.Close()
	return errorMessage, nil
}
func Iterate(n int) []struct{} {
	return make([]struct{}, n)
}
