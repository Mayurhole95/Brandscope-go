package csv_validate

import "strconv"

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
	}

	err = BrandNameValidation(data.BrandName)
	if err == errBrandNameEmpty {
		errorstring += "BrandName==> BrandName can't be empty"
	} else if err == errInvalidData {
		errorstring += "BrandName==> BrandName not valid"
	}

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
	}

	err = ProductSpecification7Validation(data.ProductSpecification7)
	if err == errInvalidData {
		errorstring += "ProductSpecification7 not valid"
	}
	err = ProductSpecification8Validation(data.ProductSpecification8)
	if err == errInvalidData {
		errorstring += "ProductSpecification8 not valid"
	}

	err = ProductSpecification9Validation(data.ProductSpecification9)
	if err == errInvalidData {
		errorstring += "ProductSpecification9 not valid"
	}
	err = ProductSpecification10Validation(data.ProductSpecification10)
	if err == errInvalidData {
		errorstring += "ProductSpecification10 not valid"
	}
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

	if errorstring == "" {
		errorstring = "ok"
	}
	errorstring += strconv.Itoa(i)
	return errorstring, nil

}
