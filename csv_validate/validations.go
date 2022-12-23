package csv_validate

import (
	"regexp"
	"strings"
	"time"

	"strconv"

	"github.com/Mayurhole95/Brandscope-go/utils"
)

func AgeGroupValidations(s string) (err error) {
	if (s == "") && (s != "Infant") && (s != "Kid") && (s != "Youth") && (s != "Adult") && (s != "Any") {
		return errAgeGroup
	}
	return nil
}

func AtsIndentValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AtsInseasonValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AttributeTypeValidation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errAttributeTypeNotValid
	}

	return nil
}

func AttributeValueValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AvailableMonthsValidations(s string) (err error) {
	if s == "" {
		return errAvailableMonthsEmpty
	}
	dates := strings.Split(s, ",")
	if !Equal(dates, dbMonths) {
		return errAvailableMonthsError
	}
	return nil
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ChangeDateFormat(delivery_month []string) (months []string, err error) {
	for i := 0; i < len(delivery_month); i++ {
		date, err := time.Parse("2006-01-02", delivery_month[i])
		utils.ReturnError(err)
		year, month, _ := date.Date()
		delivery_month[i] = strings.ToUpper(month.String()[0:3] + strconv.Itoa(year)[2:4])
	}
	return delivery_month, err
}

func BarcodeValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func BrandNameValidation(s string) (err error) {
	if s == "" {
		return errBrandNameEmpty
	}
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func BrandscopeCarryOverValidation(s string) (err error) {
	if s == "" {
		return errBrandScopeCarryOverEmpty
	}
	if (s != "Y") && (s != "N") && (s != "y") && (s != "n") {
		return errBrandScopeCarryOverNotValid
	}
	return nil

}

func BrandscopeHierarchyValidation(s string) (err error) {
	if s == "" {
		return errBrandscopeHierarchyEmpty
	}
	return nil
}

func CatalogueOrderValidation(s string) (err error) {

	if s == "" {
		return errCatalogueOrderEmpty
	}
	_, err = strconv.Atoi(s)
	if err != nil {
		return errCatalogueOrderNotANumber
	}
	return nil
}

func CategoriesValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidCategories
	}
	return nil
}

func CollectionsValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidCollections
	}
	return nil
}

func CompanyNameValidation(s string) (err error) {
	if s == "" {
		return errCompanyNameEmpty
	}
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func DisplayRetailValidation(s string) (err error) {
	if s == "" {
		return errDisplayRetailEmpty
	}
	if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidData
	}
	return nil
}

func DisplayWholesaleValidation(s string) (err error) {
	if s == "" {
		return errDisplayWholesaleEmpty
	}
	if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidDisplayWholesaleValue
	}
	if strings.ContainsAny(s, "$") {
		return errInvalidDisplayWholesaleValue
	}
	return nil
}

func DisplayWholesaleRangeValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errDisplayWholesaleRangeNotValid
	}

	return nil
}

func GenderValidations(s string) (err error) {
	if (s == "") && (s != "Male") && (s != "Female") && (s != "Unisex") {
		return errGender
	}
	return nil
}

func GenericColorValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func Integration_IDValidations(s string, i int) (err error) {
	if s == "" {
		return errIntegration_IDEmpty
	}

	for j := 1; j < i; j++ {
		if s == csvData[j].Integration_ID {
			return errIntIDExists
		}
	}
	return nil
}

func MarketingSupportValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidMarketingSupport
	}
	return nil
}

func PackUnitsValidation(s string) (err error) {
	if s == "" {
		return errPackUnitsEmpty
	} else if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidPackUnitsValue
	}

	return nil
}

func ProductColourCodeValidation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errProductColourCodeNotValid
	}
	if len(s) > 500 {
		return errLength500
	}
	return nil
}

func ProductDisplayColourValidation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errProductDisplayColourNotValid
	}
	if len(s) > 80 {
		return errLength80
	}
	return nil

}

func ProductMultipleValidation(s string) (err error) {
	if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidData
	}
	return nil
}

func ProductNameValidation(s string) (err error) {
	if s == "" {
		return errProductNameEmpty
	}
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	if len(s) > 80 {
		return errLength80
	}
	return nil
}

func RetailPriceValidation(s string, s1 string) (err error) {
	if s == "" {
		return errRetailPriceEmpty
	}

	val1, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		return
	}
	if val, err := strconv.ParseFloat(s, 32); err != nil || val < 0 || val <= val1 || strings.ContainsAny(s, "$") {
		return errInvalidRetailPriceValue
	}

	return nil
}

func RetailPriceOriginalValidation(s string, s1 string) (err error) {

	if s == "" {
		return errRetailPriceOriginalEmpty
	}
	val1, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		return
	}
	if val, err := strconv.ParseFloat(s, 32); err != nil || val < 0 || val <= val1 || strings.ContainsAny(s, "$") {
		return errInvalidRetailPriceOriginal
	}

	return nil
}

func SalesTipValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidSalesTip
	}
	return nil
}

func SegmentNameValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func SizeBreakValidation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func SKUValidations(s string, i int) (err error) {
	if s == "" {
		return errSKUEmpty
	}
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	if len(s) > 500 {
		return errLength500
	}
	for j := 1; j < i; j++ {
		if s == csvData[j].SKU {
			return errSKUExists
		}
	}
	return nil
}

func SpecificationValidation(specNum string, csvData string) (err string) {
	if isAlphaNumeric(csvData) != nil {
		err = specNum
		return err
	}
	return
}

func StateValidation(s string) (err error) {
	if (s != "Active") && (s != "Inactive") && (s != "") && (s != "active") && (s != "inactive") {
		return errInvalidState
	}
	return nil
}

func UniqueProductValidations(s string, i int) (err error) {

	for j := 1; j < i; j++ {
		if strings.EqualFold(strings.ToLower(s), strings.ToLower(csvData[j].SKU+csvData[j].ProductColourCode+csvData[j].SizeBreak)) {
			return errProductExists
		} else if csvData[i].CompanyName == csvData[j].CompanyName {
			return errCompanyDoesNotExist
		}

	}

	return nil
}

func WholesalePriceValidation(s string) (err error) {
	if s == "" {
		return errWholesalePriceEmpty
	}
	if val, err := strconv.ParseFloat(s, 32); err != nil || val < 0 {
		return errInvalidData
	}
	return nil
}

func WholesalePriceOriginalValidation(s string) (err error) {
	if s == "" {
		return errWholesalePriceOriginalEmpty
	}
	if val, err := strconv.ParseFloat(s, 32); err != nil || val < 0 {
		return errInvalidData
	}
	return nil
}

func isAlphaNumeric(s string) (err error) {
	is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9\/ .,]*$`).MatchString(s)
	if !is_alphanumeric {
		return errInvalidData
	} else {
		return nil
	}
}
