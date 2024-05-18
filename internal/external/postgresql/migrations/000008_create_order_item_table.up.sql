-- ORDER_ITEM
-- Tabela que contém dados dos produtos adicionados em um pedido.

-- TABLE

CREATE TABLE IF NOT EXISTS order_item (
    order_item_id UUID DEFAULT gen_random_uuid()
        constraint order_item_pk primary key,
    order_item_product_id UUID
        constraint order_item_product_id_fk
            references product(product_id),
    order_item_order_id UUID
        constraint order_item_order_id_fk
            references restaurant_order(restaurant_order_id),
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- INDEXES

CREATE INDEX IF NOT EXISTS order_item_order_id_idx ON restaurant_order(restaurant_order_id);


-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_order_item_timestamps
    BEFORE UPDATE
    ON public.order_item
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();