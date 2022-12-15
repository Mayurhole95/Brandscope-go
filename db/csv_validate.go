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
<<<<<<< HEAD
	INNER JOIN products ON sb.brand_id = products.brand_id where sb.brand_id=$1 AND sb.product_id=products.id;
`

	// csvheaders = `SELECT `
)

type entries struct {
	Integration_ID string `db:"Integration_id"`
	Size           string `db:"size"`
	SKU            string `db:"sku"`
	Colour_code    string `db:"colour_code"`
}

type Verify struct {
	Size        string `db:"size"`
	SKU         string `db:"sku"`
	Colour_code string `db:"colour_code"`
}

func (s *store) ListData(id string) (data map[string]Verify, err error) {
	m := make(map[string]Verify)
	rows, err := s.db.Query(checkDuplicateEntry, id)
	if err != nil {
		fmt.Println("Error occured here : ", err.Error())
	}
	for rows.Next() {
		var r entries
		err = rows.Scan(&r.Integration_ID, &r.Size, &r.SKU, &r.Colour_code)
		// fmt.Println(r)
		m[r.Integration_ID] = Verify{r.Size, r.SKU, r.Colour_code}
		if err != nil {
			fmt.Println(err)
		}
		// data = append(data, r)
	}
	fmt.Println(m)

	return m, err
}

func (s *store) FindID(ctx context.Context, id string, id2 string) (exists bool, err error) {

	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		rows, err := s.db.QueryContext(ctx, findID, id, id2)
		if err != nil {
			return err
		}

=======
	INNER JOIN products ON sb.brand_id = products.brand_id where sb.brand_id=$1 AND sb.product_id=products.id;`
	checkMonths = `SELECT delivery_months FROM releases WHERE id = $1`
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

func (s *store) FindID(ctx context.Context, brand_id string, release_id string) (exists bool, err error) {

	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		rows, err := s.db.QueryContext(ctx, findID, brand_id, release_id)
		ReturnError(err)
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
		for rows.Next() {
			err = rows.Scan(
				&exists,
			)
		}
		return err
	})
	fmt.Println("Exists : ", exists)
	if err == sql.ErrNoRows {
<<<<<<< HEAD
		return exists, ErrBookNotExist
=======
		return exists, ErrEmptyData
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
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
<<<<<<< HEAD
	if exists == true {
=======
	if exists {
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
		return true, nil
	} else {
		return false, err
	}
}
<<<<<<< HEAD

func (s *store) ShowTables(ctx context.Context) (ids []int64, err error) {

	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		rows, err := s.db.QueryContext(ctx, findtables)
		if err != nil {
			return err
		}

		ids = make([]int64, 0)
		for rows.Next() {
			var id int64
			err := rows.Scan(
				&id,
			)

			if err != nil {
				return err
			}
			fmt.Println("sasasa : ", id)
			ids = append(ids, id)
		}
		return err
	})

	if err == sql.ErrNoRows {
		return ids, ErrBookNotExist
	}
	return
=======
func ReturnError(err error) {
	if err != nil {
		fmt.Println("Error Occured :", err.Error())
		return
	}
>>>>>>> a7c0e98be135ce77bedc3d6a295fc97c0a9ad2e8
}
