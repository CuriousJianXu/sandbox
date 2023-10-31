-- name: SelectItems :many
SELECT * FROM items
ORDER BY id;

-- name: SelectOrdersByItemIDAndDate :many
SELECT * FROM orders
WHERE item_id=$1
  AND date=$2;


-- name: InsertOrders :exec
INSERT INTO orders (date, item_id, count, price) 
VALUES (
    $1,
    $2,
    $3,
    $4
);