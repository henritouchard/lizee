package products

import "lizee/pkg/utils"

// Category defines basic informations of product's categories
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CategoryQuery defines query we get for products availability corresponding to category
type CategoryQuery struct {
	CategoryID int    `form:"categoryID"`
	FromDate   string `form:"fromDate"`
	ToDate     string `form:"toDate"`
}

// ListCategories get categories from storage and returns it as serialize-ready
func ListCategories() ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	return utils.RowsToJSON(rows)
}

// CheckAvailabilityByCategory query all products available
// at given date corresponding to this category.
func CheckAvailabilityByCategory(c *CategoryQuery) ([]map[string]interface{}, error) {
	stmt, err := db.Prepare(queryProductAvailabilityByCategory)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(c.FromDate, c.ToDate, c.CategoryID)
	if err != nil {
		return nil, err
	}

	return utils.RowsToJSON(rows)
}
