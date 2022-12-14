package csv

import (
	"fmt"
	"regexp"
	"strings"

	"strconv"
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

// function to validate AttributeType
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
		// (isAlphaNumeric(s))
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

// function to validate DisplayRetail
func DisplayRetailValidation(s string) (err error) {
	if s == "" {
		return errDisplayRetailEmpty
	}
	if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidData
	}
	return nil
}

// function to validate DisplayWholesale
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

// function to validate Integration_ID
func Integration_IDValidations(s string, i int) (err error) {
	if s == "" {
		fmt.Println("Int id empty")
		return errIntegration_IDEmpty
	}
	for j := 1; j < i; j++ {
		if s == data[j].Integration_ID {
			fmt.Println("Int id exists")
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

// function to validate PackUnits
func PackUnitsValidation(s string) (err error) {
	if s == "" {
		return errPackUnitsEmpty
	} else if val, err := strconv.Atoi(s); err != nil || val < 0 {
		return errInvalidPackUnitsValue
	}

	return nil
}

// function to validate ProductColourCode
func ProductColourCodeValidation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errProductColourCodeNotValid
	}
	if len(s) > 500 {
		return errLength500
	}
	return nil
}

// function to validate ProductDisplayColour
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

// function to validate ProductName
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		if s == data[j].SKU {
			fmt.Println("SKU exists")
			return errSKUExists
		}
	}
	return nil
}

func StateValidation(s string) (err error) {
	if (s != "Active") && (s != "Inactive") && (s != "") && (s != "active") && (s != "inactive") {
		return errInvalidState
	}
	return nil
}

func UniqueProductValidations(s string, i int) (err error) {

	for j := 1; j < i; j++ {
		if strings.EqualFold(strings.ToLower(s), strings.ToLower(data[j].SKU+data[j].ProductColourCode+data[j].SizeBreak)) {
			fmt.Println("Similar product exists")
			return errProductExists
		} else if data[i].CompanyName == data[j].CompanyName {
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
