package csv_validate

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Mayurhole95/Brandscope-go/db"
	"github.com/Mayurhole95/Brandscope-go/utils"
	csvtag "github.com/artonge/go-csv-tag/v2"
)

func ValidateCSVData(dbData map[string]db.Verify) (errorMessage string, err error) {

	err = csvtag.LoadFromPath(logdata.Original_file_location,
		&csvData,
		csvtag.CsvOptions{ // Load your csv with optional options
			Separator: ',', // changes the values separator, default to ','
		})
	utils.ReturnError(err)

	errorMessage = ""
	verifiedFields := make(map[string]bool)
	records := make([][]string, 0)
	headers := readHeaders()
	headers = append(headers, "status")
	records = append(records, headers)
	for i, c := range csvData {
		record := c.ToArray()
		if _, ok := verifiedFields[c.Integration_ID]; !ok {
			continue
		}
		_, ok := dbData[c.Integration_ID]
		if ok && (c.BrandscopeCarryOver == "N" || c.BrandscopeCarryOver == "n") {
			errorMessage += errCarryOverNot
			record = append(record, errCarryOverNot)
			records = append(records, record)
			continue
		} else if ok && (dbData[c.Integration_ID].SKU != c.SKU || dbData[c.Integration_ID].Colour_code != c.ProductColourCode || dbData[c.Integration_ID].Size != c.SizeBreak) {
			errorMessage += errCarryOverYes
			record = append(record, errCarryOverYes)
			records = append(records, record)
			continue
		} else if !ok && (c.BrandscopeCarryOver == "Y" || c.BrandscopeCarryOver == "y") {
			errorMessage += errCarryOverYes
			record = append(record, errCarryOverYes)
			records = append(records, record)
			continue
		}
		resp, _ := CheckValidations(c, i)

		if resp == perfectEntry {
			verifiedFields[c.Integration_ID] = true
			fmt.Println(verifiedFields)

		}
		record = append(record, resp)
		if resp != "" {
			errorMessage += strconv.Itoa(i)
		}
		errorMessage += resp
		records = append(records, record)
	}

	csvFile, err := os.Create(file_name_errors)
	utils.ReturnError(err)

	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.WriteAll(records)
	utils.ReturnError(err)

	csvWriter.Flush()
	csvFile.Close()
	return errorMessage, nil
}
