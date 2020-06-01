CREATE TABLE IF NOT EXISTS category (
   id INT primary key generated always as identity,
   name TEXT NOT NULL
);
COMMENT ON TABLE category IS 'List all existing categories of products.';

CREATE TABLE IF NOT EXISTS product (
   id INT primary key generated always as identity,
   name TEXT NOT NULL, 
   description TEXT NOT NULL,
   cstr_category INT references category(id) NOT NULL,
   stock INT NOT NULL,
   picture TEXT NOT NULL
);
COMMENT ON TABLE product IS 'List all existing products.';

CREATE TABLE IF NOT EXISTS rental_order (
   id INT primary key generated always as identity,
   start_date DATE NOT NULL,
   end_date  DATE NOT NULL,
   returned BOOLEAN NOT NULL,
   product INT references product(id) NOT NULL
);
COMMENT ON TABLE rental_order IS 'List all orders.';


-- INSERT INTO category (name) VALUES ('tents'),('shoes'),('dish'),('sleeping bags');

INSERT INTO product (name, description, cstr_category, stock, picture) VALUES
('tente trekking UL3', 'Here is a little resumé about this tent', 1, 4, 'https://lizee-public-files-prod.s3.eu-west-3.amazonaws.com/variants/iuGMoEwWANTXVnuwvEvbh5Bb/da4eb9a43fb7f2c8ba87db5a50ee1aefbb18c69b1542105578e75baeecca46aa'),
('tente quickhiker UL 2', 'Here is a little resumé about this tent', 1, 10, 'https://lizee-public-files-prod.s3.eu-west-3.amazonaws.com/variants/RTuZhYRZbf89ZG5Lrf72sg4y/da4eb9a43fb7f2c8ba87db5a50ee1aefbb18c69b1542105578e75baeecca46aa'),
('tente trekking UL1', 'Here is a little resumé about this tent', 1, 4, 'https://lizee-public-files-prod.s3.eu-west-3.amazonaws.com/variants/iuGMoEwWANTXVnuwvEvbh5Bb/da4eb9a43fb7f2c8ba87db5a50ee1aefbb18c69b1542105578e75baeecca46aa'),
('SDC trek 900 0° plume', 'Here is a little resumé about this sleeping bag', 4, 6, 'https://lizee-public-files-prod.s3.eu-west-3.amazonaws.com/variants/GmB8XhAq5cdDnBwtP1fvWHg9/da4eb9a43fb7f2c8ba87db5a50ee1aefbb18c69b1542105578e75baeecca46aa'),
('drap de sac soie', 'Here is a little resumé about this meat bag', 1, 4, 'https://lizee-public-files-prod.s3.eu-west-3.amazonaws.com/variants/MdCVTjyDNTwJmc7VZT3grc8g/da4eb9a43fb7f2c8ba87db5a50ee1aefbb18c69b1542105578e75baeecca46aa');


-- drop table rental_order ;
-- drop table product ;
-- drop table category ;