package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Mayurhole95/Brandscope-go/db"
	csvtag "github.com/artonge/go-csv-tag/v2"
	"go.uber.org/zap"
)

type Service interface {
	Validate(ctx context.Context, id string) (success Success, err error)
}

var csvData = []BrandHeader{}

type CsvService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

var file_name string = "pride_priderelease_20221215114528.csv"
var file_name_errors string = "pride_priderelease_20221215114528_errors.csv"

func (cs *CsvService) Validate(ctx context.Context, id string) (success Success, err error) {
	brand_id := "380"
	release_id := "206"

	exist, err := cs.store.FindID(ctx, brand_id, release_id)
	ReturnError(err)

	if exist {
		missingheader, err := HeaderCheck()
		if err != nil {
			csvFile, err := os.Create(file_name)
			ReturnError(err)
			csvwriter := csv.NewWriter(csvFile)
			for i := 0; i < len(missingheader); i++ {
				_ = csvwriter.Write(missingheader[i])
			}
			csvwriter.Flush()
			csvFile.Close()
			success = SuccessMessage(false, errHeadersMissing, file_name_errors)

			return success, errNoData
		}
		success = SuccessMessage(true, errHeadersFound, "")
		csvDataMap, err := cs.store.ListData(brand_id)
		ReturnError(err)
		months, err := cs.store.ListMonths(release_id)
		dbMonths, err := ChangeDateFormat(months)
		fmt.Println(dbMonths)
		ReturnError(err)
		errorstring, err := readCSVData(csvDataMap)
		ReturnError(err)
		if errorstring == "" {
			success = SuccessMessage(true, perfectEntry, "")
		} else {
			success = SuccessMessage(false, errorstring, "")
		}
	} else {
		success = SuccessMessage(false, errBrandIDExists, "")
	}
	return success, nil
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
	return success

}

func HeaderCheck() (missingheader [][]string, err error) {
	var map_headers = make(map[string]int)
	headers := [23]string{"CatalogueOrder", "BrandscopeCarryOver", "Integration_ID", "Barcode", "SKU", "ProductName", "ProductColourCode", "ProductDisplayColour", "GenericColour", "SizeBreak", "AttributeValue", "AttributeType", "AttributeSequence", "DisplayWholesaleRange", "DisplayWholesale", "DisplayRetail", "PackUnits", "AvailableMonths", "AgeGroup", "Gender", "State", "PreOrderLeadTimeDays", "PreOrderMessage"}
	missingheaders := make([]string, 0)
	missingheaders2d := make([][]string, 0)
	readCsvFile, err := os.Open(file_name)
	ReturnError(err)
	defer readCsvFile.Close()

	csvReader := csv.NewReader(readCsvFile)
	records, err := csvReader.Read()

	ReturnError(err)
	lenarr := len(records)
	count := 0
	for column_headers := 0; column_headers < len(headers); column_headers++ {
		header_present := "false"
		for column_csv := 0; column_csv < lenarr; column_csv++ {
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
	err = csvtag.LoadFromPath(file_name,
		&csvData,
		csvtag.CsvOptions{ // Load your csv with optional options
			Separator: ',', // changes the values separator, default to ','
		})
	if err != nil {
		return "", err
	}

	var file [][]string
	readCsvFile, err := os.Open(file_name)
	ReturnError(err)

	// remember to close the file at the end of the program
	defer readCsvFile.Close()
	csvReader := csv.NewReader(readCsvFile)
	// csvReader.FieldsPerRecord = -1
	file, err = csvReader.ReadAll()
	ReturnError(err)

	errorMessage = ""
	str := ""
	file[0] = append(file[0], "status")
	count := 0
	verifiedFields := make(map[string]bool)
	for i := 0; i < len(csvData); i++ {
		count = 0
		str = ""

		print(verifiedFields)

		if _, find := verifiedFields[csvData[i].Integration_ID]; find {
			// fmt.Println("Found")
			str = errCarryOverNot
			count = 1
			// fmt.Println("same")

		}
		if count == 1 {
			continue
		}
		if _, ok := dbData[csvData[i].Integration_ID]; ok {
			if csvData[i].BrandscopeCarryOver == "N" || csvData[i].BrandscopeCarryOver == "n" {
				// fmt.Println(i, "Present")
				str = errCarryOverNot
				errorMessage += str
				file[i] = append(file[i], str)
				fmt.Println(errorMessage)
				count = 1
			} else {
				if dbData[csvData[i].Integration_ID].SKU != csvData[i].SKU || dbData[csvData[i].Integration_ID].Colour_code != csvData[i].ProductColourCode || dbData[csvData[i].Integration_ID].Size != csvData[i].SizeBreak {
					str = errCarryOverYes
					errorMessage += str
					file[i] = append(file[i], str)
					fmt.Println(errorMessage)
					count = 1
				}

			}
		} else {
			if csvData[i].BrandscopeCarryOver == "Y" || csvData[i].BrandscopeCarryOver == "y" {
				str = errCarryOverYes
				errorMessage += str
				file[i] = append(file[i], str)
				fmt.Println(errorMessage)
				count = 1
			}
		}
		if count == 1 {
			continue
		}

		str, err = CheckValidations(csvData[i], i)

		if str == perfectEntry {
			verifiedFields[csvData[i].Integration_ID] = true
			fmt.Println(verifiedFields)

		}
		file[i] = append(file[i], str)
		if str != "" {
			errorMessage += strconv.Itoa(i)
		}
		errorMessage += str
	}

	csvFile, err := os.Create(file_name)
	ReturnError(err)
	csvwriter := csv.NewWriter(csvFile)

	for i := 0; i < len(file); i++ {
		_ = csvwriter.Write(file[i])
	}
	csvwriter.Flush()
	csvFile.Close()
	return errorMessage, nil
}

func CheckValidations(data BrandHeader, i int) (errorstring string, err error) {

	err = AgeGroupValidations(data.AgeGroup)
	if err == errAgeGroup {

		errorstring += data.AgeGroup
		errorstring += InvalidAgeGroup
	}

	err = AttributeTypeValidation(data.AttributeType)
	if err == errAttributeTypeNotValid {
		errorstring += InvalidAttributeType
	}

	err = AtsInIndentValidation(data.AtsInIndent)
	if err == errInvalidAtsInIndent {
		errorstring += InvalidAtsInIndent
	}

	err = AtsInInseasonValidation(data.AtsInInSeason)
	if err == errInvalidAtsInInSeason {
		errorstring += InvalidAtsInInSeason
	}

	err = AttributeValueValidation(data.AttributeValue)
	if err == errInvalidAttributeValue {
		errorstring += InvalidAttributeValue
	}

	err = BarcodeValidation(data.Barcode)
	if err == errInvalidBarcode {
		errorstring += InvalidBarcode
	}

	err = BrandscopeCarryOverValidation(data.BrandscopeCarryOver)
	if err == errBrandScopeCarryOverEmpty {
		errorstring += EmptyBrandscopeCarryOver
	} else if err == errBrandScopeCarryOverNotValid {
		errorstring += InvalidBrandscopeCarryOver
	} else {
	}

	err = BrandscopeHierarchyValidation(data.BrandscopeHierarchy)
	if err == errBrandscopeHierarchyEmpty {
		errorstring += EmptyBrandscopeHierarchy
	}

	err = BrandNameValidation(data.BrandName)
	if err == errBrandNameEmpty {
		errorstring += EmptyBrandName
	} else if err == errInvalidBrandName {
		errorstring += InvalidBrandName
	}

	err = CatalogueOrderValidation(data.CatalogueOrder)
	if err == errCatalogueOrderempty {
		errorstring += CatalogueOrderEmpty
	} else if err == errCatalogueOrderNotaNumber {
		errorstring += CatalogueOrderNotANumber
	} else {
	}

	err = CategoriesValidation(data.Categories)

	if err == errInvalidCategories {
		errorstring += InvalidCategories
	}

	err = CollectionsValidation(data.Collections)
	if err == errInvalidCollections {
		errorstring += InvalidCollections
	}

	err = CompanyNameValidation(data.CompanyName)
	if err == errCompanyNameEmpty {
		errorstring += EmptyCompanyName
	} else if err == errInvalidCompanyName {
		errorstring += InvalidCompanyName
	}

	err = DisplayRetailValidation(data.DisplayRetail)
	if err == errDisplayRetailEmpty {
		errorstring += EmptyDisplayRetail
	}

	err = DisplayWholesaleValidation(data.DisplayWholesale)
	if err == errDisplayWholesaleEmpty {
		errorstring += EmptyDisplayWholesale
	}

	err = DisplayWholesaleRangeValidation(data.DisplayWholesaleRange)
	if err == errDisplayWholesaleRangeNotValid {
		errorstring += InvalidDisplayWholesaleRange
	}

	err = GenderValidations(data.Gender)
	if err == errGender {

		errorstring += data.Gender
		errorstring += InvalidGender
	}

	err = GenericColorValidation(data.GenericColour)
	if err == errInvalidGenericColour {
		errorstring += InvalidGenericColour
	}

	err = Integration_IDValidations(data.Integration_ID, i)
	if err == errIntegration_IDEmpty {
		errorstring += EmptyIntegration_ID
	} else if err == errIntIDAlreadyExists {
		errorstring += Integration_IDAlreadyExists
	}

	err = MarketingSupportValidation(data.MarketingSupport)
	if err == errInvalidMarketingSupport {
		errorstring += InvalidMarketingSupport
	}

	err = PackUnitsValidation(data.PackUnits)
	if err == errPackUnitsEmpty {
		errorstring += EmptyPackUnits
	} else if err == errInvalidPackUnitsValue {
		errorstring += InvalidPackUnits
	}

	err = ProductColourCodeValidation(data.ProductColourCode)
	if err == errProductColourCodeNotValid {
		errorstring += InvalidProductColourCode
	}

	err = ProductDisplayColourValidation(data.ProductDisplayColour)
	if err == errProductDisplayColourNotValid {
		errorstring += InvalidProductDisplayColour
	}

	err = ProductMultipleValidation(data.ProductMultiple)
	if err == errInvalidData {
		errorstring += "ProductMultiple==> ProductMultiple not valid"
	}

	err = ProductNameValidation(data.ProductName)
	if err == errProductNameEmpty {
		errorstring += EmptyProductName
	}

	err = RetailPriceOriginalValidation(data.RetailPriceOriginal, data.WholesalePrice)
	if err == errRetailPriceOriginalEmpty {
		errorstring += EmptyRetailPriceOriginal
	} else if err == errInvalidRetailPriceOriginal {
		errorstring += InvalidRetailPriceOriginal
	}

	err = RetailPriceValidation(data.RetailPrice, data.WholesalePrice)
	if err == errRetailPriceEmpty {
		errorstring += EmptyRetailPrice
	} else if err == errInvalidRetailPriceValue {
		errorstring += InvalidRetailPriceValue
	}

	err = SalesTipValidation(data.SalesTip)
	if err == errInvalidSalesTip {
		errorstring += InvalidSalesTip
	}

	err = SegmentNameValidation(data.SegmentNames)
	if err == errInvalidSegmentNames {
		errorstring += InvalidSegmentNames
	}

	err = SizeBreakValidation(data.SizeBreak)
	if err == errInvalidSizeBreak {
		errorstring += InvalidSizeBreak
	}

	err = SKUValidations(data.SKU, i)
	if err == errSKUEmpty {
		errorstring += EmptySKU
	} else if err == errInvalidSKU {
		errorstring += InvalidSKU
	} else if err == errLength500 {
		errorstring += InvalidSKUlength
	} else {
	}

	err = WholesalePriceValidation(data.WholesalePrice)
	if err == errWholesalePriceEmpty {
		errorstring += EmptyWholesalePrice
	} else if err == errInvalidWholesalePrice {
		errorstring += InvalidWholesalePrice
	}

	err = WholesalePriceOriginalValidation(data.WholesalePriceOriginal)
	if err == errWholesalePriceOriginalEmpty {
		errorstring += EmptyWholesalePriceOriginal
	} else if err == errInvalidWholesalePriceOriginal {
		errorstring += InvalidWholesalePriceOriginal
	}

	err = StateValidation(data.State)
	if err == errInvalidState {
		errorstring += InvalidState
	}

	err = ProductSpecification1Validation(data.ProductSpecification1)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification1
	}
	err = ProductSpecification2Validation(data.ProductSpecification2)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification2
	}
	err = ProductSpecification3Validation(data.ProductSpecification3)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification3
	}
	err = ProductSpecification4Validation(data.ProductSpecification4)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification4
	}
	err = ProductSpecification5Validation(data.ProductSpecification5)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification5
	}
	err = ProductSpecification6Validation(data.ProductSpecification6)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification6
	}

	err = ProductSpecification7Validation(data.ProductSpecification7)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification7
	}
	err = ProductSpecification8Validation(data.ProductSpecification8)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification8
	}

	err = ProductSpecification9Validation(data.ProductSpecification9)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification9
	}
	err = ProductSpecification10Validation(data.ProductSpecification10)
	if err == errInvalidProductSpecification {
		errorstring += InvalidProductSpecification10
	}

	err = ProductChanges1Validation(data.ProductChanges1)
	if err == errInvalidProductChanges {
		errorstring += InvalidProductChanges1
	}
	err = ProductChanges2Validation(data.ProductChanges2)
	if err == errInvalidProductChanges {
		errorstring += InvalidProductChanges2
	}
	err = ProductChanges3Validation(data.ProductChanges3)
	if err == errInvalidProductChanges {
		errorstring += InvalidProductChanges3
	}

	err = ProductChanges4Validation(data.ProductChanges4)
	if err == errInvalidProductChanges {
		errorstring += InvalidProductChanges4
	}
	err = ProductChanges5Validation(data.ProductChanges5)
	if err == errInvalidProductChanges {
		errorstring += InvalidProductChanges5
	}

	err = AdditionalDetail1Validation(data.AdditionalDetail1)
	if err == errInvalidAdditionalDetail {
		errorstring += InvalidAdditionalDetail1
	}
	err = AdditionalDetail2Validation(data.AdditionalDetail2)
	if err == errInvalidAdditionalDetail {
		errorstring += InvalidAdditionalDetail2
	}

	err = AdditionalDetail3Validation(data.AdditionalDetail3)
	if err == errInvalidAdditionalDetail {
		errorstring += InvalidAdditionalDetail3
	}
	err = AdditionalDetail4Validation(data.AdditionalDetail4)
	if err == errInvalidAdditionalDetail {
		errorstring += InvalidAdditionalDetail4
	}

	err = AdditionalDetail5Validation(data.AdditionalDetail5)
	if err == errInvalidAdditionalDetail {
		errorstring += InvalidAdditionalDetail5
	}

	if errorstring == "" {
		errorstring = "ok"
	}

	return errorstring, nil

}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &CsvService{
		store:  s,
		logger: l,
	}
}
