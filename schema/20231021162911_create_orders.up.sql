-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS orders (
    id serial PRIMARY KEY,
    date text NOT NULL,
    item_id integer NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    count integer NOT NULL,
    price integer NOT NULL
);

-- Index Definition ----------------------------------------------
CREATE INDEX date_item_id_index ON "orders" (date, item_id);