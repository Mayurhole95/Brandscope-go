package csv_validate

import "errors"

var (
	errCatalogueOrderNotANumber              = errors.New("CatalogueOrder should be a number")
	errCatalogueOrderEmpty                   = errors.New("CatalogueOdrder can't be empty")
	errIntIDExists                           = errors.New("Entry with similar Integration ID exists")
	errAvailableMonthsError                  = errors.New("Error Available Months")
	errIntegration_IDNotValid                = errors.New("Integration_ID not valid")
	errRetailPriceOriginalValidationNotValid = errors.New("RetailPriceOriginal Invalid")
	errInvalidProductSpecification           = errors.New("ProductSpecification Invalid")
	errInvalidProductChanges                 = errors.New("ProductChanges Invalid")
	errInvalidAdditionalDetail               = errors.New("AdditionalDetail Invalid")

	errBrandIDExists = "Brand id doesn't exist"
	errCarryOverNot  = "This product is NOT flagged as a carry-over product, but there is already a product with the SKU/Colour/Size combination."
	errCarryOverYes  = "This product is flagged as a carry-over product, but there is not a product with the SKU/Colour/Size combination."
	perfectEntry     = "ok"

	InvalidProductSpecification1  = "ProductSpecification1 contains Illegal characters"
	InvalidProductSpecification2  = "ProductSpecification2 contains Illegal characters"
	InvalidProductSpecification3  = "ProductSpecification3 contains Illegal characters"
	InvalidProductSpecification4  = "ProductSpecification4 contains Illegal characters"
	InvalidProductSpecification5  = "ProductSpecification5 contains Illegal characters"
	InvalidProductSpecification6  = "ProductSpecification6 contains Illegal characters"
	InvalidProductSpecification7  = "ProductSpecification7 contains Illegal characters"
	InvalidProductSpecification8  = "ProductSpecification8 contains Illegal characters"
	InvalidProductSpecification9  = "ProductSpecification9 contains Illegal characters"
	InvalidProductSpecification10 = "ProductSpecification10 contains Illegal characters"

	InvalidProductChanges1 = "ProductChanges1 contains Illegal characters"
	InvalidProductChanges2 = "ProductChanges2 contains Illegal characters"
	InvalidProductChanges3 = "ProductChanges3 contains Illegal characters"
	InvalidProductChanges4 = "ProductChanges4 contains Illegal characters"
	InvalidProductChanges5 = "ProductChanges5 contains Illegal characters"

	InvalidAdditionalDetail1 = "AdditionalDetail1 contains Illegal characters"
	InvalidAdditionalDetail2 = "AdditionalDetail2 contains Illegal characters"
	InvalidAdditionalDetail3 = "AdditionalDetail3 contains Illegal characters"
	InvalidAdditionalDetail4 = "AdditionalDetail4 contains Illegal characters"
	InvalidAdditionalDetail5 = "AdditionalDetail5 contains Illegal characters"

	errAgeGroup     = errors.New("<AgeGroup> is not a valid AgeGroup. Valid values are: Infant, Kid, Youth, Adult or Any")
	InvalidAgeGroup = " is not a valid AgeGroup. Valid values are: Infant, Kid, Youth, Adult or Any"

	errAttributeTypeNotValid = errors.New("AttributeType Invalid")
	InvalidAttributeType     = "AttributeType contains Illegal characters"

	errInvalidAtsInIndent = errors.New("AtsIndent Invalid")
	InvalidAtsInIndent    = "AtsIndent contains Illegal characters"

	errInvalidAtsInInSeason = errors.New("AtsInseason Invalid")
	InvalidAtsInInSeason    = "AtsInSeason contains Illegal characters"

	errInvalidAttributeValue = errors.New("Invalid AttributeValue")
	InvalidAttributeValue    = "AttributeValue contains Illegal characters"

	errInvalidBarcode = errors.New("Invalid Barcode")
	InvalidBarcode    = "Barcode contains Illegal characters"

	errBrandScopeCarryOverEmpty    = errors.New("BrandScopeCarryOver cannot be empty")
	EmptyBrandscopeCarryOver       = "BrandScopeCarryOver must be provided"
	errBrandScopeCarryOverNotValid = errors.New("BrandscopeCarryOver not Valid")

	InvalidBrandscopeCarryOver = "BrandscopeCarryOver contains Invalid characters"

	errBrandscopeHierarchyEmpty = errors.New("BrandscopeHierarchy Cannot be empty")
	EmptyBrandscopeHierarchy    = "BrandscopeHierarchy must be provided"

	errBrandNameEmpty   = errors.New("BrandName can't be empty")
	EmptyBrandName      = "BrandName must be provided"
	errInvalidBrandName = errors.New("Invalid BrandName")
	InvalidBrandName    = "BrandName contains Illegal characters"

	errCatalogueOrderNotaNumber = errors.New("CatalogueOrder should be a number")
	CatalogueOrderNotANumber    = "CatalogueOrder==> CatalogueOrder must be numeric"
	CatalogueOrderEmpty         = "CatalogueOrder==> CatalogueOrder should be a number"

	errInvalidCategories = errors.New("Invalid Categories")
	InvalidCategories    = "Categories contains Illegal characters"

	errInvalidCollections = errors.New("Invalid Collections")
	InvalidCollections    = "Collections contains Illegal characters"

	errCompanyNameEmpty   = errors.New("CompanyName cannot be empty")
	EmptyCompanyName      = "CompanyName must be provided"
	errInvalidCompanyName = errors.New("Invalid CompanyName")
	InvalidCompanyName    = "CompanyName contains Illegal characters"

	errDisplayRetailEmpty = errors.New("DisplayRetail cannot be empty")
	EmptyDisplayRetail    = "DisplayRetail must be provided"

	errDisplayWholesaleEmpty = errors.New("DisplayWholesale cannot be empty")
	EmptyDisplayWholesale    = "DisplayWholesale must be provided"

	errDisplayWholesaleRangeNotValid = errors.New("DisplayWholesaleRange Not Valid")
	InvalidDisplayWholesaleRange     = "DisplayWholesaleRange contains Illegal characters"

	errGender     = errors.New("<Gender> is not a valid Gender entered. Valid values are:  Male, Female, Unisex")
	InvalidGender = " is not a valid Gender entered. Valid values are:  Male, Female, Unisex"

	errInvalidGenericColour = errors.New("Invalid GenericColour")
	InvalidGenericColour    = "GenericColour contains Illegal characters"

	errIntegration_IDEmpty   = errors.New("Integration_ID can't be empty")
	EmptyIntegration_ID      = "Integration_ID must be provided"
	errInvalidIntegration_ID = errors.New("Integration_ID not valid")
	InvalidIntegration_ID    = "Integration_ID contains Illegal characters"

	errIntIDAlreadyExists       = errors.New("Entry with similar Integration ID exists")
	Integration_IDAlreadyExists = "Entry with similar Integration ID exists"

	errInvalidMarketingSupport = errors.New("MarketingSupport Invalid")
	InvalidMarketingSupport    = "MarketingSupport contains Illegal characters"

	errPackUnitsEmpty        = errors.New("PackUnits cannot be empty")
	EmptyPackUnits           = "PackUnits must be provided"
	errInvalidPackUnitsValue = errors.New("PackUnits should be >0")
	InvalidPackUnits         = "PackUnits should be a value >0"

	errProductColourCodeNotValid = errors.New("ProductColourCode is not Valid")
	InvalidProductColourCode     = "ProductColourCode contains Illegal characters"

	errProductDisplayColourNotValid = errors.New("ProductColourCode Invalid")
	InvalidProductDisplayColour     = "ProductDisplayColour contains Illegal characters"

	errInvalidProductMultiple = errors.New("Invalid ProductMultiple")
	InvalidProductMultiple    = "ProductMultiple contains Illegal characters"

	errProductNameEmpty = errors.New("ProductName cannot be empty")
	EmptyProductName    = "ProductName must be provided"

	errRetailPriceOriginalEmpty   = errors.New("RetailPriceOriginal cannot be empty")
	EmptyRetailPriceOriginal      = "RetailPriceOriginal must be provided"
	errInvalidRetailPriceOriginal = errors.New("RetailPriceOriginal should be >=0, should be >=WholesalePrice, and cannot have $ symbol")
	InvalidRetailPriceOriginal    = "RetailPriceOriginal should be >=0, should be >=WholesalePrice, and cannot have $ symbol"

	errRetailPriceEmpty = errors.New("RetailPrice cannot be empty")
	EmptyRetailPrice    = "RetailPrice must be provided"

	errInvalidRetailPriceValue = errors.New("RetailPrice should be >=0, should be >=WholesalePrice, and cannot have $ symbol")
	InvalidRetailPriceValue    = "RetailPriceValue should be >=0, should be >=WholesalePrice, and cannot have $ symbol"

	errInvalidSalesTip = errors.New("SalesTip Invalid")
	InvalidSalesTip    = "SalesTip contains Illegal characters"

	errInvalidSegmentNames = errors.New("Invalid SegmentName")
	InvalidSegmentNames    = "SegmentName contains Illegal characters"

	errInvalidSizeBreak = errors.New("Invalid SizeBreak")
	InvalidSizeBreak    = "Sizebreak contains Illegal characters"

	errSKUEmpty      = errors.New("SKU can't be empty")
	EmptySKU         = "SKU must be provided"
	errInvalidSKU    = errors.New("Invalid SKU")
	InvalidSKU       = "SKU contains Illegal characters"
	errLength500     = errors.New("Length should be les than 500")
	InvalidSKUlength = "SKU field length should be <500  characters"

	errWholesalePriceEmpty   = errors.New("WholesalPrice cannot be empty")
	EmptyWholesalePrice      = "WholesalePrice must be provided"
	errInvalidWholesalePrice = errors.New("Invalid WholesalePrice")
	InvalidWholesalePrice    = "WholesalePrice contains Illegal characters"

	errWholesalePriceOriginalEmpty   = errors.New("WholesalPriceOriginal cannot be empty")
	EmptyWholesalePriceOriginal      = "WholesalePriceOriginal must be provided"
	errInvalidWholesalePriceOriginal = errors.New("Invalid WholesalepriceOriginal")
	InvalidWholesalePriceOriginal    = "WholesalepriceOriginal contains Illegal characters"

	errInvalidState = errors.New("Invalid State")
	InvalidState    = "Invalid State value"

	errNoData           = errors.New("headers missing")
	sucessfulvalidation = ("Validation Successful")
	errInvalidData      = errors.New("invalid Data, please enter Alphanumeric value")

	errLength80 = errors.New("length should be less than 80")

	errSKUExists           = errors.New("entry with similar SKU exists")
	errProductExists       = errors.New("product with similar SKU+ProductColorCode+SizeBreak exists")
	errProductDoesntExist  = errors.New("this product is flagged as a carry-over, but no product can be found with the SKU/Colour/Size combination of <SKU>/<ProductColourCode>/<SizeBreak>.")
	errCompanyDoesNotExist = errors.New("comany does not exist in Brandscope.")

	errInvalidDisplayWholesaleValue    = errors.New("DisplayWholesale Value should be greater than or equal to 0")
	errDisplayWholesaleNotANumber      = errors.New("displayWholesale Not A Number")
	errDisplayWholesaleNotAValidAmount = errors.New("displayWholesale should be greater than or equal to 0")

	errEntryFound           = errors.New("Entry with similar data found")
	errDataNotAlphanumeric  = errors.New("Please enter AlphaNumeric Value")
	errAvailableMonthsEmpty = errors.New("AvailableMonths can't be empty")

	errHeadersMissing = "Headers Missing"
	errHeadersFound   = "Headers Found"
)
