CREATE TABLE wagers (
    id BIGSERIAL PRIMARY KEY,
    total_wager_value INTEGER,
    odds INTEGER,
    selling_percentage INTEGER,
    selling_price DECIMAL,
    current_selling_price DECIMAL,
    percentage_sold DECIMAL,
    amount_sold DECIMAL,
    placed_at TIMESTAMP
);

CREATE TABLE purchases (
    id BIGSERIAL PRIMARY KEY,
    wager_id BIGSERIAL,
    buying_price DECIMAL,
    bought_at TIMESTAMP
);