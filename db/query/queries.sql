-- name: SelectItems :many
SELECT * FROM items
ORDER BY id;

-- name: InsertOrders :exec
INSERT INTO orders (date, item_id, count, price) 
VALUES (
    $1,
    $2,
    $3,
    $4
);