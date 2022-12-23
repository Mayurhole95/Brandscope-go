package csv_validate

import (
	"strconv"
)

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
	}

	err = AttributeValueValidation(data.AttributeValue)
	if err == errInvalidData {
		errorstring += "AttributeValue not valid"
	}

	err = AvailableMonthsValidations(data.AvailableMonths)
	if err == errAvailableMonthsEmpty {
		errorstring += "Available Months empty"
	} else if err == errAvailableMonthsError {
		errorstring += "Available Months error"
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
	}

	err = BrandscopeHierarchyValidation(data.BrandscopeHierarchy)
	if err == errBrandscopeHierarchyEmpty {
		errorstring += "BrandscopeHierarchy cannot be empty"
	}

	err = BrandNameValidation(data.BrandName)
	if err == errBrandNameEmpty {
		errorstring += "BrandName==> BrandName can't be empty"
	} else if err == errInvalidData {
		errorstring += "BrandName==> BrandName not valid"
	} else if err == errBrandNotFound {
		errorstring += "BrandName==> BrandName not found"
	}

	err = CatalogueOrderValidation(data.CatalogueOrder)
	if err == errCatalogueOrderEmpty {
		errorstring += "CatalogueOrder==> CatalogueOrder can't be empty, "
	} else if err == errCatalogueOrderNotANumber {
		errorstring += "CatalogueOrder==> CatalogueOrder should be a number, "
	}

	err = CategoriesValidation(data.Categories)
	if err == errInvalidCategories {
		errorstring += "Categories Invalid"
	}

	err = CollectionsValidation(data.Collections)
	if err == errInvalidCollections {
		errorstring += "Collections Invalid"
	}

	err = CompanyNameValidation(data.CompanyName)
	if err == errCompanyNameEmpty {
		errorstring += "CompanyName==> CompanyName can't be empty"
	} else if err == errInvalidData {
		errorstring += "CompanyName==> CompanyName not valid"
	}

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
	}

	err = Integration_IDValidations(data.Integration_ID, i)
	if err == errIntegration_IDEmpty {
		errorstring += "Int id empty"
	} else if err == errIntIDExists {
		errorstring += "Int id exists"
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
	}

	err = ProductColourCodeValidation(data.ProductColourCode)
	if err == errProductColourCodeNotValid {
		errorstring += "ProductColourCode not valid"
	}

	err = ProductDisplayColourValidation(data.ProductDisplayColour)
	if err == errProductDisplayColourNotValid {
		errorstring += "ProductDisplayColour Invalid"
	}

	err = ProductMultipleValidation(data.ProductMultiple)
	if err == errInvalidData {
		errorstring += "ProductMultiple==> ProductMultiple not valid"
	}

	err = ProductNameValidation(data.ProductName)
	if err == errProductNameEmpty {
		errorstring += "ProductName==> ProductName cannot be empty, "
	}

	err = ReleaseNameValidation(data.ReleaseName)
	if err == errReleaseNameEmpty {
		errorstring += "ReleaseName==> ReleaseName can't be empty"
	} else if err == errInvalidData {
		errorstring += "ReleaseName==> ReleaseName not valid"
	} else if err == errReleaseNotFound {
		errorstring += "ReleaseName==> ReleaseName not found"
	}

	err = RetailPriceOriginalValidation(data.RetailPriceOriginal, data.WholesalePrice)
	if err == errRetailPriceOriginalEmpty {
		errorstring += " RetailPriceOriginal cannot be empty"
	} else if err == errInvalidRetailPriceOriginal {
		errorstring += "RetailPriceOriginal should be >=0, should be >=WholesalePrice, and cannot have $ symbol"
	}

	err = RetailPriceValidation(data.RetailPrice, data.WholesalePrice)
	if err == errRetailPriceEmpty {
		errorstring += " RetailPrice cannot be empty"
	} else if err == errInvalidRetailPriceValue {
		errorstring += "RetailPrice should be >=0, should be >=WholesalePrice, and cannot have $ symbol"
	}

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
	}

	err = StateValidation(data.State)
	if err == errInvalidState {
		errorstring += "Invalid State"
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
	if errstr == "2" {
		errorstring += InvalidAdditionalDetail3
	}
	errstr = SpecificationValidation("4", data.AdditionalDetail4)
	if errstr == "2" {
		errorstring += InvalidAdditionalDetail4
	}

	errstr = SpecificationValidation("5", data.AdditionalDetail5)
	if errstr == "5" {
		errorstring += InvalidAdditionalDetail5
	}

	if errorstring == "" {
		errorstring = "ok"
	}
	errorstring += strconv.Itoa(i)
	return errorstring, nil

}
