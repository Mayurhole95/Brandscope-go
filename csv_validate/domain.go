package csv_validate

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/Mayurhole95/Brandscope-go/db"
	"github.com/Mayurhole95/Brandscope-go/utils"
)

type BrandHeader struct {
	AttributeValue         string `csv:"AttributeValue"`
	AttributeType          string `csv:"AttributeType"`
	AttributeSequence      string `csv:"AttributeSequence"`
	AvailableMonths        string `csv:"AvailableMonths"`
	AtsInIndent            string `csv:"AtsInIndent"`
	AtsInInSeason          string `csv:"AtsInInSeason"`
	AgeGroup               string `csv:"AgeGroup"`
	BrandscopeHierarchy    string `csv:"BrandscopeHierarchy"`
	BrandscopeCarryOver    string `csv:"BrandscopeCarryOver"`
	BrandName              string `csv:"BrandName"`
	Barcode                string `csv:"Barcode"`
	CompanyName            string `csv:"CompanyName"`
	Collections            string `csv:"Collections"`
	Categories             string `csv:"Categories"`
	CatalogueOrder         string `csv:"CatalogueOrder"`
	DisplayWholesaleRange  string `csv:"DisplayWholesaleRange"`
	Divisions              string `csv:"Divisions"`
	DiscountCategory       string `csv:"DiscountCategory"`
	DisplayWholesale       string `csv:"DisplayWholesale"`
	DisplayRetail          string `csv:"DisplayRetail"`
	GenericColour          string `csv:"GenericColour"`
	Gender                 string `csv:"Gender"`
	Integration_ID         string `csv:"Integration_ID"`
	MarketingSupport       string `csv:"MarketingSupport"`
	PreOrderLeadTimeDays   string `csv:"PreOrderLeadTimeDays"`
	PreOrderMessage        string `csv:"PreOrderMessage"`
	ProductMultiple        string `csv:"ProductMultiple"`
	ProductTags            string `csv:"ProductTags"`
	ProductName            string `csv:"ProductName"`
	ProductColourCode      string `csv:"ProductColourCode"`
	ProductDisplayColour   string `csv:"ProductDisplayColour"`
	PackUnits              string `csv:"PackUnits"`
	RetailPriceOriginal    string `csv:"RetailPriceOriginal"`
	RetailPrice            string `csv:"RetailPrice"`
	ReleaseName            string `csv:"ReleaseName"`
	SegmentNames           string `csv:"SegmentNames"`
	SalesTip               string `csv:"SalesTip"`
	SizeBreak              string `csv:"SizeBreak"`
	State                  string `csv:"State"`
	SKU                    string `csv:"SKU"`
	WholesalePriceOriginal string `csv:"WholesalePriceOriginal"`
	WholesalePrice         string `csv:"WholesalePrice"`
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
}

func (b *BrandHeader) ToArray() []string {
	return []string{
		b.CatalogueOrder,
		b.BrandscopeCarryOver,
		b.Integration_ID,
		b.Barcode,
		b.SKU,
		b.ProductName,
		b.ProductColourCode,
		b.ProductDisplayColour,
		b.GenericColour,
		b.SizeBreak,
		b.AttributeValue,
		b.AttributeType,
		b.AttributeSequence,
		b.DisplayWholesale,
		b.DisplayWholesaleRange,
		b.WholesalePriceOriginal,
		b.WholesalePrice,
		b.DisplayRetail,
		b.RetailPriceOriginal,
		b.RetailPrice,
		b.PackUnits,
		b.ProductMultiple,
		b.AvailableMonths,
		b.Divisions,
		b.Collections,
		b.Categories,
		b.DiscountCategory,
		b.ProductTags,
		b.AgeGroup,
		b.Gender,
		b.BrandscopeHierarchy,
		b.State,
		b.PreOrderLeadTimeDays,
		b.PreOrderMessage,
		b.ProductRequirement1,
		b.ProductSpecification1,
		b.ProductSpecification2,
		b.ProductSpecification3,
		b.ProductSpecification4,
		b.ProductSpecification5,
		b.ProductSpecification6,
		b.ProductSpecification7,
		b.ProductSpecification8,
		b.ProductSpecification9,
		b.ProductSpecification10,
		b.ProductSpecification11,
		b.ProductSpecification12,
		b.ProductSpecification13,
		b.ProductSpecification14,
		b.ProductSpecification15,
		b.ProductChanges1,
		b.ProductChanges2,
		b.ProductChanges3,
		b.ProductChanges4,
		b.ProductChanges5,
		b.AdditionalDetail1,
		b.AdditionalDetail2,
		b.AdditionalDetail3,
		b.AdditionalDetail4,
		b.AdditionalDetail5,
		b.SalesTip,
		b.MarketingSupport,
		b.CompanyName,
		b.BrandName,
		b.ReleaseName,
		b.SegmentNames,
		b.AtsInIndent,
		b.AtsInInSeason,
	}
}

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

type LogID struct {
	Original_file_location string `db:"original_file_location"`
	ReleaseID              string `db:"release_id`
	BrandID                string `db:"brand_id"`
}

var logdata db.LogID
var dbMonths []string
var file_name_errors string = "pride_priderelease_20221215114528_errors.csv"

type Service interface {
	Validate(ctx context.Context, id string) (successmessage string, err error)
}

var csvData = []BrandHeader{}

func readHeaders() []string {
	readCsvFile, err := os.Open(logdata.Original_file_location)
	utils.ReturnError(err)
	defer readCsvFile.Close()

	csvReader := csv.NewReader(readCsvFile)
	records, err := csvReader.Read()
	utils.ReturnError(err)
	return records
}
