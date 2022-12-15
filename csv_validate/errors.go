package csv

import "errors"

var (
	errNoData                      = errors.New("Headers missing")
	errCatalogueOrderNotANumber    = errors.New("CatalogueOrder should be a number")
	errCatalogueOrderEmpty         = errors.New("CatalogueOdrder can't be empty")
	errBrandScopeCarryOverEmpty    = errors.New("BrandScopeCarryOver cannot be empty")
	errBrandScopeCarryOverNotValid = errors.New("BrandscopeCarryOver not Valid")
	errIntegration_IDEmpty         = errors.New("Integration_ID can't be empty")
	errProductNameEmpty            = errors.New("ProductName cannot be empty")
	errDisplayWholesaleEmpty       = errors.New("DisplayWholesale cannot be empty")
	errDisplayRetailEmpty          = errors.New("DisplayRetail cannot be empty")
	errDisplayPackUnitsEmpty       = errors.New("PackUnits cannot be empty")
	errPackUnitsNotANumber         = errors.New("PackUnits should be a number")
	sucessfulvalidation            = ("Validation Successful")
	errInvalidData                 = errors.New("Invalid Data .Entry should be alphanumeric")
	errLength500                   = errors.New("Length should be les than 500")
	errLength80                    = errors.New("Length should be less than 80")
	errSKUEmpty                    = errors.New("SKU can't be empty")
	errWholesalePriceOriginalEmpty = errors.New("WholesalPriceOriginal cannot be empty")
	errWholesalePriceEmpty         = errors.New("WholesalPrice cannot be empty")
	errGender                      = errors.New("<Gender> is not a valid Gender entered. Valid values are:  Male, Female, Unisex")
	errAgeGroup                    = errors.New("<AgeGroup> is not a valid AgeGroup. Valid values are: Infant, Kid, Youth, Adult or Any")
	errBrandNameEmpty              = errors.New("BrandName can't be empty")
	errCompanyNameEmpty            = errors.New("CompanyName can't be empty")
	errIntIDExists                 = errors.New("Entry with similar Integration ID exists")
	errSKUExists                   = errors.New("Entry with similar SKU exists")
	errProductExists               = errors.New("Product with similar SKU+ProductColorCode+SizeBreak exists")
	errProductDoesntExist          = errors.New("This product is flagged as a carry-over, but no product can be found with the SKU/Colour/Size combination of <SKU>/<ProductColourCode>/<SizeBreak>.")
	errCompanyDoesNotExist         = errors.New("Comany does not exist in Brandscope.")

	errIntegration_IDNotValid                = errors.New("Integration_ID not valid")
	errPackUnitsEmpty                        = errors.New("PackUnits cannot be empty")
	errInvalidPackUnitsValue                 = errors.New("PackUnits should be >0")
	errProductColourCodeNotValid             = errors.New("ProductColourCode is not Valid")
	errProductDisplayColourNotValid          = errors.New("ProductColourCode Invalid")
	errAttributeTypeNotValid                 = errors.New("AttributeType Invalid")
	errInvalidDisplayWholesaleValue          = errors.New("DisplayWholesale Value should be greater than or equal to 0")
	errDisplayWholesaleNotANumber            = errors.New("DisplayWholesale Not A Number")
	errDisplayWholesaleNotAValidAmount       = errors.New("DisplayWholesale should be greater than or equal to 0")
	errDisplayWholesaleRangeNotValid         = errors.New("DisplayWholesaleRange Not Valid")
	errRetailPriceOriginalValidationNotValid = errors.New("RetailPriceOriginal Invalid")
	errRetailPriceOriginalEmpty              = errors.New("RetailPriceOriginal cannot be empty")
	errInvalidRetailPriceOriginal            = errors.New("RetailPriceOriginal should be >=0, should be >=WholesalePrice, and cannot have $ symbol")
	errRetailPriceEmpty                      = errors.New("RetailPrice cannot be empty")
	errInvalidRetailPriceValue               = errors.New("RetailPrice should be >=0, should be >=WholesalePrice, and cannot have $ symbol")
	errInvalidCollections                    = errors.New("Collections Invalid")
	errInvalidCategories                     = errors.New("Categories Invalid")
	errBrandscopeHierarchyEmpty              = errors.New("BrandscopeHierarchy Cannot be empty")
	errInvalidState                          = errors.New("Invalid State")
	errInvalidProductSpecification           = errors.New("ProductSpecification Invalid")
	errInvalidProductChanges                 = errors.New("ProductChanges Invalid")
	errInvalidAdditionalDetail               = errors.New("AdditionalDetail Invalid")
	errInvalidSalesTip                       = errors.New("SalesTip Invalid")
	errInvalidMarketingSupport               = errors.New("MarketingSupport Invalid")
	errInvalidCompanyName                    = errors.New("Invalid CompanyName")
	errEntryFound                            = errors.New("Entry with similar data found")
	errDataNotAlphanumeric                   = errors.New("Please enter AlphaNumeric Value")
	errAvailableMonthsEmpty                  = errors.New("AvailableMonths can't be empty")

	errHeadersMissing = "Headers Missing"
	errHeadersFound   = "Headers Found"
	errBrandIDExists  = "Brand id doesn't exist"
	errCarryOverNot   = "This product is NOT flagged as a carry-over product, but there is already a product with the SKU/Colour/Size combination."
	errCarryOverYes   = "This product is flagged as a carry-over product, but there is not a product with the SKU/Colour/Size combination."
	perfectEntry      = "ok"
)
