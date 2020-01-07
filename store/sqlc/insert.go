package sqlc

import (
	"context"
	"database/sql"
)

const addProductItem = `
INSERT INTO shop_list (
    product_name, quantity, list_id, complete, category_id
) VALUES (
	$1, $2, $3, $4, $5
)`

type AddProductItemParams struct {
	ProductName sql.NullString `json:"product_name"`
	Quantity    sql.NullInt32  `json:"quantity"`
	ListID      sql.NullInt32  `json:"list_id"`
	Complete    sql.NullBool   `json:"complete"`
	CategoryID  sql.NullInt32  `json: "category_id"`
}

func (q *Queries) AddProductItem(ctx context.Context, arg AddProductItemParams) (int64, error) {
	insertResult, err := q.db.ExecContext(ctx, addProductItem, arg.ProductName, arg.Quantity, arg.ListID, arg.Complete, arg.CategoryID)
	if err != nil {
		return 0, err
	}
	insertedID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedID, nil
}

const addShop = `
INSERT INTO shop (
    name
) VALUES (
    $1
)`

func (q *Queries) AddShop(ctx context.Context, name sql.NullString) (int64, error) {
	insertResult, err := q.db.ExecContext(ctx, addShop, name)
	if err != nil {
		return 0, err
	}
	insertedID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedID, nil
}

const addShopping = `-- name: AddShopping :one
INSERT INTO shopping (
    date, shop_id, time, owner_id
) VALUES (
    $1, $2, $3, $4
)
`

type AddShoppingParams struct {
	Date    sql.NullString `json:"date"`
	ShopID  sql.NullInt32  `json:"shop_id"`
	Time    sql.NullString `json:"time"`
	OwnerID sql.NullInt32  `json:"owner_id"`
}

func (q *Queries) AddShopping(ctx context.Context, arg AddShoppingParams) (int64, error) {
	insertResult, err := q.db.ExecContext(ctx, addShopping, arg.Date, arg.ShopID, arg.Time, arg.OwnerID)
	if err != nil {
		return 0, err
	}
	insertedID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedID, nil
}
