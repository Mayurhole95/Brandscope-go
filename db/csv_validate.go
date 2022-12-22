package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	findID              = `SELECT EXISTS(SELECT 1 FROM brands WHERE id = $1) AND EXISTS(SELECT 1 FROM releases WHERE id = $2)`
	findtables          = `SELECT id FROM brands`
	checkIntegrationID  = `SELECT EXISTS(SELECT 1 FROM size_breaks WHERE brand_id = $1 AND integration_id=$2 AND size=$3) AND EXISTS(SELECT 1 FROM products WHERE brand_id = $1 AND sku=$4 AND colour_code=$5)`
	checkDuplicateEntry = `SELECT sb.Integration_id,sb.size,products.sku,products.colour_code
	FROM size_breaks as sb
	INNER JOIN products ON sb.brand_id = products.brand_id where sb.brand_id=$1 AND sb.product_id=products.id;`
	checkMonths = `SELECT delivery_months FROM releases WHERE id = $1`
	findbyLogID = `SELECT original_file_location, release_id, brand_id
	FROM release_uploads where id=$1;`
)

func (s *store) ListMonths(release_id string) (months []string, err error) {
	rows, err := s.db.Query(checkMonths, release_id)
	ReturnError(err)
	for rows.Next() {
		var row string
		err = rows.Scan(&row)
		row = strings.Trim(row, "{}")
		months = strings.Split(row, ",")
		ReturnError(err)
	}
	ReturnError(err)

	return months, err
}

func (s *store) ListData(brand_id string) (data map[string]Verify, err error) {
	csvDataMap := make(map[string]Verify)
	rows, err := s.db.Query(checkDuplicateEntry, brand_id)
	ReturnError(err)
	for rows.Next() {
		var row entries
		err = rows.Scan(&row.Integration_ID, &row.Size, &row.SKU, &row.Colour_code)
		csvDataMap[row.Integration_ID] = Verify{row.Size, row.SKU, row.Colour_code}
		ReturnError(err)
	}

	return csvDataMap, err
}

func (s *store) FindLogID(ctx context.Context, log_id string) (row LogID, err error) {

	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		rows, err := s.db.QueryContext(ctx, findbyLogID, log_id)
		ReturnError(err)
		for rows.Next() {
			err = rows.Scan(
				&row.Original_file_location, &row.ReleaseID, &row.BrandID,
			)
		}
		return err
	})
	if err == sql.ErrNoRows {
		return row, ErrEmptyData
	}
	return row, nil

}

func (s *store) FindID(ctx context.Context, brand_id string, release_id string) (exists bool, err error) {

	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		rows, err := s.db.QueryContext(ctx, findID, brand_id, release_id)
		ReturnError(err)
		for rows.Next() {
			err = rows.Scan(
				&exists,
			)
		}
		return err
	})
	if err == sql.ErrNoRows {
		return exists, ErrEmptyData
	}
	return

}

func (s *store) FindIntegrationID(brand_id string, integration_id string, size string, sku string, colour_code string) (exists bool, err error) {

	rows, err := s.db.Query(checkIntegrationID)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(
			&exists,
		)
	}
	if exists {
		return true, nil
	} else {
		return false, err
	}
}
func ReturnError(err error) {
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}
}
