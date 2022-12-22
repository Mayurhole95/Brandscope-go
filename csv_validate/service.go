package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Mayurhole95/Brandscope-go/db"
	"github.com/Mayurhole95/Brandscope-go/utils"
	csvtag "github.com/artonge/go-csv-tag/v2"
	"go.uber.org/zap"
)

type Service interface {
	Validate(ctx context.Context, id string) (success utils.Success, err error)
}

var csvData []BrandHeader

type CsvService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

var fileName string = "pride_priderelease_20221215114528.csv"
var fileNameErrors string = "pride_priderelease_20221215114528_errors.csv"

func (cs *CsvService) Validate(ctx context.Context, id string) (success utils.Success, err error) {
	brandID := "380"
	releaseID := "206"

	exist, err := cs.store.FindID(ctx, brandID, releaseID)
	utils.ReturnError(err)

	if !exist {
		success = utils.SuccessMessage(false, errBrandIDExists, "")
		return success, nil
	}
	missingheader, err := ValidateHeader()
	if err != nil {
		csvFile, err := os.Create(fileNameErrors)
		utils.ReturnError(err)
		csvWriter := csv.NewWriter(csvFile)

		// for _, j := range missingheader {
		// 	_ = csvWriter.Write(j)
		// }
		csvWriter.WriteAll(missingheader)
		csvWriter.Flush()
		csvFile.Close()
		success = utils.SuccessMessage(false, errHeadersMissing, fileNameErrors)

		return success, errNoData
	}
	csvDataMap, err := cs.store.ListData(brandID)
	utils.ReturnError(err)
	months, err := cs.store.ListMonths(releaseID)
	utils.ReturnError(err)
	dbMonths, err := ChangeDateFormat(months)
	fmt.Println(dbMonths)
	utils.ReturnError(err)
	errorstring, err := ValidateCSVData(csvDataMap)
	utils.ReturnError(err)
	if errorstring == "" {
		success = utils.SuccessMessage(true, perfectEntry, "")
	} else {
		success = utils.SuccessMessage(false, errorstring, "")
	}

	return success, nil
}

func readHeaders() []string {
	readCsvFile, err := os.Open(fileName)
	utils.ReturnError(err)
	defer readCsvFile.Close()

	csvReader := csv.NewReader(readCsvFile)
	records, err := csvReader.Read()
	utils.ReturnError(err)
	return records
}

func Iterate(n int) []struct{} {
	return make([]struct{}, n)
}

func ValidateHeader() (missingheader [][]string, err error) {
	var mapHeaders = make(map[string]int)
	headers := []string{"CatalogueOrder", "BrandscopeCarryOver", "Integration_ID", "Barcode", "SKU", "ProductName", "ProductColourCode", "ProductDisplayColour", "GenericColour", "SizeBreak", "AttributeValue", "AttributeType", "AttributeSequence", "DisplayWholesaleRange", "DisplayWholesale", "DisplayRetail", "PackUnits", "AvailableMonths", "AgeGroup", "Gender", "State", "PreOrderLeadTimeDays", "PreOrderMessage"}
	missingHeaders := make([]string, 0)
	missingHeaders2d := make([][]string, 0)
	records := readHeaders()
	lenarr := len(records)
	count := 0
	for colHeaders := range Iterate(len(headers)) {
		headerPresent := "false"
		for colCsv := range Iterate(lenarr) {
			if headers[colHeaders] == records[colCsv] {
				mapHeaders[headers[colHeaders]] = colCsv
				headerPresent = "true"
				count = count + 1
				break
			}

		}
		if headerPresent == "false" {
			missingHeaders = append(missingHeaders, headers[colHeaders])
			missingHeaders2d = append(missingHeaders2d, [][]string{missingHeaders}...)
			missingHeaders = missingHeaders[1:]
		}

	}

	if count < len(headers) {
		err = errors.New(errHeadersMissing)
		return missingHeaders2d, err
	}
	err = nil
	return missingHeaders2d, nil
}

func ValidateCSVData(dbData map[string]db.Verify) (errorMessage string, err error) {
	// var CsvData = []BrandHeader{}

	err = csvtag.LoadFromPath(fileName,
		&csvData,
		csvtag.CsvOptions{ // Load your csv with optional options
			Separator: ',', // changes the values separator, default to ','
		})
	fmt.Println(csvData)
	if err != nil {
		return "", err
	}
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
		fmt.Println("Hii")
		fmt.Println(record)
		_, ok := dbData[c.Integration_ID]
		if ok && c.BrandscopeCarryOver == "N" || c.BrandscopeCarryOver == "n" {
			errorMessage += errCarryOverNot
			record = append(record, errCarryOverNot)
			records = append(records, record)
			fmt.Println(errorMessage)
			continue
		} else if ok && dbData[c.Integration_ID].SKU != c.SKU || dbData[c.Integration_ID].Colour_code != c.ProductColourCode || dbData[c.Integration_ID].Size != c.SizeBreak {
			errorMessage += errCarryOverYes
			record = append(record, errCarryOverYes)
			records = append(records, record)
			fmt.Println(errorMessage)
			continue
		} else if c.BrandscopeCarryOver == "Y" || c.BrandscopeCarryOver == "y" {
			errorMessage += errCarryOverYes
			record = append(record, errCarryOverYes)
			records = append(records, record)
			fmt.Println(errorMessage)
			continue
		}

		resp, _ := CheckValidations(c, i)
		// check for error
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

	csvFile, err := os.Create(fileNameErrors)
	utils.ReturnError(err)
	csvWriter := csv.NewWriter(csvFile)

	err = csvWriter.WriteAll(records)
	utils.ReturnError(err)
	//handle error
	csvWriter.Flush()
	csvFile.Close()
	return errorMessage, nil
}

// func ValidateCSVData(dbData map[string]db.Verify) (errorMessage string, err error) {
// 	// var CsvData = []BrandHeader{}
// 	err = csvtag.LoadFromPath(fileName,
// 		&csvData,
// 		csvtag.CsvOptions{ // Load your csv with optional options
// 			Separator: ',', // changes the values separator, default to ','
// 		})
// 	if err != nil {
// 		return "", err
// 	}

// 	var file [][]string
// 	readCsvFile, err := os.Open(fileName)
// 	utils.ReturnError(err)(err)

// 	// remember to close the file at the end of the program
// 	defer readCsvFile.Close()
// 	csvReader := csv.NewReader(readCsvFile)
// 	// csvReader.FieldsPerRecord = -1
// 	file, err = csvReader.ReadAll()
// 	utils.ReturnError(err)(err)

// 	errorMessage = ""
// 	str := ""
// 	file[0] = append(file[0], "status")
// 	verifiedFields := make(map[string]bool)
// 	for i := 0; i < len(csvData); i++ {
// 		str = ""

// 		print(verifiedFields)

// 		if _, find := verifiedFields[csvData[i].Integration_ID]; !find {
// 			continue

// 		}

// 		if _, ok := dbData[csvData[i].Integration_ID]; ok {
// 			if csvData[i].BrandscopeCarryOver == "N" || csvData[i].BrandscopeCarryOver == "n" {
// 				// fmt.Println(i, "Present")
// 				str = errCarryOverNot
// 				errorMessage += str
// 				file[i] = append(file[i], str)
// 				fmt.Println(errorMessage)
// 				continue
// 			} else {
// 				if dbData[csvData[i].Integration_ID].SKU != csvData[i].SKU || dbData[csvData[i].Integration_ID].Colour_code != csvData[i].ProductColourCode || dbData[csvData[i].Integration_ID].Size != csvData[i].SizeBreak {
// 					str = errCarryOverYes
// 					errorMessage += str
// 					file[i] = append(file[i], str)
// 					fmt.Println(errorMessage)
// 					continue
// 				}

// 			}
// 		} else {
// 			if csvData[i].BrandscopeCarryOver == "Y" || csvData[i].BrandscopeCarryOver == "y" {
// 				str = errCarryOverYes
// 				errorMessage += str
// 				file[i] = append(file[i], str)
// 				fmt.Println(errorMessage)
// 				continue
// 			}
// 		}

// 		str, err = CheckValidations(csvData[i], i)

// 		if str == perfectEntry {
// 			verifiedFields[csvData[i].Integration_ID] = true
// 			fmt.Println(verifiedFields)

// 		}
// 		file[i] = append(file[i], str)
// 		if str != "" {
// 			errorMessage += strconv.Itoa(i)
// 		}
// 		errorMessage += str
// 	}

// 	csvFile, err := os.Create(fileNameErrors)
// 	utils.ReturnError(err)(err)
// 	csvWriter := csv.NewWriter(csvFile)

// 	for i := 0; i < len(file); i++ {
// 		_ = csvWriter.Write(file[i])
// 	}
// 	csvWriter.Flush()
// 	csvFile.Close()
// 	return errorMessage, nil
// }

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
	} else if err == errInvalidIntegration_ID {
		errorstring += InvalidIntegration_ID
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
	if err == errInvalidProductMultiple {
		errorstring += InvalidProductMultiple
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

	errstr := SpecificationValidation("1", data.ProductSpecification1)
	if errstr == "1" {
		errorstring += InvalidProductSpecification1
	}
	errstr = SpecificationValidation("2", data.ProductSpecification2)
	if errstr == "2" {
		errorstring += InvalidProductSpecification2
	}
	errstr = SpecificationValidation("3", data.ProductSpecification3)
	if errstr == "3" {
		errorstring += InvalidProductSpecification3
	}
	errstr = SpecificationValidation("4", data.ProductSpecification4)
	if errstr == "4" {
		errorstring += InvalidProductSpecification4
	}
	errstr = SpecificationValidation("5", data.ProductSpecification5)
	if errstr == "5" {
		errorstring += InvalidProductSpecification5
	}
	errstr = SpecificationValidation("6", data.ProductSpecification6)
	if errstr == "6" {
		errorstring += InvalidProductSpecification6
	}

	errstr = SpecificationValidation("7", data.ProductSpecification7)
	if errstr == "7" {
		errorstring += InvalidProductSpecification7
	}
	errstr = SpecificationValidation("8", data.ProductSpecification8)
	if errstr == "8" {
		errorstring += InvalidProductSpecification8
	}

	errstr = SpecificationValidation("9", data.ProductSpecification9)
	if errstr == "9" {
		errorstring += InvalidProductSpecification9
	}
	errstr = SpecificationValidation("10", data.ProductSpecification4)
	if errstr == "10" {
		errorstring += InvalidProductSpecification10
	}

	errstr = SpecificationValidation("1", data.ProductChanges1)
	if errstr == "1" {
		errorstring += InvalidProductChanges1
	}
	errstr = SpecificationValidation("2", data.ProductChanges2)
	if errstr == "2" {
		errorstring += InvalidProductChanges2
	}

	errstr = SpecificationValidation("3", data.ProductChanges3)
	if errstr == "3" {
		errorstring += InvalidProductChanges3
	}

	errstr = SpecificationValidation("4", data.ProductChanges4)
	if errstr == "4" {
		errorstring += InvalidProductChanges4
	}
	errstr = SpecificationValidation("5", data.ProductChanges5)
	if errstr == "5" {
		errorstring += InvalidProductChanges5
	}

	errstr = SpecificationValidation("1", data.AdditionalDetail1)
	if errstr == "1" {
		errorstring += InvalidAdditionalDetail1
	}
	errstr = SpecificationValidation("2", data.AdditionalDetail2)
	if errstr == "2" {
		errorstring += InvalidAdditionalDetail2
	}

	errstr = SpecificationValidation("3", data.AdditionalDetail3)
	if errstr == "3" {
		errorstring += InvalidAdditionalDetail3
	}
	errstr = SpecificationValidation("4", data.AdditionalDetail4)
	if errstr == "4" {
		errorstring += InvalidAdditionalDetail4
	}

	errstr = SpecificationValidation("5", data.AdditionalDetail5)
	if errstr == "5" {
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
