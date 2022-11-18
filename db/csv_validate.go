package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	findID              = `SELECT EXISTS(SELECT 1 FROM brands WHERE id = $1) AND EXISTS(SELECT 1 FROM releases WHERE id = $2)`
	findtables          = `SELECT id FROM brands`
	checkIntegrationID  = `SELECT EXISTS(SELECT 1 FROM size_breaks WHERE brand_id = $1 AND integration_id=$2 AND size=$3) AND EXISTS(SELECT 1 FROM products WHERE brand_id = $1 AND sku=$4 AND colour_code=$5)`
	checkDuplicateEntry = `SELECT sb.Integration_id,sb.size,products.sku,products.colour_code
	FROM size_breaks as sb
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

		for rows.Next() {
			err = rows.Scan(
				&exists,
			)
		}
		return err
	})
	fmt.Println("Exists : ", exists)
	if err == sql.ErrNoRows {
		return exists, ErrBookNotExist
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
	if exists == true {
		return true, nil
	} else {
		return false, err
	}
}

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
}
