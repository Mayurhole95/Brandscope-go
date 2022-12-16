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

type LogID struct {
	Original_file_location string `db:"original_file_location"`
	ReleaseID              string `db:"release_id`
	BrandID                string `db:"brand_id"`
}

// INSERT INTO public.release_uploads(
// id, file_name, no_of_products, no_of_product_releases, no_of_size_breaks, status, error_file, original_file_location, message,release_id, brand_id,  created_at, updated_at)
// VALUES (700, "Testing", 7, 7, 3, 'pending', 'Dash_Summer 21_20221201121220_errors.csv', 'Dash_Summer 21_20221201121220.csv', 'File Validation',  206, 76,  ?, ?);
// 2018-11-30 02:31:35.28
