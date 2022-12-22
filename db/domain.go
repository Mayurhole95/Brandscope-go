package db

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
