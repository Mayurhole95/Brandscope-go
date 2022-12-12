package csv

type BrandHeader struct {
	CatalogueOrder         string `csv:"CatalogueOrder"`
	Integration_ID         string `csv:"Integration_ID"`
	BrandscopeCarryOver    string `csv:"BrandscopeCarryOver"`
	Barcode                string `csv:"Barcode"`
	SKU                    string `csv:"SKU"`
	ProductName            string `csv:"ProductName"`
	ProductColourCode      string `csv:"ProductColourCode"`
	ProductDisplayColour   string `csv:"ProductDisplayColour"`
	GenericColour          string `csv:"GenericColour"`
	SizeBreak              string `csv:"SizeBreak"`
	AttributeValue         string `csv:"AttributeValue"`
	AttributeType          string `csv:"AttributeType"`
	AttributeSequence      string `csv:"AttributeSequence"`
	DisplayWholesaleRange  string `csv:"DisplayWholesaleRange"`
	PreOrderLeadTimeDays   string `csv:"PreOrderLeadTimeDays"`
	PreOrderMessage        string `csv:"PreOrderMessage"`
	DisplayWholesale       string `csv:"DisplayWholesale"`
	DisplayRetail          string `csv:"DisplayRetail"`
	PackUnits              string `csv:"PackUnits"`
	AvailableMonths        string `csv:"AvailableMonths"`
	AgeGroup               string `csv:"AgeGroup"`
	Gender                 string `csv:"Gender"`
	State                  string `csv:"State"`
	WholesalePriceOriginal string `csv:"WholesalePriceOriginal"`
	WholesalePrice         string `csv:"WholesalePrice"`
	RetailPriceOriginal    string `csv:"RetailPriceOriginal"`
	RetailPrice            string `csv:"RetailPrice"`
	ProductMultiple        string `csv:"ProductMultiple"`
	Divisions              string `csv:"Divisions"`
	Collections            string `csv:"Collections"`
	Categories             string `csv:"Categories"`
	DiscountCategory       string `csv:"DiscountCategory"`
	ProductTags            string `csv:"ProductTags"`
	BrandscopeHierarchy    string `csv:"BrandscopeHierarchy"`
	ProductRequirement1    string `csv:"ProductRequirement1"`
	ProductSpecification1  string `csv:"ProductSpecification1"`
	ProductSpecification2  string `csv:"ProductSpecification2"`
	ProductSpecification3  string `csv:"ProductSpecification3"`
	ProductSpecification4  string `csv:"ProductSpecification4"`
	ProductSpecification5  string `csv:"ProductSpecification5"`
	ProductSpecification6  string `csv:"ProductSpecification6"`
	ProductSpecification7  string `csv:"ProductSpecification7"`
	ProductSpecification8  string `csv:"ProductSpecification8"`
	ProductSpecification9  string `csv:"ProductSpecification9"`
	ProductSpecification10 string `csv:"ProductSpecification10"`
	ProductSpecification11 string `csv:"ProductSpecification11"`
	ProductSpecification12 string `csv:"ProductSpecification12"`
	ProductSpecification13 string `csv:"ProductSpecification13"`
	ProductSpecification14 string `csv:"ProductSpecification14"`
	ProductSpecification15 string `csv:"ProductSpecification15"`
	ProductChanges1        string `csv:"ProductChanges1"`
	ProductChanges2        string `csv:"ProductChanges2"`
	ProductChanges3        string `csv:"ProductChanges3"`
	ProductChanges4        string `csv:"ProductChanges4"`
	ProductChanges5        string `csv:"ProductChanges5"`
	AdditionalDetail1      string `csv:"AdditionalDetail1"`
	AdditionalDetail2      string `csv:"AdditionalDetail2"`
	AdditionalDetail3      string `csv:"AdditionalDetail3"`
	AdditionalDetail4      string `csv:"AdditionalDetail4"`
	AdditionalDetail5      string `csv:"AdditionalDetail5"`
	BrandName              string `csv:"BrandName"`
	ReleaseName            string `csv:"ReleaseName"`
	SegmentNames           string `csv:"SegmentNames"`
	AtsInIndent            string `csv:"AtsInIndent"`
	AtsInInSeason          string `csv:"AtsInInSeason"`
	SalesTip               string `csv:"SalesTip"`
	MarketingSupport       string `csv:"MarketingSupport"`
	CompanyName            string `csv:"CompanyName"`
}

// type Entries struct {
// 	Integration_ID string `db:"Integration_id"`
// 	Size           string `db:"size"`
// 	SKU            string `db:"sku"`
// 	Colour_code    string `db:"colour_code"`
// }

type Verify struct {
	Size        string `db:"size"`
	SKU         string `db:"sku"`
	Colour_code string `db:"colour_code"`
}

type PresentNValidate struct {
	Present bool
}

type Success struct {
	Success  bool   `json:"Success"`
	Message  string `json:"Message"`
	Filepath string `json:"Filepath"`
}

type File_Validation struct {
	File      string `csv:"file"`
	BrandID   string `text:"brand_id"`
	ReleaseID string `text:"release_id"`
	Format    string `text:"format"`
}

// type listTables struct {
// 	Shows []int64 `json:"tables"`
// }