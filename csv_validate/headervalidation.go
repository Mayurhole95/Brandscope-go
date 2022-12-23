package csv_validate

import (
	"errors"

	"github.com/Mayurhole95/Brandscope-go/utils"
)

func ValidateHeader() (missingHeader [][]string, err error) {
	var map_headers = make(map[string]int)
	headers := []string{"CatalogueOrder", "BrandscopeCarryOver", "Integration_ID", "Barcode", "SKU", "ProductName", "ProductColourCode", "ProductDisplayColour", "GenericColour", "SizeBreak", "AttributeValue", "AttributeType", "AttributeSequence", "DisplayWholesaleRange", "DisplayWholesale", "DisplayRetail", "PackUnits", "AvailableMonths", "AgeGroup", "Gender", "State", "PreOrderLeadTimeDays", "PreOrderMessage"}
	missingHeaders := make([]string, 0)
	missingHeaders2d := make([][]string, 0)
	records := readHeaders()
	utils.ReturnError(err)
	for _, header := range headers {
		for _, colCsv := range records {
			if header == colCsv {
				map_headers[header] = 1
				break
			}

		}

	}

	for _, header := range headers {
		if map_headers[header] != 1 {
			missingHeaders = append(missingHeaders, header)
			missingHeaders2d = append(missingHeaders2d, [][]string{missingHeaders}...)
			missingHeaders = missingHeaders[1:]
		}
	}

	if len(map_headers) < len(headers) {
		err = errors.New(errHeadersMissing)
		return missingHeaders2d, err
	}
	err = nil
	return missingHeaders2d, nil
}
