// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const getComingShoppings = `-- name: GetComingShoppings :many
SELECT id, date, sum, shop_id, complete, time, owner_id FROM shopping 
WHERE date >= $1
LIMIT 5
`

func (q *Queries) GetComingShoppings(ctx context.Context, date sql.NullString) ([]Shopping, error) {
	rows, err := q.query(ctx, q.getComingShoppingsStmt, getComingShoppings, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shopping
	for rows.Next() {
		var i Shopping
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Sum,
			&i.ShopID,
			&i.Complete,
			&i.Time,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGoodsByShoppingID = `-- name: GetGoodsByShoppingID :many
SELECT id, product_name, quantity, category_id, complete, list_id FROM shop_list
WHERE list_id = $1
`

func (q *Queries) GetGoodsByShoppingID(ctx context.Context, listID sql.NullInt32) ([]ShopList, error) {
	rows, err := q.query(ctx, q.getGoodsByShoppingIDStmt, getGoodsByShoppingID, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShopList
	for rows.Next() {
		var i ShopList
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Quantity,
			&i.CategoryID,
			&i.Complete,
			&i.ListID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLastShopping = `-- name: GetLastShopping :one
SELECT id, date, sum, shop_id, complete, time, owner_id FROM shopping 
ORDER BY rowid DESC LIMIT 1
`

func (q *Queries) GetLastShopping(ctx context.Context) (Shopping, error) {
	row := q.queryRow(ctx, q.getLastShoppingStmt, getLastShopping)
	var i Shopping
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Sum,
		&i.ShopID,
		&i.Complete,
		&i.Time,
		&i.OwnerID,
	)
	return i, err
}

const getShopByID = `-- name: GetShopByID :one
SELECT id, name FROM shop
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetShopByID(ctx context.Context, id int32) (Shop, error) {
	row := q.queryRow(ctx, q.getShopByIDStmt, getShopByID, id)
	var i Shop
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getShopByName = `-- name: GetShopByName :one
SELECT id, name FROM shop
WHERE name = $1
LIMIT 1
`

func (q *Queries) GetShopByName(ctx context.Context, name sql.NullString) (Shop, error) {
	row := q.queryRow(ctx, q.getShopByNameStmt, getShopByName, name)
	var i Shop
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getShoppingByID = `-- name: GetShoppingByID :one
SELECT id, date, sum, shop_id, complete, time, owner_id FROM shopping
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetShoppingByID(ctx context.Context, id int32) (Shopping, error) {
	row := q.queryRow(ctx, q.getShoppingByIDStmt, getShoppingByID, id)
	var i Shopping
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Sum,
		&i.ShopID,
		&i.Complete,
		&i.Time,
		&i.OwnerID,
	)
	return i, err
}
