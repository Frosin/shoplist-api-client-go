-- name: GetGoodsByShoppingID :many
SELECT * FROM shop_list
WHERE list_id = $1;

-- name: GetShopByName :one
SELECT * FROM shop
WHERE name = $1
LIMIT 1;

-- name: GetShopByID :one
SELECT * FROM shop
WHERE id = $1
LIMIT 1;

-- name: GetLastShopping :one
SELECT * FROM shopping 
ORDER BY rowid DESC LIMIT 1;

-- name: GetComingShoppings :many
SELECT * FROM shopping 
WHERE date >= $1
LIMIT 5;

-- name: GetShoppingByID :one
SELECT * FROM shopping
WHERE id = $1
LIMIT 1;

-- name: GetShoppingDays :many
SELECT date FROM shopping 
WHERE date LIKE $1;

-- name: GetShoppingsByDay :many
SELECT * FROM shopping 
WHERE date=$1;
