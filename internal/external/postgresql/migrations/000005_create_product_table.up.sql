-- PRODUCT
-- Tabela que contém dados dos produtos cadastrados.

-- TABLE

CREATE TABLE IF NOT EXISTS product
(
    product_id UUID DEFAULT gen_random_uuid()
        constraint product_pk
            primary key,
    product_category_id UUID
        constraint product_category_id_fk
            references category (category_id),
    product_name varchar(100) not null,
    product_description varchar,
    product_unitary_price NUMERIC(4, 2) not null,
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- INDEXES

CREATE UNIQUE INDEX IF NOT EXISTS product_name_idx ON product(product_name);
CREATE INDEX IF NOT EXISTS product_category_id_idx ON product(product_category_id);


-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_product_timestamps
    BEFORE UPDATE
    ON public.product
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();