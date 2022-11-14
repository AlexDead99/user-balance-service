ALTER TABLE accounts ADD CONSTRAINT balance_check CHECK (balance >= 0 AND balance IS NOT NULL);
ALTER TABLE accounts ADD CONSTRAINT owner_check CHECK (owner IS NOT NULL);

ALTER TABLE services ADD CONSTRAINT name_check CHECK (name IS NOT NULL);

ALTER TABLE products ADD CONSTRAINT product_name_check CHECK (name IS NOT NULL);
ALTER TABLE products ADD CONSTRAINT price_check CHECK (price IS NOT NULL AND price > 0);
ALTER TABLE products ADD CONSTRAINT amount_check CHECK (amount IS NOT NULL AND amount >=0);

ALTER TABLE transfers ADD CONSTRAINT total_price_check CHECK (total_price IS NOT NULL AND total_price >=0);
ALTER TABLE "ordersDetails" ADD CONSTRAINT order_amount_check CHECK (amount IS NOT NULL AND amount >=0);
