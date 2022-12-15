package csv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
<<<<<<< HEAD
	"log"
=======
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	"os"
	"strconv"

	"github.com/Mayurhole95/Brandscope-go/db"
	csvtag "github.com/artonge/go-csv-tag/v2"
	"go.uber.org/zap"
)

type Service interface {
	Validate(ctx context.Context, id string) (success Success, err error)
}

<<<<<<< HEAD
var map_headers = make(map[string]int)
=======
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
var csvData = []BrandHeader{}

type CsvService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

<<<<<<< HEAD
var brand_id string
var release_id string

func (cs *CsvService) Validate(ctx context.Context, id string) (success Success, err error) {
	brand_id = "76"
	release_id = "206"
	exist, err := cs.store.FindID(ctx, brand_id, release_id)
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}

	if exist {

		missingheader, err := HeaderCheck()
		if err != nil {
			csvFile, err := os.Create("pride_priderelease_20221122164529_errors.csv")
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
=======
var file_name string = "Dash_Summer 21_20221201121220.csv"
var file_name_errors string = "Dash_Summer 21_20221201121220_errors.csv"

func (cs *CsvService) Validate(ctx context.Context, id string) (success Success, err error) {
	brand_id := "76"
	release_id := "206"

	exist, err := cs.store.FindID(ctx, brand_id, release_id)
	ReturnError(err)

	if exist {
		missingheader, err := HeaderCheck()
		if err != nil {
			csvFile, err := os.Create(file_name)
			ReturnError(err)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
			csvwriter := csv.NewWriter(csvFile)
			for i := 0; i < len(missingheader); i++ {
				_ = csvwriter.Write(missingheader[i])
			}
			csvwriter.Flush()
			csvFile.Close()
<<<<<<< HEAD
			success.Success = false
			success.Message = "Headers Missing"
			success.Filepath = "pride_priderelease_20221122164529_errors.csv"
			return success, errNoData
		}

		success.Success = true
		success.Message = "Headers Found"
		success.Filepath = ""
		csvDataMap, err := cs.store.ListData(brand_id)
		if err != nil {
			fmt.Println("Error Occured :", err.Error())
		}
		errorstring, err := readData(csvDataMap)
		if err != nil {
			fmt.Println(err)
		}

		if errorstring == "" {
			success.Success = true
			success.Message = "ok"
			success.Filepath = ""
		} else {
			success.Success = false
			success.Message = errorstring
			success.Filepath = ""
		}
	} else {
		success.Success = false
		success.Message = "Brand id doesn't exist"
		success.Filepath = ""
=======
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
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}
	return success, nil
}

<<<<<<< HEAD
func HeaderCheck() (missingheader [][]string, err error) {
	headers := [23]string{"CatalogueOrder", "BrandscopeCarryOver", "Integration_ID", "Barcode", "SKU", "ProductName", "ProductColourCode", "ProductDisplayColour", "GenericColour", "SizeBreak", "AttributeValue", "AttributeType", "AttributeSequence", "DisplayWholesaleRange", "DisplayWholesale", "DisplayRetail", "PackUnits", "AvailableMonths", "AgeGroup", "Gender", "State", "PreOrderLeadTimeDays", "PreOrderMessage"}
	var missingheaders []string
	missingheaders2d := make([][]string, 0)
	f, err := os.Open("pride_priderelease_20221122164529.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	rec, err := csvReader.Read()

	if err != nil {
		log.Fatal(err)
	}
	lenarr := len(rec)
	count := 0
	for i := 0; i < 23; i++ {
		val := "true"
		for j := 0; j < lenarr; j++ {
			if headers[i] == rec[j] {
				map_headers[headers[i]] = j
				val = "false"
=======
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
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
				count = count + 1
				break
			}

		}
<<<<<<< HEAD
		if val == "true" {
			missingheaders = append(missingheaders, headers[i])
			missingheaders2d = append(missingheaders2d, [][]string{missingheaders}...)

=======
		if header_present == "false" {
			missingheaders = append(missingheaders, headers[column_headers])
			missingheaders2d = append(missingheaders2d, [][]string{missingheaders}...)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
			missingheaders = missingheaders[1:]
		}

	}

	if count < len(headers) {
<<<<<<< HEAD
		err = errors.New("Missinggg")
=======
		err = errors.New(errHeadersMissing)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
		return missingheaders2d, err
	}
	err = nil
	return missingheaders2d, nil
}

<<<<<<< HEAD
func readData(dbData map[string]db.Verify) (errorstring string, err error) {

	err = csvtag.LoadFromPath(
		"pride_priderelease_20221122164529.csv",
=======
func readCSVData(dbData map[string]db.Verify) (errorMessage string, err error) {
	// var CsvData = []BrandHeader{}
	err = csvtag.LoadFromPath(file_name,
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
		&csvData,
		csvtag.CsvOptions{ // Load your csv with optional options
			Separator: ',', // changes the values separator, default to ','
		})
	if err != nil {
		return "", err
	}

	var file [][]string
<<<<<<< HEAD
	f, err := os.Open("pride_priderelease_20221122164529.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	file, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	errorstring = ""
=======
	readCsvFile, err := os.Open(file_name)
	ReturnError(err)

	// remember to close the file at the end of the program
	defer readCsvFile.Close()
	csvReader := csv.NewReader(readCsvFile)
	// csvReader.FieldsPerRecord = -1
	file, err = csvReader.ReadAll()
	ReturnError(err)

	errorMessage = ""
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	str := ""
	file[0] = append(file[0], "status")
	count := 0
	verifiedFields := make(map[string]bool)
	for i := 0; i < len(csvData); i++ {
		count = 0
<<<<<<< HEAD
=======
		str = ""
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8

		print(verifiedFields)

		if _, find := verifiedFields[csvData[i].Integration_ID]; find {
<<<<<<< HEAD
			fmt.Println("Found")
			str = "This product is NOT flagged as a carry-over product, but there is already a product with the SKU/Colour/Size combination."
			count = 1
			fmt.Println("same")
=======
			// fmt.Println("Found")
			str = errCarryOverNot
			count = 1
			// fmt.Println("same")
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8

		}
		if count == 1 {
			continue
		}
		if _, ok := dbData[csvData[i].Integration_ID]; ok {
			if csvData[i].BrandscopeCarryOver == "N" || csvData[i].BrandscopeCarryOver == "n" {
<<<<<<< HEAD
				fmt.Println(i, "Present")
				str = "Present"
				errorstring += str
				file[i] = append(file[i], str)
				fmt.Println(errorstring)
				count = 1
			} else {
				if dbData[csvData[i].Integration_ID].SKU != csvData[i].SKU || dbData[csvData[i].Integration_ID].Colour_code != csvData[i].ProductColourCode || dbData[csvData[i].Integration_ID].Size != csvData[i].SizeBreak {
					fmt.Println(dbData[csvData[i].Integration_ID].SKU, csvData[i].SKU)
					fmt.Println(dbData[csvData[i].Integration_ID].Size, csvData[i].SizeBreak)
					fmt.Println(dbData[csvData[i].Integration_ID].Colour_code, csvData[i].ProductColourCode)
					fmt.Println(i, "This product is flagged as a carry-over product, but there is not a product with the SKU/Colour/Size combination.")
					str = "This product is flagged as a carry-over product, but there is not a product with the SKU/Colour/Size combination."
					errorstring += str
					file[i] = append(file[i], str)
					fmt.Println(errorstring)
=======
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
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
					count = 1
				}

			}
		} else {
			if csvData[i].BrandscopeCarryOver == "Y" || csvData[i].BrandscopeCarryOver == "y" {
<<<<<<< HEAD
				fmt.Println(i, "Not Present")

				str = "Not Present"
				errorstring += str
				file[i] = append(file[i], str)
				fmt.Println(errorstring)
=======
				str = errCarryOverYes
				errorMessage += str
				file[i] = append(file[i], str)
				fmt.Println(errorMessage)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
				count = 1
			}
		}
		if count == 1 {
			continue
		}

		str, err = CheckValidations(csvData[i], i)

<<<<<<< HEAD
		if str == "ok" {
=======
		if str == perfectEntry {
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
			verifiedFields[csvData[i].Integration_ID] = true
			fmt.Println(verifiedFields)

		}
		file[i] = append(file[i], str)
		if str != "" {
<<<<<<< HEAD
			errorstring += strconv.Itoa(i)
		}
		errorstring += str
	}

	csvFile, err := os.Create("pride_priderelease_20221122164529.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
=======
			errorMessage += strconv.Itoa(i)
		}
		errorMessage += str
	}

	csvFile, err := os.Create(file_name)
	ReturnError(err)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	csvwriter := csv.NewWriter(csvFile)

	for i := 0; i < len(file); i++ {
		_ = csvwriter.Write(file[i])
	}
	csvwriter.Flush()
	csvFile.Close()
<<<<<<< HEAD
	return errorstring, nil
}

func CheckValidations(data BrandHeader, i int) (errorstring string, err error) {
	count := 0

	err = CatalogueOrderValidation(data.CatalogueOrder)
	if err == errCatalogueOrderEmpty {
		errorstring += "CatalogueOrder==> CatalogueOrder can't be empty, "
	} else if err == errCatalogueOrderNotANumber {
		errorstring += "CatalogueOrder==> CatalogueOrder should be a number, "
	} else {
		count += 1
	}

	err = BarcodeValidation(data.Barcode)
	if err == errInvalidData {
		errorstring += "Barcode==> Invalid Data .Entry should be alphanumeric, "
	} else {
		count += 1
	}

	err = SKUValidations(data.SKU, i)
	if err == errSKUEmpty {
		errorstring += "SKU==> SKU can't be empty, "
	} else if err == errInvalidData {
		errorstring += "SKU==> Invalid Data .Entry should be alphanumeric, "
	} else if err == errLength500 {
		errorstring += "SKU==> Length should be les than 500, "
	} else {
		count += 1
	}

	err = BrandscopeCarryOverValidation(data.BrandscopeCarryOver)
	if err == errBrandScopeCarryOverEmpty {
		errorstring += "BrandscopeCarryOver==> BrandscopeCarryOver cannot be empty, "
	} else if err == errBrandScopeCarryOverNotValid {
		errorstring += "BrandscopeCarryOver==> BrandscopeCarryOver not Valid, "
	} else {
		count += 1
	}

	err = ProductNameValidation(data.ProductName)
	if err == errProductNameEmpty {
		errorstring += "ProductName==> ProductName cannot be empty, "
	} else {
		count += 1
	}

	err = GenericColorValidation(data.GenericColour)
	if err == errInvalidData {
		errorstring += "GenericColor==> GenericColor not valid"
	} else {
		count += 1
	}

	err = DisplayWholesaleValidation(data.DisplayWholesale)
	if err == errDisplayWholesaleEmpty {
		errorstring += "DisplayWholesale==> DisplayWholesale cannot be empty, "
	} else {
		count += 1
	}

	err = DisplayRetailValidation(data.DisplayRetail)
	if err == errDisplayRetailEmpty {
		errorstring += "DisplayRetail==> DisplayRetail cannot be empty "
	} else {
		count += 1

	}

	err = SizeBreakValidation(data.SizeBreak)
	if err == errInvalidData {
		errorstring += "SizeBreak==> SizeBreak not valid"
	} else {
		count += 1
=======
	return errorMessage, nil
}

func CheckValidations(data BrandHeader, i int) (errorstring string, err error) {

	err = AgeGroupValidations(data.AgeGroup)
	if err == errAgeGroup {
		errorstring += "AgeGroup==>"
		errorstring += data.AgeGroup
		errorstring += "is not a valid AgeGroup. Valid values are: Infant, Kid, Youth, Adult or Any"
	}

	err = AttributeTypeValidation(data.AttributeType)
	if err == errAttributeTypeNotValid {
		errorstring += "AttributeType Invalid"
	}

	err = AtsIndentValidation(data.AtsInIndent)
	if err == errInvalidData {
		errorstring += "AtsInIndent==> AtsInIndent not valid"
	}

	err = AtsInseasonValidation(data.AtsInInSeason)
	if err == errInvalidData {
		errorstring += "AtsInInSeason ==> AtsInInSeason not valid"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = AttributeValueValidation(data.AttributeValue)
	if err == errInvalidData {
		errorstring += "AttributeValue not valid"
<<<<<<< HEAD
	} else {
		count += 1
	}

	err = WholesalePriceOriginalValidation(data.WholesalePriceOriginal)
	if err == errWholesalePriceOriginalEmpty {
		errorstring += "WholesalePriceOriginal can't be empty"
	} else if err == errInvalidData {
		errorstring += "WholesalePriceOriginal not valid"
	} else {
		count += 1
	}

	err = WholesalePriceValidation(data.WholesalePrice)
	if err == errWholesalePriceEmpty {
		errorstring += "WholesalePrice==> WholesalePrice can't be empty"
	} else if err == errInvalidData {
		errorstring += "WholesalePrice==> WholesalePrice not valid"
	} else {
		count += 1
	}

	err = ProductMultipleValidation(data.ProductMultiple)
	if err == errInvalidData {
		errorstring += "ProductMultiple==> ProductMultiple not valid"
	} else {
		count += 1
	}

	err = GenderValidations(data.Gender)
	if err == errGender {
		errorstring += "Gender==>"
		errorstring += data.Gender
		errorstring += "is not a valid Gender entered. Valid values are:  Male, Female, Unisex"
	} else {
		count += 1
	}

	err = AgeGroupValidations(data.AgeGroup)
	if err == errAgeGroup {
		errorstring += "AgeGroup==>"
		errorstring += data.AgeGroup
		errorstring += "is not a valid AgeGroup. Valid values are: Infant, Kid, Youth, Adult or Any"
	} else {
		count += 1
=======
	}

	err = BarcodeValidation(data.Barcode)
	if err == errInvalidData {
		errorstring += "Barcode==> Invalid Data .Entry should be alphanumeric, "
	}

	err = BrandscopeCarryOverValidation(data.BrandscopeCarryOver)
	if err == errBrandScopeCarryOverEmpty {
		errorstring += "BrandscopeCarryOver==> BrandscopeCarryOver cannot be empty, "
	} else if err == errBrandScopeCarryOverNotValid {
		errorstring += "BrandscopeCarryOver==> BrandscopeCarryOver not Valid, "
	} else {
	}

	err = BrandscopeHierarchyValidation(data.BrandscopeHierarchy)
	if err == errBrandscopeHierarchyEmpty {
		errorstring += "BrandscopeHierarchy cannot be empty"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = BrandNameValidation(data.BrandName)
	if err == errBrandNameEmpty {
		errorstring += "BrandName==> BrandName can't be empty"
	} else if err == errInvalidData {
		errorstring += "BrandName==> BrandName not valid"
	}

<<<<<<< HEAD
=======
	err = CatalogueOrderValidation(data.CatalogueOrder)
	if err == errCatalogueOrderEmpty {
		errorstring += "CatalogueOrder==> CatalogueOrder can't be empty, "
	} else if err == errCatalogueOrderNotANumber {
		errorstring += "CatalogueOrder==> CatalogueOrder should be a number, "
	} else {
	}

	err = CategoriesValidation(data.Categories)
	//fmt.Println(data.Categories)
	if err == errInvalidCategories {
		errorstring += "Categories Invalid"
	}

	err = CollectionsValidation(data.Collections)
	if err == errInvalidCollections {
		errorstring += "Collections Invalid"
	}

>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	err = CompanyNameValidation(data.CompanyName)
	if err == errCompanyNameEmpty {
		errorstring += "CompanyName==> CompanyName can't be empty"
	} else if err == errInvalidData {
		errorstring += "CompanyName==> CompanyName not valid"
	}

<<<<<<< HEAD
	err = SegmentNameValidation(data.SegmentNames)
	if err == errInvalidData {
		errorstring += "SegmentNames==> SegmentNames not valid"
	} else {
		count += 1
	}

	err = AtsIndentValidation(data.AtsInIndent)
	if err == errInvalidData {
		errorstring += "AtsInIndent==> AtsInIndent not valid"
	} else {
		count += 1
	}

	err = AtsInseasonValidation(data.AtsInInSeason)
	if err == errInvalidData {
		errorstring += "AtsInInSeason ==> AtsInInSeason not valid"
	} else {
		count += 1
=======
	err = DisplayRetailValidation(data.DisplayRetail)
	if err == errDisplayRetailEmpty {
		errorstring += "DisplayRetail==> DisplayRetail cannot be empty "
	}

	err = DisplayWholesaleRangeValidation(data.DisplayWholesaleRange)
	if err == errDisplayWholesaleRangeNotValid {
		errorstring += "DisplayWholesaleRange  Invalid"
	}

	err = DisplayWholesaleValidation(data.DisplayWholesale)
	if err == errDisplayWholesaleEmpty {
		errorstring += "DisplayWholesale==> DisplayWholesale cannot be empty, "
	}

	err = GenderValidations(data.Gender)
	if err == errGender {
		errorstring += "Gender==>"
		errorstring += data.Gender
		errorstring += "is not a valid Gender entered. Valid values are:  Male, Female, Unisex"
	}

	err = GenericColorValidation(data.GenericColour)
	if err == errInvalidData {
		errorstring += "GenericColor==> GenericColor not valid"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = Integration_IDValidations(data.Integration_ID, i)
	if err == errIntegration_IDEmpty {
		errorstring += "Int id empty"
	} else if err == errIntIDExists {
		errorstring += "Int id exists"
<<<<<<< HEAD
	} else {
		count += 1
=======
	}

	err = MarketingSupportValidation(data.MarketingSupport)
	if err == errInvalidMarketingSupport {
		errorstring += "Invalid MarketingSupport"
	}

	err = PackUnitsValidation(data.PackUnits)
	if err == errPackUnitsEmpty {
		errorstring += "PackUnits cannot be empty"
	} else if err == errInvalidPackUnitsValue {
		errorstring += "PackUnits should be >=0"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = ProductColourCodeValidation(data.ProductColourCode)
	if err == errProductColourCodeNotValid {
		errorstring += "ProductColourCode not valid"
	}

	err = ProductDisplayColourValidation(data.ProductDisplayColour)
	if err == errProductDisplayColourNotValid {
		errorstring += "ProductDisplayColour Invalid"
	}

<<<<<<< HEAD
	err = AttributeTypeValidation(data.AttributeType)
	if err == errAttributeTypeNotValid {
		errorstring += "AttributeType Invalid"
	}

	// err = DisplayWholesaleValidation(data.DisplayWholesale)
	// if err == errDisplayWholesaleEmpty {
	// 	errorstring += "DisplayWholesale cannot be empty"
	// } else if err == errInvalidDisplayWholesaleValue {
	// 	errorstring += "DisplayWholesale should be >=0 and cannot have $ symbol"
	// }

	err = DisplayWholesaleRangeValidation(data.DisplayWholesaleRange)
	if err == errDisplayWholesaleRangeNotValid {
		errorstring += "DisplayWholesaleRange  Invalid"
	}

	err = RetailPriceOriginalValidation(data.RetailPriceOriginal, data.WholesalePrice)
	//fmt.Println(data.RetailPriceOriginal)
	//fmt.Println(data.WholesalePrice)
=======
	err = ProductMultipleValidation(data.ProductMultiple)
	if err == errInvalidData {
		errorstring += "ProductMultiple==> ProductMultiple not valid"
	}

	err = ProductNameValidation(data.ProductName)
	if err == errProductNameEmpty {
		errorstring += "ProductName==> ProductName cannot be empty, "
	}

	err = RetailPriceOriginalValidation(data.RetailPriceOriginal, data.WholesalePrice)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	if err == errRetailPriceOriginalEmpty {
		errorstring += " RetailPriceOriginal cannot be empty"
	} else if err == errInvalidRetailPriceOriginal {
		errorstring += "RetailPriceOriginal should be >=0, should be >=WholesalePrice, and cannot have $ symbol"
	}

	err = RetailPriceValidation(data.RetailPrice, data.WholesalePrice)
<<<<<<< HEAD
	//fmt.Println(data.RetailPrice)
=======
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	if err == errRetailPriceEmpty {
		errorstring += " RetailPrice cannot be empty"
	} else if err == errInvalidRetailPriceValue {
		errorstring += "RetailPrice should be >=0, should be >=WholesalePrice, and cannot have $ symbol"
	}

<<<<<<< HEAD
	err = PackUnitsValidation(data.PackUnits)
	if err == errPackUnitsEmpty {
		errorstring += "PackUnits cannot be empty"
	} else if err == errInvalidPackUnitsValue {
		errorstring += "PackUnits should be >=0"
	}

	err = CollectionsValidation(data.Collections)
	if err == errInvalidCollections {
		errorstring += "Collections Invalid"
	}

	err = CategoriesValidation(data.Categories)
	//fmt.Println(data.Categories)
	if err == errInvalidCategories {
		errorstring += "Categories Invalid"
	}

	err = BrandscopeHierarchyValidation(data.BrandscopeHierarchy)
	if err == errBrandscopeHierarchyEmpty {
		errorstring += "BrandscopeHierarchy cannot be empty"
=======
	err = SalesTipValidation(data.SalesTip)
	if err == errInvalidSalesTip {
		errorstring += "Invalid SalesTip"
	}

	err = SegmentNameValidation(data.SegmentNames)
	if err == errInvalidData {
		errorstring += "SegmentNames==> SegmentNames not valid"
	}

	err = SizeBreakValidation(data.SizeBreak)
	if err == errInvalidData {
		errorstring += "SizeBreak==> SizeBreak not valid"
	}

	err = SKUValidations(data.SKU, i)
	if err == errSKUEmpty {
		errorstring += "SKU==> SKU can't be empty, "
	} else if err == errInvalidData {
		errorstring += "SKU==> Invalid Data .Entry should be alphanumeric, "
	} else if err == errLength500 {
		errorstring += "SKU==> Length should be les than 500, "
	} else {
	}

	err = WholesalePriceValidation(data.WholesalePrice)
	if err == errWholesalePriceEmpty {
		errorstring += "WholesalePrice==> WholesalePrice can't be empty"
	} else if err == errInvalidData {
		errorstring += "WholesalePrice==> WholesalePrice not valid"
	}

	err = WholesalePriceOriginalValidation(data.WholesalePriceOriginal)
	if err == errWholesalePriceOriginalEmpty {
		errorstring += "WholesalePriceOriginal can't be empty"
	} else if err == errInvalidData {
		errorstring += "WholesalePriceOriginal not valid"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = StateValidation(data.State)
	if err == errInvalidState {
		errorstring += "Invalid State"
	}

	err = ProductSpecification1Validation(data.ProductSpecification1)
	//fmt.Println(data.ProductSpecification1)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification1"
	}
	err = ProductSpecification2Validation(data.ProductSpecification2)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification2"
	}
	err = ProductSpecification3Validation(data.ProductSpecification3)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification3"
	}
	err = ProductSpecification4Validation(data.ProductSpecification4)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification4"
	}
	err = ProductSpecification5Validation(data.ProductSpecification5)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification5"
	}
	err = ProductSpecification6Validation(data.ProductSpecification6)
	if err == errInvalidData {
		errorstring += "ProductSpecification6 not valid"
<<<<<<< HEAD
	} else {
		count += 1
=======
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = ProductSpecification7Validation(data.ProductSpecification7)
	if err == errInvalidData {
		errorstring += "ProductSpecification7 not valid"
<<<<<<< HEAD
	} else {
		count += 1
	}

	err = ProductSpecification8Validation(data.ProductSpecification8)
	if err == errInvalidData {
		errorstring += "ProductSpecification8 not valid"
	} else {
		count += 1
=======
	}
	err = ProductSpecification8Validation(data.ProductSpecification8)
	if err == errInvalidData {
		errorstring += "ProductSpecification8 not valid"
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	}

	err = ProductSpecification9Validation(data.ProductSpecification9)
	if err == errInvalidData {
		errorstring += "ProductSpecification9 not valid"
<<<<<<< HEAD
	} else {
		count += 1
	}

	err = ProductSpecification10Validation(data.ProductSpecification10)
	if err == errInvalidData {
		errorstring += "ProductSpecification10 not valid"
	} else {
		count += 1
	}

=======
	}
	err = ProductSpecification10Validation(data.ProductSpecification10)
	if err == errInvalidData {
		errorstring += "ProductSpecification10 not valid"
	}
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	err = ProductSpecification11Validation(data.ProductSpecification1)
	//fmt.Println(data.ProductSpecification1)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification11"
	}
	err = ProductSpecification12Validation(data.ProductSpecification2)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification12"
	}
	err = ProductSpecification13Validation(data.ProductSpecification3)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification13"
	}
	err = ProductSpecification14Validation(data.ProductSpecification4)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification14"
	}
	err = ProductSpecification15Validation(data.ProductSpecification5)
	if err == errInvalidProductSpecification {
		errorstring += "Invalid ProductSpecification15"
	}

	err = ProductChanges1Validation(data.ProductChanges1)
	if err == errInvalidProductChanges {
		errorstring += "Invalid ProductChanges"
	}
	err = ProductChanges2Validation(data.ProductChanges2)
	if err == errInvalidProductChanges {
		errorstring += "Invalid ProductChanges"
	}
	err = ProductChanges3Validation(data.ProductChanges3)
	if err == errInvalidProductChanges {
		errorstring += "Invalid ProductChanges"
	}

	err = ProductChanges4Validation(data.ProductChanges2)
	if err == errInvalidProductChanges {
		errorstring += "Invalid ProductChanges"
	}
	err = ProductChanges5Validation(data.ProductChanges3)
	if err == errInvalidProductChanges {
		errorstring += "Invalid ProductChanges"
	}

	err = AdditionalDetail1Validation(data.AdditionalDetail1)
	if err == errInvalidAdditionalDetail {
		errorstring += "Invalid AdditionalDetail"
	}
	err = AdditionalDetail2Validation(data.AdditionalDetail2)
	if err == errInvalidAdditionalDetail {
		errorstring += "Invalid AdditionalDetail"
	}

	err = AdditionalDetail3Validation(data.AdditionalDetail1)
	if err == errInvalidAdditionalDetail {
		errorstring += "Invalid AdditionalDetail"
	}
	err = AdditionalDetail4Validation(data.AdditionalDetail2)
	if err == errInvalidAdditionalDetail {
		errorstring += "Invalid AdditionalDetail"
	}

	err = AdditionalDetail5Validation(data.AdditionalDetail1)
	if err == errInvalidAdditionalDetail {
		errorstring += "Invalid AdditionalDetail"
	}

<<<<<<< HEAD
	err = SalesTipValidation(data.SalesTip)
	if err == errInvalidSalesTip {
		errorstring += "Invalid SalesTip"
	}

	err = MarketingSupportValidation(data.MarketingSupport)
	if err == errInvalidMarketingSupport {
		errorstring += "Invalid MarketingSupport"
	}
	// err = errors.New(errorstring)
=======
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
	if errorstring == "" {
		errorstring = "ok"
	}

	return errorstring, nil
<<<<<<< HEAD
	// fmt.Println(errorstring)

}

//function to match values from csv with DB

// func matchValues(s string, s1 string, s2 string, s3 string) (err error) {
// 	for _, d := range data1 {
// 		if (d.Integration_ID == s) && (d.SKU == s1) && (d.Size == s2) && (d.Colour_code == s3) {
// 			return errEntryFound
// 		}
// 	}
// 	return nil
//
// }

=======

}

>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &CsvService{
		store:  s,
		logger: l,
	}
}
