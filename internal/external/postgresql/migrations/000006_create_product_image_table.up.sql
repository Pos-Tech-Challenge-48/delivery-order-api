-- PRODUCT_IMAGE
-- Tabela que contém dados das imagens dos produtos.

-- TABLE

CREATE TABLE IF NOT EXISTS product_image (
    product_image_id UUID DEFAULT gen_random_uuid()
        constraint product_image_pk primary key,
    product_id UUID
        constraint product_id_fk
            references product (product_id) not null,
    product_image_src_uri varchar not null,
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);


-- TRIGGER

CREATE OR REPLACE TRIGGER tr_product_image_timestamps
    BEFORE UPDATE
    ON public.product_image
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();