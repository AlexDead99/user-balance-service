ALTER TABLE accounts DROP CONSTRAINT balance_check;
ALTER TABLE accounts DROP CONSTRAINT owner_check;

ALTER TABLE services DROP CONSTRAINT name_check;

ALTER TABLE products DROP CONSTRAINT product_name_check;
ALTER TABLE products DROP CONSTRAINT price_check;
ALTER TABLE products DROP CONSTRAINT amount_check;

ALTER TABLE transfers DROP CONSTRAINT total_price_check;
ALTER TABLE "ordersDetails" DROP CONSTRAINT order_amount_check;
